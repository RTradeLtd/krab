package krab_test

import (
	"reflect"
	"testing"

	"github.com/RTradeLtd/krab"
	ci "github.com/libp2p/go-libp2p-core/crypto"
)

func Test_Export_Import(t *testing.T) {
	type args struct {
		keyType int
		size    int
	}
	tests := []struct {
		name string
		args args
	}{
		{"EDKey-Success", args{ci.Ed25519, 256}},
		{"RSAKey-Success", args{ci.RSA, 2048}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pk1, _, err := ci.GenerateKeyPair(tt.args.keyType, tt.args.size)
			if err != nil {
				t.Fatal(err)
			}
			mnemonic, err := krab.ExportKeyAsMnemonic(pk1)
			if err != nil {
				t.Fatal(err)
			}
			pk2, err := krab.MnemonicToKey(mnemonic)
			if err != nil {
				t.Fatal(err)
			}
			if valid := pk1.Equals(pk2); !valid {
				t.Fatal("failed to properly recover key")
			}
			if !reflect.DeepEqual(pk1, pk2) {
				t.Fatal("bad reflect")
			}
		})
	}
}
