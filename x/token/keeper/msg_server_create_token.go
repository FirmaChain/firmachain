package keeper

import (
	"context"
	"strings"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/firmachain/firmachain/v05/x/token/types"
)

func (ms msgServer) CreateToken(goCtx context.Context, msg *types.MsgCreateToken) (*types.MsgCreateTokenResponse, error) {

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// ex) GAMETOKEN -> ugametoken
	tokenId := "u" + strings.ToLower(msg.Symbol)

	err := ms.keeper.CheckCommonError(tokenId, msg.Symbol, msg.Name, msg.TotalSupply)

	if err != nil {
		return nil, err
	}

	_, isFound := ms.keeper.GetTokenData(ctx, tokenId)

	if isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
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
		TokenId:      tokenId,
		Name:         msg.Name,
		Symbol:       msg.Symbol,
		TokenUri:     msg.TokenUri,
		TotalSupply:  msg.TotalSupply,
		Decimal:      msg.Decimal,
		Mintable:     msg.Mintable,
		Burnable:     msg.Burnable,
		MintSequence: mintSequence,
		BurnSequence: burnSequence,
	}

	ms.keeper.SetTokenData(ctx, tokenData)

	// bank module
	// mint
	newCoin := sdk.NewInt64Coin(tokenId, int64(msg.TotalSupply))
	err = ms.keeper.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins(newCoin))

	if err != nil {
		return nil, err
	}

	// transfer to account
	err = ms.keeper.bankKeeper.SendCoinsFromModuleToAccount(
		ctx,
		types.ModuleName,
		receiver,
		sdk.NewCoins(newCoin),
	)

	if err != nil {
		return nil, err
	}

	// write tokenId info to transaction event log
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		sdk.EventTypeMessage,
		sdk.NewAttribute("Owner", msg.Owner),
		sdk.NewAttribute("TokenId", tokenId),
		sdk.NewAttribute("TokenName", msg.Name),
		sdk.NewAttribute("TokenSymbol", msg.Symbol),
	))

	return &types.MsgCreateTokenResponse{}, nil
}
