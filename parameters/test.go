package parameters

import (
	"time"

	"github.com/fluxchain/core/blockchain/block"
	"github.com/fluxchain/core/blockchain/transaction"
	"github.com/fluxchain/core/consensus"
	c "github.com/fluxchain/core/crypto"
)

var Test = &Parameters{
	MinimumPoW: [32]byte{0x00, 0x00, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff},

	GenesisBlock: func() (*block.Block, error) {
		coinbase, err := transaction.NewCoinbase("rsyBe3AcPF61VFMi48phGcfsLyvho4mr", 2000, time.Now())
		if err != nil {
			return nil, err
		}
		genesisBody := block.NewBody()
		err = genesisBody.AddTransaction(coinbase)
		if err != nil {
			return nil, err
		}
		genesisBlock := block.NewGenesisBlock(time.Now(), genesisBody)

		var hash c.Hash
		hash, err = consensus.GeneratePoW(genesisBlock.Header, Current().MinimumPoW)
		if err != nil {
			return nil, err
		}
		genesisBlock.Header.Hash = hash

		return genesisBlock, nil
	},
}
