// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// BrokerBooking is an auto generated low-level Go binding around an user-defined struct.
type BrokerBooking struct {
	Index          uint64
	VmTypeId       uint64
	Miner          common.Address
	User           common.Address
	PricePerSecond uint64
	BookedAt       *big.Int
	LastPayment    *big.Int
	OfferIndex     uint64
}

// BrokerVMOffer is an auto generated low-level Go binding around an user-defined struct.
type BrokerVMOffer struct {
	Index             uint64
	Miner             common.Address
	PricePerSecond    uint64
	MachinesAvailable uint64
	VmTypeId          uint64
}

// BrokerMetaData contains all meta data concerning the Broker contract.
var BrokerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"reason\",\"type\":\"uint8\"}],\"name\":\"Complaint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Payment\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"pricePerSecond\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"vmTypeId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"machinesAvailable\",\"type\":\"uint64\"}],\"name\":\"AddOffer\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"offerIndex\",\"type\":\"uint64\"}],\"name\":\"BookVM\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"bookingId\",\"type\":\"uint64\"}],\"name\":\"ClaimPayment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"numTokens\",\"type\":\"uint256\"}],\"name\":\"DepositStablecoin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_miner\",\"type\":\"address\"}],\"name\":\"FindBookingsByMiner\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"vmTypeId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"pricePerSecond\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"bookedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastPayment\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"offerIndex\",\"type\":\"uint64\"}],\"internalType\":\"structBroker.Booking[]\",\"name\":\"filteredBookings\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"FindBookingsByUser\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"vmTypeId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"pricePerSecond\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"bookedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastPayment\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"offerIndex\",\"type\":\"uint64\"}],\"internalType\":\"structBroker.Booking[]\",\"name\":\"filteredBookings\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetAvailableOffers\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"pricePerSecond\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"machinesAvailable\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"vmTypeId\",\"type\":\"uint64\"}],\"internalType\":\"structBroker.VMOffer[]\",\"name\":\"filteredOffers\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_vmTypeId\",\"type\":\"uint64\"}],\"name\":\"GetAvailableOffersByType\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"pricePerSecond\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"machinesAvailable\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"vmTypeId\",\"type\":\"uint64\"}],\"internalType\":\"structBroker.VMOffer[]\",\"name\":\"filteredOffers\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"}],\"name\":\"GetBooking\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"vmTypeId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"pricePerSecond\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"bookedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastPayment\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"offerIndex\",\"type\":\"uint64\"}],\"internalType\":\"structBroker.Booking\",\"name\":\"booking\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"}],\"name\":\"GetMinersOffers\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"pricePerSecond\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"machinesAvailable\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"vmTypeId\",\"type\":\"uint64\"}],\"internalType\":\"structBroker.VMOffer[]\",\"name\":\"filteredOffers\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"GetStablecoinBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"GetUsersBookings\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"vmTypeId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"pricePerSecond\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"bookedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastPayment\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"offerIndex\",\"type\":\"uint64\"}],\"internalType\":\"structBroker.Booking[]\",\"name\":\"filteredBookings\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"offerIndex\",\"type\":\"uint64\"}],\"name\":\"RemoveOffer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SECONDS_IN_WEEK\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newCommunityAddress\",\"type\":\"address\"}],\"name\":\"SetCommunityContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"newStablecoinAddress\",\"type\":\"address\"}],\"name\":\"SetStablecoinAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"bookingId\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"reason\",\"type\":\"uint8\"}],\"name\":\"StopVM\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"offerIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"machinesAvailable\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"pps\",\"type\":\"uint64\"}],\"name\":\"UpdateOffer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"name\":\"WithdrawStablecoin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCommunityContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCommunityFee\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getMinerUrl\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStablecoinAddress\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"fee\",\"type\":\"uint64\"}],\"name\":\"setCommunityFee\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"url\",\"type\":\"bytes32\"}],\"name\":\"setMunerUrl\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040526000600160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506000600360006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506101f4600860146101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555034801561008f57600080fd5b5033600860006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550614eed806100e06000396000f3fe608060405234801561001057600080fd5b50600436106101a95760003560e01c80637f7a1915116100f9578063b65c4b9c11610097578063d3554ab511610071578063d3554ab51461052f578063e088b29e1461055f578063f3fe369c1461057d578063f44f11cd14610599576101a9565b8063b65c4b9c146104c5578063c2161a05146104e3578063c9b16fee14610513576101a9565b806390ad688b116100d357806390ad688b1461042b5780639a61a20e14610449578063a7bf757014610465578063ab49ec4514610495576101a9565b80637f7a1915146103af57806388317656146103df5780638f351016146103fb576101a9565b8063569e2e0f116101665780636f427a80116101405780636f427a80146103005780637559fe621461033057806375b7322e146103615780637bef4f5414610391576101a9565b8063569e2e0f146102965780636b620ed5146102b45780636d8351d6146102d0576101a9565b806318c5e502146101ae578063366445bf146101cc5780633fec2ddf146101fc578063534b3ccb146102185780635405b5e51461024857806354fd4d5014610278575b600080fd5b6101b66105b5565b6040516101c39190613c6c565b60405180910390f35b6101e660048036038101906101e19190613cea565b6105bd565b6040516101f39190613eaa565b60405180910390f35b61021660048036038101906102119190613ef8565b6109dc565b005b610232600480360381019061022d9190613cea565b610b38565b60405161023f9190613eaa565b60405180910390f35b610262600480360381019061025d9190613f4b565b610f57565b60405161026f919061408f565b60405180910390f35b610280611339565b60405161028d91906140c0565b60405180910390f35b61029e61133e565b6040516102ab91906140ea565b60405180910390f35b6102ce60048036038101906102c99190613cea565b611368565b005b6102ea60048036038101906102e59190614131565b61143c565b6040516102f79190614179565b60405180910390f35b61031a60048036038101906103159190614131565b6115b6565b6040516103279190614179565b60405180910390f35b61034a60048036038101906103459190613cea565b611790565b604051610358929190614194565b60405180910390f35b61037b60048036038101906103769190613f4b565b6117f6565b60405161038891906140c0565b60405180910390f35b610399611e2b565b6040516103a6919061421c565b60405180910390f35b6103c960048036038101906103c49190613cea565b611e55565b6040516103d69190613eaa565b60405180910390f35b6103f960048036038101906103f49190613f4b565b612274565b005b61041560048036038101906104109190613ef8565b6123ea565b60405161042291906140c0565b60405180910390f35b610433612621565b60405161044091906140c0565b60405180910390f35b610463600480360381019061045e9190614270565b61263f565b005b61047f600480360381019061047a9190613f4b565b61280b565b60405161048c9190614179565b60405180910390f35b6104af60048036038101906104aa9190613cea565b61291d565b6040516104bc91906142c9565b60405180910390f35b6104cd612966565b6040516104da919061408f565b60405180910390f35b6104fd60048036038101906104f89190613f4b565b612cea565b60405161050a9190614386565b60405180910390f35b61052d600480360381019061052891906143e0565b612eb4565b005b61054960048036038101906105449190613cea565b612f88565b604051610556919061408f565b60405180910390f35b61056761333a565b60405161057491906140c0565b60405180910390f35b61059760048036038101906105929190613f4b565b613341565b005b6105b360048036038101906105ae9190614439565b613405565b005b600042905090565b60606000600360009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff8111156105fb576105fa614466565b5b60405190808252806020026020018201604052801561063457816020015b610621613b4d565b8152602001906001900390816106195790505b509050600080600090505b600360009054906101000a900467ffffffffffffffff1667ffffffffffffffff168167ffffffffffffffff1610156108f4578473ffffffffffffffffffffffffffffffffffffffff16600260008367ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16036108e157600260008267ffffffffffffffff1667ffffffffffffffff168152602001908152602001600020604051806101000160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820160149054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff16815260200160038201548152602001600482015481526020016005820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681525050838367ffffffffffffffff16815181106108c6576108c5614495565b5b60200260200101819052506001826108de91906144f3565b91505b80806108ec90614531565b91505061063f565b508067ffffffffffffffff1667ffffffffffffffff81111561091957610918614466565b5b60405190808252806020026020018201604052801561095257816020015b61093f613b4d565b8152602001906001900390816109375790505b50925060005b8167ffffffffffffffff168167ffffffffffffffff1610156109d457828167ffffffffffffffff168151811061099157610990614495565b5b6020026020010151848267ffffffffffffffff16815181106109b6576109b5614495565b5b602002602001018190525080806109cc90614531565b915050610958565b505050919050565b3373ffffffffffffffffffffffffffffffffffffffff166000808567ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060000160089054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610a93576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a8a906145e4565b60405180910390fd5b816000808567ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060010160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550806000808567ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060010160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550505050565b60606000600360009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff811115610b7657610b75614466565b5b604051908082528060200260200182016040528015610baf57816020015b610b9c613b4d565b815260200190600190039081610b945790505b509050600080600090505b600360009054906101000a900467ffffffffffffffff1667ffffffffffffffff168167ffffffffffffffff161015610e6f578473ffffffffffffffffffffffffffffffffffffffff16600260008367ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603610e5c57600260008267ffffffffffffffff1667ffffffffffffffff168152602001908152602001600020604051806101000160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820160149054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff16815260200160038201548152602001600482015481526020016005820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681525050838367ffffffffffffffff1681518110610e4157610e40614495565b5b6020026020010181905250600182610e5991906144f3565b91505b8080610e6790614531565b915050610bba565b508067ffffffffffffffff1667ffffffffffffffff811115610e9457610e93614466565b5b604051908082528060200260200182016040528015610ecd57816020015b610eba613b4d565b815260200190600190039081610eb25790505b50925060005b8167ffffffffffffffff168167ffffffffffffffff161015610f4f57828167ffffffffffffffff1681518110610f0c57610f0b614495565b5b6020026020010151848267ffffffffffffffff1681518110610f3157610f30614495565b5b60200260200101819052508080610f4790614531565b915050610ed3565b505050919050565b60606000600160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff811115610f9557610f94614466565b5b604051908082528060200260200182016040528015610fce57816020015b610fbb613be6565b815260200190600190039081610fb35790505b509050600080600090505b600160009054906101000a900467ffffffffffffffff1667ffffffffffffffff168167ffffffffffffffff161015611251578467ffffffffffffffff166000808367ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060010160109054906101000a900467ffffffffffffffff1667ffffffffffffffff161480156110b2575060008060008367ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060010160089054906101000a900467ffffffffffffffff1667ffffffffffffffff16115b1561123e576000808267ffffffffffffffff1667ffffffffffffffff1681526020019081526020016000206040518060a00160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016001820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681525050838367ffffffffffffffff168151811061122357611222614495565b5b602002602001018190525060018261123b91906144f3565b91505b808061124990614531565b915050610fd9565b508067ffffffffffffffff1667ffffffffffffffff81111561127657611275614466565b5b6040519080825280602002602001820160405280156112af57816020015b61129c613be6565b8152602001906001900390816112945790505b50925060005b8167ffffffffffffffff168167ffffffffffffffff16101561133157828167ffffffffffffffff16815181106112ee576112ed614495565b5b6020026020010151848267ffffffffffffffff168151811061131357611312614495565b5b6020026020010181905250808061132990614531565b9150506112b5565b505050919050565b600181565b6000600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146113f8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016113ef90614676565b60405180910390fd5b80600860006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6000600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd3330856040518463ffffffff1660e01b815260040161149d93929190614696565b6020604051808303816000875af11580156114bc573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114e091906146f9565b61151f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161151690614772565b60405180910390fd5b81600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461156a9190614792565b600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555060019050919050565b6000806115c23361344c565b600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461160c91906147e8565b905082811015611651576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161164890614868565b60405180910390fd5b600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb33856040518363ffffffff1660e01b81526004016116ae929190614888565b6020604051808303816000875af11580156116cd573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116f191906146f9565b611730576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611727906148fd565b60405180910390fd5b82600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461177f91906147e8565b925050819055506001915050919050565b600080600061179e8461344c565b905080600560008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546117eb91906147e8565b819250925050915091565b6000806000808467ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060010160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1611611881576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161187890614969565b60405180910390fd5b600062093a806000808567ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060010160009054906101000a900467ffffffffffffffff166118cf9190614989565b67ffffffffffffffff166118e23361344c565b6118ec9190614792565b9050600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054811115611970576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161196790614a63565b60405180910390fd5b6000604051806101000160405280600360009054906101000a900467ffffffffffffffff1667ffffffffffffffff1681526020016000808767ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060010160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1681526020016000808767ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060000160089054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020013373ffffffffffffffffffffffffffffffffffffffff1681526020016000808767ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060010160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1681526020014281526020014281526020018567ffffffffffffffff1681525090508060026000600360009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060208201518160000160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060408201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060808201518160020160146101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060a0820151816003015560c0820151816004015560e08201518160050160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055509050506003600081819054906101000a900467ffffffffffffffff1680929190611ca090614531565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550506000808567ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060010160009054906101000a900467ffffffffffffffff16600660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282829054906101000a900467ffffffffffffffff16611d6691906144f3565b92506101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060016000808667ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060010160088282829054906101000a900467ffffffffffffffff16611dd99190614a83565b92506101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506001600360009054906101000a900467ffffffffffffffff16611e229190614a83565b92505050919050565b6000600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60606000600360009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff811115611e9357611e92614466565b5b604051908082528060200260200182016040528015611ecc57816020015b611eb9613b4d565b815260200190600190039081611eb15790505b509050600080600090505b600360009054906101000a900467ffffffffffffffff1667ffffffffffffffff168167ffffffffffffffff16101561218c578473ffffffffffffffffffffffffffffffffffffffff16600260008367ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff160361217957600260008267ffffffffffffffff1667ffffffffffffffff168152602001908152602001600020604051806101000160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820160149054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff16815260200160038201548152602001600482015481526020016005820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681525050838367ffffffffffffffff168151811061215e5761215d614495565b5b602002602001018190525060018261217691906144f3565b91505b808061218490614531565b915050611ed7565b508067ffffffffffffffff1667ffffffffffffffff8111156121b1576121b0614466565b5b6040519080825280602002602001820160405280156121ea57816020015b6121d7613b4d565b8152602001906001900390816121cf5790505b50925060005b8167ffffffffffffffff168167ffffffffffffffff16101561226c57828167ffffffffffffffff168151811061222957612228614495565b5b6020026020010151848267ffffffffffffffff168151811061224e5761224d614495565b5b6020026020010181905250808061226490614531565b9150506121f0565b505050919050565b3373ffffffffffffffffffffffffffffffffffffffff166000808367ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060000160089054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461232b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161232290614b29565b60405180910390fd5b6000808267ffffffffffffffff1667ffffffffffffffff168152602001908152602001600020600080820160006101000a81549067ffffffffffffffff02191690556000820160086101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556001820160006101000a81549067ffffffffffffffff02191690556001820160086101000a81549067ffffffffffffffff02191690556001820160106101000a81549067ffffffffffffffff0219169055505050565b60006040518060a00160405280600160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1681526020013373ffffffffffffffffffffffffffffffffffffffff1681526020018567ffffffffffffffff1681526020018367ffffffffffffffff1681526020018467ffffffffffffffff16815250600080600160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060208201518160000160086101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160010160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060608201518160010160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060808201518160010160106101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055509050506001600081819054906101000a900467ffffffffffffffff16809291906125cf90614531565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055505060018060009054906101000a900467ffffffffffffffff166126189190614a83565b90509392505050565b6000600860149054906101000a900467ffffffffffffffff16905090565b3373ffffffffffffffffffffffffffffffffffffffff16600260008467ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146126f7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016126ee90614b95565b60405180910390fd5b60008160ff16146127f5578060ff16600260008467ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16600260008567ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f1db97daa25845ee925bd84e6cf559402b161e27653e3f3187c1859f98634c1c560405160405180910390a45b6127fe826134c1565b612807826139a5565b5050565b60006127108267ffffffffffffffff161061285b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161285290614c27565b60405180910390fd5b600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146128eb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016128e290614cb9565b60405180910390fd5b81600860146101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060009050919050565b6000600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b60606000600160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff8111156129a4576129a3614466565b5b6040519080825280602002602001820160405280156129dd57816020015b6129ca613be6565b8152602001906001900390816129c25790505b509050600080600090505b600160009054906101000a900467ffffffffffffffff1667ffffffffffffffff168167ffffffffffffffff161015612c045760008060008367ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060010160089054906101000a900467ffffffffffffffff1667ffffffffffffffff161115612bf1576000808267ffffffffffffffff1667ffffffffffffffff1681526020019081526020016000206040518060a00160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016001820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681525050838367ffffffffffffffff1681518110612bd657612bd5614495565b5b6020026020010181905250600182612bee91906144f3565b91505b8080612bfc90614531565b9150506129e8565b508067ffffffffffffffff1667ffffffffffffffff811115612c2957612c28614466565b5b604051908082528060200260200182016040528015612c6257816020015b612c4f613be6565b815260200190600190039081612c475790505b50925060005b8167ffffffffffffffff168167ffffffffffffffff161015612ce457828167ffffffffffffffff1681518110612ca157612ca0614495565b5b6020026020010151848267ffffffffffffffff1681518110612cc657612cc5614495565b5b60200260200101819052508080612cdc90614531565b915050612c68565b50505090565b612cf2613b4d565b600260008367ffffffffffffffff1667ffffffffffffffff168152602001908152602001600020604051806101000160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820160149054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff16815260200160038201548152602001600482015481526020016005820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff16815250509050919050565b600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614612f44576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612f3b90614d4b565b60405180910390fd5b80600760006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60606000600160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff811115612fc657612fc5614466565b5b604051908082528060200260200182016040528015612fff57816020015b612fec613be6565b815260200190600190039081612fe45790505b509050600080600090505b600160009054906101000a900467ffffffffffffffff1667ffffffffffffffff168167ffffffffffffffff161015613252578473ffffffffffffffffffffffffffffffffffffffff166000808367ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060000160089054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff160361323f576000808267ffffffffffffffff1667ffffffffffffffff1681526020019081526020016000206040518060a00160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016001820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681525050838367ffffffffffffffff168151811061322457613223614495565b5b602002602001018190525060018261323c91906144f3565b91505b808061324a90614531565b91505061300a565b508067ffffffffffffffff1667ffffffffffffffff81111561327757613276614466565b5b6040519080825280602002602001820160405280156132b057816020015b61329d613be6565b8152602001906001900390816132955790505b50925060005b8167ffffffffffffffff168167ffffffffffffffff16101561333257828167ffffffffffffffff16815181106132ef576132ee614495565b5b6020026020010151848267ffffffffffffffff168151811061331457613313614495565b5b6020026020010181905250808061332a90614531565b9150506132b6565b505050919050565b62093a8081565b3373ffffffffffffffffffffffffffffffffffffffff16600260008367ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146133f9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016133f090614ddd565b60405180910390fd5b613402816134c1565b50565b80600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555050565b600062093a80600660008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900467ffffffffffffffff166134b09190614989565b67ffffffffffffffff169050919050565b6000600260008367ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060040154426134f991906147e8565b90506000600260008467ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060020160149054906101000a900467ffffffffffffffff1667ffffffffffffffff16826135519190614dfd565b90508060056000600260008767ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205410156136715760056000600260008667ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490505b6000612710600860149054906101000a900467ffffffffffffffff1667ffffffffffffffff16836136a29190614dfd565b6136ac9190614e86565b9050600081836136bc91906147e8565b905042600260008767ffffffffffffffff1667ffffffffffffffff168152602001908152602001600020600401819055508160056000600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461375e9190614792565b925050819055508060056000600260008967ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546137fe9190614792565b925050819055508260056000600260008967ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461389e91906147e8565b92505081905550600260008667ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16600260008767ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8ed7db30df47fbd87811a5e8a95a94838f0c0241263d9a1865d6a2a3e10516e2836040516139969190613c6c565b60405180910390a35050505050565b6001600080600260008567ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060050160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060010160088282829054906101000a900467ffffffffffffffff16613a3091906144f3565b92506101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550600260008267ffffffffffffffff1667ffffffffffffffff168152602001908152602001600020600080820160006101000a81549067ffffffffffffffff02191690556000820160086101000a81549067ffffffffffffffff02191690556001820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556002820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556002820160146101000a81549067ffffffffffffffff0219169055600382016000905560048201600090556005820160006101000a81549067ffffffffffffffff0219169055505050565b604051806101000160405280600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600067ffffffffffffffff1681526020016000815260200160008152602001600067ffffffffffffffff1681525090565b6040518060a00160405280600067ffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff1681525090565b6000819050919050565b613c6681613c53565b82525050565b6000602082019050613c816000830184613c5d565b92915050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000613cb782613c8c565b9050919050565b613cc781613cac565b8114613cd257600080fd5b50565b600081359050613ce481613cbe565b92915050565b600060208284031215613d0057613cff613c87565b5b6000613d0e84828501613cd5565b91505092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b600067ffffffffffffffff82169050919050565b613d6081613d43565b82525050565b613d6f81613cac565b82525050565b613d7e81613c53565b82525050565b61010082016000820151613d9b6000850182613d57565b506020820151613dae6020850182613d57565b506040820151613dc16040850182613d66565b506060820151613dd46060850182613d66565b506080820151613de76080850182613d57565b5060a0820151613dfa60a0850182613d75565b5060c0820151613e0d60c0850182613d75565b5060e0820151613e2060e0850182613d57565b50505050565b6000613e328383613d84565b6101008301905092915050565b6000602082019050919050565b6000613e5782613d17565b613e618185613d22565b9350613e6c83613d33565b8060005b83811015613e9d578151613e848882613e26565b9750613e8f83613e3f565b925050600181019050613e70565b5085935050505092915050565b60006020820190508181036000830152613ec48184613e4c565b905092915050565b613ed581613d43565b8114613ee057600080fd5b50565b600081359050613ef281613ecc565b92915050565b600080600060608486031215613f1157613f10613c87565b5b6000613f1f86828701613ee3565b9350506020613f3086828701613ee3565b9250506040613f4186828701613ee3565b9150509250925092565b600060208284031215613f6157613f60613c87565b5b6000613f6f84828501613ee3565b91505092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b60a082016000820151613fba6000850182613d57565b506020820151613fcd6020850182613d66565b506040820151613fe06040850182613d57565b506060820151613ff36060850182613d57565b5060808201516140066080850182613d57565b50505050565b60006140188383613fa4565b60a08301905092915050565b6000602082019050919050565b600061403c82613f78565b6140468185613f83565b935061405183613f94565b8060005b83811015614082578151614069888261400c565b975061407483614024565b925050600181019050614055565b5085935050505092915050565b600060208201905081810360008301526140a98184614031565b905092915050565b6140ba81613d43565b82525050565b60006020820190506140d560008301846140b1565b92915050565b6140e481613cac565b82525050565b60006020820190506140ff60008301846140db565b92915050565b61410e81613c53565b811461411957600080fd5b50565b60008135905061412b81614105565b92915050565b60006020828403121561414757614146613c87565b5b60006141558482850161411c565b91505092915050565b60008115159050919050565b6141738161415e565b82525050565b600060208201905061418e600083018461416a565b92915050565b60006040820190506141a96000830185613c5d565b6141b66020830184613c5d565b9392505050565b6000819050919050565b60006141e26141dd6141d884613c8c565b6141bd565b613c8c565b9050919050565b60006141f4826141c7565b9050919050565b6000614206826141e9565b9050919050565b614216816141fb565b82525050565b6000602082019050614231600083018461420d565b92915050565b600060ff82169050919050565b61424d81614237565b811461425857600080fd5b50565b60008135905061426a81614244565b92915050565b6000806040838503121561428757614286613c87565b5b600061429585828601613ee3565b92505060206142a68582860161425b565b9150509250929050565b6000819050919050565b6142c3816142b0565b82525050565b60006020820190506142de60008301846142ba565b92915050565b610100820160008201516142fb6000850182613d57565b50602082015161430e6020850182613d57565b5060408201516143216040850182613d66565b5060608201516143346060850182613d66565b5060808201516143476080850182613d57565b5060a082015161435a60a0850182613d75565b5060c082015161436d60c0850182613d75565b5060e082015161438060e0850182613d57565b50505050565b60006101008201905061439c60008301846142e4565b92915050565b60006143ad82613cac565b9050919050565b6143bd816143a2565b81146143c857600080fd5b50565b6000813590506143da816143b4565b92915050565b6000602082840312156143f6576143f5613c87565b5b6000614404848285016143cb565b91505092915050565b614416816142b0565b811461442157600080fd5b50565b6000813590506144338161440d565b92915050565b60006020828403121561444f5761444e613c87565b5b600061445d84828501614424565b91505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006144fe82613d43565b915061450983613d43565b92508267ffffffffffffffff03821115614526576145256144c4565b5b828201905092915050565b600061453c82613d43565b915067ffffffffffffffff8203614556576145556144c4565b5b600182019050919050565b600082825260208201905092915050565b7f4f6e6c7920746865206f776e65722063616e2075706461746520616e206f666660008201527f6572000000000000000000000000000000000000000000000000000000000000602082015250565b60006145ce602283614561565b91506145d982614572565b604082019050919050565b600060208201905081810360008301526145fd816145c1565b9050919050565b7f6f6e6c7920636f6d6d756e69747920636f6e74726163742063616e207365742060008201527f6e657720636f6d6d756e69747920636f6e747261637400000000000000000000602082015250565b6000614660603683614561565b915061466b82614604565b604082019050919050565b6000602082019050818103600083015261468f81614653565b9050919050565b60006060820190506146ab60008301866140db565b6146b860208301856140db565b6146c56040830184613c5d565b949350505050565b6146d68161415e565b81146146e157600080fd5b50565b6000815190506146f3816146cd565b92915050565b60006020828403121561470f5761470e613c87565b5b600061471d848285016146e4565b91505092915050565b7f4661696c656420746f207472616e7366657220746f6b656e7300000000000000600082015250565b600061475c601983614561565b915061476782614726565b602082019050919050565b6000602082019050818103600083015261478b8161474f565b9050919050565b600061479d82613c53565b91506147a883613c53565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff038211156147dd576147dc6144c4565b5b828201905092915050565b60006147f382613c53565b91506147fe83613c53565b925082821015614811576148106144c4565b5b828203905092915050565b7f4e6f7420656e6f7567682062616c616e636520746f2077697468647261770000600082015250565b6000614852601e83614561565b915061485d8261481c565b602082019050919050565b6000602082019050818103600083015261488181614845565b9050919050565b600060408201905061489d60008301856140db565b6148aa6020830184613c5d565b9392505050565b7f4552433230207472616e73666572206661696c65640000000000000000000000600082015250565b60006148e7601583614561565b91506148f2826148b1565b602082019050919050565b60006020820190508181036000830152614916816148da565b9050919050565b7f4e6f206d616368696e657320617661696c61626c650000000000000000000000600082015250565b6000614953601583614561565b915061495e8261491d565b602082019050919050565b6000602082019050818103600083015261498281614946565b9050919050565b600061499482613d43565b915061499f83613d43565b92508167ffffffffffffffff04831182151516156149c0576149bf6144c4565b5b828202905092915050565b7f596f7520646f6e2774206861766520656e6f7567682062616c616e636520746f60008201527f2070617920666f72207468697320616e6420616c6c206f7468657220564d732060208201527f666f722037206461797320000000000000000000000000000000000000000000604082015250565b6000614a4d604b83614561565b9150614a58826149cb565b606082019050919050565b60006020820190508181036000830152614a7c81614a40565b9050919050565b6000614a8e82613d43565b9150614a9983613d43565b925082821015614aac57614aab6144c4565b5b828203905092915050565b7f4f6e6c7920746865206f776e65722063616e2072656d6f766520616e206f666660008201527f6572000000000000000000000000000000000000000000000000000000000000602082015250565b6000614b13602283614561565b9150614b1e82614ab7565b604082019050919050565b60006020820190508181036000830152614b4281614b06565b9050919050565b7f4f6e6c792074686520757365722063616e2073746f70206120564d0000000000600082015250565b6000614b7f601b83614561565b9150614b8a82614b49565b602082019050919050565b60006020820190508181036000830152614bae81614b72565b9050919050565b7f636f6d6d756e697479206665652073686f756c6420626520696e2072616e676560008201527f206f662030202830252920746f20313030303020283130302529000000000000602082015250565b6000614c11603a83614561565b9150614c1c82614bb5565b604082019050919050565b60006020820190508181036000830152614c4081614c04565b9050919050565b7f6f6e6c7920636f6d6d756e69747920636f6e74726163742063616e207365742060008201527f636f6d6d756e6974792066656500000000000000000000000000000000000000602082015250565b6000614ca3602d83614561565b9150614cae82614c47565b604082019050919050565b60006020820190508181036000830152614cd281614c96565b9050919050565b7f6f6e6c7920636f6d6d756e69747920636f6e74726163742063616e207365742060008201527f737461626c65636f696e00000000000000000000000000000000000000000000602082015250565b6000614d35602a83614561565b9150614d4082614cd9565b604082019050919050565b60006020820190508181036000830152614d6481614d28565b9050919050565b7f4f6e6c7920746865206d696e65722063616e20636c61696d2061207061796d6560008201527f6e74000000000000000000000000000000000000000000000000000000000000602082015250565b6000614dc7602283614561565b9150614dd282614d6b565b604082019050919050565b60006020820190508181036000830152614df681614dba565b9050919050565b6000614e0882613c53565b9150614e1383613c53565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615614e4c57614e4b6144c4565b5b828202905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b6000614e9182613c53565b9150614e9c83613c53565b925082614eac57614eab614e57565b5b82820490509291505056fea264697066735822122065368b249a97db53e1f3f3b30097638be5673a2e5a53ca63a11df0a99d2c626b64736f6c634300080f0033",
}

// BrokerABI is the input ABI used to generate the binding from.
// Deprecated: Use BrokerMetaData.ABI instead.
var BrokerABI = BrokerMetaData.ABI

// BrokerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BrokerMetaData.Bin instead.
var BrokerBin = BrokerMetaData.Bin

// DeployBroker deploys a new Ethereum contract, binding an instance of Broker to it.
func DeployBroker(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Broker, error) {
	parsed, err := BrokerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BrokerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Broker{BrokerCaller: BrokerCaller{contract: contract}, BrokerTransactor: BrokerTransactor{contract: contract}, BrokerFilterer: BrokerFilterer{contract: contract}}, nil
}

// Broker is an auto generated Go binding around an Ethereum contract.
type Broker struct {
	BrokerCaller     // Read-only binding to the contract
	BrokerTransactor // Write-only binding to the contract
	BrokerFilterer   // Log filterer for contract events
}

// BrokerCaller is an auto generated read-only Go binding around an Ethereum contract.
type BrokerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BrokerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BrokerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BrokerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BrokerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BrokerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BrokerSession struct {
	Contract     *Broker           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BrokerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BrokerCallerSession struct {
	Contract *BrokerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BrokerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BrokerTransactorSession struct {
	Contract     *BrokerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BrokerRaw is an auto generated low-level Go binding around an Ethereum contract.
type BrokerRaw struct {
	Contract *Broker // Generic contract binding to access the raw methods on
}

// BrokerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BrokerCallerRaw struct {
	Contract *BrokerCaller // Generic read-only contract binding to access the raw methods on
}

// BrokerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BrokerTransactorRaw struct {
	Contract *BrokerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBroker creates a new instance of Broker, bound to a specific deployed contract.
func NewBroker(address common.Address, backend bind.ContractBackend) (*Broker, error) {
	contract, err := bindBroker(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Broker{BrokerCaller: BrokerCaller{contract: contract}, BrokerTransactor: BrokerTransactor{contract: contract}, BrokerFilterer: BrokerFilterer{contract: contract}}, nil
}

// NewBrokerCaller creates a new read-only instance of Broker, bound to a specific deployed contract.
func NewBrokerCaller(address common.Address, caller bind.ContractCaller) (*BrokerCaller, error) {
	contract, err := bindBroker(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BrokerCaller{contract: contract}, nil
}

// NewBrokerTransactor creates a new write-only instance of Broker, bound to a specific deployed contract.
func NewBrokerTransactor(address common.Address, transactor bind.ContractTransactor) (*BrokerTransactor, error) {
	contract, err := bindBroker(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BrokerTransactor{contract: contract}, nil
}

// NewBrokerFilterer creates a new log filterer instance of Broker, bound to a specific deployed contract.
func NewBrokerFilterer(address common.Address, filterer bind.ContractFilterer) (*BrokerFilterer, error) {
	contract, err := bindBroker(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BrokerFilterer{contract: contract}, nil
}

// bindBroker binds a generic wrapper to an already deployed contract.
func bindBroker(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BrokerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Broker *BrokerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Broker.Contract.BrokerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Broker *BrokerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Broker.Contract.BrokerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Broker *BrokerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Broker.Contract.BrokerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Broker *BrokerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Broker.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Broker *BrokerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Broker.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Broker *BrokerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Broker.Contract.contract.Transact(opts, method, params...)
}

// FindBookingsByMiner is a free data retrieval call binding the contract method 0x366445bf.
//
// Solidity: function FindBookingsByMiner(address _miner) view returns((uint64,uint64,address,address,uint64,uint256,uint256,uint64)[] filteredBookings)
func (_Broker *BrokerCaller) FindBookingsByMiner(opts *bind.CallOpts, _miner common.Address) ([]BrokerBooking, error) {
	var out []interface{}
	err := _Broker.contract.Call(opts, &out, "FindBookingsByMiner", _miner)

	if err != nil {
		return *new([]BrokerBooking), err
	}

	out0 := *abi.ConvertType(out[0], new([]BrokerBooking)).(*[]BrokerBooking)

	return out0, err

}

// FindBookingsByMiner is a free data retrieval call binding the contract method 0x366445bf.
//
// Solidity: function FindBookingsByMiner(address _miner) view returns((uint64,uint64,address,address,uint64,uint256,uint256,uint64)[] filteredBookings)
func (_Broker *BrokerSession) FindBookingsByMiner(_miner common.Address) ([]BrokerBooking, error) {
	return _Broker.Contract.FindBookingsByMiner(&_Broker.CallOpts, _miner)
}

// FindBookingsByMiner is a free data retrieval call binding the contract method 0x366445bf.
//
// Solidity: function FindBookingsByMiner(address _miner) view returns((uint64,uint64,address,address,uint64,uint256,uint256,uint64)[] filteredBookings)
func (_Broker *BrokerCallerSession) FindBookingsByMiner(_miner common.Address) ([]BrokerBooking, error) {
	return _Broker.Contract.FindBookingsByMiner(&_Broker.CallOpts, _miner)
}

// FindBookingsByUser is a free data retrieval call binding the contract method 0x534b3ccb.
//
// Solidity: function FindBookingsByUser(address _owner) view returns((uint64,uint64,address,address,uint64,uint256,uint256,uint64)[] filteredBookings)
func (_Broker *BrokerCaller) FindBookingsByUser(opts *bind.CallOpts, _owner common.Address) ([]BrokerBooking, error) {
	var out []interface{}
	err := _Broker.contract.Call(opts, &out, "FindBookingsByUser", _owner)

	if err != nil {
		return *new([]BrokerBooking), err
	}

	out0 := *abi.ConvertType(out[0], new([]BrokerBooking)).(*[]BrokerBooking)

	return out0, err

}

// FindBookingsByUser is a free data retrieval call binding the contract method 0x534b3ccb.
//
// Solidity: function FindBookingsByUser(address _owner) view returns((uint64,uint64,address,address,uint64,uint256,uint256,uint64)[] filteredBookings)
func (_Broker *BrokerSession) FindBookingsByUser(_owner common.Address) ([]BrokerBooking, error) {
	return _Broker.Contract.FindBookingsByUser(&_Broker.CallOpts, _owner)
}

// FindBookingsByUser is a free data retrieval call binding the contract method 0x534b3ccb.
//
// Solidity: function FindBookingsByUser(address _owner) view returns((uint64,uint64,address,address,uint64,uint256,uint256,uint64)[] filteredBookings)
func (_Broker *BrokerCallerSession) FindBookingsByUser(_owner common.Address) ([]BrokerBooking, error) {
	return _Broker.Contract.FindBookingsByUser(&_Broker.CallOpts, _owner)
}

// GetAvailableOffers is a free data retrieval call binding the contract method 0xb65c4b9c.
//
// Solidity: function GetAvailableOffers() view returns((uint64,address,uint64,uint64,uint64)[] filteredOffers)
func (_Broker *BrokerCaller) GetAvailableOffers(opts *bind.CallOpts) ([]BrokerVMOffer, error) {
	var out []interface{}
	err := _Broker.contract.Call(opts, &out, "GetAvailableOffers")

	if err != nil {
		return *new([]BrokerVMOffer), err
	}

	out0 := *abi.ConvertType(out[0], new([]BrokerVMOffer)).(*[]BrokerVMOffer)

	return out0, err

}

// GetAvailableOffers is a free data retrieval call binding the contract method 0xb65c4b9c.
//
// Solidity: function GetAvailableOffers() view returns((uint64,address,uint64,uint64,uint64)[] filteredOffers)
func (_Broker *BrokerSession) GetAvailableOffers() ([]BrokerVMOffer, error) {
	return _Broker.Contract.GetAvailableOffers(&_Broker.CallOpts)
}

// GetAvailableOffers is a free data retrieval call binding the contract method 0xb65c4b9c.
//
// Solidity: function GetAvailableOffers() view returns((uint64,address,uint64,uint64,uint64)[] filteredOffers)
func (_Broker *BrokerCallerSession) GetAvailableOffers() ([]BrokerVMOffer, error) {
	return _Broker.Contract.GetAvailableOffers(&_Broker.CallOpts)
}

// GetAvailableOffersByType is a free data retrieval call binding the contract method 0x5405b5e5.
//
// Solidity: function GetAvailableOffersByType(uint64 _vmTypeId) view returns((uint64,address,uint64,uint64,uint64)[] filteredOffers)
func (_Broker *BrokerCaller) GetAvailableOffersByType(opts *bind.CallOpts, _vmTypeId uint64) ([]BrokerVMOffer, error) {
	var out []interface{}
	err := _Broker.contract.Call(opts, &out, "GetAvailableOffersByType", _vmTypeId)

	if err != nil {
		return *new([]BrokerVMOffer), err
	}

	out0 := *abi.ConvertType(out[0], new([]BrokerVMOffer)).(*[]BrokerVMOffer)

	return out0, err

}

// GetAvailableOffersByType is a free data retrieval call binding the contract method 0x5405b5e5.
//
// Solidity: function GetAvailableOffersByType(uint64 _vmTypeId) view returns((uint64,address,uint64,uint64,uint64)[] filteredOffers)
func (_Broker *BrokerSession) GetAvailableOffersByType(_vmTypeId uint64) ([]BrokerVMOffer, error) {
	return _Broker.Contract.GetAvailableOffersByType(&_Broker.CallOpts, _vmTypeId)
}

// GetAvailableOffersByType is a free data retrieval call binding the contract method 0x5405b5e5.
//
// Solidity: function GetAvailableOffersByType(uint64 _vmTypeId) view returns((uint64,address,uint64,uint64,uint64)[] filteredOffers)
func (_Broker *BrokerCallerSession) GetAvailableOffersByType(_vmTypeId uint64) ([]BrokerVMOffer, error) {
	return _Broker.Contract.GetAvailableOffersByType(&_Broker.CallOpts, _vmTypeId)
}

// GetBooking is a free data retrieval call binding the contract method 0xc2161a05.
//
// Solidity: function GetBooking(uint64 index) view returns((uint64,uint64,address,address,uint64,uint256,uint256,uint64) booking)
func (_Broker *BrokerCaller) GetBooking(opts *bind.CallOpts, index uint64) (BrokerBooking, error) {
	var out []interface{}
	err := _Broker.contract.Call(opts, &out, "GetBooking", index)

	if err != nil {
		return *new(BrokerBooking), err
	}

	out0 := *abi.ConvertType(out[0], new(BrokerBooking)).(*BrokerBooking)

	return out0, err

}

// GetBooking is a free data retrieval call binding the contract method 0xc2161a05.
//
// Solidity: function GetBooking(uint64 index) view returns((uint64,uint64,address,address,uint64,uint256,uint256,uint64) booking)
func (_Broker *BrokerSession) GetBooking(index uint64) (BrokerBooking, error) {
	return _Broker.Contract.GetBooking(&_Broker.CallOpts, index)
}

// GetBooking is a free data retrieval call binding the contract method 0xc2161a05.
//
// Solidity: function GetBooking(uint64 index) view returns((uint64,uint64,address,address,uint64,uint256,uint256,uint64) booking)
func (_Broker *BrokerCallerSession) GetBooking(index uint64) (BrokerBooking, error) {
	return _Broker.Contract.GetBooking(&_Broker.CallOpts, index)
}

// GetMinersOffers is a free data retrieval call binding the contract method 0xd3554ab5.
//
// Solidity: function GetMinersOffers(address miner) view returns((uint64,address,uint64,uint64,uint64)[] filteredOffers)
func (_Broker *BrokerCaller) GetMinersOffers(opts *bind.CallOpts, miner common.Address) ([]BrokerVMOffer, error) {
	var out []interface{}
	err := _Broker.contract.Call(opts, &out, "GetMinersOffers", miner)

	if err != nil {
		return *new([]BrokerVMOffer), err
	}

	out0 := *abi.ConvertType(out[0], new([]BrokerVMOffer)).(*[]BrokerVMOffer)

	return out0, err

}

// GetMinersOffers is a free data retrieval call binding the contract method 0xd3554ab5.
//
// Solidity: function GetMinersOffers(address miner) view returns((uint64,address,uint64,uint64,uint64)[] filteredOffers)
func (_Broker *BrokerSession) GetMinersOffers(miner common.Address) ([]BrokerVMOffer, error) {
	return _Broker.Contract.GetMinersOffers(&_Broker.CallOpts, miner)
}

// GetMinersOffers is a free data retrieval call binding the contract method 0xd3554ab5.
//
// Solidity: function GetMinersOffers(address miner) view returns((uint64,address,uint64,uint64,uint64)[] filteredOffers)
func (_Broker *BrokerCallerSession) GetMinersOffers(miner common.Address) ([]BrokerVMOffer, error) {
	return _Broker.Contract.GetMinersOffers(&_Broker.CallOpts, miner)
}

// GetStablecoinBalance is a free data retrieval call binding the contract method 0x7559fe62.
//
// Solidity: function GetStablecoinBalance(address user) view returns(uint256, uint256)
func (_Broker *BrokerCaller) GetStablecoinBalance(opts *bind.CallOpts, user common.Address) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Broker.contract.Call(opts, &out, "GetStablecoinBalance", user)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetStablecoinBalance is a free data retrieval call binding the contract method 0x7559fe62.
//
// Solidity: function GetStablecoinBalance(address user) view returns(uint256, uint256)
func (_Broker *BrokerSession) GetStablecoinBalance(user common.Address) (*big.Int, *big.Int, error) {
	return _Broker.Contract.GetStablecoinBalance(&_Broker.CallOpts, user)
}

// GetStablecoinBalance is a free data retrieval call binding the contract method 0x7559fe62.
//
// Solidity: function GetStablecoinBalance(address user) view returns(uint256, uint256)
func (_Broker *BrokerCallerSession) GetStablecoinBalance(user common.Address) (*big.Int, *big.Int, error) {
	return _Broker.Contract.GetStablecoinBalance(&_Broker.CallOpts, user)
}

// GetTime is a free data retrieval call binding the contract method 0x18c5e502.
//
// Solidity: function GetTime() view returns(uint256)
func (_Broker *BrokerCaller) GetTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Broker.contract.Call(opts, &out, "GetTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTime is a free data retrieval call binding the contract method 0x18c5e502.
//
// Solidity: function GetTime() view returns(uint256)
func (_Broker *BrokerSession) GetTime() (*big.Int, error) {
	return _Broker.Contract.GetTime(&_Broker.CallOpts)
}

// GetTime is a free data retrieval call binding the contract method 0x18c5e502.
//
// Solidity: function GetTime() view returns(uint256)
func (_Broker *BrokerCallerSession) GetTime() (*big.Int, error) {
	return _Broker.Contract.GetTime(&_Broker.CallOpts)
}

// GetUsersBookings is a free data retrieval call binding the contract method 0x7f7a1915.
//
// Solidity: function GetUsersBookings(address user) view returns((uint64,uint64,address,address,uint64,uint256,uint256,uint64)[] filteredBookings)
func (_Broker *BrokerCaller) GetUsersBookings(opts *bind.CallOpts, user common.Address) ([]BrokerBooking, error) {
	var out []interface{}
	err := _Broker.contract.Call(opts, &out, "GetUsersBookings", user)

	if err != nil {
		return *new([]BrokerBooking), err
	}

	out0 := *abi.ConvertType(out[0], new([]BrokerBooking)).(*[]BrokerBooking)

	return out0, err

}

// GetUsersBookings is a free data retrieval call binding the contract method 0x7f7a1915.
//
// Solidity: function GetUsersBookings(address user) view returns((uint64,uint64,address,address,uint64,uint256,uint256,uint64)[] filteredBookings)
func (_Broker *BrokerSession) GetUsersBookings(user common.Address) ([]BrokerBooking, error) {
	return _Broker.Contract.GetUsersBookings(&_Broker.CallOpts, user)
}

// GetUsersBookings is a free data retrieval call binding the contract method 0x7f7a1915.
//
// Solidity: function GetUsersBookings(address user) view returns((uint64,uint64,address,address,uint64,uint256,uint256,uint64)[] filteredBookings)
func (_Broker *BrokerCallerSession) GetUsersBookings(user common.Address) ([]BrokerBooking, error) {
	return _Broker.Contract.GetUsersBookings(&_Broker.CallOpts, user)
}

// SECONDSINWEEK is a free data retrieval call binding the contract method 0xe088b29e.
//
// Solidity: function SECONDS_IN_WEEK() view returns(uint64)
func (_Broker *BrokerCaller) SECONDSINWEEK(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Broker.contract.Call(opts, &out, "SECONDS_IN_WEEK")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// SECONDSINWEEK is a free data retrieval call binding the contract method 0xe088b29e.
//
// Solidity: function SECONDS_IN_WEEK() view returns(uint64)
func (_Broker *BrokerSession) SECONDSINWEEK() (uint64, error) {
	return _Broker.Contract.SECONDSINWEEK(&_Broker.CallOpts)
}

// SECONDSINWEEK is a free data retrieval call binding the contract method 0xe088b29e.
//
// Solidity: function SECONDS_IN_WEEK() view returns(uint64)
func (_Broker *BrokerCallerSession) SECONDSINWEEK() (uint64, error) {
	return _Broker.Contract.SECONDSINWEEK(&_Broker.CallOpts)
}

// GetCommunityContract is a free data retrieval call binding the contract method 0x569e2e0f.
//
// Solidity: function getCommunityContract() view returns(address)
func (_Broker *BrokerCaller) GetCommunityContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Broker.contract.Call(opts, &out, "getCommunityContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetCommunityContract is a free data retrieval call binding the contract method 0x569e2e0f.
//
// Solidity: function getCommunityContract() view returns(address)
func (_Broker *BrokerSession) GetCommunityContract() (common.Address, error) {
	return _Broker.Contract.GetCommunityContract(&_Broker.CallOpts)
}

// GetCommunityContract is a free data retrieval call binding the contract method 0x569e2e0f.
//
// Solidity: function getCommunityContract() view returns(address)
func (_Broker *BrokerCallerSession) GetCommunityContract() (common.Address, error) {
	return _Broker.Contract.GetCommunityContract(&_Broker.CallOpts)
}

// GetCommunityFee is a free data retrieval call binding the contract method 0x90ad688b.
//
// Solidity: function getCommunityFee() view returns(uint64)
func (_Broker *BrokerCaller) GetCommunityFee(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Broker.contract.Call(opts, &out, "getCommunityFee")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetCommunityFee is a free data retrieval call binding the contract method 0x90ad688b.
//
// Solidity: function getCommunityFee() view returns(uint64)
func (_Broker *BrokerSession) GetCommunityFee() (uint64, error) {
	return _Broker.Contract.GetCommunityFee(&_Broker.CallOpts)
}

// GetCommunityFee is a free data retrieval call binding the contract method 0x90ad688b.
//
// Solidity: function getCommunityFee() view returns(uint64)
func (_Broker *BrokerCallerSession) GetCommunityFee() (uint64, error) {
	return _Broker.Contract.GetCommunityFee(&_Broker.CallOpts)
}

// GetMinerUrl is a free data retrieval call binding the contract method 0xab49ec45.
//
// Solidity: function getMinerUrl(address _user) view returns(bytes32)
func (_Broker *BrokerCaller) GetMinerUrl(opts *bind.CallOpts, _user common.Address) ([32]byte, error) {
	var out []interface{}
	err := _Broker.contract.Call(opts, &out, "getMinerUrl", _user)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetMinerUrl is a free data retrieval call binding the contract method 0xab49ec45.
//
// Solidity: function getMinerUrl(address _user) view returns(bytes32)
func (_Broker *BrokerSession) GetMinerUrl(_user common.Address) ([32]byte, error) {
	return _Broker.Contract.GetMinerUrl(&_Broker.CallOpts, _user)
}

// GetMinerUrl is a free data retrieval call binding the contract method 0xab49ec45.
//
// Solidity: function getMinerUrl(address _user) view returns(bytes32)
func (_Broker *BrokerCallerSession) GetMinerUrl(_user common.Address) ([32]byte, error) {
	return _Broker.Contract.GetMinerUrl(&_Broker.CallOpts, _user)
}

// GetStablecoinAddress is a free data retrieval call binding the contract method 0x7bef4f54.
//
// Solidity: function getStablecoinAddress() view returns(address)
func (_Broker *BrokerCaller) GetStablecoinAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Broker.contract.Call(opts, &out, "getStablecoinAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetStablecoinAddress is a free data retrieval call binding the contract method 0x7bef4f54.
//
// Solidity: function getStablecoinAddress() view returns(address)
func (_Broker *BrokerSession) GetStablecoinAddress() (common.Address, error) {
	return _Broker.Contract.GetStablecoinAddress(&_Broker.CallOpts)
}

// GetStablecoinAddress is a free data retrieval call binding the contract method 0x7bef4f54.
//
// Solidity: function getStablecoinAddress() view returns(address)
func (_Broker *BrokerCallerSession) GetStablecoinAddress() (common.Address, error) {
	return _Broker.Contract.GetStablecoinAddress(&_Broker.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint64)
func (_Broker *BrokerCaller) Version(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Broker.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint64)
func (_Broker *BrokerSession) Version() (uint64, error) {
	return _Broker.Contract.Version(&_Broker.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint64)
func (_Broker *BrokerCallerSession) Version() (uint64, error) {
	return _Broker.Contract.Version(&_Broker.CallOpts)
}

// AddOffer is a paid mutator transaction binding the contract method 0x8f351016.
//
// Solidity: function AddOffer(uint64 pricePerSecond, uint64 vmTypeId, uint64 machinesAvailable) returns(uint64)
func (_Broker *BrokerTransactor) AddOffer(opts *bind.TransactOpts, pricePerSecond uint64, vmTypeId uint64, machinesAvailable uint64) (*types.Transaction, error) {
	return _Broker.contract.Transact(opts, "AddOffer", pricePerSecond, vmTypeId, machinesAvailable)
}

// AddOffer is a paid mutator transaction binding the contract method 0x8f351016.
//
// Solidity: function AddOffer(uint64 pricePerSecond, uint64 vmTypeId, uint64 machinesAvailable) returns(uint64)
func (_Broker *BrokerSession) AddOffer(pricePerSecond uint64, vmTypeId uint64, machinesAvailable uint64) (*types.Transaction, error) {
	return _Broker.Contract.AddOffer(&_Broker.TransactOpts, pricePerSecond, vmTypeId, machinesAvailable)
}

// AddOffer is a paid mutator transaction binding the contract method 0x8f351016.
//
// Solidity: function AddOffer(uint64 pricePerSecond, uint64 vmTypeId, uint64 machinesAvailable) returns(uint64)
func (_Broker *BrokerTransactorSession) AddOffer(pricePerSecond uint64, vmTypeId uint64, machinesAvailable uint64) (*types.Transaction, error) {
	return _Broker.Contract.AddOffer(&_Broker.TransactOpts, pricePerSecond, vmTypeId, machinesAvailable)
}

// BookVM is a paid mutator transaction binding the contract method 0x75b7322e.
//
// Solidity: function BookVM(uint64 offerIndex) returns(uint64)
func (_Broker *BrokerTransactor) BookVM(opts *bind.TransactOpts, offerIndex uint64) (*types.Transaction, error) {
	return _Broker.contract.Transact(opts, "BookVM", offerIndex)
}

// BookVM is a paid mutator transaction binding the contract method 0x75b7322e.
//
// Solidity: function BookVM(uint64 offerIndex) returns(uint64)
func (_Broker *BrokerSession) BookVM(offerIndex uint64) (*types.Transaction, error) {
	return _Broker.Contract.BookVM(&_Broker.TransactOpts, offerIndex)
}

// BookVM is a paid mutator transaction binding the contract method 0x75b7322e.
//
// Solidity: function BookVM(uint64 offerIndex) returns(uint64)
func (_Broker *BrokerTransactorSession) BookVM(offerIndex uint64) (*types.Transaction, error) {
	return _Broker.Contract.BookVM(&_Broker.TransactOpts, offerIndex)
}

// ClaimPayment is a paid mutator transaction binding the contract method 0xf3fe369c.
//
// Solidity: function ClaimPayment(uint64 bookingId) returns()
func (_Broker *BrokerTransactor) ClaimPayment(opts *bind.TransactOpts, bookingId uint64) (*types.Transaction, error) {
	return _Broker.contract.Transact(opts, "ClaimPayment", bookingId)
}

// ClaimPayment is a paid mutator transaction binding the contract method 0xf3fe369c.
//
// Solidity: function ClaimPayment(uint64 bookingId) returns()
func (_Broker *BrokerSession) ClaimPayment(bookingId uint64) (*types.Transaction, error) {
	return _Broker.Contract.ClaimPayment(&_Broker.TransactOpts, bookingId)
}

// ClaimPayment is a paid mutator transaction binding the contract method 0xf3fe369c.
//
// Solidity: function ClaimPayment(uint64 bookingId) returns()
func (_Broker *BrokerTransactorSession) ClaimPayment(bookingId uint64) (*types.Transaction, error) {
	return _Broker.Contract.ClaimPayment(&_Broker.TransactOpts, bookingId)
}

// DepositStablecoin is a paid mutator transaction binding the contract method 0x6d8351d6.
//
// Solidity: function DepositStablecoin(uint256 numTokens) returns(bool)
func (_Broker *BrokerTransactor) DepositStablecoin(opts *bind.TransactOpts, numTokens *big.Int) (*types.Transaction, error) {
	return _Broker.contract.Transact(opts, "DepositStablecoin", numTokens)
}

// DepositStablecoin is a paid mutator transaction binding the contract method 0x6d8351d6.
//
// Solidity: function DepositStablecoin(uint256 numTokens) returns(bool)
func (_Broker *BrokerSession) DepositStablecoin(numTokens *big.Int) (*types.Transaction, error) {
	return _Broker.Contract.DepositStablecoin(&_Broker.TransactOpts, numTokens)
}

// DepositStablecoin is a paid mutator transaction binding the contract method 0x6d8351d6.
//
// Solidity: function DepositStablecoin(uint256 numTokens) returns(bool)
func (_Broker *BrokerTransactorSession) DepositStablecoin(numTokens *big.Int) (*types.Transaction, error) {
	return _Broker.Contract.DepositStablecoin(&_Broker.TransactOpts, numTokens)
}

// RemoveOffer is a paid mutator transaction binding the contract method 0x88317656.
//
// Solidity: function RemoveOffer(uint64 offerIndex) returns()
func (_Broker *BrokerTransactor) RemoveOffer(opts *bind.TransactOpts, offerIndex uint64) (*types.Transaction, error) {
	return _Broker.contract.Transact(opts, "RemoveOffer", offerIndex)
}

// RemoveOffer is a paid mutator transaction binding the contract method 0x88317656.
//
// Solidity: function RemoveOffer(uint64 offerIndex) returns()
func (_Broker *BrokerSession) RemoveOffer(offerIndex uint64) (*types.Transaction, error) {
	return _Broker.Contract.RemoveOffer(&_Broker.TransactOpts, offerIndex)
}

// RemoveOffer is a paid mutator transaction binding the contract method 0x88317656.
//
// Solidity: function RemoveOffer(uint64 offerIndex) returns()
func (_Broker *BrokerTransactorSession) RemoveOffer(offerIndex uint64) (*types.Transaction, error) {
	return _Broker.Contract.RemoveOffer(&_Broker.TransactOpts, offerIndex)
}

// SetCommunityContract is a paid mutator transaction binding the contract method 0x6b620ed5.
//
// Solidity: function SetCommunityContract(address newCommunityAddress) returns()
func (_Broker *BrokerTransactor) SetCommunityContract(opts *bind.TransactOpts, newCommunityAddress common.Address) (*types.Transaction, error) {
	return _Broker.contract.Transact(opts, "SetCommunityContract", newCommunityAddress)
}

// SetCommunityContract is a paid mutator transaction binding the contract method 0x6b620ed5.
//
// Solidity: function SetCommunityContract(address newCommunityAddress) returns()
func (_Broker *BrokerSession) SetCommunityContract(newCommunityAddress common.Address) (*types.Transaction, error) {
	return _Broker.Contract.SetCommunityContract(&_Broker.TransactOpts, newCommunityAddress)
}

// SetCommunityContract is a paid mutator transaction binding the contract method 0x6b620ed5.
//
// Solidity: function SetCommunityContract(address newCommunityAddress) returns()
func (_Broker *BrokerTransactorSession) SetCommunityContract(newCommunityAddress common.Address) (*types.Transaction, error) {
	return _Broker.Contract.SetCommunityContract(&_Broker.TransactOpts, newCommunityAddress)
}

// SetStablecoinAddress is a paid mutator transaction binding the contract method 0xc9b16fee.
//
// Solidity: function SetStablecoinAddress(address newStablecoinAddress) returns()
func (_Broker *BrokerTransactor) SetStablecoinAddress(opts *bind.TransactOpts, newStablecoinAddress common.Address) (*types.Transaction, error) {
	return _Broker.contract.Transact(opts, "SetStablecoinAddress", newStablecoinAddress)
}

// SetStablecoinAddress is a paid mutator transaction binding the contract method 0xc9b16fee.
//
// Solidity: function SetStablecoinAddress(address newStablecoinAddress) returns()
func (_Broker *BrokerSession) SetStablecoinAddress(newStablecoinAddress common.Address) (*types.Transaction, error) {
	return _Broker.Contract.SetStablecoinAddress(&_Broker.TransactOpts, newStablecoinAddress)
}

// SetStablecoinAddress is a paid mutator transaction binding the contract method 0xc9b16fee.
//
// Solidity: function SetStablecoinAddress(address newStablecoinAddress) returns()
func (_Broker *BrokerTransactorSession) SetStablecoinAddress(newStablecoinAddress common.Address) (*types.Transaction, error) {
	return _Broker.Contract.SetStablecoinAddress(&_Broker.TransactOpts, newStablecoinAddress)
}

// StopVM is a paid mutator transaction binding the contract method 0x9a61a20e.
//
// Solidity: function StopVM(uint64 bookingId, uint8 reason) returns()
func (_Broker *BrokerTransactor) StopVM(opts *bind.TransactOpts, bookingId uint64, reason uint8) (*types.Transaction, error) {
	return _Broker.contract.Transact(opts, "StopVM", bookingId, reason)
}

// StopVM is a paid mutator transaction binding the contract method 0x9a61a20e.
//
// Solidity: function StopVM(uint64 bookingId, uint8 reason) returns()
func (_Broker *BrokerSession) StopVM(bookingId uint64, reason uint8) (*types.Transaction, error) {
	return _Broker.Contract.StopVM(&_Broker.TransactOpts, bookingId, reason)
}

// StopVM is a paid mutator transaction binding the contract method 0x9a61a20e.
//
// Solidity: function StopVM(uint64 bookingId, uint8 reason) returns()
func (_Broker *BrokerTransactorSession) StopVM(bookingId uint64, reason uint8) (*types.Transaction, error) {
	return _Broker.Contract.StopVM(&_Broker.TransactOpts, bookingId, reason)
}

// UpdateOffer is a paid mutator transaction binding the contract method 0x3fec2ddf.
//
// Solidity: function UpdateOffer(uint64 offerIndex, uint64 machinesAvailable, uint64 pps) returns()
func (_Broker *BrokerTransactor) UpdateOffer(opts *bind.TransactOpts, offerIndex uint64, machinesAvailable uint64, pps uint64) (*types.Transaction, error) {
	return _Broker.contract.Transact(opts, "UpdateOffer", offerIndex, machinesAvailable, pps)
}

// UpdateOffer is a paid mutator transaction binding the contract method 0x3fec2ddf.
//
// Solidity: function UpdateOffer(uint64 offerIndex, uint64 machinesAvailable, uint64 pps) returns()
func (_Broker *BrokerSession) UpdateOffer(offerIndex uint64, machinesAvailable uint64, pps uint64) (*types.Transaction, error) {
	return _Broker.Contract.UpdateOffer(&_Broker.TransactOpts, offerIndex, machinesAvailable, pps)
}

// UpdateOffer is a paid mutator transaction binding the contract method 0x3fec2ddf.
//
// Solidity: function UpdateOffer(uint64 offerIndex, uint64 machinesAvailable, uint64 pps) returns()
func (_Broker *BrokerTransactorSession) UpdateOffer(offerIndex uint64, machinesAvailable uint64, pps uint64) (*types.Transaction, error) {
	return _Broker.Contract.UpdateOffer(&_Broker.TransactOpts, offerIndex, machinesAvailable, pps)
}

// WithdrawStablecoin is a paid mutator transaction binding the contract method 0x6f427a80.
//
// Solidity: function WithdrawStablecoin(uint256 amt) returns(bool)
func (_Broker *BrokerTransactor) WithdrawStablecoin(opts *bind.TransactOpts, amt *big.Int) (*types.Transaction, error) {
	return _Broker.contract.Transact(opts, "WithdrawStablecoin", amt)
}

// WithdrawStablecoin is a paid mutator transaction binding the contract method 0x6f427a80.
//
// Solidity: function WithdrawStablecoin(uint256 amt) returns(bool)
func (_Broker *BrokerSession) WithdrawStablecoin(amt *big.Int) (*types.Transaction, error) {
	return _Broker.Contract.WithdrawStablecoin(&_Broker.TransactOpts, amt)
}

// WithdrawStablecoin is a paid mutator transaction binding the contract method 0x6f427a80.
//
// Solidity: function WithdrawStablecoin(uint256 amt) returns(bool)
func (_Broker *BrokerTransactorSession) WithdrawStablecoin(amt *big.Int) (*types.Transaction, error) {
	return _Broker.Contract.WithdrawStablecoin(&_Broker.TransactOpts, amt)
}

// SetCommunityFee is a paid mutator transaction binding the contract method 0xa7bf7570.
//
// Solidity: function setCommunityFee(uint64 fee) returns(bool)
func (_Broker *BrokerTransactor) SetCommunityFee(opts *bind.TransactOpts, fee uint64) (*types.Transaction, error) {
	return _Broker.contract.Transact(opts, "setCommunityFee", fee)
}

// SetCommunityFee is a paid mutator transaction binding the contract method 0xa7bf7570.
//
// Solidity: function setCommunityFee(uint64 fee) returns(bool)
func (_Broker *BrokerSession) SetCommunityFee(fee uint64) (*types.Transaction, error) {
	return _Broker.Contract.SetCommunityFee(&_Broker.TransactOpts, fee)
}

// SetCommunityFee is a paid mutator transaction binding the contract method 0xa7bf7570.
//
// Solidity: function setCommunityFee(uint64 fee) returns(bool)
func (_Broker *BrokerTransactorSession) SetCommunityFee(fee uint64) (*types.Transaction, error) {
	return _Broker.Contract.SetCommunityFee(&_Broker.TransactOpts, fee)
}

// SetMunerUrl is a paid mutator transaction binding the contract method 0xf44f11cd.
//
// Solidity: function setMunerUrl(bytes32 url) returns()
func (_Broker *BrokerTransactor) SetMunerUrl(opts *bind.TransactOpts, url [32]byte) (*types.Transaction, error) {
	return _Broker.contract.Transact(opts, "setMunerUrl", url)
}

// SetMunerUrl is a paid mutator transaction binding the contract method 0xf44f11cd.
//
// Solidity: function setMunerUrl(bytes32 url) returns()
func (_Broker *BrokerSession) SetMunerUrl(url [32]byte) (*types.Transaction, error) {
	return _Broker.Contract.SetMunerUrl(&_Broker.TransactOpts, url)
}

// SetMunerUrl is a paid mutator transaction binding the contract method 0xf44f11cd.
//
// Solidity: function setMunerUrl(bytes32 url) returns()
func (_Broker *BrokerTransactorSession) SetMunerUrl(url [32]byte) (*types.Transaction, error) {
	return _Broker.Contract.SetMunerUrl(&_Broker.TransactOpts, url)
}

// BrokerComplaintIterator is returned from FilterComplaint and is used to iterate over the raw logs and unpacked data for Complaint events raised by the Broker contract.
type BrokerComplaintIterator struct {
	Event *BrokerComplaint // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BrokerComplaintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BrokerComplaint)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BrokerComplaint)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BrokerComplaintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BrokerComplaintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BrokerComplaint represents a Complaint event raised by the Broker contract.
type BrokerComplaint struct {
	User   common.Address
	Miner  common.Address
	Reason uint8
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterComplaint is a free log retrieval operation binding the contract event 0x1db97daa25845ee925bd84e6cf559402b161e27653e3f3187c1859f98634c1c5.
//
// Solidity: event Complaint(address indexed user, address indexed miner, uint8 indexed reason)
func (_Broker *BrokerFilterer) FilterComplaint(opts *bind.FilterOpts, user []common.Address, miner []common.Address, reason []uint8) (*BrokerComplaintIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}
	var reasonRule []interface{}
	for _, reasonItem := range reason {
		reasonRule = append(reasonRule, reasonItem)
	}

	logs, sub, err := _Broker.contract.FilterLogs(opts, "Complaint", userRule, minerRule, reasonRule)
	if err != nil {
		return nil, err
	}
	return &BrokerComplaintIterator{contract: _Broker.contract, event: "Complaint", logs: logs, sub: sub}, nil
}

// WatchComplaint is a free log subscription operation binding the contract event 0x1db97daa25845ee925bd84e6cf559402b161e27653e3f3187c1859f98634c1c5.
//
// Solidity: event Complaint(address indexed user, address indexed miner, uint8 indexed reason)
func (_Broker *BrokerFilterer) WatchComplaint(opts *bind.WatchOpts, sink chan<- *BrokerComplaint, user []common.Address, miner []common.Address, reason []uint8) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}
	var reasonRule []interface{}
	for _, reasonItem := range reason {
		reasonRule = append(reasonRule, reasonItem)
	}

	logs, sub, err := _Broker.contract.WatchLogs(opts, "Complaint", userRule, minerRule, reasonRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BrokerComplaint)
				if err := _Broker.contract.UnpackLog(event, "Complaint", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseComplaint is a log parse operation binding the contract event 0x1db97daa25845ee925bd84e6cf559402b161e27653e3f3187c1859f98634c1c5.
//
// Solidity: event Complaint(address indexed user, address indexed miner, uint8 indexed reason)
func (_Broker *BrokerFilterer) ParseComplaint(log types.Log) (*BrokerComplaint, error) {
	event := new(BrokerComplaint)
	if err := _Broker.contract.UnpackLog(event, "Complaint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BrokerPaymentIterator is returned from FilterPayment and is used to iterate over the raw logs and unpacked data for Payment events raised by the Broker contract.
type BrokerPaymentIterator struct {
	Event *BrokerPayment // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BrokerPaymentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BrokerPayment)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BrokerPayment)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BrokerPaymentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BrokerPaymentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BrokerPayment represents a Payment event raised by the Broker contract.
type BrokerPayment struct {
	User   common.Address
	Miner  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPayment is a free log retrieval operation binding the contract event 0x8ed7db30df47fbd87811a5e8a95a94838f0c0241263d9a1865d6a2a3e10516e2.
//
// Solidity: event Payment(address indexed user, address indexed miner, uint256 amount)
func (_Broker *BrokerFilterer) FilterPayment(opts *bind.FilterOpts, user []common.Address, miner []common.Address) (*BrokerPaymentIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}

	logs, sub, err := _Broker.contract.FilterLogs(opts, "Payment", userRule, minerRule)
	if err != nil {
		return nil, err
	}
	return &BrokerPaymentIterator{contract: _Broker.contract, event: "Payment", logs: logs, sub: sub}, nil
}

// WatchPayment is a free log subscription operation binding the contract event 0x8ed7db30df47fbd87811a5e8a95a94838f0c0241263d9a1865d6a2a3e10516e2.
//
// Solidity: event Payment(address indexed user, address indexed miner, uint256 amount)
func (_Broker *BrokerFilterer) WatchPayment(opts *bind.WatchOpts, sink chan<- *BrokerPayment, user []common.Address, miner []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}

	logs, sub, err := _Broker.contract.WatchLogs(opts, "Payment", userRule, minerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BrokerPayment)
				if err := _Broker.contract.UnpackLog(event, "Payment", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePayment is a log parse operation binding the contract event 0x8ed7db30df47fbd87811a5e8a95a94838f0c0241263d9a1865d6a2a3e10516e2.
//
// Solidity: event Payment(address indexed user, address indexed miner, uint256 amount)
func (_Broker *BrokerFilterer) ParsePayment(log types.Log) (*BrokerPayment, error) {
	event := new(BrokerPayment)
	if err := _Broker.contract.UnpackLog(event, "Payment", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
