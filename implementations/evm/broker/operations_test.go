package broker_test

import (
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"

	"github.com/p2pcloud/protocol/implementations/evm"
	"github.com/p2pcloud/protocol/implementations/evm/token"
)

func TestDepositCoin(t *testing.T) {
	userIdx := 0

	var tests = []struct {
		name                string
		gifts               *evm.Gifts
		decimals            uint8
		deposit             float64
		wantUserRealBalance float64
		wantUserBalance     float64
		wantContractBalance float64
		wantAllowance       float64
		wantErrContains     string
	}{
		{
			name: "basic",
			gifts: evm.NewGifts(map[int]float64{
				userIdx: 1.5,
			}, map[int]float64{
				userIdx: 1.2,
			},
			),
			decimals:            6,
			deposit:             1.2,
			wantUserBalance:     1.2,
			wantUserRealBalance: 0.3,
			wantContractBalance: 1.2,
			wantAllowance:       0,
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
			deposit:             1.199999,
			wantUserBalance:     1.199999,
			wantUserRealBalance: 0.300001,
			wantContractBalance: 1.199999,
			wantAllowance:       0.000001,
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
			deposit:             1.1999999,
			wantUserBalance:     1.199999,
			wantUserRealBalance: 0.300001,
			wantContractBalance: 1.199999,
			wantAllowance:       0.000001,
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
			deposit:             1.1999990,
			wantUserBalance:     1.199999,
			wantUserRealBalance: 0.300001,
			wantContractBalance: 1.199999,
			wantAllowance:       0.000001,
		},
		{
			name: "insufficient allowance",
			gifts: evm.NewGifts(map[int]float64{
				userIdx: 1.5,
			}, map[int]float64{
				userIdx: 1.5,
			},
			),
			decimals:        6,
			deposit:         2,
			wantErrContains: "insufficient allowance",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			blockchain := evm.NewWrappedSimulatedBlockchainEnv(t)

			communityPk, err := blockchain.GetNextPrivateKey()
			require.NoError(t, err)

			testInstances, err := evm.InitializeTestInstances(
				1, tt.decimals, tt.gifts,
				blockchain.Origin.Backend, blockchain, communityPk,
			)
			require.NoError(t, err)

			user := testInstances.Contracts[0]

			err = user.DepositCoin(tt.deposit)
			if err != nil {
				require.ErrorContains(t, err, tt.wantErrContains)
				return
			}

			balance, err := user.Balance()
			require.NoError(t, err)
			require.Equal(t, tt.wantUserBalance, balance)

			realBalance, err := testInstances.DeployerToken.BalanceOf(
				crypto.PubkeyToAddress(user.GetPrivateKey().PublicKey),
			)
			require.NoError(t, err)
			require.Equal(t, tt.wantUserRealBalance, realBalance)

			contractBalance, err := testInstances.DeployerToken.BalanceOf(user.ContractAddress())
			require.NoError(t, err)
			require.Equal(t, tt.wantContractBalance, contractBalance)

			allowance, err := testInstances.DeployerToken.Allowance(
				crypto.PubkeyToAddress(user.GetPrivateKey().PublicKey),
				user.ContractAddress(),
			)
			require.NoError(t, err)
			require.Equal(t, tt.wantAllowance, allowance)
		})
	}
}

func TestWithdrawCoin(t *testing.T) {
	userIdx := 0

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
			blockchain := evm.NewWrappedSimulatedBlockchainEnv(t)

			communityPk, err := blockchain.GetNextPrivateKey()
			require.NoError(t, err)

			testInstances, err := evm.InitializeTestInstances(
				1, tt.decimals, tt.gifts,
				blockchain.Origin.Backend, blockchain, communityPk,
			)
			require.NoError(t, err)

			user := testInstances.Contracts[0]

			require.NoError(t, user.DepositCoin(tt.deposit))

			err = user.WithdrawCoin(tt.withdraw)
			if err != nil {
				require.ErrorContains(t, err, tt.wantErrContains)
				return
			}

			balance, err := user.Balance()
			require.NoError(t, err)
			require.Equal(t, tt.wantUserBalance, balance)

			realBalance, err := testInstances.DeployerToken.BalanceOf(
				crypto.PubkeyToAddress(user.GetPrivateKey().PublicKey),
			)
			require.NoError(t, err)
			require.Equal(t, tt.wantUserRealBalance, realBalance)

			contractBalance, err := testInstances.DeployerToken.BalanceOf(user.ContractAddress())
			require.NoError(t, err)
			require.Equal(t, tt.wantContractBalance, contractBalance)
		})
	}
}

func TestGetStablecoinAddress(t *testing.T) {
	blockchain := evm.NewWrappedSimulatedBlockchainEnv(t)

	communityPk, err := blockchain.GetNextPrivateKey()
	require.NoError(t, err)

	ti, err := evm.InitializeTestInstances(
		2, 6, nil,
		blockchain.Origin.Backend, blockchain, communityPk,
	)
	require.NoError(t, err)

	newTokenPk, err := blockchain.GetNextPrivateKey()
	require.NoError(t, err)

	decimals := uint8(8)
	tkn := token.NewToken(&token.Params{
		Decimals:   &decimals,
		Backend:    blockchain.Origin.Backend,
		PrivateKey: newTokenPk,
		ChainID:    evm.ChainIDSimulated,
		WaitForTx:  blockchain.WaitForTx,
	})
	newTokenAddr, err := tkn.(*token.Token).DeployContract(0)
	require.NoError(t, err)

	require.NoError(t, ti.CommunityAccount.SetStablecoinAddress(*newTokenAddr))

	got, err := ti.CommunityAccount.GetStablecoinAddress()
	require.NoError(t, err)
	require.Equal(t, *newTokenAddr, got)
}
