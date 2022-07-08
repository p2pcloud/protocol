package broker

import (
	"fmt"
	"math/big"

	"github.com/p2pcloud/protocol"
)

func (b *Broker) BookVM(offerIndex, seconds int) error {
	_, err := b.session.BookVM(big.NewInt(int64(offerIndex)), big.NewInt(int64(seconds)))
	return err
}

func (b *Broker) GetUsersBookings() ([]protocol.VMBooking, error) {
	bookings, err := b.session.FindBookingsByUser(b.transactOpts.From)
	if err != nil {
		return nil, err
	}
	var result []protocol.VMBooking
	for _, booking := range bookings {
		result = append(result, protocol.VMBooking{
			VmTypeId:   int(booking.VmTypeId.Int64()),
			PPS:        int(booking.PricePerSecond.Int64()),
			Miner:      &booking.Miner,
			Index:      int(booking.Index.Int64()),
			User:       &booking.User,
			BookedTill: int(booking.BookedTill.Int64()),
		})
	}
	return result, nil
}

func (b *Broker) GetBooking(index int) (*protocol.VMBooking, error) {
	booking, err := b.session.GetBooking(uint64(index))
	if err != nil {
		return nil, err
	}

	if booking.Miner.Hex() == "0x0000000000000000000000000000000000000000" {
		return nil, fmt.Errorf("booking %d not found", index)
	}

	return &protocol.VMBooking{
		VmTypeId:   int(booking.VmTypeId.Int64()),
		PPS:        int(booking.PricePerSecond.Int64()),
		Miner:      &booking.Miner,
		Index:      int(booking.Index.Int64()),
		User:       &booking.User,
		BookedTill: int(booking.BookedTill.Int64()),
	}, nil
}

func (b *Broker) GetTime() (int, error) {
	t, err := b.session.GetTime()
	if err != nil {
		return 0, err
	}
	return int(t.Int64()), nil
}

func (b *Broker) GetMinersBookings() ([]protocol.VMBooking, error) {
	bookings, err := b.session.FindBookingsByMiner(b.transactOpts.From)
	if err != nil {
		return nil, fmt.Errorf("error executing FindBookingsByMiner: %v", err)
	}
	var result []protocol.VMBooking
	for _, booking := range bookings {

		result = append(result, protocol.VMBooking{
			VmTypeId:   int(booking.VmTypeId.Int64()),
			PPS:        int(booking.PricePerSecond.Int64()),
			Miner:      &booking.Miner,
			Index:      int(booking.Index.Int64()),
			User:       &booking.User,
			BookedTill: int(booking.BookedTill.Int64()),
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
