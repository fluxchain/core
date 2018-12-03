package main

import (
	"os"
	"time"

	"github.com/fluxchain/core/blockchain"
	"github.com/fluxchain/core/blockchain/block"
	"github.com/fluxchain/core/blockchain/storage"
	"github.com/fluxchain/core/blockchain/transaction"
	"github.com/fluxchain/core/consensus"
	c "github.com/fluxchain/core/crypto"
	"github.com/fluxchain/core/parameters"
)

func main() {
	var err error
	var coinbase *transaction.Transaction
	var hash c.Hash

	err = storage.OpenDatabase("database.db")
	if err != nil {
		panic(err)
	}
	defer storage.CloseDatabase()

	err = storage.Migrate()
	if err != nil {
		panic(err)
	}

	parameters.Set(parameters.Main)
	mainchain := blockchain.NewBlockchain()

	genesisBlock, err := parameters.Current().GenesisBlock()
	if err != nil {
		panic(err)
	}

	err = mainchain.AddBlock(genesisBlock)
	if err != nil {
		panic(err)
	}

	body := block.NewBody()
	coinbase, err = transaction.NewCoinbase("rsyBe3AcPF61VFMi48phGcfsLyvho4mr", 1500, time.Now())
	if err != nil {
		panic(err)
	}

	err = body.AddTransaction(coinbase)
	if err != nil {
		panic(err)
	}

	nextBlock := block.NewBlock(genesisBlock, time.Now(), body)

	hash, err = consensus.GeneratePoW(nextBlock.Header, parameters.Current().MinimumPoW)
	if err != nil {
		panic(err)
	}
	nextBlock.Header.Hash = hash

	err = mainchain.AddBlock(nextBlock)
	if err != nil {
		panic(err)
	}

	file, err := os.Create("blockchain.json")

	if err != nil {
		panic(err)
	}

	mainchain.Serialize(file)

	file.Close()
}
