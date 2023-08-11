package keeper_test

import (
	"cosmos-lottery/testutil/sample"
	"strconv"
	"testing"

	keepertest "cosmos-lottery/testutil/keeper"
	"cosmos-lottery/testutil/nullify"
	"cosmos-lottery/x/lottery/keeper"
	"cosmos-lottery/x/lottery/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNLottery(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Lottery {
	items := make([]types.Lottery, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetLottery(ctx, items[i])
	}
	return items
}

func TestLotteryGet(t *testing.T) {
	keeper, ctx := keepertest.LotteryKeeper(t)
	items := createNLottery(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetLottery(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestLotteryRemove(t *testing.T) {
	keeper, ctx := keepertest.LotteryKeeper(t)
	items := createNLottery(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveLottery(ctx,
			item.Index,
		)
		_, found := keeper.GetLottery(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestLotteryGetAll(t *testing.T) {
	keeper, ctx := keepertest.LotteryKeeper(t)
	items := createNLottery(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllLottery(ctx)),
	)
}

func TestLottery_UpdatePool(t *testing.T) {
	keeper, ctx := keepertest.SetupLotteryKeeperWithGenesis(t)
	coins := sdk.Coins{
		sdk.NewInt64Coin("token", 10),
		sdk.NewInt64Coin("token", 10),
	}

	for _, coin := range coins {
		err := keeper.UpdateLotteryPool(ctx, "1", coin)
		require.NoError(t, err)
	}

	// Index "2" does not exist
	err := keeper.UpdateLotteryPool(ctx, "2", sdk.NewInt64Coin("token", 10))
	require.Error(t, err)

	currentLottery, _ := keeper.GetLottery(ctx, "1")
	require.Equal(t, uint64(20), currentLottery.Pool.Amount.Uint64())
}

func TestLotteryEndBlock_NotEnoughTx(t *testing.T) {
	k, ctx := keepertest.SetupLotteryKeeperWithGenesis(t)

	createNLotteryTransaction(k, ctx, 9)

	addr, _ := sdk.AccAddressFromBech32(sample.AccAddress())

	err := k.LotteryEndBlock(ctx, addr)
	require.NoError(t, err)

	require.Equal(t, uint64(9), k.GetLotteryTransactionCount(ctx))

	currentLotteryId, found := k.GetActiveLottery(ctx)
	require.True(t, found)

	// current lottery id should not be changed because there weren't enough transactions
	require.Equal(t, uint64(1), currentLotteryId.LotteryId)
}

func TestLotteryEndBlock_LowestBetWinner(t *testing.T) {
	k, ctx := keepertest.SetupLotteryKeeperWithGenesis(t)
	lotteryTx := make([]types.LotteryTransaction, 10)
	bets := [10]int{33, 71, 54, 31, 11, 99, 76, 4, 65, 82}
	alice, _ := sdk.AccAddressFromBech32(sample.AccAddress())

	currentLotteryId, found := k.GetActiveLottery(ctx)
	require.True(t, found)

	lottery, found := k.GetLottery(ctx, strconv.FormatUint(currentLotteryId.LotteryId, 10))
	require.True(t, found)
	require.Equal(t, strconv.FormatUint(currentLotteryId.LotteryId, 10), lottery.Index)

	for i := range lotteryTx {
		if bets[i] == 11 {
			lotteryTx[i].CreatedBy = alice.String()
		} else {
			lotteryTx[i].CreatedBy = sample.AccAddress()
		}

		lotteryTx[i].Bet = sdk.NewInt64Coin("token", int64(bets[i]))
		lotteryTx[i].Id = k.AppendLotteryTransaction(ctx, lotteryTx[i])
	}

	err := k.LotteryEndBlock(ctx, alice)
	require.NoError(t, err)

	// Verify LotteryTransactionCounter has been set to 0
	k.SetLotteryTransactionCount(ctx, 0)
	require.Equal(t, uint64(0), k.GetLotteryTransactionCount(ctx))
	require.Equal(t, 0, len(k.GetAllLotteryTransaction(ctx)))

	// Verify active lottery has been incremented
	next, found := k.GetActiveLottery(ctx)
	require.True(t, found)

	nextLottery, found := k.GetLottery(ctx, strconv.FormatUint(next.LotteryId, 10))
	require.True(t, found)
	require.Equal(t, strconv.FormatUint(next.LotteryId, 10), nextLottery.Index)

	// Since the winner placed the lowest bet
	// verify nextLottery.Pool is equal to the previous
	require.Equal(t, lottery.Pool.Amount.Uint64(), nextLottery.Pool.Amount.Uint64())
}

func TestLotteryEndBlock_HighestBetWinner(t *testing.T) {}

func TestLotteryEndBlock_Winner(t *testing.T) {}
