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
	VmTypeId      uint64
	PPS           uint64
	Availablility uint64
	Miner         common.Address
	Index         uint64
}

type VMBooking struct {
	VmTypeId    uint64
	PPS         uint64
	Miner       *common.Address
	Index       uint64
	User        *common.Address
	BookedAt    uint64
	LastPayment uint64
}

type P2PCloudProtocolIface interface {
	//Offers
	AddOffer(offer Offer, callbackUrl string) error
	GetMyOffers() ([]Offer, error)
	UpdateOffer(offer Offer) error
	GetAvailableOffers(vmTypeId uint64) ([]Offer, error)
	RemoveOffer(id uint64) error

	//Balance
	GetStablecoinBalance() (uint64, uint64, error)
	DepositStablecoin(amount uint64) error
	WithdrawStablecoin(amount uint64) error

	//Booking
	BookVM(offerIndex uint64) error
	ClaimPayment(bookingIndex uint64) error

	//Wallet
	GetPrivateKey() *ecdsa.PrivateKey
	GetMyAddress() *common.Address

	//WIP
	// DepositCoin(coins float64) error
	// WithdrawCoin() error
	// Balance() (float64, error)
	// DepositBalance() (float64, error)
	// LockedBalance() (float64, error)
	// SetStablecoinAddress(address common.Address) error
	// GetStablecoinAddress() (common.Address, error)
	// UserAllowance() (float64, error)
	// UserTokenBalance() (float64, error)

	//Non-refactored yet
	ContractAddress() common.Address
	GetBooking(index uint64) (*VMBooking, error)
	GetUsersBookings() ([]VMBooking, error)
	GetMinerUrl(address *common.Address) (string, error)
	SetMinerUrlIfNeeded(newUrl string) error
	GetTime() (uint64, error)
	GetMinersBookings() ([]VMBooking, error)
	SetCommunityContract(address common.Address) error
	GetCommunityContract() (common.Address, error)
	SetCommunityFee(fee uint64) error
	GetCommunityFee() (uint64, error)
	// AbortBooking(index uint64, abortType AbortType) error
	// ClaimExpired(index uint64) error
	// ExtendBooking(index uint64, secs uint64) error
}
