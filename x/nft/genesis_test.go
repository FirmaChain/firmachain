package nft_test

import (
	"testing"

	keepertest "github.com/firmachain/firmachain/testutil/keeper"
	"github.com/firmachain/firmachain/testutil/nullify"
	"github.com/firmachain/firmachain/x/nft"
	"github.com/firmachain/firmachain/x/nft/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.NftKeeper(t)
	nft.InitGenesis(ctx, *k, genesisState)
	got := nft.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
