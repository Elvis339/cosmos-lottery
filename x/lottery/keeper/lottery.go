package keeper

import (
	"context"
	"cosmos-lottery/x/lottery/types"
	"fmt"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"strconv"
)

// SetLottery set a specific lottery in the store from its index
func (k Keeper) SetLottery(ctx sdk.Context, lottery types.Lottery) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LotteryKeyPrefix))
	b := k.cdc.MustMarshal(&lottery)
	store.Set(types.LotteryKey(
		lottery.Index,
	), b)
}

// GetLottery returns a lottery from its index
func (k Keeper) GetLottery(
	ctx sdk.Context,
	index string,

) (val types.Lottery, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LotteryKeyPrefix))

	b := store.Get(types.LotteryKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveLottery removes a lottery from the store
func (k Keeper) RemoveLottery(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LotteryKeyPrefix))
	store.Delete(types.LotteryKey(
		index,
	))
}

// GetAllLottery returns all lottery
func (k Keeper) GetAllLottery(ctx sdk.Context) (list []types.Lottery) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LotteryKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Lottery
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// UpdateLotteryPool sum LotteryTransaction with Fee and MinBet
func (k Keeper) UpdateLotteryPool(ctx sdk.Context, index string, amount sdk.Coin) error {
	lottery, found := k.GetLottery(ctx, index)

	if !found {
		return sdkerrors.ErrNotFound.Wrapf(fmt.Sprintf("lottery with index %s", index))
	}

	newAmount := lottery.Pool.Add(amount)
	lottery.Pool = newAmount

	k.SetLottery(ctx, lottery)

	return nil
}

func (k Keeper) EndLottery(goCtx context.Context, winner sdk.AccAddress) error {
	ctx := sdk.UnwrapSDKContext(goCtx)

	currentLotteryId, found := k.GetActiveLottery(ctx)
	if !found {
		panic("active lottery is not set!")
	}

	lotteryTransactionCount := k.GetLotteryTransactionCount(ctx)

	// Early exit if we have less than 10 lottery tx in the block
	if lotteryTransactionCount < 10 {
		return nil
	}

	lottery, found := k.GetLottery(ctx, strconv.FormatUint(currentLotteryId.LotteryId, 10))
	if !found {
		panic(fmt.Sprintf("lottery %d does not exist", currentLotteryId.LotteryId))
	}

	var nextLottery types.Lottery

	highestBetFound, _, highestBetAddress := k.LotteryTransactionMetadata.GetMaxBet()
	lowestBetFound, _, lowestBetAddress := k.LotteryTransactionMetadata.GetMinBet()

	// If winner placed the lowest bet, no payment is issued, current lottery pool is carried over
	if lowestBetFound == true && winner.String() == lowestBetAddress {
		nextLottery.Pool = lottery.Pool
	} else {
		var paymentAmount sdk.Coin

		// If the winner placed the highest bet, the entire pool is paid to the winner
		if highestBetFound == true && winner.String() == highestBetAddress {
			paymentAmount = lottery.Pool
		} else {
			// Winner did not place highest or lowest bet, the winner is paid the sum of all bets (without fees)
			paymentAmount = k.LotteryTransactionMetadata.GetBetSum()
		}
		nextLottery.Pool = types.Pool

		// Issue payment
		err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, winner, sdk.Coins{paymentAmount})
		if err != nil {
			return err
		}
	}

	lottery.Winner = winner.String()
	k.SetLottery(ctx, lottery)

	// Prune in-memory data structure
	k.LotteryTransactionMetadata.Prune()

	// next active lottery id
	nextLotteryId := k.IncrementActiveLottery(ctx)

	// Remove prev lotteries
	k.PruneLotteryTransactions(ctx, nextLotteryId)

	// Reset counter
	k.SetLotteryTransactionCount(ctx, 0)

	// set new lottery with the new incremented active lottery
	nextLottery.Index = strconv.FormatUint(nextLotteryId, 10)
	nextLottery.Fee = types.Fee
	k.SetLottery(ctx, nextLottery)

	return nil
}
