package types

import (
	"github.com/cosmos/cosmos-sdk/x/auth"
	"strings"
)

type EstimateGasReq struct {
	StdTx         auth.StdTx `json:"tx"`
	GasAdjustment string     `json:"adj"`
}

type EstimateGasResp struct {
	Gas      uint64 `json:"estimated"`
	Adjusted uint64 `json:"adjusted"`
}

func NewEstimateGasReq(StdTx auth.StdTx, GasAdjustment string) EstimateGasReq {
	return EstimateGasReq{
		StdTx:         StdTx,
		GasAdjustment: strings.TrimSpace(GasAdjustment),
	}
}

func NewEstimateGasResp(Gas uint64, Adjusted uint64) EstimateGasResp {
	return EstimateGasResp{
		Gas:      Gas,
		Adjusted: Adjusted,
	}
}

func (egr EstimateGasReq) Sanitize() EstimateGasReq {
	if len(egr.StdTx.GetSignatures()) == 0 {
		egr.StdTx.Signatures = []auth.StdSignature{{}}
	}

	return NewEstimateGasReq(egr.StdTx, egr.GasAdjustment)
}

func (egr EstimateGasReq) ValidateBasic() error {
	err := egr.StdTx.ValidateBasic()
	if err != nil {
		return err
	}

	return nil
}
