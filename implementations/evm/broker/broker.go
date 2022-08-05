package broker

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/p2pcloud/protocol/implementations/evm/contracts"
)

type Broker struct {
	backend         bind.ContractBackend
	transactOpts    *bind.TransactOpts
	contractAddress common.Address
	session         contracts.BrokerSession
	privateKey      *ecdsa.PrivateKey
	waitForTx       func(hash common.Hash) error
	updateCh        chan<- common.Address

	mu                *sync.Mutex
	decimals          uint8
	stableCoinAddress *common.Address
}

func NewBroker(
	backend bind.ContractBackend,
	privateKey *ecdsa.PrivateKey,
	contractAddressStr string,
	chanId int64,
	waitForTx func(hash common.Hash) error,
	updCh chan<- common.Address,
) (*Broker, error) {
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
		waitForTx:       waitForTx,
		mu:              &sync.Mutex{},
		updateCh:        updCh,
	}

	b.RegenerateSession()

	return b, nil
}

func (b *Broker) GetPrivateKey() *ecdsa.PrivateKey {
	return b.privateKey
}

func (b *Broker) ContractAddress() common.Address {
	return b.contractAddress
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

func (b *Broker) DeployContract() (string, error) {
	address, tx, _, err := contracts.DeployBroker(
		b.transactOpts,
		b.backend,
		// community,
	)

	if err != nil {
		return "", fmt.Errorf("could not deploy broker: %v", err)
	}
	b.contractAddress = address

	b.RegenerateSession()

	return address.String(), b.waitForTx(tx.Hash())
}

func (b *Broker) GetMyAddress() *common.Address {
	return &b.transactOpts.From
}

// func (b *Broker) DepositCoin(coins float64) error {
// 	if err := b.setDecimals(); err != nil {
// 		return err
// 	}

// 	tx, err := b.session.DepositCoin(b.coinsToAmount(coins))
// 	if err != nil {
// 		return err
// 	}

// 	return b.waitForTx(tx.Hash())
// }

// func (b *Broker) UserAllowance() (float64, error) {
// 	if err := b.setDecimals(); err != nil {
// 		return 0, err
// 	}

// 	amount, err := b.session.UserAllowance()
// 	if err != nil {
// 		return 0, err
// 	}

// 	return b.amountToCoins(amount), nil
// }

// func (b *Broker) setDecimals() error {
// 	if b.decimals > 0 {
// 		return nil
// 	}

// 	decimals, err := b.session.GetCoinDecimals()
// 	if err != nil {
// 		return err
// 	}

// 	b.decimals = decimals

// 	return nil
// }
