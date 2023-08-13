package keeper

import (
	"cosmos-lottery/x/lottery/types"
	"encoding/binary"
	"errors"
	"fmt"
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
1. Check if the user's address exists in the in-memory map (LotteryTransactionMetadata).
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
	isInMetadata, lotteryTxId := k.LotteryTransactionMetadata.GetLotteryTransactionId(lotteryTransaction.GetCreatedBy())

	count := k.GetLotteryTransactionCount(ctx)

	lotteryTransaction.Id = count

	if isInMetadata == true {
		oldLotteryTx, oldLotteryTxExist := k.GetLotteryTransaction(ctx, lotteryTxId)
		if oldLotteryTxExist == true {
			if oldLotteryTx.CreatedBy == lotteryTransaction.CreatedBy {
				if oldLotteryTx.LotteryId != lotteryTransaction.LotteryId {
					// Because their lottery id is different `mesamo babe i zabe`
					k.LotteryTransactionMetadata.RemoveLotteryTransactionId(lotteryTransaction.GetCreatedBy())
					k.LotteryTransactionMetadata.Set(lotteryTransaction)
					k.SetLotteryTransactionCount(ctx, count+1)
				} else {
					lotteryTransaction.Id = oldLotteryTx.Id
				}
			}
		}
	} else {
		// Update meta
		k.LotteryTransactionMetadata.Set(lotteryTransaction)

		// Update lotteryTransaction count
		k.SetLotteryTransactionCount(ctx, count+1)
	}

	appendedValue := k.cdc.MustMarshal(&lotteryTransaction)
	store.Set(GetLotteryTransactionIDBytes(lotteryTransaction.Id), appendedValue)

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

func (k Keeper) GetWinner(ctx sdk.Context) (*types.LotteryTransaction, error) {
	all := k.GetAllLotteryTransaction(ctx)
	hash := types.Hash(all)
	numOfTx := len(all)

	if numOfTx == 0 {
		return nil, errors.New("no transactions in the block. Winner cannot be determined")
	}

	result := binary.LittleEndian.Uint16(hash[len(hash)-2:])
	winnerIndex := int(result) % numOfTx

	ltryTx := &all[winnerIndex]

	if ltryTx != nil {
		return ltryTx, nil
	}

	return nil, errors.New(fmt.Sprintf("lottery tx with index=%d does not exist", winnerIndex))
}

func (k Keeper) PruneLotteryTransactions(ctx sdk.Context, activeLottery uint64) {
	allLotteryTx := k.GetAllLotteryTransaction(ctx)

	prevLotteryId := activeLottery - 1

	if prevLotteryId == 0 {
		prevLotteryId = 1
	}

	ctx.Logger().Info(fmt.Sprintf("--------- REMOVING ALL LOTTERY TX WITH LOTTERY ID %d ---------", prevLotteryId))

	for index, tx := range allLotteryTx {
		if allLotteryTx[index].LotteryId == prevLotteryId {
			k.RemoveLotteryTransaction(ctx, tx.Id)
		}
	}
}
