package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) TriggerEndLottery(ctx sdk.Context) {
	blockHeader := ctx.BlockHeader()
	blockProposerAddress, err := sdk.Bech32ifyAddressBytes("cosmos", blockHeader.GetProposerAddress())
	if err != nil {
		return
	}

	blockProposerHasBets, _ := k.LotteryTransactionMetadata.GetLotteryTransactionId(blockProposerAddress)

	if blockProposerHasBets == true {
		ctx.Logger().Info("block proposer has placed bets skipping...")
		return
	}

	tx, err := k.GetWinner(ctx)

	if err != nil {
		return
	}

	if tx == nil {
		return
	}

	winnerAddr, err := sdk.AccAddressFromBech32(tx.GetCreatedBy())
	if err != nil {
		return
	}

	err = k.EndLottery(ctx, winnerAddr)
	if err != nil {
		ctx.Logger().Error(err.Error())
	}
}
