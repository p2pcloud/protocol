package wallet_test

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"testing"

	"github.com/Incognida/protocol/implementations/evm"
	"github.com/Incognida/protocol/implementations/evm/wallet"
)

const (
	ChainIDSimulated = 1337
	GasLimit = 8000000
)

func TestGetbalance(t *testing.T) {
	blockchainSim, err := evm.NewSimulatedBlockchainEnv()
	if err != nil {
		t.Fatal(err)
	}

	pk, err := blockchainSim.GetNextPrivateKey()
	if err != nil {
		t.Fatal(err)
	}

	w := wallet.NewWalletFromPK(pk, blockchainSim.Backend, ChainIDSimulated)
	balance, err := w.GetBalanceETH()
	if err != nil {
		t.Fatal(err)
	}
	if balance != 10 {
		t.Fatal("balance is not 10")
	}
}

func TestTransfer(t *testing.T) {
	blockchainSim, err := evm.NewSimulatedBlockchainEnv()
	if err != nil {
		t.Fatal(err)
	}

	senderPk, err := blockchainSim.GetNextPrivateKey()
	if err != nil {
		t.Fatal(err)
	}
	senderWallet := wallet.NewWalletFromPK(senderPk, blockchainSim.Backend, ChainIDSimulated)

	receiverPk, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		t.Fatal(err)
	}
	receiverWallet := wallet.NewWalletFromPK(receiverPk, blockchainSim.Backend, ChainIDSimulated)

	balance, err := receiverWallet.GetBalanceETH()
	if err != nil {
		t.Fatal(err)
	}
	if balance != 0 {
		t.Fatal("balance is not 0")
	}

	receiverAdress, err := receiverWallet.GetAddress()
	if err != nil {
		t.Fatal(err)
	}
	err = senderWallet.TransferEth(*receiverAdress, 1.555)
	if err != nil {
		t.Fatal(err)
	}
	blockchainSim.Backend.Commit()

	balance, err = receiverWallet.GetBalanceETH()
	if err != nil {
		t.Fatal(err)
	}
	if balance != 1.555 {
		t.Fatalf("balance is not 1.555. it is %f", balance)
	}
}
