package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client/context"
)

func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
	r.HandleFunc("/contract/{hash}", QueryContractHandlerFn(cliCtx)).Methods("GET")
	r.HandleFunc("/contract", AddContractRequest(cliCtx)).Methods("POST")
}
