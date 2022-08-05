package protocol

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/common"
)

const (
	UnknownAbortType AbortType = 0
	ReportAbortType  AbortType = 1
	StopAbortType    AbortType = 2
)

type AbortType uint8

func (a AbortType) ToSolidityType() uint8 {
	return uint8(a)
}

type Offer struct {
	VmTypeId      int
	PPS           float64
	Availablility int
	Miner         common.Address
	Index         int
}

type VMBooking struct {
	VmTypeId   int
	PPS        float64
	Miner      *common.Address
	Index      int
	User       *common.Address
	BookedAt   int
	BookedTill int
}

type BrokerIface interface {
	DeployContracts(address common.Address) ([]string, error)
	AddOffer(offer Offer, callbackUrl string) error
	GetMyOffers() ([]Offer, error)
	UpdateOffer(offer Offer) error
	GetPrivateKey() *ecdsa.PrivateKey
	ContractAddress() common.Address
	GetBooking(index int) (*VMBooking, error)
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
	DepositBalance() (float64, error)
	LockedBalance() (float64, error)
	UserTokenBalance() (float64, error)
	UserAllowance() (float64, error)
	SetStablecoinAddress(address common.Address) error
	GetStablecoinAddress() (common.Address, error)
	SetCommunityContract(address common.Address) error
	GetCommunityContract() (common.Address, error)
	SetCommunityFee(fee int64) error
	GetCommunityFee() (int64, error)
	AbortBooking(index uint64, abortType AbortType) error
	ClaimExpired(index uint64) error
	ExtendBooking(index uint64, secs int) error
}

type TokenIface interface {
	BalanceOf(address common.Address) (float64, error)
	Transfer(to common.Address, coins float64) error
	Approve(to common.Address, coins float64) error
	Allowance(from, address common.Address) (float64, error)

	StartUp() error
}
