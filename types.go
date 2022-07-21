package protocol

import (
	"crypto/ecdsa"
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
	DepositCoin(coins float64) error
	WithdrawCoin(coins float64) error
	Balance() (float64, error)
	UserTokenBalance() (float64, error)
	UserAllowance() (float64, error)
	SetStablecoinAddress(address common.Address) error
	GetStablecoinAddress() (common.Address, error)
	SetCommunityContract(address common.Address) error
	GetCommunityContract() (common.Address, error)
	SetCommunityFee(fee int64) error
	GetCommunityFee() (int64, error)
}

type TokenIface interface {
	BalanceOf(address common.Address) (float64, error)
	Transfer(to common.Address, coins float64) error
	Approve(to common.Address, coins float64) error
	Allowance(from, address common.Address) (float64, error)

	StartUp() error
}
