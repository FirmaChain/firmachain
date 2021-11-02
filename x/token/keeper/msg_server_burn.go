package keeper

import (
	"context"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/firmachain/firmachain/x/token/types"
)

func (k msgServer) Burn(goCtx context.Context, msg *types.MsgBurn) (*types.MsgBurnResponse, error) {
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

	// bank module

	newCoin := sdk.NewInt64Coin(msg.TokenID, int64(msg.Amount))

	// send minted coins to receiver
	receiver, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return nil, err
	}

	err = k.bankKeeper.SendCoinsFromAccountToModule(
		ctx,
		receiver,
		types.ModuleName,
		sdk.NewCoins(newCoin),
	)

	if err != nil {
		return nil, err
	}

	err = k.bankKeeper.BurnCoins(ctx, types.ModuleName, sdk.NewCoins(newCoin))

	if err != nil {
		return nil, err
	}

	valFound.BurnSequence++
	valFound.TotalSupply -= msg.Amount

	k.SetTokenData(ctx, valFound)

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		sdk.EventTypeMessage,
		sdk.NewAttribute("Owner", msg.Owner),
		sdk.NewAttribute("TokenID", msg.TokenID),
		sdk.NewAttribute("BurnAmount", strconv.FormatUint(msg.Amount, 10)),
		sdk.NewAttribute("TotalSupply", strconv.FormatUint(valFound.TotalSupply, 10)),
	))

	return &types.MsgBurnResponse{}, nil
}
