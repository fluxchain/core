package rpc

import (
	"encoding/hex"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/fluxchain/core/blockchain/block"
	"github.com/fluxchain/core/blockchain/storage"
	c "github.com/fluxchain/core/crypto"
	"github.com/gorilla/mux"
)

func GetBlock(w http.ResponseWriter, req *http.Request) {
	var err error
	var hash c.Hash
	var block *block.Block
	vars := mux.Vars(req)

	hash, err = hex.DecodeString(vars["block"])
	if err != nil {
		panic(err)
	}

	block, err = storage.GetBlock(hash)
	if err != nil {
		panic(err)
	}

	err = json.NewEncoder(w).Encode(block)
	if err != nil {
		panic(err)
	}
}

func GetBlockHeight(w http.ResponseWriter, req *http.Request) {
	var err error
	var height uint64
	var block *block.Block

	vars := mux.Vars(req)

	height, err = strconv.ParseUint(vars["block"], 10, 64)
	if err != nil {
		panic(err)
	}

	block, err = storage.GetBlockByHeight(height)
	if err != nil {
		panic(err)
	}

	err = json.NewEncoder(w).Encode(block)
	if err != nil {
		panic(err)
	}
}
