package broker

// func (b *Broker) WithdrawCoins() error {
// 	if err := b.setDecimals(); err != nil {
// 		return err
// 	}

// 	tx, err := b.session.WithdrawCoin()
// 	if err != nil {
// 		return err
// 	}

// 	return b.waitForTx(tx.Hash())
// }

// func (b *Broker) Balance() (int, error) {
// 	if err := b.setDecimals(); err != nil {
// 		return 0, err
// 	}

// 	amount, err := b.session.UserBalance()
// 	if err != nil {
// 		return 0, err
// 	}

// 	return int(amount.Int64()), nil
// }

// func (b *Broker) DepositBalance() (int, error) {
// 	if err := b.setDecimals(); err != nil {
// 		return 0, err
// 	}

// 	amount, err := b.session.UserDeposit()
// 	if err != nil {
// 		return 0, err
// 	}

// 	return int(amount.Int64()), nil
// }

// func (b *Broker) LockedBalance() (int, error) {
// 	if err := b.setDecimals(); err != nil {
// 		return 0, err
// 	}

// 	amount, err := b.session.UserLockedBalance()
// 	if err != nil {
// 		return 0, err
// 	}

// 	return int(amount.Int64()), nil
// }

// func (b *Broker) UserTokenBalance() (int, error) {
// 	if err := b.setDecimals(); err != nil {
// 		return 0, err
// 	}

// 	amount, err := b.session.UserTokenBalance()
// 	if err != nil {
// 		return 0, err
// 	}

// 	return int(amount.Int64()), nil
// }

// // func (b *Broker) amountToCoins(amount *big.Int) float64 {
// // 	coin := math.Pow(10, float64(b.decimals))

// // 	return float64(amount.Int64()) / coin
// // }

// // func (b *Broker) coinsToAmount(coins float64) *big.Int {
// // 	coinsInt := int64(coins * math.Pow(10, float64(b.decimals)))

// // 	return big.NewInt(coinsInt)
// // }
