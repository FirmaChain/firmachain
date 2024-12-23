package keeper

import (
	"github.com/firmachain/firmachain/v05/x/nft/types"
)

var _ types.QueryServer = Keeper{}
