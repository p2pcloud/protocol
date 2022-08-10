package evm_test

import (
	"testing"
	"time"

	"github.com/p2pcloud/protocol"
	"github.com/stretchr/testify/require"
)

func TestPayment(t *testing.T) {
	var err error

	const PPS = 1000
	const INITATL_BALANCE = 10000000000

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

	//book vm
	err = user.BookVM(0)
	require.NoError(t, err)
	bookings, err := user.GetUsersBookings()
	require.Len(t, bookings, 1)
	require.NoError(t, err)

	booking := bookings[0]

	//0 to withdraw
	err = miner.ClaimPayment(0)
	require.NoError(t, err)

	//check time
	ts, err := miner.GetTime()
	require.NoError(t, err)
	secondsPassed := ts - booking.BookedAt

	//5% is a default comission
	minerCoins, _, err := miner.GetStablecoinBalance()
	require.NoError(t, err)
	require.Equal(t, secondsPassed*PPS*95/100, minerCoins) //95%

	communityCoins, _, err := testEnv.Admin.GetStablecoinBalance()
	require.NoError(t, err)
	require.Equal(t, secondsPassed*PPS*5/100, communityCoins) //5%

	//skip time
	testEnv.AdjustTime(time.Minute * 10)

	//claim payment
	err = miner.ClaimPayment(0)
	require.NoError(t, err)

	//check time
	ts, err = miner.GetTime()
	require.NoError(t, err)

	//check last payment date updates
	booking, err = user.GetBooking(0)
	require.NoError(t, err)
	require.Equal(t, ts, booking.LastPayment)

	//check balance
	secondsPassed = ts - booking.BookedAt
	minerCoins, _, err = miner.GetStablecoinBalance()

	//miner gets 95% of payment
	require.NoError(t, err)
	require.Equal(t, secondsPassed*PPS*95/100, minerCoins) //95%

	//community gets 5% of payment
	communityCoins, _, err = testEnv.Admin.GetStablecoinBalance()
	require.NoError(t, err)
	require.Equal(t, secondsPassed*PPS*5/100, communityCoins) //5%
}

//TODO: check for events
//TODO: try to claim stopped booking
//TODO: test other values community fee
//TODO: try to break math in ClaimPayment method
//TODO: check for rounding errors
//TODO: check booking gets deleted if it is not enough money to cover all costs