package keeper

import (
	"context"

	"cosmossdk.io/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/firmachain/firmachain/v05/x/nft/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) NftItemAll(c context.Context, req *types.NftItemAllRequest) (*types.NftItemAllResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var nftItems []*types.NftItem
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	nftItemStore := prefix.NewStore(store, types.KeyPrefix(types.NftItemDataKey))

	pageRes, err := query.Paginate(nftItemStore, req.Pagination, func(key []byte, value []byte) error {
		var nftItem types.NftItem
		if err := k.cdc.Unmarshal(value, &nftItem); err != nil {
			return err
		}

		nftItems = append(nftItems, &nftItem)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.NftItemAllResponse{NftItem: nftItems, Pagination: pageRes}, nil
}

func (k Keeper) NftItem(c context.Context, req *types.NftItemRequest) (*types.NftItemResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var nftItem types.NftItem
	ctx := sdk.UnwrapSDKContext(c)

	if !k.HasNftItem(ctx, req.Id) {
		return nil, sdkerrors.ErrKeyNotFound
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NftItemDataKey))
	k.cdc.MustUnmarshal(store.Get(GetBytesFromUInt64(req.Id)), &nftItem)

	return &types.NftItemResponse{NftItem: &nftItem}, nil
}
