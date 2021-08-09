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
		case MsgMintNFT:
			return handleMsgMintNFT(ctx, keeper, msg)
		case MsgDelegateMintNFT:
			return handleMsgDelegateMintNFT(ctx, keeper, msg)
		case MsgBurnNFT:
			return handleMsgBurnNFT(ctx, keeper, msg)
		case MsgTransferNFT:
			return handleMsgTransferNFT(ctx, keeper, msg)
		case MsgMultiTransferNFT:
			return handleMsgMultiTransferNFT(ctx, keeper, msg)
		default:
			errMsg := fmt.Sprintf("Unrecognized contract Msg type: %v", msg.Type())
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}
func handleMsgMintNFT(ctx sdk.Context, keeper Keeper, msg MsgMintNFT) (*sdk.Result, error) {
	err := keeper.Mint(ctx, msg.Hash, msg.TokenURI, msg.Owner, msg.Description, msg.Image)

	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		sdk.EventTypeMessage,
		sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory)))

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func handleMsgDelegateMintNFT(ctx sdk.Context, keeper Keeper, msg MsgDelegateMintNFT) (*sdk.Result, error) {
	err := keeper.Mint(ctx, msg.Hash, msg.TokenURI, msg.Owner, msg.Description, msg.Image)

	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		sdk.EventTypeMessage,
		sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory)))

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func handleMsgBurnNFT(ctx sdk.Context, keeper Keeper, msg MsgBurnNFT) (*sdk.Result, error) {
	err := keeper.Burn(ctx, msg.Hash, msg.Owner)

	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		sdk.EventTypeMessage,
		sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory)))

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func handleMsgTransferNFT(ctx sdk.Context, keeper Keeper, msg MsgTransferNFT) (*sdk.Result, error) {
	err := keeper.Transfer(ctx, msg.Hash, msg.Owner, msg.Recipient)

	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		sdk.EventTypeMessage,
		sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory)))

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func handleMsgMultiTransferNFT(ctx sdk.Context, keeper Keeper, msg MsgMultiTransferNFT) (*sdk.Result, error) {
	err := keeper.MultiTransfer(ctx, msg.Owner, msg.Outputs)

	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		sdk.EventTypeMessage,
		sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory)))

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
