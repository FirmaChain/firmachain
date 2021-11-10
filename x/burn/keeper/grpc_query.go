package keeper

import (
	"github.com/firmachain/firmachain/x/burn/types"
)

var _ types.QueryServer = Keeper{}
