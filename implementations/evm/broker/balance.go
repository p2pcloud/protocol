package broker

import (
	"math/big"
)

func (b *Broker) GetStablecoinBalance() (int, int, error) {
	free, locked, err := b.session.GetStablecoinBalance(*b.GetMyAddress())
	if err != nil {
		return 0, 0, err
	}

	return int(free.Int64()), int(locked.Int64()), nil
}

func (b *Broker) DepositStablecoin(amount int) error {
	amt := big.NewInt(int64(amount))

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

func (b *Broker) WithdrawStablecoin(amount int) error {
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
