package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/firmachain/firmachain/x/contract/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ContractFileAll(c context.Context, req *types.QueryAllContractFileRequest) (*types.QueryAllContractFileResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var contractFiles []*types.ContractFile
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	contractFileStore := prefix.NewStore(store, types.KeyPrefix(types.ContractFileKey))

	pageRes, err := query.Paginate(contractFileStore, req.Pagination, func(key []byte, value []byte) error {
		var contractFile types.ContractFile
		if err := k.cdc.Unmarshal(value, &contractFile); err != nil {
			return err
		}

		contractFiles = append(contractFiles, &contractFile)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllContractFileResponse{ContractFile: contractFiles, Pagination: pageRes}, nil
}

func (k Keeper) ContractFile(c context.Context, req *types.QueryGetContractFileRequest) (*types.QueryGetContractFileResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetContractFile(ctx, req.Index)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryGetContractFileResponse{ContractFile: &val}, nil
}
