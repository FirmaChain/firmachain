package simulation

import (
	"encoding/hex"
	"github.com/cosmos/cosmos-sdk/simapp/helpers"
	"math/rand"

	"github.com/tendermint/tendermint/crypto"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/firmachain/FirmaChain/x/contract/keeper"
	"github.com/firmachain/FirmaChain/x/contract/types"
)

// Simulation operation weights constants
const (
	OpWeightMsgAddContract = "op_weight_msg_add_contract"
)

// WeightedOperations returns all the operations from the module with their respective weights
func WeightedOperations(appParams simulation.AppParams, cdc *codec.Codec, ak authkeeper.AccountKeeper,
	ck keeper.Keeper) simulation.WeightedOperations {

	var weightMsgAddContract int
	appParams.GetOrGenerate(cdc, OpWeightMsgAddContract, &weightMsgAddContract, nil,
		func(_ *rand.Rand) {
			weightMsgAddContract = simappparams.DefaultWeightMsgSend
		},
	)

	return simulation.WeightedOperations{
		simulation.NewWeightedOperation(
			weightMsgAddContract,
			SimulateMsgAddContract(ak),
		),
	}
}

// SimulateMsgAddContract tests and runs a MsgAddContract
// nolint: funlen
func SimulateMsgAddContract(ak authkeeper.AccountKeeper) simulation.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simulation.Account, chainID string,
	) (simulation.OperationMsg, []simulation.FutureOperation, error) {

		path := dummyIPFSPath()
		hash := randomFileHash()
		simAccount, _ := simulation.RandomAcc(r, accs)

		msg := types.NewMsgAddContract(path, hash, simAccount.Address)

		err := sendMsgAddContract(r, app, ak, msg, ctx, chainID, []crypto.PrivKey{simAccount.PrivKey})
		if err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, err
		}

		return simulation.NewOperationMsg(msg, true, ""), nil, nil
	}
}

// sendMsgAddContract sends a transaction with a MsgAddContract from a provided random account.
func sendMsgAddContract(
	r *rand.Rand, app *baseapp.BaseApp, ak authkeeper.AccountKeeper, // nolint:interfacer
	msg types.MsgAddContract, ctx sdk.Context, chainID string, privkeys []crypto.PrivKey,
) error {

	account := ak.GetAccount(ctx, msg.Owner)
	coins := account.SpendableCoins(ctx.BlockTime())

	var (
		fees sdk.Coins
		err  error
	)

	fees, err = simulation.RandomFees(r, ctx, coins)
	if err != nil {
		return err
	}

	tx := helpers.GenTx(
		[]sdk.Msg{msg},
		fees,
		helpers.DefaultGenTxGas,
		chainID,
		[]uint64{account.GetAccountNumber()},
		[]uint64{account.GetSequence()},
		privkeys...,
	)

	_, _, err = app.Deliver(tx)
	if err != nil {
		return err
	}

	return nil
}

func dummyIPFSPath() string {
	return "QmTF7NerdGZhnDPJj3Yj51gqH18o8kLtgkgtVjMLk1V9tx"
}

func randomFileHash() string {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return "0000000000000000000000000000000000000000000000000000000000000000"
	}
	return hex.EncodeToString(bytes)
}
