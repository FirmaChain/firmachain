package rest

import (
	"net/http"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client/context"

	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"

	"github.com/cosmos/cosmos-sdk/x/auth"
)

func QueryEstimateGasHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req EstimateGasReq

		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			return
		}

		req = req.Sanitize()
		err := req.ValidateBasic()

		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		adjustment, err := strconv.ParseFloat(req.GasAdjustment, 64)

		if adjustment == 0 {
			adjustment = 1
		}

		var tx = req.StdTx

		tx.Signatures = []auth.StdSignature{{}}
		txBytes, err := utils.GetTxEncoder(cliCtx.Codec)(tx)

		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		estimated, adjusted, err := utils.CalculateGas(cliCtx.QueryWithData, cliCtx.Codec, txBytes, adjustment)

		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		rest.PostProcessResponse(w, cliCtx, EstimateGasResp{Gas: estimated, Adjusted: adjusted})
	}
}
