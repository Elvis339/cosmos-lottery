package keeper_test

import (
	"context"
	"cosmos-lottery/testutil/sample"
	"cosmos-lottery/x/lottery"
	"fmt"
	"strconv"
	"testing"

	keepertest "cosmos-lottery/testutil/keeper"
	"cosmos-lottery/x/lottery/keeper"
	"cosmos-lottery/x/lottery/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.LotteryKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}

func TestMsgServer(t *testing.T) {
	ms, ctx := setupMsgServer(t)
	require.NotNil(t, ms)
	require.NotNil(t, ctx)
}

func mockGenesis(t *testing.T, lotteryId uint64, lotteryList ...types.Lottery) (types.MsgServer, context.Context, keeper.Keeper) {
	// Setup genesis
	k, ctx := keepertest.LotteryKeeper(t)
	lottery.InitGenesis(ctx, *k, types.GenesisState{
		ActiveLottery: types.ActiveLottery{
			LotteryId: lotteryId,
		},
		LotteryList: lotteryList,
	})
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx), *k
}

func TestMsgServer_PlaceBetError(t *testing.T) {
	testCases := []struct {
		desc     string
		msg      *types.MsgPlaceBet
		valid    bool
		callback func(t *testing.T) (types.MsgServer, context.Context, keeper.Keeper)
	}{
		{
			desc: "should error because active lottery is not set",
			msg: &types.MsgPlaceBet{
				Creator: sample.AccAddress(),
				Bet:     1,
			},
			valid:    false,
			callback: nil,
		},
		{
			desc: "should error if lottery is not found",
			msg: &types.MsgPlaceBet{
				Creator: sample.AccAddress(),
				Bet:     4,
			},
			valid: false,
			callback: func(t *testing.T) (types.MsgServer, context.Context, keeper.Keeper) {
				lotteryList := []types.Lottery{{Index: "5"}}
				return mockGenesis(t, 1, lotteryList...)
			},
		},
		{
			desc: "should not error if active lottery and lottery is set correctly",
			msg: &types.MsgPlaceBet{
				Creator: sample.AccAddress(),
				Bet:     0,
			},
			valid: true,
			callback: func(t *testing.T) (types.MsgServer, context.Context, keeper.Keeper) {
				activeLotteryId := 1
				lotteryList := []types.Lottery{{Index: strconv.FormatUint(uint64(activeLotteryId), 10)}}
				return mockGenesis(t, uint64(activeLotteryId), lotteryList...)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			var ms types.MsgServer
			var ctx context.Context

			if tc.callback != nil {
				callbackMs, callBackCtx, _ := tc.callback(t)
				ms = callbackMs
				ctx = callBackCtx
			} else {
				msgServerMs, msgServerCtx := setupMsgServer(t)
				ms = msgServerMs
				ctx = msgServerCtx
			}

			_, err := ms.PlaceBet(ctx, tc.msg)
			fmt.Println(err)
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}

func TestMsgServer_PlaceBetNotEnoughBalance(t *testing.T) {
	mock := keepertest.NewMockLotteryWithGenesis(t)
	k := mock.LotteryKeeper
	bankKeeper := mock.BankKeeper
	ctx := mock.Ctx

	victor, _ := sdk.AccAddressFromBech32(sample.AccAddress())

	msgServer := keeper.NewMsgServerImpl(*k)

	bet := sdk.NewInt64Coin(types.TokenDenom, 3)
	amountCharged := bet.Add(types.Fee).Add(types.MinBet)

	bankKeeper.EXPECT().GetBalance(ctx, victor, types.TokenDenom).AnyTimes().Return(sdk.NewInt64Coin(types.TokenDenom, 0))
	bankKeeper.EXPECT().SendCoinsFromAccountToModule(ctx, victor, types.ModuleName, sdk.Coins{amountCharged}).AnyTimes()

	// Place a bet
	_, err := msgServer.PlaceBet(ctx, &types.MsgPlaceBet{
		Creator: victor.String(),
		Bet:     bet.Amount.Uint64(),
	})
	require.Error(t, err)
}

func TestMsgServer_PlaceBetSuccess(t *testing.T) {
	mock := keepertest.NewMockLotteryWithGenesis(t)
	k := mock.LotteryKeeper
	bankKeeper := mock.BankKeeper
	ctx := mock.Ctx

	victor, _ := sdk.AccAddressFromBech32(sample.AccAddress())

	msgServer := keeper.NewMsgServerImpl(*k)

	bet := sdk.NewInt64Coin(types.TokenDenom, 3)
	amountCharged := bet.Add(types.Fee).Add(types.MinBet)

	bankKeeper.EXPECT().GetBalance(ctx, victor, types.TokenDenom).AnyTimes().Return(sdk.NewInt64Coin(types.TokenDenom, 10))
	bankKeeper.EXPECT().SendCoinsFromAccountToModule(ctx, victor, types.ModuleName, sdk.Coins{amountCharged}).AnyTimes()

	// Place a bet
	_, err := msgServer.PlaceBet(ctx, &types.MsgPlaceBet{
		Creator: victor.String(),
		Bet:     bet.Amount.Uint64(),
	})
	require.NoError(t, err)

	// === Assert state has been changed correctly ===
	getLottery, found := k.GetLottery(ctx, "1")
	require.Equal(t, found, true)

	// LotteryTransactions length should be increased by 1
	require.Equal(t, 1, len(k.GetAllLotteryTransaction(ctx)))
	require.Equal(t, uint64(1), k.GetLotteryTransactionCount(ctx))

	// Pool should be increased by: Bet + MinBet + Fee
	require.Equal(t, amountCharged.Amount.Uint64(), getLottery.Pool.Amount.Uint64())
}
