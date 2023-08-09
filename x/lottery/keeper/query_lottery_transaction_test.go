package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "cosmos-lottery/testutil/keeper"
	"cosmos-lottery/testutil/nullify"
	"cosmos-lottery/x/lottery/types"
)

func TestLotteryTransactionQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.LotteryKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNLotteryTransaction(keeper, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetLotteryTransactionRequest
		response *types.QueryGetLotteryTransactionResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetLotteryTransactionRequest{Id: msgs[0].Id},
			response: &types.QueryGetLotteryTransactionResponse{LotteryTransaction: msgs[0]},
		},
		{
			desc:     "Second",
			request:  &types.QueryGetLotteryTransactionRequest{Id: msgs[1].Id},
			response: &types.QueryGetLotteryTransactionResponse{LotteryTransaction: msgs[1]},
		},
		{
			desc:    "KeyNotFound",
			request: &types.QueryGetLotteryTransactionRequest{Id: uint64(len(msgs))},
			err:     sdkerrors.ErrKeyNotFound,
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.LotteryTransaction(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestLotteryTransactionQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.LotteryKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNLotteryTransaction(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllLotteryTransactionRequest {
		return &types.QueryAllLotteryTransactionRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.LotteryTransactionAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.LotteryTransaction), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.LotteryTransaction),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.LotteryTransactionAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.LotteryTransaction), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.LotteryTransaction),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.LotteryTransactionAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.LotteryTransaction),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.LotteryTransactionAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
