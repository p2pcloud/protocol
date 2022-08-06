package token

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/p2pcloud/protocol/implementations/evm/contracts"
)

type Token struct {
	backend      bind.ContractBackend
	transactOpts *bind.TransactOpts
	session      contracts.TokenSession
	privateKey   *ecdsa.PrivateKey

	waitForTx func(tx *types.Transaction) error

	mu              *sync.Mutex
	decimals        int16
	contractAddress common.Address
}

func NewToken(
	backend bind.ContractBackend,
	privateKey *ecdsa.PrivateKey,
	contractAddressStr string,
	chainId int64,
	waitForTx func(tx *types.Transaction) error,
) (*Token, error) {
	transactOpts, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(chainId))

	b := &Token{
		backend:         backend,
		transactOpts:    transactOpts,
		contractAddress: common.HexToAddress(contractAddressStr),
		privateKey:      privateKey,
		waitForTx:       waitForTx,
		mu:              &sync.Mutex{},
		decimals:        -1,
	}

	instance, err := contracts.NewToken(b.contractAddress, b.backend)
	if err != nil {
		return nil, err
	}

	b.session = contracts.TokenSession{
		Contract:     instance,
		TransactOpts: *b.transactOpts,
		CallOpts: bind.CallOpts{
			Pending: true,                // Whether to operate on the pending state or the last known one
			From:    b.transactOpts.From, // Optional the sender address, otherwise the first account is used
			Context: context.Background(),
		},
	}

	return b, nil
}

func (t *Token) GetPrivateKey() *ecdsa.PrivateKey {
	return t.privateKey
}

func (t *Token) GetContractAddress() common.Address {
	return t.contractAddress
}

func (t *Token) DeployContract() (*common.Address, error) {
	address, tx, _, err := contracts.DeployToken(
		t.transactOpts,
		t.backend,
	)

	if err != nil {
		return nil, fmt.Errorf("could not deploy token: %v", err)
	}

	return &address, t.waitForTx(tx)
}

func (t *Token) BalanceOf(address common.Address) (int, error) {
	amount, err := t.session.BalanceOf(address)
	if err != nil {
		return 0, err
	}

	return int(amount.Int64()), nil
}

func (t *Token) Transfer(to common.Address, amount int) error {
	tx, err := t.session.Transfer(to, big.NewInt(int64(amount)))
	if err != nil {
		return err
	}

	return t.waitForTx(tx)
}

func (t *Token) Approve(to common.Address, amount int) error {
	tx, err := t.session.Approve(to, big.NewInt(int64(amount)))
	if err != nil {
		return err
	}

	return t.waitForTx(tx)
}

func (t *Token) Allowance(from, address common.Address) (float64, error) {
	amount, err := t.session.Allowance(from, address)
	if err != nil {
		return 0, err
	}

	return t.amountToCoins(amount), nil
}

func (t *Token) amountToCoins(amount *big.Int) float64 {
	coin := math.Pow(10, float64(t.decimals))

	return float64(amount.Int64()) / coin
}

func (t *Token) coinsToAmount(coins float64) *big.Int {
	coinsInt := int64(coins * math.Pow(10, float64(t.decimals)))

	return big.NewInt(coinsInt)
}

func (t *Token) GetDecimals() (int16, error) {
	if t.decimals != -1 {
		return t.decimals, nil
	}

	t.mu.Lock()
	decimals, err := t.session.Decimals()
	if err != nil {
		return -1, err
	}

	t.decimals = int16(decimals)
	t.mu.Unlock()

	return t.decimals, nil
}
