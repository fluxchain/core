package node

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

type Node struct {
	Chain *blockchain.Blockchain
}

// bootstraps the node initing the specified database, setting the
// appropriate parameters, storing the genesis block if not available
// and cycling over all the blocks in the database.
func (n *Node) Bootstrap(databasePath string, currentParameters *parameters.Parameters) {
	logrus.WithFields(logrus.Fields{
		"path": databasePath,
	}).Info("opening database")

	if err := storage.OpenDatabase(databasePath); err != nil {
		logrus.Fatal("could not open local database", err)
	}

	if err := storage.Migrate(); err != nil {
		logrus.Fatal("could not migrate database structure to local database: ", err)
	}

	parameters.Set(currentParameters)
	n.Chain = blockchain.NewBlockchain()

	hasGenesis, err := n.Chain.HasGenesis()
	if err != nil {
		logrus.Fatal("error looking up genesis existence", err)
	}

	// if the genesis isn't available in the local database, fetch one from the
	// selected parameters and store it.
	if !hasGenesis {
		logrus.Info("database does not seem to include genesis, adding it")

		genesis, err := parameters.Current().GenesisBlock()
		if err != nil {
			logrus.Fatal("could not create genesis block from selected parameters: ", err)
		}

		if err := n.Chain.AddGenesisBlock(genesis); err != nil {
			logrus.Fatal("could not add genesis block to local database: ", err)
		}
	}

	if err := n.Chain.Hydrate(); err != nil {
		logrus.Fatal("could not read local database during hydrate: ", err)
	}
}

// clean up the state and close the database, connections and outstanding requests
func (n *Node) Teardown() {
	storage.CloseDatabase()
}

// mines an abritrary amount of blocks in accordance with the selected chain parameters
func (n *Node) Mine(amount uint64) {
	for i := 0; i < 10; i++ {
		coinbase, err := transaction.NewCoinbase("rsyBe3AcPF61VFMi48phGcfsLyvho4mr", 1500, time.Now())
		if err != nil {
			logrus.Error("could not create coinbase for block: ", err)
		}

		body := block.NewBody()
		if err := body.AddTransaction(coinbase); err != nil {
			logrus.Error("could not add coinbase transaction to block body: ", err)
		}

		nextBlock := block.NewBlock(n.Chain.Tip, time.Now(), body)
		hash, err := consensus.GeneratePoW(nextBlock.Header, parameters.Current().MinimumPoW)
		if err != nil {
			logrus.Error("could not generate PoW for block: ", err)
		}
		nextBlock.Header.Hash = hash

		if err := n.Chain.AddBlock(nextBlock); err != nil {
			logrus.Error("could not add newly mined block to local chain: ", err)
		}
	}
}

func New() *Node {
	return &Node{}
}
