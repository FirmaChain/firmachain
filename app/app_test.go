package app_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/cosmos/cosmos-sdk/x/auth/signing"

	"github.com/firmachain/firmachain/v05/app/apptesting"

	abci "github.com/cometbft/cometbft/abci/types"
)

type AppTestSuite struct {
	apptesting.TestSuite
}

func (s *AppTestSuite) SetupTest() {
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(AppTestSuite))
}

// Check if a tx can be run without encountering internal application errors.
func (s *AppTestSuite) NoInternalErrorOnRunTx(tx signing.Tx) {
	txConfig := s.App.GetTxConfig()
	txBytes, err := txConfig.TxEncoder()(tx)
	require.NoError(s.T(), err)
	req := &abci.RequestCheckTx{Tx: txBytes}
	_, err = s.App.BaseApp.CheckTx(req)
	require.NoError(s.T(), err)
}

// Check if dummy tx from all modules can be correctly handled by the application,
// without focusing on the actual meaning of these transactions.
func (s *AppTestSuite) TestModulesTxHandlers() {
	s.Setup()

	s.NoInternalErrorOnRunTx(s.NewDummyTxContractMsgCreateContractFile())
	s.NoInternalErrorOnRunTx(s.NewDummyTxContractMsgAddContractLog())

	s.NoInternalErrorOnRunTx(s.NewDummyTxNftMsgBurn())
	s.NoInternalErrorOnRunTx(s.NewDummyTxNftMsgMint())
	s.NoInternalErrorOnRunTx(s.NewDummyTxNftMsgTransfer())

	s.NoInternalErrorOnRunTx(s.NewDummyTxTokenMsgBurn())
	s.NoInternalErrorOnRunTx(s.NewDummyTxTokenMsgCreateToken())
	s.NoInternalErrorOnRunTx(s.NewDummyTxTokenMsgMint())
	s.NoInternalErrorOnRunTx(s.NewDummyTxTokenMsgUpdateTokenURI())
}
