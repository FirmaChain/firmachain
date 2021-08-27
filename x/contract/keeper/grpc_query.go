package keeper

import (
	"github.com/firmachain/firmachain/x/contract/types"
)

var _ types.QueryServer = Keeper{}
