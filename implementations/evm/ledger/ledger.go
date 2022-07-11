package ledger

import (
	"context"
	"errors"
)

var ErrWithdrawNotAvailable = errors.New("withdraw currently not available due " +
	"it is locked or balance is less than withdraw amount")

const (
	UnknownTxType TransactionType = iota
	DepositTxType
	WithdrawTxType
)

type TransactionType int

type Transaction struct {
	Amount uint64
	TxType TransactionType
}

type Balance struct {
	Free uint64
	Locked uint64
}

type Ledger interface {
	Deposit(ctx context.Context, transaction Transaction, user string) error
	Withdraw(ctx context.Context, transaction Transaction, user string) error
	CheckBalance(ctx context.Context, transaction Transaction, user string) error
}
