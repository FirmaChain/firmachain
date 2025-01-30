package apptesting

import (
	//"fmt"
	"os"
	"time"

	"github.com/stretchr/testify/suite"

	"cosmossdk.io/log"
	"cosmossdk.io/math"

	coreheader "cosmossdk.io/core/header"
	abci "github.com/cometbft/cometbft/abci/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	signingtypes "github.com/cosmos/cosmos-sdk/types/tx/signing"
	"github.com/cosmos/cosmos-sdk/x/auth/signing"

	upgradetypes "cosmossdk.io/x/upgrade/types"
	"github.com/cosmos/cosmos-sdk/baseapp"

	sdk "github.com/cosmos/cosmos-sdk/types"

	stakinghelper "github.com/cosmos/cosmos-sdk/x/staking/testutil"

	"github.com/firmachain/firmachain/v05/app"
	appparams "github.com/firmachain/firmachain/v05/app/params"
	contracttypes "github.com/firmachain/firmachain/v05/x/contract/types"
	nfttypes "github.com/firmachain/firmachain/v05/x/nft/types"
	tokentypes "github.com/firmachain/firmachain/v05/x/token/types"

	header "cosmossdk.io/core/header"
)

type AddressWithKeys struct {
	PrivKey *secp256k1.PrivKey
	PubKey  cryptotypes.PubKey
	Address sdk.AccAddress
}

type TxRequirements struct {
	cfg           client.TxConfig
	signerPrivKey *secp256k1.PrivKey
	sequence      uint64
	memo          string
	gasLimit      uint64
	feeAmount     sdk.Coins
}

type TestSuite struct {
	suite.Suite

	App         *app.App
	Ctx         sdk.Context
	QueryHelper *baseapp.QueryServiceTestHelper
	TestAccs    []AddressWithKeys

	Logger log.Logger

	StakingHelper *stakinghelper.Helper
}

// ========================================================================
// 								  SETUP
// ========================================================================

// Setup sets up basic environment for suite (App, Ctx, and test accounts)
func (s *TestSuite) Setup() {
	t := s.T()
	s.Logger = log.NewLogger(os.Stderr)
	s.App, s.Ctx, s.TestAccs = SetupApp(t)
	s.Commit()

	s.SetupHelpers()

	// FIXME: begin block make the app panic
	// New block to test if everything goes smoothly.
	//s.MakeBlock()
}

// Setup helpers sets the suit helpers
func (s *TestSuite) SetupHelpers() {
	s.QueryHelper = &baseapp.QueryServiceTestHelper{
		GRPCQueryRouter: s.App.GRPCQueryRouter(),
		Ctx:             s.Ctx,
	}
	s.StakingHelper = stakinghelper.NewHelper(s.Suite.T(), s.Ctx, s.App.AppKeepers.StakingKeeper)
	s.StakingHelper.Denom = "ufct"
}

// FinalizeBlock prepares the block to be committed: it runs begin-blockers, then delivers txs, then end-blockers.
func (s *TestSuite) FinalizeBlock() {
	_, err := s.App.FinalizeBlock(&abci.RequestFinalizeBlock{Height: s.Ctx.BlockHeight(), Time: s.Ctx.BlockTime()})
	s.NoError(err)
}

// Commit commits the store changes
func (s *TestSuite) Commit() {
	_, err := s.App.Commit()
	s.NoError(err)

	// Update the suite context
	newBlockTime := s.Ctx.BlockTime().Add(time.Second)
	header := s.Ctx.BlockHeader()
	header.Time = newBlockTime
	header.Height++
	s.Ctx = s.App.BaseApp.NewUncachedContext(false, header).WithHeaderInfo(coreheader.Info{
		Height: header.Height,
		Time:   header.Time,
	})
}

// MakeBlock creates the block, commits it and updates the suite context.
func (s *TestSuite) MakeBlock() {
	s.FinalizeBlock()
	s.Commit()
}

// ========================================================================
// 								  UPGRADE
// ========================================================================

func (s *TestSuite) ConfirmUpgradeSucceeded(upgradeName string, upgradeHeight int64) {
	s.Ctx = s.Ctx.WithBlockHeight(upgradeHeight - 1)
	plan := upgradetypes.Plan{Name: upgradeName, Height: upgradeHeight}
	err := s.App.AppKeepers.UpgradeKeeper.ScheduleUpgrade(s.Ctx, plan)
	s.Require().NoError(err)
	planGot, err := s.App.AppKeepers.UpgradeKeeper.GetUpgradePlan(s.Ctx)
	_ = planGot
	s.Require().NoError(err)

	s.Ctx = s.Ctx.WithBlockHeight(upgradeHeight)

	s.Ctx = s.Ctx.WithHeaderInfo(header.Info{Height: upgradeHeight, Time: s.Ctx.BlockTime().Add(time.Second)}).WithBlockHeight(upgradeHeight)
	/*
		s.Require().NotPanics(func() {
			res, err := s.App.PreBlocker(s.Ctx, nil)
			_ = res
			s.Require().NoError(err)
		})
	*/
	res, err := s.App.PreBlocker(s.Ctx, nil)
	_ = res
	s.Require().NoError(err)
}

// ========================================================================
// 								TRANSACTIONS
// ========================================================================

func (s *TestSuite) SetTxSignature(builder client.TxBuilder, privKey *secp256k1.PrivKey, nonce uint64) {
	pubKey := privKey.PubKey()
	err := builder.SetSignatures(
		signingtypes.SignatureV2{
			PubKey:   pubKey,
			Sequence: nonce,
			Data:     &signingtypes.SingleSignatureData{},
		},
	)
	s.NoError(err)
}

func (s *TestSuite) MakeBondDenomFeeAmount(amount int64) sdk.Coins {
	return sdk.NewCoins(sdk.NewCoin(appparams.DefaultBondDenom, math.NewInt(1000)))
}

func (s *TestSuite) MakeDummyTxRequirements() TxRequirements {
	return TxRequirements{
		cfg:           s.App.GetTxConfig(),
		signerPrivKey: s.TestAccs[0].PrivKey,
		sequence:      0,
		memo:          "dummy memo",
		gasLimit:      999999999,
		feeAmount:     s.MakeBondDenomFeeAmount(9999),
	}
}

func (s *TestSuite) MakeAndSingTx(
	cfg client.TxConfig,
	msg sdk.Msg,
	memo string,
	gasLimit uint64,
	feeAmount sdk.Coins,
	signerPrivKey *secp256k1.PrivKey,
	sequence uint64,
) signing.Tx {
	msgs := make([]sdk.Msg, 0, 1)
	msgs = append(msgs, msg)
	builder := cfg.NewTxBuilder()
	builder.SetMsgs(msgs...)
	builder.SetMemo(memo)
	builder.SetGasLimit(gasLimit)
	builder.SetFeeAmount(feeAmount)
	s.SetTxSignature(builder, signerPrivKey, sequence)
	return builder.GetTx()
}

// ==== Contract messages ====

func (s *TestSuite) NewTxContractMsgCreateContractFile(
	creatorAddr sdk.AccAddress,
	fileHash string,
	timeStamp uint64,
	metadataJsonString string,
	ownerList []sdk.AccAddress,
	txR TxRequirements,
) signing.Tx {
	ownerListString := make([]string, 0, len(ownerList))
	for _, owner := range ownerList {
		ownerListString = append(ownerListString, owner.String())
	}
	msg := contracttypes.NewMsgCreateContractFile(
		creatorAddr.String(),
		fileHash,
		timeStamp,
		ownerListString,
		metadataJsonString,
	)
	return s.MakeAndSingTx(txR.cfg, msg, txR.memo, txR.gasLimit, txR.feeAmount, txR.signerPrivKey, txR.sequence)
}

func (s *TestSuite) NewDummyTxContractMsgCreateContractFile() signing.Tx {
	return s.NewTxContractMsgCreateContractFile(
		s.TestAccs[0].Address,
		"dummyfilehash",
		1234,
		"{}",
		[]sdk.AccAddress{s.TestAccs[0].Address},
		s.MakeDummyTxRequirements(),
	)
}

func (s *TestSuite) NewTxContractMsgAddContractLog(
	creatorAddr sdk.AccAddress,
	contractHash string,
	timeStamp uint64,
	eventName string,
	ownerAddr sdk.AccAddress,
	jsonString string,
	txR TxRequirements,
) signing.Tx {
	msg := contracttypes.NewMsgAddContractLog(creatorAddr.String(), contractHash, timeStamp, eventName, ownerAddr.String(), jsonString)
	return s.MakeAndSingTx(txR.cfg, msg, txR.memo, txR.gasLimit, txR.feeAmount, txR.signerPrivKey, txR.sequence)
}

func (s *TestSuite) NewDummyTxContractMsgAddContractLog() signing.Tx {
	return s.NewTxContractMsgAddContractLog(
		s.TestAccs[0].Address,
		"dummyfilehash",
		1234,
		"dummyevent",
		s.TestAccs[0].Address,
		"{}",
		s.MakeDummyTxRequirements(),
	)
}

// ==== Nft messages ====

func (s *TestSuite) NewTxNftMsgBurn(
	ownerAddr sdk.AccAddress,
	nftId uint64,
	txR TxRequirements,
) signing.Tx {
	msg := nfttypes.NewMsgBurn(ownerAddr.String(), nftId)
	return s.MakeAndSingTx(txR.cfg, msg, txR.memo, txR.gasLimit, txR.feeAmount, txR.signerPrivKey, txR.sequence)
}

func (s *TestSuite) NewDummyTxNftMsgBurn() signing.Tx {
	return s.NewTxNftMsgBurn(
		s.TestAccs[0].Address,
		1,
		s.MakeDummyTxRequirements(),
	)
}

func (s *TestSuite) NewTxNftMsgMint(
	ownerAddr sdk.AccAddress,
	tokenURI string,
	txR TxRequirements,
) signing.Tx {
	msg := nfttypes.NewMsgMint(ownerAddr.String(), tokenURI)
	return s.MakeAndSingTx(txR.cfg, msg, txR.memo, txR.gasLimit, txR.feeAmount, txR.signerPrivKey, txR.sequence)
}

func (s *TestSuite) NewDummyTxNftMsgMint() signing.Tx {
	return s.NewTxNftMsgMint(
		s.TestAccs[0].Address,
		"dummy tokenURI",
		s.MakeDummyTxRequirements(),
	)
}

func (s *TestSuite) NewTxNftMsgTransfer(
	ownerAddr sdk.AccAddress,
	nftId uint64,
	toAddr sdk.AccAddress,
	txR TxRequirements,
) signing.Tx {
	msg := nfttypes.NewMsgTransfer(ownerAddr.String(), nftId, toAddr.String())
	return s.MakeAndSingTx(txR.cfg, msg, txR.memo, txR.gasLimit, txR.feeAmount, txR.signerPrivKey, txR.sequence)
}

func (s *TestSuite) NewDummyTxNftMsgTransfer() signing.Tx {
	return s.NewTxNftMsgTransfer(
		s.TestAccs[0].Address,
		1,
		s.TestAccs[1].Address,
		s.MakeDummyTxRequirements(),
	)
}

// ==== Token messages ====

func (s *TestSuite) NewTxTokenMsgBurn(
	ownerAddr sdk.AccAddress,
	tokenID string,
	amount uint64,
	txR TxRequirements,
) signing.Tx {
	msg := tokentypes.NewMsgBurn(ownerAddr.String(), tokenID, amount)
	return s.MakeAndSingTx(txR.cfg, msg, txR.memo, txR.gasLimit, txR.feeAmount, txR.signerPrivKey, txR.sequence)
}

func (s *TestSuite) NewDummyTxTokenMsgBurn() signing.Tx {
	return s.NewTxTokenMsgBurn(
		s.TestAccs[0].Address,
		"dummytokenID",
		1234,
		s.MakeDummyTxRequirements(),
	)
}

func (s *TestSuite) NewTxTokenMsgCreateToken(
	ownerAddr sdk.AccAddress,
	name string,
	symbol string,
	tokenUri string,
	totalSupply uint64,
	decimal uint64,
	mintable bool,
	burnable bool,
	txR TxRequirements,
) signing.Tx {
	msg := tokentypes.NewMsgCreateToken(ownerAddr.String(), name, symbol, tokenUri, totalSupply, decimal, mintable, burnable)
	return s.MakeAndSingTx(txR.cfg, msg, txR.memo, txR.gasLimit, txR.feeAmount, txR.signerPrivKey, txR.sequence)
}

func (s *TestSuite) NewDummyTxTokenMsgCreateToken() signing.Tx {
	return s.NewTxTokenMsgCreateToken(
		s.TestAccs[0].Address,
		"dummytokenID",
		"dt1",
		"dummytokenURI",
		123456789,
		6,
		true,
		true,
		s.MakeDummyTxRequirements(),
	)
}

func (s *TestSuite) NewTxTokenMsgMint(
	ownerAddr sdk.AccAddress,
	tokenID string,
	amount uint64,
	toAddr sdk.AccAddress,
	txR TxRequirements,
) signing.Tx {
	msg := tokentypes.NewMsgMint(ownerAddr.String(), tokenID, amount, toAddr.String())
	return s.MakeAndSingTx(txR.cfg, msg, txR.memo, txR.gasLimit, txR.feeAmount, txR.signerPrivKey, txR.sequence)
}

func (s *TestSuite) NewDummyTxTokenMsgMint() signing.Tx {
	return s.NewTxTokenMsgMint(
		s.TestAccs[0].Address,
		"dummytokenID",
		1234,
		s.TestAccs[1].Address,
		s.MakeDummyTxRequirements(),
	)
}

func (s *TestSuite) NewTxTokenMsgUpdateTokenURI(
	ownerAddr sdk.AccAddress,
	tokenID string,
	tokenURI string,
	txR TxRequirements,
) signing.Tx {
	msg := tokentypes.NewMsgUpdateTokenURI(ownerAddr.String(), tokenID, tokenURI)
	return s.MakeAndSingTx(txR.cfg, msg, txR.memo, txR.gasLimit, txR.feeAmount, txR.signerPrivKey, txR.sequence)
}

func (s *TestSuite) NewDummyTxTokenMsgUpdateTokenURI() signing.Tx {
	return s.NewTxTokenMsgUpdateTokenURI(
		s.TestAccs[0].Address,
		"dummytokenID",
		"newDummyTokenURI",
		s.MakeDummyTxRequirements(),
	)
}
