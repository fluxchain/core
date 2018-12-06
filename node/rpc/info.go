package rpc

import (
	"encoding/json"
	"net/http"

	"github.com/fluxchain/core/blockchain"
)

func GetInfo(w http.ResponseWriter, req *http.Request) {
	info := &nodeInfo{
		Height: blockchain.Active.Tip.Header.Height,
	}

	json.NewEncoder(w).Encode(info)
}

type nodeInfo struct {
	Height uint64 `json:"height"`
}
