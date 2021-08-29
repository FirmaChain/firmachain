package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/firmachain/firmachain/x/nft/types"
)

func (k msgServer) Mint(goCtx context.Context, msg *types.MsgMint) (*types.MsgMintResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var nftItem = types.NftItem{
		Owner:    msg.Owner,
		TokenURI: msg.TokenURI,
	}

	id := k.AppendNftItem(
		ctx,
		nftItem,
	)

	k.AddNftItemToAccount(ctx, msg.Owner, id)

	return &types.MsgMintResponse{NftId: id}, nil
}
