package keeper

import (
	"github.com/firmachain/firmachain/v5/x/nft/types"
)

var _ types.QueryServer = Keeper{}
