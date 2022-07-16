package broker_test

import (
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"

	"github.com/p2pcloud/protocol"
	"github.com/p2pcloud/protocol/implementations/evm"
	"github.com/p2pcloud/protocol/implementations/evm/broker"
	"github.com/p2pcloud/protocol/implementations/evm/token"
)

const ChainIDSimulated = 1337

func getTestInstances(t *testing.T, supply int64, count int) ([]protocol.BrokerIface, *backends.SimulatedBackend) {
	blockchainSim, err := evm.NewSimulatedBlockchainEnv()
	if err != nil {
		t.Fatal(err)
	}

	tokenPk, err := blockchainSim.GetNextPrivateKey()
	if err != nil {
		t.Fatal(err)
	}

	tkn := token.NewToken(&token.Params{
		Decimals:   6,
		Backend:    blockchainSim.Backend,
		PrivateKey: tokenPk,
		ChainID:    ChainIDSimulated,
		Commit:     blockchainSim.Backend.Commit,
	})

	pk, err := blockchainSim.GetNextPrivateKey()
	if err != nil {
		t.Fatal(err)
	}

	tokenAddress, err := tkn.DeployContract(supply)
	if err != nil {
		t.Fatal(err)
	}

	contract0, err := broker.NewBroker(
		blockchainSim.Backend, pk, "", ChainIDSimulated, *tokenAddress,
	)
	if err != nil {
		t.Fatal(err)
	}
	contractAddress, err := contract0.DeployContracts()
	if err != nil {
		t.Fatal(err)
	}
	blockchainSim.Backend.Commit()

	result := make([]protocol.BrokerIface, 0)
	for i := 0; i < count; i++ {
		pk, err := blockchainSim.GetNextPrivateKey()
		if err != nil {
			t.Fatal(err)
		}
		contract, err := broker.NewBroker(
			blockchainSim.Backend, pk, contractAddress[0], ChainIDSimulated, *tokenAddress,
		)
		if err != nil {
			t.Fatal(err)
		}
		result = append(result, contract)
	}
	return result, blockchainSim.Backend
}

func TestGetMtlsHash(t *testing.T) {
	testInstances, simChain := getTestInstances(t, 0, 2)
	defer simChain.Close()

	contr1 := testInstances[0]

	hash := strings.ToLower("0x9b920402017534cd42AC81e738E0F23a7162c4D7")

	err := contr1.RegisterMtlsHashIfNeeded(hash)
	check(t, err)
	simChain.Commit()

	contr2 := testInstances[1]

	loadedHash, err := contr2.GetMtlsHash(contr1.GetMyAddress())
	check(t, err)

	if loadedHash != hash {
		t.Fatalf("hashes are not equal: expected %s, got %s", hash, loadedHash)
	}
}
