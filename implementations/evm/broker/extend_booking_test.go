package broker_test

// import (
// 	"testing"
// 	"time"

// 	"github.com/stretchr/testify/require"

// 	"github.com/p2pcloud/protocol"
// 	"github.com/p2pcloud/protocol/implementations/evm"
// )

// func TestExtendBooking(t *testing.T) {
// 	var tests = []struct {
// 		name                  string
// 		deposit               float64
// 		offer                 protocol.Offer
// 		bookSecs              int64
// 		waitSecs              int64
// 		extendSecs            int
// 		wantLifetime          int
// 		wantUserBalance       float64
// 		wantUserLockedBalance float64
// 		wantErrContains       string
// 	}{
// 		{
// 			name:    "success, last second",
// 			deposit: 1500,
// 			offer: protocol.Offer{
// 				VmTypeId:      1,
// 				PPS:           10,
// 				Availablility: 1,
// 			},
// 			bookSecs:              100,
// 			extendSecs:            50,
// 			waitSecs:              79,
// 			wantLifetime:          150,
// 			wantUserBalance:       0,
// 			wantUserLockedBalance: 1500,
// 		},
// 		{
// 			name:    "booking expired",
// 			deposit: 1500,
// 			offer: protocol.Offer{
// 				VmTypeId:      1,
// 				PPS:           10,
// 				Availablility: 1,
// 			},
// 			bookSecs:        100,
// 			extendSecs:      50,
// 			waitSecs:        80,
// 			wantErrContains: "booking is expired",
// 		},
// 		{
// 			name:    "insufficient balance",
// 			deposit: 1500,
// 			offer: protocol.Offer{
// 				VmTypeId:      1,
// 				PPS:           10,
// 				Availablility: 1,
// 			},
// 			bookSecs:        100,
// 			extendSecs:      51,
// 			waitSecs:        79,
// 			wantErrContains: "insufficient funds to extend booking",
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
// 			require.NoError(t, user.BookVM(offer.Index, int(tt.bookSecs)))

// 			bookings, err := user.GetUsersBookings()
// 			require.NoError(t, err)
// 			booking := bookings[0]

// 			require.NoError(t, blockchain.Origin.Backend.AdjustTime(time.Second*time.Duration(tt.waitSecs)))
// 			blockchain.Origin.Backend.Commit()

// 			err = user.ExtendBooking(uint64(booking.Index), tt.extendSecs)
// 			if err != nil {
// 				require.ErrorContains(t, err, tt.wantErrContains)
// 				require.NotEqual(t, "", tt.wantErrContains)
// 				return
// 			}

// 			bookings, err = user.GetUsersBookings()
// 			require.NoError(t, err)
// 			require.Equal(t, tt.wantLifetime, bookings[0].BookedTill-bookings[0].BookedAt)

// 			userBalance, err := user.Balance()
// 			require.NoError(t, err)
// 			require.Equal(t, tt.wantUserBalance, userBalance)

// 			lockedBalance, err := user.LockedBalance()
// 			require.NoError(t, err)
// 			require.Equal(t, tt.wantUserLockedBalance, lockedBalance)
// 		})
// 	}

// }
