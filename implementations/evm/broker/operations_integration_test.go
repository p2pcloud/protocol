package broker_test

import (
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"

	"github.com/p2pcloud/protocol/implementations/evm"
	"github.com/p2pcloud/protocol/implementations/evm/token"
)

func TestDepositCoinIntegration(t *testing.T) {
	rpcEndpoint := os.Getenv("GANACHE_RPC_ENDPOINT")
	if rpcEndpoint == "" {
		t.Skip()
	}

	web3Client, err := ethclient.Dial(rpcEndpoint)
	require.NoError(t, err)

	var tests = []struct {
		name                string
		deposit             float64
		wantUserRealBalance float64
		wantUserBalance     float64
		wantContractBalance float64
		wantErrContains     string
	}{
		{
			name:                "basic",
			deposit:             1.2,
			wantUserBalance:     1.2,
			wantUserRealBalance: 0.3,
			wantContractBalance: 1.2,
		},
		{
			name:                "6 decimals",
			deposit:             1.199999,
			wantUserBalance:     1.199999,
			wantUserRealBalance: 0.300001,
			wantContractBalance: 1.199999,
		},
		{
			name:                "overflow decimals, last 9",
			deposit:             1.1999999,
			wantUserBalance:     1.199999,
			wantUserRealBalance: 0.300001,
			wantContractBalance: 1.199999,
		},
		{
			name:                "overflow decimals, last 0",
			deposit:             1.1999990,
			wantUserBalance:     1.199999,
			wantUserRealBalance: 0.300001,
			wantContractBalance: 1.199999,
		},
		{
			name:            "insufficient balance",
			deposit:         2,
			wantErrContains: "exceeds balance",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			blockchain, err := evm.NewGanacheBCHelper(1, web3Client)
			require.NoError(t, err)

			userIdx := 0
			ti, err := evm.InitializeTestInstances(
				1, 6, evm.NewGifts(map[int]float64{0: 1.5}, nil),
				web3Client, blockchain,
			)
			require.NoError(t, err)

			contract := ti.Contracts[0]

			user, err := evm.NewEVMImplementation(
				blockchain.GetUserPrivateKeyByIndexStr(userIdx),
				contract.ContractAddress().Hex(),
				rpcEndpoint,
				evm.ChainIDSimulated,
			)
			require.NoError(t, err)

			err = user.DepositCoin(tt.deposit)
			if err != nil {
				require.ErrorContains(t, err, tt.wantErrContains)
				return
			}

			balance, err := user.Balance()
			require.NoError(t, err)
			require.Equal(t, tt.wantUserBalance, balance)

			realBalance, err := ti.DeployerToken.BalanceOf(
				crypto.PubkeyToAddress(user.GetPrivateKey().PublicKey),
			)
			require.NoError(t, err)
			require.Equal(t, tt.wantUserRealBalance, realBalance)

			contractBalance, err := ti.DeployerToken.BalanceOf(user.ContractAddress())
			require.NoError(t, err)
			require.Equal(t, tt.wantContractBalance, contractBalance)
		})
	}
}

func TestWithdrawCoinIntegration(t *testing.T) {
	userIdx := 0

	rpcEndpoint := os.Getenv("GANACHE_RPC_ENDPOINT")
	if rpcEndpoint == "" {
		t.Skip()
	}

	web3Client, err := ethclient.Dial(rpcEndpoint)
	require.NoError(t, err)

	var tests = []struct {
		name                string
		gifts               *evm.Gifts
		decimals            uint8
		deposit             float64
		withdraw            float64
		wantUserBalance     float64
		wantUserRealBalance float64
		wantContractBalance float64
		wantAllowance       float64
		wantErrContains     string
	}{
		{
			name: "success",
			gifts: evm.NewGifts(map[int]float64{
				userIdx: 1.5,
			}, map[int]float64{
				userIdx: 1.2,
			},
			),
			decimals:            6,
			deposit:             1.2,
			withdraw:            1.2,
			wantUserBalance:     0,
			wantUserRealBalance: 1.5,
			wantContractBalance: 0,
		},
		{
			name: "6 decimals",
			gifts: evm.NewGifts(map[int]float64{
				userIdx: 1.5,
			}, map[int]float64{
				userIdx: 1.2,
			},
			),
			decimals:            6,
			deposit:             1.2,
			withdraw:            1.199999,
			wantUserBalance:     0.000001,
			wantUserRealBalance: 1.499999,
			wantContractBalance: 0.000001,
		},
		{
			name: "overflow decimals, last 9",
			gifts: evm.NewGifts(map[int]float64{
				userIdx: 1.5,
			}, map[int]float64{
				userIdx: 1.2,
			},
			),
			decimals:            6,
			deposit:             1.2,
			withdraw:            1.1999999,
			wantUserBalance:     0.000001,
			wantUserRealBalance: 1.499999,
			wantContractBalance: 0.000001,
		},
		{
			name: "overflow decimals, last 0",
			gifts: evm.NewGifts(map[int]float64{
				userIdx: 1.5,
			}, map[int]float64{
				userIdx: 1.2,
			},
			),
			decimals:            6,
			deposit:             1.2,
			withdraw:            1.1999990,
			wantUserBalance:     0.000001,
			wantUserRealBalance: 1.499999,
			wantContractBalance: 0.000001,
		},
		{
			name: "overflow deposit + withdraw decimals, last 9",
			gifts: evm.NewGifts(map[int]float64{
				userIdx: 1.5,
			}, map[int]float64{
				userIdx: 1.2,
			},
			),
			decimals:            6,
			deposit:             1.1999999,
			withdraw:            0.3333339,
			wantUserBalance:     0.866666,
			wantUserRealBalance: 0.633334,
			wantContractBalance: 0.866666,
		},
		{
			name: "overflow deposit + withdraw decimals, last 0",
			gifts: evm.NewGifts(map[int]float64{
				userIdx: 1.5,
			}, map[int]float64{
				userIdx: 1.2,
			},
			),
			decimals:            6,
			deposit:             1.1999990,
			withdraw:            0.3333330,
			wantUserBalance:     0.866666,
			wantUserRealBalance: 0.633334,
			wantContractBalance: 0.866666,
		},
		{
			name: "insufficient funds",
			gifts: evm.NewGifts(map[int]float64{
				userIdx: 1.5,
			}, map[int]float64{
				userIdx: 1.5,
			},
			),
			decimals:        6,
			deposit:         1.5,
			withdraw:        2,
			wantErrContains: "insufficient funds to withdraw",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			blockchain, err := evm.NewGanacheBCHelper(1, web3Client)
			require.NoError(t, err)

			ti, err := evm.InitializeTestInstances(
				1, 6, evm.NewGifts(map[int]float64{0: 1.5}, nil),
				web3Client, blockchain,
			)
			require.NoError(t, err)

			contract := ti.Contracts[0]

			user, err := evm.NewEVMImplementation(
				blockchain.GetUserPrivateKeyByIndexStr(userIdx),
				contract.ContractAddress().Hex(),
				rpcEndpoint,
				evm.ChainIDSimulated,
			)
			require.NoError(t, err)

			require.NoError(t, user.DepositCoin(tt.deposit))

			err = user.WithdrawCoin(tt.withdraw)
			if err != nil {
				require.ErrorContains(t, err, tt.wantErrContains)
				return
			}

			balance, err := user.Balance()
			require.NoError(t, err)
			require.Equal(t, tt.wantUserBalance, balance)

			realBalance, err := ti.DeployerToken.BalanceOf(
				crypto.PubkeyToAddress(user.GetPrivateKey().PublicKey),
			)
			require.NoError(t, err)
			require.Equal(t, tt.wantUserRealBalance, realBalance)

			contractBalance, err := ti.DeployerToken.BalanceOf(user.ContractAddress())
			require.NoError(t, err)
			require.Equal(t, tt.wantContractBalance, contractBalance)
		})
	}
}

func TestGetStablecoinAddressIntegration(t *testing.T) {
	communityIdx := 0

	rpcEndpoint := os.Getenv("GANACHE_RPC_ENDPOINT")
	if rpcEndpoint == "" {
		t.Skip()
	}

	web3Client, err := ethclient.Dial(rpcEndpoint)
	require.NoError(t, err)

	blockchain, err := evm.NewGanacheBCHelper(1, web3Client)
	require.NoError(t, err)

	ti, err := evm.InitializeTestInstances(
		1, 6, nil,
		web3Client, blockchain,
	)
	require.NoError(t, err)

	comm := ti.Contracts[communityIdx]

	newTokenPk, err := blockchain.GetNextPrivateKey()
	require.NoError(t, err)

	decimals := uint8(8)
	tkn := token.NewToken(&token.Params{
		Decimals:   &decimals,
		Backend:    web3Client,
		PrivateKey: newTokenPk,
		ChainID:    evm.ChainIDSimulated,
		WaitForTx:  blockchain.WaitForTx,
	})
	newTokenAddr, err := tkn.(*token.Token).DeployContract(0)
	require.NoError(t, err)

	require.NoError(t, comm.SetStablecoinAddress(*newTokenAddr))

	got, err := comm.GetStablecoinAddress()
	require.NoError(t, err)
	require.Equal(t, *newTokenAddr, got)
}
