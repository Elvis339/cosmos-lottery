package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "cosmos-lottery/testutil/keeper"
	"cosmos-lottery/testutil/nullify"
	"cosmos-lottery/x/lottery/keeper"
	"cosmos-lottery/x/lottery/types"
)

func createTestActiveLottery(keeper *keeper.Keeper, ctx sdk.Context, id uint64) types.ActiveLottery {
	item := types.ActiveLottery{
		LotteryId: id,
	}
	keeper.SetActiveLottery(ctx, item)
	return item
}

func TestActiveLotteryGet(t *testing.T) {
	keeper, ctx := keepertest.LotteryKeeper(t)
	item := createTestActiveLottery(keeper, ctx, 1)
	rst, found := keeper.GetActiveLottery(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
	require.Equal(t, uint64(1), rst.LotteryId)
}

func TestActiveLotteryRemove(t *testing.T) {
	keeper, ctx := keepertest.LotteryKeeper(t)
	createTestActiveLottery(keeper, ctx, 1)
	keeper.RemoveActiveLottery(ctx)
	_, found := keeper.GetActiveLottery(ctx)
	require.False(t, found)
}

func TestActiveLottery_Increment(t *testing.T) {
	keeper, ctx := keepertest.LotteryKeeper(t)
	createTestActiveLottery(keeper, ctx, 1)
	keeper.IncrementActiveLottery(ctx)
	rst, found := keeper.GetActiveLottery(ctx)
	require.True(t, found)
	require.Equal(t, uint64(2), rst.LotteryId)
}
