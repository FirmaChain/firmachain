package types

import "strings"

const (
	QueryContract = "contract"
)

type QueryResContract []string

func (n QueryResContract) String() string {
	return strings.Join(n[:], "\n")
}
