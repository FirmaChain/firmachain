package contract

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case MsgAddContract:
			return HandleMsgAddContract(ctx, keeper, msg)
		default:
			errMsg := fmt.Sprintf("Unrecognized contract Msg type: %v", msg.Type())
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}
func HandleMsgAddContract(ctx sdk.Context, keeper Keeper, msg MsgAddContract) sdk.Result {
	keeper.SetContract(ctx, msg.Path, msg.Hash, msg.Owner)
	return sdk.Result{}
}
