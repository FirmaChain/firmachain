package auth

import (
	"github.com/cosmos/cosmos-sdk/x/auth"
)

// SignatureVerificationGasConsumer is the type of function that is used to both consume gas when verifying signatures
// and also to accept or reject different types of PubKey's. This is where apps can define their own PubKey
type SignatureVerificationGasConsumer = auth.SignatureVerificationGasConsumer

// NewAnteHandler returns an AnteHandler that checks and increments sequence
// numbers, checks signatures & account numbers, and deducts fees from the first
// signer.
var NewAnteHandler = auth.NewAnteHandler

// DefaultSigVerificationGasConsumer is the default implementation of SignatureVerificationGasConsumer. It consumes gas
// for signature verification based upon the public key type. The cost is fetched from the given params and is matched
// by the concrete type.
var DefaultSigVerificationGasConsumer = auth.DefaultSigVerificationGasConsumer
