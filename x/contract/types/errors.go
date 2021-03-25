package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrContractDoesNotExist = sdkerrors.Register(ModuleName, 101, "Contract does not exist")
	ErrContractInvalid      = sdkerrors.Register(ModuleName, 102, "Contract hash invalid or manipulated.")
	ErrContractDuplicated   = sdkerrors.Register(ModuleName, 103, "Duplicate contract")
)
