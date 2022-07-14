package broker

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/sirupsen/logrus"

	"github.com/p2pcloud/protocol"
	"github.com/p2pcloud/protocol/implementations/evm/contracts"
)

type Broker struct {
	backend         bind.ContractBackend
	transactOpts    *bind.TransactOpts
	contractAddress common.Address
	session         contracts.BrokerSession
	privateKey      *ecdsa.PrivateKey
}

func NewBroker(
	backend bind.ContractBackend, privateKey *ecdsa.PrivateKey, contractAddressStr string, chanId int64,
) (protocol.BlockchainIface, error) {
	if chanId == 0 {
		return nil, fmt.Errorf("chanId is 0. please set it to a valid value")
	}

	transactOpts, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(chanId))
	if err != nil {
		return nil, err
	}

	b := &Broker{
		backend:         backend,
		transactOpts:    transactOpts,
		contractAddress: common.HexToAddress(contractAddressStr),
		privateKey:      privateKey,
	}

	b.RegenerateSession()

	return b, nil
}

func (b *Broker) GetPrivateKey() *ecdsa.PrivateKey {
	return b.privateKey
}

func (b *Broker) RegenerateSession() error {
	instance, err := contracts.NewBroker(b.contractAddress, b.backend)
	if err != nil {
		return err
	}

	b.session = contracts.BrokerSession{
		Contract:     instance,
		TransactOpts: *b.transactOpts,
		CallOpts: bind.CallOpts{
			Pending: false,               // Whether to operate on the pending state or the last known one
			From:    b.transactOpts.From, // Optional the sender address, otherwise the first account is used
			Context: context.Background(),
		},
	}
	return nil
}

func (b *Broker) DeployContracts() ([]string, error) {
	address, _, _, err := contracts.DeployBroker(
		b.transactOpts,
		b.backend,
	)

	if err != nil {
		return nil, fmt.Errorf("could not deploy broker: %v", err)
	}
	b.contractAddress = address

	b.RegenerateSession()

	return []string{address.String()}, nil
}

func (b *Broker) GetMyAddress() *common.Address {
	return &b.transactOpts.From
}

func (b *Broker) GetMtlsHash(address *common.Address) (string, error) {
	hashBytes, err := b.session.GetMtlsHash(*address)
	if err != nil {
		return "", err
	}
	logrus.WithField("result", "0x"+hex.EncodeToString(hashBytes[:])).Debug("GetMtlsHash called")
	return "0x" + hex.EncodeToString(hashBytes[:]), nil
}

func (b *Broker) RegisterMtlsHashIfNeeded(mtlsHash string) error {
	logrus.WithField("mtlsHash", mtlsHash).Debug("RegisterMtlsHashIfNeeded called")
	oldMtlsHash, err := b.GetMtlsHash(b.GetMyAddress())
	if err != nil {
		return err
	}

	if oldMtlsHash == mtlsHash {
		return nil
	}

	var bytes20 [20]byte

	hashBytes := common.FromHex(mtlsHash)

	if len(hashBytes) != 20 {
		return fmt.Errorf("mtls hash is not 20 bytes long")
	}

	copy(bytes20[:], hashBytes)

	_, err = b.session.SetMtlsHash(bytes20)
	if err != nil {
		return fmt.Errorf("could not register mtls hash: %v", err)
	}
	return nil
}

func (b *Broker) DepositCoin(amount int64) error {
	usdc := new(big.Int).Exp(big.NewInt(10), big.NewInt(6), nil)

	_, err := b.session.DepositCoin(new(big.Int).Mul(usdc, big.NewInt(amount)))

	return err
}

func (b *Broker) WithdrawCoin(amount int64) error {
	usdc := new(big.Int).Exp(big.NewInt(10), big.NewInt(6), nil)

	_, err := b.session.WithdrawCoin(new(big.Int).Mul(usdc, big.NewInt(amount)))

	return err
}

func (b *Broker) Balance() (int64, error) {
	tx, err := b.session.UserBalance()
	if err != nil {
		return 0, err
	}

	usdc := new(big.Int).Exp(big.NewInt(10), big.NewInt(6), nil)

	return new(big.Int).Div(tx.Value(), usdc).Int64(), nil
}

func (b *Broker) TestApprove(to *common.Address, amount int64) error {
	usdc := new(big.Int).Exp(big.NewInt(10), big.NewInt(6), nil)

	_, err := b.session.TestApprove(*to, new(big.Int).Mul(usdc, big.NewInt(amount)))

	return err
}

func (b *Broker) UserTokenBalance() (int64, error) {
	usdc := new(big.Int).Exp(big.NewInt(10), big.NewInt(6), nil)

	tx, err := b.session.UserTokenBalance()
	if err != nil {
		return 0, err
	}

	return new(big.Int).Div(tx.Value(), usdc).Int64(), nil
}

func (b *Broker) UserAllowance(address common.Address) (int64, error) {
	usdc := new(big.Int).Exp(big.NewInt(10), big.NewInt(6), nil)

	tx, err := b.session.UserAllowance(address)
	if err != nil {
		return 0, err
	}

	return new(big.Int).Div(tx.Value(), usdc).Int64(), nil
}
