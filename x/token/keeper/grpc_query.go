package keeper

import (
	"github.com/firmachain/firmachain/x/token/types"
)

var _ types.QueryServer = Keeper{}
