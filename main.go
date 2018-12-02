package main

import (
	"os"
	"time"

	"github.com/fluxchain/core/blockchain"
	"github.com/fluxchain/core/blockchain/block"
	"github.com/fluxchain/core/blockchain/transaction"
	"github.com/fluxchain/core/parameters"
)

func main() {
	var err error
	var coinbase *transaction.Transaction

	parameters.Set(parameters.Main)

	mainchain := blockchain.NewBlockchain()

	coinbase, err = transaction.NewCoinbase("rsyBe3AcPF61VFMi48phGcfsLyvho4mr", 2000, time.Now())
	if err != nil {
		panic(err)
	}

	genesisBody := block.NewBody()
	err = genesisBody.AddTransaction(coinbase)
	if err != nil {
		panic(err)
	}

	genesisBlock := block.NewGenesisBlock(time.Now(), genesisBody)
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
