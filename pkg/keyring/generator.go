package keyring

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/crypto"
)

func GenPrivateKey() (*ecdsa.PrivateKey, error) {
	return crypto.GenerateKey()
}
