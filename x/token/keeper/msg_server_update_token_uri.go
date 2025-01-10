package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/firmachain/firmachain/v05/x/token/types"
)

func (ms msgServer) UpdateTokenURI(goCtx context.Context, msg *types.MsgUpdateTokenURI) (*types.MsgUpdateTokenURIResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	tokenData, isFound := ms.keeper.GetTokenData(
		ctx,
		msg.TokenID,
	)

	if !isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	err := ms.keeper.CheckCommonError(tokenData.TokenID, tokenData.Symbol, tokenData.Name, tokenData.TotalSupply)

	if err != nil {
		return nil, err
	}

	// Checks if the the msg owner is the same as the current owner
	if msg.Owner != tokenData.Owner {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	tokenData.TokenURI = msg.TokenURI

	ms.keeper.SetTokenData(ctx, tokenData)

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		sdk.EventTypeMessage,
		sdk.NewAttribute("Owner", msg.Owner),
		sdk.NewAttribute("TokenID", msg.TokenID),
		sdk.NewAttribute("TokenURI", tokenData.TokenURI),
	))

	return &types.MsgUpdateTokenURIResponse{}, nil
}
