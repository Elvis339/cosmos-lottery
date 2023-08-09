package keeper

import (
	"cosmos-lottery/x/lottery/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetActiveLottery set activeLottery in the store
func (k Keeper) SetActiveLottery(ctx sdk.Context, activeLottery types.ActiveLottery) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ActiveLotteryKey))
	b := k.cdc.MustMarshal(&activeLottery)
	store.Set([]byte{0}, b)
}

// GetActiveLottery returns activeLottery
func (k Keeper) GetActiveLottery(ctx sdk.Context) (val types.ActiveLottery, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ActiveLotteryKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveActiveLottery removes activeLottery from the store
func (k Keeper) RemoveActiveLottery(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ActiveLotteryKey))
	store.Delete([]byte{0})
}
