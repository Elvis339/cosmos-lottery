package keeper

import (
	"cosmos-lottery/x/lottery"
	"cosmos-lottery/x/lottery/keeper"
	"cosmos-lottery/x/lottery/types"
	tmdb "github.com/cometbft/cometbft-db"
	"github.com/cometbft/cometbft/libs/log"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	typesparams "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/stretchr/testify/require"
	"strconv"
	"testing"
)

func LotteryKeeper(t testing.TB) (*keeper.Keeper, sdk.Context) {
	storeKey := sdk.NewKVStoreKey(types.StoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(types.MemStoreKey)

	db := tmdb.NewMemDB()
	stateStore := store.NewCommitMultiStore(db)
	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(memStoreKey, storetypes.StoreTypeMemory, nil)
	require.NoError(t, stateStore.LoadLatestVersion())

	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)

	paramsSubspace := typesparams.NewSubspace(cdc,
		types.Amino,
		storeKey,
		memStoreKey,
		"LotteryParams",
	)
	k := keeper.NewKeeper(
		cdc,
		storeKey,
		memStoreKey,
		paramsSubspace,
		nil, // @TODO: Figure out to mock bankKeper so tests in msg_server_test.go can pass
	)

	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, log.NewNopLogger())

	// Initialize params
	k.SetParams(ctx, types.DefaultParams())

	return k, ctx
}

func SetupLotteryKeeperWithGenesis(t testing.TB) (*keeper.Keeper, sdk.Context) {
	activeLotteryId := uint64(1)
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		ActiveLottery: types.ActiveLottery{
			LotteryId: activeLotteryId,
		},
		LotteryList: []types.Lottery{
			{
				Index: strconv.FormatUint(activeLotteryId, 10),
				Fee:   types.Fee,
				Pool:  types.Pool,
			},
		},
	}

	k, ctx := LotteryKeeper(t)
	lottery.InitGenesis(ctx, *k, genesisState)

	return k, ctx
}
