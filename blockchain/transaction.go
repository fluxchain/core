package blockchain

import (
	"github.com/cbergoon/merkletree"
	"github.com/fluxchain/core/crypto"
	"golang.org/x/crypto/ed25519"
)

// Transaction object
type Transaction struct {
	Inputs     []*Input           `json:"inputs"`
	Outputs    []*Output          `json:"outputs"`
	Kernel     bool               `json:"kernel"`
	IOHash     crypto.Hash        `json:"hash"`
	Signatures []crypto.Signature `json:"signature"`
}

// Input are spent coins
type Input struct {
	TxID      crypto.Hash      `json:"txid"`
	Index     uint64           `json:"index"`
	PublicKey crypto.PublicKey `json:"publickey"`
}

// Output are newly created coins
type Output struct {
	Value      uint64      `json:"amount"`
	PubKeyHash crypto.Hash `json:"pubkeyhash"`
	Locktime   uint64      `json:"locktime"`
}

// Hash calculates the transaction inner hash by concatting the
// inputs and outputs this is verified hash the one we sign different from txid which is h(TX)
func (tx Transaction) Hash() (crypto.Hash, error) {

	// TODO : Add serialization routines
	return crypto.NilHash(), nil
}

// Equals compares two transaction 
func (tx Transaction) Equals(other merkletree.Content) (bool, error) {
	return t.Hash.String() == other.(Transaction).Hash.String(), nil
}

// NewKernelInput returns the default input for a kernel transaction .
// Kernel fields can be modified to help find a nonce for the proof of work .
func NewKernelInput() *Input {

	var zeroHash c.Hash
	var pk ed25519.PublicKey

	return &Input{
		TxHash:    zeroHash,
		Index:     0,
		PublicKey: pk,
	}

}

// NewKernel creates a new kernel transaction
func NewKernel(outputs []*Output) *Transaction {

	kerinputs := []*Input{NewKernelInput()}
	return &Transaction{Kernel: true,
		Inputs:  kerinputs,
		Outputs: outputs}
}

// NewTransaction creates a new regular transaction
func NewTransaction(inputs []*Input, outputs []*Output, sigs []*crypto.Signature) *Transaction {
	return &Transaction{
		Version:    P2PKH,
		Kernel:     false,
		IOHash:     c.NilHash(),
		Inputs:     inputs,
		Outputs:    outputs,
		Signatures: sigs,
	}
}
