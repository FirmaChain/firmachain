package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgUpdateTokenURI{}

func NewMsgUpdateTokenURI(owner string, tokenID string, tokenURI string) *MsgUpdateTokenURI {
	return &MsgUpdateTokenURI{
		Owner:    owner,
		TokenID:  tokenID,
		TokenURI: tokenURI,
	}
}

func (msg *MsgUpdateTokenURI) Route() string {
	return RouterKey
}

func (msg *MsgUpdateTokenURI) Type() string {
	return "UpdateTokenURI"
}

func (msg *MsgUpdateTokenURI) GetSigners() []sdk.AccAddress {
	owner, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{owner}
}

func (msg *MsgUpdateTokenURI) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateTokenURI) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", err)
	}
	return nil
}
