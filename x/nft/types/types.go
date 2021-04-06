package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type NFToken struct {
	Hash     string         `json:"hash"`
	TokenURI string         `json:"tokenURI"`
	Owner    sdk.AccAddress `json:"owner"`
}

func NewNFToken() NFToken {
	return NFToken{}
}

func (nft NFToken) String() string {

	return strings.TrimSpace(fmt.Sprintf(`Hash %s
TokenURI %s
Owner %s`, nft.Hash, nft.TokenURI, nft.Owner))
}
