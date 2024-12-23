package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/firmachain/firmachain/v05/x/nft/types"
)

func (k msgServer) Transfer(goCtx context.Context, msg *types.MsgTransfer) (*types.MsgTransferResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.HasNftItem(ctx, msg.NftId) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.NftId))
	}

	item := k.GetNftItem(ctx, msg.NftId)

	if msg.Owner != item.Owner {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	oldOwner := item.Owner
	newOwnder := msg.ToAddress

	item.Owner = newOwnder

	k.SetNftItem(ctx, item)
	k.RemoveNftItemToAccount(ctx, oldOwner, msg.NftId)

	return &types.MsgTransferResponse{}, nil
}
