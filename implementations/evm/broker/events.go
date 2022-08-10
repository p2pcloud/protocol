package broker

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/p2pcloud/protocol"
)

func (b *Broker) GetComplaintEvents(filter protocol.ComplaintFilterOpts) ([]protocol.ComplaintEvent, error) {
	var usersFilter []common.Address
	if filter.User != nil {
		usersFilter = make([]common.Address, 0)
		usersFilter = append(usersFilter, *filter.User)
	}

	var minersFilter []common.Address
	if filter.Miner != nil {
		minersFilter = make([]common.Address, 0)
		minersFilter = append(minersFilter, *filter.Miner)
	}

	var reasonsFilter []uint8
	if filter.StopReason != nil {
		reasonsFilter = make([]uint8, 0)
		reasonsFilter = append(reasonsFilter, uint8(*filter.StopReason))
	}

	opts := &bind.FilterOpts{
		Start:   filter.StartBlock,
		End:     filter.EndBlock,
		Context: nil,
	}

	filterer, err := b.EventsFilter.FilterComplaint(opts, usersFilter, minersFilter, reasonsFilter)
	if err != nil {
		return nil, err
	}
	events := make([]protocol.ComplaintEvent, 0)

	for filterer.Next() {
		event := filterer.Event
		events = append(events, protocol.ComplaintEvent{
			User:   &event.User,
			Miner:  &event.Miner,
			Reason: protocol.StopReason(event.Reason),
		})
	}
	return events, nil
}

func (b *Broker) GetPaymentEvents(filter protocol.PaymentFilterOpts) ([]protocol.PaymentEvent, error) {
	var usersFilter []common.Address
	if filter.User != nil {
		usersFilter = make([]common.Address, 0)
		usersFilter = append(usersFilter, *filter.User)
	}

	var minersFilter []common.Address
	if filter.Miner != nil {
		minersFilter = make([]common.Address, 0)
		minersFilter = append(minersFilter, *filter.Miner)
	}

	opts := &bind.FilterOpts{
		Start:   filter.StartBlock,
		End:     filter.EndBlock,
		Context: nil,
	}

	filterer, err := b.EventsFilter.FilterPayment(opts, usersFilter, minersFilter)
	if err != nil {
		return nil, err
	}
	events := make([]protocol.PaymentEvent, 0)

	for filterer.Next() {
		event := filterer.Event
		events = append(events, protocol.PaymentEvent{
			User:   &event.User,
			Miner:  &event.Miner,
			Amount: event.Amount.Uint64(),
		})
	}
	return events, nil
}
