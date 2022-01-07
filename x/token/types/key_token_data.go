package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// TokenDataKeyPrefix is the prefix to retrieve all TokenData
	TokenDataKeyPrefix = "TokenDataKeyPrefix"
)

// TokenDataKey returns the store key to retrieve a TokenData from the index fields
func TokenDataKey(
	tokenID string,
) []byte {
	var key []byte

	tokenIDBytes := []byte(tokenID)
	key = append(key, tokenIDBytes...)
	key = append(key, []byte("/")...)

	return key
}
