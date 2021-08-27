package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/firmachain/firmachain/x/contract/types"
)

func (k msgServer) CreateContractFile(goCtx context.Context, msg *types.MsgCreateContractFile) (*types.MsgCreateContractFileResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetContractFile(ctx, msg.FileHash)

	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("index %v already set", msg.FileHash))
	}

	var contractFile = types.ContractFile{
		FileHash:           msg.FileHash,
		Creator:            msg.Creator,
		TimeStamp:          msg.TimeStamp,
		OwnerList:          msg.OwnerList,
		MetaDataJsonString: msg.MetaDataJsonString,
	}

	k.SetContractFile(
		ctx,
		contractFile,
	)

	return &types.MsgCreateContractFileResponse{}, nil
}
