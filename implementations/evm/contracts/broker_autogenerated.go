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
	Bin: "0x60806040526000600160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506000600360006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506101f4600860146101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555034801561008f57600080fd5b5033600860006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550614f13806100e06000396000f3fe608060405234801561001057600080fd5b50600436106101a95760003560e01c80637f7a1915116100f9578063b65c4b9c11610097578063d3554ab511610071578063d3554ab51461052f578063e088b29e1461055f578063f3fe369c1461057d578063f44f11cd14610599576101a9565b8063b65c4b9c146104c5578063c2161a05146104e3578063c9b16fee14610513576101a9565b806390ad688b116100d357806390ad688b1461042b5780639a61a20e14610449578063a7bf757014610465578063ab49ec4514610495576101a9565b80637f7a1915146103af57806388317656146103df5780638f351016146103fb576101a9565b8063569e2e0f116101665780636f427a80116101405780636f427a80146103005780637559fe621461033057806375b7322e146103615780637bef4f5414610391576101a9565b8063569e2e0f146102965780636b620ed5146102b45780636d8351d6146102d0576101a9565b806318c5e502146101ae578063366445bf146101cc5780633fec2ddf146101fc578063534b3ccb146102185780635405b5e51461024857806354fd4d5014610278575b600080fd5b6101b66105b5565b6040516101c39190613c92565b60405180910390f35b6101e660048036038101906101e19190613d10565b6105bd565b6040516101f39190613ed0565b60405180910390f35b61021660048036038101906102119190613f1e565b6109dc565b005b610232600480360381019061022d9190613d10565b610b38565b60405161023f9190613ed0565b60405180910390f35b610262600480360381019061025d9190613f71565b610f57565b60405161026f91906140b5565b60405180910390f35b610280611339565b60405161028d91906140e6565b60405180910390f35b61029e61133e565b6040516102ab9190614110565b60405180910390f35b6102ce60048036038101906102c99190613d10565b611368565b005b6102ea60048036038101906102e59190614157565b61143c565b6040516102f7919061419f565b60405180910390f35b61031a60048036038101906103159190614157565b6115b6565b604051610327919061419f565b60405180910390f35b61034a60048036038101906103459190613d10565b611790565b6040516103589291906141ba565b60405180910390f35b61037b60048036038101906103769190613f71565b6117f6565b60405161038891906140e6565b60405180910390f35b610399611e2b565b6040516103a69190614242565b60405180910390f35b6103c960048036038101906103c49190613d10565b611e55565b6040516103d69190613ed0565b60405180910390f35b6103f960048036038101906103f49190613f71565b612274565b005b61041560048036038101906104109190613f1e565b6123ea565b60405161042291906140e6565b60405180910390f35b610433612621565b60405161044091906140e6565b60405180910390f35b610463600480360381019061045e9190614296565b61263f565b005b61047f600480360381019061047a9190613f71565b61280c565b60405161048c919061419f565b60405180910390f35b6104af60048036038101906104aa9190613d10565b61291e565b6040516104bc91906142ef565b60405180910390f35b6104cd612967565b6040516104da91906140b5565b60405180910390f35b6104fd60048036038101906104f89190613f71565b612ceb565b60405161050a91906143ac565b60405180910390f35b61052d60048036038101906105289190614406565b612eb5565b005b61054960048036038101906105449190613d10565b612f89565b60405161055691906140b5565b60405180910390f35b61056761333b565b60405161057491906140e6565b60405180910390f35b61059760048036038101906105929190613f71565b613342565b005b6105b360048036038101906105ae919061445f565b61341a565b005b600042905090565b60606000600360009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff8111156105fb576105fa61448c565b5b60405190808252806020026020018201604052801561063457816020015b610621613b73565b8152602001906001900390816106195790505b509050600080600090505b600360009054906101000a900467ffffffffffffffff1667ffffffffffffffff168167ffffffffffffffff1610156108f4578473ffffffffffffffffffffffffffffffffffffffff16600260008367ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16036108e157600260008267ffffffffffffffff1667ffffffffffffffff168152602001908152602001600020604051806101000160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820160149054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff16815260200160038201548152602001600482015481526020016005820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681525050838367ffffffffffffffff16815181106108c6576108c56144bb565b5b60200260200101819052506001826108de9190614519565b91505b80806108ec90614557565b91505061063f565b508067ffffffffffffffff1667ffffffffffffffff8111156109195761091861448c565b5b60405190808252806020026020018201604052801561095257816020015b61093f613b73565b8152602001906001900390816109375790505b50925060005b8167ffffffffffffffff168167ffffffffffffffff1610156109d457828167ffffffffffffffff1681518110610991576109906144bb565b5b6020026020010151848267ffffffffffffffff16815181106109b6576109b56144bb565b5b602002602001018190525080806109cc90614557565b915050610958565b505050919050565b3373ffffffffffffffffffffffffffffffffffffffff166000808567ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060000160089054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610a93576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a8a9061460a565b60405180910390fd5b816000808567ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060010160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550806000808567ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060010160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550505050565b60606000600360009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff811115610b7657610b7561448c565b5b604051908082528060200260200182016040528015610baf57816020015b610b9c613b73565b815260200190600190039081610b945790505b509050600080600090505b600360009054906101000a900467ffffffffffffffff1667ffffffffffffffff168167ffffffffffffffff161015610e6f578473ffffffffffffffffffffffffffffffffffffffff16600260008367ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603610e5c57600260008267ffffffffffffffff1667ffffffffffffffff168152602001908152602001600020604051806101000160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820160149054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff16815260200160038201548152602001600482015481526020016005820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681525050838367ffffffffffffffff1681518110610e4157610e406144bb565b5b6020026020010181905250600182610e599190614519565b91505b8080610e6790614557565b915050610bba565b508067ffffffffffffffff1667ffffffffffffffff811115610e9457610e9361448c565b5b604051908082528060200260200182016040528015610ecd57816020015b610eba613b73565b815260200190600190039081610eb25790505b50925060005b8167ffffffffffffffff168167ffffffffffffffff161015610f4f57828167ffffffffffffffff1681518110610f0c57610f0b6144bb565b5b6020026020010151848267ffffffffffffffff1681518110610f3157610f306144bb565b5b60200260200101819052508080610f4790614557565b915050610ed3565b505050919050565b60606000600160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff811115610f9557610f9461448c565b5b604051908082528060200260200182016040528015610fce57816020015b610fbb613c0c565b815260200190600190039081610fb35790505b509050600080600090505b600160009054906101000a900467ffffffffffffffff1667ffffffffffffffff168167ffffffffffffffff161015611251578467ffffffffffffffff166000808367ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060010160109054906101000a900467ffffffffffffffff1667ffffffffffffffff161480156110b2575060008060008367ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060010160089054906101000a900467ffffffffffffffff1667ffffffffffffffff16115b1561123e576000808267ffffffffffffffff1667ffffffffffffffff1681526020019081526020016000206040518060a00160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016001820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681525050838367ffffffffffffffff1681518110611223576112226144bb565b5b602002602001018190525060018261123b9190614519565b91505b808061124990614557565b915050610fd9565b508067ffffffffffffffff1667ffffffffffffffff8111156112765761127561448c565b5b6040519080825280602002602001820160405280156112af57816020015b61129c613c0c565b8152602001906001900390816112945790505b50925060005b8167ffffffffffffffff168167ffffffffffffffff16101561133157828167ffffffffffffffff16815181106112ee576112ed6144bb565b5b6020026020010151848267ffffffffffffffff1681518110611313576113126144bb565b5b6020026020010181905250808061132990614557565b9150506112b5565b505050919050565b600181565b6000600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146113f8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016113ef9061469c565b60405180910390fd5b80600860006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6000600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd3330856040518463ffffffff1660e01b815260040161149d939291906146bc565b6020604051808303816000875af11580156114bc573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114e0919061471f565b61151f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161151690614798565b60405180910390fd5b81600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461156a91906147b8565b600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555060019050919050565b6000806115c233613461565b600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461160c919061480e565b905082811015611651576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016116489061488e565b60405180910390fd5b600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb33856040518363ffffffff1660e01b81526004016116ae9291906148ae565b6020604051808303816000875af11580156116cd573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116f1919061471f565b611730576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161172790614923565b60405180910390fd5b82600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461177f919061480e565b925050819055506001915050919050565b600080600061179e84613461565b905080600560008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546117eb919061480e565b819250925050915091565b6000806000808467ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060010160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1611611881576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016118789061498f565b60405180910390fd5b600062093a806000808567ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060010160009054906101000a900467ffffffffffffffff166118cf91906149af565b67ffffffffffffffff166118e233613461565b6118ec91906147b8565b9050600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054811115611970576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161196790614a89565b60405180910390fd5b6000604051806101000160405280600360009054906101000a900467ffffffffffffffff1667ffffffffffffffff1681526020016000808767ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060010160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1681526020016000808767ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060000160089054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020013373ffffffffffffffffffffffffffffffffffffffff1681526020016000808767ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060010160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1681526020014281526020014281526020018567ffffffffffffffff1681525090508060026000600360009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060208201518160000160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060408201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060808201518160020160146101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060a0820151816003015560c0820151816004015560e08201518160050160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055509050506003600081819054906101000a900467ffffffffffffffff1680929190611ca090614557565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550506000808567ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060010160009054906101000a900467ffffffffffffffff16600660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282829054906101000a900467ffffffffffffffff16611d669190614519565b92506101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060016000808667ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060010160088282829054906101000a900467ffffffffffffffff16611dd99190614aa9565b92506101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506001600360009054906101000a900467ffffffffffffffff16611e229190614aa9565b92505050919050565b6000600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60606000600360009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff811115611e9357611e9261448c565b5b604051908082528060200260200182016040528015611ecc57816020015b611eb9613b73565b815260200190600190039081611eb15790505b509050600080600090505b600360009054906101000a900467ffffffffffffffff1667ffffffffffffffff168167ffffffffffffffff16101561218c578473ffffffffffffffffffffffffffffffffffffffff16600260008367ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff160361217957600260008267ffffffffffffffff1667ffffffffffffffff168152602001908152602001600020604051806101000160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820160149054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff16815260200160038201548152602001600482015481526020016005820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681525050838367ffffffffffffffff168151811061215e5761215d6144bb565b5b60200260200101819052506001826121769190614519565b91505b808061218490614557565b915050611ed7565b508067ffffffffffffffff1667ffffffffffffffff8111156121b1576121b061448c565b5b6040519080825280602002602001820160405280156121ea57816020015b6121d7613b73565b8152602001906001900390816121cf5790505b50925060005b8167ffffffffffffffff168167ffffffffffffffff16101561226c57828167ffffffffffffffff1681518110612229576122286144bb565b5b6020026020010151848267ffffffffffffffff168151811061224e5761224d6144bb565b5b6020026020010181905250808061226490614557565b9150506121f0565b505050919050565b3373ffffffffffffffffffffffffffffffffffffffff166000808367ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060000160089054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461232b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161232290614b4f565b60405180910390fd5b6000808267ffffffffffffffff1667ffffffffffffffff168152602001908152602001600020600080820160006101000a81549067ffffffffffffffff02191690556000820160086101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556001820160006101000a81549067ffffffffffffffff02191690556001820160086101000a81549067ffffffffffffffff02191690556001820160106101000a81549067ffffffffffffffff0219169055505050565b60006040518060a00160405280600160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1681526020013373ffffffffffffffffffffffffffffffffffffffff1681526020018567ffffffffffffffff1681526020018367ffffffffffffffff1681526020018467ffffffffffffffff16815250600080600160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060208201518160000160086101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160010160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060608201518160010160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060808201518160010160106101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055509050506001600081819054906101000a900467ffffffffffffffff16809291906125cf90614557565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055505060018060009054906101000a900467ffffffffffffffff166126189190614aa9565b90509392505050565b6000600860149054906101000a900467ffffffffffffffff16905090565b3373ffffffffffffffffffffffffffffffffffffffff16600260008467ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146126f7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016126ee90614bbb565b60405180910390fd5b60008160ff16146127f5578060ff16600260008467ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16600260008567ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f1db97daa25845ee925bd84e6cf559402b161e27653e3f3187c1859f98634c1c560405160405180910390a45b6127fe826134d6565b50612808826139cb565b5050565b60006127108267ffffffffffffffff161061285c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161285390614c4d565b60405180910390fd5b600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146128ec576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016128e390614cdf565b60405180910390fd5b81600860146101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060009050919050565b6000600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b60606000600160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff8111156129a5576129a461448c565b5b6040519080825280602002602001820160405280156129de57816020015b6129cb613c0c565b8152602001906001900390816129c35790505b509050600080600090505b600160009054906101000a900467ffffffffffffffff1667ffffffffffffffff168167ffffffffffffffff161015612c055760008060008367ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060010160089054906101000a900467ffffffffffffffff1667ffffffffffffffff161115612bf2576000808267ffffffffffffffff1667ffffffffffffffff1681526020019081526020016000206040518060a00160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016001820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681525050838367ffffffffffffffff1681518110612bd757612bd66144bb565b5b6020026020010181905250600182612bef9190614519565b91505b8080612bfd90614557565b9150506129e9565b508067ffffffffffffffff1667ffffffffffffffff811115612c2a57612c2961448c565b5b604051908082528060200260200182016040528015612c6357816020015b612c50613c0c565b815260200190600190039081612c485790505b50925060005b8167ffffffffffffffff168167ffffffffffffffff161015612ce557828167ffffffffffffffff1681518110612ca257612ca16144bb565b5b6020026020010151848267ffffffffffffffff1681518110612cc757612cc66144bb565b5b60200260200101819052508080612cdd90614557565b915050612c69565b50505090565b612cf3613b73565b600260008367ffffffffffffffff1667ffffffffffffffff168152602001908152602001600020604051806101000160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820160149054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff16815260200160038201548152602001600482015481526020016005820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff16815250509050919050565b600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614612f45576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612f3c90614d71565b60405180910390fd5b80600760006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60606000600160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff811115612fc757612fc661448c565b5b60405190808252806020026020018201604052801561300057816020015b612fed613c0c565b815260200190600190039081612fe55790505b509050600080600090505b600160009054906101000a900467ffffffffffffffff1667ffffffffffffffff168167ffffffffffffffff161015613253578473ffffffffffffffffffffffffffffffffffffffff166000808367ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060000160089054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603613240576000808267ffffffffffffffff1667ffffffffffffffff1681526020019081526020016000206040518060a00160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016001820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681525050838367ffffffffffffffff1681518110613225576132246144bb565b5b602002602001018190525060018261323d9190614519565b91505b808061324b90614557565b91505061300b565b508067ffffffffffffffff1667ffffffffffffffff8111156132785761327761448c565b5b6040519080825280602002602001820160405280156132b157816020015b61329e613c0c565b8152602001906001900390816132965790505b50925060005b8167ffffffffffffffff168167ffffffffffffffff16101561333357828167ffffffffffffffff16815181106132f0576132ef6144bb565b5b6020026020010151848267ffffffffffffffff1681518110613315576133146144bb565b5b6020026020010181905250808061332b90614557565b9150506132b7565b505050919050565b62093a8081565b3373ffffffffffffffffffffffffffffffffffffffff16600260008367ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146133fa576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016133f190614e03565b60405180910390fd5b6000613405826134d6565b90508061341657613415826139cb565b5b5050565b80600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555050565b600062093a80600660008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900467ffffffffffffffff166134c591906149af565b67ffffffffffffffff169050919050565b600080600190506000600260008567ffffffffffffffff1667ffffffffffffffff1681526020019081526020016000206004015442613515919061480e565b90506000600260008667ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060020160149054906101000a900467ffffffffffffffff1667ffffffffffffffff168261356d9190614e23565b90508060056000600260008967ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205410156136915760056000600260008867ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050600092505b6000612710600860149054906101000a900467ffffffffffffffff1667ffffffffffffffff16836136c29190614e23565b6136cc9190614eac565b9050600081836136dc919061480e565b905042600260008967ffffffffffffffff1667ffffffffffffffff168152602001908152602001600020600401819055508160056000600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461377e91906147b8565b925050819055508060056000600260008b67ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461381e91906147b8565b925050819055508260056000600260008b67ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546138be919061480e565b92505081905550600260008867ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16600260008967ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8ed7db30df47fbd87811a5e8a95a94838f0c0241263d9a1865d6a2a3e10516e2836040516139b69190613c92565b60405180910390a38495505050505050919050565b6001600080600260008567ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060050160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060010160088282829054906101000a900467ffffffffffffffff16613a569190614519565b92506101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550600260008267ffffffffffffffff1667ffffffffffffffff168152602001908152602001600020600080820160006101000a81549067ffffffffffffffff02191690556000820160086101000a81549067ffffffffffffffff02191690556001820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556002820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556002820160146101000a81549067ffffffffffffffff0219169055600382016000905560048201600090556005820160006101000a81549067ffffffffffffffff0219169055505050565b604051806101000160405280600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600067ffffffffffffffff1681526020016000815260200160008152602001600067ffffffffffffffff1681525090565b6040518060a00160405280600067ffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff1681525090565b6000819050919050565b613c8c81613c79565b82525050565b6000602082019050613ca76000830184613c83565b92915050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000613cdd82613cb2565b9050919050565b613ced81613cd2565b8114613cf857600080fd5b50565b600081359050613d0a81613ce4565b92915050565b600060208284031215613d2657613d25613cad565b5b6000613d3484828501613cfb565b91505092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b600067ffffffffffffffff82169050919050565b613d8681613d69565b82525050565b613d9581613cd2565b82525050565b613da481613c79565b82525050565b61010082016000820151613dc16000850182613d7d565b506020820151613dd46020850182613d7d565b506040820151613de76040850182613d8c565b506060820151613dfa6060850182613d8c565b506080820151613e0d6080850182613d7d565b5060a0820151613e2060a0850182613d9b565b5060c0820151613e3360c0850182613d9b565b5060e0820151613e4660e0850182613d7d565b50505050565b6000613e588383613daa565b6101008301905092915050565b6000602082019050919050565b6000613e7d82613d3d565b613e878185613d48565b9350613e9283613d59565b8060005b83811015613ec3578151613eaa8882613e4c565b9750613eb583613e65565b925050600181019050613e96565b5085935050505092915050565b60006020820190508181036000830152613eea8184613e72565b905092915050565b613efb81613d69565b8114613f0657600080fd5b50565b600081359050613f1881613ef2565b92915050565b600080600060608486031215613f3757613f36613cad565b5b6000613f4586828701613f09565b9350506020613f5686828701613f09565b9250506040613f6786828701613f09565b9150509250925092565b600060208284031215613f8757613f86613cad565b5b6000613f9584828501613f09565b91505092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b60a082016000820151613fe06000850182613d7d565b506020820151613ff36020850182613d8c565b5060408201516140066040850182613d7d565b5060608201516140196060850182613d7d565b50608082015161402c6080850182613d7d565b50505050565b600061403e8383613fca565b60a08301905092915050565b6000602082019050919050565b600061406282613f9e565b61406c8185613fa9565b935061407783613fba565b8060005b838110156140a857815161408f8882614032565b975061409a8361404a565b92505060018101905061407b565b5085935050505092915050565b600060208201905081810360008301526140cf8184614057565b905092915050565b6140e081613d69565b82525050565b60006020820190506140fb60008301846140d7565b92915050565b61410a81613cd2565b82525050565b60006020820190506141256000830184614101565b92915050565b61413481613c79565b811461413f57600080fd5b50565b6000813590506141518161412b565b92915050565b60006020828403121561416d5761416c613cad565b5b600061417b84828501614142565b91505092915050565b60008115159050919050565b61419981614184565b82525050565b60006020820190506141b46000830184614190565b92915050565b60006040820190506141cf6000830185613c83565b6141dc6020830184613c83565b9392505050565b6000819050919050565b60006142086142036141fe84613cb2565b6141e3565b613cb2565b9050919050565b600061421a826141ed565b9050919050565b600061422c8261420f565b9050919050565b61423c81614221565b82525050565b60006020820190506142576000830184614233565b92915050565b600060ff82169050919050565b6142738161425d565b811461427e57600080fd5b50565b6000813590506142908161426a565b92915050565b600080604083850312156142ad576142ac613cad565b5b60006142bb85828601613f09565b92505060206142cc85828601614281565b9150509250929050565b6000819050919050565b6142e9816142d6565b82525050565b600060208201905061430460008301846142e0565b92915050565b610100820160008201516143216000850182613d7d565b5060208201516143346020850182613d7d565b5060408201516143476040850182613d8c565b50606082015161435a6060850182613d8c565b50608082015161436d6080850182613d7d565b5060a082015161438060a0850182613d9b565b5060c082015161439360c0850182613d9b565b5060e08201516143a660e0850182613d7d565b50505050565b6000610100820190506143c2600083018461430a565b92915050565b60006143d382613cd2565b9050919050565b6143e3816143c8565b81146143ee57600080fd5b50565b600081359050614400816143da565b92915050565b60006020828403121561441c5761441b613cad565b5b600061442a848285016143f1565b91505092915050565b61443c816142d6565b811461444757600080fd5b50565b60008135905061445981614433565b92915050565b60006020828403121561447557614474613cad565b5b60006144838482850161444a565b91505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061452482613d69565b915061452f83613d69565b92508267ffffffffffffffff0382111561454c5761454b6144ea565b5b828201905092915050565b600061456282613d69565b915067ffffffffffffffff820361457c5761457b6144ea565b5b600182019050919050565b600082825260208201905092915050565b7f4f6e6c7920746865206f776e65722063616e2075706461746520616e206f666660008201527f6572000000000000000000000000000000000000000000000000000000000000602082015250565b60006145f4602283614587565b91506145ff82614598565b604082019050919050565b60006020820190508181036000830152614623816145e7565b9050919050565b7f6f6e6c7920636f6d6d756e69747920636f6e74726163742063616e207365742060008201527f6e657720636f6d6d756e69747920636f6e747261637400000000000000000000602082015250565b6000614686603683614587565b91506146918261462a565b604082019050919050565b600060208201905081810360008301526146b581614679565b9050919050565b60006060820190506146d16000830186614101565b6146de6020830185614101565b6146eb6040830184613c83565b949350505050565b6146fc81614184565b811461470757600080fd5b50565b600081519050614719816146f3565b92915050565b60006020828403121561473557614734613cad565b5b60006147438482850161470a565b91505092915050565b7f4661696c656420746f207472616e7366657220746f6b656e7300000000000000600082015250565b6000614782601983614587565b915061478d8261474c565b602082019050919050565b600060208201905081810360008301526147b181614775565b9050919050565b60006147c382613c79565b91506147ce83613c79565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115614803576148026144ea565b5b828201905092915050565b600061481982613c79565b915061482483613c79565b925082821015614837576148366144ea565b5b828203905092915050565b7f4e6f7420656e6f7567682062616c616e636520746f2077697468647261770000600082015250565b6000614878601e83614587565b915061488382614842565b602082019050919050565b600060208201905081810360008301526148a78161486b565b9050919050565b60006040820190506148c36000830185614101565b6148d06020830184613c83565b9392505050565b7f4552433230207472616e73666572206661696c65640000000000000000000000600082015250565b600061490d601583614587565b9150614918826148d7565b602082019050919050565b6000602082019050818103600083015261493c81614900565b9050919050565b7f4e6f206d616368696e657320617661696c61626c650000000000000000000000600082015250565b6000614979601583614587565b915061498482614943565b602082019050919050565b600060208201905081810360008301526149a88161496c565b9050919050565b60006149ba82613d69565b91506149c583613d69565b92508167ffffffffffffffff04831182151516156149e6576149e56144ea565b5b828202905092915050565b7f596f7520646f6e2774206861766520656e6f7567682062616c616e636520746f60008201527f2070617920666f72207468697320616e6420616c6c206f7468657220564d732060208201527f666f722037206461797320000000000000000000000000000000000000000000604082015250565b6000614a73604b83614587565b9150614a7e826149f1565b606082019050919050565b60006020820190508181036000830152614aa281614a66565b9050919050565b6000614ab482613d69565b9150614abf83613d69565b925082821015614ad257614ad16144ea565b5b828203905092915050565b7f4f6e6c7920746865206f776e65722063616e2072656d6f766520616e206f666660008201527f6572000000000000000000000000000000000000000000000000000000000000602082015250565b6000614b39602283614587565b9150614b4482614add565b604082019050919050565b60006020820190508181036000830152614b6881614b2c565b9050919050565b7f4f6e6c792074686520757365722063616e2073746f70206120564d0000000000600082015250565b6000614ba5601b83614587565b9150614bb082614b6f565b602082019050919050565b60006020820190508181036000830152614bd481614b98565b9050919050565b7f636f6d6d756e697479206665652073686f756c6420626520696e2072616e676560008201527f206f662030202830252920746f20313030303020283130302529000000000000602082015250565b6000614c37603a83614587565b9150614c4282614bdb565b604082019050919050565b60006020820190508181036000830152614c6681614c2a565b9050919050565b7f6f6e6c7920636f6d6d756e69747920636f6e74726163742063616e207365742060008201527f636f6d6d756e6974792066656500000000000000000000000000000000000000602082015250565b6000614cc9602d83614587565b9150614cd482614c6d565b604082019050919050565b60006020820190508181036000830152614cf881614cbc565b9050919050565b7f6f6e6c7920636f6d6d756e69747920636f6e74726163742063616e207365742060008201527f737461626c65636f696e00000000000000000000000000000000000000000000602082015250565b6000614d5b602a83614587565b9150614d6682614cff565b604082019050919050565b60006020820190508181036000830152614d8a81614d4e565b9050919050565b7f4f6e6c7920746865206d696e65722063616e20636c61696d2061207061796d6560008201527f6e74000000000000000000000000000000000000000000000000000000000000602082015250565b6000614ded602283614587565b9150614df882614d91565b604082019050919050565b60006020820190508181036000830152614e1c81614de0565b9050919050565b6000614e2e82613c79565b9150614e3983613c79565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615614e7257614e716144ea565b5b828202905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b6000614eb782613c79565b9150614ec283613c79565b925082614ed257614ed1614e7d565b5b82820490509291505056fea264697066735822122070c428489d39c969541b9c13b23fc57c3465194393b51c90a87b9ce87471820464736f6c634300080f0033",
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
