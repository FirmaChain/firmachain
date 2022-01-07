package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/firmachain/firmachain/x/contract/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetContractListFromHash(goCtx context.Context, req *types.QueryGetContractListFromHashRequest) (*types.QueryGetContractListFromHashResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ContractLogHashKey))

	hashStore := prefix.NewStore(store, []byte(req.Hash))

	iterator := hashStore.Iterator(nil, nil)
	defer iterator.Close()

	var idList []uint64

	for ; iterator.Valid(); iterator.Next() {
		id := GetUInt64FromBytes(iterator.Key())
		idList = append(idList, id)
	}

	return &types.QueryGetContractListFromHashResponse{IdList: idList}, nil
}
