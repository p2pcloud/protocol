package broker_test

import (
	"errors"

	"github.com/p2pcloud/protocol"
)

type vmWithSecs struct {
	offer protocol.Offer
	secs  int
}

// func TestBookVMIntegration(t *testing.T) {
// 	rpcEndpoint := os.Getenv("GANACHE_RPC_ENDPOINT")
// 	if rpcEndpoint == "" {
// 		t.Skip()
// 	}

// 	web3Client, err := ethclient.Dial(rpcEndpoint)
// 	require.NoError(t, err)

// 	defaultVms := []vmWithSecs{
// 		{
// 			offer: protocol.Offer{
// 				VmTypeId:      1,
// 				PPS:           10,
// 				Availablility: 1,
// 			},
// 			secs: 2,
// 		},
// 		{
// 			offer: protocol.Offer{
// 				VmTypeId:      1,
// 				PPS:           5,
// 				Availablility: 1,
// 			},
// 			secs: 6,
// 		},
// 	}

// 	var tests = []struct {
// 		name            string
// 		deposit         float64
// 		vms             []vmWithSecs
// 		wantUserBalance float64
// 		wantErrContains string
// 	}{
// 		{
// 			name:            "user successfully books one",
// 			deposit:         50,
// 			vms:             []vmWithSecs{defaultVms[0]},
// 			wantUserBalance: 30,
// 		},
// 		{
// 			name:            "user successfully books many",
// 			deposit:         60,
// 			vms:             defaultVms,
// 			wantUserBalance: 10,
// 		},
// 		{
// 			name:    "user tries to book for 0 seconds",
// 			deposit: 100,
// 			vms: []vmWithSecs{
// 				{
// 					offer: protocol.Offer{
// 						VmTypeId:      1,
// 						PPS:           5,
// 						Availablility: 1,
// 					},
// 					secs: 0,
// 				},
// 			},
// 			wantUserBalance: 100,
// 			wantErrContains: "secs should be positive",
// 		},
// 		{
// 			name:            "user books without depositing",
// 			deposit:         0,
// 			vms:             defaultVms,
// 			wantErrContains: "user does not have required amount to book machine",
// 		},
// 		{
// 			name:    "expensive machine",
// 			deposit: 50,
// 			vms: []vmWithSecs{
// 				{
// 					offer: protocol.Offer{
// 						VmTypeId:      1,
// 						PPS:           11,
// 						Availablility: 1,
// 					},
// 					secs: 5,
// 				},
// 			},
// 			wantUserBalance: 50,
// 			wantErrContains: "user does not have required amount to book machine",
// 		},
// 		{
// 			name:            "user ends up with locked balance",
// 			deposit:         40,
// 			vms:             defaultVms,
// 			wantUserBalance: 20,
// 			wantErrContains: "user does not have required amount to book machine",
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			blockchain, err := evm.NewGanacheBCHelper(2, web3Client)
// 			require.NoError(t, err)

// 			communityPk, err := blockchain.GetNextPrivateKey()
// 			require.NoError(t, err)

// 			userIdx := 0
// 			minerIdx := 1

// 			ti, err := evm.InitializeTestInstances(
// 				2, 6, evm.NewGifts(map[int]float64{userIdx: tt.deposit}, nil),
// 				web3Client, blockchain,
// 				communityPk,
// 			)
// 			require.NoError(t, err)

// 			userContract := ti.Contracts[userIdx]

// 			user, err := evm.NewEVMImplementation(
// 				blockchain.GetPrivateKeyString(userContract.GetPrivateKey()),
// 				userContract.ContractAddress().Hex(),
// 				rpcEndpoint,
// 				evm.ChainIDSimulated,
// 			)
// 			require.NoError(t, err)

// 			minerContract := ti.Contracts[minerIdx]

// 			miner, err := evm.NewEVMImplementation(
// 				blockchain.GetPrivateKeyString(minerContract.GetPrivateKey()),
// 				minerContract.ContractAddress().Hex(),
// 				rpcEndpoint,
// 				evm.ChainIDSimulated,
// 			)
// 			require.NoError(t, err)

// 			for i := range tt.vms {
// 				require.NoError(t, miner.AddOffer(tt.vms[i].offer, "https://hello.world"))
// 			}

// 			if tt.deposit > 1 {
// 				require.NoError(t, user.DepositCoin(tt.deposit))
// 			}

// 			offers, err := user.GetAvailableOffers(1)
// 			require.NoError(t, err)

// 			for i := range tt.vms {
// 				// secs, err := findOfferSecs(tt.vms[i], offers)
// 				require.NoError(t, err)

// 				err = user.BookVM(offers[i].Index)
// 				if err != nil {
// 					require.ErrorContains(t, err, tt.wantErrContains)
// 				}
// 			}

// 			userBalance, err := user.Balance()
// 			require.NoError(t, err)
// 			require.Equal(t, tt.wantUserBalance, userBalance)
// 		})
// 	}
// }

func findOfferSecs(vm vmWithSecs, offers []protocol.Offer) (int, error) {
	for _, o := range offers {
		offer := vm.offer
		if o.VmTypeId == offer.VmTypeId && o.PPS == offer.PPS && o.Availablility == offer.Availablility {
			return vm.secs, nil
		}
	}

	return 0, errors.New("not found")
}
