package parameters

import "github.com/fluxchain/core/blockchain/block"

var current *Parameters = nil

type Parameters struct {
	Name         string
	MinimumPoW   [32]byte
	GenesisBlock GenesisBlockFunc
}

type GenesisBlockFunc func() (*block.Block, error)

func IsSet() bool {
	return current != nil
}

func Set(p *Parameters) {
	current = p
}

func Current() *Parameters {
	if !IsSet() {
		panic("no chain parameters set!")
	}

	return current
}
