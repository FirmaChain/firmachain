package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	DefaultCodespace sdk.CodespaceType = ModuleName

	CodeContractDoesNotExist sdk.CodeType = 101
	CodeContractInvalid      sdk.CodeType = 102
)

func ErrContractDoesNotExist(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeContractDoesNotExist, "Contract does not exist")
}

func ErrContractInvalid(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeContractInvalid, "Contract hash invalid or manipulated.")
}
