package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/firmachain/firmachain/x/token/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetTokenList(goCtx context.Context, req *types.QueryGetTokenListRequest) (*types.QueryGetTokenListResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TokenDataAccountMapKey))
	accountStore := prefix.NewStore(store, []byte(req.OwnerAddress))

	iterator := accountStore.Iterator(nil, nil)
	defer iterator.Close()

	var tokenDataArray []string

	for ; iterator.Valid(); iterator.Next() {

		// bytes to string
		tokenId := string(iterator.Value()[:])
		tokenDataArray = append(tokenDataArray, tokenId)
	}

	return &types.QueryGetTokenListResponse{TokenID: tokenDataArray}, nil
}
