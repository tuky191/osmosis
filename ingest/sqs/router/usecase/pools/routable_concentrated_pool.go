package pools

import (
	"encoding/json"
	"fmt"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	// "github.com/gogo/protobuf/codec"

	"github.com/osmosis-labs/osmosis/osmomath"
	"github.com/osmosis-labs/osmosis/v20/ingest/sqs/domain"
	clmath "github.com/osmosis-labs/osmosis/v20/x/concentrated-liquidity/math"
	concentratedmodel "github.com/osmosis-labs/osmosis/v20/x/concentrated-liquidity/model"
	"github.com/osmosis-labs/osmosis/v20/x/concentrated-liquidity/swapstrategy"
	"github.com/osmosis-labs/osmosis/v20/x/poolmanager"
	poolmanagertypes "github.com/osmosis-labs/osmosis/v20/x/poolmanager/types"
)

var _ domain.RoutablePool = &routableConcentratedPoolImpl{}

type routableConcentratedPoolImpl struct {
	// codec codec.Codec `json:"-"`
	domain.PoolI
	TokenOutDenom string       "json:\"token_out_denom\""
	TakerFee      osmomath.Dec "json:\"taker_fee\""
}

// GetTakerFee implements domain.RoutablePool.
func (r *routableConcentratedPoolImpl) GetTakerFee() math.LegacyDec {
	return r.TakerFee
}

// CalculateTokenOutByTokenIn implements domain.RoutablePool.
// It calculates the amount of token out given the amount of token in for a concentrated liquidity pool.
// Fails if:
// - the underlying chain pool set on the routable pool is not of concentrated type
// - fails to retrieve the tick model for the pool
// - the current tick is not within the specified current bucket range
// - tick model has no liquidity flag set
// - the current sqrt price is zero
// - rans out of ticks during swap (token in is too high for liquidity in the pool)
func (r *routableConcentratedPoolImpl) CalculateTokenOutByTokenIn(tokenIn sdk.Coin) (sdk.Coin, error) {
	poolType := r.GetType()

	// Esnure that the pool is concentrated
	if poolType != poolmanagertypes.Concentrated {
		return sdk.Coin{}, domain.InvalidPoolTypeError{PoolType: int32(poolType)}
	}

	chainPool := r.GetUnderlyingPool()
	// Defense in depth casting check to confirm that the pool is concentrated
	concentratedPool, ok := chainPool.(*concentratedmodel.Pool)
	if !ok {
		return sdk.Coin{}, fmt.Errorf("failed to cast pool (%d) to concentrated pool", r.GetId())
	}

	tickModel, err := r.GetTickModel()
	if err != nil {
		return sdk.Coin{}, err
	}

	// Ensure pool has liquidity.
	if tickModel.HasNoLiquidity {
		return sdk.Coin{}, domain.ConcentratedNoLiquidityError{
			PoolId: r.GetId(),
		}
	}

	// Ensure that the current bucket is within the available bucket range.
	currentBucketIndex := tickModel.CurrentTickIndex

	if currentBucketIndex < 0 || currentBucketIndex >= int64(len(tickModel.Ticks)) {
		return sdk.Coin{}, domain.ConcentratedCurrentTickNotWithinBucketError{
			PoolId:             r.GetId(),
			CurrentBucketIndex: currentBucketIndex,
			TotalBuckets:       int64(len(tickModel.Ticks)),
		}
	}

	currentBucket := tickModel.Ticks[currentBucketIndex]

	isCurrentTickWithinBucket := concentratedPool.IsCurrentTickInRange(currentBucket.LowerTick, currentBucket.UpperTick)
	if !isCurrentTickWithinBucket {
		return sdk.Coin{}, domain.ConcentratedCurrentTickAndBucketMismatchError{
			CurrentTick: concentratedPool.CurrentTick,
			LowerTick:   currentBucket.LowerTick,
			UpperTick:   currentBucket.UpperTick,
		}
	}

	// Set the appropriate token out denom.
	isZeroForOne := tokenIn.Denom == concentratedPool.Token0
	tokenOutDenom := concentratedPool.Token0
	if isZeroForOne {
		tokenOutDenom = concentratedPool.Token1
	}

	// Initialize the swap strategy.
	swapStrategy := swapstrategy.New(isZeroForOne, osmomath.ZeroBigDec(), &sdk.KVStoreKey{}, concentratedPool.SpreadFactor)

	var (
		// Swap state
		currentSqrtPrice = concentratedPool.GetCurrentSqrtPrice()

		amountRemainingIn = tokenIn.Amount.ToLegacyDec()
		amountOutTotal    = osmomath.ZeroDec()
	)

	if currentSqrtPrice.IsZero() {
		return sdk.Coin{}, domain.ConcentratedZeroCurrentSqrtPriceError{
			PoolId: r.GetId(),
		}
	}

	// Compute swap over all buckets.
	for amountRemainingIn.GT(osmomath.ZeroDec()) {
		if currentBucketIndex >= int64(len(tickModel.Ticks)) {
			// This happens when there is not enough liquidity in the pool to complete the swap
			// for a given amount of token in.
			return sdk.Coin{}, domain.ConcentratedNotEnoughLiquidityToCompleteSwapError{
				PoolId:   r.GetId(),
				AmountIn: sdk.NewCoins(tokenIn).String(),
			}
		}

		currentBucket = tickModel.Ticks[currentBucketIndex]

		// Compute the next initialized tick index depending on the swap direction.
		// Zero for one - in the lower tick direction.
		// One for zero - in the upper tick direction.
		var nextInitializedTickIndex int64
		if isZeroForOne {
			nextInitializedTickIndex = currentBucket.LowerTick
			currentBucketIndex--
		} else {
			nextInitializedTickIndex = currentBucket.UpperTick
			currentBucketIndex++
		}

		// Get the sqrt price for the next initialized tick index.
		sqrtPriceTarget, err := clmath.TickToSqrtPrice(nextInitializedTickIndex)
		if err != nil {
			return sdk.Coin{}, err
		}

		// Compute the swap within current bucket
		sqrtPriceNext, amountInConsumed, amountOutComputed, spreadRewardChargeTotal := swapStrategy.ComputeSwapWithinBucketOutGivenIn(currentSqrtPrice, sqrtPriceTarget, currentBucket.LiquidityAmount, amountRemainingIn)

		// Update swap state for next iteration
		amountRemainingIn = amountRemainingIn.SubMut(amountInConsumed).SubMut(spreadRewardChargeTotal)
		amountOutTotal = amountOutTotal.AddMut(amountOutComputed)

		// Logs
		r.emitSwapDebugLogs(currentSqrtPrice, sqrtPriceNext, currentBucket.LiquidityAmount, amountInConsumed, amountOutComputed, spreadRewardChargeTotal)

		// Update current sqrt price
		currentSqrtPrice = sqrtPriceNext
	}

	// Return the total amount out.
	return sdk.NewCoin(tokenOutDenom, amountOutTotal.TruncateInt()), nil
}

// GetTokenOutDenom implements RoutablePool.
func (r *routableConcentratedPoolImpl) GetTokenOutDenom() string {
	return r.TokenOutDenom
}

// String implements domain.RoutablePool.
func (r *routableConcentratedPoolImpl) String() string {
	return fmt.Sprintf("pool (%d), pool type (%d), pool denoms (%v)", r.PoolI.GetId(), r.PoolI.GetType(), r.PoolI.GetPoolDenoms())
}

// ChargeTakerFee implements domain.RoutablePool.
// Charges the taker fee for the given token in and returns the token in after the fee has been charged.
func (r *routableConcentratedPoolImpl) ChargeTakerFeeExactIn(tokenIn sdk.Coin) (tokenInAfterFee sdk.Coin) {
	tokenInAfterTakerFee, _ := poolmanager.CalcTakerFeeExactIn(tokenIn, r.GetTakerFee())
	return tokenInAfterTakerFee
}

// MarshalJSON implements domain.RoutablePool.
func (r *routableConcentratedPoolImpl) MarshalJSON() ([]byte, error) {
	var (
		routablePoolSerialized routableSerializedPool
		err                    error
	)
	bz, err := json.Marshal(r.PoolI)
	if err != nil {
		return nil, err
	}

	routablePoolSerialized.PoolData = bz
	routablePoolSerialized.TokenOutDenom = r.TokenOutDenom
	routablePoolSerialized.TakerFee = r.TakerFee

	return json.Marshal(routablePoolSerialized)
}

// UnmarshalJSON implements domain.RoutablePool.
func (r *routableConcentratedPoolImpl) UnmarshalJSON(data []byte) error {
	var routablePoolSerialized routableSerializedPool
	err := json.Unmarshal(data, &routablePoolSerialized)
	if err != nil {
		return err
	}

	r.PoolI = &domain.PoolWrapper{}

	err = json.Unmarshal(routablePoolSerialized.PoolData, r.PoolI)
	if err != nil {
		return err
	}

	r.TokenOutDenom = routablePoolSerialized.TokenOutDenom
	r.TakerFee = routablePoolSerialized.TakerFee

	return nil
}

// TODO: switch to proper logging
func (r *routableConcentratedPoolImpl) emitSwapDebugLogs(currentSqrtPrice, reachedPrice osmomath.BigDec, liquidity osmomath.Dec, amountIn, amountOut, spreadCharge osmomath.Dec) {
	fmt.Println("start sqrt price", currentSqrtPrice)
	fmt.Println("reached sqrt price", reachedPrice)
	fmt.Println("liquidity", liquidity)
	fmt.Println("amountIn", amountIn)
	fmt.Println("amountOut", amountOut)
	fmt.Println("spreadRewardChargeTotal", spreadCharge)
}