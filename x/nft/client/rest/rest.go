package rest

import (
	"github.com/cosmos/cosmos-sdk/client/context"

	"github.com/gorilla/mux"
)

func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
	r.HandleFunc("/nft/{hash}", QueryNFTHandlerFn(cliCtx)).Methods("GET")
	r.HandleFunc("/nft/add", AddNFTokenHandlerFn(cliCtx)).Methods("POST")
	r.HandleFunc("/nft/transfer", TransferNFTokenHandlerFn(cliCtx)).Methods("POST")
}
