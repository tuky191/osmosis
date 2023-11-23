package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/osmosis-labs/osmosis/osmoutils"
)

type SwapAmountInRoutes []SwapAmountInRoute

func (routes SwapAmountInRoutes) Validate() error {
	if len(routes) == 0 {
		return ErrEmptyRoutes
	}

	for _, route := range routes {
		err := sdk.ValidateDenom(route.TokenOutDenom)
		if err != nil {
			return err
		}
	}

	return nil
}

func (routes SwapAmountInRoutes) IntermediateDenoms() []string {
	if len(routes) < 2 {
		return nil
	}
	intermediateDenoms := make([]string, 0, len(routes)-1)
	for _, route := range routes[:len(routes)-1] {
		intermediateDenoms = append(intermediateDenoms, route.TokenOutDenom)
	}

	return intermediateDenoms
}

func (routes SwapAmountInRoutes) PoolIds() []uint64 {
	poolIds := make([]uint64, 0, len(routes))
	for _, route := range routes {
		poolIds = append(poolIds, route.PoolId)
	}
	return poolIds
}

func (routes SwapAmountInRoutes) Length() int {
	return len(routes)
}

type SwapAmountOutRoutes []SwapAmountOutRoute

func (routes SwapAmountOutRoutes) Validate() error {
	if len(routes) == 0 {
		return ErrEmptyRoutes
	}

	for _, route := range routes {
		err := sdk.ValidateDenom(route.TokenInDenom)
		if err != nil {
			return err
		}
	}

	return nil
}

func (routes SwapAmountOutRoutes) IntermediateDenoms() []string {
	if len(routes) < 2 {
		return nil
	}
	intermediateDenoms := make([]string, 0, len(routes)-1)
	for _, route := range routes[1:] {
		intermediateDenoms = append(intermediateDenoms, route.TokenInDenom)
	}

	return intermediateDenoms
}

func (routes SwapAmountOutRoutes) PoolIds() []uint64 {
	poolIds := make([]uint64, 0, len(routes))
	for _, route := range routes {
		poolIds = append(poolIds, route.PoolId)
	}
	return poolIds
}

func (routes SwapAmountOutRoutes) Length() int {
	return len(routes)
}

// ValidateSwapAmountInSplitRoute validates a slice of SwapAmountInSplitRoute.
//
// returns an error if any of the following are true:
// - the slice is empty
// - any SwapAmountInRoute in the slice is invalid
// - the last TokenOutDenom of any SwapAmountInRoute in the slice does not match the TokenOutDenom of the previous SwapAmountInRoute in the slice
// - there are duplicate SwapAmountInRoutes in the slice
func ValidateSwapAmountInSplitRoute(splitRoutes []SwapAmountInSplitRoute) error {
	if len(splitRoutes) == 0 {
		return ErrEmptyRoutes
	}

	// validate every multihop path
	previousLastDenomOut := ""
	multihopRoutes := make([]SwapAmountInRoutes, 0, len(splitRoutes))
	for _, splitRoute := range splitRoutes {
		multihopRoute := splitRoute.Pools

		err := SwapAmountInRoutes(multihopRoute).Validate()
		if err != nil {
			return err
		}

		lastDenomOut := multihopRoute[len(multihopRoute)-1].TokenOutDenom

		if previousLastDenomOut != "" && lastDenomOut != previousLastDenomOut {
			return InvalidFinalTokenOutError{TokenOutGivenA: previousLastDenomOut, TokenOutGivenB: lastDenomOut}
		}

		previousLastDenomOut = lastDenomOut

		multihopRoutes = append(multihopRoutes, multihopRoute)
	}

	if osmoutils.ContainsDuplicateDeepEqual(multihopRoutes) {
		return ErrDuplicateRoutesNotAllowed
	}

	return nil
}

// ValidateSwapAmountOutSplitRoute validates a slice of SwapAmountOutSplitRoute and returns an error if any of the following are true:
// - the slice is empty
// - any SwapAmountOutRoute in the slice is invalid
// - the first TokenInDenom of any SwapAmountOutRoute in the slice does not match the TokenInDenom of the previous SwapAmountOutRoute in the slice
// - there are duplicate SwapAmountOutRoutes in the slice
func ValidateSwapAmountOutSplitRoute(splitRoutes []SwapAmountOutSplitRoute) error {
	if len(splitRoutes) == 0 {
		return ErrEmptyRoutes
	}

	// validate every multihop path
	previousFirstDenomIn := ""
	multihopRoutes := make([]SwapAmountOutRoutes, 0, len(splitRoutes))
	for _, splitRoute := range splitRoutes {
		multihopRoute := splitRoute.Pools

		err := SwapAmountOutRoutes(multihopRoute).Validate()
		if err != nil {
			return err
		}

		firstDenomIn := multihopRoute[0].TokenInDenom

		if previousFirstDenomIn != "" && firstDenomIn != previousFirstDenomIn {
			return InvalidFinalTokenOutError{TokenOutGivenA: previousFirstDenomIn, TokenOutGivenB: firstDenomIn}
		}

		previousFirstDenomIn = firstDenomIn

		multihopRoutes = append(multihopRoutes, multihopRoute)
	}

	if osmoutils.ContainsDuplicateDeepEqual(multihopRoutes) {
		return ErrDuplicateRoutesNotAllowed
	}

	return nil
}

// AddEdge adds the given pool id and token as a route in the end entry.
// Creates start entry or end entry if they do not exist
func (g *RoutingGraph) AddEdge(start, end, token string, poolID uint64) {
	if g.Entries == nil {
		g.Entries = make([]*RoutingGraphEntry, 0)
	}

	var startEntry *Inner
	for _, entry := range g.Entries {
		if entry.Key == start {
			startEntry = entry.Value
			break
		}
	}

	// if we did not have any startEntry, initialize start entry
	if startEntry == nil {
		startEntry = &Inner{Entries: make([]*InnerMapEntry, 0)}
		g.Entries = append(g.Entries, &RoutingGraphEntry{Key: start, Value: startEntry})
	}

	var endEntry *Routes
	for _, entry := range startEntry.Entries {
		if entry.Key == end {
			endEntry = entry.Value
			break
		}
	}

	// if we did not have any endEntry, initialize end entry
	if endEntry == nil {
		endEntry = &Routes{Routes: make([]*Route, 0)}
		startEntry.Entries = append(startEntry.Entries, &InnerMapEntry{Key: end, Value: endEntry})
	}

	endEntry.Routes = append(endEntry.Routes, &Route{PoolId: poolID, Token: token})
}
