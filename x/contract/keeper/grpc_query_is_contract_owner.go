package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/firmachain/firmachain/v05/x/contract/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) IsContractOwner(goCtx context.Context, req *types.QueryIsContractOwnerRequest) (*types.QueryIsContractOwnerResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetContractFile(ctx, req.FileHash)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	result := contains(val.OwnerList, req.OwnerAddress)

	return &types.QueryIsContractOwnerResponse{Result: result}, nil
}

func contains(s []string, substr string) bool {
	for _, v := range s {
		if v == substr {
			return true
		}
	}
	return false
}
