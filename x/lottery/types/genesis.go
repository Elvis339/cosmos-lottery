package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"strconv"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

var (
	Fee  = sdk.NewInt64Coin("token", 1)
	Pool = sdk.NewInt64Coin("token", 0)
)

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		ActiveLottery: ActiveLottery{LotteryId: DefaultIndex},
		LotteryList: []Lottery{
			{
				Index:               strconv.FormatUint(DefaultIndex, 10),
				Fee:                 Fee,
				Pool:                Pool,
				LotteryTransactions: nil,
			},
		},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	lotteryId := gs.ActiveLottery.LotteryId

	// Check for duplicated index in lottery
	lotteryIndexMap := make(map[string]struct{})

	for _, elem := range gs.LotteryList {
		index := string(LotteryKey(elem.Index))
		currentLotteryId, err := strconv.ParseUint(elem.Index, 10, 64)

		if err != nil {
			return err
		}

		if currentLotteryId > lotteryId {
			return fmt.Errorf("lottery have higher id then the current active lottery")
		}

		if _, ok := lotteryIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for lottery")
		}
		lotteryIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
