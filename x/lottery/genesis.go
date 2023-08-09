package lottery

import (
	"cosmos-lottery/x/lottery/keeper"
	"cosmos-lottery/x/lottery/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set if defined
	if genState.ActiveLottery != nil {
		k.SetActiveLottery(ctx, *genState.ActiveLottery)
	}
	// Set all the lotteryTransaction
	for _, elem := range genState.LotteryTransactionList {
		k.SetLotteryTransaction(ctx, elem)
	}

	// Set lotteryTransaction count
	k.SetLotteryTransactionCount(ctx, genState.LotteryTransactionCount)
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	// Get all activeLottery
	activeLottery, found := k.GetActiveLottery(ctx)
	if found {
		genesis.ActiveLottery = &activeLottery
	}
	genesis.LotteryTransactionList = k.GetAllLotteryTransaction(ctx)
	genesis.LotteryTransactionCount = k.GetLotteryTransactionCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}