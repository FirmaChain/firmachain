package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateToken{}

func NewMsgCreateToken(owner string, name string, symbol string, tokenURI string, totalSupply uint64, decimal uint64, mintable bool, burnable bool) *MsgCreateToken {
	return &MsgCreateToken{
		Owner:       owner,
		Name:        name,
		Symbol:      symbol,
		TokenURI:    tokenURI,
		TotalSupply: totalSupply,
		Decimal:     decimal,
		Mintable:    mintable,
		Burnable:    burnable,
	}
}

func (msg *MsgCreateToken) Route() string {
	return RouterKey
}

func (msg *MsgCreateToken) Type() string {
	return "CreateToken"
}

func (msg *MsgCreateToken) GetSigners() []sdk.AccAddress {
	owner, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{owner}
}

func (msg *MsgCreateToken) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateToken) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", err)
	}
	return nil
}
