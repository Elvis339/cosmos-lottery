package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		ActiveLottery:          nil,
		LotteryTransactionList: []LotteryTransaction{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in lotteryTransaction
	lotteryTransactionIdMap := make(map[uint64]bool)
	lotteryTransactionCount := gs.GetLotteryTransactionCount()
	for _, elem := range gs.LotteryTransactionList {
		if _, ok := lotteryTransactionIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for lotteryTransaction")
		}
		if elem.Id >= lotteryTransactionCount {
			return fmt.Errorf("lotteryTransaction id should be lower or equal than the last id")
		}
		lotteryTransactionIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
