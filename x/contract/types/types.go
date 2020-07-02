package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Contract struct {
	Path  string         `json:"path"`
	Hash  string         `json:"hash"`
	Owner sdk.AccAddress `json:"owner"`
}

func NewContract() Contract {
	return Contract{}
}

func (c Contract) String() string {
	return strings.TrimSpace(fmt.Sprintf(`Owner: %s
Path: %s
Hash: %s`, c.Owner, c.Path, c.Hash))
}
