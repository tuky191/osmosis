package mocks

import (
	"cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/osmosis-labs/osmosis/osmomath"
	"github.com/osmosis-labs/osmosis/v20/ingest/sqs/domain"
	poolmanagertypes "github.com/osmosis-labs/osmosis/v20/x/poolmanager/types"
)

type MockRoutablePool struct {
	ChainPoolModel       poolmanagertypes.PoolI
	TickModel            *domain.TickModel
	ID                   uint64
	Balances             sdk.Coins
	Denoms               []string
	TotalValueLockedUSDC osmomath.Int
	PoolType             poolmanagertypes.PoolType
	TokenOutDenom        string
	TakerFee             osmomath.Dec
	SpreadFactor         osmomath.Dec
}

var DefaultSpreadFactor = osmomath.MustNewDecFromStr("0.005")

// MarshalJSON implements domain.RoutablePool.
func (*MockRoutablePool) MarshalJSON() ([]byte, error) {
	panic("unimplemented")
}

// UnmarshalJSON implements domain.RoutablePool.
func (*MockRoutablePool) UnmarshalJSON([]byte) error {
	panic("unimplemented")
}

var (
	_ domain.PoolI        = &MockRoutablePool{}
	_ domain.RoutablePool = &MockRoutablePool{}
)

// GetUnderlyingPool implements routerusecase.RoutablePool.
func (mp *MockRoutablePool) GetUnderlyingPool() poolmanagertypes.PoolI {
	return mp.ChainPoolModel
}

// GetSQSPoolModel implements domain.PoolI.
func (mp *MockRoutablePool) GetSQSPoolModel() domain.SQSPool {
	return domain.SQSPool{
		Balances:             mp.Balances,
		TotalValueLockedUSDC: mp.TotalValueLockedUSDC,
		SpreadFactor:         DefaultSpreadFactor,
	}
}

// CalculateTokenOutByTokenIn implements routerusecase.RoutablePool.
func (*MockRoutablePool) CalculateTokenOutByTokenIn(tokenIn sdk.Coin) (sdk.Coin, error) {
	panic("unimplemented")
}

// String implements domain.RoutablePool.
func (*MockRoutablePool) String() string {
	panic("unimplemented")
}

// GetTickModel implements domain.RoutablePool.
func (mp *MockRoutablePool) GetTickModel() (*domain.TickModel, error) {
	return mp.TickModel, nil
}

// Validate implements domain.PoolI.
func (*MockRoutablePool) Validate(minUOSMOTVL math.Int) error {
	// Note: always valid for tests.
	return nil
}

// GetTokenOutDenom implements routerusecase.RoutablePool.
func (mp *MockRoutablePool) GetTokenOutDenom() string {
	return mp.TokenOutDenom
}

// ChargeTakerFee implements domain.RoutablePool.
func (*MockRoutablePool) ChargeTakerFeeExactIn(tokenIn sdk.Coin) (tokenInAfterFee sdk.Coin) {
	panic("unimplemented")
}

// GetTakerFee implements domain.PoolI.
func (mp *MockRoutablePool) GetTakerFee() math.LegacyDec {
	return mp.TakerFee
}

var _ domain.PoolI = &MockRoutablePool{}
var _ domain.RoutablePool = &MockRoutablePool{}

// GetId implements domain.PoolI.
func (mp *MockRoutablePool) GetId() uint64 {
	return mp.ID
}

// GetPoolDenoms implements domain.PoolI.
func (mp *MockRoutablePool) GetPoolDenoms() []string {
	return mp.Denoms
}

// GetTotalValueLockedUOSMO implements domain.PoolI.
func (mp *MockRoutablePool) GetTotalValueLockedUOSMO() math.Int {
	return mp.TotalValueLockedUSDC
}

// GetType implements domain.PoolI.
func (mp *MockRoutablePool) GetType() poolmanagertypes.PoolType {
	return mp.PoolType
}

func deepCopyPool(mp *MockRoutablePool) *MockRoutablePool {
	newDenoms := make([]string, len(mp.Denoms))
	copy(newDenoms, mp.Denoms)

	newTotalValueLocker := osmomath.NewIntFromBigInt(mp.TotalValueLockedUSDC.BigInt())

	newBalances := sdk.NewCoins(mp.Balances...)

	return &MockRoutablePool{
		ID:                   mp.ID,
		Denoms:               newDenoms,
		TotalValueLockedUSDC: newTotalValueLocker,
		PoolType:             mp.PoolType,

		// Note these are not deep copied.
		ChainPoolModel: mp.ChainPoolModel,
		TokenOutDenom:  mp.TokenOutDenom,
		Balances:       newBalances,
		TakerFee:       mp.TakerFee.Clone(),
		SpreadFactor:   mp.SpreadFactor.Clone(),
	}
}

func WithPoolID(mockPool *MockRoutablePool, id uint64) *MockRoutablePool {
	newPool := deepCopyPool(mockPool)
	newPool.ID = id
	return newPool
}

func WithDenoms(mockPool *MockRoutablePool, denoms []string) *MockRoutablePool {
	newPool := deepCopyPool(mockPool)
	newPool.Denoms = denoms
	return newPool
}

func WithTokenOutDenom(mockPool *MockRoutablePool, tokenOutDenom string) *MockRoutablePool {
	newPool := deepCopyPool(mockPool)
	newPool.TokenOutDenom = tokenOutDenom
	return newPool
}

func WithChainPoolModel(mockPool *MockRoutablePool, chainPool poolmanagertypes.PoolI) *MockRoutablePool {
	newPool := deepCopyPool(mockPool)
	newPool.ChainPoolModel = chainPool
	return newPool
}

func WithTakerFee(mockPool *MockRoutablePool, takerFee osmomath.Dec) *MockRoutablePool {
	newPool := deepCopyPool(mockPool)
	newPool.TakerFee = takerFee
	return newPool
}