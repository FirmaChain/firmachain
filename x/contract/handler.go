package contract

import (
	"fmt"
	"github.com/firmachain/FirmaChain/x/contract/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case MsgAddContract:
			return handleMsgAddContract(ctx, keeper, msg)
		default:
			errMsg := fmt.Sprintf("Unrecognized contract Msg type: %v", msg.Type())
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}
func handleMsgAddContract(ctx sdk.Context, keeper Keeper, msg MsgAddContract) sdk.Result {
	error := keeper.SetContract(ctx, msg.Hash, msg.Path, msg.Owner)

	if error != nil {
		return error.Result()
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		sdk.EventTypeMessage,
		sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		sdk.NewAttribute(sdk.AttributeKeySender, msg.Owner.String())))

	return sdk.Result{Events: ctx.EventManager().Events()}
}
