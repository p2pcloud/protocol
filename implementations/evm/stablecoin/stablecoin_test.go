package stablecoin_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/p2pcloud/protocol/implementations/evm"
	"github.com/p2pcloud/protocol/implementations/evm/stablecoin"
)

func TestStableCoin_DepositCoin(t *testing.T) {
	blockchain := evm.NewWrappedSimulatedBlockchainEnv(t)

	userIdx := 0

	p, err := evm.InitializeTestInstances(
		2,
		5,
		evm.NewGifts(map[int]int64{userIdx: 3},
			map[int]int64{userIdx: 3}),
		blockchain.Origin.Backend,
		blockchain,
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

func TestStableCoin_WithdrawCoin(t *testing.T) {
	blockchain := evm.NewWrappedSimulatedBlockchainEnv(t)

	userIdx := 0

	p, err := evm.InitializeTestInstances(
		2,
		5,
		evm.NewGifts(map[int]int64{userIdx: 3},
			map[int]int64{userIdx: 3}),
		blockchain.Origin.Backend,
		blockchain,
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
