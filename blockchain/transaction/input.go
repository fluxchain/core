package transaction

import "github.com/fluxchain/core/crypto"

type Input struct {
	Amount      uint32      `json:"amount"`
	Transaction crypto.Hash `json:"transaction"`
	PublicKey   []byte      `json:"publickey"`
	Signature   []byte      `json:"signature"`
}
