package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/firmachain/firmachain/v05/x/contract/types"
)

func (ms msgServer) AddContractLog(goCtx context.Context, msg *types.MsgAddContractLog) (*types.MsgAddContractLogResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := ms.keeper.CheckCommonError(msg)

	if err != nil {
		return nil, err
	}

	var contractLog = types.ContractLog{
		Creator:      msg.Creator,
		ContractHash: msg.ContractHash,
		TimeStamp:    msg.TimeStamp,
		EventName:    msg.EventName,
		OwnerAddress: msg.OwnerAddress,
		JsonString:   msg.JsonString,
	}

	id := ms.keeper.AppendContractLog(
		ctx,
		contractLog,
	)

	return &types.MsgAddContractLogResponse{
		Id: id,
	}, nil
}

func (ms Keeper) CheckCommonError(msg *types.MsgAddContractLog) error {
	if len(msg.ContractHash) == 0 {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "input ContractHash lengh cannot be zero.")
	}

	if len(msg.EventName) == 0 {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "input EventName lengh cannot be zero.")
	}

	return nil
}
