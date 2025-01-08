package keeper

import (
	"context"

	"cosmossdk.io/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/firmachain/firmachain/v05/x/contract/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ContractLogAll(c context.Context, req *types.QueryAllContractLogRequest) (*types.QueryAllContractLogResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var contractLogs []*types.ContractLog
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	contractLogStore := prefix.NewStore(store, types.KeyPrefix(types.ContractLogDataKey))

	pageRes, err := query.Paginate(contractLogStore, req.Pagination, func(key []byte, value []byte) error {
		var contractLog types.ContractLog
		if err := k.cdc.Unmarshal(value, &contractLog); err != nil {
			return err
		}

		contractLogs = append(contractLogs, &contractLog)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllContractLogResponse{ContractLog: contractLogs, Pagination: pageRes}, nil
}

func (k Keeper) ContractLog(c context.Context, req *types.QueryGetContractLogRequest) (*types.QueryGetContractLogResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var contractLog types.ContractLog
	ctx := sdk.UnwrapSDKContext(c)

	if !k.HasContractLog(ctx, req.Id) {
		return nil, sdkerrors.ErrKeyNotFound
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ContractLogDataKey))
	k.cdc.MustUnmarshal(store.Get(GetBytesFromUInt64(req.Id)), &contractLog)

	return &types.QueryGetContractLogResponse{ContractLog: &contractLog}, nil
}
