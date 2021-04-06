package types

import "strings"

const (
	QueryNFToken = "nft"
)

type QueryResNFToken []string

func (n QueryResNFToken) String() string {
	return strings.Join(n[:], "\n")
}
