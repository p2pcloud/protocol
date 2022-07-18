package broker_test

import (
	"strings"
	"testing"

	"github.com/p2pcloud/protocol/implementations/evm"
)

func TestGetMtlsHash(t *testing.T) {
	blockchain := evm.NewWrappedSimulatedBlockchainEnv(t)

	testInstances, err := evm.InitializeTestInstances(
		2, 6, nil, blockchain.Origin.Backend, blockchain,
	)
	check(t, err)

	contr1 := testInstances.Contracts[0]

	hash := strings.ToLower("0x9b920402017534cd42AC81e738E0F23a7162c4D7")

	err = contr1.RegisterMtlsHashIfNeeded(hash)
	check(t, err)

	contr2 := testInstances.Contracts[1]

	loadedHash, err := contr2.GetMtlsHash(contr1.GetMyAddress())
	check(t, err)

	if loadedHash != hash {
		t.Fatalf("hashes are not equal: expected %s, got %s", hash, loadedHash)
	}
}
