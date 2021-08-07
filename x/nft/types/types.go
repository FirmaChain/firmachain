package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type NFT struct {
	Hash    string         `json:"hash"`
	Owner   sdk.AccAddress `json:"owner"`
	Creator sdk.AccAddress `json:"creator"` // nft creator, the origianl author

	TokenURI    string `json:"tokenURI"`    // meta data url (it's only data link url. not chain data)
	Description string `json:"description"` // basic unique description of the NFT (chain data)
	Image       string `json:"image"`       // basic image path (chain data)
}

// below NFT code is data structure for next version (like ERC 1155)
// the detail implementation will be added soon.

// NFT non fungible token interface
type NFTInterface interface {
	GetID() string
	GetOwner() sdk.AccAddress
	SetOwner(address sdk.AccAddress) NFT
	GetName() string
	GetDescription() string
	GetImage() string
	GetTokenURI() string
	EditMetadata(name, description, image, tokenURI string) NFT
	String() string
}

// NFT array for collection item like ERC 1155
type NFTs []NFT

// Collection of non fungible tokens
type Collection struct {
	Denom string `json:"denom,omitempty"` // name of the collection; not exported to clients
	NFTs  NFTs   `json:"nfts"`            // NFTs that belong to a collection
}

// Collections define an array of Collection
type Collections []Collection

func NewNFT() NFT {
	return NFT{}
}

func (nft NFT) String() string {

	return strings.TrimSpace(fmt.Sprintf(`Hash %s
TokenURI %s
Owner %s
Creator %s`, nft.Hash, nft.TokenURI, nft.Owner, nft.Creator))
}
