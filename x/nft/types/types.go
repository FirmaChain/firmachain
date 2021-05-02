package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type NFT struct {
	Hash     string         `json:"hash"`
	TokenURI string         `json:"tokenURI"`
	Owner    sdk.AccAddress `json:"owner"`
}

func NewNFT() NFT {
	return NFT{}
}

func (nft NFT) String() string {

	return strings.TrimSpace(fmt.Sprintf(`Hash %s
TokenURI %s
Owner %s`, nft.Hash, nft.TokenURI, nft.Owner))
}
