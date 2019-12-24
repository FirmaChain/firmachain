package firmachain

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/staking"
	abci "github.com/tendermint/tendermint/abci/types"
)

type GenesisState struct {
	StakingData     staking.GenesisState `json:"staking"`
	ContractRecords []Contract           `json:"contract_records"`
}

func NewGenesisState(contractRecords []Contract, stakingData staking.GenesisState) GenesisState {

	return GenesisState{
		ContractRecords: nil,
		StakingData:     stakingData,
	}
}

func ValidateGenesis(data GenesisState) error {
	for _, record := range data.ContractRecords {
		if record.Owner == nil {
			return fmt.Errorf("invalid ContractRecords: Value: %s. Error: Missing Owner", record.Owner)
		}
		if record.Path == "" {
			return fmt.Errorf("invalid ContractRecords: Value: %s. Error: Missing Path", record.Path)
		}
		if record.Hash == "" {
			return fmt.Errorf("invalid ContractRecords: Value: %s. Error: Missing Hash", record.Hash)
		}

	}
	return nil
}

func DefaultGenesisState() GenesisState {
	// Set custom params
	stakingGenState := staking.DefaultGenesisState()
	stakingGenState.Params.UnbondingTime = time.Second * 60 * 60 * 24 * 7 * 3 // three weeks
	stakingGenState.Params.MaxValidators = 11
	stakingGenState.Params.BondDenom = StoreKey

	return GenesisState{
		StakingData:     stakingGenState,
		ContractRecords: []Contract{},
	}
}

func InitGenesis(ctx sdk.Context, keeper Keeper, data GenesisState) []abci.ValidatorUpdate {
	for _, record := range data.ContractRecords {
		keeper.SetContract(ctx, record.Path, record.Hash, record.Owner)
	}
	return []abci.ValidatorUpdate{}
}

func ExportGenesis(ctx sdk.Context, k Keeper) GenesisState {
	var records []Contract
	iterator := k.GetContractsIterator(ctx)
	for ; iterator.Valid(); iterator.Next() {

		hash := string(iterator.Key())
		whois := k.GetContract(ctx, hash)
		records = append(records, whois)

	}
	return GenesisState{ContractRecords: records}
}
