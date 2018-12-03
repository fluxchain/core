package storage

import (
	"testing"
	"time"

	"github.com/fluxchain/core/blockchain/block"
	"github.com/fluxchain/core/blockchain/transaction"
	"github.com/fluxchain/core/parameters"
)

func setupTestCase(t *testing.T) func(t *testing.T) {
	parameters.Set(parameters.UnitTest)
	err := OpenDatabase("database.db")
	if err != nil {
		t.Error(err)
	}

	return func(t *testing.T) {
		CloseDatabase()
	}
}

func TestStoringBlock(t *testing.T) {
	teardown := setupTestCase(t)
	defer teardown(t)

	coinbase, err := transaction.NewCoinbase("rsyBe3AcPF61VFMi48phGcfsLyvho4mr", 2000, time.Now())
	if err != nil {
		t.Error(err)
	}

	b := block.NewGenesisBlock(time.Now(), &block.Body{coinbase})
	err = BlockRepository.Store(b)
	if err != nil {
		t.Error(err)
	}

	BlockRepository.Get(b.Header.Hash)
}
