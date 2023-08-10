package keeper

import (
	"cosmos-lottery/x/lottery/types"
	"encoding/binary"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetLotteryTransactionCount get the total number of lotteryTransaction
func (k Keeper) GetLotteryTransactionCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.LotteryTransactionCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetLotteryTransactionCount set the total number of lotteryTransaction
func (k Keeper) SetLotteryTransactionCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.LotteryTransactionCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

/*
AppendLotteryTransaction
On a new bet:
1. Check if the user's address exists in the in-memory map (lotteryTxMeta).
2. If the user's address is NOT present:
  - Add the user's address to the hash map.
  - Increment the transaction count.

3. If the user's address IS present:
  - Update existing transaction bet with the new bet details retaining same insertion order.
  - Maintain the same transaction ID.
*/
func (k Keeper) AppendLotteryTransaction(
	ctx sdk.Context,
	lotteryTransaction types.LotteryTransaction,
) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LotteryTransactionKey))
	ltryTx := lotteryTransaction

	count := k.GetLotteryTransactionCount(ctx)
	id, found := k.lotteryTxMeta.Get(lotteryTransaction.GetCreatedBy())

	if found {
		lotteryTx, exist := k.GetLotteryTransaction(ctx, id)
		if exist {
			lotteryTx.Bet = ltryTx.Bet
		}
		ltryTx = lotteryTx
	} else {
		k.lotteryTxMeta.Set(lotteryTransaction.GetCreatedBy(), count)
		ltryTx.Id = count

		// Update lotteryTransaction count
		k.SetLotteryTransactionCount(ctx, count+1)
	}

	appendedValue := k.cdc.MustMarshal(&ltryTx)
	store.Set(GetLotteryTransactionIDBytes(ltryTx.Id), appendedValue)

	return count
}

// SetLotteryTransaction set a specific lotteryTransaction in the store
func (k Keeper) SetLotteryTransaction(ctx sdk.Context, lotteryTransaction types.LotteryTransaction) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LotteryTransactionKey))
	b := k.cdc.MustMarshal(&lotteryTransaction)
	store.Set(GetLotteryTransactionIDBytes(lotteryTransaction.Id), b)
}

// GetLotteryTransaction returns a lotteryTransaction from its id
func (k Keeper) GetLotteryTransaction(ctx sdk.Context, id uint64) (val types.LotteryTransaction, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LotteryTransactionKey))
	b := store.Get(GetLotteryTransactionIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveLotteryTransaction removes a lotteryTransaction from the store
func (k Keeper) RemoveLotteryTransaction(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LotteryTransactionKey))
	store.Delete(GetLotteryTransactionIDBytes(id))
}

// GetAllLotteryTransaction returns all lotteryTransaction
func (k Keeper) GetAllLotteryTransaction(ctx sdk.Context) (list []types.LotteryTransaction) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LotteryTransactionKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.LotteryTransaction
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetLotteryTransactionIDBytes returns the byte representation of the ID
func GetLotteryTransactionIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetLotteryTransactionIDFromBytes returns ID in uint64 format from a byte array
func GetLotteryTransactionIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
