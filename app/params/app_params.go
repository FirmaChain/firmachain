package params

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	HumanCoinUnit     = "fct"
	BaseCoinUnit      = "ufct"
	HumanCoinExponent = 6

	DefaultBondDenom = BaseCoinUnit
)

const (
	AccountAddressPrefix = "firma"
	Name                 = "firmachain"
	NodeDir              = ".firmachain"
)

var (
// If EnabledSpecificProposals is "", and this is "true", then enable all x/wasm proposals.
// If EnabledSpecificProposals is "", and this is not "true", then disable all x/wasm proposals.
// ProposalsEnabled = "true"
// If set to non-empty string it must be comma-separated list of values that are all a subset
// of "EnableAllProposals" (takes precedence over ProposalsEnabled)
// https://github.com/CosmWasm/wasmd/blob/02a54d33ff2c064f3539ae12d75d027d9c665f05/x/wasm/internal/types/proposal.go#L28-L34
// EnableSpecificProposals = ""
)

// These constants are derived from the above variables.
// These are the ones we will want to use in the code, based on
// any overrides above
var (
	Bech32Prefix = AccountAddressPrefix

	// Bech32PrefixAccAddr defines the Bech32 prefix of an account's address
	Bech32PrefixAccAddr = Bech32Prefix
	// Bech32PrefixAccPub defines the Bech32 prefix of an account's public key
	Bech32PrefixAccPub = Bech32Prefix + sdk.PrefixPublic
	// Bech32PrefixValAddr defines the Bech32 prefix of a validator's operator address
	Bech32PrefixValAddr = Bech32Prefix + sdk.PrefixValidator + sdk.PrefixOperator
	// Bech32PrefixValPub defines the Bech32 prefix of a validator's operator public key
	Bech32PrefixValPub = Bech32Prefix + sdk.PrefixValidator + sdk.PrefixOperator + sdk.PrefixPublic
	// Bech32PrefixConsAddr defines the Bech32 prefix of a consensus node address
	Bech32PrefixConsAddr = Bech32Prefix + sdk.PrefixValidator + sdk.PrefixConsensus
	// Bech32PrefixConsPub defines the Bech32 prefix of a consensus node public key
	Bech32PrefixConsPub = Bech32Prefix + sdk.PrefixValidator + sdk.PrefixConsensus + sdk.PrefixPublic
)

// Bech32Prefix: sdk.GetConfig().GetBech32AccountAddrPrefix()
