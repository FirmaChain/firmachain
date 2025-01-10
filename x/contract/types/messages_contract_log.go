package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgAddContractLog{}

func NewMsgAddContractLog(creator string, contractHash string, timeStamp uint64, eventName string, ownerAddress string, jsonString string) *MsgAddContractLog {
	return &MsgAddContractLog{
		Creator:      creator,
		ContractHash: contractHash,
		TimeStamp:    timeStamp,
		EventName:    eventName,
		OwnerAddress: ownerAddress,
		JsonString:   jsonString,
	}
}

func (msg *MsgAddContractLog) Route() string {
	return RouterKey
}

func (msg *MsgAddContractLog) Type() string {
	return "CreateContractLog"
}

func (msg *MsgAddContractLog) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAddContractLog) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAddContractLog) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
