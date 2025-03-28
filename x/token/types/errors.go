package types

import (
	errorsmod "cosmossdk.io/errors"
)

// x/token module sentinel errors
var (
	ErrSample = errorsmod.Register(ModuleName, 1100, "sample error")
)
