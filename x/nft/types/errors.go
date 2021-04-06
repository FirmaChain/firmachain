package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrTokenDoesNotExist = sdkerrors.Register(ModuleName, 201, "Token does not exist.")
	ErrNFTokenInvalid    = sdkerrors.Register(ModuleName, 202, "The hash is invalid or manipulated.")
	ErrNFTokenDuplicated = sdkerrors.Register(ModuleName, 203, "Existed hash.")
)
