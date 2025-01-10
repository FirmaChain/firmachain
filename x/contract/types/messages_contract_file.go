package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateContractFile{}

func NewMsgCreateContractFile(creator string, fileHash string, timeStamp uint64, ownerList []string, metaDataJsonString string) *MsgCreateContractFile {
	return &MsgCreateContractFile{
		Creator:            creator,
		FileHash:           fileHash,
		TimeStamp:          timeStamp,
		OwnerList:          ownerList,
		MetaDataJsonString: metaDataJsonString,
	}
}

func (msg *MsgCreateContractFile) Route() string {
	return RouterKey
}

func (msg *MsgCreateContractFile) Type() string {
	return "CreateContractFile"
}

func (msg *MsgCreateContractFile) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateContractFile) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateContractFile) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
