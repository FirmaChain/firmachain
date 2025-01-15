package keeper

import (
	"context"
	"fmt"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/firmachain/firmachain/v05/x/nft/types"
)

func (ms msgServer) Burn(goCtx context.Context, msg *types.MsgBurn) (*types.MsgBurnResponse, error) {

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	if !ms.keeper.HasNftItem(ctx, msg.NftId) {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.NftId))
	}
	if msg.Owner != ms.keeper.GetNftItemOwner(ctx, msg.NftId) {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	ms.keeper.RemoveNftItem(ctx, msg.NftId)

	ms.keeper.RemoveNftItemToAccount(ctx, msg.Owner, msg.NftId)

	return &types.MsgBurnResponse{}, nil
}
