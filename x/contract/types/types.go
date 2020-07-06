package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Contract struct {
	Hash   string           `json:"hash"`
	Paths  []string         `json:"paths"`
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
Owner %s`, strings.Join(owners[:], ","), strings.Join(c.Paths[:], ","), c.Hash))
}
