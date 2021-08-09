package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

type MsgMintNFT struct {
	Hash     string         `json:"hash"`
	TokenURI string         `json:"tokenURI"`
	Owner    sdk.AccAddress `json:"owner"`

	Description string `json:"description"`
	Image       string `json:"image"`
}

func NewMsgMintNFT(hash string, tokenURI string, owner sdk.AccAddress, description string, image string) MsgMintNFT {
	return MsgMintNFT{
		Hash:        hash,
		TokenURI:    tokenURI,
		Owner:       owner,
		Description: description,
		Image:       image,
	}
}

func (msg MsgMintNFT) Route() string { return RouterKey }

func (msg MsgMintNFT) Type() string { return "mint_nft" }

func (msg MsgMintNFT) ValidateBasic() error {
	if msg.Owner.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Owner.String())
	}
	if len(msg.Hash) == 0 || len(msg.TokenURI) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "Hash or TokenURI cannot be empty")
	}

	return nil
}

func (msg MsgMintNFT) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg MsgMintNFT) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}

type MsgDelegateMintNFT struct {
	Hash     string         `json:"hash"`
	TokenURI string         `json:"tokenURI"`
	Minter   sdk.AccAddress `json:"minter"`
	Owner    sdk.AccAddress `json:"owner"`

	Description string `json:"description"`
	Image       string `json:"image"`
}

func NewMsgDelegateMintNFT(hash string, tokenURI string, minter sdk.AccAddress, owner sdk.AccAddress, description string, image string) MsgDelegateMintNFT {
	return MsgDelegateMintNFT{
		Hash:        hash,
		TokenURI:    tokenURI,
		Minter:      minter,
		Owner:       owner,
		Description: description,
		Image:       image,
	}
}

func (msg MsgDelegateMintNFT) Route() string { return RouterKey }

func (msg MsgDelegateMintNFT) Type() string { return "delegate_mint_nft" }

func (msg MsgDelegateMintNFT) ValidateBasic() error {
	if msg.Minter.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Minter.String())
	}
	if msg.Owner.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Owner.String())
	}
	if len(msg.Hash) == 0 || len(msg.TokenURI) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "Hash or TokenURI cannot be empty")
	}

	return nil
}

func (msg MsgDelegateMintNFT) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg MsgDelegateMintNFT) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Minter}
}

type MsgBurnNFT struct {
	Hash  string         `json:"hash"`
	Owner sdk.AccAddress `json:"owner"`
}

func NewMsgBurnNFT(hash string, owner sdk.AccAddress) MsgBurnNFT {
	return MsgBurnNFT{
		Hash:  hash,
		Owner: owner,
	}
}

func (msg MsgBurnNFT) Route() string { return RouterKey }

func (msg MsgBurnNFT) Type() string { return "burn_nft" }

func (msg MsgBurnNFT) ValidateBasic() error {
	if msg.Owner.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Owner.String())
	}
	if len(msg.Hash) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "Hash cannot be empty")
	}

	return nil
}

func (msg MsgBurnNFT) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg MsgBurnNFT) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}

type MsgTransferNFT struct {
	Hash      string         `json:"hash"`
	Owner     sdk.AccAddress `json:"owner"`
	Recipient sdk.AccAddress `json:"recipient"`
}

func NewMsgTransferNFT(hash string, owner sdk.AccAddress, recipient sdk.AccAddress) MsgTransferNFT {
	return MsgTransferNFT{
		Hash:      hash,
		Owner:     owner,
		Recipient: recipient,
	}
}

func (msg MsgTransferNFT) Route() string { return RouterKey }

func (msg MsgTransferNFT) Type() string { return "transfer_nft" }

func (msg MsgTransferNFT) ValidateBasic() error {
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

func (msg MsgTransferNFT) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg MsgTransferNFT) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}

type MsgMultiTransferNFT struct {
	Owner   sdk.AccAddress `json:"owner"`
	Outputs []Output       `json:"outputs"`
}

func NewMsgMultiTransferNFT(owner sdk.AccAddress, outputs []Output) MsgMultiTransferNFT {
	return MsgMultiTransferNFT{
		Owner:   owner,
		Outputs: outputs,
	}
}

func (msg MsgMultiTransferNFT) Route() string { return RouterKey }

func (msg MsgMultiTransferNFT) Type() string { return "multi_transfer_nft" }

func (msg MsgMultiTransferNFT) ValidateBasic() error {
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

func (msg MsgMultiTransferNFT) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg MsgMultiTransferNFT) GetSigners() []sdk.AccAddress {
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
