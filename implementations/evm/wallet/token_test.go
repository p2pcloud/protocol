package wallet_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/stretchr/testify/require"

	"github.com/p2pcloud/protocol/implementations/evm"
	"github.com/p2pcloud/protocol/implementations/evm/wallet"
)

func getTestInstances(t *testing.T, count int) ([]*wallet.Token, *backends.SimulatedBackend) {
	blockchainSim, err := evm.NewSimulatedBlockchainEnv()
	if err != nil {
		t.Fatal(err)
	}

	pk, err := blockchainSim.GetNextPrivateKey()
	if err != nil {
		t.Fatal(err)
	}
	contract0, err := wallet.NewToken(blockchainSim.Backend, pk, "", ChainIDSimulated, GasLimit)
	if err != nil {
		t.Fatal(err)
	}
	contractAddress, err := contract0.Deploy()
	if err != nil {
		t.Fatal(err)
	}
	blockchainSim.Backend.Commit()

	result := make([]*wallet.Token, 0)
	for i := 0; i < count; i++ {
		pk, err := blockchainSim.GetNextPrivateKey()
		if err != nil {
			t.Fatal(err)
		}
		contract, err := wallet.NewToken(blockchainSim.Backend, pk, contractAddress, ChainIDSimulated, GasLimit)
		if err != nil {
			t.Fatal(err)
		}
		result = append(result, contract)
	}
	return result, blockchainSim.Backend
}

func TestToken_TransferTokens(t *testing.T) {
	ctx := context.Background()

	testInstances, simChain := getTestInstances(t, 2)
	defer simChain.Close()

	contr1 := testInstances[0]
	contr2 := testInstances[1]

	amount, err := contr1.BalanceTokens()
	require.NoError(t, err)

	ethBalance, err := contr1.GetBalanceWei()
	require.NoError(t, err)
	fmt.Println(ethBalance)

	err = contr1.TransferTokens(ctx, *contr1.GetMyAddress(), *contr2.GetMyAddress(), 10)
	require.NoError(t, err)
	simChain.Commit()

	amount, err = contr2.BalanceTokens()
	require.NoError(t, err)
	fmt.Println(amount)
}
