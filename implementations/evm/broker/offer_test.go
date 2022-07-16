package broker_test

import (
	"testing"

	"github.com/p2pcloud/protocol"
	"github.com/stretchr/testify/require"
)

func check(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}

func TestAddOffer(t *testing.T) {
	testInstances, simChain := getTestInstances(t, 0, 2)
	defer simChain.Close()

	contract1 := testInstances[0]
	contract2 := testInstances[1]

	//test add offer
	err := contract1.AddOffer(protocol.Offer{
		VmTypeId:      3,
		PPS:           1,
		Availablility: 1,
	}, "https://hello.world")
	require.NoError(t, err)
	err = contract1.AddOffer(protocol.Offer{
		VmTypeId:      3,
		PPS:           1,
		Availablility: 1,
	}, "https://hello.world")
	require.NoError(t, err)

	simChain.Commit()

	offers, err := contract1.GetMyOffers()
	require.NoError(t, err)
	if len(offers) != 2 {
		t.Fatalf("expected 1 offer, got %d", len(offers))
	}
	if offers[0].VmTypeId != 3 {
		t.Fatalf("expected 3, got %d", offers[0].VmTypeId)
	}
	if offers[1].Index != 1 {
		t.Fatalf("expected 1, got %d", offers[1].Index)
	}

	offers, err = contract2.GetMyOffers()
	require.NoError(t, err)
	if len(offers) != 0 {
		t.Fatalf("expected 0 offer, got %d", len(offers))
	}
}

func assertEqual(t *testing.T, a, b interface{}) {
	if a != b {
		t.Fatalf("expected %v, got %v", a, b)
	}
}

func TestUrlUpdate(t *testing.T) {
	testInstances, simChain := getTestInstances(t, 0, 2)
	defer simChain.Close()

	contract := testInstances[0]

	//test add offer
	err := contract.AddOffer(protocol.Offer{
		VmTypeId:      3,
		PPS:           1,
		Availablility: 1,
	}, "https://hello.1")
	require.NoError(t, err)
	simChain.Commit()

	url, err := contract.GetMinerUrl(contract.GetMyAddress())
	require.NoError(t, err)

	assertEqual(t, "https://hello.1", url)

	//update manually
	err = contract.SetMinerUrlIfNeeded("https://hello.2")
	require.NoError(t, err)
	simChain.Commit()

	url, err = contract.GetMinerUrl(contract.GetMyAddress())
	require.NoError(t, err)

	assertEqual(t, "https://hello.2", url)
}

// func TestRemoveOffer(t *testing.T) {
// 	testInstances, simChain := getTestInstances(t, 2)
// 	defer simChain.Close()

// 	contract1 := testInstances[0]

// 	//test add offer
// 	err := contract1.AddOffer(protocol.Offer{
// 		VmTypeId:      3,
// 		PPS:           1,
// 		Availablility: 1,
// 	}, "http://hello.world")
// 	require.NoError(t, err)
// 	err = contract1.AddOffer(protocol.Offer{
// 		VmTypeId:      100,
// 		PPS:           1,
// 		Availablility: 1,
// 	}, "http://hello.world")
// 	require.NoError(t, err)
// 	simChain.Commit()

// 	offers, err := contract1.GetMyOffers()
// 	require.NoError(t, err)
// 	assertEqual(t, len(offers), 2)

// 	err = contract1.RemoveOffer(1)
// 	require.NoError(t, err)
// 	simChain.Commit()

// 	offers, err = contract1.GetMyOffers()
// 	require.NoError(t, err)
// 	assertEqual(t, len(offers), 1)
// 	assertEqual(t, offers[0].VmTypeId, 3)
// }

func TestUpdateOffer(t *testing.T) {
	testInstances, simChain := getTestInstances(t, 0, 2)
	defer simChain.Close()

	contract1 := testInstances[0]

	//test add offer
	err := contract1.AddOffer(protocol.Offer{
		VmTypeId:      3,
		PPS:           1,
		Availablility: 1,
	}, "http://hello.world")
	require.NoError(t, err)
	simChain.Commit()

	offers, err := contract1.GetMyOffers()
	require.NoError(t, err)
	require.Len(t, offers, 1)

	offerUpdate := offers[0]
	offerUpdate.PPS = 77
	offerUpdate.Availablility = 88
	offerUpdate.VmTypeId = 99

	err = contract1.UpdateOffer(offerUpdate)
	require.NoError(t, err)
	simChain.Commit()

	updatedOffers, err := contract1.GetMyOffers()
	require.NoError(t, err)
	require.Len(t, offers, 1)

	assertEqual(t, updatedOffers[0].VmTypeId, 99)
	assertEqual(t, updatedOffers[0].Availablility, 88)
	assertEqual(t, updatedOffers[0].PPS, 77)
}
