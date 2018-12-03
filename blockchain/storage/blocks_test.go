package storage

import (
	"bytes"
	"testing"
	"time"

	"github.com/fluxchain/core/blockchain/block"
	"github.com/fluxchain/core/blockchain/transaction"
	"github.com/fluxchain/core/parameters"
)

func setupTestCase(t *testing.T) func(t *testing.T) {
	parameters.Set(parameters.UnitTest)
	err := OpenDatabase("../../unittest.db")
	Setup()

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

	mockBlock, err := mockGenesisBlock()
	if err != nil {
		t.Error(err)
	}

	err = StoreBlock(mockBlock)
	if err != nil {
		t.Error(err)
	}
}

func TestStoringAndGettingBlock(t *testing.T) {
	teardown := setupTestCase(t)
	defer teardown(t)

	mockBlock, err := mockGenesisBlock()
	if err != nil {
		t.Error(err)
	}

	err = StoreBlock(mockBlock)
	if err != nil {
		t.Error(err)
	}

	var resultBlock *block.Block
	resultBlock, err = GetBlock(mockBlock.Header.Hash)
	if err != nil {
		t.Error(err)
	}

	if bytes.Compare(mockBlock.Header.Hash, resultBlock.Header.Hash) != 0 {
		t.Fatalf("block hash of origin and result block are not the same! original: %s, result: %s",
			mockBlock.Header.Hash,
			resultBlock.Header.Hash)
	}

	if mockBlock.Header.Height != resultBlock.Header.Height {
		t.Fatalf("block height of origin and result block are not the same! original: %#v, result: %#v",
			mockBlock.Header.Height,
			resultBlock.Header.Height)
	}
}

func mockGenesisBlock() (*block.Block, error) {
	coinbase, err := transaction.NewCoinbase("rsyBe3AcPF61VFMi48phGcfsLyvho4mr", 2000, time.Now())
	if err != nil {
		return nil, err
	}

	result := block.NewGenesisBlock(time.Now(), &block.Body{coinbase})

	return result, nil
}
