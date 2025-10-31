package apptesting

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"testing"
	"time"

	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	"github.com/stretchr/testify/require"

	coreheader "cosmossdk.io/core/header"
	"cosmossdk.io/log"
	abci "github.com/cometbft/cometbft/abci/types"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	tmtypes "github.com/cometbft/cometbft/types"
	dbm "github.com/cosmos/cosmos-db"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"

	"cosmossdk.io/math"

	bam "github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client/flags"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	simtestutil "github.com/cosmos/cosmos-sdk/testutil/sims"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	"github.com/firmachain/firmachain/app"
	apphelpers "github.com/firmachain/firmachain/app/helpers"
	appparams "github.com/firmachain/firmachain/app/params"
)

func SetupApp(t *testing.T, chainId string, bondDenom string) (*app.App, sdk.Context, []AddressWithKeys) {
	t.Helper()

	appparams.SetSdkConfigAndSeal()

	privVal := apphelpers.NewPV()
	pubKey, err := privVal.GetPubKey()
	require.NoError(t, err)

	// create validator set with single validator
	validator := tmtypes.NewValidator(pubKey, 1)
	valSet := tmtypes.NewValidatorSet([]*tmtypes.Validator{validator})

	// Genesis ------------------------------------------------------------------
	testAddrsWithKeys := make([]AddressWithKeys, 3)
	genesisAccounts := make([]authtypes.GenesisAccount, 0, 4)
	genesisBalances := make([]banktypes.Balance, 0, 4)
	// generate the validator genesis accounts
	senderPrivKey := secp256k1.GenPrivKey()
	acc := authtypes.NewBaseAccount(senderPrivKey.PubKey().Address().Bytes(), senderPrivKey.PubKey(), 0, 0)
	balance := banktypes.Balance{
		Address: acc.GetAddress().String(),
		Coins:   sdk.NewCoins(sdk.NewCoin(bondDenom, math.NewInt(100000000000000))),
	}
	genesisAccounts = append(genesisAccounts, acc)
	genesisBalances = append(genesisBalances, balance)

	// generate three other genesis accounts
	for i := 0; i < 3; i++ {
		priv := secp256k1.GenPrivKey()
		pub := priv.PubKey()
		testAddrsWithKeys[i].PrivKey = priv
		testAddrsWithKeys[i].PubKey = pub
		testAddrsWithKeys[i].Address = sdk.AccAddress(pub.Address())
		acc := authtypes.NewBaseAccount(pub.Address().Bytes(), pub, 0, 0)
		balance = banktypes.Balance{
			Address: acc.GetAddress().String(),
			Coins:   sdk.NewCoins(sdk.NewCoin(bondDenom, math.NewInt(100000000000000))),
		}
		genesisAccounts = append(genesisAccounts, acc)
		genesisBalances = append(genesisBalances, balance)
	}

	timenow := time.Now()
	initialHeight := int64(1)

	app := SetupWithGenesisValSet(t, chainId, bondDenom, timenow, initialHeight, valSet, genesisAccounts, genesisBalances...)

	ctx := app.BaseApp.NewContextLegacy(false, tmproto.Header{
		ChainID: chainId,
		Height:  initialHeight,
		Time:    timenow,
	}).WithHeaderInfo(coreheader.Info{
		ChainID: chainId,
		Height:  initialHeight,
		Time:    timenow,
	})

	return app, ctx, testAddrsWithKeys
}

// SetupWithGenesisValSet initializes a new firmachainApp with a validator set and genesis accounts
// that also act as delegators. For simplicity, each validator is bonded with a delegation
// of one consensus engine unit in the default token of the firmachainApp from first genesis
// account. A Nop logger is set in firmachainApp.
func SetupWithGenesisValSet(
	t *testing.T,
	chainId string,
	bondDenom string,
	initialTime time.Time,
	initialHeight int64,
	valSet *tmtypes.ValidatorSet,
	genAccs []authtypes.GenesisAccount,
	balances ...banktypes.Balance,
) *app.App {
	t.Helper()

	const withGenesis = true
	firmachainApp, genesisState := setup(t, withGenesis, chainId)
	genesisState = genesisStateWithValSet(t, firmachainApp, bondDenom, genesisState, valSet, genAccs, balances...)

	stateBytes, err := json.MarshalIndent(genesisState, "", " ")
	require.NoError(t, err)

	// init chain will set the validator set and initialize the genesis accounts
	_, err = firmachainApp.InitChain(
		&abci.RequestInitChain{
			Validators:      []abci.ValidatorUpdate{},
			ConsensusParams: DefaultConsensusParams,
			AppStateBytes:   stateBytes,
			ChainId:         chainId,
			Time:            initialTime,
			InitialHeight:   initialHeight,
		},
	)
	require.NoError(t, err, "Failed to setup app: InitChain failed.")

	return firmachainApp
}

// DefaultConsensusParams defines the default Tendermint consensus params used
// in firmachainApp testing.
var DefaultConsensusParams = &tmproto.ConsensusParams{
	Block: &tmproto.BlockParams{
		MaxBytes: 200000,
		MaxGas:   2000000,
	},
	Evidence: &tmproto.EvidenceParams{
		MaxAgeNumBlocks: 302400,
		MaxAgeDuration:  504 * time.Hour, // 3 weeks is the max duration
		MaxBytes:        10000,
	},
	Validator: &tmproto.ValidatorParams{
		PubKeyTypes: []string{
			tmtypes.ABCIPubKeyTypeEd25519,
		},
	},
}

type EmptyAppOptions struct{}

func (EmptyAppOptions) Get(_ string) interface{} { return nil }

// EmptyBaseAppOptions is a stub implementing AppOptions
type EmptyBaseAppOptions struct{}

// Get implements AppOptions
func (ao EmptyBaseAppOptions) Get(_ string) interface{} {
	return nil
}

func setup(t *testing.T, withGenesis bool, chainId string, opts ...wasmkeeper.Option) (*app.App, app.GenesisState) {
	db := dbm.NewMemDB()
	nodeHome := t.TempDir()
	snapshotDir := filepath.Join(nodeHome, "data", "snapshots")

	snapshotDB, err := dbm.NewDB("metadata", dbm.GoLevelDBBackend, snapshotDir)
	require.NoError(t, err)
	t.Cleanup(func() { snapshotDB.Close() })
	require.NoError(t, err)

	// var emptyWasmOpts []wasm.Option
	appOptions := make(simtestutil.AppOptionsMap, 0)
	appOptions[flags.FlagHome] = nodeHome // ensure unique folder

	app := app.New(
		log.NewNopLogger(),
		db,
		nil,
		true,
		EmptyAppOptions{},
		opts,
		bam.SetChainID(chainId),
	)
	if withGenesis {
		genesisState := app.NewDefaultGenesisState(app.AppCodec())
		_ = genesisState
		return app, genesisState
	}

	return app, nil
}

func genesisStateWithValSet(
	t *testing.T,
	app *app.App,
	bondDenom string,
	genesisState app.GenesisState,
	valSet *tmtypes.ValidatorSet,
	genAccs []authtypes.GenesisAccount,
	balances ...banktypes.Balance,
) app.GenesisState {
	codec := app.AppCodec()

	// set genesis accounts
	authGenesis := authtypes.NewGenesisState(authtypes.DefaultParams(), genAccs)
	genesisState[authtypes.ModuleName] = codec.MustMarshalJSON(authGenesis)

	validators := make([]stakingtypes.Validator, 0, len(valSet.Validators))
	delegations := make([]stakingtypes.Delegation, 0, len(valSet.Validators))

	bondAmt := sdk.DefaultPowerReduction

	for _, val := range valSet.Validators {
		pk, err := cryptocodec.FromCmtPubKeyInterface(val.PubKey)
		require.NoError(t, err)
		pkAny, err := codectypes.NewAnyWithValue(pk)
		require.NoError(t, err)
		validator := stakingtypes.Validator{
			OperatorAddress:   sdk.ValAddress(val.Address).String(),
			ConsensusPubkey:   pkAny,
			Jailed:            false,
			Status:            stakingtypes.Bonded,
			Tokens:            bondAmt,
			DelegatorShares:   math.LegacyOneDec(),
			Description:       stakingtypes.Description{},
			UnbondingHeight:   int64(0),
			UnbondingTime:     time.Unix(0, 0).UTC(),
			Commission:        stakingtypes.NewCommission(math.LegacyZeroDec(), math.LegacyZeroDec(), math.LegacyZeroDec()),
			MinSelfDelegation: math.ZeroInt(),
		}
		validators = append(validators, validator)
		delegations = append(delegations, stakingtypes.NewDelegation(genAccs[0].GetAddress().String(), sdk.ValAddress(val.Address).String(), math.LegacyOneDec()))

	}

	defaultStParams := stakingtypes.DefaultParams()
	stParams := stakingtypes.NewParams(
		defaultStParams.UnbondingTime,
		defaultStParams.MaxValidators,
		defaultStParams.MaxEntries,
		defaultStParams.HistoricalEntries,
		bondDenom,
		defaultStParams.MinCommissionRate, // 5%
	)

	// set validators and delegations
	stakingGenesis := stakingtypes.NewGenesisState(stParams, validators, delegations)
	genesisState[stakingtypes.ModuleName] = codec.MustMarshalJSON(stakingGenesis)

	// add bonded amount to bonded pool module account
	balances = append(balances, banktypes.Balance{
		Address: authtypes.NewModuleAddress(stakingtypes.BondedPoolName).String(),
		Coins:   sdk.Coins{sdk.NewCoin(bondDenom, bondAmt.MulRaw(int64(len(valSet.Validators))))},
	})

	totalSupply := sdk.NewCoins()
	for _, b := range balances {
		// add genesis acc tokens to total supply
		totalSupply = totalSupply.Add(b.Coins...)
	}

	// update total supply
	bankGenesis := banktypes.NewGenesisState(banktypes.DefaultGenesisState().Params, balances, totalSupply, []banktypes.Metadata{}, []banktypes.SendEnabled{})
	genesisState[banktypes.ModuleName] = codec.MustMarshalJSON(bankGenesis)
	// println("genesisStateWithValSet bankState:", string(genesisState[banktypes.ModuleName]))

	// update mint genesis
	mintGenesis := minttypes.DefaultGenesisState()
	mintGenesis.Params.MintDenom = bondDenom
	genesisState[minttypes.ModuleName] = codec.MustMarshalJSON(mintGenesis)

	// update crisis genesis
	crisisGenesis := crisistypes.DefaultGenesisState()
	crisisGenesis.ConstantFee.Denom = bondDenom
	genesisState[crisistypes.ModuleName] = codec.MustMarshalJSON(crisisGenesis)

	return genesisState
}

// KeyTestPubAddr generates a new secp256k1 keypair.
func KeyTestPubAddr() (cryptotypes.PrivKey, cryptotypes.PubKey, sdk.AccAddress) {
	key := secp256k1.GenPrivKey()
	pub := key.PubKey()
	addr := sdk.AccAddress(pub.Address())
	return key, pub, addr
}

// CreateRandomAccounts is a function return a list of randomly generated AccAddresses
func CreateRandomAccounts(numAccts int) []AddressWithKeys {
	testAddrsWithKeys := make([]AddressWithKeys, numAccts)
	for i := 0; i < numAccts; i++ {
		priv := secp256k1.GenPrivKey()
		pub := priv.PubKey()
		testAddrsWithKeys[i].PrivKey = priv
		testAddrsWithKeys[i].PubKey = pub
		testAddrsWithKeys[i].Address = sdk.AccAddress(pub.Address())
	}

	return testAddrsWithKeys
}

// Convert bech32 into sdk.Address and returns it. Panic otherwise.
func MustAcc(s string) sdk.AccAddress {
	a, err := sdk.AccAddressFromBech32(s)
	if err != nil {
		panic(err)
	}
	return a
}

// Convert bech32 into sdk.ValAddress and return it. Panic otherwise.
func MustVal(s string) sdk.ValAddress {
	v, err := sdk.ValAddressFromBech32(s)
	if err != nil {
		panic(err)
	}
	return v
}

// Ensure a validator exists, and return it. Panic otherwise.
func MustExistValidator(app *app.App, ctx sdk.Context, valAddr sdk.ValAddress) stakingtypes.Validator {
	v, err := app.AppKeepers.StakingKeeper.GetValidator(ctx, valAddr)
	if err != nil {
		panic(err)
	}
	return v
}

// Create a minimal validator if it does not exist yet.
func MakeValidator(app *app.App, ctx sdk.Context, valAddr sdk.ValAddress) (stakingtypes.Validator, error) {
	val, err := app.AppKeepers.StakingKeeper.GetValidator(ctx, valAddr)
	if err == nil {
		return val, fmt.Errorf("validator already exists: %s", valAddr.String())
	}
	// Create a dummy validator with small power
	val = stakingtypes.Validator{
		OperatorAddress: valAddr.String(),
		Status:          stakingtypes.Bonded,
		Tokens:          math.NewInt(1),
		DelegatorShares: math.LegacyOneDec(),
		Commission:      stakingtypes.NewCommission(math.LegacyZeroDec(), math.LegacyZeroDec(), math.LegacyZeroDec()),
	}
	err = app.AppKeepers.StakingKeeper.SetValidator(ctx, val)
	if err != nil {
		return stakingtypes.Validator{}, err
	}
	err = app.AppKeepers.DistrKeeper.Hooks().AfterValidatorCreated(ctx, valAddr)
	if err != nil {
		return stakingtypes.Validator{}, err
	}
	return val, err
}

// Create a minimal validator if it does not exist yet. Panic otherwise.
func MustMakeValidator(app *app.App, ctx sdk.Context, valAddr sdk.ValAddress) stakingtypes.Validator {
	val, err := MakeValidator(app, ctx, valAddr)
	if err != nil {
		panic(err)
	}
	return val
}

// Mint and send coin to target account. Panic on error.
func MustFundAccount(app *app.App, ctx sdk.Context, addr sdk.AccAddress, denom string, amount int64) {
	amt := math.NewInt(amount)
	err := app.AppKeepers.BankKeeper.MintCoins(ctx, minttypes.ModuleName, sdk.NewCoins(sdk.NewCoin(denom, amt)))
	if err != nil {
		panic(err)
	}
	err = app.AppKeepers.BankKeeper.SendCoinsFromModuleToAccount(ctx, minttypes.ModuleName, addr, sdk.NewCoins(sdk.NewCoin(denom, amt)))
	if err != nil {
		panic(err)
	}
}

// Create a delegation. Panic otherwise.
func MustCreateDelegation(app *app.App, ctx sdk.Context, delegator sdk.AccAddress, valAddr sdk.ValAddress, amount int64) {
	val := MustExistValidator(app, ctx, valAddr)
	_, err := app.AppKeepers.StakingKeeper.Delegate(ctx, delegator, math.NewInt(amount), stakingtypes.Unbonded, val, true)
	if err != nil {
		panic(err)
	}
}

// Create an unbonding delegation.
func CreateUnbondingDelegation(app *app.App, ctx sdk.Context, delegator sdk.AccAddress, valAddr sdk.ValAddress, amount int64) (time.Time, math.Int, error) {
	shares := math.LegacyNewDec(amount)
	return app.AppKeepers.StakingKeeper.Undelegate(ctx, delegator, valAddr, shares)
}

// Create a redelegation.
func CreateRedelegation(app *app.App, ctx sdk.Context, delegator sdk.AccAddress, valAddrSrc sdk.ValAddress, valAddrDst sdk.ValAddress, amount int64) (completionTime time.Time, err error) {
	shares := math.LegacyNewDecFromInt(math.NewInt(amount))
	return app.AppKeepers.StakingKeeper.BeginRedelegation(ctx, delegator, valAddrSrc, valAddrDst, shares)
}

// Create an unbonding delegation. Panic otherwise.
func MustCreateUnbondingDelegation(app *app.App, ctx sdk.Context, delegator sdk.AccAddress, valAddr sdk.ValAddress, amount int64) {
	_, _, err := CreateUnbondingDelegation(app, ctx, delegator, valAddr, amount)
	if err != nil {
		panic(err)
	}
}

// Create a redelegation. Panic otherwise.
func MustCreateRedelegation(app *app.App, ctx sdk.Context, delegator sdk.AccAddress, valAddrSrc sdk.ValAddress, valAddrDst sdk.ValAddress, amount int64) {
	_, err := CreateRedelegation(app, ctx, delegator, valAddrSrc, valAddrDst, amount)
	if err != nil {
		panic(err)
	}
}
