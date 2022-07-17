package stablecoin_test

import (
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"

	"github.com/p2pcloud/protocol/implementations/evm"
	"github.com/p2pcloud/protocol/implementations/evm/stablecoin"
)

func TestIntegrationStableCoin_DepositCoin(t *testing.T) {
	ganacheJsonKeys := os.Getenv("GANACHE_KEYS")
	ganacheRPCEndpoint := os.Getenv("GANACHE_RPC_ENDPOINT")

	if ganacheJsonKeys == "" || ganacheRPCEndpoint == "" {
		t.Skip()
	}

	web3Client, err := ethclient.Dial(ganacheRPCEndpoint)
	require.NoError(t, err)

	helper, err := evm.NewGanacheBCHelper(5, ganacheJsonKeys)
	require.NoError(t, err)

	userIdx := 0

	p, err := evm.InitializeTestInstances(
		2,
		5,
		*evm.NewGifts(map[int]int64{userIdx: 3},
			map[int]int64{userIdx: 3}),
		web3Client,
		helper,
	)
	require.NoError(t, err)

	userBroker := p.Contracts[userIdx]

	user1 := stablecoin.New(&stablecoin.Params{
		Decimals: p.Decimals,
		Backend:  p.Backend,
		Session:  userBroker.GetStableCoinSession(),
		Commit:   p.BcHelper.Commit,
	})

	// deposit for contract address and then check balances
	require.NoError(t, user1.DepositCoin(2))

	tokens, err := user1.UserTokenBalance()
	require.NoError(t, err)
	require.Equal(t, int64(1), tokens)

	brokerDeployerBalance, err := p.DeployerToken.Balance(*p.BrokerDeployAddress)
	require.NoError(t, err)
	require.Equal(t, int64(2), brokerDeployerBalance)

	balance, err := user1.Balance()
	require.NoError(t, err)
	require.Equal(t, int64(2), balance)
}

func TestIntegrationStableCoin_WithdrawCoin(t *testing.T) {
	ganacheJsonKeys := os.Getenv("GANACHE_KEYS")
	ganacheRPCEndpoint := os.Getenv("GANACHE_RPC_ENDPOINT")

	if ganacheJsonKeys == "" || ganacheRPCEndpoint == "" {
		t.Skip()
	}

	web3Client, err := ethclient.Dial(ganacheRPCEndpoint)
	require.NoError(t, err)

	helper, err := evm.NewGanacheBCHelper(5, ganacheJsonKeys)
	require.NoError(t, err)

	userIdx := 0

	p, err := evm.InitializeTestInstances(
		2,
		5,
		*evm.NewGifts(map[int]int64{userIdx: 3},
			map[int]int64{userIdx: 3}),
		web3Client,
		helper,
	)
	require.NoError(t, err)

	userBroker := p.Contracts[userIdx]

	user1 := stablecoin.New(&stablecoin.Params{
		Decimals: p.Decimals,
		Backend:  p.Backend,
		Session:  userBroker.GetStableCoinSession(),
		Commit:   p.BcHelper.Commit,
	})

	// deposit for contract address and then check balances
	require.NoError(t, user1.DepositCoin(2))

	tokens, err := user1.UserTokenBalance()
	require.NoError(t, err)
	require.Equal(t, int64(1), tokens)

	brokerDeployerBalance, err := p.DeployerToken.Balance(*p.BrokerDeployAddress)
	require.NoError(t, err)
	require.Equal(t, int64(2), brokerDeployerBalance)

	balance, err := user1.Balance()
	require.NoError(t, err)
	require.Equal(t, int64(2), balance)

	require.NoError(t, user1.WithdrawCoin(2))

	tokens, err = user1.UserTokenBalance()
	require.NoError(t, err)
	require.Equal(t, int64(3), tokens)

	brokerDeployerBalance, err = p.DeployerToken.Balance(*p.BrokerDeployAddress)
	require.NoError(t, err)
	require.Equal(t, int64(0), brokerDeployerBalance)

	balance, err = user1.Balance()
	require.NoError(t, err)
	require.Equal(t, int64(0), balance)
}
