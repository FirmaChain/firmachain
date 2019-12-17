package types

import "strings"

type QueryResContract []string

func (n QueryResContract) String() string {
	return strings.Join(n[:], "\n")
}
