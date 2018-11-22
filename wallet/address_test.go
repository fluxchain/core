package wallet

import (
	"fmt"
	"testing"
)

func getTestKey() []byte {
	return []byte{
		0xe4, 0x47, 0xae, 0x7,
		0x93, 0xf4, 0x70, 0xf4,
		0x97, 0x35, 0x25, 0xf8,
		0xb, 0x77, 0x48, 0xfe,
		0xd3, 0x0, 0x6a, 0x74,
		0xcd, 0x15, 0x5f, 0xc9,
		0x26, 0x68, 0x7c, 0x1a,
		0xad, 0x13, 0x2a, 0xae,
		0xf6, 0xef, 0x44, 0xed,
		0xd9, 0xa2, 0xe9, 0x51,
		0xda, 0x7e, 0xde, 0x41,
		0xf8, 0xc0, 0x34, 0xf1,
		0x6a, 0xa8, 0xdb, 0x8,
		0x74, 0xd8, 0xe8, 0xcc,
		0xed, 0xa6, 0x5d, 0x2,
		0x49, 0xe6, 0x12, 0xd0}
}

func TestAddressValidationFromKey(t *testing.T) {
	pk := getTestKey()
	address := NewAddressFromPublicKey(pk)

	if !address.Valid() {
		t.Error("could not validate address derived from public key")
	}
}

func TestAddressValidationString(t *testing.T) {
	tests := []struct {
		Address  string
		Expected bool
	}{
		{"rsyBe3AcPF61VFMi48phGcfsLyvho4mr", true},
		{"asyBe3AcPF61VFMi48phGcfsLyvho4mr", false},
	}

	for _, tc := range tests {
		tc := tc // capture range variable
		t.Run(tc.Address, func(t *testing.T) {
			t.Parallel()
			address := NewAddressFromString(tc.Address)

			if result := address.Valid(); result != tc.Expected {
				t.Errorf("validation for %v resulted in %v where %v was expected",
					tc.Address,
					result,
					tc.Expected)
			}
		})

	}
}

func ExampleAddressCreationFromSetPublicKey() {
	pk := getTestKey()
	address := NewAddressFromPublicKey(pk)

	fmt.Println(address.String)
	// Output: rsyBe3AcPF61VFMi48phGcfsLyvho4mr
}
