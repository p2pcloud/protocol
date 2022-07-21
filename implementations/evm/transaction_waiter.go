package evm

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/cenkalti/backoff"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

var ErrTxFailed = errors.New("receipt failed")

type TransactionWaiter struct {
	client  *ethclient.Client
	timeout time.Duration
	backoff backoff.BackOff
}

func NewTransactionWaiter(
	client *ethclient.Client, timeout time.Duration,
) *TransactionWaiter {
	b := backoff.NewExponentialBackOff()
	b.MaxElapsedTime = time.Minute

	return &TransactionWaiter{
		client:  client,
		timeout: timeout,
		backoff: b,
	}
}

func (t *TransactionWaiter) WaitForTx(hash common.Hash) error {
	// run instantly, cause ticker firstly waits the period, helpful in tests
	ctx, cancel := context.WithTimeout(context.Background(), t.timeout)
	defer cancel()

	finished, err := t.waitForTx(ctx, hash)
	if err != nil {
		return err
	}

	if finished {
		return nil
	}

	// polling
	ticker := time.NewTicker(time.Second * 1)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-ticker.C:
			finished, err = t.waitForTx(ctx, hash)
			if err != nil {
				return err
			}

			if finished {
				return nil
			}
		}
	}
}

func (t *TransactionWaiter) waitForTx(ctx context.Context, hash common.Hash) (bool, error) {
	pending, err := t.isTxPending(ctx, hash)
	if err != nil {
		return false, err
	}

	if pending {
		return pending, nil
	}

	receipt, err := t.getTxReceipt(ctx, hash)
	if err != nil {
		return false, err
	}

	if receipt.Status == types.ReceiptStatusFailed {
		receiptData, _ := receipt.MarshalJSON()

		return false, fmt.Errorf("receipt failed, data %s\nerr: %w", string(receiptData), ErrTxFailed)
	}

	return true, nil
}

func (t *TransactionWaiter) isTxPending(ctx context.Context, hash common.Hash) (bool, error) {
	var pending bool

	err := backoff.Retry(func() error {
		_, txPending, err := t.client.TransactionByHash(ctx, hash)
		if err != nil {
			return err
		}

		pending = txPending

		return nil
	}, t.backoff)

	return pending, err
}

func (t *TransactionWaiter) getTxReceipt(ctx context.Context, hash common.Hash) (*types.Receipt, error) {
	var receipt *types.Receipt

	err := backoff.Retry(func() error {
		result, err := t.client.TransactionReceipt(ctx, hash)
		if err != nil {
			return err
		}

		receipt = result

		return nil
	}, t.backoff)

	return receipt, err
}
