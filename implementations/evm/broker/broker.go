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
	tokenAddress    common.Address
	session         contracts.BrokerSession
	privateKey      *ecdsa.PrivateKey
	commit          func()
}

func NewBroker(
	backend bind.ContractBackend,
	privateKey *ecdsa.PrivateKey,
	contractAddressStr string,
	chanId int64,
	tokenAddress common.Address,
	commit func(),
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
		tokenAddress:    tokenAddress,
		commit:          commit,
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
		b.tokenAddress,
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

func (b *Broker) GetStableCoinSession() protocol.StableCoinSessionIface {
	return &b.session
}
