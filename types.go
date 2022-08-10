package protocol

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/common"
)

type StopReason uint8

const (
	StopReasonNotNeeded StopReason = 0
	StopReasonComplain  StopReason = 1
)

type ComplainEvent struct {
	User   *common.Address
	Miner  *common.Address
	Reason StopReason
}

type PaymentEvent struct {
	User   *common.Address
	Miner  *common.Address
	Amount uint64
}
type PaymentFilterOpts struct {
	User       *common.Address
	Miner      *common.Address
	StartBlock uint64
	EndBlock   *uint64
}

type ComplainFilterOpts struct {
	User       *common.Address
	Miner      *common.Address
	StopReason *StopReason
	StartBlock uint64
	EndBlock   *uint64
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
	StopVM(index uint64, reason StopReason) error
	GetBooking(index uint64) (VMBooking, error)
	GetUsersBookings() ([]VMBooking, error)
	GetMinerUrl(address *common.Address) (string, error)
	SetMinerUrlIfNeeded(newUrl string) error
	GetMinersBookings() ([]VMBooking, error)

	//Wallet
	GetPrivateKey() *ecdsa.PrivateKey
	GetMyAddress() *common.Address

	//Admin
	SetCommunityContract(address common.Address) error
	GetCommunityContract() (common.Address, error)
	SetCommunityFee(fee uint64) error
	GetCommunityFee() (uint64, error)

	//Utility
	GetTime() (uint64, error)
	ContractAddress() common.Address

	//Events
	GetComplainEvents(filter ComplainFilterOpts) ([]ComplainEvent, error)
	GetPaymentEvents(filter PaymentFilterOpts) ([]PaymentEvent, error)
}
