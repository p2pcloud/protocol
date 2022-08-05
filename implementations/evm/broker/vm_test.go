package broker_test

import (
	"github.com/p2pcloud/protocol/implementations/evm"
	"testing"

	"github.com/p2pcloud/protocol"
	"github.com/stretchr/testify/require"
)

func TestBookingNotFound(t *testing.T) {
	blockchain := evm.NewWrappedSimulatedBlockchainEnv(t)

	communityPk, err := blockchain.GetNextPrivateKey()
	require.NoError(t, err)

	gifts := evm.NewGifts(
		map[int]float64{1: 1000},
		map[int]float64{1: 1000},
	)

	testInstances, err := evm.InitializeTestInstances(
		2, 6, gifts, blockchain.Origin.Backend, blockchain, communityPk,
	)
	check(t, err)

	minerContr := testInstances.Contracts[0]
	userContr := testInstances.Contracts[1]

	//test add offer
	err = minerContr.AddOffer(protocol.Offer{
		VmTypeId:      3,
		PPS:           1,
		Availablility: 1,
	}, "https://hello.world")
	require.NoError(t, err)

	require.NoError(t, userContr.DepositCoin(1000))

	offers, err := userContr.GetAvailableOffers(3)
	require.NoError(t, err)
	if len(offers) != 1 {
		t.Errorf("Expected 1 offer, got %d", len(offers))
	}

	//test book
	err = userContr.BookVM(offers[0].Index, 1000)
	require.NoError(t, err)

	_, err = userContr.GetBooking(99999)
	if err == nil {
		t.Fatal("Expected error for non existent booking")
	}
}

func TestGetBooking(t *testing.T) {
	blockchain := evm.NewWrappedSimulatedBlockchainEnv(t)

	communityPk, err := blockchain.GetNextPrivateKey()
	require.NoError(t, err)

	gifts := evm.NewGifts(
		map[int]float64{1: 1000},
		map[int]float64{1: 1000},
	)

	testInstances, err := evm.InitializeTestInstances(
		2, 6, gifts, blockchain.Origin.Backend, blockchain, communityPk,
	)
	check(t, err)

	minerContr := testInstances.Contracts[0]
	userContr := testInstances.Contracts[1]

	//test add offer
	err = minerContr.AddOffer(protocol.Offer{
		VmTypeId:      3,
		PPS:           1,
		Availablility: 1,
	}, "https://hello.world")
	require.NoError(t, err)

	offers, err := userContr.GetAvailableOffers(3)
	require.NoError(t, err)
	if len(offers) != 1 {
		t.Errorf("Expected 1 offer, got %d", len(offers))
	}

	require.NoError(t, userContr.DepositCoin(1000))

	//test book
	err = userContr.BookVM(offers[0].Index, 1000)
	require.NoError(t, err)

	booking, err := userContr.GetBooking(0)
	require.NoError(t, err)
	assertEqual(t, 0, booking.Index)
	assertEqual(t, 1.0, booking.PPS)
	assertEqual(t, 3, booking.VmTypeId)
}

func TestBook(t *testing.T) {
	blockchain := evm.NewWrappedSimulatedBlockchainEnv(t)

	communityPk, err := blockchain.GetNextPrivateKey()
	require.NoError(t, err)

	gifts := evm.NewGifts(
		map[int]float64{1: 17777},
		map[int]float64{1: 17777},
	)

	testInstances, err := evm.InitializeTestInstances(
		2, 6, gifts, blockchain.Origin.Backend, blockchain, communityPk,
	)
	check(t, err)

	minerContr := testInstances.Contracts[0]
	userContr := testInstances.Contracts[1]

	//test add offer
	err = minerContr.AddOffer(protocol.Offer{
		VmTypeId:      3,
		PPS:           1,
		Availablility: 1,
	}, "https://hello.world")
	require.NoError(t, err)

	offers, err := userContr.GetAvailableOffers(3)
	require.NoError(t, err)
	if len(offers) != 1 {
		t.Errorf("Expected 1 offer, got %d", len(offers))
	}

	userContr.DepositCoin(17777)

	//test book
	err = userContr.BookVM(offers[0].Index, 17777)
	require.NoError(t, err)

	bookings, err := userContr.GetUsersBookings()
	require.NoError(t, err)
	if len(bookings) != 1 {
		t.Errorf("Expected 1 booking, got %d", len(bookings))
	}

	//check time
	now, err := userContr.GetTime()
	require.NoError(t, err)

	require.Equal(t, 17777, bookings[0].BookedTill-now)

	bookings, err = minerContr.GetMinersBookings()
	require.NoError(t, err)
	if len(bookings) != 1 {
		t.Errorf("Expected 1 booking, got %d", len(bookings))
	}
}
