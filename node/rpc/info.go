package rpc

import (
	"encoding/json"
	"net/http"
)

func GetInfo(w http.ResponseWriter, req *http.Request) {
	info := &nodeInfo{
		Height: 100,
	}

	json.NewEncoder(w).Encode(info)
}

type nodeInfo struct {
	Height uint64 `json:"height"`
}
