package stablecoin

import (
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/p2pcloud/protocol"
)

type StableCoin struct {
	coin         *big.Int
	backend      bind.ContractBackend
	transactOpts *bind.TransactOpts
	session      protocol.StableCoinSessionIface
	privateKey   *ecdsa.PrivateKey
	commit       func()
}

type Params struct {
	Decimals int64
	Backend  bind.ContractBackend
	Session  protocol.StableCoinSessionIface
	Commit   func()
}

func New(params *Params) *StableCoin {
	return &StableCoin{
		coin:    new(big.Int).Exp(big.NewInt(10), big.NewInt(params.Decimals), nil),
		backend: params.Backend,
		session: params.Session,
		commit:  params.Commit,
	}
}

func (s *StableCoin) DepositCoin(amount int64) error {
	defer s.commit()

	_, err := s.session.DepositCoin(new(big.Int).Mul(big.NewInt(amount), s.coin))

	return err
}

func (s *StableCoin) WithdrawCoin(amount int64) error {
	defer s.commit()

	_, err := s.session.WithdrawCoin(new(big.Int).Mul(big.NewInt(amount), s.coin))

	return err
}

func (s *StableCoin) Balance() (int64, error) {
	amount, err := s.session.UserBalance()
	if err != nil {
		return 0, err
	}

	return new(big.Int).Div(amount, s.coin).Int64(), nil
}

func (s *StableCoin) UserTokenBalance() (int64, error) {
	amount, err := s.session.UserTokenBalance()
	if err != nil {
		return 0, err
	}

	return new(big.Int).Div(amount, s.coin).Int64(), nil
}

func (s *StableCoin) UserAllowance(address common.Address) (int64, error) {
	amount, err := s.session.UserAllowance(address)
	if err != nil {
		return 0, err
	}

	return new(big.Int).Div(amount, s.coin).Int64(), nil
}
