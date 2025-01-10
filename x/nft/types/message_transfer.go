package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgTransfer{}

func NewMsgTransfer(owner string, nftId uint64, toAddress string) *MsgTransfer {
	return &MsgTransfer{
		Owner:     owner,
		NftId:     nftId,
		ToAddress: toAddress,
	}
}

func (msg *MsgTransfer) Route() string {
	return RouterKey
}

func (msg *MsgTransfer) Type() string {
	return "Transfer"
}

func (msg *MsgTransfer) GetSigners() []sdk.AccAddress {
	owner, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{owner}
}

func (msg *MsgTransfer) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgTransfer) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", err)
	}
	return nil
}
