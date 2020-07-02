package rest

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"strings"

	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/x/auth"
	authrest "github.com/cosmos/cosmos-sdk/x/auth/client/rest"
)

type EstimateGasReq struct {
	StdTx         auth.StdTx `json:"tx"`
	GasAdjustment string     `json:"adj"`
}

type EstimateGasResp struct {
	Gas      uint64 `json:"estimated"`
	Adjusted uint64 `json:"adjusted"`
}

func NewBaseReq(stdTx auth.StdTx, gasAdjustment string) EstimateGasReq {
	return EstimateGasReq{
		StdTx:         stdTx,
		GasAdjustment: strings.TrimSpace(gasAdjustment),
	}
}

func (egr EstimateGasReq) Sanitize() EstimateGasReq {
	return NewBaseReq(egr.StdTx, egr.GasAdjustment)
}

func (egr EstimateGasReq) ValidateBasic() error {
	err := egr.StdTx.ValidateBasic()
	if err != nil {
		return err
	}

	return nil
}

func RegisterTxRoutes(cliCtx context.CLIContext, r *mux.Router) {
	authrest.RegisterTxRoutes(cliCtx, r)
	r.HandleFunc("/txs/estimateGas", QueryEstimateGasHandlerFn(cliCtx)).Methods("GET")
}
