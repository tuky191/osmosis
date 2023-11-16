package usecase_test

import (
	"fmt"

	"github.com/osmosis-labs/osmosis/osmomath"
	"github.com/osmosis-labs/osmosis/v20/ingest/sqs/domain"
	"github.com/osmosis-labs/osmosis/v20/ingest/sqs/domain/mocks"
	routerusecase "github.com/osmosis-labs/osmosis/v20/ingest/sqs/router/usecase"
	"github.com/osmosis-labs/osmosis/v20/ingest/sqs/router/usecase/route"
	poolmanagertypes "github.com/osmosis-labs/osmosis/v20/x/poolmanager/types"
)

type routesTestCase struct {
	pools []domain.PoolI

	maxHops   int
	maxRoutes int

	tokenInDenom           string
	tokenOutDenom          string
	currentRoute           domain.Route
	poolsUsed              []bool
	previousTokenOutDenoms []string

	expectedRoutes []domain.Route
	expectedError  error
}

// Tests that find routes is a greedy algorithm where it does not prioritize the best route
// in terms of the number of hops. It prioritizes the first route that it finds via DFS.
func (s *RouterTestSuite) TestFindRoutes() {
	denomOne := DenomOne
	denomTwo := DenomTwo

	defaultPool := &mocks.MockRoutablePool{
		ID:                   1,
		Denoms:               []string{denomOne, denomTwo},
		TotalValueLockedUSDC: osmomath.NewInt(10),
		PoolType:             poolmanagertypes.Balancer,
		TakerFee:             osmomath.ZeroDec(),
		SpreadFactor:         osmomath.ZeroDec(),
	}

	tests := map[string]routesTestCase{
		"no pools -> no routes": {
			pools: []domain.PoolI{},

			tokenInDenom:   denomOne,
			tokenOutDenom:  denomTwo,
			currentRoute:   &route.RouteImpl{},
			poolsUsed:      []bool{},
			expectedRoutes: []domain.Route{},
		},
		"one pool; tokens in & out match -> route created": {
			pools: []domain.PoolI{
				defaultPool,
			},

			maxHops:   1,
			maxRoutes: 1,

			tokenInDenom:  denomOne,
			tokenOutDenom: denomTwo,
			currentRoute:  &route.RouteImpl{},
			poolsUsed:     []bool{false},
			expectedRoutes: []domain.Route{
				WithRoutePools(&route.RouteImpl{}, []domain.RoutablePool{mocks.WithTokenOutDenom(defaultPool, denomTwo)}),
			},
		},
		"one pool; tokens in & out match but max hops stops route from being found": {
			pools: []domain.PoolI{
				defaultPool,
			},

			maxHops:   0,
			maxRoutes: 3,

			tokenInDenom:   denomOne,
			tokenOutDenom:  denomTwo,
			currentRoute:   &route.RouteImpl{},
			poolsUsed:      []bool{false},
			expectedRoutes: []domain.Route{},
		},
		"one pool; tokens in & out match but max router stops route from being found": {
			pools: []domain.PoolI{
				defaultPool,
			},

			maxHops:   3,
			maxRoutes: 0,

			tokenInDenom:   denomOne,
			tokenOutDenom:  denomTwo,
			currentRoute:   &route.RouteImpl{},
			poolsUsed:      []bool{false},
			expectedRoutes: []domain.Route{},
		},
		"one pool; token out does not match -> no route": {
			pools: []domain.PoolI{
				defaultPool,
			},

			maxHops:   1,
			maxRoutes: 1,

			tokenInDenom:   denomOne,
			tokenOutDenom:  DenomThree,
			currentRoute:   &route.RouteImpl{},
			poolsUsed:      []bool{false},
			expectedRoutes: []domain.Route{},
		},
		"one pool; token in does not match -> no route": {
			pools: []domain.PoolI{
				defaultPool,
			},

			maxHops:   1,
			maxRoutes: 1,

			tokenInDenom:   DenomThree,
			tokenOutDenom:  denomTwo,
			currentRoute:   &route.RouteImpl{},
			poolsUsed:      []bool{false},
			expectedRoutes: []domain.Route{},
		},
		"two pools; valid 2 hop route": {
			pools: []domain.PoolI{
				defaultPool,
				mocks.WithDenoms(defaultPool, []string{denomTwo, DenomThree}),
			},

			maxHops:   2,
			maxRoutes: 1,

			tokenInDenom:  denomOne,
			tokenOutDenom: DenomThree,
			currentRoute:  &route.RouteImpl{},
			poolsUsed:     []bool{false, false},
			expectedRoutes: []domain.Route{
				WithRoutePools(&route.RouteImpl{}, []domain.RoutablePool{
					mocks.WithTokenOutDenom(defaultPool, denomTwo),
					mocks.WithTokenOutDenom(mocks.WithDenoms(defaultPool, []string{denomTwo, DenomThree}), DenomThree),
				}),
			},
		},
		"two pools; max hops of one does not let route to be found": {
			pools: []domain.PoolI{
				defaultPool,
				mocks.WithDenoms(defaultPool, []string{denomTwo, DenomThree}),
			},

			maxHops:   1,
			maxRoutes: 1,

			tokenInDenom:   denomOne,
			tokenOutDenom:  DenomThree,
			currentRoute:   &route.RouteImpl{},
			poolsUsed:      []bool{false, false},
			expectedRoutes: []domain.Route{},
		},
		"4 pools; valid 4 hop route (not in order)": {
			pools: []domain.PoolI{
				defaultPool, // A: denom 1, 2
				mocks.WithDenoms(defaultPool, []string{denomTwo, DenomThree}), // B: denom 2, 3
				mocks.WithDenoms(defaultPool, []string{DenomFour, denomOne}),  // C: denom 4, 1
				mocks.WithDenoms(defaultPool, []string{DenomFour, DenomFive}), // D: denom 4, 5
			},

			maxHops:   4,
			maxRoutes: 1,

			// D (denom5 for denom4) -> C (denom4 for denom1) -> A (denom1 for denom2) -> B (denom2 for denom3)
			tokenInDenom:  DenomFive,
			tokenOutDenom: DenomThree,
			currentRoute:  &route.RouteImpl{},
			poolsUsed:     []bool{false, false, false, false},
			expectedRoutes: []domain.Route{
				WithRoutePools(&route.RouteImpl{}, []domain.RoutablePool{
					mocks.WithTokenOutDenom(mocks.WithDenoms(defaultPool, []string{DenomFour, DenomFive}), DenomFour),
					mocks.WithTokenOutDenom(mocks.WithDenoms(defaultPool, []string{DenomFour, denomOne}), denomOne),
					mocks.WithTokenOutDenom(defaultPool, denomTwo),
					mocks.WithTokenOutDenom(mocks.WithDenoms(defaultPool, []string{denomTwo, DenomThree}), DenomThree),
				}),
			},
		},
		"2 routes; direct and 2 hop": {
			pools: []domain.PoolI{
				defaultPool, // A: denom 1, 2
				mocks.WithDenoms(defaultPool, []string{denomTwo, DenomThree}), // B: denom 2, 3
				mocks.WithDenoms(defaultPool, []string{denomOne, DenomThree}), // C: denom 1, 3
			},

			maxHops:   2,
			maxRoutes: 2,

			// Route 1: A (denom1 for denom2)
			// Route 2: A (denom1 for denom3) -> B (denom3 for denom2)
			tokenInDenom:  denomOne,
			tokenOutDenom: denomTwo,
			currentRoute:  &route.RouteImpl{},
			poolsUsed:     []bool{false, false, false},
			expectedRoutes: []domain.Route{
				WithRoutePools(&route.RouteImpl{}, []domain.RoutablePool{
					mocks.WithTokenOutDenom(defaultPool, denomTwo),
				}),

				WithRoutePools(&route.RouteImpl{}, []domain.RoutablePool{
					mocks.WithTokenOutDenom(mocks.WithDenoms(defaultPool, []string{denomOne, DenomThree}), DenomThree),
					mocks.WithTokenOutDenom(mocks.WithDenoms(defaultPool, []string{denomTwo, DenomThree}), denomTwo),
				}),
			},
		},
		"routes: first over 4 hops, second 4 hops hop. Shorter subroute exists but not selected. Token in in intermediary path.": {
			pools: []domain.PoolI{
				mocks.WithDenoms(defaultPool, []string{denomOne, DenomThree}),  // A: denom 1, 3
				mocks.WithDenoms(defaultPool, []string{DenomThree, DenomFour}), // B: denom 3, 4
				mocks.WithDenoms(defaultPool, []string{DenomFour, denomOne}),   // C: denom 4, 1
				mocks.WithDenoms(defaultPool, []string{denomOne, denomTwo}),    // D: denom 1, 2
			},

			maxHops:   4,
			maxRoutes: 2,

			// Route 1: A (denom1 for denom3) -> B (denom3 for denom4) -> C (denom4 for denom1) -> D (denom1 for denom2)
			// Route 2: D(denom1 for denom2)
			//
			// Note that the algorithm detects that in the first route, the A -> B -> C part is obsolete since
			// D can be swapped directly. As a result, it returns duplicate routes.
			tokenInDenom:  denomOne,
			tokenOutDenom: denomTwo,
			currentRoute:  &route.RouteImpl{},
			poolsUsed:     []bool{false, false, false, false},
			expectedRoutes: []domain.Route{
				WithRoutePools(&route.RouteImpl{}, []domain.RoutablePool{
					mocks.WithTokenOutDenom(mocks.WithDenoms(defaultPool, []string{denomOne, DenomThree}), DenomThree),
					mocks.WithTokenOutDenom(mocks.WithDenoms(defaultPool, []string{DenomThree, DenomFour}), DenomFour),
					mocks.WithTokenOutDenom(mocks.WithDenoms(defaultPool, []string{DenomFour, denomOne}), denomOne),
					mocks.WithTokenOutDenom(defaultPool, denomTwo)},
				),

				WithRoutePools(&route.RouteImpl{}, []domain.RoutablePool{
					mocks.WithTokenOutDenom(mocks.WithDenoms(defaultPool, []string{DenomFour, denomOne}), DenomFour),
					mocks.WithTokenOutDenom(mocks.WithDenoms(defaultPool, []string{DenomThree, DenomFour}), DenomThree),
					mocks.WithTokenOutDenom(mocks.WithDenoms(defaultPool, []string{denomOne, DenomThree}), denomOne),
					mocks.WithTokenOutDenom(mocks.WithDenoms(defaultPool, []string{denomOne, denomTwo}), denomTwo),
				},
				),
			},
		},
		"2 possible routes with overlap ": {
			pools: []domain.PoolI{
				mocks.WithDenoms(defaultPool, []string{denomOne, denomTwo}),    // A: denom 1, 2
				mocks.WithDenoms(defaultPool, []string{denomTwo, DenomThree}),  // B: denom 2, 3
				mocks.WithDenoms(defaultPool, []string{DenomThree, DenomFour}), // C: denom 3, 4
				mocks.WithDenoms(defaultPool, []string{DenomFive, DenomFour}),  // D: denom 5, 4
				mocks.WithDenoms(defaultPool, []string{DenomThree, DenomFive}), // E: denom 3, 5
			},

			maxHops:   4,
			maxRoutes: 2,

			tokenInDenom:  denomOne,
			tokenOutDenom: DenomFive,
			currentRoute:  &route.RouteImpl{},
			poolsUsed:     []bool{false, false, false, false, false},
			// Possible routes:
			// Route 1: A -> B -> C -> D
			// Route 2: A -> B -> E
			//
			// Note that we expect the first one (which is longer) to not be accounted for.
			expectedRoutes: []domain.Route{
				WithRoutePools(&route.RouteImpl{}, []domain.RoutablePool{
					mocks.WithTokenOutDenom(mocks.WithDenoms(defaultPool, []string{denomOne, denomTwo}), denomTwo),
					mocks.WithTokenOutDenom(mocks.WithDenoms(defaultPool, []string{denomTwo, DenomThree}), DenomThree),
					mocks.WithTokenOutDenom(mocks.WithDenoms(defaultPool, []string{DenomThree, DenomFour}), DenomFour),
					mocks.WithTokenOutDenom(mocks.WithDenoms(defaultPool, []string{DenomFive, DenomFour}), DenomFive),
				}),
				WithRoutePools(&route.RouteImpl{}, []domain.RoutablePool{
					mocks.WithTokenOutDenom(mocks.WithDenoms(defaultPool, []string{denomOne, denomTwo}), denomTwo),
					mocks.WithTokenOutDenom(mocks.WithDenoms(defaultPool, []string{denomTwo, DenomThree}), DenomThree),
					mocks.WithTokenOutDenom(mocks.WithDenoms(defaultPool, []string{DenomThree, DenomFive}), DenomFive),
				}),
			},
		},
		"possible routes; overlapping in the beginning but second one is shorter (second not filtered out)": {
			pools: []domain.PoolI{
				mocks.WithDenoms(defaultPool, []string{denomTwo, DenomThree}),            // A: denom 2, 3
				mocks.WithDenoms(defaultPool, []string{DenomThree, denomTwo}),            // B: denom 3, 2
				mocks.WithDenoms(defaultPool, []string{denomOne, denomTwo}),              // C: denom 1, 2
				mocks.WithDenoms(defaultPool, []string{denomTwo, DenomFour, DenomThree}), // D: denom 2, 4, 3
			},

			maxHops:   4,
			maxRoutes: 2,

			tokenInDenom:  denomOne,
			tokenOutDenom: DenomFour,
			currentRoute:  &route.RouteImpl{},
			poolsUsed:     []bool{false, false, false, false},
			// Possible routes:
			// Route 1: C -> A -> B -> D
			// Route 2: C -> A -> D
			expectedRoutes: []domain.Route{
				WithRoutePools(&route.RouteImpl{}, []domain.RoutablePool{
					mocks.WithTokenOutDenom(mocks.WithDenoms(defaultPool, []string{denomOne, denomTwo}), denomTwo),
					mocks.WithTokenOutDenom(mocks.WithDenoms(defaultPool, []string{denomTwo, DenomThree}), DenomThree),
					mocks.WithTokenOutDenom(mocks.WithDenoms(defaultPool, []string{DenomThree, denomTwo}), denomTwo),
					mocks.WithTokenOutDenom(mocks.WithDenoms(defaultPool, []string{denomTwo, DenomFour, DenomThree}), DenomFour),
				}),
				WithRoutePools(&route.RouteImpl{}, []domain.RoutablePool{
					mocks.WithTokenOutDenom(mocks.WithDenoms(defaultPool, []string{denomOne, denomTwo}), denomTwo),
					mocks.WithTokenOutDenom(mocks.WithDenoms(defaultPool, []string{denomTwo, DenomThree}), DenomThree),
					mocks.WithTokenOutDenom(mocks.WithDenoms(defaultPool, []string{denomTwo, DenomFour, DenomThree}), DenomFour),
				}),
			},
		},
		// If this test is used with max hops of 10, it will select direct route as the last one.
		"3 routes limit; 4 hop, 4 hop, and 3 hop (better routes not selected)": {
			pools: []domain.PoolI{
				defaultPool, // A: denom 1, 2
				mocks.WithDenoms(defaultPool, []string{denomTwo, DenomThree}),  // B: denom 2, 3
				mocks.WithDenoms(defaultPool, []string{DenomFour, DenomSix}),   // C: denom 4, 6
				mocks.WithDenoms(defaultPool, []string{DenomThree, DenomFour}), // D: denom 3, 4
				mocks.WithDenoms(defaultPool, []string{denomOne, DenomThree}),  // E: denom 1, 3
				mocks.WithDenoms(defaultPool, []string{DenomThree, DenomFive}), // F: denom 3, 5
				mocks.WithDenoms(defaultPool, []string{denomTwo, DenomFour}),   // G: denom 2, 4
				mocks.WithDenoms(defaultPool, []string{denomOne, DenomFive}),   // H: denom 1, 5 // note that direct route is not selected due to max routes
				mocks.WithDenoms(defaultPool, []string{DenomFour, DenomFive}),  // I: denom 4, 5
			},

			maxHops:   4,
			maxRoutes: 3,

			// Top 3 routes are selected out:
			// Route 1: A (denom1 for denom2) -> B (denom2 for denom3) -> D (denom3 for denom4) -> I (denom4 for denom5)
			// Route 2: A (denom1 for denom2) -> B (denom2 for denom3) -> E (denom3 for denom1) -> F (denom1 for denom5)
			//    - Note that since F is the direct route, the route is truncated to only have the direct part
			// Route 3: A (denom1 for denom2) -> B (denom2 for denom4) -> I (denom4 for denom5)
			// Route 4: E (denom1 for denom3) -> D (denom3 for denom4) -> I (denom4 for denom5)
			// Route 5: E (denom1 for denom3) -> F (denom3 for denom5) -> G (denom2 for denom4) -> I (denom4 for denom5)
			// Route 6: F (denom1 for denom5)
			tokenInDenom:  denomOne,
			tokenOutDenom: DenomFive,
			currentRoute:  &route.RouteImpl{},
			poolsUsed:     []bool{false, false, false, false, false, false, false, false, false},
			expectedRoutes: []domain.Route{
				// Route 1
				WithRoutePools(&route.RouteImpl{}, []domain.RoutablePool{
					mocks.WithTokenOutDenom(defaultPool, denomTwo),
					mocks.WithTokenOutDenom(mocks.WithDenoms(defaultPool, []string{denomTwo, DenomThree}), DenomThree),
					mocks.WithTokenOutDenom(mocks.WithDenoms(defaultPool, []string{DenomThree, DenomFour}), DenomFour),
					mocks.WithTokenOutDenom(mocks.WithDenoms(defaultPool, []string{DenomFour, DenomFive}), DenomFive),
				}),

				// Route 2
				WithRoutePools(&route.RouteImpl{}, []domain.RoutablePool{
					// Note that a component of the route is obsolete due to having a direct route in the intermediary path
					mocks.WithTokenOutDenom(defaultPool, denomTwo),
					mocks.WithTokenOutDenom(mocks.WithDenoms(defaultPool, []string{denomTwo, DenomThree}), DenomThree),
					mocks.WithTokenOutDenom(mocks.WithDenoms(defaultPool, []string{denomOne, DenomThree}), denomOne),
					mocks.WithTokenOutDenom(mocks.WithDenoms(defaultPool, []string{denomOne, DenomFive}), DenomFive),
				}),

				// Route 3
				WithRoutePools(&route.RouteImpl{}, []domain.RoutablePool{
					mocks.WithTokenOutDenom(defaultPool, denomTwo),
					mocks.WithTokenOutDenom(mocks.WithDenoms(defaultPool, []string{denomTwo, DenomThree}), DenomThree),
					mocks.WithTokenOutDenom(mocks.WithDenoms(defaultPool, []string{DenomThree, DenomFive}), DenomFive),
				}),
			},
		},
		// errors
		"error: nil route": {
			pools: []domain.PoolI{},

			tokenInDenom:   denomOne,
			tokenOutDenom:  denomTwo,
			currentRoute:   nil,
			poolsUsed:      []bool{},
			expectedRoutes: []domain.Route{},

			expectedError: routerusecase.ErrNilCurrentRoute,
		},
		"error: sorted pools and pools used mismatch": {
			pools: []domain.PoolI{},

			tokenInDenom:  denomOne,
			tokenOutDenom: denomTwo,
			currentRoute:  &route.RouteImpl{},
			poolsUsed:     []bool{true, false},

			expectedError: routerusecase.SortedPoolsAndPoolsUsedLengthMismatchError{
				SortedPoolsLen: 0,
				PoolsUsedLen:   2,
			},
		},
		"error: no pools but non empty pools in route": {
			pools: []domain.PoolI{},

			tokenInDenom:  denomOne,
			tokenOutDenom: denomTwo,
			currentRoute:  WithRoutePools(&route.RouteImpl{}, []domain.RoutablePool{defaultPool}),
			poolsUsed:     []bool{},

			expectedError: routerusecase.SortedPoolsAndPoolsInRouteLengthMismatchError{
				SortedPoolsLen: 0,
				PoolsInRoute:   1,
			},
		},
	}

	for name, tc := range tests {
		s.Run(name, func() {

			// Get taker fees for all pools.
			takerFees := s.getTakerFeeMapForAllPoolTokenPairs(tc.pools)

			r := routerusecase.NewRouter([]uint64{}, takerFees, tc.maxHops, tc.maxRoutes, 0, 0, nil)
			r = routerusecase.WithSortedPools(r, tc.pools)

			routes, err := r.FindRoutes(tc.tokenInDenom, tc.tokenOutDenom, tc.currentRoute, tc.poolsUsed, tc.previousTokenOutDenoms)

			if tc.expectedError != nil {
				s.Require().Error(err)
				s.Require().Equal(tc.expectedError.Error(), err.Error())
				return
			}

			s.Require().NoError(err)

			s.validateFoundRoutes(tc, routes)
		})
	}
}

func (s *RouterTestSuite) TestGetCandidateRoutes() {
	tests := map[string]struct {
		pools []domain.PoolI

		maxHops   int
		maxRoutes int

		tokenInDenom           string
		tokenOutDenom          string
		currentRoute           domain.Route
		poolsUsed              []bool
		previousTokenOutDenoms []string

		expectedRoutes []domain.Route
		expectedError  error
	}{
		"2 possible routes with overlap ": {
			pools: []domain.PoolI{
				mocks.WithDenoms(DefaultPool, []string{DenomOne, DenomTwo}),                                       // A: denom 1, 2
				mocks.WithPoolID(mocks.WithDenoms(DefaultPool, []string{DenomTwo, DenomThree}), defaultPoolID+1),  // B: denom 2, 3
				mocks.WithPoolID(mocks.WithDenoms(DefaultPool, []string{DenomThree, DenomFour}), defaultPoolID+2), // C: denom 3, 4
				mocks.WithPoolID(mocks.WithDenoms(DefaultPool, []string{DenomFive, DenomFour}), defaultPoolID+3),  // D: denom 5, 4
				mocks.WithPoolID(mocks.WithDenoms(DefaultPool, []string{DenomThree, DenomFive}), defaultPoolID+4), // E: denom 3, 5
			},

			maxHops:   4,
			maxRoutes: 2,

			tokenInDenom:  DenomOne,
			tokenOutDenom: DenomFive,
			currentRoute:  &route.RouteImpl{},
			poolsUsed:     []bool{false, false, false, false, false},
			// Possible routes:
			// Route 1: A -> B -> C -> D
			// Route 2: A -> B -> E
			//
			// Note that we expect the first one (which is longer) to not be accounted for
			// due to overlapping pool IDs.
			expectedRoutes: []domain.Route{
				WithRoutePools(&route.RouteImpl{}, []domain.RoutablePool{
					mocks.WithTokenOutDenom(mocks.WithDenoms(DefaultPool, []string{DenomOne, DenomTwo}), DenomTwo),
					mocks.WithPoolID(mocks.WithTokenOutDenom(mocks.WithDenoms(DefaultPool, []string{DenomTwo, DenomThree}), DenomThree), defaultPoolID+1),
					mocks.WithPoolID(mocks.WithTokenOutDenom(mocks.WithDenoms(DefaultPool, []string{DenomThree, DenomFive}), DenomFive), defaultPoolID+4),
				}),
			},
		},
		// If this test is used with max hops of 10, it will select direct route as the last one.
		"routes limit; 3 hop preferred, same pool IDs are filtered out (better routes not selected)": {
			pools: []domain.PoolI{
				DefaultPool, // A: denom 1, 2
				mocks.WithPoolID(mocks.WithDenoms(DefaultPool, []string{DenomTwo, DenomThree}), defaultPoolID+1),  // B: denom 2, 3
				mocks.WithPoolID(mocks.WithDenoms(DefaultPool, []string{DenomFour, DenomSix}), defaultPoolID+2),   // C: denom 4, 6
				mocks.WithPoolID(mocks.WithDenoms(DefaultPool, []string{DenomThree, DenomFour}), defaultPoolID+3), // D: denom 3, 4
				mocks.WithPoolID(mocks.WithDenoms(DefaultPool, []string{DenomOne, DenomThree}), defaultPoolID+4),  // E: denom 1, 3
				mocks.WithPoolID(mocks.WithDenoms(DefaultPool, []string{DenomThree, DenomFive}), defaultPoolID+5), // F: denom 3, 5
				mocks.WithPoolID(mocks.WithDenoms(DefaultPool, []string{DenomTwo, DenomFour}), defaultPoolID+6),   // G: denom 2, 4
				mocks.WithPoolID(mocks.WithDenoms(DefaultPool, []string{DenomOne, DenomFive}), defaultPoolID+7),   // H: denom 1, 5 // note that direct route is not selected due to max routes
				mocks.WithPoolID(mocks.WithDenoms(DefaultPool, []string{DenomFour, DenomFive}), defaultPoolID+8),  // I: denom 4, 5
			},

			maxHops:   4,
			maxRoutes: 3,

			// Top 3 routes are selected out:
			// Route 1: A (denom1 for denom2) -> B (denom2 for denom3) -> D (denom3 for denom4) -> I (denom4 for denom5)
			// Route 2: A (denom1 for denom2) -> B (denom2 for denom3) -> E (denom3 for denom1) -> F (denom1 for denom5)
			//    - Note that since F is the direct route, the route is truncated to only have the direct part
			// Route 3: A (denom1 for denom2) -> B (denom2 for denom4) -> I (denom4 for denom5)
			// Route 4: E (denom1 for denom3) -> D (denom3 for denom4) -> I (denom4 for denom5)
			// Route 5: E (denom1 for denom3) -> F (denom3 for denom5) -> G (denom2 for denom4) -> I (denom4 for denom5)
			// Route 6: F (denom1 for denom5)
			tokenInDenom:  DenomOne,
			tokenOutDenom: DenomFive,
			currentRoute:  &route.RouteImpl{},
			poolsUsed:     []bool{false, false, false, false, false, false, false, false, false},
			expectedRoutes: []domain.Route{
				// Note that routes get reordered by the number of hops.
				// See similar test in TestFindRoutes for comparison.

				// Note that route 1 & 2 are removed due to overlapping pool IDS. preferring the shorter route.
				// This is actually suboptimal and we should revisit this.

				// Route 3
				WithRoutePools(&route.RouteImpl{}, []domain.RoutablePool{
					mocks.WithTokenOutDenom(DefaultPool, DenomTwo),
					mocks.WithPoolID(mocks.WithTokenOutDenom(mocks.WithDenoms(DefaultPool, []string{DenomTwo, DenomThree}), DenomThree), defaultPoolID+1),
					mocks.WithPoolID(mocks.WithTokenOutDenom(mocks.WithDenoms(DefaultPool, []string{DenomThree, DenomFive}), DenomFive), defaultPoolID+5),
				}),

				// Note that the third route is removed. See similar test in TestFindRoutes for comparison.
			},
		},
	}

	for name, tc := range tests {
		s.Run(name, func() {

			takerFees := s.getTakerFeeMapForAllPoolTokenPairs(tc.pools)

			r := routerusecase.NewRouter([]uint64{}, takerFees, tc.maxHops, tc.maxRoutes, 3, 0, nil)
			r = routerusecase.WithSortedPools(r, tc.pools)

			routes, err := r.GetCandidateRoutes(tc.tokenInDenom, tc.tokenOutDenom)

			if tc.expectedError != nil {
				s.Require().Error(err)
				return
			}
			s.Require().NoError(err)

			s.validateFoundRoutes(tc, routes)
		})
	}
}

// validateFoundRoutes validates that the routes are as expected.
func (s *RouterTestSuite) validateFoundRoutes(tc routesTestCase, routes []domain.Route) {
	s.Require().Equal(len(tc.expectedRoutes), len(routes))
	for i, expectedRoute := range tc.expectedRoutes {
		actualRoute := routes[i]

		expectedPools := expectedRoute.GetPools()
		actualPools := actualRoute.GetPools()

		s.Require().Equal(len(expectedPools), len(actualPools), fmt.Sprintf("expected route: %s, \nactual route %s", expectedRoute, actualRoute))

		for j, expectedPool := range expectedPools {
			s.Require().Equal(expectedPool.GetId(), actualPools[j].GetId())
			s.Require().Equal(expectedPool.GetTokenOutDenom(), actualPools[j].GetTokenOutDenom())
			s.Require().Equal(expectedPool.GetPoolDenoms(), actualPools[j].GetPoolDenoms())
		}
	}
}