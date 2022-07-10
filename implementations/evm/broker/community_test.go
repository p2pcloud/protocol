package broker_test

import (
	"strings"
	"testing"
)

func TestGetCommunityWallet(t *testing.T) {
	testInstances, simChain := getTestInstances(t, 2)
	defer simChain.Close()

	contr1 := testInstances[0]

	hash := strings.ToLower("0x9b920402017534cd42AC81e738E0F23a7162c4D7")

	err := contr1.SetCommunityWalletIfNeeded(hash)
	check(t, err)
	simChain.Commit()

	contr2 := testInstances[1]

	loaded, err := contr2.GetCommunityWallet()
	check(t, err)

	if loaded != hash {
		t.Fatalf("hashes are not equal: expected %s, got %s", hash, loaded)
	}
}

func TestGetCommunityFee(t *testing.T) {
	testInstances, simChain := getTestInstances(t, 2)
	defer simChain.Close()

	contr1 := testInstances[0]

	fee := int64(123)

	err := contr1.SetCommunityFeeIfNeeded(fee)
	check(t, err)
	simChain.Commit()

	contr2 := testInstances[1]

	loaded, err := contr2.GetCommunityFee()
	check(t, err)

	if loaded.Int64() != fee {
		t.Fatalf("hashes are not equal: expected %d, got %d", fee, loaded.Int64())
	}
}
