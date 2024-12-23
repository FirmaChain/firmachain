package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/firmachain/firmachain/v05/x/nft/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/cosmos/cosmos-sdk/store/prefix"
)

func (k Keeper) BalanceOf(goCtx context.Context, req *types.QueryBalanceOfRequest) (*types.QueryBalanceOfResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OwnerOfNftTotalKey))
	accountStore := prefix.NewStore(store, []byte(req.OwnerAddress))

	fmt.Println("0")
	byteKey := types.KeyPrefix(types.OwnerOfNftTotalKey)
	byteTotal := accountStore.Get(byteKey)

	// Count doesn't exist: no element
	if byteTotal == nil {
		return &types.QueryBalanceOfResponse{Total: uint64(0)}, nil
	}

	count := GetUInt64FromBytes(byteTotal)
	return &types.QueryBalanceOfResponse{Total: uint64(count)}, nil
}
