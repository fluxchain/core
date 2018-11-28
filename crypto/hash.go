package crypto

import (
	"encoding/hex"
)

// Simple wrapper for byte arrays so I can easily write it to strings in an
// user understandable manner.
type Hash []byte

func (h Hash) String() string {
	return hex.EncodeToString(h)
}
