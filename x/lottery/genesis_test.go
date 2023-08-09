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

		ActiveLottery: &types.ActiveLottery{
			LotteryId: 65,
		},
		LotteryTransactionList: []types.LotteryTransaction{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		LotteryTransactionCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.LotteryKeeper(t)
	lottery.InitGenesis(ctx, *k, genesisState)
	got := lottery.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.ActiveLottery, got.ActiveLottery)
	require.ElementsMatch(t, genesisState.LotteryTransactionList, got.LotteryTransactionList)
	require.Equal(t, genesisState.LotteryTransactionCount, got.LotteryTransactionCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
