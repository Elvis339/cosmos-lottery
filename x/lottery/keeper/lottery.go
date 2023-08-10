package keeper

import (
	"cosmos-lottery/x/lottery/types"
	"fmt"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
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
func (k Keeper) UpdateLotteryPool(ctx sdk.Context, index string, lotteryTxs []types.LotteryTransaction) error {
	lottery, found := k.GetLottery(ctx, index)

	if !found {
		return sdkerrors.ErrNotFound.Wrapf(fmt.Sprintf("lottery with index %s", index))
	}

	sum := sdk.NewInt64Coin("token", 0)
	for _, lotteryTx := range lotteryTxs {
		sum = sum.Add(lotteryTx.Bet).Add(types.Fee).Add(types.MinBet)
	}
	lottery.Pool = sum
	k.SetLottery(ctx, lottery)

	return nil
}
