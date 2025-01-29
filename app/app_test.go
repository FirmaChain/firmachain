package app_test

import (
	"testing"

	math "cosmossdk.io/math"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	signingtypes "github.com/cosmos/cosmos-sdk/types/tx/signing"
	"github.com/cosmos/cosmos-sdk/x/auth/signing"
	"github.com/firmachain/firmachain/v05/app/apptesting"

	appparams "github.com/firmachain/firmachain/v05/app/params"
	contracttypes "github.com/firmachain/firmachain/v05/x/contract/types"

	abci "github.com/cometbft/cometbft/abci/types"
)

type AppTestSuite struct {
	apptesting.KeeperTestHelper
}

func (s *AppTestSuite) SetupTest() {
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(AppTestSuite))
}

// KeyTestPubAddr generates a new secp256k1 keypair.
func KeyTestPubAddr() (cryptotypes.PrivKey, cryptotypes.PubKey, sdk.AccAddress) {
	key := secp256k1.GenPrivKey()
	pub := key.PubKey()
	addr := sdk.AccAddress(pub.Address())
	return key, pub, addr
}

func SetTxSignature(t *testing.T, builder client.TxBuilder, privKey *secp256k1.PrivKey, nonce uint64) {
	pubKey := privKey.PubKey()
	err := builder.SetSignatures(
		signingtypes.SignatureV2{
			PubKey:   pubKey,
			Sequence: nonce,
			Data:     &signingtypes.SingleSignatureData{},
		},
	)
	require.NoError(t, err)
}

func NewTxMsgCreateContractFile(t *testing.T, cfg client.TxConfig, signerPrivKey *secp256k1.PrivKey, sequence uint64) signing.Tx {
	_, _, addr := KeyTestPubAddr()
	msgs := make([]sdk.Msg, 0, 1)
	msg := contracttypes.NewMsgCreateContractFile(
		addr.String(),
		"dummyfilehash",
		1234,
		[]string{addr.String()},
		"{}",
	)
	msgs = append(msgs, msg)

	builder := cfg.NewTxBuilder()
	builder.SetMsgs(msgs...)
	builder.SetMemo("test MsgCreateContractFile")
	builder.SetGasLimit(9999999999999)
	builder.SetFeeAmount(sdk.NewCoins(sdk.NewCoin(appparams.DefaultBondDenom, math.NewInt(1000))))
	SetTxSignature(t, builder, signerPrivKey, sequence)

	return builder.GetTx()
}

// Ensures the test does not error out.
func (s *AppTestSuite) TestModulesTxHandlers() {
	s.Setup()
	txConfig := s.App.GetTxConfig()

	tx := NewTxMsgCreateContractFile(s.T(), txConfig, s.TestAccs[0].PrivKey, 1)
	txBytes, err := txConfig.TxEncoder()(tx)
	require.NoError(s.T(), err)

	req := &abci.RequestCheckTx{Tx: txBytes}
	var res *abci.ResponseCheckTx
	res, err = s.App.BaseApp.CheckTx(req)
	require.NoError(s.T(), err)
	res.IsOK()

}
