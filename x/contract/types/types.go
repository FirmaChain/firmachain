package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Contract struct {
	Hash   string           `json:"hash"`
	Path   string           `json:"path"`
	Owners []sdk.AccAddress `json:"owners"`
}

func NewContract() Contract {
	return Contract{}
}

func (c Contract) String() string {
	var owners []string
	for _, address := range c.Owners {
		owners = append(owners, address.String())
	}

	return strings.TrimSpace(fmt.Sprintf(`Hash %s
Paths %s
Owner %s`, strings.Join(owners[:], ","), c.Path, c.Hash))
}
