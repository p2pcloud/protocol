package token

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/p2pcloud/protocol/implementations/evm/contracts"
	"math/big"
)

type Token struct {
	coin            *big.Int
	backend         bind.ContractBackend
	transactOpts    *bind.TransactOpts
	contractAddress common.Address
	Session         contracts.TokenSession
	privateKey      *ecdsa.PrivateKey

	commit func()
}

type Params struct {
	Decimals           int64
	Backend            bind.ContractBackend
	PrivateKey         *ecdsa.PrivateKey
	ContractAddressStr string
	ChainID            int64
	Commit             func()
}

func NewToken(params *Params) *Token {
	transactOpts, _ := bind.NewKeyedTransactorWithChainID(params.PrivateKey, big.NewInt(params.ChainID))

	b := &Token{
		coin:            new(big.Int).Exp(big.NewInt(10), big.NewInt(params.Decimals), nil),
		backend:         params.Backend,
		transactOpts:    transactOpts,
		contractAddress: common.HexToAddress(params.ContractAddressStr),
		privateKey:      params.PrivateKey,
		commit:          params.Commit,
	}

	return b
}

func (t *Token) GetPrivateKey() *ecdsa.PrivateKey {
	return t.privateKey
}

func (t *Token) GetContractAddress() common.Address {
	return t.contractAddress
}

func (t *Token) DeployContract(initialSupply int64) (*common.Address, error) {
	defer t.commit()

	address, _, _, err := contracts.DeployToken(
		t.transactOpts,
		t.backend,
		new(big.Int).Mul(big.NewInt(initialSupply), t.coin),
	)

	if err != nil {
		return nil, fmt.Errorf("could not deploy token: %v", err)
	}
	t.contractAddress = address

	t.RegenerateSession()

	return &address, nil
}

func (t *Token) RegenerateSession() error {
	instance, err := contracts.NewToken(t.contractAddress, t.backend)
	if err != nil {
		return err
	}

	t.Session = contracts.TokenSession{
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

func (t *Token) Balance(address common.Address) (int64, error) {
	balance, err := t.Session.BalanceOf(address)
	if err != nil {
		return 0, err
	}

	return new(big.Int).Div(balance, t.coin).Int64(), nil
}

func (t *Token) TransferForTests(from, to common.Address, amount int64) error {
	defer t.commit()

	_, err := t.Session.TransferForTests(
		from, to,
		new(big.Int).Mul(big.NewInt(amount), t.coin),
	)

	return err
}

func (t *Token) TestApprove(to common.Address, amount int64) error {
	defer t.commit()

	_, err := t.Session.Approve(to, new(big.Int).Mul(big.NewInt(amount), t.coin))

	return err
}
