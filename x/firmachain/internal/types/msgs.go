package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/firmachain/FirmaChain/x/firmachain/internal/utils"
)

const RouterKey = ModuleName

type MsgAddContract struct {
	Path  string         `json:"path"`
	Hash  string         `json:"hash"`
	Owner sdk.AccAddress `json:"owner"`
}

func NewMsgAddContract(path string, hash string, owner sdk.AccAddress) MsgAddContract {
	return MsgAddContract{
		Path:  path,
		Hash:  hash,
		Owner: owner,
	}
}

func (msg MsgAddContract) Route() string { return RouterKey }

func (msg MsgAddContract) Type() string { return "add_contract" }

func (msg MsgAddContract) ValidateBasic() sdk.Error {
	if msg.Owner.Empty() {
		return sdk.ErrInvalidAddress(msg.Owner.String())
	}
	if len(msg.Path) == 0 || len(msg.Hash) == 0 {
		return sdk.ErrUnknownRequest("Path or Hash cannot be empty")
	}
	if err := utils.VerifyFile(msg.Path, msg.Hash); err != nil {
		return sdk.ErrUnknownRequest("Contract has been manipulated or invalid.")
	}
	return nil
}

func (msg MsgAddContract) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg MsgAddContract) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}
