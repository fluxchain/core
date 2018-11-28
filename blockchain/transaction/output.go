package transaction

import "github.com/fluxchain/core/wallet"

type Output struct {
	Amount    uint32          `json:"amount"`
	Recipient *wallet.Address `json:"recipient"`
}
