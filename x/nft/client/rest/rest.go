package rest

import (
	"github.com/cosmos/cosmos-sdk/client/context"

	"github.com/gorilla/mux"
)

func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
	r.HandleFunc("/nft/{hash}", QueryNFTHandlerFn(cliCtx)).Methods("GET")
	r.HandleFunc("/nft/mint", MintHandlerFn(cliCtx)).Methods("POST")
	r.HandleFunc("/nft/burn", BurnHandlerFn(cliCtx)).Methods("POST")
	r.HandleFunc("/nft/transfer", TransferHandlerFn(cliCtx)).Methods("POST")
	r.HandleFunc("/nft/transfer", TransferHandlerFn(cliCtx)).Methods("POST")
}
