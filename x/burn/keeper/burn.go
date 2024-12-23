package keeper

import (
	"github.com/firmachain/firmachain/v05/x/burn/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) BurnFromModuleAccount(ctx sdk.Context) error {

	// firma1sk06e3dyexuq4shw77y3dsv480xv42mqkx0y4x  (fixed)
	burnAddress := k.accountKeeper.GetModuleAccount(ctx, types.ModuleName).GetAddress()
	balances := k.bankKeeper.GetAllBalances(ctx, burnAddress)

	//fmt.Printf("burnaddress : %s\n", moduleAccount.GetAddress().String())
	//const ufctTokenName = "ufct"
	//balance := k.bankKeeper.GetBalance(ctx, burnAddress, ufctTokenName)

	if !balances.IsZero() {
		err := k.bankKeeper.BurnCoins(ctx, "burn", balances)

		if err != nil {
			return err
		}
	}

	return nil
}
