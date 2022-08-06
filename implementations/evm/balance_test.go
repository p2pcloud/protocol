package evm_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRevertedTx(t *testing.T) {
	testEnv := CreateTestEnv(t, 2)

	user := testEnv.Users[0]

	err := user.WithdrawStablecoin(701)
	require.Error(t, err)
}

func TestDepositWithdraw(t *testing.T) {
	var err error
	testEnv := CreateTestEnv(t, 2)
	testEnv.AddSomeStablecoin(*testEnv.Users[0].GetMyAddress(), 1000)
	testEnv.RequireStablecoinBalance(*testEnv.Users[0].GetMyAddress(), 1000)

	user := testEnv.Users[0]

	free, locked, err := user.GetStablecoinBalance()
	require.NoError(t, err)
	require.Equal(t, 0, locked)
	require.Equal(t, 0, free)

	err = user.DepositStablecoin(1000)
	require.NoError(t, err)

	testEnv.RequireStablecoinBalance(*testEnv.Users[0].GetMyAddress(), 0)

	free, locked, err = user.GetStablecoinBalance()
	require.NoError(t, err)
	require.Equal(t, 0, locked)
	require.Equal(t, 1000, free)

	err = user.WithdrawStablecoin(300)
	require.NoError(t, err)

	free, locked, err = user.GetStablecoinBalance()
	require.NoError(t, err)
	require.Equal(t, 0, locked)
	require.Equal(t, 700, free)

	testEnv.RequireStablecoinBalance(*testEnv.Users[0].GetMyAddress(), 300)

	//withdraw more than available
	err = user.WithdrawStablecoin(701)
	require.Error(t, err)

	err = user.WithdrawStablecoin(700)
	require.NoError(t, err)
}

//TODO: book a VM and check that you cannot withdraw pps*week funds
