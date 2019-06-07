package krab

import (
	"fmt"
	"strings"

	mnemonics "github.com/RTradeLtd/entropy-mnemonics"
	ci "github.com/libp2p/go-libp2p-core/crypto"
)

// all of this is taken from https://github.com/ipfs/go-ipfs-keystore

var (
	// ErrNoSuchKey is returned if a key of the given name is not found in the store
	ErrNoSuchKey = "no key by the given name was found"
	// ErrKeyExists is returned when writing a key would overwrite an existing key
	ErrKeyExists = "key by that name already exists, refusing to overwrite"
	// ErrKeyFmt is returned when the key's format is invalid
	ErrKeyFmt = "key has invalid format"
)

// ExportKeyAsMnemonic is used to take an IPFS key, and return a human-readable friendly version.
// The idea is to allow users to easily export the keys they create, allowing them to take control of their records (ipns, tns, etc..)
func ExportKeyAsMnemonic(pk ci.PrivKey) (string, error) {
	pkBytes, err := pk.Bytes()
	if err != nil {
		return "", err
	}
	phrase, err := mnemonics.ToPhrase(pkBytes, mnemonics.English)
	if err != nil {
		return "", err
	}
	return phrase.String(), nil
}

// MnemonicToKey takes an exported mnemonic phrase, and converts it to a private key
func MnemonicToKey(phrase string) (ci.PrivKey, error) {
	mnemonicBytes, err := mnemonics.FromString(phrase, mnemonics.English)
	if err != nil {
		return nil, err
	}
	return ci.UnmarshalPrivateKey(mnemonicBytes)
}

func validateName(name string) error {
	if name == "" {
		return fmt.Errorf("%s: key names must be at least one character", ErrKeyFmt)
	}

	if strings.Contains(name, "/") {
		return fmt.Errorf("%s: key names may not contain slashes", ErrKeyFmt)
	}

	if strings.HasPrefix(name, ".") {
		return fmt.Errorf("%s: key names may not begin with a period", ErrKeyFmt)
	}

	return nil
}
