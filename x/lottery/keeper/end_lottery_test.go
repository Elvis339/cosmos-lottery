package keeper_test

import (
	"cosmos-lottery/testutil/keeper"
	"cosmos-lottery/testutil/sample"
	"cosmos-lottery/x/lottery/types"
	"github.com/cometbft/cometbft/libs/log"
	"github.com/cometbft/cometbft/libs/rand"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestTriggerEndLottery_WithBlockProposerAsTxSender(t *testing.T) {
	mock := keeper.NewMockLotteryWithGenesis(t)
	stateStore := mock.StateStore
	k := mock.LotteryKeeper
	bankKeeper := mock.BankKeeper

	proposer, _ := sdk.AccAddressFromBech32(sample.AccAddress())

	ctx := sdk.NewContext(stateStore, tmproto.Header{
		ProposerAddress: proposer.Bytes(),
	}, false, log.NewNopLogger())

	bankKeeper.EXPECT().SendCoinsFromModuleToAccount(ctx, types.ModuleName, gomock.Any(), gomock.Any())

	for i := 0; i < 20; i++ {
		lotteryTx := sample.CreateLotteryTx(sample.AccAddress(), int64(rand.Intn(100-1)+1))
		k.AppendLotteryTransaction(ctx, lotteryTx)
	}

	// Block proposer placed a bet
	k.AppendLotteryTransaction(ctx, sample.CreateLotteryTx(proposer.String(), int64(rand.Intn(100-1)+1)))

	k.TriggerEndLottery(ctx)

	// Verify active lottery has not been incremented
	activeLottery, _ := k.GetActiveLottery(ctx)
	require.Equal(t, uint64(1), activeLottery.LotteryId)

	// Verify counter has not been set to 0
	require.Equal(t, uint64(21), k.GetLotteryTransactionCount(ctx))
}
