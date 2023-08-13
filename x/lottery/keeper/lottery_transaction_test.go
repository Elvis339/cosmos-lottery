package keeper_test

import (
	"cosmos-lottery/testutil/sample"
	"github.com/cometbft/cometbft/libs/rand"
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
		items[i].CreatedBy = sample.AccAddress()
		items[i].Bet = sdk.NewInt64Coin(types.TokenDenom, int64(rand.Intn(100-1)+1))
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

func TestLotteryTransactionAppend(t *testing.T) {
	keeper, ctx := keepertest.LotteryKeeper(t)

	alice := sample.AccAddress()
	bob := sample.AccAddress()

	items := []types.LotteryTransaction{{
		Id:        0,
		Bet:       sdk.NewInt64Coin(types.TokenDenom, 5),
		CreatedBy: sample.AccAddress(),
		LotteryId: 1,
	}, {
		Id:        1,
		Bet:       sdk.NewInt64Coin(types.TokenDenom, 3),
		CreatedBy: bob,
		LotteryId: 1,
	}, {
		Id:        2,
		Bet:       sdk.NewInt64Coin(types.TokenDenom, 10),
		CreatedBy: alice,
		LotteryId: 1,
	}, {
		Id:        3,
		Bet:       sdk.NewInt64Coin(types.TokenDenom, 2),
		CreatedBy: sample.AccAddress(),
		LotteryId: 1,
	}}

	for i := range items {
		keeper.AppendLotteryTransaction(ctx, items[i])
	}

	// Alice figures out that her bet is the highest one, she now wants to de-risk, so she sends another LotteryTx
	// with lower bet
	aliceLotteryTx := types.LotteryTransaction{
		Bet:       sdk.NewInt64Coin(types.TokenDenom, 4),
		CreatedBy: alice,
		LotteryId: 1,
	}
	keeper.AppendLotteryTransaction(ctx, aliceLotteryTx)

	newAliceLotteryTx, found := keeper.GetLotteryTransaction(ctx, 2)
	// Ensure Alice can be found with the same id as her initial bet id
	// Initial position is retained and only unique transaction is recorded
	require.True(t, found)

	// Ensure Alice's old bet has been updated
	require.Equal(t, uint64(4), newAliceLotteryTx.Bet.Amount.Uint64())

	// Ensure the count has not changed
	count := keeper.GetLotteryTransactionCount(ctx)
	require.Equal(t, count, uint64(4))
}

func TestLotteryTransactionAppend_WhenDifferentLotteryId(t *testing.T) {
	keeper, ctx := keepertest.LotteryKeeper(t)

	alice := sample.AccAddress()
	bob := sample.AccAddress()

	items := []types.LotteryTransaction{{
		Id:        0,
		Bet:       sdk.NewInt64Coin(types.TokenDenom, 5),
		CreatedBy: sample.AccAddress(),
		LotteryId: 2,
	}, {
		Id:        1,
		Bet:       sdk.NewInt64Coin(types.TokenDenom, 3),
		CreatedBy: bob,
		LotteryId: 2,
	}, {
		Id:        2,
		Bet:       sdk.NewInt64Coin(types.TokenDenom, 10),
		CreatedBy: alice,
		// Alice Lottery id is 1
		LotteryId: 1,
	}}

	for i := range items {
		keeper.AppendLotteryTransaction(ctx, items[i])
	}

	keeper.AppendLotteryTransaction(ctx, types.LotteryTransaction{
		Bet:       sdk.NewInt64Coin(types.TokenDenom, 3),
		CreatedBy: alice,
		// Different lottery ID
		LotteryId: 2,
	})

	found, ltryTxId := keeper.LotteryTransactionMetadata.GetLotteryTransactionId(alice)
	require.True(t, found)

	val, found := keeper.GetLotteryTransaction(ctx, ltryTxId)
	require.True(t, found)
	require.Equal(t, uint64(2), val.LotteryId)

	require.Equal(t, uint64(4), keeper.GetLotteryTransactionCount(ctx))
}
