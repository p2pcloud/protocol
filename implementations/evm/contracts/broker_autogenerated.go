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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"minerPayout\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"timeUsed\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"vmTypeId\",\"type\":\"uint64\"}],\"name\":\"BookingClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"vmTypeId\",\"type\":\"uint64\"}],\"name\":\"BookingExtended\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"minerPayout\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"timeUsed\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"vmTypeId\",\"type\":\"uint64\"}],\"name\":\"BookingReported\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"vmTypeId\",\"type\":\"uint64\"}],\"name\":\"BookingStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"minerPayout\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"timeUsed\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"vmTypeId\",\"type\":\"uint64\"}],\"name\":\"BookingStopped\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"pricePerSecond\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"vmTypeId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"machinesAvailable\",\"type\":\"uint64\"}],\"name\":\"AddOffer\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"offerIndex\",\"type\":\"uint64\"}],\"name\":\"BookVM\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"bookingId\",\"type\":\"uint64\"}],\"name\":\"ClaimPayment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"numTokens\",\"type\":\"uint256\"}],\"name\":\"DepositStablecoin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_miner\",\"type\":\"address\"}],\"name\":\"FindBookingsByMiner\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"vmTypeId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"pricePerSecond\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"bookedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastPayment\",\"type\":\"uint256\"}],\"internalType\":\"structBroker.Booking[]\",\"name\":\"filteredBookings\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"FindBookingsByUser\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"vmTypeId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"pricePerSecond\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"bookedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastPayment\",\"type\":\"uint256\"}],\"internalType\":\"structBroker.Booking[]\",\"name\":\"filteredBookings\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetAvailableOffers\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"pricePerSecond\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"machinesAvailable\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"vmTypeId\",\"type\":\"uint64\"}],\"internalType\":\"structBroker.VMOffer[]\",\"name\":\"filteredOffers\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_vmTypeId\",\"type\":\"uint64\"}],\"name\":\"GetAvailableOffersByType\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"pricePerSecond\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"machinesAvailable\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"vmTypeId\",\"type\":\"uint64\"}],\"internalType\":\"structBroker.VMOffer[]\",\"name\":\"filteredOffers\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"}],\"name\":\"GetBooking\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"vmTypeId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"pricePerSecond\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"bookedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastPayment\",\"type\":\"uint256\"}],\"internalType\":\"structBroker.Booking\",\"name\":\"booking\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"}],\"name\":\"GetMinersOffers\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"pricePerSecond\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"machinesAvailable\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"vmTypeId\",\"type\":\"uint64\"}],\"internalType\":\"structBroker.VMOffer[]\",\"name\":\"filteredOffers\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"GetStablecoinBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"GetUsersBookings\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"vmTypeId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"pricePerSecond\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"bookedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastPayment\",\"type\":\"uint256\"}],\"internalType\":\"structBroker.Booking[]\",\"name\":\"filteredBookings\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"offerIndex\",\"type\":\"uint64\"}],\"name\":\"RemoveOffer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SECONDS_IN_WEEK\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newCommunityAddress\",\"type\":\"address\"}],\"name\":\"SetCommunityContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"newStablecoinAddress\",\"type\":\"address\"}],\"name\":\"SetStablecoinAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"offerIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"machinesAvailable\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"pps\",\"type\":\"uint64\"}],\"name\":\"UpdateOffer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"name\":\"WithdrawStablecoin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCommunityContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCommunityFee\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getMinerUrl\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStablecoinAddress\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"fee\",\"type\":\"uint64\"}],\"name\":\"setCommunityFee\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"url\",\"type\":\"bytes32\"}],\"name\":\"setMunerUrl\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040526000600160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506000600360006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506101f4600860146101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555034801561008f57600080fd5b5033600860006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550614810806100e06000396000f3fe608060405234801561001057600080fd5b506004361061018e5760003560e01c80637f7a1915116100de578063b65c4b9c11610097578063d3554ab511610071578063d3554ab5146104f8578063e088b29e14610528578063f3fe369c14610546578063f44f11cd146105625761018e565b8063b65c4b9c1461048e578063c2161a05146104ac578063c9b16fee146104dc5761018e565b80637f7a19151461039457806388317656146103c45780638f351016146103e057806390ad688b14610410578063a7bf75701461042e578063ab49ec451461045e5761018e565b8063569e2e0f1161014b5780636f427a80116101255780636f427a80146102e55780637559fe621461031557806375b7322e146103465780637bef4f54146103765761018e565b8063569e2e0f1461027b5780636b620ed5146102995780636d8351d6146102b55761018e565b806318c5e50214610193578063366445bf146101b15780633fec2ddf146101e1578063534b3ccb146101fd5780635405b5e51461022d57806354fd4d501461025d575b600080fd5b61019b61057e565b6040516101a8919061369e565b60405180910390f35b6101cb60048036038101906101c6919061371c565b610586565b6040516101d891906138c7565b60405180910390f35b6101fb60048036038101906101f69190613915565b610972565b005b6102176004803603810190610212919061371c565b610ace565b60405161022491906138c7565b60405180910390f35b61024760048036038101906102429190613968565b610eba565b6040516102549190613aac565b60405180910390f35b61026561129c565b6040516102729190613add565b60405180910390f35b6102836112a1565b6040516102909190613b07565b60405180910390f35b6102b360048036038101906102ae919061371c565b6112cb565b005b6102cf60048036038101906102ca9190613b4e565b61139f565b6040516102dc9190613b96565b60405180910390f35b6102ff60048036038101906102fa9190613b4e565b611519565b60405161030c9190613b96565b60405180910390f35b61032f600480360381019061032a919061371c565b6116f3565b60405161033d929190613bb1565b60405180910390f35b610360600480360381019061035b9190613968565b611759565b60405161036d9190613add565b60405180910390f35b61037e611d4e565b60405161038b9190613c39565b60405180910390f35b6103ae60048036038101906103a9919061371c565b611d78565b6040516103bb91906138c7565b60405180910390f35b6103de60048036038101906103d99190613968565b612164565b005b6103fa60048036038101906103f59190613915565b6122da565b6040516104079190613add565b60405180910390f35b610418612511565b6040516104259190613add565b60405180910390f35b61044860048036038101906104439190613968565b61252f565b6040516104559190613b96565b60405180910390f35b6104786004803603810190610473919061371c565b612641565b6040516104859190613c6d565b60405180910390f35b61049661268a565b6040516104a39190613aac565b60405180910390f35b6104c660048036038101906104c19190613968565b612a0e565b6040516104d39190613d16565b60405180910390f35b6104f660048036038101906104f19190613d6f565b612ba5565b005b610512600480360381019061050d919061371c565b612c79565b60405161051f9190613aac565b60405180910390f35b61053061302b565b60405161053d9190613add565b60405180910390f35b610560600480360381019061055b9190613968565b613032565b005b61057c60048036038101906105779190613dc8565b6134d5565b005b600042905090565b60606000600360009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff8111156105c4576105c3613df5565b5b6040519080825280602002602001820160405280156105fd57816020015b6105ea613591565b8152602001906001900390816105e25790505b509050600080600090505b600360009054906101000a900467ffffffffffffffff1667ffffffffffffffff168167ffffffffffffffff16101561088a578473ffffffffffffffffffffffffffffffffffffffff16600260008367ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff160361087757600260008267ffffffffffffffff1667ffffffffffffffff1681526020019081526020016000206040518060e00160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820160149054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff16815260200160038201548152602001600482015481525050838367ffffffffffffffff168151811061085c5761085b613e24565b5b60200260200101819052506001826108749190613e82565b91505b808061088290613ec0565b915050610608565b508067ffffffffffffffff1667ffffffffffffffff8111156108af576108ae613df5565b5b6040519080825280602002602001820160405280156108e857816020015b6108d5613591565b8152602001906001900390816108cd5790505b50925060005b8167ffffffffffffffff168167ffffffffffffffff16101561096a57828167ffffffffffffffff168151811061092757610926613e24565b5b6020026020010151848267ffffffffffffffff168151811061094c5761094b613e24565b5b6020026020010181905250808061096290613ec0565b9150506108ee565b505050919050565b3373ffffffffffffffffffffffffffffffffffffffff166000808567ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060000160089054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610a29576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a2090613f73565b60405180910390fd5b816000808567ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060010160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550806000808567ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060010160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550505050565b60606000600360009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff811115610b0c57610b0b613df5565b5b604051908082528060200260200182016040528015610b4557816020015b610b32613591565b815260200190600190039081610b2a5790505b509050600080600090505b600360009054906101000a900467ffffffffffffffff1667ffffffffffffffff168167ffffffffffffffff161015610dd2578473ffffffffffffffffffffffffffffffffffffffff16600260008367ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603610dbf57600260008267ffffffffffffffff1667ffffffffffffffff1681526020019081526020016000206040518060e00160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820160149054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff16815260200160038201548152602001600482015481525050838367ffffffffffffffff1681518110610da457610da3613e24565b5b6020026020010181905250600182610dbc9190613e82565b91505b8080610dca90613ec0565b915050610b50565b508067ffffffffffffffff1667ffffffffffffffff811115610df757610df6613df5565b5b604051908082528060200260200182016040528015610e3057816020015b610e1d613591565b815260200190600190039081610e155790505b50925060005b8167ffffffffffffffff168167ffffffffffffffff161015610eb257828167ffffffffffffffff1681518110610e6f57610e6e613e24565b5b6020026020010151848267ffffffffffffffff1681518110610e9457610e93613e24565b5b60200260200101819052508080610eaa90613ec0565b915050610e36565b505050919050565b60606000600160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff811115610ef857610ef7613df5565b5b604051908082528060200260200182016040528015610f3157816020015b610f1e613618565b815260200190600190039081610f165790505b509050600080600090505b600160009054906101000a900467ffffffffffffffff1667ffffffffffffffff168167ffffffffffffffff1610156111b4578467ffffffffffffffff166000808367ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060010160109054906101000a900467ffffffffffffffff1667ffffffffffffffff16148015611015575060008060008367ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060010160089054906101000a900467ffffffffffffffff1667ffffffffffffffff16115b156111a1576000808267ffffffffffffffff1667ffffffffffffffff1681526020019081526020016000206040518060a00160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016001820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681525050838367ffffffffffffffff168151811061118657611185613e24565b5b602002602001018190525060018261119e9190613e82565b91505b80806111ac90613ec0565b915050610f3c565b508067ffffffffffffffff1667ffffffffffffffff8111156111d9576111d8613df5565b5b60405190808252806020026020018201604052801561121257816020015b6111ff613618565b8152602001906001900390816111f75790505b50925060005b8167ffffffffffffffff168167ffffffffffffffff16101561129457828167ffffffffffffffff168151811061125157611250613e24565b5b6020026020010151848267ffffffffffffffff168151811061127657611275613e24565b5b6020026020010181905250808061128c90613ec0565b915050611218565b505050919050565b600181565b6000600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461135b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161135290614005565b60405180910390fd5b80600860006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6000600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd3330856040518463ffffffff1660e01b815260040161140093929190614025565b6020604051808303816000875af115801561141f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114439190614088565b611482576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161147990614101565b60405180910390fd5b81600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546114cd9190614121565b600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555060019050919050565b6000806115253361351c565b600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461156f9190614177565b9050828110156115b4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016115ab906141f7565b60405180910390fd5b600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb33856040518363ffffffff1660e01b8152600401611611929190614217565b6020604051808303816000875af1158015611630573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116549190614088565b611693576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161168a9061428c565b60405180910390fd5b82600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546116e29190614177565b925050819055506001915050919050565b60008060006117018461351c565b905080600560008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461174e9190614177565b819250925050915091565b6000806000808467ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060010160089054906101000a900467ffffffffffffffff1667ffffffffffffffff16116117e4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016117db906142f8565b60405180910390fd5b600062093a806000808567ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060010160009054906101000a900467ffffffffffffffff166118329190614318565b67ffffffffffffffff166118453361351c565b61184f9190614121565b9050600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020548111156118d3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016118ca906143f2565b60405180910390fd5b60006040518060e00160405280600360009054906101000a900467ffffffffffffffff1667ffffffffffffffff1681526020016000808767ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060010160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1681526020016000808767ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060000160089054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020013373ffffffffffffffffffffffffffffffffffffffff1681526020016000808767ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060010160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1681526020014281526020014281525090508060026000600360009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060208201518160000160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060408201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060808201518160020160146101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060a0820151816003015560c082015181600401559050506003600081819054906101000a900467ffffffffffffffff1680929190611bc390613ec0565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550506000808567ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060010160009054906101000a900467ffffffffffffffff16600660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282829054906101000a900467ffffffffffffffff16611c899190613e82565b92506101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060016000808667ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060010160088282829054906101000a900467ffffffffffffffff16611cfc9190614412565b92506101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506001600360009054906101000a900467ffffffffffffffff16611d459190614412565b92505050919050565b6000600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60606000600360009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff811115611db657611db5613df5565b5b604051908082528060200260200182016040528015611def57816020015b611ddc613591565b815260200190600190039081611dd45790505b509050600080600090505b600360009054906101000a900467ffffffffffffffff1667ffffffffffffffff168167ffffffffffffffff16101561207c578473ffffffffffffffffffffffffffffffffffffffff16600260008367ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff160361206957600260008267ffffffffffffffff1667ffffffffffffffff1681526020019081526020016000206040518060e00160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820160149054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff16815260200160038201548152602001600482015481525050838367ffffffffffffffff168151811061204e5761204d613e24565b5b60200260200101819052506001826120669190613e82565b91505b808061207490613ec0565b915050611dfa565b508067ffffffffffffffff1667ffffffffffffffff8111156120a1576120a0613df5565b5b6040519080825280602002602001820160405280156120da57816020015b6120c7613591565b8152602001906001900390816120bf5790505b50925060005b8167ffffffffffffffff168167ffffffffffffffff16101561215c57828167ffffffffffffffff168151811061211957612118613e24565b5b6020026020010151848267ffffffffffffffff168151811061213e5761213d613e24565b5b6020026020010181905250808061215490613ec0565b9150506120e0565b505050919050565b3373ffffffffffffffffffffffffffffffffffffffff166000808367ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060000160089054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461221b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612212906144b8565b60405180910390fd5b6000808267ffffffffffffffff1667ffffffffffffffff168152602001908152602001600020600080820160006101000a81549067ffffffffffffffff02191690556000820160086101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556001820160006101000a81549067ffffffffffffffff02191690556001820160086101000a81549067ffffffffffffffff02191690556001820160106101000a81549067ffffffffffffffff0219169055505050565b60006040518060a00160405280600160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1681526020013373ffffffffffffffffffffffffffffffffffffffff1681526020018567ffffffffffffffff1681526020018367ffffffffffffffff1681526020018467ffffffffffffffff16815250600080600160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060208201518160000160086101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160010160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060608201518160010160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060808201518160010160106101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055509050506001600081819054906101000a900467ffffffffffffffff16809291906124bf90613ec0565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055505060018060009054906101000a900467ffffffffffffffff166125089190614412565b90509392505050565b6000600860149054906101000a900467ffffffffffffffff16905090565b60006127108267ffffffffffffffff161061257f576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016125769061454a565b60405180910390fd5b600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461260f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612606906145dc565b60405180910390fd5b81600860146101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060009050919050565b6000600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b60606000600160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff8111156126c8576126c7613df5565b5b60405190808252806020026020018201604052801561270157816020015b6126ee613618565b8152602001906001900390816126e65790505b509050600080600090505b600160009054906101000a900467ffffffffffffffff1667ffffffffffffffff168167ffffffffffffffff1610156129285760008060008367ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060010160089054906101000a900467ffffffffffffffff1667ffffffffffffffff161115612915576000808267ffffffffffffffff1667ffffffffffffffff1681526020019081526020016000206040518060a00160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016001820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681525050838367ffffffffffffffff16815181106128fa576128f9613e24565b5b60200260200101819052506001826129129190613e82565b91505b808061292090613ec0565b91505061270c565b508067ffffffffffffffff1667ffffffffffffffff81111561294d5761294c613df5565b5b60405190808252806020026020018201604052801561298657816020015b612973613618565b81526020019060019003908161296b5790505b50925060005b8167ffffffffffffffff168167ffffffffffffffff161015612a0857828167ffffffffffffffff16815181106129c5576129c4613e24565b5b6020026020010151848267ffffffffffffffff16815181106129ea576129e9613e24565b5b60200260200101819052508080612a0090613ec0565b91505061298c565b50505090565b612a16613591565b600260008367ffffffffffffffff1667ffffffffffffffff1681526020019081526020016000206040518060e00160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820160149054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff168152602001600382015481526020016004820154815250509050919050565b600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614612c35576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612c2c9061466e565b60405180910390fd5b80600760006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60606000600160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff811115612cb757612cb6613df5565b5b604051908082528060200260200182016040528015612cf057816020015b612cdd613618565b815260200190600190039081612cd55790505b509050600080600090505b600160009054906101000a900467ffffffffffffffff1667ffffffffffffffff168167ffffffffffffffff161015612f43578473ffffffffffffffffffffffffffffffffffffffff166000808367ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060000160089054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603612f30576000808267ffffffffffffffff1667ffffffffffffffff1681526020019081526020016000206040518060a00160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016001820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681525050838367ffffffffffffffff1681518110612f1557612f14613e24565b5b6020026020010181905250600182612f2d9190613e82565b91505b8080612f3b90613ec0565b915050612cfb565b508067ffffffffffffffff1667ffffffffffffffff811115612f6857612f67613df5565b5b604051908082528060200260200182016040528015612fa157816020015b612f8e613618565b815260200190600190039081612f865790505b50925060005b8167ffffffffffffffff168167ffffffffffffffff16101561302357828167ffffffffffffffff1681518110612fe057612fdf613e24565b5b6020026020010151848267ffffffffffffffff168151811061300557613004613e24565b5b6020026020010181905250808061301b90613ec0565b915050612fa7565b505050919050565b62093a8081565b3373ffffffffffffffffffffffffffffffffffffffff16600260008367ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146130ea576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016130e190614700565b60405180910390fd5b6000600260008367ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060040154426131229190614177565b90506000600260008467ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060020160149054906101000a900467ffffffffffffffff1667ffffffffffffffff168261317a9190614720565b90508060056000600260008767ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054101561329a5760056000600260008667ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490505b6000612710600860149054906101000a900467ffffffffffffffff1667ffffffffffffffff16836132cb9190614720565b6132d591906147a9565b9050600081836132e59190614177565b905042600260008767ffffffffffffffff1667ffffffffffffffff168152602001908152602001600020600401819055508160056000600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546133879190614121565b925050819055508060056000600260008967ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546134279190614121565b925050819055508260056000600260008967ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546134c79190614177565b925050819055505050505050565b80600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555050565b600062093a80600660008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900467ffffffffffffffff166135809190614318565b67ffffffffffffffff169050919050565b6040518060e00160405280600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600067ffffffffffffffff16815260200160008152602001600081525090565b6040518060a00160405280600067ffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff1681525090565b6000819050919050565b61369881613685565b82525050565b60006020820190506136b3600083018461368f565b92915050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006136e9826136be565b9050919050565b6136f9816136de565b811461370457600080fd5b50565b600081359050613716816136f0565b92915050565b600060208284031215613732576137316136b9565b5b600061374084828501613707565b91505092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b600067ffffffffffffffff82169050919050565b61379281613775565b82525050565b6137a1816136de565b82525050565b6137b081613685565b82525050565b60e0820160008201516137cc6000850182613789565b5060208201516137df6020850182613789565b5060408201516137f26040850182613798565b5060608201516138056060850182613798565b5060808201516138186080850182613789565b5060a082015161382b60a08501826137a7565b5060c082015161383e60c08501826137a7565b50505050565b600061385083836137b6565b60e08301905092915050565b6000602082019050919050565b600061387482613749565b61387e8185613754565b935061388983613765565b8060005b838110156138ba5781516138a18882613844565b97506138ac8361385c565b92505060018101905061388d565b5085935050505092915050565b600060208201905081810360008301526138e18184613869565b905092915050565b6138f281613775565b81146138fd57600080fd5b50565b60008135905061390f816138e9565b92915050565b60008060006060848603121561392e5761392d6136b9565b5b600061393c86828701613900565b935050602061394d86828701613900565b925050604061395e86828701613900565b9150509250925092565b60006020828403121561397e5761397d6136b9565b5b600061398c84828501613900565b91505092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b60a0820160008201516139d76000850182613789565b5060208201516139ea6020850182613798565b5060408201516139fd6040850182613789565b506060820151613a106060850182613789565b506080820151613a236080850182613789565b50505050565b6000613a3583836139c1565b60a08301905092915050565b6000602082019050919050565b6000613a5982613995565b613a6381856139a0565b9350613a6e836139b1565b8060005b83811015613a9f578151613a868882613a29565b9750613a9183613a41565b925050600181019050613a72565b5085935050505092915050565b60006020820190508181036000830152613ac68184613a4e565b905092915050565b613ad781613775565b82525050565b6000602082019050613af26000830184613ace565b92915050565b613b01816136de565b82525050565b6000602082019050613b1c6000830184613af8565b92915050565b613b2b81613685565b8114613b3657600080fd5b50565b600081359050613b4881613b22565b92915050565b600060208284031215613b6457613b636136b9565b5b6000613b7284828501613b39565b91505092915050565b60008115159050919050565b613b9081613b7b565b82525050565b6000602082019050613bab6000830184613b87565b92915050565b6000604082019050613bc6600083018561368f565b613bd3602083018461368f565b9392505050565b6000819050919050565b6000613bff613bfa613bf5846136be565b613bda565b6136be565b9050919050565b6000613c1182613be4565b9050919050565b6000613c2382613c06565b9050919050565b613c3381613c18565b82525050565b6000602082019050613c4e6000830184613c2a565b92915050565b6000819050919050565b613c6781613c54565b82525050565b6000602082019050613c826000830184613c5e565b92915050565b60e082016000820151613c9e6000850182613789565b506020820151613cb16020850182613789565b506040820151613cc46040850182613798565b506060820151613cd76060850182613798565b506080820151613cea6080850182613789565b5060a0820151613cfd60a08501826137a7565b5060c0820151613d1060c08501826137a7565b50505050565b600060e082019050613d2b6000830184613c88565b92915050565b6000613d3c826136de565b9050919050565b613d4c81613d31565b8114613d5757600080fd5b50565b600081359050613d6981613d43565b92915050565b600060208284031215613d8557613d846136b9565b5b6000613d9384828501613d5a565b91505092915050565b613da581613c54565b8114613db057600080fd5b50565b600081359050613dc281613d9c565b92915050565b600060208284031215613dde57613ddd6136b9565b5b6000613dec84828501613db3565b91505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000613e8d82613775565b9150613e9883613775565b92508267ffffffffffffffff03821115613eb557613eb4613e53565b5b828201905092915050565b6000613ecb82613775565b915067ffffffffffffffff8203613ee557613ee4613e53565b5b600182019050919050565b600082825260208201905092915050565b7f4f6e6c7920746865206f776e65722063616e2075706461746520616e206f666660008201527f6572000000000000000000000000000000000000000000000000000000000000602082015250565b6000613f5d602283613ef0565b9150613f6882613f01565b604082019050919050565b60006020820190508181036000830152613f8c81613f50565b9050919050565b7f6f6e6c7920636f6d6d756e69747920636f6e74726163742063616e207365742060008201527f6e657720636f6d6d756e69747920636f6e747261637400000000000000000000602082015250565b6000613fef603683613ef0565b9150613ffa82613f93565b604082019050919050565b6000602082019050818103600083015261401e81613fe2565b9050919050565b600060608201905061403a6000830186613af8565b6140476020830185613af8565b614054604083018461368f565b949350505050565b61406581613b7b565b811461407057600080fd5b50565b6000815190506140828161405c565b92915050565b60006020828403121561409e5761409d6136b9565b5b60006140ac84828501614073565b91505092915050565b7f4661696c656420746f207472616e7366657220746f6b656e7300000000000000600082015250565b60006140eb601983613ef0565b91506140f6826140b5565b602082019050919050565b6000602082019050818103600083015261411a816140de565b9050919050565b600061412c82613685565b915061413783613685565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0382111561416c5761416b613e53565b5b828201905092915050565b600061418282613685565b915061418d83613685565b9250828210156141a05761419f613e53565b5b828203905092915050565b7f4e6f7420656e6f7567682062616c616e636520746f2077697468647261770000600082015250565b60006141e1601e83613ef0565b91506141ec826141ab565b602082019050919050565b60006020820190508181036000830152614210816141d4565b9050919050565b600060408201905061422c6000830185613af8565b614239602083018461368f565b9392505050565b7f4552433230207472616e73666572206661696c65640000000000000000000000600082015250565b6000614276601583613ef0565b915061428182614240565b602082019050919050565b600060208201905081810360008301526142a581614269565b9050919050565b7f4e6f206d616368696e657320617661696c61626c650000000000000000000000600082015250565b60006142e2601583613ef0565b91506142ed826142ac565b602082019050919050565b60006020820190508181036000830152614311816142d5565b9050919050565b600061432382613775565b915061432e83613775565b92508167ffffffffffffffff048311821515161561434f5761434e613e53565b5b828202905092915050565b7f596f7520646f6e2774206861766520656e6f7567682062616c616e636520746f60008201527f2070617920666f72207468697320616e6420616c6c206f7468657220564d732060208201527f666f722037206461797320000000000000000000000000000000000000000000604082015250565b60006143dc604b83613ef0565b91506143e78261435a565b606082019050919050565b6000602082019050818103600083015261440b816143cf565b9050919050565b600061441d82613775565b915061442883613775565b92508282101561443b5761443a613e53565b5b828203905092915050565b7f4f6e6c7920746865206f776e65722063616e2072656d6f766520616e206f666660008201527f6572000000000000000000000000000000000000000000000000000000000000602082015250565b60006144a2602283613ef0565b91506144ad82614446565b604082019050919050565b600060208201905081810360008301526144d181614495565b9050919050565b7f636f6d6d756e697479206665652073686f756c6420626520696e2072616e676560008201527f206f662030202830252920746f20313030303020283130302529000000000000602082015250565b6000614534603a83613ef0565b915061453f826144d8565b604082019050919050565b6000602082019050818103600083015261456381614527565b9050919050565b7f6f6e6c7920636f6d6d756e69747920636f6e74726163742063616e207365742060008201527f636f6d6d756e6974792066656500000000000000000000000000000000000000602082015250565b60006145c6602d83613ef0565b91506145d18261456a565b604082019050919050565b600060208201905081810360008301526145f5816145b9565b9050919050565b7f6f6e6c7920636f6d6d756e69747920636f6e74726163742063616e207365742060008201527f737461626c65636f696e00000000000000000000000000000000000000000000602082015250565b6000614658602a83613ef0565b9150614663826145fc565b604082019050919050565b600060208201905081810360008301526146878161464b565b9050919050565b7f4f6e6c7920746865206d696e65722063616e20636c61696d2061207061796d6560008201527f6e74000000000000000000000000000000000000000000000000000000000000602082015250565b60006146ea602283613ef0565b91506146f58261468e565b604082019050919050565b60006020820190508181036000830152614719816146dd565b9050919050565b600061472b82613685565b915061473683613685565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048311821515161561476f5761476e613e53565b5b828202905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b60006147b482613685565b91506147bf83613685565b9250826147cf576147ce61477a565b5b82820490509291505056fea2646970667358221220bfa13d3897b3d8be80aea60bd1e0099fa9dc43bcf0e31fd26c29fa0b8930dbec64736f6c634300080f0033",
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
// Solidity: function FindBookingsByMiner(address _miner) view returns((uint64,uint64,address,address,uint64,uint256,uint256)[] filteredBookings)
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
// Solidity: function FindBookingsByMiner(address _miner) view returns((uint64,uint64,address,address,uint64,uint256,uint256)[] filteredBookings)
func (_Broker *BrokerSession) FindBookingsByMiner(_miner common.Address) ([]BrokerBooking, error) {
	return _Broker.Contract.FindBookingsByMiner(&_Broker.CallOpts, _miner)
}

// FindBookingsByMiner is a free data retrieval call binding the contract method 0x366445bf.
//
// Solidity: function FindBookingsByMiner(address _miner) view returns((uint64,uint64,address,address,uint64,uint256,uint256)[] filteredBookings)
func (_Broker *BrokerCallerSession) FindBookingsByMiner(_miner common.Address) ([]BrokerBooking, error) {
	return _Broker.Contract.FindBookingsByMiner(&_Broker.CallOpts, _miner)
}

// FindBookingsByUser is a free data retrieval call binding the contract method 0x534b3ccb.
//
// Solidity: function FindBookingsByUser(address _owner) view returns((uint64,uint64,address,address,uint64,uint256,uint256)[] filteredBookings)
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
// Solidity: function FindBookingsByUser(address _owner) view returns((uint64,uint64,address,address,uint64,uint256,uint256)[] filteredBookings)
func (_Broker *BrokerSession) FindBookingsByUser(_owner common.Address) ([]BrokerBooking, error) {
	return _Broker.Contract.FindBookingsByUser(&_Broker.CallOpts, _owner)
}

// FindBookingsByUser is a free data retrieval call binding the contract method 0x534b3ccb.
//
// Solidity: function FindBookingsByUser(address _owner) view returns((uint64,uint64,address,address,uint64,uint256,uint256)[] filteredBookings)
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
// Solidity: function GetBooking(uint64 index) view returns((uint64,uint64,address,address,uint64,uint256,uint256) booking)
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
// Solidity: function GetBooking(uint64 index) view returns((uint64,uint64,address,address,uint64,uint256,uint256) booking)
func (_Broker *BrokerSession) GetBooking(index uint64) (BrokerBooking, error) {
	return _Broker.Contract.GetBooking(&_Broker.CallOpts, index)
}

// GetBooking is a free data retrieval call binding the contract method 0xc2161a05.
//
// Solidity: function GetBooking(uint64 index) view returns((uint64,uint64,address,address,uint64,uint256,uint256) booking)
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
// Solidity: function GetUsersBookings(address user) view returns((uint64,uint64,address,address,uint64,uint256,uint256)[] filteredBookings)
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
// Solidity: function GetUsersBookings(address user) view returns((uint64,uint64,address,address,uint64,uint256,uint256)[] filteredBookings)
func (_Broker *BrokerSession) GetUsersBookings(user common.Address) ([]BrokerBooking, error) {
	return _Broker.Contract.GetUsersBookings(&_Broker.CallOpts, user)
}

// GetUsersBookings is a free data retrieval call binding the contract method 0x7f7a1915.
//
// Solidity: function GetUsersBookings(address user) view returns((uint64,uint64,address,address,uint64,uint256,uint256)[] filteredBookings)
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

// BrokerBookingClaimedIterator is returned from FilterBookingClaimed and is used to iterate over the raw logs and unpacked data for BookingClaimed events raised by the Broker contract.
type BrokerBookingClaimedIterator struct {
	Event *BrokerBookingClaimed // Event containing the contract specifics and raw log

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
func (it *BrokerBookingClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BrokerBookingClaimed)
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
		it.Event = new(BrokerBookingClaimed)
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
func (it *BrokerBookingClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BrokerBookingClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BrokerBookingClaimed represents a BookingClaimed event raised by the Broker contract.
type BrokerBookingClaimed struct {
	MinerPayout uint64
	Index       uint64
	Miner       common.Address
	User        common.Address
	TimeUsed    uint64
	VmTypeId    uint64
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBookingClaimed is a free log retrieval operation binding the contract event 0x746851bdc47e0b3b82bedf2c554843f922fe045ab2b51b00ba83c33f4b76073d.
//
// Solidity: event BookingClaimed(uint64 minerPayout, uint64 index, address indexed miner, address indexed user, uint64 timeUsed, uint64 vmTypeId)
func (_Broker *BrokerFilterer) FilterBookingClaimed(opts *bind.FilterOpts, miner []common.Address, user []common.Address) (*BrokerBookingClaimedIterator, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Broker.contract.FilterLogs(opts, "BookingClaimed", minerRule, userRule)
	if err != nil {
		return nil, err
	}
	return &BrokerBookingClaimedIterator{contract: _Broker.contract, event: "BookingClaimed", logs: logs, sub: sub}, nil
}

// WatchBookingClaimed is a free log subscription operation binding the contract event 0x746851bdc47e0b3b82bedf2c554843f922fe045ab2b51b00ba83c33f4b76073d.
//
// Solidity: event BookingClaimed(uint64 minerPayout, uint64 index, address indexed miner, address indexed user, uint64 timeUsed, uint64 vmTypeId)
func (_Broker *BrokerFilterer) WatchBookingClaimed(opts *bind.WatchOpts, sink chan<- *BrokerBookingClaimed, miner []common.Address, user []common.Address) (event.Subscription, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Broker.contract.WatchLogs(opts, "BookingClaimed", minerRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BrokerBookingClaimed)
				if err := _Broker.contract.UnpackLog(event, "BookingClaimed", log); err != nil {
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

// ParseBookingClaimed is a log parse operation binding the contract event 0x746851bdc47e0b3b82bedf2c554843f922fe045ab2b51b00ba83c33f4b76073d.
//
// Solidity: event BookingClaimed(uint64 minerPayout, uint64 index, address indexed miner, address indexed user, uint64 timeUsed, uint64 vmTypeId)
func (_Broker *BrokerFilterer) ParseBookingClaimed(log types.Log) (*BrokerBookingClaimed, error) {
	event := new(BrokerBookingClaimed)
	if err := _Broker.contract.UnpackLog(event, "BookingClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BrokerBookingExtendedIterator is returned from FilterBookingExtended and is used to iterate over the raw logs and unpacked data for BookingExtended events raised by the Broker contract.
type BrokerBookingExtendedIterator struct {
	Event *BrokerBookingExtended // Event containing the contract specifics and raw log

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
func (it *BrokerBookingExtendedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BrokerBookingExtended)
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
		it.Event = new(BrokerBookingExtended)
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
func (it *BrokerBookingExtendedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BrokerBookingExtendedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BrokerBookingExtended represents a BookingExtended event raised by the Broker contract.
type BrokerBookingExtended struct {
	Index    uint64
	Miner    common.Address
	User     common.Address
	VmTypeId uint64
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBookingExtended is a free log retrieval operation binding the contract event 0x20244eb216168895873e86b6fd72b44c1f87c0709b129d6a35db75cd7d492062.
//
// Solidity: event BookingExtended(uint64 index, address indexed miner, address indexed user, uint64 vmTypeId)
func (_Broker *BrokerFilterer) FilterBookingExtended(opts *bind.FilterOpts, miner []common.Address, user []common.Address) (*BrokerBookingExtendedIterator, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Broker.contract.FilterLogs(opts, "BookingExtended", minerRule, userRule)
	if err != nil {
		return nil, err
	}
	return &BrokerBookingExtendedIterator{contract: _Broker.contract, event: "BookingExtended", logs: logs, sub: sub}, nil
}

// WatchBookingExtended is a free log subscription operation binding the contract event 0x20244eb216168895873e86b6fd72b44c1f87c0709b129d6a35db75cd7d492062.
//
// Solidity: event BookingExtended(uint64 index, address indexed miner, address indexed user, uint64 vmTypeId)
func (_Broker *BrokerFilterer) WatchBookingExtended(opts *bind.WatchOpts, sink chan<- *BrokerBookingExtended, miner []common.Address, user []common.Address) (event.Subscription, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Broker.contract.WatchLogs(opts, "BookingExtended", minerRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BrokerBookingExtended)
				if err := _Broker.contract.UnpackLog(event, "BookingExtended", log); err != nil {
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

// ParseBookingExtended is a log parse operation binding the contract event 0x20244eb216168895873e86b6fd72b44c1f87c0709b129d6a35db75cd7d492062.
//
// Solidity: event BookingExtended(uint64 index, address indexed miner, address indexed user, uint64 vmTypeId)
func (_Broker *BrokerFilterer) ParseBookingExtended(log types.Log) (*BrokerBookingExtended, error) {
	event := new(BrokerBookingExtended)
	if err := _Broker.contract.UnpackLog(event, "BookingExtended", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BrokerBookingReportedIterator is returned from FilterBookingReported and is used to iterate over the raw logs and unpacked data for BookingReported events raised by the Broker contract.
type BrokerBookingReportedIterator struct {
	Event *BrokerBookingReported // Event containing the contract specifics and raw log

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
func (it *BrokerBookingReportedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BrokerBookingReported)
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
		it.Event = new(BrokerBookingReported)
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
func (it *BrokerBookingReportedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BrokerBookingReportedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BrokerBookingReported represents a BookingReported event raised by the Broker contract.
type BrokerBookingReported struct {
	MinerPayout uint64
	Index       uint64
	Miner       common.Address
	User        common.Address
	TimeUsed    uint64
	VmTypeId    uint64
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBookingReported is a free log retrieval operation binding the contract event 0x0446ef521b810a98a051f6038ddd4cbec3a24402261df37e7258d69eae43773c.
//
// Solidity: event BookingReported(uint64 minerPayout, uint64 index, address indexed miner, address indexed user, uint64 timeUsed, uint64 vmTypeId)
func (_Broker *BrokerFilterer) FilterBookingReported(opts *bind.FilterOpts, miner []common.Address, user []common.Address) (*BrokerBookingReportedIterator, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Broker.contract.FilterLogs(opts, "BookingReported", minerRule, userRule)
	if err != nil {
		return nil, err
	}
	return &BrokerBookingReportedIterator{contract: _Broker.contract, event: "BookingReported", logs: logs, sub: sub}, nil
}

// WatchBookingReported is a free log subscription operation binding the contract event 0x0446ef521b810a98a051f6038ddd4cbec3a24402261df37e7258d69eae43773c.
//
// Solidity: event BookingReported(uint64 minerPayout, uint64 index, address indexed miner, address indexed user, uint64 timeUsed, uint64 vmTypeId)
func (_Broker *BrokerFilterer) WatchBookingReported(opts *bind.WatchOpts, sink chan<- *BrokerBookingReported, miner []common.Address, user []common.Address) (event.Subscription, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Broker.contract.WatchLogs(opts, "BookingReported", minerRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BrokerBookingReported)
				if err := _Broker.contract.UnpackLog(event, "BookingReported", log); err != nil {
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

// ParseBookingReported is a log parse operation binding the contract event 0x0446ef521b810a98a051f6038ddd4cbec3a24402261df37e7258d69eae43773c.
//
// Solidity: event BookingReported(uint64 minerPayout, uint64 index, address indexed miner, address indexed user, uint64 timeUsed, uint64 vmTypeId)
func (_Broker *BrokerFilterer) ParseBookingReported(log types.Log) (*BrokerBookingReported, error) {
	event := new(BrokerBookingReported)
	if err := _Broker.contract.UnpackLog(event, "BookingReported", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BrokerBookingStartedIterator is returned from FilterBookingStarted and is used to iterate over the raw logs and unpacked data for BookingStarted events raised by the Broker contract.
type BrokerBookingStartedIterator struct {
	Event *BrokerBookingStarted // Event containing the contract specifics and raw log

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
func (it *BrokerBookingStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BrokerBookingStarted)
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
		it.Event = new(BrokerBookingStarted)
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
func (it *BrokerBookingStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BrokerBookingStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BrokerBookingStarted represents a BookingStarted event raised by the Broker contract.
type BrokerBookingStarted struct {
	Index    uint64
	Miner    common.Address
	User     common.Address
	VmTypeId uint64
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBookingStarted is a free log retrieval operation binding the contract event 0x6034230bd8e825821683243dcdff2e45036039550fda19285f1083c8eefdf05e.
//
// Solidity: event BookingStarted(uint64 index, address indexed miner, address indexed user, uint64 vmTypeId)
func (_Broker *BrokerFilterer) FilterBookingStarted(opts *bind.FilterOpts, miner []common.Address, user []common.Address) (*BrokerBookingStartedIterator, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Broker.contract.FilterLogs(opts, "BookingStarted", minerRule, userRule)
	if err != nil {
		return nil, err
	}
	return &BrokerBookingStartedIterator{contract: _Broker.contract, event: "BookingStarted", logs: logs, sub: sub}, nil
}

// WatchBookingStarted is a free log subscription operation binding the contract event 0x6034230bd8e825821683243dcdff2e45036039550fda19285f1083c8eefdf05e.
//
// Solidity: event BookingStarted(uint64 index, address indexed miner, address indexed user, uint64 vmTypeId)
func (_Broker *BrokerFilterer) WatchBookingStarted(opts *bind.WatchOpts, sink chan<- *BrokerBookingStarted, miner []common.Address, user []common.Address) (event.Subscription, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Broker.contract.WatchLogs(opts, "BookingStarted", minerRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BrokerBookingStarted)
				if err := _Broker.contract.UnpackLog(event, "BookingStarted", log); err != nil {
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

// ParseBookingStarted is a log parse operation binding the contract event 0x6034230bd8e825821683243dcdff2e45036039550fda19285f1083c8eefdf05e.
//
// Solidity: event BookingStarted(uint64 index, address indexed miner, address indexed user, uint64 vmTypeId)
func (_Broker *BrokerFilterer) ParseBookingStarted(log types.Log) (*BrokerBookingStarted, error) {
	event := new(BrokerBookingStarted)
	if err := _Broker.contract.UnpackLog(event, "BookingStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BrokerBookingStoppedIterator is returned from FilterBookingStopped and is used to iterate over the raw logs and unpacked data for BookingStopped events raised by the Broker contract.
type BrokerBookingStoppedIterator struct {
	Event *BrokerBookingStopped // Event containing the contract specifics and raw log

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
func (it *BrokerBookingStoppedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BrokerBookingStopped)
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
		it.Event = new(BrokerBookingStopped)
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
func (it *BrokerBookingStoppedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BrokerBookingStoppedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BrokerBookingStopped represents a BookingStopped event raised by the Broker contract.
type BrokerBookingStopped struct {
	MinerPayout uint64
	Index       uint64
	Miner       common.Address
	User        common.Address
	TimeUsed    uint64
	VmTypeId    uint64
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBookingStopped is a free log retrieval operation binding the contract event 0x07761517f60030901f4220da817be81d470b14927f180f9118412191c066dd02.
//
// Solidity: event BookingStopped(uint64 minerPayout, uint64 index, address indexed miner, address indexed user, uint64 timeUsed, uint64 vmTypeId)
func (_Broker *BrokerFilterer) FilterBookingStopped(opts *bind.FilterOpts, miner []common.Address, user []common.Address) (*BrokerBookingStoppedIterator, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Broker.contract.FilterLogs(opts, "BookingStopped", minerRule, userRule)
	if err != nil {
		return nil, err
	}
	return &BrokerBookingStoppedIterator{contract: _Broker.contract, event: "BookingStopped", logs: logs, sub: sub}, nil
}

// WatchBookingStopped is a free log subscription operation binding the contract event 0x07761517f60030901f4220da817be81d470b14927f180f9118412191c066dd02.
//
// Solidity: event BookingStopped(uint64 minerPayout, uint64 index, address indexed miner, address indexed user, uint64 timeUsed, uint64 vmTypeId)
func (_Broker *BrokerFilterer) WatchBookingStopped(opts *bind.WatchOpts, sink chan<- *BrokerBookingStopped, miner []common.Address, user []common.Address) (event.Subscription, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Broker.contract.WatchLogs(opts, "BookingStopped", minerRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BrokerBookingStopped)
				if err := _Broker.contract.UnpackLog(event, "BookingStopped", log); err != nil {
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

// ParseBookingStopped is a log parse operation binding the contract event 0x07761517f60030901f4220da817be81d470b14927f180f9118412191c066dd02.
//
// Solidity: event BookingStopped(uint64 minerPayout, uint64 index, address indexed miner, address indexed user, uint64 timeUsed, uint64 vmTypeId)
func (_Broker *BrokerFilterer) ParseBookingStopped(log types.Log) (*BrokerBookingStopped, error) {
	event := new(BrokerBookingStopped)
	if err := _Broker.contract.UnpackLog(event, "BookingStopped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
