package evm

import (
	"crypto/ecdsa"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/p2pcloud/protocol"
	"github.com/p2pcloud/protocol/implementations/evm/broker"
	"github.com/p2pcloud/protocol/implementations/evm/token"
	"github.com/p2pcloud/protocol/pkg/keyring"
)

type EVMImplementation struct {
	Broker *broker.Broker
	Token  *token.Token
}

var _ protocol.P2PCloudProtocolIface = (*EVMImplementation)(nil)

func NewEvmImplementationFromBackend(
	contractBackend bind.ContractBackend,
	privateKey string, brokerContractAddress string,
	chanId int64,
	waitForTx func(tx *types.Transaction) error,
) (*EVMImplementation, error) {

	privateKeyDecoded, err := keyring.DecodePrivateKey(privateKey)
	if err != nil {
		return nil, err
	}

	updCh := make(chan common.Address, 1)

	myBroker, err := broker.NewBroker(
		contractBackend, privateKeyDecoded, brokerContractAddress, chanId, waitForTx, updCh,
	)
	if err != nil {
		return nil, err
	}

	tokenAddr, err := myBroker.GetStablecoinAddress()
	if err != nil {
		return nil, err
	}

	tkn, err := token.NewToken(
		contractBackend, privateKeyDecoded, tokenAddr.Hex(), chanId, waitForTx,
	)
	if err != nil {
		return nil, err
	}

	return &EVMImplementation{
		Broker: myBroker,
		Token:  tkn,
	}, nil
}

func NewEVMImplementation(
	privateKey string, contractAddress, rpcEndpoint string, chanId int64,
) (*EVMImplementation, error) {
	if rpcEndpoint == "" {
		return nil, fmt.Errorf("rpc_endpoint is not set")
	}

	web3Client, err := ethclient.Dial(rpcEndpoint)
	if err != nil {
		return nil, fmt.Errorf("could not connect to web3 at \"%s\": %v", rpcEndpoint, err)
	}

	txWaiter := NewTransactionWaiter(web3Client, time.Minute)

	return NewEvmImplementationFromBackend(web3Client, privateKey, contractAddress, chanId, txWaiter.WaitForTx)
}

func (a *EVMImplementation) AddOffer(offer protocol.Offer, callbackUrl string) error {
	return a.Broker.AddOffer(offer, callbackUrl)
}

func (a *EVMImplementation) GetMyOffers() ([]protocol.Offer, error) {
	return a.Broker.GetMyOffers()
}

func (a *EVMImplementation) UpdateOffer(offer protocol.Offer) error {
	return a.Broker.UpdateOffer(offer)
}

func (a *EVMImplementation) GetPrivateKey() *ecdsa.PrivateKey {
	return a.Broker.GetPrivateKey()
}

func (a *EVMImplementation) ContractAddress() common.Address {
	return a.Broker.ContractAddress()
}

func (a *EVMImplementation) GetBooking(index uint64) (protocol.VMBooking, error) {
	return a.Broker.GetBooking(index)
}

func (a *EVMImplementation) GetAvailableOffers(vmTypeId uint64) ([]protocol.Offer, error) {
	return a.Broker.GetAvailableOffers(vmTypeId)
}

func (a *EVMImplementation) RemoveOffer(id uint64) error {
	return a.Broker.RemoveOffer(id)
}

func (a *EVMImplementation) DepositStablecoin(amount uint64) error {
	err := a.Token.Approve(a.Broker.ContractAddress(), amount)
	if err != nil {
		return err
	}

	return a.Broker.DepositStablecoin(amount)
}

func (a *EVMImplementation) WithdrawStablecoin(amount uint64) error {
	return a.Broker.WithdrawStablecoin(amount)
}

func (a *EVMImplementation) BookVM(offerIndex uint64) error {
	return a.Broker.BookVM(offerIndex)
}

func (a *EVMImplementation) StopVM(offerIndex uint64, reason protocol.StopReason) error {
	return a.Broker.StopVM(offerIndex, uint8(reason))
}

func (a *EVMImplementation) ClaimPayment(offerIndex uint64) error {
	return a.Broker.ClaimPayment(offerIndex)
}

func (a *EVMImplementation) GetUsersBookings() ([]protocol.VMBooking, error) {
	return a.Broker.GetUsersBookings()
}

func (a *EVMImplementation) GetMyAddress() *common.Address {
	return a.Broker.GetMyAddress()
}

func (a *EVMImplementation) GetMinerUrl(address *common.Address) (string, error) {
	return a.Broker.GetMinerUrl(address)
}

func (a *EVMImplementation) GetComplaintEvents(filter protocol.ComplaintFilterOpts) ([]protocol.ComplaintEvent, error) {
	return a.Broker.GetComplaintEvents(filter)
}

func (a *EVMImplementation) GetPaymentEvents(filter protocol.PaymentFilterOpts) ([]protocol.PaymentEvent, error) {
	return a.Broker.GetPaymentEvents(filter)
}

func (a *EVMImplementation) SetMinerUrlIfNeeded(newUrl string) error {
	return a.Broker.SetMinerUrlIfNeeded(newUrl)
}

func (a *EVMImplementation) GetTime() (uint64, error) {
	return a.Broker.GetTime()
}

func (a *EVMImplementation) GetMinersBookings() ([]protocol.VMBooking, error) {
	return a.Broker.GetMinersBookings()
}

func (a *EVMImplementation) GetStablecoinBalance() (uint64, uint64, error) {
	return a.Broker.GetStablecoinBalance()
}

// func (a *EVMImplementation) SetStablecoinAddress(address common.Address) error {
// 	return a.Broker.SetStablecoinAddress(address)
// }

func (a *EVMImplementation) GetStablecoinAddress() (common.Address, error) {
	return a.Broker.GetStablecoinAddress()
}

func (a *EVMImplementation) SetCommunityContract(address common.Address) error {
	return a.Broker.SetCommunityContract(address)
}

func (a *EVMImplementation) GetCommunityContract() (common.Address, error) {
	return a.Broker.GetCommunityContract()
}

func (a *EVMImplementation) SetCommunityFee(fee uint64) error {
	return a.Broker.SetCommunityFee(fee)
}

func (a *EVMImplementation) GetCommunityFee() (uint64, error) {
	return a.Broker.GetCommunityFee()
}

// func (a *EVMImplementation) AbortBooking(index uint64, abortType protocol.AbortType) error {
// 	return a.Broker.AbortBooking(index, abortType)
// }

// func (a *EVMImplementation) ClaimExpired(index uint64) error {
// 	return a.Broker.ClaimExpired(index)
// }

// func (a *EVMImplementation) ExtendBooking(index uint64, secs int) error {
// 	return a.Broker.ExtendBooking(index, secs)
// }

// func (a *EVMImplementation) DepositBalance() (float64, error) {
// 	return a.Broker.DepositBalance()
// }

// func (a *EVMImplementation) LockedBalance() (float64, error) {
// 	return a.Broker.LockedBalance()
// }
