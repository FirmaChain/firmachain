package keeper

import (
	"context"
	"encoding/binary"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/firmachain/firmachain/x/nft/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/cosmos/cosmos-sdk/store/prefix"
)

func (k Keeper) TokenOfOwnerByIndex(goCtx context.Context, req *types.QueryTokenOfOwnerByIndexRequest) (*types.QueryTokenOfOwnerByIndexResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NftItemAccountMapKey))
	accountStore := prefix.NewStore(store, []byte(req.OwnerAddress))

	iterator := accountStore.Iterator(nil, nil)
	defer iterator.Close()

	var totalBalance uint64 = 0

	for ; iterator.Valid(); iterator.Next() {
		if totalBalance == req.Index {

			nftId := GetNftItemIDFromBytesTemp(iterator.Key())
			return &types.QueryTokenOfOwnerByIndexResponse{TokenId: nftId}, nil
		}

		totalBalance++
	}

	return nil, status.Error(codes.InvalidArgument, "no valid index data")
}

func GetNftItemIDFromBytesTemp(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
