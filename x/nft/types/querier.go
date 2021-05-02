package types

import "strings"

const (
	QueryNFT = "nft"
)

type QueryResNFT []string

func (n QueryResNFT) String() string {
	return strings.Join(n[:], "\n")
}
