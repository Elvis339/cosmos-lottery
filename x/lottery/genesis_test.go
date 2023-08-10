package lottery_test

import (
	"testing"

	keepertest "cosmos-lottery/testutil/keeper"
	"cosmos-lottery/testutil/nullify"
	"cosmos-lottery/x/lottery"
	"cosmos-lottery/x/lottery/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		ActiveLottery: types.ActiveLottery{
			LotteryId: 1,
		},
		LotteryList: []types.Lottery{
			{
				Index: "1",
				Fee:   types.Fee,
				Pool:  types.Fee,
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.LotteryKeeper(t)
	lottery.InitGenesis(ctx, *k, genesisState)
	got := lottery.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.ActiveLottery, got.ActiveLottery)
	require.ElementsMatch(t, genesisState.LotteryList, got.LotteryList)
	// this line is used by starport scaffolding # genesis/test/assert
}
