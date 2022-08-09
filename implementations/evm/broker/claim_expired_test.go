package broker_test

// func TestClaimExpiredInvalidUser(t *testing.T) {
// 	blockchain := evm.NewWrappedSimulatedBlockchainEnv(t)

// 	communityPk, err := blockchain.GetNextPrivateKey()
// 	require.NoError(t, err)

// 	userIdx := 0
// 	minerIdx := 1
// 	otherUser := 2

// 	ti, err := evm.InitializeTestInstances(
// 		3, 6, evm.NewGifts(
// 			map[int]float64{userIdx: 3000}, map[int]float64{userIdx: 3000},
// 		),
// 		blockchain.Origin.Backend, blockchain,
// 		communityPk,
// 	)
// 	require.NoError(t, err)

// 	user := ti.Contracts[userIdx]

// 	miner := ti.Contracts[minerIdx]

// 	require.NoError(t, miner.AddOffer(protocol.Offer{
// 		VmTypeId:      1,
// 		PPS:           10,
// 		Availablility: 1,
// 	}, "https://hello.world"))

// 	require.NoError(t, user.DepositCoin(3000))

// 	offers, err := user.GetAvailableOffers(1)
// 	require.NoError(t, err)
// 	require.Len(t, offers, 1)

// 	require.NoError(t, user.BookVM(offers[0].Index, 60))

// 	bookings, err := user.GetUsersBookings()
// 	require.NoError(t, err)
// 	booking := bookings[0]

// 	require.NoError(t, blockchain.Origin.Backend.AdjustTime(time.Second*999))
// 	blockchain.Origin.Backend.Commit()

// 	err = ti.Contracts[otherUser].ClaimExpired(uint64(booking.Index))
// 	require.ErrorContains(t, err, "only miner of booking can claim")
// }

// func TestClaimExpired(t *testing.T) {
// 	var tests = []struct {
// 		name                  string
// 		deposit               float64
// 		offer                 protocol.Offer
// 		secs                  int64
// 		waitSecs              int64
// 		wantUserBalance       float64
// 		wantUserLockedBalance float64
// 		wantMinersBalance     float64
// 		wantCommunityBalance  float64
// 		wantErrContains       string
// 	}{
// 		{
// 			name:    "success",
// 			deposit: 600,
// 			offer: protocol.Offer{
// 				VmTypeId:      1,
// 				PPS:           10,
// 				Availablility: 1,
// 			},
// 			secs:                  50,
// 			waitSecs:              999,
// 			wantUserBalance:       100,
// 			wantUserLockedBalance: 0,
// 			wantMinersBalance:     475,
// 			wantCommunityBalance:  25,
// 		},
// 		{
// 			name:    "overflow 6 decimals",
// 			deposit: 500,
// 			offer: protocol.Offer{
// 				VmTypeId:      1,
// 				PPS:           800,
// 				Availablility: 1,
// 			},
// 			secs:                  50,
// 			waitSecs:              999,
// 			wantUserBalance:       0.00005,
// 			wantUserLockedBalance: 0,
// 			wantMinersBalance:     474.999952,
// 			wantCommunityBalance:  24.999998,
// 		},
// 		{
// 			name:    "lowest pps",
// 			deposit: 500,
// 			offer: protocol.Offer{
// 				VmTypeId:      1,
// 				PPS:           1,
// 				Availablility: 1,
// 			},
// 			secs:                  50,
// 			waitSecs:              999,
// 			wantUserBalance:       499.99995,
// 			wantUserLockedBalance: 0,
// 			wantMinersBalance:     0.000047,
// 			wantCommunityBalance:  0.000003,
// 		},
// 		{
// 			name:    "currently used error",
// 			deposit: 500,
// 			offer: protocol.Offer{
// 				VmTypeId:      1,
// 				PPS:           10,
// 				Availablility: 1,
// 			},
// 			secs:            50,
// 			waitSecs:        1,
// 			wantErrContains: "booking is still in use",
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			blockchain := evm.NewWrappedSimulatedBlockchainEnv(t)

// 			communityPk, err := blockchain.GetNextPrivateKey()
// 			require.NoError(t, err)

// 			userIdx := 0
// 			minerIdx := 1

// 			ti, err := evm.InitializeTestInstances(
// 				2, 6, evm.NewGifts(
// 					map[int]float64{userIdx: tt.deposit}, map[int]float64{userIdx: tt.deposit},
// 				),
// 				blockchain.Origin.Backend, blockchain,
// 				communityPk,
// 			)
// 			require.NoError(t, err)

// 			user := ti.Contracts[userIdx]

// 			miner := ti.Contracts[minerIdx]

// 			require.NoError(t, miner.AddOffer(tt.offer, "https://hello.world"))

// 			if tt.deposit > 1 {
// 				require.NoError(t, user.DepositCoin(tt.deposit))
// 			}

// 			offers, err := user.GetAvailableOffers(1)
// 			require.NoError(t, err)
// 			require.Len(t, offers, 1)

// 			offer := offers[0]
// 			require.NoError(t, user.BookVM(offer.Index, int(tt.secs)))

// 			bookings, err := user.GetUsersBookings()
// 			require.NoError(t, err)
// 			booking := bookings[0]

// 			require.NoError(t, blockchain.Origin.Backend.AdjustTime(time.Second*time.Duration(tt.waitSecs)))
// 			blockchain.Origin.Backend.Commit()

// 			err = miner.ClaimExpired(uint64(booking.Index))
// 			if err != nil {
// 				require.ErrorContains(t, err, tt.wantErrContains)
// 				require.NotEqual(t, "", tt.wantErrContains)
// 				return
// 			}

// 			userBalance, err := user.Balance()
// 			require.NoError(t, err)

// 			commBalance, err := ti.DeployerToken.BalanceOf(crypto.PubkeyToAddress(communityPk.PublicKey))
// 			require.NoError(t, err)
// 			t.Logf("community balance is %.6f", commBalance)

// 			minerBalance, err := miner.UserTokenBalance()
// 			require.NoError(t, err)
// 			require.Equal(t, tt.wantMinersBalance, minerBalance)
// 			require.Equal(t, tt.wantUserBalance, userBalance)
// 		})
// 	}
// }
