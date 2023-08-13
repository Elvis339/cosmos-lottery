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

// AppendLotteryTransaction adds or updates a LotteryTransaction in the store.
// If a LotteryTransaction with the same creator and LotteryId already exists in the store, it updates the existing record handling only the most recent tx.
// Otherwise, it increments the transaction count and appends a new transaction.
// The function returns the current transaction count.
func (k Keeper) AppendLotteryTransaction(ctx sdk.Context, lotteryTransaction types.LotteryTransaction) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LotteryTransactionKey))
	isInMetadata, lotteryTxId := k.LotteryTransactionMetadata.GetLotteryTransactionId(lotteryTransaction.GetCreatedBy())

	count := k.GetLotteryTransactionCount(ctx)
	lotteryTransaction.Id = count

	if isInMetadata {
		val, exists := k.GetLotteryTransaction(ctx, lotteryTxId)

		if exists && val.GetCreatedBy() == lotteryTransaction.GetCreatedBy() && val.LotteryId == lotteryTransaction.LotteryId {
			lotteryTransaction.Id = val.Id
		} else {
			k.LotteryTransactionMetadata.RemoveLotteryTransactionId(lotteryTransaction.GetCreatedBy())
			k.SetLotteryTransactionCount(ctx, count+1)
		}
		k.LotteryTransactionMetadata.Set(lotteryTransaction)
		store.Set(GetLotteryTransactionIDBytes(lotteryTransaction.Id), k.cdc.MustMarshal(&lotteryTransaction))
	} else {
		k.LotteryTransactionMetadata.Set(lotteryTransaction)
		store.Set(GetLotteryTransactionIDBytes(lotteryTransaction.Id), k.cdc.MustMarshal(&lotteryTransaction))
		k.SetLotteryTransactionCount(ctx, count+1)
	}

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

	for index, tx := range allLotteryTx {
		if allLotteryTx[index].LotteryId == prevLotteryId {
			k.RemoveLotteryTransaction(ctx, tx.Id)
		}
	}
}
