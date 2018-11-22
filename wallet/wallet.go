package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/json"
	"io"
	"log"
)

type Wallet struct {
	PrivateKey *ecdsa.PrivateKey `json:"private"`
	PublicKey  []byte            `json:"public"`
	Address    string            `json:"address"`
}

func (w *Wallet) ExportWallet(out io.Writer) error {
	return json.NewEncoder(out).Encode(w)
}

func NewWallet() *Wallet {
	// get a curve
	c := elliptic.P256()
	// generate a private key
	private, err := ecdsa.GenerateKey(c, rand.Reader)
	if err != nil {
		log.Panic(err)
	}

	// create a public key
	public := append(private.PublicKey.X.Bytes(), private.PublicKey.Y.Bytes()...)

	wallet := &Wallet{PrivateKey: private,
		PublicKey: public}

	return wallet
}
