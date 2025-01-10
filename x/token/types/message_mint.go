package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgMint{}

func NewMsgMint(owner string, tokenID string, amount uint64, toAddress string) *MsgMint {
	return &MsgMint{
		Owner:     owner,
		TokenID:   tokenID,
		Amount:    amount,
		ToAddress: toAddress,
	}
}

func (msg *MsgMint) Route() string {
	return RouterKey
}

func (msg *MsgMint) Type() string {
	return "Mint"
}

func (msg *MsgMint) GetSigners() []sdk.AccAddress {
	owner, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{owner}
}

func (msg *MsgMint) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgMint) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", err)
	}
	return nil
}
