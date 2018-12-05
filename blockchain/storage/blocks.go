package storage

import (
	"bytes"
	"encoding/binary"
	"encoding/json"

	"github.com/boltdb/bolt"
	"github.com/fluxchain/core/blockchain/block"
	c "github.com/fluxchain/core/crypto"
)

func StoreBlock(b *block.Block) error {
	serialized, err := serializeBlock(b)
	if err != nil {
		return err
	}

	err = db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BLOCK_BUCKET))
		err := bucket.Put([]byte(b.Header.Hash), serialized)

		return err
	})
	if err != nil {
		return err
	}

	buffer := new(bytes.Buffer)
	if err := binary.Write(buffer, binary.BigEndian, b.Header.Height); err != nil {
		return err
	}

	err = db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BLOCK_HEIGHT_BUCKET))
		err := bucket.Put(buffer.Bytes(), b.Header.Hash)

		return err
	})

	return err
}

func GetBlock(hash c.Hash) (*block.Block, error) {
	var result *block.Block
	err := db.View(func(tx *bolt.Tx) error {
		var err error
		b := tx.Bucket([]byte(BLOCK_BUCKET))

		data := b.Get(hash)
		result, err = deserializeBlock(data)
		if err != nil {
			return err
		}

		return nil
	})

	return result, err
}

func GetBlockByHeight(height uint64) (*block.Block, error) {
	var hash []byte

	buffer := new(bytes.Buffer)
	if err := binary.Write(buffer, binary.BigEndian, height); err != nil {
		return nil, err
	}

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BLOCK_HEIGHT_BUCKET))

		hash = b.Get(buffer.Bytes())

		return nil
	})

	if hash == nil {
		return nil, nil
	}

	return getBlockByIndex(hash)
}

func WalkBlocks(walkFn func(*block.Block) error) {
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BLOCK_HEIGHT_BUCKET))

		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			block, err := getBlockByIndex(v)
			if err != nil {
				return err
			}
			if err := walkFn(block); err != nil {
				return err
			}
		}

		return nil
	})
}

func getBlockByIndex(index []byte) (*block.Block, error) {
	var result *block.Block

	err := db.View(func(tx *bolt.Tx) error {
		var err error
		b := tx.Bucket([]byte(BLOCK_BUCKET))

		data := b.Get(index)
		result, err = deserializeBlock(data)
		if err != nil {
			return err
		}

		return nil
	})

	return result, err
}

func hasBlockIndex(index []byte) (bool, error) {
	var result bool

	err := db.View(func(tx *bolt.Tx) error {
		var err error
		b := tx.Bucket([]byte(BLOCK_BUCKET))

		data := b.Get(index)
		if err != nil {
			result = false
			return err
		}

		result = data != nil

		return nil
	})

	return result, err
}

func LastBlock() (*block.Block, error) {
	var result *block.Block
	err := db.View(func(tx *bolt.Tx) error {
		var err error
		b := tx.Bucket([]byte(BLOCK_BUCKET))

		c := b.Cursor()
		_, data := c.Last()

		result, err = deserializeBlock(data)
		if err != nil {
			return err
		}

		return nil
	})

	return result, err
}

func serializeBlock(b *block.Block) ([]byte, error) {
	return json.Marshal(b)
}

func deserializeBlock(data []byte) (*block.Block, error) {
	var result *block.Block

	err := json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
