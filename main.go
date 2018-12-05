package main

import (
	"time"

	"github.com/fluxchain/core/blockchain"
	"github.com/fluxchain/core/blockchain/block"
	"github.com/fluxchain/core/blockchain/storage"
	"github.com/fluxchain/core/blockchain/transaction"
	"github.com/fluxchain/core/consensus"
	"github.com/fluxchain/core/parameters"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	logrus.SetLevel(logrus.DebugLevel)

	logrus.Info("starting flux...")
	logrus.Info("opening database")
	if err := storage.OpenDatabase("database.db"); err != nil {
		logrus.Error("could not open local database", err)
	}
	defer storage.CloseDatabase()

	if err := storage.Migrate(); err != nil {
		logrus.Error("could not migrate database structure to local database: ", err)
	}

	parameters.Set(parameters.Main)
	mainchain := blockchain.NewBlockchain()

	hasGenesis, err := mainchain.HasGenesis()
	if err != nil {
		logrus.Error("error looking up genesis existence", err)
	}

	if !hasGenesis {
		logrus.Info("database does not seem to include genesis, adding it")

		genesis, err := parameters.Current().GenesisBlock()
		if err != nil {
			logrus.Error("could not create genesis block from selected parameters: ", err)
		}

		if err := mainchain.AddBlock(genesis); err != nil {
			logrus.Error("could not add genesis block to local database: ", err)
		}
	}

	if err := mainchain.Hydrate(); err != nil {
		logrus.Error("could not read local database during hydrate: ", err)
	}

	for i := 0; i < 10; i++ {
		coinbase, err := transaction.NewCoinbase("rsyBe3AcPF61VFMi48phGcfsLyvho4mr", 1500, time.Now())
		if err != nil {
			logrus.Error("could not create coinbase for block: ", err)
		}

		body := block.NewBody()
		if err := body.AddTransaction(coinbase); err != nil {
			logrus.Error("could not add coinbase transaction to block body: ", err)
		}

		nextBlock := block.NewBlock(mainchain.Tip, time.Now(), body)
		hash, err := consensus.GeneratePoW(nextBlock.Header, parameters.Current().MinimumPoW)
		if err != nil {
			logrus.Error("could not generate PoW for block: ", err)
		}
		nextBlock.Header.Hash = hash

		if err := mainchain.AddBlock(nextBlock); err != nil {
			logrus.Error("could not add newly mined block to local chain: ", err)
		}
	}
}
