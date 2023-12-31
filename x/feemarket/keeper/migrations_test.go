package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	feemarketkeeper "github.com/tokenize-titan/ethermint/x/feemarket/keeper"
	"github.com/tokenize-titan/ethermint/x/feemarket/types"
)

type mockSubspace struct {
	ps types.Params
}

func newMockSubspace(ps types.Params) mockSubspace {
	return mockSubspace{ps: ps}
}

func (ms mockSubspace) GetParamSetIfExists(_ sdk.Context, ps types.LegacyParams) {
	*ps.(*types.Params) = ms.ps
}

func (suite *KeeperTestSuite) TestMigrations() {
	legacySubspace := newMockSubspace(types.DefaultParams())
	_ = feemarketkeeper.NewMigrator(suite.app.FeeMarketKeeper, legacySubspace) // migrator :=

	testCases := []struct {
		name        string
		migrateFunc func(ctx sdk.Context) error
	}{}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			err := tc.migrateFunc(suite.ctx)
			suite.Require().NoError(err)
		})
	}
}
