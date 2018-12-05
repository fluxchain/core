package main

import (
	"time"

	"github.com/fluxchain/core/blockchain"
	"github.com/fluxchain/core/blockchain/block"
	"github.com/fluxchain/core/blockchain/storage"
	"github.com/fluxchain/core/blockchain/transaction"
	"github.com/fluxchain/core/consensus"
	"github.com/fluxchain/core/parameters"
)

func main() {
	if err := storage.OpenDatabase("database.db"); err != nil {
		panic(err)
	}
	defer storage.CloseDatabase()

	if err := storage.Migrate(); err != nil {
		panic(err)
	}

	parameters.Set(parameters.UnitTest)
	mainchain := blockchain.NewBlockchain()

	if !mainchain.HasGenesis() {
		genesis, err := parameters.Current().GenesisBlock()

		if err != nil {
			panic(err)
		}

		if err := mainchain.AddBlock(genesis); err != nil {
			panic(err)
		}
	}

	if err := mainchain.Hydrate(); err != nil {
		panic(err)
	}

	for i := 0; i < 100; i++ {
		coinbase, err := transaction.NewCoinbase("rsyBe3AcPF61VFMi48phGcfsLyvho4mr", 1500, time.Now())
		if err != nil {
			panic(err)
		}

		body := block.NewBody()
		if err := body.AddTransaction(coinbase); err != nil {
			panic(err)
		}

		nextBlock := block.NewBlock(mainchain.Tip, time.Now(), body)
		hash, err := consensus.GeneratePoW(nextBlock.Header, parameters.Current().MinimumPoW)
		if err != nil {
			panic(err)
		}
		nextBlock.Header.Hash = hash

		if err := mainchain.AddBlock(nextBlock); err != nil {
			panic(err)
		}
	}
}
