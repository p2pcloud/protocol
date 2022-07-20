package broker

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"math"
	"math/big"
	"sync"

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
	commit          func()
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
	commit func(),
	updCh chan<- common.Address,
) (protocol.BrokerIface, error) {
	if chanId == 0 {
		return nil, fmt.Errorf("chanId is 0. please set it to a valid value")
	}

	if commit == nil {
		commit = func() {}
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
		commit:          commit,
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

func (b *Broker) DeployContracts() ([]string, error) {
	defer b.commit()

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
	defer b.commit()

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

func (b *Broker) DepositCoin(coins float64) error {
	defer b.commit()

	if err := b.setDecimals(); err != nil {
		return err
	}

	_, err := b.session.DepositCoin(b.coinsToAmount(coins))

	return err
}

func (b *Broker) WithdrawCoin(coins float64) error {
	defer b.commit()

	if err := b.setDecimals(); err != nil {
		return err
	}

	_, err := b.session.WithdrawCoin(b.coinsToAmount(coins))

	return err
}

func (b *Broker) Balance() (float64, error) {
	if err := b.setDecimals(); err != nil {
		return 0, err
	}

	amount, err := b.session.UserBalance()
	if err != nil {
		return 0, err
	}

	return b.amountToCoins(amount), nil
}

func (b *Broker) UserTokenBalance() (float64, error) {
	if err := b.setDecimals(); err != nil {
		return 0, err
	}

	amount, err := b.session.UserTokenBalance()
	if err != nil {
		return 0, err
	}

	return b.amountToCoins(amount), nil
}

func (b *Broker) UserAllowance() (float64, error) {
	if err := b.setDecimals(); err != nil {
		return 0, err
	}

	amount, err := b.session.UserAllowance()
	if err != nil {
		return 0, err
	}

	return b.amountToCoins(amount), nil
}

func (b *Broker) SetStablecoinAddress(address common.Address) error {
	_, err := b.session.SetStablecoinAddress(address)
	if err != nil {
		return err
	}
	b.commit()

	decimals, err := b.session.GetCoinDecimals()
	if err != nil {
		return err
	}

	select {
	case b.updateCh <- address:
	default:
		logrus.WithField("updated", address.Hex()).Error("failed to send to upd ch")
	}

	b.mu.Lock()
	defer b.mu.Unlock()

	b.stableCoinAddress = nil
	b.decimals = decimals

	return nil
}

func (b *Broker) GetStablecoinAddress() (common.Address, error) {
	if b.stableCoinAddress != nil {
		return *b.stableCoinAddress, nil
	}

	addr, err := b.session.GetStablecoinAddress()
	if err != nil {
		return [20]byte{}, err
	}

	b.stableCoinAddress = &addr

	return addr, nil
}

func (b *Broker) setDecimals() error {
	if b.decimals > 0 {
		return nil
	}

	decimals, err := b.session.GetCoinDecimals()
	if err != nil {
		return err
	}

	b.decimals = decimals

	return nil
}

func (b *Broker) amountToCoins(amount *big.Int) float64 {
	coin := math.Pow(10, float64(b.decimals))

	return float64(amount.Int64()) / coin
}

func (b *Broker) coinsToAmount(coins float64) *big.Int {
	coinsInt := int64(coins * math.Pow(10, float64(b.decimals)))

	return big.NewInt(coinsInt)
}
