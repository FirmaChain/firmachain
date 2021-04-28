package nft

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/firmachain/FirmaChain/x/nft/types"
)

func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		switch msg := msg.(type) {
		case MsgMint:
			return handleMsgMint(ctx, keeper, msg)
		case MsgBurn:
			return handleMsgBurn(ctx, keeper, msg)
		case MsgTransfer:
			return handleMsgTransfer(ctx, keeper, msg)
		case MsgMultiTransfer:
			return handleMsgMultiTransfer(ctx, keeper, msg)
		default:
			errMsg := fmt.Sprintf("Unrecognized contract Msg type: %v", msg.Type())
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}
func handleMsgMint(ctx sdk.Context, keeper Keeper, msg MsgMint) (*sdk.Result, error) {
	err := keeper.Mint(ctx, msg.Hash, msg.TokenURI, msg.Owner)

	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		sdk.EventTypeMessage,
		sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory)))

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func handleMsgBurn(ctx sdk.Context, keeper Keeper, msg MsgBurn) (*sdk.Result, error) {
	err := keeper.Burn(ctx, msg.Hash, msg.Owner)

	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		sdk.EventTypeMessage,
		sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory)))

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func handleMsgTransfer(ctx sdk.Context, keeper Keeper, msg MsgTransfer) (*sdk.Result, error) {
	err := keeper.Transfer(ctx, msg.Hash, msg.Owner, msg.Recipient)

	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		sdk.EventTypeMessage,
		sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory)))

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func handleMsgMultiTransfer(ctx sdk.Context, keeper Keeper, msg MsgMultiTransfer) (*sdk.Result, error) {
	err := keeper.MultiTransfer(ctx, msg.Owner, msg.Outputs)

	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		sdk.EventTypeMessage,
		sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory)))

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
