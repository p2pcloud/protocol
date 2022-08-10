package broker

import (
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"

	"github.com/p2pcloud/protocol"
)

func (b *Broker) BookVM(offerIndex uint64) error {
	_, err := b.EstimateGas("BookVM", offerIndex)
	if err != nil {
		return err
	}

	tx, err := b.session.BookVM(offerIndex)
	if err != nil {
		return err
	}

	return b.waitForTx(tx)
}

func (b *Broker) StopVM(offerIndex uint64, reason uint8) error {
	_, err := b.EstimateGas("StopVM", offerIndex, reason)
	if err != nil {
		return err
	}

	tx, err := b.session.StopVM(offerIndex, reason)
	if err != nil {
		return err
	}

	return b.waitForTx(tx)
}

func (b *Broker) GetUsersBookings() ([]protocol.VMBooking, error) {
	// if err := b.setDecimals(); err != nil {
	// 	return nil, err
	// }

	bookings, err := b.session.FindBookingsByUser(b.transactOpts.From)
	if err != nil {
		return nil, err
	}
	var result []protocol.VMBooking
	for _, booking := range bookings {
		result = append(result, protocol.VMBooking{
			VmTypeId:    booking.VmTypeId,
			PPS:         booking.PricePerSecond,
			Miner:       &booking.Miner,
			Index:       booking.Index,
			User:        &booking.User,
			BookedAt:    booking.BookedAt.Uint64(),
			LastPayment: booking.LastPayment.Uint64(),
		})
	}
	return result, nil
}

func (b *Broker) GetBooking(index uint64) (protocol.VMBooking, error) {
	booking, err := b.session.GetBooking(index)
	if err != nil {
		return protocol.VMBooking{}, err
	}

	if booking.Miner.Hex() == "0x0000000000000000000000000000000000000000" {
		return protocol.VMBooking{}, fmt.Errorf("booking %d not found", index)
	}

	return protocol.VMBooking{
		VmTypeId:    booking.VmTypeId,
		PPS:         booking.PricePerSecond,
		Miner:       &booking.Miner,
		Index:       booking.Index,
		User:        &booking.User,
		BookedAt:    booking.BookedAt.Uint64(),
		LastPayment: booking.LastPayment.Uint64(),
	}, nil
}

func (b *Broker) GetTime() (uint64, error) {
	t, err := b.session.GetTime()
	if err != nil {
		return 0, err
	}
	return t.Uint64(), nil
}

func (b *Broker) GetMinersBookings() ([]protocol.VMBooking, error) {
	// if err := b.setDecimals(); err != nil {
	// 	return nil, err
	// }

	bookings, err := b.session.FindBookingsByMiner(b.transactOpts.From)
	if err != nil {
		return nil, fmt.Errorf("error executing FindBookingsByMiner: %v", err)
	}
	var result []protocol.VMBooking
	for _, booking := range bookings {

		result = append(result, protocol.VMBooking{
			VmTypeId:    booking.VmTypeId,
			PPS:         booking.PricePerSecond,
			Miner:       &booking.Miner,
			Index:       booking.Index,
			User:        &booking.User,
			BookedAt:    booking.BookedAt.Uint64(),
			LastPayment: booking.LastPayment.Uint64(),
		})
	}
	return result, nil
}

// func (b *Broker) AbortBooking(index uint64, abortType protocol.AbortType) error {
// 	tx, err := b.session.AbortBooking(index, abortType.ToSolidityType())
// 	if err != nil {
// 		return err
// 	}

// 	return b.waitForTx(tx)
// }

// func (b *Broker) ClaimExpired(index uint64) error {
// 	tx, err := b.session.ClaimExpired(index)
// 	if err != nil {
// 		return err
// 	}

// 	return b.waitForTx(tx)
// }

// func (b *Broker) ExtendBooking(index uint64, secs int) error {
// 	tx, err := b.session.ExtendBooking(index, big.NewInt(int64(secs)))
// 	if err != nil {
// 		return err
// 	}

// 	return b.waitForTx(tx)
// }

func (b *Broker) GetUserBookings() ([]protocol.VMBooking, error) {
	userBookings, err := b.session.GetUsersBookings(crypto.PubkeyToAddress(b.GetPrivateKey().PublicKey))
	if err != nil {
		return nil, err
	}

	result := make([]protocol.VMBooking, 0, len(userBookings))

	for i := range userBookings {
		result = append(result, protocol.VMBooking{
			VmTypeId:    userBookings[i].VmTypeId,
			PPS:         userBookings[i].PricePerSecond,
			Miner:       &userBookings[i].Miner,
			Index:       userBookings[i].Index,
			User:        &userBookings[i].User,
			BookedAt:    userBookings[i].BookedAt.Uint64(),
			LastPayment: userBookings[i].LastPayment.Uint64(),
		})
	}

	return result, nil
}

// func (b *Broker) GetMinerUrl(address *common.Address) (string, error) {
// 	urlBytes, err := b.session.GetMinerUrl(*address)
// 	if err != nil {
// 		return "", err
// 	}
// 	url, err := converters.Bytes32ToUrl(urlBytes)
// 	return url, err
// }

// func (b *Broker) SetMinerUrlIfNeeded(newUrl string) error {
// 	oldUrl, err := b.GetMinerUrl(&b.transactOpts.From)
// 	if err != nil {
// 		return err
// 	}
// 	if oldUrl == newUrl {
// 		return nil

// 	}

// 	urlBytes, err := converters.UrlToBytes32(newUrl)
// 	if err != nil {
// 		return err
// 	}

// 	_, err = b.session.SetMunerUrl(urlBytes)
// 	return err
// }
