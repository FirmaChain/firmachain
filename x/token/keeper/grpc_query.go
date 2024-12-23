package keeper

import (
	"github.com/firmachain/firmachain/v05/x/token/types"
)

var _ types.QueryServer = Keeper{}
