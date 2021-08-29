package keeper

import (
	"github.com/firmachain/firmachain/x/nft/types"
)

var _ types.QueryServer = Keeper{}
