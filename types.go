package protocol

import (
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"

	"github.com/ethereum/go-ethereum/common"
)

type Offer struct {
	VmTypeId      int
	PPS           int
	Availablility int
	Miner         common.Address
	Index         int
}

type VMBooking struct {
	VmTypeId   int
	PPS        int
	Miner      *common.Address
	Index      int
	User       *common.Address
	BookedTill int
}

type StableCoinIface interface {
	DepositCoin(amount int64) error
	WithdrawCoin(amount int64) error
	Balance() (int64, error)
	UserTokenBalance() (int64, error)
	UserAllowance() (int64, error)
}

type StableCoinSessionIface interface {
	DepositCoin(numTokens *big.Int) (*types.Transaction, error)
	WithdrawCoin(numTokens *big.Int) (*types.Transaction, error)
	UserBalance() (*big.Int, error)
	UserAllowance() (*big.Int, error)
	UserTokenBalance() (*big.Int, error)
}

type BrokerIface interface {
	DeployContracts() ([]string, error)
	AddOffer(offer Offer, callbackUrl string) error
	GetMyOffers() ([]Offer, error)
	UpdateOffer(offer Offer) error
	GetPrivateKey() *ecdsa.PrivateKey
	ContractAddress() common.Address
	GetMtlsHash(*common.Address) (string, error)
	GetBooking(index int) (*VMBooking, error)
	RegisterMtlsHashIfNeeded(mtlsHash string) error
	GetAvailableOffers(vmTypeId int) ([]Offer, error)
	BookVM(offerIndex, seconds int) error
	GetUsersBookings() ([]VMBooking, error)
	GetMyAddress() *common.Address
	GetMinerUrl(address *common.Address) (string, error)
	SetMinerUrlIfNeeded(newUrl string) error
	GetTime() (int, error)
	GetMinersBookings() ([]VMBooking, error)

	RegenerateSession() error
	GetStableCoinSession() StableCoinSessionIface
}

type BlockchainIface interface {
	BrokerIface
	StableCoinIface
}
