package inmem

import (
	"context"
	"sync"

	"github.com/p2pcloud/protocol/implementations/evm/ledger"
)

type Ledger struct {
	mu             *sync.Mutex
	transactionLog map[string]ledger.Transaction
	ledger         map[string]ledger.Balance
}

func NewMemoryLedger() ledger.Ledger {
	return &Ledger{
		mu:             &sync.Mutex{},
		transactionLog: make(map[string]ledger.Transaction),
		ledger:         make(map[string]ledger.Balance),
	}
}

// Deposit adds free money for user and saves transaction in transaction log
func (l *Ledger) Deposit(ctx context.Context, transaction ledger.Transaction, user string) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.transactionLog[user] = transaction

	balance := l.ledger[user]
	balance.Free += transaction.Amount
	l.ledger[user] = balance

	return nil
}

// Withdraw removes free money from user and saves the transaction
func (l *Ledger) Withdraw(ctx context.Context, transaction ledger.Transaction, user string) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	balance := l.ledger[user]
	if balance.Free-transaction.Amount < 0 {
		return ledger.ErrWithdrawNotAvailable
	}

	l.transactionLog[user] = transaction

	balance.Free -= transaction.Amount
	l.ledger[user] = balance

	return nil
}

// CheckBalance checks whether free money is greater than or equal to transaction amount
func (l *Ledger) CheckBalance(ctx context.Context, transaction ledger.Transaction, user string) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	balance := l.ledger[user]
	if balance.Free-transaction.Amount < 0 {
		return ledger.ErrWithdrawNotAvailable
	}

	return nil
}
