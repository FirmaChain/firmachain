package keeper

import (
	"github.com/firmachain/firmachain/v05/x/contract/types"
)

var _ types.QueryServer = Keeper{}
