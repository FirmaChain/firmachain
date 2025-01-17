package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// TokenDataKeyPrefix is the prefix to retrieve all TokenData
	TokenDataKeyPrefix = "TokenDataKeyPrefix"
)

// TokenDataKey returns the store key to retrieve a TokenData from the index fields
func TokenDataKey(
	tokenId string,
) []byte {
	var key []byte

	tokenIdBytes := []byte(tokenId)
	key = append(key, tokenIdBytes...)
	key = append(key, []byte("/")...)

	return key
}
