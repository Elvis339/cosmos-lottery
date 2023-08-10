package keeper_test

import (
	"cosmos-lottery/testutil/sample"
	"cosmos-lottery/x/lottery"
	"strconv"
	"testing"

	keepertest "cosmos-lottery/testutil/keeper"
	"cosmos-lottery/testutil/nullify"
	"cosmos-lottery/x/lottery/keeper"
	"cosmos-lottery/x/lottery/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNLottery(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Lottery {
	items := make([]types.Lottery, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetLottery(ctx, items[i])
	}
	return items
}

func TestLotteryGet(t *testing.T) {
	keeper, ctx := keepertest.LotteryKeeper(t)
	items := createNLottery(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetLottery(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestLotteryRemove(t *testing.T) {
	keeper, ctx := keepertest.LotteryKeeper(t)
	items := createNLottery(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveLottery(ctx,
			item.Index,
		)
		_, found := keeper.GetLottery(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestLotteryGetAll(t *testing.T) {
	keeper, ctx := keepertest.LotteryKeeper(t)
	items := createNLottery(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllLottery(ctx)),
	)
}
func TestUpdateLottery(t *testing.T) {
	tests := []struct {
		desc       string
		index      string
		lotteryTxs []types.LotteryTransaction
		pool       uint64
		valid      bool
	}{
		{
			desc:       "should error if lottery does not exist",
			index:      "10",
			lotteryTxs: make([]types.LotteryTransaction, 0),
			pool:       0,
			valid:      false,
		},
		{
			desc:  "should update lottery pool",
			index: "1",
			lotteryTxs: []types.LotteryTransaction{{
				Id:        0,
				Bet:       sdk.NewInt64Coin("token", 2),
				CreatedBy: sample.AccAddress(),
				LotteryId: 1,
			}, {
				Id:        1,
				Bet:       sdk.NewInt64Coin("token", 6),
				CreatedBy: sample.AccAddress(),
				LotteryId: 1,
			}},
			pool:  20,
			valid: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.LotteryKeeper(t)
			lottery.InitGenesis(ctx, *k, types.GenesisState{
				ActiveLottery: types.ActiveLottery{
					LotteryId: 1,
				},
				LotteryList: []types.Lottery{
					{
						Index: "1",
						Fee:   types.Fee,
						Pool:  types.Fee,
					},
				},
			})

			err := k.UpdateLotteryPool(ctx, tc.index, tc.lotteryTxs)
			if tc.valid {
				require.NoError(t, err)

				lottery, _ := k.GetLottery(ctx, tc.index)
				require.Equal(t, lottery.Pool.Amount.Uint64(), tc.pool)
			} else {
				require.Error(t, err)
			}
		})
	}
}
