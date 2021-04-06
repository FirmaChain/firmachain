package nft

import (
	"fmt"
	"github.com/firmachain/FirmaChain/x/nft/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		switch msg := msg.(type) {
		case MsgAddNFToken:
			return handleMsgAddNFToken(ctx, keeper, msg)
		case MsgTransferNFToken:
			return handleMsgTransferNFToken(ctx, keeper, msg)
		default:
			errMsg := fmt.Sprintf("Unrecognized contract Msg type: %v", msg.Type())
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}
func handleMsgAddNFToken(ctx sdk.Context, keeper Keeper, msg MsgAddNFToken) (*sdk.Result, error) {
	err := keeper.AddNFToken(ctx, msg.Hash, msg.TokenURI, msg.Owner)

	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		sdk.EventTypeMessage,
		sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		sdk.NewAttribute(sdk.AttributeKeySender, msg.Owner.String())))

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func handleMsgTransferNFToken(ctx sdk.Context, keeper Keeper, msg MsgTransferNFToken) (*sdk.Result, error) {
	err := keeper.TransferNFToken(ctx, msg.Hash, msg.Owner, msg.Recipient)

	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		sdk.EventTypeMessage,
		sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		sdk.NewAttribute(sdk.AttributeKeySender, msg.Owner.String())))

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
