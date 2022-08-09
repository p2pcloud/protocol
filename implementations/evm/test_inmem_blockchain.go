package evm

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/p2pcloud/protocol/pkg/keyring"
)

const (
	CommunityFee = 5
)

type InMemBlockChain struct {
	Origin *SimulatedBlockchainEnv
}

func (b *InMemBlockChain) GetNextPrivateKey() (*ecdsa.PrivateKey, error) {
	return b.Origin.GetNextPrivateKey()
}

func (b *InMemBlockChain) WaitForTx(tx *types.Transaction) error {
	b.Origin.Backend.Commit()

	receipt, err := b.Origin.Backend.TransactionReceipt(context.Background(), tx.Hash())
	_ = receipt
	_ = err
	// receipt.Status

	return nil
}

func NewWrappedSimulatedBlockchainEnv(t *testing.T) *InMemBlockChain {
	bc, err := NewSimulatedBlockchainEnv()
	if err != nil {
		t.Fatal(err)
	}

	return &InMemBlockChain{Origin: bc}
}

type SimulatedBlockchainEnv struct {
	Backend      *backends.SimulatedBackend
	privateKeys  []*ecdsa.PrivateKey
	nextKeyIndex int
}

func (s *SimulatedBlockchainEnv) GetNextPrivateKey() (*ecdsa.PrivateKey, error) {
	if s.nextKeyIndex >= len(s.privateKeys) {
		return nil, fmt.Errorf("no more private keys to test, got %d", len(s.privateKeys))
	}
	key := s.privateKeys[s.nextKeyIndex]
	s.nextKeyIndex++
	return key, nil
}

func NewSimulatedBlockchainEnv() (*SimulatedBlockchainEnv, error) {
	s := &SimulatedBlockchainEnv{nextKeyIndex: 0}

	privateKeys, err := getPrivateKeys()
	s.privateKeys = privateKeys

	if err != nil {
		return nil, err
	}

	genesisAlloc := map[common.Address]core.GenesisAccount{}

	balance := new(big.Int)
	balance.SetString("10000000000000000000", 10) // 10 eth in wei

	for _, key := range privateKeys {
		address := crypto.PubkeyToAddress(key.PublicKey)
		genesisAlloc[address] = core.GenesisAccount{
			Balance: balance,
		}
	}

	blockGasLimit := uint64(8000000)
	backend := backends.NewSimulatedBackend(genesisAlloc, blockGasLimit)
	s.Backend = backend
	return s, nil
}

func getPrivateKeys() ([]*ecdsa.PrivateKey, error) {
	privateKeys := make([]*ecdsa.PrivateKey, 0)

	privateKeyHexes := []string{
		"1b6e407dd3395c73b8599faf5fdda54c4111999217012c81e19e3b0599eee423",
		"153f16b54b94bc80f6f53e5e57fece88732732fed35b7d4623e84f913f462db0",
		"ac124c244c46256ddef17b8fdfc9c470aee7bb2357e69ca842d08726ca2e4459",
		"6cf569a5ca13672cd3d8aeb337174d9a3bbbedfbdf3ba4cd0e9dcfcf9f1ff876",
		"0be9e3a9890e85b24a7d10aa80e12bfd1a0ff1826e0bc14c0eaadf82781c8987",
		"648fc67b513e03688adbd99a7143704f7641e3df84b833896807f04a1af03041",
		"efa076cc0c402ff9339a3ea8ab68e635ab4a9a10fed9449034421837cb956f70",
		"eefe2c282a4c82e4f3b00d6a3af975874189cf3287a7441d58853883e8f93a39",
		"f75fee772a7bfd826fbba2e95361191da52ab9327cecdbee97aeb0abdee8b3cd",
		"c1493bda5ce8684674a6dc8617ab6e8b28367d46fcec22e65c736781790139de",
	}

	for _, privateKeyHex := range privateKeyHexes {
		privateKey, err := keyring.DecodePrivateKey(privateKeyHex)
		if err != nil {
			return nil, err
		}

		privateKeys = append(privateKeys, privateKey)
	}
	return privateKeys, nil
}
