package types

// nft module event types
const (
	EventTypeMint     = "nft_mint"
	EventTypeTransfer = "nft_transfer"
	EventTypeBurn     = "nft_burn"

	AttributeKeyOwner     = "owner"
	AttributeKeySender    = "sender"
	AttributeKeyRecipient = "recipient"
	AttributeKeyHash      = "hash"
)
