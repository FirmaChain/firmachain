package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/firmachain/firmachain/v05/x/token/types"
)

func (ms msgServer) UpdateTokenUri(goCtx context.Context, msg *types.MsgUpdateTokenUri) (*types.MsgUpdateTokenUriResponse, error) {

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	tokenData, isFound := ms.keeper.GetTokenData(
		ctx,
		msg.TokenId,
	)

	if !isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	err := ms.keeper.CheckCommonError(tokenData.TokenId, tokenData.Symbol, tokenData.Name, tokenData.TotalSupply)

	if err != nil {
		return nil, err
	}

	// Checks if the the msg owner is the same as the current owner
	if msg.Owner != tokenData.Owner {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	tokenData.TokenUri = msg.TokenUri

	ms.keeper.SetTokenData(ctx, tokenData)

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		sdk.EventTypeMessage,
		sdk.NewAttribute("Owner", msg.Owner),
		sdk.NewAttribute("TokenId", msg.TokenId),
		sdk.NewAttribute("TokenUri", tokenData.TokenUri),
	))

	return &types.MsgUpdateTokenUriResponse{}, nil
}
