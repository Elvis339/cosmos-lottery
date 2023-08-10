package keeper

import (
	"context"
	"cosmos-lottery/x/lottery/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"strconv"
)

func (k msgServer) PlaceBet(goCtx context.Context, msg *types.MsgPlaceBet) (*types.MsgPlaceBetResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Will error if there's no active lottery set
	// because active lottery is deterministic id generator that represents current active lottery set & is managed by the blockchain
	activeLottery, found := k.Keeper.GetActiveLottery(ctx)
	if !found {
		return nil, sdkerrors.ErrNotFound.Wrapf("active lottery is not set")
	}

	addr, err := sdk.AccAddressFromBech32(msg.GetCreator())
	if err != nil {
		return nil, err
	}

	bet := sdk.NewInt64Coin("token", int64(msg.GetBet()))
	amount := bet.Add(types.Fee).Add(types.MinBet)

	balance := k.bankKeeper.GetBalance(ctx, addr, "token")

	if balance.IsLT(amount) {
		return nil, sdkerrors.ErrInsufficientFunds.Wrapf("could not place a bet")
	}

	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, addr, types.ModuleName, sdk.Coins{amount})
	if err != nil {
		return nil, err
	}

	k.AppendLotteryTransaction(ctx, types.LotteryTransaction{
		Bet:       bet,
		CreatedBy: msg.GetCreator(),
		LotteryId: activeLottery.LotteryId,
	})

	lotteryTxs := k.GetAllLotteryTransaction(ctx)

	// Update lottery pool
	err = k.UpdateLotteryPool(ctx, strconv.FormatUint(activeLottery.LotteryId, 10), lotteryTxs)
	if err != nil {
		return nil, err
	}

	return &types.MsgPlaceBetResponse{}, nil
}
