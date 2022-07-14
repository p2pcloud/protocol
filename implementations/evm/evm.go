package evm

import (
	"crypto/ecdsa"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/p2pcloud/protocol"
	"github.com/p2pcloud/protocol/implementations/evm/broker"
	"github.com/p2pcloud/protocol/pkg/keyring"
)

type EVMImplementation struct {
	broker protocol.BlockchainIface
}

var _ protocol.BlockchainIface = (*EVMImplementation)(nil)

func NewEVMImplementation(privateKey string, contractAddress string, rpcEndpoint string, chanId int64) (*EVMImplementation, error) {
	privateKeyDecoded, err := keyring.DecodePrivateKey(privateKey)
	if err != nil {
		return nil, err
	}

	if rpcEndpoint == "" {
		return nil, fmt.Errorf("rpc_endpoint is not set")
	}

	web3Client, err := ethclient.Dial(rpcEndpoint)
	if err != nil {
		return nil, fmt.Errorf("could not connect to web3 at \"%s\": %v", rpcEndpoint, err)
	}

	myBroker, err := broker.NewBroker(web3Client, privateKeyDecoded, contractAddress, chanId)
	if err != nil {
		return nil, err
	}
	return &EVMImplementation{
		broker: myBroker,
	}, nil

}

func (a *EVMImplementation) DeployContracts() ([]string, error) {
	addresses, err := a.broker.DeployContracts()
	return addresses, err
}

func (a *EVMImplementation) AddOffer(offer protocol.Offer, callbackUrl string) error {
	return a.broker.AddOffer(offer, callbackUrl)
}

func (a *EVMImplementation) GetMyOffers() ([]protocol.Offer, error) {
	return a.broker.GetMyOffers()
}

func (a *EVMImplementation) UpdateOffer(offer protocol.Offer) error {
	return a.broker.UpdateOffer(offer)
}

func (a *EVMImplementation) GetPrivateKey() *ecdsa.PrivateKey {
	return a.broker.GetPrivateKey()
}

func (a *EVMImplementation) GetMtlsHash(address *common.Address) (string, error) {
	return a.broker.GetMtlsHash(address)
}

func (a *EVMImplementation) GetBooking(index int) (*protocol.VMBooking, error) {
	return a.broker.GetBooking(index)
}

func (a *EVMImplementation) RegisterMtlsHashIfNeeded(mtlsHash string) error {
	return a.broker.RegisterMtlsHashIfNeeded(mtlsHash)
}

func (a *EVMImplementation) GetAvailableOffers(vmTypeId int) ([]protocol.Offer, error) {
	return a.broker.GetAvailableOffers(vmTypeId)
}

func (a *EVMImplementation) BookVM(offerIndex, seconds int) error {
	return a.broker.BookVM(offerIndex, seconds)
}

func (a *EVMImplementation) GetUsersBookings() ([]protocol.VMBooking, error) {
	return a.broker.GetUsersBookings()
}

func (a *EVMImplementation) GetMyAddress() *common.Address {
	return a.broker.GetMyAddress()
}

func (a *EVMImplementation) GetMinerUrl(address *common.Address) (string, error) {
	return a.broker.GetMinerUrl(address)
}

func (a *EVMImplementation) SetMinerUrlIfNeeded(newUrl string) error {
	return a.broker.SetMinerUrlIfNeeded(newUrl)
}

func (a *EVMImplementation) GetTime() (int, error) {
	return a.broker.GetTime()
}

func (a *EVMImplementation) GetMinersBookings() ([]protocol.VMBooking, error) {
	return a.broker.GetMinersBookings()
}

func (a *EVMImplementation) RegenerateSession() error {
	return a.broker.RegenerateSession()
}

func (a *EVMImplementation) DepositCoin(amount int64) error {
	return a.broker.DepositCoin(amount)
}

func (a *EVMImplementation) WithdrawCoin(amount int64) error {
	return a.broker.WithdrawCoin(amount)
}

func (a *EVMImplementation) Balance() (int64, error) {
	return a.broker.Balance()
}

func (a *EVMImplementation) TestApprove(to *common.Address, amount int64) error {
	return a.broker.TestApprove(to, amount)
}

func (a *EVMImplementation) UserTokenBalance() (int64, error) {
	return a.broker.UserTokenBalance()
}

func (a *EVMImplementation) UserAllowance(address common.Address) (int64, error) {
	return a.broker.UserAllowance(address)
}
