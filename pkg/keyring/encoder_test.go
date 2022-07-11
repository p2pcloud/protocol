package keyring_test

import (
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/Incognida/protocol/pkg/keyring"
)

func TestEncodeDecode(t *testing.T) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		t.Fatal(err)
	}

	encPriv := keyring.EncodePrivateKey(privateKey)

	priv2, err := keyring.DecodePrivateKey(encPriv)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(privateKey, priv2) {
		t.Fatal("Private keys do not match.")
	}
}
