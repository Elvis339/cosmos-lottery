package keeper

import (
	"cosmos-lottery/testutil"
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
	"go.uber.org/mock/gomock"
	"strconv"
	"testing"
)

type MockLotteryKeeper struct {
	Ctx           sdk.Context
	LotteryKeeper *keeper.Keeper
	BankKeeper    *testutil.MockBankKeeper
}

// Keep this to support existing interface
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
		nil,
	)

	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, log.NewNopLogger())

	// Initialize params
	k.SetParams(ctx, types.DefaultParams())

	return k, ctx
}

func NewMockLotteryKeeper(t *testing.T) MockLotteryKeeper {
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

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, log.NewNopLogger())

	bankKeeperMock := testutil.NewMockBankKeeper(ctrl)

	k := keeper.NewKeeper(
		cdc,
		storeKey,
		memStoreKey,
		paramsSubspace,
		bankKeeperMock,
	)

	// Initialize params
	k.SetParams(ctx, types.DefaultParams())

	return MockLotteryKeeper{
		Ctx:           ctx,
		LotteryKeeper: k,
		BankKeeper:    bankKeeperMock,
	}
}
func NewMockLotteryWithGenesis(t *testing.T) MockLotteryKeeper {
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
	mock := NewMockLotteryKeeper(t)
	lottery.InitGenesis(mock.Ctx, *mock.LotteryKeeper, genesisState)
	return mock
}
