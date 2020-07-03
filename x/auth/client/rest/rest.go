package rest

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/gorilla/mux"

	rest "github.com/cosmos/cosmos-sdk/x/auth/client/rest"
)

func RegisterTxRoutes(cliCtx context.CLIContext, r *mux.Router) {
	r.HandleFunc("/txs/estimate_gas", QueryEstimateGasHandlerFn(cliCtx)).Methods("POST")
	rest.RegisterTxRoutes(cliCtx, r)
}
