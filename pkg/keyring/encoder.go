package keyring

import (
	"crypto/ecdsa"
	"encoding/hex"

	"github.com/ethereum/go-ethereum/crypto"
)

func EncodePrivateKey(privateKey *ecdsa.PrivateKey) string {
	return hex.EncodeToString(crypto.FromECDSA(privateKey))
}

func DecodePrivateKey(keyHex string) (*ecdsa.PrivateKey, error) {
	if keyHex[:2] == "0x" {
		keyHex = keyHex[2:]
	}
	return crypto.HexToECDSA(keyHex)

}
