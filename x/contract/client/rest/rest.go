package rest

import (
	"github.com/cosmos/cosmos-sdk/client/context"

	"github.com/gorilla/mux"
)

func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
	r.HandleFunc("/contract/{hash}", QueryContractHandlerFn(cliCtx)).Methods("GET")
	r.HandleFunc("/contract", AddContractHandlerFn(cliCtx)).Methods("POST")
}
