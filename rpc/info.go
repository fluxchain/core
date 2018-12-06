package rpc

import (
	"encoding/json"
	"net/http"

	"github.com/fluxchain/core/blockchain"
	"github.com/fluxchain/core/parameters"
)

func GetInfo(w http.ResponseWriter, req *http.Request) {
	info := &nodeInfo{
		Network: parameters.Current().Name,
		Height:  blockchain.Active.Tip.Header.Height,
	}

	json.NewEncoder(w).Encode(info)
}

type nodeInfo struct {
	Network string `json:"network"`
	Height  uint64 `json:"height"`
}
