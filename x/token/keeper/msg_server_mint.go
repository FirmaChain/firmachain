package keeper

import (
	"context"
	"fmt"
	"strconv"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/firmachain/firmachain/v05/x/token/types"
)

func (ms msgServer) Mint(goCtx context.Context, msg *types.MsgMint) (*types.MsgMintResponse, error) {

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	tokenData, isFound := ms.keeper.GetTokenData(
		ctx,
		msg.TokenID,
	)

	if !isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	if !tokenData.Mintable {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "mint is not allowed.")
	}

	err := ms.keeper.CheckCommonError(tokenData.TokenID, tokenData.Symbol, tokenData.Name, tokenData.TotalSupply)

	if err != nil {
		return nil, err
	}

	// Checks if the the msg owner is the same as the current owner
	if msg.Owner != tokenData.Owner {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	// bank module
	newTotal := msg.Amount + tokenData.TotalSupply

	if newTotal > maxTokenValue {
		errStr := fmt.Sprintf("TotalSupply cannot exceed  %d", maxTokenValue)
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, errStr)
	}

	newCoin := sdk.NewInt64Coin(msg.TokenID, int64(msg.Amount))
	err = ms.keeper.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins(newCoin))

	if err != nil {
		return nil, err
	}

	// send minted coins to receiver
	receiver, err := sdk.AccAddressFromBech32(msg.ToAddress)
	if err != nil {
		return nil, err
	}

	err = ms.keeper.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, receiver, sdk.NewCoins(newCoin))

	if err != nil {
		return nil, err
	}

	tokenData.MintSequence++
	tokenData.TotalSupply += msg.Amount

	ms.keeper.SetTokenData(ctx, tokenData)

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		sdk.EventTypeMessage,
		sdk.NewAttribute("Owner", msg.Owner),
		sdk.NewAttribute("TokenID", msg.TokenID),
		sdk.NewAttribute("MintAmount", strconv.FormatUint(msg.Amount, 10)),
		sdk.NewAttribute("TotalSupply", strconv.FormatUint(tokenData.TotalSupply, 10)),
	))

	return &types.MsgMintResponse{}, nil
}
