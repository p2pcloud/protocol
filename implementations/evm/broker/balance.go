package broker

import (
	"math/big"
)

func (b *Broker) GetStablecoinBalance() (uint64, uint64, error) {
	free, locked, err := b.session.GetStablecoinBalance(*b.GetMyAddress())
	if err != nil {
		return 0, 0, err
	}

	return free.Uint64(), locked.Uint64(), nil
}

func (b *Broker) DepositStablecoin(amount uint64) error {
	amt := big.NewInt(0).SetUint64(amount)

	_, err := b.EstimateGas("DepositStablecoin", amt)
	if err != nil {
		return err
	}

	tx, err := b.session.DepositStablecoin(amt)
	if err != nil {
		return err
	}

	return b.waitForTx(tx)
}

func (b *Broker) WithdrawStablecoin(amount uint64) error {
	amt := big.NewInt(int64(amount))

	_, err := b.EstimateGas("WithdrawStablecoin", amt)
	if err != nil {
		return err
	}

	tx, err := b.session.WithdrawStablecoin(amt)
	if err != nil {
		return err
	}

	return b.waitForTx(tx)
}

func (b *Broker) ClaimPayment(offerIndex uint64) error {
	_, err := b.EstimateGas("ClaimPayment", offerIndex)
	if err != nil {
		return err
	}

	tx, err := b.session.ClaimPayment(offerIndex)
	if err != nil {
		return err
	}

	return b.waitForTx(tx)
}
