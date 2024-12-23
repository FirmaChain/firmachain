package keeper

import (
	"github.com/firmachain/firmachain/v05/x/burn/types"
)

var _ types.QueryServer = Keeper{}
