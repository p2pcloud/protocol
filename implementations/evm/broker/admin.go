package broker

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// func (b *Broker) SetStablecoinAddress(address common.Address) error {
// 	tx, err := b.session.SetStablecoinAddress(address)
// 	if err != nil {
// 		return err
// 	}

// 	if err = b.waitForTx(tx.Hash()); err != nil {
// 		return err
// 	}

// 	decimals, err := b.session.GetCoinDecimals()
// 	if err != nil {
// 		return err
// 	}

// 	select {
// 	case b.updateCh <- address:
// 	default:
// 		logrus.WithField("updated", address.Hex()).Error("failed to send to upd ch")
// 	}

// 	b.mu.Lock()
// 	defer b.mu.Unlock()

// 	b.stableCoinAddress = nil
// 	b.decimals = decimals

// 	return nil
// }

func (b *Broker) GetStablecoinAddress() (common.Address, error) {
	if b.stableCoinAddress != nil {
		return *b.stableCoinAddress, nil
	}

	addr, err := b.session.GetStablecoinAddress()
	if err != nil {
		return [20]byte{}, err
	}

	b.stableCoinAddress = &addr

	return addr, nil
}

func (b *Broker) SetCommunityContract(address common.Address) error {
	tx, err := b.session.SetCommunityContract(address)
	if err != nil {
		return err
	}

	return b.waitForTx(tx.Hash())
}

func (b *Broker) GetCommunityContract() (common.Address, error) {
	return b.session.GetCommunityContract()
}

func (b *Broker) SetCommunityFee(fee int64) error {
	tx, err := b.session.SetCommunityFee(big.NewInt(fee))
	if err != nil {
		return err
	}

	return b.waitForTx(tx.Hash())
}

func (b *Broker) GetCommunityFee() (int64, error) {
	fee, err := b.session.GetCommunityFee()
	if err != nil {
		return 0, err
	}

	return fee.Int64(), nil
}
