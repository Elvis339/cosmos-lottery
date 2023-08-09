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

				ActiveLottery: types.ActiveLottery{
					LotteryId: 2,
				},
				LotteryList: []types.Lottery{
					{
						Index: "1",
					},
					{
						Index: "2",
					},
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated lottery",
			genState: &types.GenesisState{
				LotteryList: []types.Lottery{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "should not have higher id then active",
			genState: &types.GenesisState{

				ActiveLottery: types.ActiveLottery{
					LotteryId: 1,
				},
				LotteryList: []types.Lottery{{Index: "1"}, {Index: "2"}},
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
