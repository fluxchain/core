package crypto

import (
	"encoding/json"
	"testing"
)

var marshalCases = []struct {
	Name    string
	Hash    Hash
	Encoded string
}{
	{"8 zero bytes", Hash{0x00, 0x00, 0x00, 0x00}, "00000000"},
	{"encoded output is lowercase", Hash{0xAB, 0xCD, 0xEF}, "abcdef"},
}

func TestAddressJsonMarshal(t *testing.T) {
	for _, tc := range marshalCases {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()

			marshalled, err := json.Marshal(tc.Hash)

			if err != nil {
				t.Errorf("could not serialize hash, json error: %v", err.Error())
			}

			if string(marshalled) != "\""+tc.Encoded+"\"" {
				t.Errorf("encoded hash [\"%v\"] does not seem identical to expected output [%v]", tc.Encoded, string(marshalled))
			}

		})
	}
}

func TestAddressJsonUnmarshal(t *testing.T) {
	for _, tc := range marshalCases {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			hash := &Hash{}

			data := []byte("\"" + tc.Encoded + "\"")
			err := json.Unmarshal(data, hash)
			if err != nil {
				t.Errorf("could not deserialize hash, json error: %v", err.Error())
			}

			if hash.String() != tc.Encoded {
				t.Errorf("decoded hash [\"%v\"] does not seem identical to expected output [%v]",
					hash.String(),
					tc.Encoded)
			}
		})
	}
}
