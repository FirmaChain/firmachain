package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/firmachain/firmachain/x/contract/types"
)

func TestContractLogMsgServerCreate(t *testing.T) {
	srv, ctx := setupMsgServer(t)
	creator := "A"
	for i := 0; i < 5; i++ {
		resp, err := srv.AddContractLog(ctx, &types.MsgAddContractLog{Creator: creator})
		require.NoError(t, err)
		require.Equal(t, i, int(resp.Id))
	}
}
