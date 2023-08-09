package keeper_test

import (
	"testing"

	keepertest "cosmos-lottery/testutil/keeper"
	"cosmos-lottery/testutil/nullify"
	"cosmos-lottery/x/lottery/keeper"
	"cosmos-lottery/x/lottery/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func createNLotteryTransaction(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.LotteryTransaction {
	items := make([]types.LotteryTransaction, n)
	for i := range items {
		items[i].Id = keeper.AppendLotteryTransaction(ctx, items[i])
	}
	return items
}

func TestLotteryTransactionGet(t *testing.T) {
	keeper, ctx := keepertest.LotteryKeeper(t)
	items := createNLotteryTransaction(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetLotteryTransaction(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestLotteryTransactionRemove(t *testing.T) {
	keeper, ctx := keepertest.LotteryKeeper(t)
	items := createNLotteryTransaction(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveLotteryTransaction(ctx, item.Id)
		_, found := keeper.GetLotteryTransaction(ctx, item.Id)
		require.False(t, found)
	}
}

func TestLotteryTransactionGetAll(t *testing.T) {
	keeper, ctx := keepertest.LotteryKeeper(t)
	items := createNLotteryTransaction(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllLotteryTransaction(ctx)),
	)
}

func TestLotteryTransactionCount(t *testing.T) {
	keeper, ctx := keepertest.LotteryKeeper(t)
	items := createNLotteryTransaction(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetLotteryTransactionCount(ctx))
}
