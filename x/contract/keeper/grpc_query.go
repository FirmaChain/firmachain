package keeper

import (
	"github.com/firmachain/firmachain/v5/x/contract/types"
)

var _ types.QueryServer = Keeper{}
