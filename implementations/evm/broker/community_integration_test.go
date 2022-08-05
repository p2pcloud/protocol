package broker_test

import (
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/p2pcloud/protocol/implementations/evm"

	"github.com/stretchr/testify/require"
)

func TestSetCommunityContract(t *testing.T) {
	rpcEndpoint := os.Getenv("GANACHE_RPC_ENDPOINT")
	if rpcEndpoint == "" {
		t.Skip()
	}

	web3Client, err := ethclient.Dial(rpcEndpoint)
	require.NoError(t, err)

	bcHelper, err := evm.NewGanacheBCHelper(1, web3Client)
	require.NoError(t, err)

	initialPk, err := bcHelper.GetNextPrivateKey()
	require.NoError(t, err)

	p := &evm.TestInstances{
		Count:              1,
		Decimals:           6,
		Backend:            web3Client,
		BcHelper:           bcHelper,
		UpdateCh:           make(chan common.Address, 1),
		CommunityInitialPk: initialPk,
	}
	require.NoError(t, evm.BuildToken(p))
	require.NoError(t, evm.BuildBroker(p))
	require.NoError(t, evm.SetFeeAndStablecoin(p))

	newCommunityPk, err := bcHelper.GetNextPrivateKey()
	require.NoError(t, err)

	community, err := evm.NewEVMImplementation(
		bcHelper.GetPrivateKeyString(initialPk),
		p.DeployerBroker.ContractAddress().Hex(),
		rpcEndpoint,
		evm.ChainIDSimulated,
	)
	require.NoError(t, err)

	require.NoError(t, community.SetCommunityContract(crypto.PubkeyToAddress(newCommunityPk.PublicKey)))

	wrongPk, err := bcHelper.GetNextPrivateKey()
	require.NoError(t, err)

	wrongContract, err := evm.NewEVMImplementation(
		bcHelper.GetPrivateKeyString(wrongPk),
		p.DeployerBroker.ContractAddress().Hex(),
		rpcEndpoint,
		evm.ChainIDSimulated,
	)
	require.NoError(t, err)

	require.Error(t, wrongContract.SetCommunityContract(crypto.PubkeyToAddress(wrongPk.PublicKey)))

	got, err := community.GetCommunityContract()
	require.NoError(t, err)
	require.Equal(t, crypto.PubkeyToAddress(newCommunityPk.PublicKey), got)
}

func TestSetCommunityFee(t *testing.T) {
	rpcEndpoint := os.Getenv("GANACHE_RPC_ENDPOINT")
	if rpcEndpoint == "" {
		t.Skip()
	}

	web3Client, err := ethclient.Dial(rpcEndpoint)
	require.NoError(t, err)

	bcHelper, err := evm.NewGanacheBCHelper(1, web3Client)
	require.NoError(t, err)

	initialPk, err := bcHelper.GetNextPrivateKey()
	require.NoError(t, err)

	p := &evm.TestInstances{
		Count:              1,
		Decimals:           6,
		Backend:            web3Client,
		BcHelper:           bcHelper,
		UpdateCh:           make(chan common.Address, 1),
		CommunityInitialPk: initialPk,
	}
	require.NoError(t, evm.BuildToken(p))
	require.NoError(t, evm.BuildBroker(p))
	require.NoError(t, evm.SetFeeAndStablecoin(p))

	community, err := evm.NewEVMImplementation(
		bcHelper.GetPrivateKeyString(initialPk),
		p.DeployerBroker.ContractAddress().Hex(),
		rpcEndpoint,
		evm.ChainIDSimulated,
	)
	require.NoError(t, err)

	require.NoError(t, community.SetCommunityFee(90))

	wrongPk, err := bcHelper.GetNextPrivateKey()
	require.NoError(t, err)

	wrongContract, err := evm.NewEVMImplementation(
		bcHelper.GetPrivateKeyString(wrongPk),
		p.DeployerBroker.ContractAddress().Hex(),
		rpcEndpoint,
		evm.ChainIDSimulated,
	)
	require.NoError(t, err)

	require.Error(t, wrongContract.SetCommunityFee(30))

	fee, err := community.GetCommunityFee()
	require.NoError(t, err)
	require.Equal(t, int64(90), fee)
}
