package broker

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/p2pcloud/protocol"
	"github.com/p2pcloud/protocol/pkg/converters"
)

func (b *Broker) AddOffer(offer protocol.Offer, callbackUrl string) error {
	err := b.SetMinerUrlIfNeeded(callbackUrl)
	if err != nil {
		return err
	}

	tx, err := b.session.AddOffer(
		big.NewInt(int64(offer.PPS)),
		big.NewInt(int64(offer.VmTypeId)),
		big.NewInt(int64(offer.Availablility)),
	)
	if err != nil {
		return err
	}

	return b.waitForTx(tx.Hash())
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
			VmTypeId:      int(offer.VmTypeId.Int64()),
			PPS:           int(offer.PricePerSecond.Int64()),
			Availablility: int(offer.MachinesAvailable.Int64()),
			Miner:         offer.Miner,
			Index:         int(offer.Index.Int64()),
		})
	}
	return result, nil
}

func (b *Broker) GetAvailableOffers(vmTypeId int) ([]protocol.Offer, error) {
	// if err := b.setDecimals(); err != nil {
	// 	return nil, err
	// }

	offers, err := b.session.GetAvailableOffersByType(big.NewInt(int64(vmTypeId)))
	if err != nil {
		return nil, err
	}
	var result []protocol.Offer
	for _, offer := range offers {
		result = append(result, protocol.Offer{
			VmTypeId:      int(offer.VmTypeId.Int64()),
			PPS:           int(offer.PricePerSecond.Int64()),
			Availablility: int(offer.MachinesAvailable.Int64()),
			Miner:         offer.Miner,
			Index:         int(offer.Index.Int64()),
		})
	}
	return result, nil
}

func (b *Broker) RemoveOffer(offerId int) error {
	tx, err := b.session.RemoveOffer(big.NewInt(int64(offerId)))
	if err != nil {
		return err
	}
	return b.waitForTx(tx.Hash())
}

func (b *Broker) UpdateOffer(offer protocol.Offer) error {
	// if err := b.setDecimals(); err != nil {
	// 	return err
	// }

	tx, err := b.session.UpdateOffer(
		big.NewInt(int64(offer.Index)),
		big.NewInt(int64(offer.Availablility)),
		big.NewInt(int64(offer.PPS)),
	)
	if err != nil {
		return err
	}

	return b.waitForTx(tx.Hash())
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

	return b.waitForTx(tx.Hash())
}
