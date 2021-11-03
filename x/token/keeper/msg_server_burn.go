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
	tokenData, isFound := k.GetTokenData(
		ctx,
		msg.TokenID,
	)

	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	err := k.CheckCommonError(tokenData.TokenID, tokenData.Symbol, tokenData.Name, tokenData.TotalSupply)

	if err != nil {
		return nil, err
	}

	// Checks if the the msg owner is the same as the current owner
	if msg.Owner != tokenData.Owner {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	// send minted coins to ownerAccAddress
	ownerAccAddress, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return nil, err
	}

	balance := k.bankKeeper.GetBalance(ctx, ownerAccAddress, msg.TokenID).Amount

	if balance.Uint64() < msg.Amount {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "account balance is not enough to burn")
	}

	newCoin := sdk.NewInt64Coin(msg.TokenID, int64(msg.Amount))

	err = k.bankKeeper.SendCoinsFromAccountToModule(
		ctx,
		ownerAccAddress,
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

	tokenData.BurnSequence++
	tokenData.TotalSupply -= msg.Amount

	k.SetTokenData(ctx, tokenData)

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		sdk.EventTypeMessage,
		sdk.NewAttribute("Owner", msg.Owner),
		sdk.NewAttribute("TokenID", msg.TokenID),
		sdk.NewAttribute("BurnAmount", strconv.FormatUint(msg.Amount, 10)),
		sdk.NewAttribute("TotalSupply", strconv.FormatUint(tokenData.TotalSupply, 10)),
	))

	return &types.MsgBurnResponse{}, nil
}
