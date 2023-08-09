package keeper

import (
	"context"

	"cosmos-lottery/x/lottery/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) LotteryTransactionAll(goCtx context.Context, req *types.QueryAllLotteryTransactionRequest) (*types.QueryAllLotteryTransactionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var lotteryTransactions []types.LotteryTransaction
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	lotteryTransactionStore := prefix.NewStore(store, types.KeyPrefix(types.LotteryTransactionKey))

	pageRes, err := query.Paginate(lotteryTransactionStore, req.Pagination, func(key []byte, value []byte) error {
		var lotteryTransaction types.LotteryTransaction
		if err := k.cdc.Unmarshal(value, &lotteryTransaction); err != nil {
			return err
		}

		lotteryTransactions = append(lotteryTransactions, lotteryTransaction)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllLotteryTransactionResponse{LotteryTransaction: lotteryTransactions, Pagination: pageRes}, nil
}

func (k Keeper) LotteryTransaction(goCtx context.Context, req *types.QueryGetLotteryTransactionRequest) (*types.QueryGetLotteryTransactionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	lotteryTransaction, found := k.GetLotteryTransaction(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetLotteryTransactionResponse{LotteryTransaction: lotteryTransaction}, nil
}
