package keyring_test

import (
	"testing"

	"github.com/p2pcloud/protocol/pkg/keyring"
	"github.com/stretchr/testify/require"
)

func TestGenPrivateKey(t *testing.T) {
	testPrivateKey, err := keyring.GenPrivateKey()
	require.NoError(t, err)
	require.NotNil(t, testPrivateKey)
}
