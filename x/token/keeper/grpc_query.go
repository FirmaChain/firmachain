package keeper

import (
	"github.com/firmachain/firmachain/v5/x/token/types"
)

var _ types.QueryServer = Keeper{}
