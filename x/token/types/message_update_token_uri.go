package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgUpdateTokenUri{}

func NewMsgUpdateTokenUri(owner string, tokenId string, tokenUri string) *MsgUpdateTokenUri {
	return &MsgUpdateTokenUri{
		Owner:    owner,
		TokenId:  tokenId,
		TokenUri: tokenUri,
	}
}

// SDK 0.50: ValidateBasic is no more required to fullfil the sdg.Msg interface implementation.
// The msg's validation is recommended to be performed directly in the msg server and not in the cli command's RunE.
// We still keep it to wrap the basic stateless checks and use it directly in the msg server.
func (msg *MsgUpdateTokenUri) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", err)
	}
	return nil
}
