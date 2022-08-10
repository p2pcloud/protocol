package evm_test

import (
	"testing"
	"time"

	"github.com/p2pcloud/protocol"
	"github.com/stretchr/testify/require"
)

func TestStop(t *testing.T) {
	//TODO: skip time
	//TODO: claim payment
	//TODO: skip more time
	//TODO: stop
	//TODO: check miner balance
	//TODO: failure to claim payment

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
	booking, _ := user.GetBooking(0)

	// skip time
	testEnv.AdjustTime(time.Minute * 10)

	//claim payment
	err = miner.ClaimPayment(0)
	require.NoError(t, err)

	//skip more time
	testEnv.AdjustTime(time.Minute * 30)

	//stop
	err = user.StopVM(0, protocol.StopReasonNotNeeded)
	require.NoError(t, err)

	//check time
	ts, err := miner.GetTime()
	require.NoError(t, err)
	secondsPassed := ts - booking.BookedAt

	//check miner was paid exactly the amount of seconds passed
	minerCoins, _, err := miner.GetStablecoinBalance()
	require.NoError(t, err)
	require.Equal(t, secondsPassed*PPS*95/100, minerCoins) //95%

	//community also gets 5%
	communityCoins, _, err := testEnv.Admin.GetStablecoinBalance()
	require.NoError(t, err)
	require.Equal(t, secondsPassed*PPS*5/100, communityCoins) //5%

	//check for the absense of Complaint event
	events, err := user.GetComplaintEvents(protocol.ComplaintFilterOpts{})
	require.NoError(t, err)
	require.Len(t, events, 0)

	//claim payment causes error
	err = miner.ClaimPayment(0)
	require.Error(t, err)

	//stop booking causes error
	err = user.StopVM(0, protocol.StopReasonNotNeeded)
	require.Error(t, err)
}

//TODO: check availability comes back
func TestStopComplaint(t *testing.T) {
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
	booking, _ := user.GetBooking(0)

	// skip time
	testEnv.AdjustTime(time.Minute * 10)

	//stop
	err = user.StopVM(0, protocol.StopReasonComplaint)
	require.NoError(t, err)

	//check time
	ts, err := miner.GetTime()
	require.NoError(t, err)
	secondsPassed := ts - booking.BookedAt

	expectedPayment := secondsPassed * PPS * 95 / 100

	//check Complain event
	events, err := user.GetComplaintEvents(protocol.ComplaintFilterOpts{})
	require.NoError(t, err)
	require.Len(t, events, 1)
	require.Equal(t, protocol.StopReasonComplaint, events[0].Reason)
	require.Equal(t, miner.GetMyAddress().Hex(), events[0].Miner.Hex())
	require.Equal(t, user.GetMyAddress().Hex(), events[0].User.Hex())

	//check payment event
	payEvents, err := user.GetPaymentEvents(protocol.PaymentFilterOpts{})
	require.NoError(t, err)
	require.Len(t, payEvents, 1)
	require.Equal(t, miner.GetMyAddress().Hex(), payEvents[0].Miner.Hex())
	require.Equal(t, user.GetMyAddress().Hex(), payEvents[0].User.Hex())
	require.Equal(t, expectedPayment, payEvents[0].Amount)

}
