package keeper

import (
	"context"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/firmachain/firmachain/v05/x/nft/types"
)

func (ms msgServer) Mint(goCtx context.Context, msg *types.MsgMint) (*types.MsgMintResponse, error) {

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	var nftItem = types.NftItem{
		Owner:    msg.Owner,
		TokenUri: msg.TokenUri,
	}

	id := ms.keeper.AppendNftItem(
		ctx,
		nftItem,
	)

	//k.AddNftItemToAccount(ctx, msg.Owner, id)

	// write nftID info to transaction event log
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		sdk.EventTypeMessage,
		sdk.NewAttribute("Owner", msg.Owner),
		sdk.NewAttribute("nftID", strconv.FormatUint(id, 10)),
	))

	return &types.MsgMintResponse{NftId: id}, nil
}
