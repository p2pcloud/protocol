package broker

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/p2pcloud/protocol/implementations/evm/contracts"
)

type Broker struct {
	backend         bind.ContractBackend
	transactOpts    *bind.TransactOpts
	contractAddress common.Address
	session         contracts.BrokerSession
	privateKey      *ecdsa.PrivateKey
	waitForTx       func(tx *types.Transaction) error
	updateCh        chan<- common.Address

	mu                *sync.Mutex
	stableCoinAddress *common.Address
}

func NewBroker(
	backend bind.ContractBackend,
	privateKey *ecdsa.PrivateKey,
	contractAddressStr string,
	chanId int64,
	waitForTx func(tx *types.Transaction) error,
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

	instance, err := contracts.NewBroker(b.contractAddress, b.backend)
	if err != nil {
		return nil, err
	}

	b.session = contracts.BrokerSession{
		Contract:     instance,
		TransactOpts: *b.transactOpts,
		CallOpts: bind.CallOpts{
			Pending: false,                // Whether to operate on the pending state or the last known one
			From:    b.transactOpts.From, // Optional the sender address, otherwise the first account is used
			Context: context.Background(),
		},
	}

	return b, nil
}

func (b *Broker) EstimateGas(method string, args ...interface{}) (uint64, error) {
	abi, err := abi.JSON(strings.NewReader(contracts.BrokerABI))
	if err != nil {
		return 0, err
	}

	input, err := abi.Pack(method, args...)
	if err != nil {
		return 0, err
	}

	callMsg := ethereum.CallMsg{
		From:     b.transactOpts.From,
		To:       &b.contractAddress,
		Gas:      0,
		GasPrice: big.NewInt(0),
		Value:    big.NewInt(0),
		Data:     input,
	}

	gasLimit, err := b.backend.EstimateGas(context.Background(), callMsg)
	return gasLimit, err
}

func (b *Broker) GetPrivateKey() *ecdsa.PrivateKey {
	return b.privateKey
}

func (b *Broker) ContractAddress() common.Address {
	return b.contractAddress
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

	return address.String(), b.waitForTx(tx)
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

// 	return b.waitForTx(tx)
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
