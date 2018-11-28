package transaction

import "github.com/fluxchain/core/util"

type Input struct {
	Amount      uint32    `json:"amount"`
	Transaction util.Hash `json:"transaction"`
	PublicKey   []byte    `json:"publickey"`
	Signature   []byte    `json:"signature"`
}
