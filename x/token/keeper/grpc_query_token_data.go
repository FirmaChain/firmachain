package keeper

import (
	"context"

	"cosmossdk.io/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/firmachain/firmachain/v05/x/token/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) TokenDataAll(c context.Context, req *types.TokenDataAllRequest) (*types.TokenDataAllResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var tokenDatas []types.TokenData
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	tokenDataStore := prefix.NewStore(store, types.KeyPrefix(types.TokenDataKeyPrefix))

	pageRes, err := query.Paginate(tokenDataStore, req.Pagination, func(key []byte, value []byte) error {
		var tokenData types.TokenData
		if err := k.cdc.Unmarshal(value, &tokenData); err != nil {
			return err
		}

		tokenDatas = append(tokenDatas, tokenData)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.TokenDataAllResponse{TokenData: tokenDatas, Pagination: pageRes}, nil
}

func (k Keeper) TokenData(c context.Context, req *types.TokenDataRequest) (*types.TokenDataResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetTokenData(
		ctx,
		req.TokenId,
	)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.TokenDataResponse{TokenData: val}, nil
}
