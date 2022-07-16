package stablecoin_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/p2pcloud/protocol/implementations/evm/stablecoin"
)

func TestStableCoin_DepositCoin(t *testing.T) {
	userIdx := 0

	p, err := instantiate(2, 5, gifts{
		userInitialBalances: map[int]int64{userIdx: 3},
		userAllowances:      map[int]int64{userIdx: 3},
	})
	require.NoError(t, err)

	userBroker := p.contracts[userIdx]

	user1 := stablecoin.New(&stablecoin.Params{
		Decimals: p.decimals,
		Backend:  p.bcSim.Backend,
		Session:  userBroker.GetStableCoinSession(),
		Commit:   p.bcSim.Backend.Commit,
	})

	// deposit for contract address and then check balances
	require.NoError(t, user1.DepositCoin(2))

	tokens, err := user1.UserTokenBalance()
	require.NoError(t, err)
	require.Equal(t, int64(1), tokens)

	brokerDeployerBalance, err := p.deployerToken.Balance(*p.brokerDeployAddress)
	require.NoError(t, err)
	require.Equal(t, int64(2), brokerDeployerBalance)

	balance, err := user1.Balance()
	require.NoError(t, err)
	require.Equal(t, int64(2), balance)
}

func TestStableCoin_WithdrawCoin(t *testing.T) {
	userIdx := 0

	p, err := instantiate(2, 5, gifts{
		userInitialBalances: map[int]int64{userIdx: 3},
		userAllowances:      map[int]int64{userIdx: 3},
	})
	require.NoError(t, err)

	userBroker := p.contracts[userIdx]

	user1 := stablecoin.New(&stablecoin.Params{
		Decimals: p.decimals,
		Backend:  p.bcSim.Backend,
		Session:  userBroker.GetStableCoinSession(),
		Commit:   p.bcSim.Backend.Commit,
	})

	// deposit for contract address and then check balances
	require.NoError(t, user1.DepositCoin(2))

	tokens, err := user1.UserTokenBalance()
	require.NoError(t, err)
	require.Equal(t, int64(1), tokens)

	brokerDeployerBalance, err := p.deployerToken.Balance(*p.brokerDeployAddress)
	require.NoError(t, err)
	require.Equal(t, int64(2), brokerDeployerBalance)

	balance, err := user1.Balance()
	require.NoError(t, err)
	require.Equal(t, int64(2), balance)

	require.NoError(t, user1.WithdrawCoin(2))

	tokens, err = user1.UserTokenBalance()
	require.NoError(t, err)
	require.Equal(t, int64(3), tokens)

	brokerDeployerBalance, err = p.deployerToken.Balance(*p.brokerDeployAddress)
	require.NoError(t, err)
	require.Equal(t, int64(0), brokerDeployerBalance)

	balance, err = user1.Balance()
	require.NoError(t, err)
	require.Equal(t, int64(0), balance)
}
