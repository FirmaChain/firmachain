package keeper

import (
	"context"
	"fmt"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/firmachain/firmachain/v05/x/nft/types"
)

func (ms msgServer) Transfer(goCtx context.Context, msg *types.MsgTransfer) (*types.MsgTransferResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !ms.keeper.HasNftItem(ctx, msg.NftId) {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.NftId))
	}

	item := ms.keeper.GetNftItem(ctx, msg.NftId)

	if msg.Owner != item.Owner {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	oldOwner := item.Owner
	newOwnder := msg.ToAddress

	item.Owner = newOwnder

	ms.keeper.SetNftItem(ctx, item)
	ms.keeper.RemoveNftItemToAccount(ctx, oldOwner, msg.NftId)

	return &types.MsgTransferResponse{}, nil
}
