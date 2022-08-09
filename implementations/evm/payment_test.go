package evm_test

import (
	"log"
	"testing"
	"time"

	"github.com/p2pcloud/protocol"
	"github.com/stretchr/testify/require"
)

/*
book, skip time, stop
check money withdrawn from client to miner
check percentage is sent to community
*/

/*
stop when not enough money to pay
*/

/*

 */

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
	secondsPassed = ts - booking.BookedAt

	minerCoins, _, err = miner.GetStablecoinBalance()
	log.Println("miner coins", minerCoins)
	log.Println("secondsPassed", secondsPassed)
	log.Println("secondsPassed*PPS*95/100", secondsPassed*PPS*95/100)

	require.NoError(t, err)
	require.Equal(t, secondsPassed*PPS*95/100, minerCoins) //95%

	communityCoins, _, err = testEnv.Admin.GetStablecoinBalance()
	require.NoError(t, err)
	require.Equal(t, secondsPassed*PPS*5/100, communityCoins) //5%
}

//TODO: try to claim stopped booking
