package evm_test

import (
	"testing"

	"github.com/p2pcloud/protocol"
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
	require.EqualValues(t, 0, locked)
	require.EqualValues(t, 0, free)

	err = user.DepositStablecoin(1000)
	require.NoError(t, err)

	testEnv.RequireStablecoinBalance(*testEnv.Users[0].GetMyAddress(), 0)

	free, locked, err = user.GetStablecoinBalance()
	require.NoError(t, err)
	require.EqualValues(t, 0, locked)
	require.EqualValues(t, 1000, free)

	err = user.WithdrawStablecoin(300)
	require.NoError(t, err)

	free, locked, err = user.GetStablecoinBalance()
	require.NoError(t, err)
	require.EqualValues(t, 0, locked)
	require.EqualValues(t, 700, free)

	testEnv.RequireStablecoinBalance(*testEnv.Users[0].GetMyAddress(), 300)

	//withdraw more than available
	err = user.WithdrawStablecoin(701)
	require.Error(t, err)

	err = user.WithdrawStablecoin(700)
	require.NoError(t, err)

	free, locked, err = user.GetStablecoinBalance()
	require.NoError(t, err)
	require.EqualValues(t, 0, locked)
	require.EqualValues(t, 0, free)
}

//book a VM and check that you cannot withdraw pps*week funds
func TestWithdrawLocked(t *testing.T) {
	var err error

	const PPS = 1
	const INITATL_BALANCE = uint64(10000000)
	const BALANCE_REQUIRED_FOR_1_VM = uint64(60 * 60 * 24 * 7 * PPS)

	testEnv := CreateTestEnv(t, 2)
	user := testEnv.Users[0]
	miner := testEnv.Users[1]

	//add offer
	miner.AddOffer(protocol.Offer{
		VmTypeId:      0,
		PPS:           PPS,
		Availablility: 1,
	}, "https://hello.world")
	offers, err := user.GetAvailableOffers(0)
	require.Len(t, offers, 1)
	require.NoError(t, err)

	//add some money
	testEnv.AddSomeStablecoin(*user.GetMyAddress(), INITATL_BALANCE)
	err = user.DepositStablecoin(INITATL_BALANCE)
	require.NoError(t, err)

	free, locked, err := user.GetStablecoinBalance()
	require.NoError(t, err)
	require.EqualValues(t, 0, locked)
	require.Equal(t, INITATL_BALANCE, free)

	err = user.BookVM(0)
	require.NoError(t, err)

	free, locked, err = user.GetStablecoinBalance()
	require.NoError(t, err)
	require.Equal(t, BALANCE_REQUIRED_FOR_1_VM, locked)
	require.Equal(t, INITATL_BALANCE-BALANCE_REQUIRED_FOR_1_VM, free)

	err = user.WithdrawStablecoin(free + 1)
	require.Error(t, err)

	err = user.WithdrawStablecoin(free)
	require.NoError(t, err)
}

//TODO: call brokers transfer without calling allowance
