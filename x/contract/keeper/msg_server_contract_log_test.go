package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/firmachain/firmachain/x/contract/types"
)

func TestContractLogMsgServerCreate(t *testing.T) {
	srv, ctx := setupMsgServer(t)
	creator := "A"
	contractHash := "testContractHash001"
	eventName := "testEvent"

	for i := 0; i < 5; i++ {
		resp, err := srv.AddContractLog(ctx, &types.MsgAddContractLog{Creator: creator, ContractHash: contractHash, EventName: eventName})
		require.NoError(t, err)
		require.Equal(t, i, int(resp.Id))
	}
}
