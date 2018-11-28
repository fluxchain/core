package crypto

import (
	"encoding/hex"
	"encoding/json"
)

// Simple wrapper for byte arrays so I can easily write it to strings in an
// user understandable manner.
type Hash []byte

func (h Hash) String() string {
	return hex.EncodeToString(h)
}

func (h Hash) MarshalJSON() ([]byte, error) {
	data, err := json.Marshal(h.String())
	if err != nil {
		return nil, err
	}

	return data, nil
}
