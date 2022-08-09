package evm_test

import (
	"crypto/ecdsa"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/p2pcloud/protocol"
	"github.com/p2pcloud/protocol/implementations/evm"
	"github.com/p2pcloud/protocol/implementations/evm/broker"
	"github.com/p2pcloud/protocol/implementations/evm/token"
	"github.com/p2pcloud/protocol/pkg/keyring"
	"github.com/stretchr/testify/require"
)

type TestEnv struct {
	Admin                   protocol.P2PCloudProtocolIface
	Users                   []protocol.P2PCloudProtocolIface
	blockchain              *evm.InMemBlockChain
	adminStablecoinContract *token.Token
	t                       *testing.T
}

const ChainIDSimulated = 1337

func CreateTestEnv(t *testing.T, usersCount int) *TestEnv {
	blockchain := evm.NewWrappedSimulatedBlockchainEnv(t)

	testEnv := &TestEnv{
		blockchain: blockchain,
		t:          t,
	}

	adminPk, err := blockchain.Origin.GetNextPrivateKey()
	require.NoError(t, err)

	tokenAddress := testEnv.deployToken(adminPk)
	brokerAddress := testEnv.deployBroker(adminPk, tokenAddress)
	testEnv.createAdmin(adminPk, brokerAddress)
	testEnv.createUsers(usersCount)
	return testEnv
}

func (te *TestEnv) AdjustTime(adjustment time.Duration) {
	err := te.blockchain.Origin.Backend.AdjustTime(adjustment)
	require.NoError(te.t, err)
	te.blockchain.Origin.Backend.Commit()
}

func (te *TestEnv) AddSomeStablecoin(to common.Address, amount uint64) {
	err := te.adminStablecoinContract.Transfer(to, amount)
	require.NoError(te.t, err)
}

func (te *TestEnv) RequireStablecoinBalance(address common.Address, expectedBalance uint64) {
	balance, err := te.adminStablecoinContract.BalanceOf(address)
	require.NoError(te.t, err)
	require.Equal(te.t, expectedBalance, balance)
}

func (te *TestEnv) deployBroker(adminPk *ecdsa.PrivateKey, tokenAddress common.Address) string {
	idkWhyWeNeedThat := make(chan<- common.Address)

	adminBrokerContract, err := broker.NewBroker(
		te.blockchain.Origin.Backend,
		adminPk,
		"",
		ChainIDSimulated,
		te.blockchain.WaitForTx,
		idkWhyWeNeedThat,
	)
	require.NoError(te.t, err)

	brokerContractAddress, err := adminBrokerContract.DeployContract()

	adminBrokerContract, err = broker.NewBroker(
		te.blockchain.Origin.Backend,
		adminPk,
		brokerContractAddress,
		ChainIDSimulated,
		te.blockchain.WaitForTx,
		idkWhyWeNeedThat,
	)
	require.NoError(te.t, err)

	err = adminBrokerContract.SetStablecoinAddress(tokenAddress)
	require.NoError(te.t, err)

	require.NoError(te.t, err)
	return brokerContractAddress
}

func (te *TestEnv) deployToken(adminPk *ecdsa.PrivateKey) common.Address {
	adminStablecoinContract, err := token.NewToken(
		te.blockchain.Origin.Backend,
		adminPk,
		"",
		ChainIDSimulated,
		te.blockchain.WaitForTx,
	)
	require.NoError(te.t, err)

	tokenAddress, err := adminStablecoinContract.DeployContract()
	require.NoError(te.t, err)

	//recreate with address
	te.adminStablecoinContract, err = token.NewToken(
		te.blockchain.Origin.Backend,
		adminPk,
		tokenAddress.Hex(),
		ChainIDSimulated,
		te.blockchain.WaitForTx,
	)
	require.NoError(te.t, err)

	return *tokenAddress
}

func (te *TestEnv) createAdmin(adminPk *ecdsa.PrivateKey, brokerContractAddress string) {
	admin, err := evm.NewEvmImplementationFromBackend(
		te.blockchain.Origin.Backend,
		keyring.EncodePrivateKey(adminPk),
		brokerContractAddress,
		ChainIDSimulated,
		te.blockchain.WaitForTx,
	)
	require.NoError(te.t, err)

	te.Admin = admin
}

func (te *TestEnv) createUsers(usersCount int) {
	users := make([]protocol.P2PCloudProtocolIface, usersCount)

	for i := 0; i < usersCount; i++ {
		userPk, err := te.blockchain.Origin.GetNextPrivateKey()
		require.NoError(te.t, err)

		users[i], err = evm.NewEvmImplementationFromBackend(
			te.blockchain.Origin.Backend,
			keyring.EncodePrivateKey(userPk),
			te.Admin.ContractAddress().Hex(),
			ChainIDSimulated,
			te.blockchain.WaitForTx,
		)
		require.NoError(te.t, err)
	}

	te.Users = users
}
