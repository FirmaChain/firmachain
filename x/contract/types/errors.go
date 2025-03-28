package types

import (
	errorsmod "cosmossdk.io/errors"
)

// x/contract module sentinel errors
var (
	ErrSample = errorsmod.Register(ModuleName, 1100, "sample error")
)
