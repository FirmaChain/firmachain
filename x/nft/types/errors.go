package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrTokenNotFound = sdkerrors.Register(ModuleName, 201, "token not found")
	ErrInvalidHash   = sdkerrors.Register(ModuleName, 202, "invalid hash")
	ErrExistedHash   = sdkerrors.Register(ModuleName, 203, "existed hash")
	ErrNotOwnerToken = sdkerrors.Register(ModuleName, 204, "not owner of token")
)
