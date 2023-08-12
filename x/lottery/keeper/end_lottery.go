package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k Keeper) TriggerEndLottery(ctx sdk.Context) {
	blockHeader := ctx.BlockHeader()
	blockProposerAddress, err := sdk.Bech32ifyAddressBytes("cosmos", blockHeader.GetProposerAddress())
	if err != nil {
		ctx.Logger().Error(sdkerrors.ErrInvalidAddress.Error())
		return
	}

	blockProposerHasBets, _ := k.lotteryTxMeta.GetLotteryTransactionId(blockProposerAddress)

	if blockProposerHasBets == true {
		return
	}

	tx, err := k.GetWinner(ctx)
	if err != nil {
		return
	}
	winnerAddr, _ := sdk.AccAddressFromBech32(tx.CreatedBy)

	err = k.EndLottery(ctx, winnerAddr)
	if err != nil {
		ctx.Logger().Error(err.Error())
	}
}
