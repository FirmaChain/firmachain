package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/firmachain/firmachain/x/contract/types"
)

func (k msgServer) AddContractLog(goCtx context.Context, msg *types.MsgAddContractLog) (*types.MsgAddContractLogResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var contractLog = types.ContractLog{
		Creator:      msg.Creator,
		ContractHash: msg.ContractHash,
		TimeStamp:    msg.TimeStamp,
		EventName:    msg.EventName,
		JsonString:   msg.JsonString,
	}

	id := k.AppendContractLog(
		ctx,
		contractLog,
	)

	return &types.MsgAddContractLogResponse{
		Id: id,
	}, nil
}
