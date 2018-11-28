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

func (h *Hash) UnmarshalJSON(data []byte) error {
	var buf string
	var err error

	err = json.Unmarshal(data, &buf)
	if err != nil {
		return err
	}

	*h, err = hex.DecodeString(buf)
	if err != nil {
		return err
	}

	return nil
}
