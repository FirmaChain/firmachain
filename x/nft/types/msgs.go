package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

type MsgMint struct {
	Hash     string         `json:"hash"`
	TokenURI string         `json:"tokenURI"`
	Owner    sdk.AccAddress `json:"owner"`
}

func NewMsgMint(hash string, tokenURI string, owner sdk.AccAddress) MsgMint {
	return MsgMint{
		Hash:     hash,
		TokenURI: tokenURI,
		Owner:    owner,
	}
}

func (msg MsgMint) Route() string { return RouterKey }

func (msg MsgMint) Type() string { return "mint_nft" }

func (msg MsgMint) ValidateBasic() error {
	if msg.Owner.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Owner.String())
	}
	if len(msg.Hash) == 0 || len(msg.TokenURI) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "Hash or TokenURI cannot be empty")
	}

	return nil
}

func (msg MsgMint) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg MsgMint) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}

type MsgBurn struct {
	Hash  string         `json:"hash"`
	Owner sdk.AccAddress `json:"owner"`
}

func NewMsgBurn(hash string, owner sdk.AccAddress) MsgBurn {
	return MsgBurn{
		Hash:  hash,
		Owner: owner,
	}
}

func (msg MsgBurn) Route() string { return RouterKey }

func (msg MsgBurn) Type() string { return "burn_nft" }

func (msg MsgBurn) ValidateBasic() error {
	if msg.Owner.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Owner.String())
	}
	if len(msg.Hash) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "Hash cannot be empty")
	}

	return nil
}

func (msg MsgBurn) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg MsgBurn) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}

type MsgTransfer struct {
	Hash      string         `json:"hash"`
	Owner     sdk.AccAddress `json:"owner"`
	Recipient sdk.AccAddress `json:"recipient"`
}

func NewMsgTransfer(hash string, owner sdk.AccAddress, recipient sdk.AccAddress) MsgTransfer {
	return MsgTransfer{
		Hash:      hash,
		Owner:     owner,
		Recipient: recipient,
	}
}

func (msg MsgTransfer) Route() string { return RouterKey }

func (msg MsgTransfer) Type() string { return "transfer_nft" }

func (msg MsgTransfer) ValidateBasic() error {
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

func (msg MsgTransfer) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg MsgTransfer) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}

type MsgMultiTransfer struct {
	Owner   sdk.AccAddress `json:"owner"`
	Outputs []Output       `json:"outputs"`
}

func NewMsgMultiTransfer(owner sdk.AccAddress, outputs []Output) MsgMultiTransfer {
	return MsgMultiTransfer{
		Owner:   owner,
		Outputs: outputs,
	}
}

func (msg MsgMultiTransfer) Route() string { return RouterKey }

func (msg MsgMultiTransfer) Type() string { return "multi_transfer_nft" }

func (msg MsgMultiTransfer) ValidateBasic() error {
	if msg.Owner.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Owner.String())
	}

	if len(msg.Outputs) == 0 {
		return ErrNoOutputs
	}

	for _, out := range msg.Outputs {
		if err := out.ValidateBasic(); err != nil {
			return err
		}
	}

	return nil
}

func (msg MsgMultiTransfer) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg MsgMultiTransfer) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}

// Output models transaction outputs
type Output struct {
	Hash      string         `json:"hash"`
	Recipient sdk.AccAddress `json:"recipient"`
}

// ValidateBasic - validate transaction output
func (out Output) ValidateBasic() error {
	if len(out.Recipient) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "Recipient missing")
	}
	if len(out.Hash) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "Hash cannot be empty")
	}
	return nil
}
