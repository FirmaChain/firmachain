package keeper

import (
	"context"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/firmachain/firmachain/v05/x/token/types"
)

func (k msgServer) CreateToken(goCtx context.Context, msg *types.MsgCreateToken) (*types.MsgCreateTokenResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// ex) GAMETOKEN -> ugametoken
	tokenID := "u" + strings.ToLower(msg.Symbol)

	err := k.CheckCommonError(tokenID, msg.Symbol, msg.Name, msg.TotalSupply)

	if err != nil {
		return nil, err
	}

	_, isFound := k.GetTokenData(ctx, tokenID)

	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	// send minted coins to receiver
	receiver, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return nil, err
	}

	const mintSequence = 1
	const burnSequence = 0

	var tokenData = types.TokenData{
		Owner:        msg.Owner,
		TokenID:      tokenID,
		Name:         msg.Name,
		Symbol:       msg.Symbol,
		TokenURI:     msg.TokenURI,
		TotalSupply:  msg.TotalSupply,
		Decimal:      msg.Decimal,
		Mintable:     msg.Mintable,
		Burnable:     msg.Burnable,
		MintSequence: mintSequence,
		BurnSequence: burnSequence,
	}

	k.SetTokenData(ctx, tokenData)

	// bank module
	// mint
	newCoin := sdk.NewInt64Coin(tokenID, int64(msg.TotalSupply))
	err = k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins(newCoin))

	if err != nil {
		return nil, err
	}

	// transfer to account
	err = k.bankKeeper.SendCoinsFromModuleToAccount(
		ctx,
		types.ModuleName,
		receiver,
		sdk.NewCoins(newCoin),
	)

	if err != nil {
		return nil, err
	}

	// write tokenID info to transaction event log
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		sdk.EventTypeMessage,
		sdk.NewAttribute("Owner", msg.Owner),
		sdk.NewAttribute("TokenID", tokenID),
		sdk.NewAttribute("TokenName", msg.Name),
		sdk.NewAttribute("TokenSymbol", msg.Symbol),
	))

	return &types.MsgCreateTokenResponse{}, nil
}
