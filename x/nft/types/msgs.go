package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

type MsgAddNFToken struct {
	Hash     string         `json:"hash"`
	TokenURI string         `json:"tokenURI"`
	Owner    sdk.AccAddress `json:"owner"`
}

func NewMsgAddNFToken(hash string, tokenURI string, owner sdk.AccAddress) MsgAddNFToken {
	return MsgAddNFToken{
		Hash:     hash,
		TokenURI: tokenURI,
		Owner:    owner,
	}
}

func (msg MsgAddNFToken) Route() string { return RouterKey }

func (msg MsgAddNFToken) Type() string { return "add_nft" }

func (msg MsgAddNFToken) ValidateBasic() error {
	if msg.Owner.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Owner.String())
	}
	if len(msg.Hash) == 0 || len(msg.TokenURI) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "Hash or TokenURI cannot be empty")
	}

	return nil
}

func (msg MsgAddNFToken) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg MsgAddNFToken) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}

type MsgTransferNFToken struct {
	Hash      string         `json:"hash"`
	Owner     sdk.AccAddress `json:"owner"`
	Recipient sdk.AccAddress `json:"recipient"`
}

func NewMsgTransferNFToken(hash string, owner sdk.AccAddress, recipient sdk.AccAddress) MsgTransferNFToken {
	return MsgTransferNFToken{
		Hash:      hash,
		Owner:     owner,
		Recipient: recipient,
	}
}

func (msg MsgTransferNFToken) Route() string { return RouterKey }

func (msg MsgTransferNFToken) Type() string { return "transfer_nft" }

func (msg MsgTransferNFToken) ValidateBasic() error {
	if msg.Owner.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Owner.String())
	}

	if msg.Recipient.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Owner.String())
	}

	if len(msg.Hash) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "Hash cannot be empty")
	}

	return nil
}

func (msg MsgTransferNFToken) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg MsgTransferNFToken) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}
