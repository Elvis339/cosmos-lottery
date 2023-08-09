package lottery

import (
	"cosmos-lottery/x/lottery/keeper"
	"cosmos-lottery/x/lottery/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	k.SetParams(ctx, genState.Params)
	k.SetActiveLottery(ctx, genState.ActiveLottery)

	// Set all the lottery
	for _, elem := range genState.LotteryList {
		k.SetLottery(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	// Get all activeLottery
	activeLottery, found := k.GetActiveLottery(ctx)
	if found {
		genesis.ActiveLottery = activeLottery
	}
	genesis.LotteryList = k.GetAllLottery(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
