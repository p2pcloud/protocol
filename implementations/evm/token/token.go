package token

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/p2pcloud/protocol"
	"github.com/p2pcloud/protocol/implementations/evm/contracts"
	"github.com/sirupsen/logrus"
	"math"
	"math/big"
	"sync"
)

type Token struct {
	coin         *big.Int
	backend      bind.ContractBackend
	transactOpts *bind.TransactOpts
	session      contracts.TokenSession
	privateKey   *ecdsa.PrivateKey

	commit func()
	upd    <-chan common.Address

	mu              *sync.Mutex
	decimals        uint8
	contractAddress common.Address
}

type Params struct {
	Decimals           *uint8
	Backend            bind.ContractBackend
	PrivateKey         *ecdsa.PrivateKey
	ContractAddressStr string
	ChainID            int64
	Commit             func()
	UpdCh              <-chan common.Address
}

func NewToken(params *Params) protocol.TokenIface {
	commit := func() {}
	if params.Commit != nil {
		commit = params.Commit
	}

	transactOpts, _ := bind.NewKeyedTransactorWithChainID(params.PrivateKey, big.NewInt(params.ChainID))

	b := &Token{
		backend:         params.Backend,
		transactOpts:    transactOpts,
		contractAddress: common.HexToAddress(params.ContractAddressStr),
		privateKey:      params.PrivateKey,
		commit:          commit,
		mu:              &sync.Mutex{},
		upd:             params.UpdCh,
	}

	if params.Decimals != nil {
		b.decimals = *params.Decimals
	}

	return b
}

func (t *Token) GetPrivateKey() *ecdsa.PrivateKey {
	return t.privateKey
}

func (t *Token) GetContractAddress() common.Address {
	return t.contractAddress
}

func (t *Token) DeployContract(initialSupply float64) (*common.Address, error) {
	defer t.commit()
	address, _, _, err := contracts.DeployToken(
		t.transactOpts,
		t.backend,
		t.coinsToAmount(initialSupply),
	)
	if err != nil {
		return nil, fmt.Errorf("could not deploy token: %v", err)
	}

	t.commit()

	if err = t.updateToken(address); err != nil {
		return nil, err
	}

	return &address, nil
}

func (t *Token) RegenerateSession() error {
	instance, err := contracts.NewToken(t.contractAddress, t.backend)
	if err != nil {
		return err
	}

	t.session = contracts.TokenSession{
		Contract:     instance,
		TransactOpts: *t.transactOpts,
		CallOpts: bind.CallOpts{
			Pending: false,               // Whether to operate on the pending state or the last known one
			From:    t.transactOpts.From, // Optional the sender address, otherwise the first account is used
			Context: context.Background(),
		},
	}
	return nil
}

func (t *Token) BalanceOf(address common.Address) (float64, error) {
	amount, err := t.session.BalanceOf(address)
	if err != nil {
		return 0, err
	}

	return t.amountToCoins(amount), nil
}

func (t *Token) Transfer(to common.Address, coins float64) error {
	defer t.commit()

	_, err := t.session.Transfer(to, t.coinsToAmount(coins))

	return err
}

func (t *Token) Approve(to common.Address, coins float64) error {
	defer t.commit()

	_, err := t.session.Approve(to, t.coinsToAmount(coins))

	return err
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

func (t *Token) StartUp() error {
	return t.updateToken(t.contractAddress)
}

func (t *Token) asyncUpdate() {
	for {
		select {
		case addr := <-t.upd:
			if err := t.updateToken(addr); err != nil {
				logrus.WithField("updated", addr.Hex()).Error("failed update token")
			}
		}
	}
}

func (t *Token) updateToken(contractAddr common.Address) error {
	t.mu.Lock()
	t.contractAddress = contractAddr
	t.mu.Unlock()

	t.RegenerateSession()

	decimals, err := t.session.Decimals()
	if err != nil {
		return err
	}

	t.mu.Lock()
	t.decimals = decimals
	t.mu.Unlock()

	return nil
}
