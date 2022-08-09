package evm_test

import (
	"log"
	"testing"

	"github.com/p2pcloud/protocol"
	"github.com/stretchr/testify/require"
)

func TestGetBooking(t *testing.T) {
	var err error

	const PPS = uint64(1)
	const INITATL_BALANCE = uint64(10000000)

	testEnv := CreateTestEnv(t, 2)
	user := testEnv.Users[0]
	miner := testEnv.Users[1]

	//add some money
	testEnv.AddSomeStablecoin(*user.GetMyAddress(), INITATL_BALANCE)
	err = user.DepositStablecoin(INITATL_BALANCE)
	require.NoError(t, err)

	//add offer
	miner.AddOffer(protocol.Offer{
		VmTypeId:      333,
		PPS:           PPS,
		Availablility: 2,
	}, "https://hello.world")
	require.NoError(t, err)

	//book
	err = user.BookVM(0)
	require.NoError(t, err)

	bookedAt, err := user.GetTime()
	require.NoError(t, err)

	bookings, err := user.GetUsersBookings()
	require.Len(t, bookings, 1)
	require.NoError(t, err)

	require.EqualValues(t, 333, bookings[0].VmTypeId)
	require.EqualValues(t, PPS, bookings[0].PPS)
	require.EqualValues(t, miner.GetMyAddress().Hex(), bookings[0].Miner.Hex())
	require.EqualValues(t, 0, bookings[0].Index)
	require.EqualValues(t, user.GetMyAddress().Hex(), bookings[0].User.Hex())
	log.Printf("booked at: %d, now: %d", bookings[0].BookedAt, bookedAt)
	require.EqualValues(t, bookedAt, bookings[0].BookedAt)
	require.EqualValues(t, bookedAt, bookings[0].LastPayment)
}

func TestBooking(t *testing.T) {
	var err error

	const PPS = uint64(1)
	const INITATL_BALANCE = uint64(10000000)

	testEnv := CreateTestEnv(t, 2)
	user := testEnv.Users[0]
	miner := testEnv.Users[1]

	//add some money
	testEnv.AddSomeStablecoin(*user.GetMyAddress(), INITATL_BALANCE)
	err = user.DepositStablecoin(INITATL_BALANCE)
	require.NoError(t, err)

	free, locked, err := user.GetStablecoinBalance()
	require.NoError(t, err)
	require.EqualValues(t, 0, locked)
	require.Equal(t, INITATL_BALANCE, free)

	//add offer
	miner.AddOffer(protocol.Offer{
		VmTypeId:      0,
		PPS:           PPS,
		Availablility: 1,
	}, "https://hello.world")
	offers, err := user.GetAvailableOffers(0)
	require.Len(t, offers, 1)
	require.NoError(t, err)

	//book
	err = user.BookVM(0)
	require.NoError(t, err)

	//check availability changed
	offers, err = user.GetAvailableOffers(0)
	require.NoError(t, err)
	require.Len(t, offers, 0)

	//check locked amount changes
	free, locked, err = user.GetStablecoinBalance()
	require.NoError(t, err)
	require.Equal(t, 60*60*24*7*PPS, locked)
	require.Equal(t, INITATL_BALANCE-locked, free)
}

//test user can't book over his available money
func TestNotEnoughMoney(t *testing.T) {
	var err error

	testEnv := CreateTestEnv(t, 2)
	user := testEnv.Users[0]
	miner := testEnv.Users[1]

	const PPS = uint64(2)
	const BALANCE_REQUIRED_FOR_1_VM = uint64(60 * 60 * 24 * 7 * PPS)

	//add offer
	_ = miner.AddOffer(protocol.Offer{
		VmTypeId:      0,
		PPS:           PPS,
		Availablility: 3,
	}, "https://hello.world")

	//add some money
	testEnv.AddSomeStablecoin(*user.GetMyAddress(), BALANCE_REQUIRED_FOR_1_VM*10)
	err = user.DepositStablecoin(BALANCE_REQUIRED_FOR_1_VM - 1)
	require.NoError(t, err)

	free, _, _ := user.GetStablecoinBalance()
	require.Equal(t, BALANCE_REQUIRED_FOR_1_VM-1, free)

	//book 1st VM
	err = user.BookVM(0)
	require.Error(t, err)

	//check balance the same
	free, _, _ = user.GetStablecoinBalance()
	require.Equal(t, BALANCE_REQUIRED_FOR_1_VM-1, free)

	//deposit enough
	err = user.DepositStablecoin(BALANCE_REQUIRED_FOR_1_VM)
	require.NoError(t, err)

	//book 1st vm
	err = user.BookVM(0)
	require.NoError(t, err)

	//not enough for 2nd VM
	err = user.BookVM(0)
	require.Error(t, err)

	//add 1 coin + 1 leftover
	err = user.DepositStablecoin(2)
	require.NoError(t, err)

	//book 2nd vm
	err = user.BookVM(0)
	require.NoError(t, err)

	//check balance is 0
	free, locked, err := user.GetStablecoinBalance()
	require.NoError(t, err)
	require.Equal(t, BALANCE_REQUIRED_FOR_1_VM*2, locked)
	require.EqualValues(t, 1, free)
}

//check can't book over availability
func TestNotEnoughVms(t *testing.T) {
	var err error

	testEnv := CreateTestEnv(t, 2)
	user := testEnv.Users[0]
	miner := testEnv.Users[1]

	const PPS = 2
	const BALANCE_REQUIRED_FOR_1_VM = 60 * 60 * 24 * 7 * PPS

	//add offer
	_ = miner.AddOffer(protocol.Offer{
		VmTypeId:      0,
		PPS:           PPS,
		Availablility: 2,
	}, "https://hello.world")

	//add some money
	testEnv.AddSomeStablecoin(*user.GetMyAddress(), BALANCE_REQUIRED_FOR_1_VM*10)
	err = user.DepositStablecoin(BALANCE_REQUIRED_FOR_1_VM * 10)
	require.NoError(t, err)

	// 1 ok
	err = user.BookVM(0)
	require.NoError(t, err)

	// 2 ok
	err = user.BookVM(0)
	require.NoError(t, err)

	//3 ok
	err = user.BookVM(0)
	require.Error(t, err)
}
