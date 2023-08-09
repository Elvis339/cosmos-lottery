package types_test

import (
	"testing"

	"cosmos-lottery/x/lottery/types"
	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	tests := []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{

				ActiveLottery: &types.ActiveLottery{
					LotteryId: 10,
				},
				LotteryTransactionList: []types.LotteryTransaction{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
				},
				LotteryTransactionCount: 2,
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated lotteryTransaction",
			genState: &types.GenesisState{
				LotteryTransactionList: []types.LotteryTransaction{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid lotteryTransaction count",
			genState: &types.GenesisState{
				LotteryTransactionList: []types.LotteryTransaction{
					{
						Id: 1,
					},
				},
				LotteryTransactionCount: 0,
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
