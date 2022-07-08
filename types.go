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

type BlockchainIface interface {
	DeployContracts() ([]string, error)
	AddOffer(offer Offer, callbackUrl string) error
	GetMyOffers() ([]Offer, error)
	UpdateOffer(offer Offer) error
	GetPrivateKey() *ecdsa.PrivateKey
	GetMtlsHash(*common.Address) (string, error)
	GetBooking(index int) (*VMBooking, error)
	RegisterMtlsHashIfNeeded(mtlsHash string) error
	GetAvailableOffers(vmTypeId int) ([]Offer, error)
	BookVM(offerIndex, seconds int) error
	GetUsersBookings() ([]VMBooking, error)
	GetMyAddress() *common.Address
	GetMinerUrl(address *common.Address) (string, error)
}
