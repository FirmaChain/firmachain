package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/firmachain/firmachain/x/token/types"
)

func (k msgServer) UpdateTokenURI(goCtx context.Context, msg *types.MsgUpdateTokenURI) (*types.MsgUpdateTokenURIResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetTokenData(
		ctx,
		msg.TokenID,
	)

	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg owner is the same as the current owner
	if msg.Owner != valFound.Owner {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	valFound.TokenURI = msg.TokenURI

	k.SetTokenData(ctx, valFound)

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		sdk.EventTypeMessage,
		sdk.NewAttribute("Owner", msg.Owner),
		sdk.NewAttribute("TokenID", msg.TokenID),
		sdk.NewAttribute("TokenURI", valFound.TokenURI),
	))

	return &types.MsgUpdateTokenURIResponse{}, nil
}
