// Copyright 2018 The Cloq Authors
// This file is part of the cloq-core library .
//
// The cloq-core library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The cloq-core library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the cloq-core library. If not, see <http://www.gnu.org/licenses/>.

// Package crypto provides wrappers for the ed25519 implementation
package crypto

import (
	"crypto/rand"
	"encoding/base32"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"golang.org/x/crypto/ed25519"
)

// PublicKey type for ed25519 keys
type PublicKey ed25519.PublicKey

// PrivateKey type for ed25518 keys
type PrivateKey ed25519.PrivateKey

var (
	errBadPrivateKeySize = errors.New("wrong private key size ,private keys are 64 byte long")
	errBadPublicKeySize  = errors.New("wrong public key size , public keys are 32 byte long")
)

// GenerateKey generates an ed25519 key pair
func GenerateKey() (ed25519.PublicKey, ed25519.PrivateKey, error) {

	publicKey, privateKey, err := ed25519.GenerateKey(rand.Reader)

	if err != nil {

		return nil, nil, err

	}

	return publicKey, privateKey, nil

}

// Sign signs a message using ed25519 private key
// Note that ed25519 is a EdDSA algorithm .
// It includes hashing by default using SHA-512
// Golang implementation does two passes thus all message are signed as s .
// Ref: https://tools.ietf.org/html/rfc8032
func Sign(privateKey ed25519.PrivateKey, message []byte) ([]byte, error) {

	// Sign will panic if privateKey is not 32 bytes therefore we check before

	if len(privateKey) != ed25519.PrivateKeySize {
		return nil, errors.New("bad private key size")
	}
	sig := ed25519.Sign(privateKey, message)

	return sig, nil
}

// Verify verifies whether a signature passed as parameter has been signed
// by the given public key
func Verify(publicKey ed25519.PublicKey, message []byte, signature []byte) (bool, error) {
	if len(publicKey) != ed25519.PublicKeySize {
		return false, errors.New("bad public key size")
	}
	if !ed25519.Verify(publicKey, message, signature) {
		return false, errors.New("wrong signature")
	}
	return true, nil
}

// SavePrivateKey will save a base32 encoded private key to a file
func SavePrivateKey(filename string, key ed25519.PrivateKey) error {

	k := base32.StdEncoding.EncodeToString(key)

	err := ioutil.WriteFile(filename, []byte(k), 0600)

	if err != nil {
		return err
	}
	return nil

}

// LoadPrivateKey will load a base32 encoded private key from a file
func LoadPrivateKey(filename string) (ed25519.PrivateKey, error) {
	var privateKey ed25519.PrivateKey
	var err error
	privateKey, err = ioutil.ReadFile(filename)

	if err != nil {
		return nil, errors.New("failed to read from file : ")
	}

	return privateKey, nil
}

// SavePrivateKeyHex will save a hex encoded private key to a file
func SavePrivateKeyHex(filename string, privateKey ed25519.PrivateKey) error {

	k := ToHex(privateKey)

	err := ioutil.WriteFile(filename, []byte(k), 0600)

	if err != nil {
		return err
	}
	return nil
}

// LoadPrivateKeyHex will load a hex encoded private key from a file
func LoadPrivateKeyHex(filename string) (ed25519.PrivateKey, error) {

	// buffer to hold bytes read from file
	buf := make([]byte, ed25519.PrivateKeySize)

	// get a file descriptor / handle to a file
	handle, err := os.Open(filename)

	if err != nil {
		return nil, errors.New("failed to open file for reading")
	}

	// defer the handle close
	defer handle.Close()

	// read key size bytes from file
	n, err := handle.Read(buf)

	if n != ed25519.PrivateKeySize {
		return nil, errBadPrivateKeySize
	}

	key, err := ToPrivateKey(buf)

	if err != nil {
		return nil, err
	}

	return key, nil
}

// HexToPrivateKey decodes a hex string to a ed25519 PrivateKey
func HexToPrivateKey(hexstring string) (ed25519.PrivateKey, error) {

	b, err := FromHex(hexstring)

	if err != nil {
		return nil, errors.New("failed to decode hexstring")
	}

	key, err := ToPrivateKey(b)

	if err != nil {
		return nil, err
	}

	return key, nil

}

// ToPrivateKey encodes bytes as a ed25519 PrivateKey
func ToPrivateKey(b []byte) (ed25519.PrivateKey, error) {

	return toEd(b)

}

// GetPublicKey derives the Public Key from an ed25519 private key
// and does type assertion
func GetPublicKey(privateKey ed25519.PrivateKey) (ed25519.PublicKey, error) {

	publicKey := privateKey.Public()

	edPublicKey, ok := publicKey.(ed25519.PublicKey)

	if !ok {
		return nil, errors.New("failed to derive a correct ed25519 public key")
	}

	return edPublicKey, nil
}

// toEd encodes bytes as ed25519 private keys with verifications
func toEd(b []byte) (ed25519.PrivateKey, error) {

	var key ed25519.PrivateKey

	if len(b) != ed25519.PrivateKeySize {
		return nil, errBadPrivateKeySize
	}

	key = ed25519.PrivateKey(b)

	return key, nil

}

// saveKeyPair saves a PEM encoded Private Key and Public Key in separate files this is just
// for debugging purposes as ed25519 keys are small there are faster and better ways to encode
// and save to disk
func saveKeyPair(publicKey ed25519.PublicKey, privateKey ed25519.PrivateKey) error {

	pubKeyBlock := &pem.Block{
		Type:    "EdDSA-PUBLICKEY",
		Headers: nil,
		Bytes:   publicKey,
	}

	privKeyBlock := &pem.Block{
		Type:    "EdDSA-PRIVATEKEY",
		Headers: nil,
		Bytes:   privateKey,
	}

	keyOut, err := os.OpenFile("privatekey.pem", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		fmt.Println("failed to open publickey.pem with : ", err)

	}

	err = pem.Encode(keyOut, privKeyBlock)

	keyOut.Close()

	if err != nil {
		return errors.New("failed to save private key")
	}

	keyOut, err = os.OpenFile("publickey.pem", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		fmt.Println("failed to open publickey.pem with : ", err)
	}

	err = pem.Encode(keyOut, pubKeyBlock)

	keyOut.Close()

	if err != nil {
		return errors.New("failed to save public key")
	}

	return nil
}
