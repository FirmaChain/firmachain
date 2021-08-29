package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/firmachain/firmachain/x/nft/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/cosmos/cosmos-sdk/store/prefix"
)

func (k Keeper) BalanceOf(goCtx context.Context, req *types.QueryBalanceOfRequest) (*types.QueryBalanceOfResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NftItemAccountMapKey))
	accountStore := prefix.NewStore(store, []byte(req.OwnerAddress))

	iterator := accountStore.Iterator(nil, nil)
	defer iterator.Close()

	totalBalance := 0

	for ; iterator.Valid(); iterator.Next() {
		totalBalance++
	}

	return &types.QueryBalanceOfResponse{Total: uint64(totalBalance)}, nil
}
