package transaction

import (
	"crypto/sha256"
	"encoding/binary"
	"time"

	"github.com/cbergoon/merkletree"
	"github.com/fluxchain/core/util"
	"github.com/fluxchain/core/wallet"
)

type Transaction struct {
	Inputs      []*Input  `json:"inputs"`
	Outputs     []*Output `json:"outputs"`
	Description string    `json:"description"`
	Hash        util.Hash `json:"hash"`
	Timestamp   time.Time `json:"timestamp"`
}

type Input struct {
	Amount      uint32    `json:"amount"`
	Transaction util.Hash `json:"transaction"`
	PublicKey   []byte    `json:"publickey"`
	Signature   []byte    `json:"signature"`
}

type Output struct {
	Amount    uint32          `json:"amount"`
	Recipient *wallet.Address `json:"recipient"`
}

// Calculates the transactions hash by concatting the binary buffers of
// the tx description, the unix timestamp, the inputs their amount, public and
// signature and the outputs their amount and recipient address. And passing that
// through a round of SHA256.
func (tx Transaction) CalculateHash() ([]byte, error) {
	hash := sha256.New()
	hash.Write([]byte(tx.Description))

	// Put the transaction UNIX timestamp into the buffer as BE uint64.
	tsBuf := make([]byte, 8)
	binary.BigEndian.PutUint64(tsBuf, uint64(tx.Timestamp.Unix()))
	hash.Write(tsBuf)

	for _, input := range tx.Inputs {
		hash.Write(input.Transaction)
		buf := make([]byte, 4)
		binary.BigEndian.PutUint32(buf, input.Amount)
		hash.Write(buf)
		hash.Write(input.PublicKey)
		hash.Write(input.Signature)
	}

	for _, output := range tx.Outputs {
		buf := make([]byte, 4)
		binary.BigEndian.PutUint32(buf, output.Amount)
		hash.Write(buf)

		hash.Write([]byte(output.Recipient.String))
	}

	md := hash.Sum(nil)
	return md, nil
}

// Figures out if two transactions are the same, used for the merkletree lib.
func (t Transaction) Equals(other merkletree.Content) (bool, error) {
	return t.Hash.String() == other.(Transaction).Hash.String(), nil
}

// Creates a coinbase transaction which requires no inputs.
func NewCoinbase(recipient string, amount uint32, timestamp time.Time) (*Transaction, error) {
	var err error
	address := wallet.NewAddressFromString(recipient)

	output := &Output{
		Recipient: address,
		Amount:    amount,
	}

	tx := &Transaction{Outputs: []*Output{output}, Timestamp: timestamp}
	tx.Hash, err = tx.CalculateHash()

	return tx, err
}
