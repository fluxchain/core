package main

import (
	"os"
	"time"

	"github.com/fluxchain/core/blockchain"
)

func main() {
	var err error
	var coinbase *blockchain.Transaction

	mainchain := blockchain.NewBlockchain()

	coinbase, err = blockchain.NewCoinbase("rsyBe3AcPF61VFMi48phGcfsLyvho4mr", 2000, time.Now())
	if err != nil {
		panic(err)
	}

	genesisBody := blockchain.NewBlockBody()
	err = genesisBody.AddTransaction(coinbase)
	if err != nil {
		panic(err)
	}

	genesisBlock := blockchain.NewGenesisBlock(time.Now(), genesisBody)
	err = mainchain.AddBlock(genesisBlock)
	if err != nil {
		panic(err)
	}

	body := blockchain.NewBlockBody()
	coinbase, err = blockchain.NewCoinbase("rsyBe3AcPF61VFMi48phGcfsLyvho4mr", 1500, time.Now())
	if err != nil {
		panic(err)
	}

	err = body.AddTransaction(coinbase)
	if err != nil {
		panic(err)
	}

	nextBlock := blockchain.NewBlock(genesisBlock, time.Now(), body)
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
