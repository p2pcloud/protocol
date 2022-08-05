package broker_test

import (
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"

	"github.com/p2pcloud/protocol"
	"github.com/p2pcloud/protocol/implementations/evm"
)

func TestAbortBookingWithMany(t *testing.T) {
	blockchain := evm.NewWrappedSimulatedBlockchainEnv(t)

	communityPk, err := blockchain.GetNextPrivateKey()
	require.NoError(t, err)

	userIdx := 0
	minerIdx := 1

	ti, err := evm.InitializeTestInstances(
		2, 6, evm.NewGifts(
			map[int]float64{userIdx: 3000}, map[int]float64{userIdx: 3000},
		),
		blockchain.Origin.Backend, blockchain,
		communityPk,
	)
	require.NoError(t, err)

	user := ti.Contracts[userIdx]

	miner := ti.Contracts[minerIdx]

	require.NoError(t, miner.AddOffer(protocol.Offer{
		VmTypeId:      1,
		PPS:           10,
		Availablility: 1,
	}, "https://hello.world"))
	require.NoError(t, miner.AddOffer(protocol.Offer{
		VmTypeId:      1,
		PPS:           20,
		Availablility: 1,
	}, "https://hello.world"))

	require.NoError(t, user.DepositCoin(3000))

	offers, err := user.GetAvailableOffers(1)
	require.NoError(t, err)
	require.Len(t, offers, 2)

	require.NoError(t, user.BookVM(offers[0].Index, 60))
	require.NoError(t, user.BookVM(offers[1].Index, 80))

	lockedBalance, err := user.LockedBalance()
	require.NoError(t, err)
	require.Equal(t, 2200.0, lockedBalance)

	userBalance, err := user.Balance()
	require.NoError(t, err)
	require.Equal(t, 800.0, userBalance)

	bookings, err := user.GetUsersBookings()
	require.NoError(t, err)
	booking := bookings[0]

	require.NoError(t, blockchain.Origin.Backend.AdjustTime(time.Second*5))
	blockchain.Origin.Backend.Commit()

	require.NoError(t, user.AbortBooking(uint64(booking.Index), protocol.ReportAbortType))

	userBalance, err = user.Balance()
	require.NoError(t, err)
	require.Equal(t, 1050.0, userBalance)

	lockedBalance, err = user.LockedBalance()
	require.NoError(t, err)
	require.Equal(t, 1600.0, lockedBalance)

	commBalance, err := ti.DeployerToken.BalanceOf(crypto.PubkeyToAddress(communityPk.PublicKey))
	require.NoError(t, err)

	minerBalance, err := miner.UserTokenBalance()
	require.NoError(t, err)

	require.Equal(t, 175.0, commBalance)
	require.Equal(t, 175.0, minerBalance)
}

func TestAbortBookingInvalidUser(t *testing.T) {
	blockchain := evm.NewWrappedSimulatedBlockchainEnv(t)

	communityPk, err := blockchain.GetNextPrivateKey()
	require.NoError(t, err)

	userIdx := 0
	minerIdx := 1
	otherUser := 2

	ti, err := evm.InitializeTestInstances(
		3, 6, evm.NewGifts(
			map[int]float64{userIdx: 3000}, map[int]float64{userIdx: 3000},
		),
		blockchain.Origin.Backend, blockchain,
		communityPk,
	)
	require.NoError(t, err)

	user := ti.Contracts[userIdx]

	miner := ti.Contracts[minerIdx]

	require.NoError(t, miner.AddOffer(protocol.Offer{
		VmTypeId:      1,
		PPS:           10,
		Availablility: 1,
	}, "https://hello.world"))

	require.NoError(t, user.DepositCoin(3000))

	offers, err := user.GetAvailableOffers(1)
	require.NoError(t, err)
	require.Len(t, offers, 1)

	require.NoError(t, user.BookVM(offers[0].Index, 60))

	bookings, err := user.GetUsersBookings()
	require.NoError(t, err)
	booking := bookings[0]

	require.NoError(t, blockchain.Origin.Backend.AdjustTime(time.Second*5))
	blockchain.Origin.Backend.Commit()

	err = ti.Contracts[otherUser].AbortBooking(uint64(booking.Index), protocol.ReportAbortType)
	require.ErrorContains(t, err, "only owner of booking can report")
}

func TestAbortBooking(t *testing.T) {
	var tests = []struct {
		name                  string
		deposit               float64
		offer                 protocol.Offer
		secs                  int64
		waitSecs              int64
		communityFee          int64
		abortType             protocol.AbortType
		wantUserBalance       float64
		wantUserLockedBalance float64
		wantMinersBalance     float64
		wantCommunityBalance  float64
		wantErrContains       string
	}{
		{
			name:    "user successfully reports",
			deposit: 600,
			offer: protocol.Offer{
				VmTypeId:      1,
				PPS:           10,
				Availablility: 1,
			},
			secs:                  50,
			waitSecs:              20,
			abortType:             protocol.ReportAbortType,
			wantUserBalance:       200,
			wantUserLockedBalance: 0,
			wantMinersBalance:     200,
			wantCommunityBalance:  200,
		},
		{
			name:    "user successfully stops",
			deposit: 600,
			offer: protocol.Offer{
				VmTypeId:      1,
				PPS:           10,
				Availablility: 1,
			},
			secs:                  50,
			waitSecs:              20,
			communityFee:          5,
			abortType:             protocol.StopAbortType,
			wantUserBalance:       200,
			wantUserLockedBalance: 0,
			wantMinersBalance:     380,
			wantCommunityBalance:  20,
		},
		{
			name:    "user reports, overflow 6 decimals",
			deposit: 500,
			offer: protocol.Offer{
				VmTypeId:      1,
				PPS:           12,
				Availablility: 1,
			},
			secs:                  50,
			waitSecs:              29,
			abortType:             protocol.ReportAbortType,
			wantUserBalance:       10.000049,
			wantUserLockedBalance: 0,
			wantMinersBalance:     244.999975,
			wantCommunityBalance:  244.999976,
		},
		{
			name:    "user stops, overflow 6 decimals",
			deposit: 500,
			offer: protocol.Offer{
				VmTypeId:      1,
				PPS:           12,
				Availablility: 1,
			},
			secs:                  50,
			waitSecs:              29,
			communityFee:          5,
			abortType:             protocol.StopAbortType,
			wantUserBalance:       10.000049,
			wantUserLockedBalance: 0,
			wantMinersBalance:     465.499953,
			wantCommunityBalance:  24.499998,
		},
		{
			name:    "user reports, lowest pps",
			deposit: 500,
			offer: protocol.Offer{
				VmTypeId:      1,
				PPS:           1,
				Availablility: 1,
			},
			secs:                  50,
			waitSecs:              29,
			abortType:             protocol.ReportAbortType,
			wantUserBalance:       499.999951,
			wantUserLockedBalance: 0,
			wantMinersBalance:     0.000024,
			wantCommunityBalance:  0.000025,
		},
		{
			name:    "user stops, lowest pps",
			deposit: 500,
			offer: protocol.Offer{
				VmTypeId:      1,
				PPS:           1,
				Availablility: 1,
			},
			secs:                  50,
			waitSecs:              29,
			abortType:             protocol.StopAbortType,
			wantUserBalance:       499.999951,
			wantUserLockedBalance: 0,
			wantMinersBalance:     0.000046,
			wantCommunityBalance:  0.000003,
		},
		{
			name:    "booking expired",
			deposit: 500,
			offer: protocol.Offer{
				VmTypeId:      1,
				PPS:           10,
				Availablility: 1,
			},
			secs:            50,
			waitSecs:        999,
			abortType:       protocol.StopAbortType,
			wantErrContains: "expired",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			blockchain := evm.NewWrappedSimulatedBlockchainEnv(t)

			communityPk, err := blockchain.GetNextPrivateKey()
			require.NoError(t, err)

			userIdx := 0
			minerIdx := 1

			ti, err := evm.InitializeTestInstances(
				2, 6, evm.NewGifts(
					map[int]float64{userIdx: tt.deposit}, map[int]float64{userIdx: tt.deposit},
				),
				blockchain.Origin.Backend, blockchain,
				communityPk,
			)
			require.NoError(t, err)

			user := ti.Contracts[userIdx]

			miner := ti.Contracts[minerIdx]

			require.NoError(t, miner.AddOffer(tt.offer, "https://hello.world"))

			if tt.deposit > 1 {
				require.NoError(t, user.DepositCoin(tt.deposit))
			}

			offers, err := user.GetAvailableOffers(1)
			require.NoError(t, err)
			require.Len(t, offers, 1)

			offer := offers[0]
			require.NoError(t, user.BookVM(offer.Index, int(tt.secs)))

			bookings, err := user.GetUsersBookings()
			require.NoError(t, err)
			booking := bookings[0]

			require.NoError(t, blockchain.Origin.Backend.AdjustTime(time.Second*time.Duration(tt.waitSecs)))
			blockchain.Origin.Backend.Commit()

			err = user.AbortBooking(uint64(booking.Index), tt.abortType)
			if err != nil {
				require.ErrorContains(t, err, tt.wantErrContains)
				require.NotEqual(t, "", tt.wantErrContains)
				return
			}

			userBalance, err := user.Balance()
			require.NoError(t, err)

			commBalance, err := ti.DeployerToken.BalanceOf(crypto.PubkeyToAddress(communityPk.PublicKey))
			require.NoError(t, err)
			t.Logf("community balance is %.6f", commBalance)

			minerBalance, err := miner.UserTokenBalance()
			require.NoError(t, err)
			require.Equal(t, tt.wantMinersBalance, minerBalance)
			require.Equal(t, tt.wantUserBalance, userBalance)
		})
	}
}
