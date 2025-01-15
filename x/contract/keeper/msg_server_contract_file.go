package keeper

import (
	"context"
	"fmt"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/firmachain/firmachain/v05/x/contract/types"
)

func (ms msgServer) CreateContractFile(goCtx context.Context, msg *types.MsgCreateContractFile) (*types.MsgCreateContractFileResponse, error) {

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := ms.keeper.GetContractFile(ctx, msg.FileHash)

	if isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("index %v already set", msg.FileHash))
	}

	var contractFile = types.ContractFile{
		FileHash:           msg.FileHash,
		Creator:            msg.Creator,
		TimeStamp:          msg.TimeStamp,
		OwnerList:          msg.OwnerList,
		MetaDataJsonString: msg.MetaDataJsonString,
	}

	ms.keeper.SetContractFile(
		ctx,
		contractFile,
	)

	return &types.MsgCreateContractFileResponse{}, nil
}
