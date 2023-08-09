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

func createTestActiveLottery(keeper *keeper.Keeper, ctx sdk.Context) types.ActiveLottery {
	item := types.ActiveLottery{}
	keeper.SetActiveLottery(ctx, item)
	return item
}

func TestActiveLotteryGet(t *testing.T) {
	keeper, ctx := keepertest.LotteryKeeper(t)
	item := createTestActiveLottery(keeper, ctx)
	rst, found := keeper.GetActiveLottery(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestActiveLotteryRemove(t *testing.T) {
	keeper, ctx := keepertest.LotteryKeeper(t)
	createTestActiveLottery(keeper, ctx)
	keeper.RemoveActiveLottery(ctx)
	_, found := keeper.GetActiveLottery(ctx)
	require.False(t, found)
}
