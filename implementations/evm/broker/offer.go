package broker

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/p2pcloud/protocol"
	"github.com/p2pcloud/protocol/pkg/converters"
)

func (b *Broker) AddOffer(offer protocol.Offer, callbackUrl string) error {
	_, err := b.EstimateGas("AddOffer", offer.PPS, offer.VmTypeId, offer.Availablility)
	if err != nil {
		return err
	}

	err = b.SetMinerUrlIfNeeded(callbackUrl)
	if err != nil {
		return err
	}

	tx, err := b.session.AddOffer(
		offer.PPS,
		offer.VmTypeId,
		offer.Availablility,
	)
	if err != nil {
		return err
	}

	return b.waitForTx(tx)
}

func (b *Broker) GetMyOffers() ([]protocol.Offer, error) {
	// if err := b.setDecimals(); err != nil {
	// 	return nil, err
	// }

	offers, err := b.session.GetMinersOffers(b.transactOpts.From)
	if err != nil {
		return nil, err
	}
	var result []protocol.Offer
	for _, offer := range offers {
		result = append(result, protocol.Offer{
			VmTypeId:      offer.VmTypeId,
			PPS:           offer.PricePerSecond,
			Availablility: offer.MachinesAvailable,
			Miner:         offer.Miner,
			Index:         offer.Index,
		})
	}
	return result, nil
}

func (b *Broker) GetAvailableOffers(vmTypeId uint64) ([]protocol.Offer, error) {
	// if err := b.setDecimals(); err != nil {
	// 	return nil, err
	// }

	offers, err := b.session.GetAvailableOffersByType(vmTypeId)
	if err != nil {
		return nil, err
	}
	var result []protocol.Offer
	for _, offer := range offers {
		result = append(result, protocol.Offer{
			VmTypeId:      offer.VmTypeId,
			PPS:           offer.PricePerSecond,
			Availablility: offer.MachinesAvailable,
			Miner:         offer.Miner,
			Index:         offer.Index,
		})
	}
	return result, nil
}

func (b *Broker) RemoveOffer(offerId uint64) error {
	_, err := b.EstimateGas("RemoveOffer", offerId)
	if err != nil {
		return err
	}

	tx, err := b.session.RemoveOffer(offerId)
	if err != nil {
		return err
	}
	return b.waitForTx(tx)
}

func (b *Broker) UpdateOffer(offer protocol.Offer) error {
	// if err := b.setDecimals(); err != nil {
	// 	return err
	// }

	tx, err := b.session.UpdateOffer(
		offer.Index,
		offer.Availablility,
		offer.PPS,
	)
	if err != nil {
		return err
	}

	return b.waitForTx(tx)
}

func (b *Broker) GetMinerUrl(address *common.Address) (string, error) {
	urlBytes, err := b.session.GetMinerUrl(*address)
	if err != nil {
		return "", err
	}
	url, err := converters.Bytes32ToUrl(urlBytes)
	return url, err
}

func (b *Broker) SetMinerUrlIfNeeded(newUrl string) error {
	oldUrl, err := b.GetMinerUrl(&b.transactOpts.From)
	if err != nil {
		return err
	}
	if oldUrl == newUrl {
		return nil

	}

	urlBytes, err := converters.UrlToBytes32(newUrl)
	if err != nil {
		return err
	}

	tx, err := b.session.SetMunerUrl(urlBytes)
	if err != nil {
		return err
	}

	return b.waitForTx(tx)
}
