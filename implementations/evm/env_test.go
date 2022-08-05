package evm_test

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/p2pcloud/protocol"
	"github.com/p2pcloud/protocol/implementations/evm"
	"github.com/p2pcloud/protocol/implementations/evm/broker"
	"github.com/p2pcloud/protocol/pkg/keyring"
	"github.com/stretchr/testify/require"
)

type TestEnv struct {
	Admin      protocol.P2PCloudProtocolIface
	Users      []protocol.P2PCloudProtocolIface
	blockchain *evm.InMemBlockChain
}

const ChainIDSimulated = 1337

func CreateTestEnv(t *testing.T, usersCount int) *TestEnv {
	blockchain := evm.NewWrappedSimulatedBlockchainEnv(t)

	testEnv := &TestEnv{
		blockchain: blockchain,
	}
	testEnv.CreateAdmin(t)
	testEnv.CreateUsers(t, usersCount)
	return testEnv
}

func (te *TestEnv) CreateAdmin(t *testing.T) {
	adminPk, err := te.blockchain.Origin.GetNextPrivateKey()
	require.NoError(t, err)

	idkWhyWeNeedThat := make(chan<- common.Address)

	adminBrokerContract, err := broker.NewBroker(
		te.blockchain.Origin.Backend,
		adminPk,
		"",
		ChainIDSimulated,
		te.blockchain.WaitForTx,
		idkWhyWeNeedThat,
	)
	require.NoError(t, err)

	brokerContractAddress, err := adminBrokerContract.DeployContract()
	require.NoError(t, err)

	admin, err := evm.NewEvmImplementationFromBackend(
		te.blockchain.Origin.Backend,
		keyring.EncodePrivateKey(adminPk),
		brokerContractAddress,
		ChainIDSimulated,
		te.blockchain.WaitForTx,
	)
	require.NoError(t, err)

	te.Admin = admin
}

func (te *TestEnv) CreateUsers(t *testing.T, usersCount int) {
	users := make([]protocol.P2PCloudProtocolIface, usersCount)

	for i := 0; i < usersCount; i++ {
		userPk, err := te.blockchain.Origin.GetNextPrivateKey()
		require.NoError(t, err)

		users[i], err = evm.NewEvmImplementationFromBackend(
			te.blockchain.Origin.Backend,
			keyring.EncodePrivateKey(userPk),
			te.Admin.ContractAddress().Hex(),
			ChainIDSimulated,
			te.blockchain.WaitForTx,
		)
	}

	te.Users = users
}
