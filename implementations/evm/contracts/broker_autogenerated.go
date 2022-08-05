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
	Index          *big.Int
	VmTypeId       *big.Int
	Miner          common.Address
	User           common.Address
	PricePerSecond *big.Int
	BookedAt       *big.Int
	BookedTill     *big.Int
}

// BrokerVMOffer is an auto generated low-level Go binding around an user-defined struct.
type BrokerVMOffer struct {
	Index             *big.Int
	Miner             common.Address
	PricePerSecond    *big.Int
	MachinesAvailable *big.Int
	VmTypeId          *big.Int
}

// BrokerMetaData contains all meta data concerning the Broker contract.
var BrokerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"communityAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minerPayout\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timeUsed\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vmTypeId\",\"type\":\"uint256\"}],\"name\":\"BookingClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vmTypeId\",\"type\":\"uint256\"}],\"name\":\"BookingExtended\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minerPayout\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timeUsed\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vmTypeId\",\"type\":\"uint256\"}],\"name\":\"BookingReported\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vmTypeId\",\"type\":\"uint256\"}],\"name\":\"BookingStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minerPayout\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timeUsed\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vmTypeId\",\"type\":\"uint256\"}],\"name\":\"BookingStopped\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"GetTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offerIndex\",\"type\":\"uint256\"}],\"name\":\"RemoveOffer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offerIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"machinesAvailable\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pps\",\"type\":\"uint256\"}],\"name\":\"UpdateOffer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"abortType\",\"type\":\"uint8\"}],\"name\":\"abortBooking\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pricePerSecond\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vmTypeId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"machinesAvailable\",\"type\":\"uint256\"}],\"name\":\"addOffer\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offerIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"secs\",\"type\":\"uint256\"}],\"name\":\"bookVM\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"}],\"name\":\"claimExpired\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"numTokens\",\"type\":\"uint256\"}],\"name\":\"depositCoin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"secs\",\"type\":\"uint256\"}],\"name\":\"extendBooking\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_miner\",\"type\":\"address\"}],\"name\":\"findBookingsByMiner\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vmTypeId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pricePerSecond\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bookedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bookedTill\",\"type\":\"uint256\"}],\"internalType\":\"structBroker.Booking[]\",\"name\":\"filteredBookings\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"findBookingsByUser\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vmTypeId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pricePerSecond\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bookedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bookedTill\",\"type\":\"uint256\"}],\"internalType\":\"structBroker.Booking[]\",\"name\":\"filteredBookings\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_vmTypeId\",\"type\":\"uint256\"}],\"name\":\"getAvailableOffers\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pricePerSecond\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"machinesAvailable\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vmTypeId\",\"type\":\"uint256\"}],\"internalType\":\"structBroker.VMOffer[]\",\"name\":\"filteredOffers\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"}],\"name\":\"getBooking\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vmTypeId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pricePerSecond\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bookedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bookedTill\",\"type\":\"uint256\"}],\"internalType\":\"structBroker.Booking\",\"name\":\"booking\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCoinDecimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCommunityContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCommunityFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getMinerUrl\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"}],\"name\":\"getMinersOffers\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pricePerSecond\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"machinesAvailable\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vmTypeId\",\"type\":\"uint256\"}],\"internalType\":\"structBroker.VMOffer[]\",\"name\":\"filteredOffers\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStablecoinAddress\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUsersBookings\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vmTypeId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pricePerSecond\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bookedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bookedTill\",\"type\":\"uint256\"}],\"internalType\":\"structBroker.Booking[]\",\"name\":\"filteredBookings\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newCommunityAddress\",\"type\":\"address\"}],\"name\":\"setCommunityContract\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"setCommunityFee\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"url\",\"type\":\"bytes32\"}],\"name\":\"setMunerUrl\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"t\",\"type\":\"address\"}],\"name\":\"setStablecoinAddress\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"userAllowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"userBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"userDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"userLockedBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"userTokenBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"numTokens\",\"type\":\"uint256\"}],\"name\":\"withdrawCoin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052600060015560006003553480156200001b57600080fd5b5060405162004d5c38038062004d5c8339818101604052810190620000419190620000f3565b80600860006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505062000125565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620000bb826200008e565b9050919050565b620000cd81620000ae565b8114620000d957600080fd5b50565b600081519050620000ed81620000c2565b92915050565b6000602082840312156200010c576200010b62000089565b5b60006200011c84828501620000dc565b91505092915050565b614c2780620001356000396000f3fe608060405234801561001057600080fd5b50600436106101e45760003560e01c806362c296171161010f578063ab49ec45116100a2578063ca3b3b4b11610071578063ca3b3b4b146105e1578063cbb6bfbd14610611578063cf696d1414610641578063f44f11cd1461065f576101e4565b8063ab49ec4514610545578063bb0090c914610575578063bf15276514610593578063c700ff0a146105b1576101e4565b80639ecd0500116100de5780639ecd0500146104ad5780639f60c50c146104dd578063a35310d4146104f9578063a5701e6314610529576101e4565b806362c29617146104255780637bef4f54146104415780637ee080c61461045f57806390ad688b1461048f576101e4565b8063301bcaae1161018757806354fd4d501161015657806354fd4d5014610389578063569e2e0f146103a75780635a610b57146103c5578063626bbb27146103f5576101e4565b8063301bcaae146102db5780633b3e328b1461030b578063481461131461033b5780634eb264e814610359576101e4565b806320255c7e116101c357806320255c7e1461025357806320e56d93146102715780632d0d6c331461028d5780632dc8a2b9146102ab576101e4565b8062cf9021146101e95780630d12bbdb1461020557806318c5e50214610235575b600080fd5b61020360048036038101906101fe919061369d565b61067b565b005b61021f600480360381019061021a919061369d565b61077d565b60405161022c91906136e5565b60405180910390f35b61023d610836565b60405161024a919061370f565b60405180910390f35b61025b61083e565b604051610268919061370f565b60405180910390f35b61028b600480360381019061028691906137a3565b6108e3565b005b610295610ff8565b6040516102a2919061370f565b60405180910390f35b6102c560048036038101906102c091906137e3565b61109b565b6040516102d291906138ee565b60405180910390f35b6102f560048036038101906102f09190613935565b6111b0565b6040516103029190613a9f565b60405180910390f35b61032560048036038101906103209190613ac1565b61147a565b604051610332919061370f565b60405180910390f35b610343611570565b604051610350919061370f565b60405180910390f35b610373600480360381019061036e919061369d565b6115b7565b60405161038091906136e5565b60405180910390f35b610391611748565b60405161039e919061370f565b60405180910390f35b6103af61174d565b6040516103bc9190613b23565b60405180910390f35b6103df60048036038101906103da9190613935565b611777565b6040516103ec9190613a9f565b60405180910390f35b61040f600480360381019061040a9190613935565b611a41565b60405161041c9190613a9f565b60405180910390f35b61043f600480360381019061043a91906137e3565b611d0b565b005b6104496122f9565b6040516104569190613b9d565b60405180910390f35b6104796004803603810190610474919061369d565b612323565b6040516104869190613ccf565b60405180910390f35b610497612561565b6040516104a4919061370f565b60405180910390f35b6104c760048036038101906104c29190613d2f565b61256b565b6040516104d491906136e5565b60405180910390f35b6104f760048036038101906104f29190613d5c565b612647565b005b610513600480360381019061050e9190613935565b612a4a565b60405161052091906136e5565b60405180910390f35b610543600480360381019061053e9190613ac1565b612b26565b005b61055f600480360381019061055a9190613935565b612c02565b60405161056c9190613db5565b60405180910390f35b61057d612c4b565b60405161058a919061370f565b60405180910390f35b61059b612c92565b6040516105a8919061370f565b60405180910390f35b6105cb60048036038101906105c69190613935565b612d23565b6040516105d89190613ccf565b60405180910390f35b6105fb60048036038101906105f6919061369d565b612f8b565b60405161060891906136e5565b60405180910390f35b61062b60048036038101906106269190613dd0565b6130d4565b604051610638919061370f565b60405180910390f35b6106496134d5565b6040516106569190613e1f565b60405180910390f35b61067960048036038101906106749190613e66565b61356d565b005b3373ffffffffffffffffffffffffffffffffffffffff1660008083815260200190815260200160002060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461071e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161071590613f16565b60405180910390fd5b6000808281526020019081526020016000206000808201600090556001820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff0219169055600282016000905560038201600090556004820160009055505050565b6000808211801561078e5750606582105b80156107e75750600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b610826576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161081d90613fa8565b60405180910390fd5b8160098190555060009050919050565b600042905090565b6000600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663dd62ed3e33306040518363ffffffff1660e01b815260040161089d929190613fc8565b602060405180830381865afa1580156108ba573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108de9190614006565b905090565b6000600260008467ffffffffffffffff1681526020019081526020016000206040518060e001604052908160008201548152602001600182015481526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016003820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016004820154815260200160058201548152602001600682015481525050905060006109f7610836565b9050816060015173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610a6b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a629061407f565b60405180910390fd5b60018360ff161480610a80575060028360ff16145b610abf576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ab690614111565b60405180910390fd5b8160c001518110610b05576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610afc906141a3565b60405180910390fd5b60008060018560ff1603610b20576032915060329050610b37565b6009546064610b2f91906141f2565b915060095490505b60008460a0015184610b4991906141f2565b8560800151610b589190614226565b9050600060648483610b6a9190614226565b610b7491906142af565b905060008183610b8491906141f2565b905082600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610bd591906141f2565b925050819055508660a001518760c00151610bf091906141f2565b8760800151610bff9190614226565b600660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610c4d91906141f2565b92505081905550600260008860000151815260200190815260200160002060008082016000905560018201600090556002820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556003820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff0219169055600482016000905560058201600090556006820160009055505060018860ff1603610d7e57866060015173ffffffffffffffffffffffffffffffffffffffff16876040015173ffffffffffffffffffffffffffffffffffffffff167f26c93dbbeb58948b01dbae3d96c9cbe902a2f92082c8d91ed5f12e8579b4cf41848a600001518b60a001518b610d5c91906141f2565b8c60200151604051610d7194939291906142e0565b60405180910390a3610e09565b866060015173ffffffffffffffffffffffffffffffffffffffff16876040015173ffffffffffffffffffffffffffffffffffffffff167f232ba2cbdb135f5ee58e4023da7094472434ff90d2d63f3c10232a45ee4c00d1848a600001518b60a001518b610deb91906141f2565b8c60200151604051610e0094939291906142e0565b60405180910390a35b600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb8860400151846040518363ffffffff1660e01b8152600401610e6a929190614325565b6020604051808303816000875af1158015610e89573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ead919061437a565b610eec576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ee3906143f3565b60405180910390fd5b600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16836040518363ffffffff1660e01b8152600401610f6b929190614325565b6020604051808303816000875af1158015610f8a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610fae919061437a565b610fed576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610fe49061445f565b60405180910390fd5b505050505050505050565b6000600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231336040518263ffffffff1660e01b81526004016110559190613b23565b602060405180830381865afa158015611072573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110969190614006565b905090565b6110a36135b4565b600260008367ffffffffffffffff1681526020019081526020016000206040518060e001604052908160008201548152602001600182015481526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016003820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160048201548152602001600582015481526020016006820154815250509050919050565b6060600060035467ffffffffffffffff8111156111d0576111cf61447f565b5b60405190808252806020026020018201604052801561120957816020015b6111f66135b4565b8152602001906001900390816111ee5790505b509050600080600090505b6003548110156113c4578473ffffffffffffffffffffffffffffffffffffffff166002600083815260200190815260200160002060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16036113b157600260008281526020019081526020016000206040518060e001604052908160008201548152602001600182015481526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016003820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016004820154815260200160058201548152602001600682015481525050838381518110611396576113956144ae565b5b60200260200101819052506001826113ae91906144dd565b91505b80806113bc90614533565b915050611214565b508067ffffffffffffffff8111156113df576113de61447f565b5b60405190808252806020026020018201604052801561141857816020015b6114056135b4565b8152602001906001900390816113fd5790505b50925060005b8181101561147257828181518110611439576114386144ae565b5b6020026020010151848281518110611454576114536144ae565b5b6020026020010181905250808061146a90614533565b91505061141e565b505050919050565b60006040518060a0016040528060015481526020013373ffffffffffffffffffffffffffffffffffffffff1681526020018581526020018381526020018481525060008060015481526020019081526020016000206000820151816000015560208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506040820151816002015560608201518160030155608082015181600401559050506001600081548092919061155490614533565b91905055506001805461156791906141f2565b90509392505050565b6000600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905090565b6000816115c2612c92565b1015611603576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016115fa906145c7565b60405180910390fd5b600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb33846040518363ffffffff1660e01b8152600401611660929190614325565b6020604051808303816000875af115801561167f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116a3919061437a565b6116b05760009050611743565b81600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546116fb91906141f2565b600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550600190505b919050565b600181565b6000600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6060600060035467ffffffffffffffff8111156117975761179661447f565b5b6040519080825280602002602001820160405280156117d057816020015b6117bd6135b4565b8152602001906001900390816117b55790505b509050600080600090505b60035481101561198b578473ffffffffffffffffffffffffffffffffffffffff166002600083815260200190815260200160002060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff160361197857600260008281526020019081526020016000206040518060e001604052908160008201548152602001600182015481526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016003820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600482015481526020016005820154815260200160068201548152505083838151811061195d5761195c6144ae565b5b602002602001018190525060018261197591906144dd565b91505b808061198390614533565b9150506117db565b508067ffffffffffffffff8111156119a6576119a561447f565b5b6040519080825280602002602001820160405280156119df57816020015b6119cc6135b4565b8152602001906001900390816119c45790505b50925060005b81811015611a3957828181518110611a00576119ff6144ae565b5b6020026020010151848281518110611a1b57611a1a6144ae565b5b60200260200101819052508080611a3190614533565b9150506119e5565b505050919050565b6060600060035467ffffffffffffffff811115611a6157611a6061447f565b5b604051908082528060200260200182016040528015611a9a57816020015b611a876135b4565b815260200190600190039081611a7f5790505b509050600080600090505b600354811015611c55578473ffffffffffffffffffffffffffffffffffffffff166002600083815260200190815260200160002060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603611c4257600260008281526020019081526020016000206040518060e001604052908160008201548152602001600182015481526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016003820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016004820154815260200160058201548152602001600682015481525050838381518110611c2757611c266144ae565b5b6020026020010181905250600182611c3f91906144dd565b91505b8080611c4d90614533565b915050611aa5565b508067ffffffffffffffff811115611c7057611c6f61447f565b5b604051908082528060200260200182016040528015611ca957816020015b611c966135b4565b815260200190600190039081611c8e5790505b50925060005b81811015611d0357828181518110611cca57611cc96144ae565b5b6020026020010151848281518110611ce557611ce46144ae565b5b60200260200101819052508080611cfb90614533565b915050611caf565b505050919050565b6000600260008367ffffffffffffffff1681526020019081526020016000206040518060e001604052908160008201548152602001600182015481526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016003820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600482015481526020016005820154815260200160068201548152505090506000611e1f610836565b9050816040015173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611e93576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611e8a90614633565b60405180910390fd5b8160c00151811015611eda576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611ed19061469f565b60405180910390fd5b60008260a001518360c00151611ef091906141f2565b8360800151611eff9190614226565b9050600060646009546064611f1491906141f2565b83611f1f9190614226565b611f2991906142af565b905060008183611f3991906141f2565b90508260056000876060015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254611f8e91906141f2565b925050819055508260066000876060015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254611fe891906141f2565b92505081905550600260008660000151815260200190815260200160002060008082016000905560018201600090556002820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556003820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556004820160009055600582016000905560068201600090555050846060015173ffffffffffffffffffffffffffffffffffffffff16856040015173ffffffffffffffffffffffffffffffffffffffff167f75d2566632c43a31ab3bc2a82efa0c6d7797ad569fb798ae857716022614d90c8488600001518960a001518a60c001516120f091906141f2565b8a6020015160405161210594939291906142e0565b60405180910390a3600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb8660400151846040518363ffffffff1660e01b815260040161216e929190614325565b6020604051808303816000875af115801561218d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906121b1919061437a565b6121f0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016121e7906143f3565b60405180910390fd5b600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16836040518363ffffffff1660e01b815260040161226f929190614325565b6020604051808303816000875af115801561228e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906122b2919061437a565b6122f1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016122e89061445f565b60405180910390fd5b505050505050565b6000600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6060600060015467ffffffffffffffff8111156123435761234261447f565b5b60405190808252806020026020018201604052801561237c57816020015b61236961361d565b8152602001906001900390816123615790505b509050600080600090505b6001548110156124ab5784600080838152602001908152602001600020600401541480156123ca5750600080600083815260200190815260200160002060030154115b15612498576000808281526020019081526020016000206040518060a0016040529081600082015481526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600282015481526020016003820154815260200160048201548152505083838151811061247d5761247c6144ae565b5b602002602001018190525060018261249591906144dd565b91505b80806124a390614533565b915050612387565b508067ffffffffffffffff8111156124c6576124c561447f565b5b6040519080825280602002602001820160405280156124ff57816020015b6124ec61361d565b8152602001906001900390816124e45790505b50925060005b81811015612559578281815181106125205761251f6144ae565b5b602002602001015184828151811061253b5761253a6144ae565b5b6020026020010181905250808061255190614533565b915050612505565b505050919050565b6000600954905090565b6000600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146125fd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016125f490614731565b60405180910390fd5b81600760006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060009050919050565b6000600260008467ffffffffffffffff1681526020019081526020016000206040518060e001604052908160008201548152602001600182015481526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016003820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160048201548152602001600582015481526020016006820154815250509050600061275b610836565b9050816060015173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146127cf576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016127c69061479d565b60405180910390fd5b8160c001518110612815576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161280c90614809565b60405180910390fd5b8282608001516128259190614226565b61282d612c92565b101561286e576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016128659061489b565b60405180910390fd5b82826080015161287e9190614226565b600660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546128cc91906144dd565b92505081905550828260c0018181516128e591906144dd565b9150818152505081600260008667ffffffffffffffff168152602001908152602001600020600082015181600001556020820151816001015560408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506080820151816004015560a0820151816005015560c08201518160060155905050816060015173ffffffffffffffffffffffffffffffffffffffff16826040015173ffffffffffffffffffffffffffffffffffffffff167f075b18a02d51e5a00f1bf8c8adbeba55d6fc957987d4285794bcd9029814746884600001518560200151604051612a3c9291906148bb565b60405180910390a350505050565b6000600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614612adc576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612ad390614956565b60405180910390fd5b81600860006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060009050919050565b3373ffffffffffffffffffffffffffffffffffffffff1660008085815260200190815260200160002060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614612bc9576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612bc0906149e8565b60405180910390fd5b81600080858152602001908152602001600020600301819055508060008085815260200190815260200160002060020181905550505050565b6000600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b6000600660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905090565b6000600660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054612d1e91906141f2565b905090565b6060600060015467ffffffffffffffff811115612d4357612d4261447f565b5b604051908082528060200260200182016040528015612d7c57816020015b612d6961361d565b815260200190600190039081612d615790505b509050600080600090505b600154811015612ed5578473ffffffffffffffffffffffffffffffffffffffff1660008083815260200190815260200160002060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603612ec2576000808281526020019081526020016000206040518060a0016040529081600082015481526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820154815260200160038201548152602001600482015481525050838381518110612ea757612ea66144ae565b5b6020026020010181905250600182612ebf91906144dd565b91505b8080612ecd90614533565b915050612d87565b508067ffffffffffffffff811115612ef057612eef61447f565b5b604051908082528060200260200182016040528015612f2957816020015b612f1661361d565b815260200190600190039081612f0e5790505b50925060005b81811015612f8357828181518110612f4a57612f496144ae565b5b6020026020010151848281518110612f6557612f646144ae565b5b60200260200101819052508080612f7b90614533565b915050612f2f565b505050919050565b6000600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd3330856040518463ffffffff1660e01b8152600401612fec93929190614a08565b6020604051808303816000875af115801561300b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061302f919061437a565b61303c57600090506130cf565b81600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461308791906144dd565b600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550600190505b919050565b600080600080858152602001908152602001600020600301541180156130fa5750600082115b801561312d575081600080858152602001908152602001600020600201546131229190614226565b61312a612c92565b10155b61316c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161316390614afd565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff1660008085815260200190815260200160002060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603613210576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161320790614b8f565b60405180910390fd5b81600080858152602001908152602001600020600201546132319190614226565b600660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461327f91906144dd565b9250508190555060006040518060e00160405280600354815260200160008087815260200190815260200160002060040154815260200160008087815260200190815260200160002060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020013373ffffffffffffffffffffffffffffffffffffffff168152602001600080878152602001908152602001600020600201548152602001428152602001844261335091906144dd565b815250905080600260006003548152602001908152602001600020600082015181600001556020820151816001015560408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506080820151816004015560a0820151816005015560c082015181600601559050506003600081548092919061344190614533565b9190505550806060015173ffffffffffffffffffffffffffffffffffffffff16816040015173ffffffffffffffffffffffffffffffffffffffff167f246cbcaa18d1e49adb9a7db733d35a6e5a2690726bc6b2326bd3c53fa602f1b6836000015184602001516040516134b59291906148bb565b60405180910390a360016003546134cc91906141f2565b91505092915050565b6000600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663313ce5676040518163ffffffff1660e01b8152600401602060405180830381865afa158015613544573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906135689190614bc4565b905090565b80600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555050565b6040518060e001604052806000815260200160008152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff1681526020016000815260200160008152602001600081525090565b6040518060a0016040528060008152602001600073ffffffffffffffffffffffffffffffffffffffff1681526020016000815260200160008152602001600081525090565b600080fd5b6000819050919050565b61367a81613667565b811461368557600080fd5b50565b60008135905061369781613671565b92915050565b6000602082840312156136b3576136b2613662565b5b60006136c184828501613688565b91505092915050565b60008115159050919050565b6136df816136ca565b82525050565b60006020820190506136fa60008301846136d6565b92915050565b61370981613667565b82525050565b60006020820190506137246000830184613700565b92915050565b600067ffffffffffffffff82169050919050565b6137478161372a565b811461375257600080fd5b50565b6000813590506137648161373e565b92915050565b600060ff82169050919050565b6137808161376a565b811461378b57600080fd5b50565b60008135905061379d81613777565b92915050565b600080604083850312156137ba576137b9613662565b5b60006137c885828601613755565b92505060206137d98582860161378e565b9150509250929050565b6000602082840312156137f9576137f8613662565b5b600061380784828501613755565b91505092915050565b61381981613667565b82525050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061384a8261381f565b9050919050565b61385a8161383f565b82525050565b60e0820160008201516138766000850182613810565b5060208201516138896020850182613810565b50604082015161389c6040850182613851565b5060608201516138af6060850182613851565b5060808201516138c26080850182613810565b5060a08201516138d560a0850182613810565b5060c08201516138e860c0850182613810565b50505050565b600060e0820190506139036000830184613860565b92915050565b6139128161383f565b811461391d57600080fd5b50565b60008135905061392f81613909565b92915050565b60006020828403121561394b5761394a613662565b5b600061395984828501613920565b91505092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b60e0820160008201516139a46000850182613810565b5060208201516139b76020850182613810565b5060408201516139ca6040850182613851565b5060608201516139dd6060850182613851565b5060808201516139f06080850182613810565b5060a0820151613a0360a0850182613810565b5060c0820151613a1660c0850182613810565b50505050565b6000613a28838361398e565b60e08301905092915050565b6000602082019050919050565b6000613a4c82613962565b613a56818561396d565b9350613a618361397e565b8060005b83811015613a92578151613a798882613a1c565b9750613a8483613a34565b925050600181019050613a65565b5085935050505092915050565b60006020820190508181036000830152613ab98184613a41565b905092915050565b600080600060608486031215613ada57613ad9613662565b5b6000613ae886828701613688565b9350506020613af986828701613688565b9250506040613b0a86828701613688565b9150509250925092565b613b1d8161383f565b82525050565b6000602082019050613b386000830184613b14565b92915050565b6000819050919050565b6000613b63613b5e613b598461381f565b613b3e565b61381f565b9050919050565b6000613b7582613b48565b9050919050565b6000613b8782613b6a565b9050919050565b613b9781613b7c565b82525050565b6000602082019050613bb26000830184613b8e565b92915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b60a082016000820151613bfa6000850182613810565b506020820151613c0d6020850182613851565b506040820151613c206040850182613810565b506060820151613c336060850182613810565b506080820151613c466080850182613810565b50505050565b6000613c588383613be4565b60a08301905092915050565b6000602082019050919050565b6000613c7c82613bb8565b613c868185613bc3565b9350613c9183613bd4565b8060005b83811015613cc2578151613ca98882613c4c565b9750613cb483613c64565b925050600181019050613c95565b5085935050505092915050565b60006020820190508181036000830152613ce98184613c71565b905092915050565b6000613cfc8261383f565b9050919050565b613d0c81613cf1565b8114613d1757600080fd5b50565b600081359050613d2981613d03565b92915050565b600060208284031215613d4557613d44613662565b5b6000613d5384828501613d1a565b91505092915050565b60008060408385031215613d7357613d72613662565b5b6000613d8185828601613755565b9250506020613d9285828601613688565b9150509250929050565b6000819050919050565b613daf81613d9c565b82525050565b6000602082019050613dca6000830184613da6565b92915050565b60008060408385031215613de757613de6613662565b5b6000613df585828601613688565b9250506020613e0685828601613688565b9150509250929050565b613e198161376a565b82525050565b6000602082019050613e346000830184613e10565b92915050565b613e4381613d9c565b8114613e4e57600080fd5b50565b600081359050613e6081613e3a565b92915050565b600060208284031215613e7c57613e7b613662565b5b6000613e8a84828501613e51565b91505092915050565b600082825260208201905092915050565b7f4f6e6c7920746865206f776e65722063616e2072656d6f766520616e206f666660008201527f6572000000000000000000000000000000000000000000000000000000000000602082015250565b6000613f00602283613e93565b9150613f0b82613ea4565b604082019050919050565b60006020820190508181036000830152613f2f81613ef3565b9050919050565b7f636f6d6d756e697479206665652073686f756c6420626520696e2072616e676560008201527f206f66203120746f203130300000000000000000000000000000000000000000602082015250565b6000613f92602c83613e93565b9150613f9d82613f36565b604082019050919050565b60006020820190508181036000830152613fc181613f85565b9050919050565b6000604082019050613fdd6000830185613b14565b613fea6020830184613b14565b9392505050565b60008151905061400081613671565b92915050565b60006020828403121561401c5761401b613662565b5b600061402a84828501613ff1565b91505092915050565b7f6f6e6c79206f776e6572206f6620626f6f6b696e672063616e207265706f7274600082015250565b6000614069602083613e93565b915061407482614033565b602082019050919050565b600060208201905081810360008301526140988161405c565b9050919050565b7f7265706f7274206f722073746f7020626f6f6b696e672069732072657175697260008201527f6564000000000000000000000000000000000000000000000000000000000000602082015250565b60006140fb602283613e93565b91506141068261409f565b604082019050919050565b6000602082019050818103600083015261412a816140ee565b9050919050565b7f626f6f6b696e6720697320657870697265642c206d696e65722073686f756c6460008201527f20636c61696d206578706972656420626f6f6b696e6700000000000000000000602082015250565b600061418d603683613e93565b915061419882614131565b604082019050919050565b600060208201905081810360008301526141bc81614180565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006141fd82613667565b915061420883613667565b92508282101561421b5761421a6141c3565b5b828203905092915050565b600061423182613667565b915061423c83613667565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615614275576142746141c3565b5b828202905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b60006142ba82613667565b91506142c583613667565b9250826142d5576142d4614280565b5b828204905092915050565b60006080820190506142f56000830187613700565b6143026020830186613700565b61430f6040830185613700565b61431c6060830184613700565b95945050505050565b600060408201905061433a6000830185613b14565b6143476020830184613700565b9392505050565b614357816136ca565b811461436257600080fd5b50565b6000815190506143748161434e565b92915050565b6000602082840312156143905761438f613662565b5b600061439e84828501614365565b91505092915050565b7f636f756c64206e6f74207061796f757420666f72206d696e6572000000000000600082015250565b60006143dd601a83613e93565b91506143e8826143a7565b602082019050919050565b6000602082019050818103600083015261440c816143d0565b9050919050565b7f636f756c64206e6f74207061796f757420666f7220636f6d6d756e6974790000600082015250565b6000614449601e83613e93565b915061445482614413565b602082019050919050565b600060208201905081810360008301526144788161443c565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60006144e882613667565b91506144f383613667565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115614528576145276141c3565b5b828201905092915050565b600061453e82613667565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036145705761456f6141c3565b5b600182019050919050565b7f696e73756666696369656e742066756e647320746f2077697468647261770000600082015250565b60006145b1601e83613e93565b91506145bc8261457b565b602082019050919050565b600060208201905081810360008301526145e0816145a4565b9050919050565b7f6f6e6c79206d696e6572206f6620626f6f6b696e672063616e20636c61696d00600082015250565b600061461d601f83613e93565b9150614628826145e7565b602082019050919050565b6000602082019050818103600083015261464c81614610565b9050919050565b7f626f6f6b696e67206973207374696c6c20696e20757365000000000000000000600082015250565b6000614689601783613e93565b915061469482614653565b602082019050919050565b600060208201905081810360008301526146b88161467c565b9050919050565b7f6f6e6c7920636f6d6d756e69747920636f6e74726163742063616e207365742060008201527f737461626c65636f696e00000000000000000000000000000000000000000000602082015250565b600061471b602a83613e93565b9150614726826146bf565b604082019050919050565b6000602082019050818103600083015261474a8161470e565b9050919050565b7f6f6e6c79206f776e6572206f6620626f6f6b696e672063616e20657874656e64600082015250565b6000614787602083613e93565b915061479282614751565b602082019050919050565b600060208201905081810360008301526147b68161477a565b9050919050565b7f626f6f6b696e6720697320657870697265640000000000000000000000000000600082015250565b60006147f3601283613e93565b91506147fe826147bd565b602082019050919050565b60006020820190508181036000830152614822816147e6565b9050919050565b7f696e73756666696369656e742066756e647320746f20657874656e6420626f6f60008201527f6b696e6700000000000000000000000000000000000000000000000000000000602082015250565b6000614885602483613e93565b915061489082614829565b604082019050919050565b600060208201905081810360008301526148b481614878565b9050919050565b60006040820190506148d06000830185613700565b6148dd6020830184613700565b9392505050565b7f6f6e6c7920636f6d6d756e69747920636f6e74726163742063616e207365742060008201527f6e657720636f6d6d756e69747920636f6e747261637400000000000000000000602082015250565b6000614940603683613e93565b915061494b826148e4565b604082019050919050565b6000602082019050818103600083015261496f81614933565b9050919050565b7f4f6e6c7920746865206f776e65722063616e2075706461746520616e206f666660008201527f6572000000000000000000000000000000000000000000000000000000000000602082015250565b60006149d2602283613e93565b91506149dd82614976565b604082019050919050565b60006020820190508181036000830152614a01816149c5565b9050919050565b6000606082019050614a1d6000830186613b14565b614a2a6020830185613b14565b614a376040830184613700565b949350505050565b7f4e6f206d616368696e657320617661696c61626c65204f52207365637320736860008201527f6f756c6420626520706f736974697665204f52207573657220646f6573206e6f60208201527f74206861766520726571756972656420616d6f756e7420746f20626f6f6b206d60408201527f616368696e650000000000000000000000000000000000000000000000000000606082015250565b6000614ae7606683613e93565b9150614af282614a3f565b608082019050919050565b60006020820190508181036000830152614b1681614ada565b9050919050565b7f6f666665722077697468207468617420696e64657820646f6573206e6f74206560008201527f7869737400000000000000000000000000000000000000000000000000000000602082015250565b6000614b79602483613e93565b9150614b8482614b1d565b604082019050919050565b60006020820190508181036000830152614ba881614b6c565b9050919050565b600081519050614bbe81613777565b92915050565b600060208284031215614bda57614bd9613662565b5b6000614be884828501614baf565b9150509291505056fea2646970667358221220114693761340a1e1235ee430bf0b9f34aa8d3bf86673ba7ba657f22159c87c0664736f6c634300080f0033",
}

// BrokerABI is the input ABI used to generate the binding from.
// Deprecated: Use BrokerMetaData.ABI instead.
var BrokerABI = BrokerMetaData.ABI

// BrokerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BrokerMetaData.Bin instead.
var BrokerBin = BrokerMetaData.Bin

// DeployBroker deploys a new Ethereum contract, binding an instance of Broker to it.
func DeployBroker(auth *bind.TransactOpts, backend bind.ContractBackend, communityAddress common.Address) (common.Address, *types.Transaction, *Broker, error) {
	parsed, err := BrokerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BrokerBin), backend, communityAddress)
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

// FindBookingsByMiner is a free data retrieval call binding the contract method 0x301bcaae.
//
// Solidity: function findBookingsByMiner(address _miner) view returns((uint256,uint256,address,address,uint256,uint256,uint256)[] filteredBookings)
func (_Broker *BrokerCaller) FindBookingsByMiner(opts *bind.CallOpts, _miner common.Address) ([]BrokerBooking, error) {
	var out []interface{}
	err := _Broker.contract.Call(opts, &out, "findBookingsByMiner", _miner)

	if err != nil {
		return *new([]BrokerBooking), err
	}

	out0 := *abi.ConvertType(out[0], new([]BrokerBooking)).(*[]BrokerBooking)

	return out0, err

}

// FindBookingsByMiner is a free data retrieval call binding the contract method 0x301bcaae.
//
// Solidity: function findBookingsByMiner(address _miner) view returns((uint256,uint256,address,address,uint256,uint256,uint256)[] filteredBookings)
func (_Broker *BrokerSession) FindBookingsByMiner(_miner common.Address) ([]BrokerBooking, error) {
	return _Broker.Contract.FindBookingsByMiner(&_Broker.CallOpts, _miner)
}

// FindBookingsByMiner is a free data retrieval call binding the contract method 0x301bcaae.
//
// Solidity: function findBookingsByMiner(address _miner) view returns((uint256,uint256,address,address,uint256,uint256,uint256)[] filteredBookings)
func (_Broker *BrokerCallerSession) FindBookingsByMiner(_miner common.Address) ([]BrokerBooking, error) {
	return _Broker.Contract.FindBookingsByMiner(&_Broker.CallOpts, _miner)
}

// FindBookingsByUser is a free data retrieval call binding the contract method 0x5a610b57.
//
// Solidity: function findBookingsByUser(address _owner) view returns((uint256,uint256,address,address,uint256,uint256,uint256)[] filteredBookings)
func (_Broker *BrokerCaller) FindBookingsByUser(opts *bind.CallOpts, _owner common.Address) ([]BrokerBooking, error) {
	var out []interface{}
	err := _Broker.contract.Call(opts, &out, "findBookingsByUser", _owner)

	if err != nil {
		return *new([]BrokerBooking), err
	}

	out0 := *abi.ConvertType(out[0], new([]BrokerBooking)).(*[]BrokerBooking)

	return out0, err

}

// FindBookingsByUser is a free data retrieval call binding the contract method 0x5a610b57.
//
// Solidity: function findBookingsByUser(address _owner) view returns((uint256,uint256,address,address,uint256,uint256,uint256)[] filteredBookings)
func (_Broker *BrokerSession) FindBookingsByUser(_owner common.Address) ([]BrokerBooking, error) {
	return _Broker.Contract.FindBookingsByUser(&_Broker.CallOpts, _owner)
}

// FindBookingsByUser is a free data retrieval call binding the contract method 0x5a610b57.
//
// Solidity: function findBookingsByUser(address _owner) view returns((uint256,uint256,address,address,uint256,uint256,uint256)[] filteredBookings)
func (_Broker *BrokerCallerSession) FindBookingsByUser(_owner common.Address) ([]BrokerBooking, error) {
	return _Broker.Contract.FindBookingsByUser(&_Broker.CallOpts, _owner)
}

// GetAvailableOffers is a free data retrieval call binding the contract method 0x7ee080c6.
//
// Solidity: function getAvailableOffers(uint256 _vmTypeId) view returns((uint256,address,uint256,uint256,uint256)[] filteredOffers)
func (_Broker *BrokerCaller) GetAvailableOffers(opts *bind.CallOpts, _vmTypeId *big.Int) ([]BrokerVMOffer, error) {
	var out []interface{}
	err := _Broker.contract.Call(opts, &out, "getAvailableOffers", _vmTypeId)

	if err != nil {
		return *new([]BrokerVMOffer), err
	}

	out0 := *abi.ConvertType(out[0], new([]BrokerVMOffer)).(*[]BrokerVMOffer)

	return out0, err

}

// GetAvailableOffers is a free data retrieval call binding the contract method 0x7ee080c6.
//
// Solidity: function getAvailableOffers(uint256 _vmTypeId) view returns((uint256,address,uint256,uint256,uint256)[] filteredOffers)
func (_Broker *BrokerSession) GetAvailableOffers(_vmTypeId *big.Int) ([]BrokerVMOffer, error) {
	return _Broker.Contract.GetAvailableOffers(&_Broker.CallOpts, _vmTypeId)
}

// GetAvailableOffers is a free data retrieval call binding the contract method 0x7ee080c6.
//
// Solidity: function getAvailableOffers(uint256 _vmTypeId) view returns((uint256,address,uint256,uint256,uint256)[] filteredOffers)
func (_Broker *BrokerCallerSession) GetAvailableOffers(_vmTypeId *big.Int) ([]BrokerVMOffer, error) {
	return _Broker.Contract.GetAvailableOffers(&_Broker.CallOpts, _vmTypeId)
}

// GetBooking is a free data retrieval call binding the contract method 0x2dc8a2b9.
//
// Solidity: function getBooking(uint64 index) view returns((uint256,uint256,address,address,uint256,uint256,uint256) booking)
func (_Broker *BrokerCaller) GetBooking(opts *bind.CallOpts, index uint64) (BrokerBooking, error) {
	var out []interface{}
	err := _Broker.contract.Call(opts, &out, "getBooking", index)

	if err != nil {
		return *new(BrokerBooking), err
	}

	out0 := *abi.ConvertType(out[0], new(BrokerBooking)).(*BrokerBooking)

	return out0, err

}

// GetBooking is a free data retrieval call binding the contract method 0x2dc8a2b9.
//
// Solidity: function getBooking(uint64 index) view returns((uint256,uint256,address,address,uint256,uint256,uint256) booking)
func (_Broker *BrokerSession) GetBooking(index uint64) (BrokerBooking, error) {
	return _Broker.Contract.GetBooking(&_Broker.CallOpts, index)
}

// GetBooking is a free data retrieval call binding the contract method 0x2dc8a2b9.
//
// Solidity: function getBooking(uint64 index) view returns((uint256,uint256,address,address,uint256,uint256,uint256) booking)
func (_Broker *BrokerCallerSession) GetBooking(index uint64) (BrokerBooking, error) {
	return _Broker.Contract.GetBooking(&_Broker.CallOpts, index)
}

// GetCoinDecimals is a free data retrieval call binding the contract method 0xcf696d14.
//
// Solidity: function getCoinDecimals() view returns(uint8)
func (_Broker *BrokerCaller) GetCoinDecimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Broker.contract.Call(opts, &out, "getCoinDecimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetCoinDecimals is a free data retrieval call binding the contract method 0xcf696d14.
//
// Solidity: function getCoinDecimals() view returns(uint8)
func (_Broker *BrokerSession) GetCoinDecimals() (uint8, error) {
	return _Broker.Contract.GetCoinDecimals(&_Broker.CallOpts)
}

// GetCoinDecimals is a free data retrieval call binding the contract method 0xcf696d14.
//
// Solidity: function getCoinDecimals() view returns(uint8)
func (_Broker *BrokerCallerSession) GetCoinDecimals() (uint8, error) {
	return _Broker.Contract.GetCoinDecimals(&_Broker.CallOpts)
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
// Solidity: function getCommunityFee() view returns(uint256)
func (_Broker *BrokerCaller) GetCommunityFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Broker.contract.Call(opts, &out, "getCommunityFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCommunityFee is a free data retrieval call binding the contract method 0x90ad688b.
//
// Solidity: function getCommunityFee() view returns(uint256)
func (_Broker *BrokerSession) GetCommunityFee() (*big.Int, error) {
	return _Broker.Contract.GetCommunityFee(&_Broker.CallOpts)
}

// GetCommunityFee is a free data retrieval call binding the contract method 0x90ad688b.
//
// Solidity: function getCommunityFee() view returns(uint256)
func (_Broker *BrokerCallerSession) GetCommunityFee() (*big.Int, error) {
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

// GetMinersOffers is a free data retrieval call binding the contract method 0xc700ff0a.
//
// Solidity: function getMinersOffers(address miner) view returns((uint256,address,uint256,uint256,uint256)[] filteredOffers)
func (_Broker *BrokerCaller) GetMinersOffers(opts *bind.CallOpts, miner common.Address) ([]BrokerVMOffer, error) {
	var out []interface{}
	err := _Broker.contract.Call(opts, &out, "getMinersOffers", miner)

	if err != nil {
		return *new([]BrokerVMOffer), err
	}

	out0 := *abi.ConvertType(out[0], new([]BrokerVMOffer)).(*[]BrokerVMOffer)

	return out0, err

}

// GetMinersOffers is a free data retrieval call binding the contract method 0xc700ff0a.
//
// Solidity: function getMinersOffers(address miner) view returns((uint256,address,uint256,uint256,uint256)[] filteredOffers)
func (_Broker *BrokerSession) GetMinersOffers(miner common.Address) ([]BrokerVMOffer, error) {
	return _Broker.Contract.GetMinersOffers(&_Broker.CallOpts, miner)
}

// GetMinersOffers is a free data retrieval call binding the contract method 0xc700ff0a.
//
// Solidity: function getMinersOffers(address miner) view returns((uint256,address,uint256,uint256,uint256)[] filteredOffers)
func (_Broker *BrokerCallerSession) GetMinersOffers(miner common.Address) ([]BrokerVMOffer, error) {
	return _Broker.Contract.GetMinersOffers(&_Broker.CallOpts, miner)
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

// GetUsersBookings is a free data retrieval call binding the contract method 0x626bbb27.
//
// Solidity: function getUsersBookings(address user) view returns((uint256,uint256,address,address,uint256,uint256,uint256)[] filteredBookings)
func (_Broker *BrokerCaller) GetUsersBookings(opts *bind.CallOpts, user common.Address) ([]BrokerBooking, error) {
	var out []interface{}
	err := _Broker.contract.Call(opts, &out, "getUsersBookings", user)

	if err != nil {
		return *new([]BrokerBooking), err
	}

	out0 := *abi.ConvertType(out[0], new([]BrokerBooking)).(*[]BrokerBooking)

	return out0, err

}

// GetUsersBookings is a free data retrieval call binding the contract method 0x626bbb27.
//
// Solidity: function getUsersBookings(address user) view returns((uint256,uint256,address,address,uint256,uint256,uint256)[] filteredBookings)
func (_Broker *BrokerSession) GetUsersBookings(user common.Address) ([]BrokerBooking, error) {
	return _Broker.Contract.GetUsersBookings(&_Broker.CallOpts, user)
}

// GetUsersBookings is a free data retrieval call binding the contract method 0x626bbb27.
//
// Solidity: function getUsersBookings(address user) view returns((uint256,uint256,address,address,uint256,uint256,uint256)[] filteredBookings)
func (_Broker *BrokerCallerSession) GetUsersBookings(user common.Address) ([]BrokerBooking, error) {
	return _Broker.Contract.GetUsersBookings(&_Broker.CallOpts, user)
}

// UserAllowance is a free data retrieval call binding the contract method 0x20255c7e.
//
// Solidity: function userAllowance() view returns(uint256)
func (_Broker *BrokerCaller) UserAllowance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Broker.contract.Call(opts, &out, "userAllowance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserAllowance is a free data retrieval call binding the contract method 0x20255c7e.
//
// Solidity: function userAllowance() view returns(uint256)
func (_Broker *BrokerSession) UserAllowance() (*big.Int, error) {
	return _Broker.Contract.UserAllowance(&_Broker.CallOpts)
}

// UserAllowance is a free data retrieval call binding the contract method 0x20255c7e.
//
// Solidity: function userAllowance() view returns(uint256)
func (_Broker *BrokerCallerSession) UserAllowance() (*big.Int, error) {
	return _Broker.Contract.UserAllowance(&_Broker.CallOpts)
}

// UserBalance is a free data retrieval call binding the contract method 0xbf152765.
//
// Solidity: function userBalance() view returns(uint256)
func (_Broker *BrokerCaller) UserBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Broker.contract.Call(opts, &out, "userBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserBalance is a free data retrieval call binding the contract method 0xbf152765.
//
// Solidity: function userBalance() view returns(uint256)
func (_Broker *BrokerSession) UserBalance() (*big.Int, error) {
	return _Broker.Contract.UserBalance(&_Broker.CallOpts)
}

// UserBalance is a free data retrieval call binding the contract method 0xbf152765.
//
// Solidity: function userBalance() view returns(uint256)
func (_Broker *BrokerCallerSession) UserBalance() (*big.Int, error) {
	return _Broker.Contract.UserBalance(&_Broker.CallOpts)
}

// UserDeposit is a free data retrieval call binding the contract method 0x48146113.
//
// Solidity: function userDeposit() view returns(uint256)
func (_Broker *BrokerCaller) UserDeposit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Broker.contract.Call(opts, &out, "userDeposit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserDeposit is a free data retrieval call binding the contract method 0x48146113.
//
// Solidity: function userDeposit() view returns(uint256)
func (_Broker *BrokerSession) UserDeposit() (*big.Int, error) {
	return _Broker.Contract.UserDeposit(&_Broker.CallOpts)
}

// UserDeposit is a free data retrieval call binding the contract method 0x48146113.
//
// Solidity: function userDeposit() view returns(uint256)
func (_Broker *BrokerCallerSession) UserDeposit() (*big.Int, error) {
	return _Broker.Contract.UserDeposit(&_Broker.CallOpts)
}

// UserLockedBalance is a free data retrieval call binding the contract method 0xbb0090c9.
//
// Solidity: function userLockedBalance() view returns(uint256)
func (_Broker *BrokerCaller) UserLockedBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Broker.contract.Call(opts, &out, "userLockedBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserLockedBalance is a free data retrieval call binding the contract method 0xbb0090c9.
//
// Solidity: function userLockedBalance() view returns(uint256)
func (_Broker *BrokerSession) UserLockedBalance() (*big.Int, error) {
	return _Broker.Contract.UserLockedBalance(&_Broker.CallOpts)
}

// UserLockedBalance is a free data retrieval call binding the contract method 0xbb0090c9.
//
// Solidity: function userLockedBalance() view returns(uint256)
func (_Broker *BrokerCallerSession) UserLockedBalance() (*big.Int, error) {
	return _Broker.Contract.UserLockedBalance(&_Broker.CallOpts)
}

// UserTokenBalance is a free data retrieval call binding the contract method 0x2d0d6c33.
//
// Solidity: function userTokenBalance() view returns(uint256)
func (_Broker *BrokerCaller) UserTokenBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Broker.contract.Call(opts, &out, "userTokenBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserTokenBalance is a free data retrieval call binding the contract method 0x2d0d6c33.
//
// Solidity: function userTokenBalance() view returns(uint256)
func (_Broker *BrokerSession) UserTokenBalance() (*big.Int, error) {
	return _Broker.Contract.UserTokenBalance(&_Broker.CallOpts)
}

// UserTokenBalance is a free data retrieval call binding the contract method 0x2d0d6c33.
//
// Solidity: function userTokenBalance() view returns(uint256)
func (_Broker *BrokerCallerSession) UserTokenBalance() (*big.Int, error) {
	return _Broker.Contract.UserTokenBalance(&_Broker.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_Broker *BrokerCaller) Version(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Broker.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_Broker *BrokerSession) Version() (*big.Int, error) {
	return _Broker.Contract.Version(&_Broker.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_Broker *BrokerCallerSession) Version() (*big.Int, error) {
	return _Broker.Contract.Version(&_Broker.CallOpts)
}

// RemoveOffer is a paid mutator transaction binding the contract method 0x00cf9021.
//
// Solidity: function RemoveOffer(uint256 offerIndex) returns()
func (_Broker *BrokerTransactor) RemoveOffer(opts *bind.TransactOpts, offerIndex *big.Int) (*types.Transaction, error) {
	return _Broker.contract.Transact(opts, "RemoveOffer", offerIndex)
}

// RemoveOffer is a paid mutator transaction binding the contract method 0x00cf9021.
//
// Solidity: function RemoveOffer(uint256 offerIndex) returns()
func (_Broker *BrokerSession) RemoveOffer(offerIndex *big.Int) (*types.Transaction, error) {
	return _Broker.Contract.RemoveOffer(&_Broker.TransactOpts, offerIndex)
}

// RemoveOffer is a paid mutator transaction binding the contract method 0x00cf9021.
//
// Solidity: function RemoveOffer(uint256 offerIndex) returns()
func (_Broker *BrokerTransactorSession) RemoveOffer(offerIndex *big.Int) (*types.Transaction, error) {
	return _Broker.Contract.RemoveOffer(&_Broker.TransactOpts, offerIndex)
}

// UpdateOffer is a paid mutator transaction binding the contract method 0xa5701e63.
//
// Solidity: function UpdateOffer(uint256 offerIndex, uint256 machinesAvailable, uint256 pps) returns()
func (_Broker *BrokerTransactor) UpdateOffer(opts *bind.TransactOpts, offerIndex *big.Int, machinesAvailable *big.Int, pps *big.Int) (*types.Transaction, error) {
	return _Broker.contract.Transact(opts, "UpdateOffer", offerIndex, machinesAvailable, pps)
}

// UpdateOffer is a paid mutator transaction binding the contract method 0xa5701e63.
//
// Solidity: function UpdateOffer(uint256 offerIndex, uint256 machinesAvailable, uint256 pps) returns()
func (_Broker *BrokerSession) UpdateOffer(offerIndex *big.Int, machinesAvailable *big.Int, pps *big.Int) (*types.Transaction, error) {
	return _Broker.Contract.UpdateOffer(&_Broker.TransactOpts, offerIndex, machinesAvailable, pps)
}

// UpdateOffer is a paid mutator transaction binding the contract method 0xa5701e63.
//
// Solidity: function UpdateOffer(uint256 offerIndex, uint256 machinesAvailable, uint256 pps) returns()
func (_Broker *BrokerTransactorSession) UpdateOffer(offerIndex *big.Int, machinesAvailable *big.Int, pps *big.Int) (*types.Transaction, error) {
	return _Broker.Contract.UpdateOffer(&_Broker.TransactOpts, offerIndex, machinesAvailable, pps)
}

// AbortBooking is a paid mutator transaction binding the contract method 0x20e56d93.
//
// Solidity: function abortBooking(uint64 index, uint8 abortType) returns()
func (_Broker *BrokerTransactor) AbortBooking(opts *bind.TransactOpts, index uint64, abortType uint8) (*types.Transaction, error) {
	return _Broker.contract.Transact(opts, "abortBooking", index, abortType)
}

// AbortBooking is a paid mutator transaction binding the contract method 0x20e56d93.
//
// Solidity: function abortBooking(uint64 index, uint8 abortType) returns()
func (_Broker *BrokerSession) AbortBooking(index uint64, abortType uint8) (*types.Transaction, error) {
	return _Broker.Contract.AbortBooking(&_Broker.TransactOpts, index, abortType)
}

// AbortBooking is a paid mutator transaction binding the contract method 0x20e56d93.
//
// Solidity: function abortBooking(uint64 index, uint8 abortType) returns()
func (_Broker *BrokerTransactorSession) AbortBooking(index uint64, abortType uint8) (*types.Transaction, error) {
	return _Broker.Contract.AbortBooking(&_Broker.TransactOpts, index, abortType)
}

// AddOffer is a paid mutator transaction binding the contract method 0x3b3e328b.
//
// Solidity: function addOffer(uint256 pricePerSecond, uint256 vmTypeId, uint256 machinesAvailable) returns(uint256)
func (_Broker *BrokerTransactor) AddOffer(opts *bind.TransactOpts, pricePerSecond *big.Int, vmTypeId *big.Int, machinesAvailable *big.Int) (*types.Transaction, error) {
	return _Broker.contract.Transact(opts, "addOffer", pricePerSecond, vmTypeId, machinesAvailable)
}

// AddOffer is a paid mutator transaction binding the contract method 0x3b3e328b.
//
// Solidity: function addOffer(uint256 pricePerSecond, uint256 vmTypeId, uint256 machinesAvailable) returns(uint256)
func (_Broker *BrokerSession) AddOffer(pricePerSecond *big.Int, vmTypeId *big.Int, machinesAvailable *big.Int) (*types.Transaction, error) {
	return _Broker.Contract.AddOffer(&_Broker.TransactOpts, pricePerSecond, vmTypeId, machinesAvailable)
}

// AddOffer is a paid mutator transaction binding the contract method 0x3b3e328b.
//
// Solidity: function addOffer(uint256 pricePerSecond, uint256 vmTypeId, uint256 machinesAvailable) returns(uint256)
func (_Broker *BrokerTransactorSession) AddOffer(pricePerSecond *big.Int, vmTypeId *big.Int, machinesAvailable *big.Int) (*types.Transaction, error) {
	return _Broker.Contract.AddOffer(&_Broker.TransactOpts, pricePerSecond, vmTypeId, machinesAvailable)
}

// BookVM is a paid mutator transaction binding the contract method 0xcbb6bfbd.
//
// Solidity: function bookVM(uint256 offerIndex, uint256 secs) returns(uint256)
func (_Broker *BrokerTransactor) BookVM(opts *bind.TransactOpts, offerIndex *big.Int, secs *big.Int) (*types.Transaction, error) {
	return _Broker.contract.Transact(opts, "bookVM", offerIndex, secs)
}

// BookVM is a paid mutator transaction binding the contract method 0xcbb6bfbd.
//
// Solidity: function bookVM(uint256 offerIndex, uint256 secs) returns(uint256)
func (_Broker *BrokerSession) BookVM(offerIndex *big.Int, secs *big.Int) (*types.Transaction, error) {
	return _Broker.Contract.BookVM(&_Broker.TransactOpts, offerIndex, secs)
}

// BookVM is a paid mutator transaction binding the contract method 0xcbb6bfbd.
//
// Solidity: function bookVM(uint256 offerIndex, uint256 secs) returns(uint256)
func (_Broker *BrokerTransactorSession) BookVM(offerIndex *big.Int, secs *big.Int) (*types.Transaction, error) {
	return _Broker.Contract.BookVM(&_Broker.TransactOpts, offerIndex, secs)
}

// ClaimExpired is a paid mutator transaction binding the contract method 0x62c29617.
//
// Solidity: function claimExpired(uint64 index) returns()
func (_Broker *BrokerTransactor) ClaimExpired(opts *bind.TransactOpts, index uint64) (*types.Transaction, error) {
	return _Broker.contract.Transact(opts, "claimExpired", index)
}

// ClaimExpired is a paid mutator transaction binding the contract method 0x62c29617.
//
// Solidity: function claimExpired(uint64 index) returns()
func (_Broker *BrokerSession) ClaimExpired(index uint64) (*types.Transaction, error) {
	return _Broker.Contract.ClaimExpired(&_Broker.TransactOpts, index)
}

// ClaimExpired is a paid mutator transaction binding the contract method 0x62c29617.
//
// Solidity: function claimExpired(uint64 index) returns()
func (_Broker *BrokerTransactorSession) ClaimExpired(index uint64) (*types.Transaction, error) {
	return _Broker.Contract.ClaimExpired(&_Broker.TransactOpts, index)
}

// DepositCoin is a paid mutator transaction binding the contract method 0xca3b3b4b.
//
// Solidity: function depositCoin(uint256 numTokens) returns(bool)
func (_Broker *BrokerTransactor) DepositCoin(opts *bind.TransactOpts, numTokens *big.Int) (*types.Transaction, error) {
	return _Broker.contract.Transact(opts, "depositCoin", numTokens)
}

// DepositCoin is a paid mutator transaction binding the contract method 0xca3b3b4b.
//
// Solidity: function depositCoin(uint256 numTokens) returns(bool)
func (_Broker *BrokerSession) DepositCoin(numTokens *big.Int) (*types.Transaction, error) {
	return _Broker.Contract.DepositCoin(&_Broker.TransactOpts, numTokens)
}

// DepositCoin is a paid mutator transaction binding the contract method 0xca3b3b4b.
//
// Solidity: function depositCoin(uint256 numTokens) returns(bool)
func (_Broker *BrokerTransactorSession) DepositCoin(numTokens *big.Int) (*types.Transaction, error) {
	return _Broker.Contract.DepositCoin(&_Broker.TransactOpts, numTokens)
}

// ExtendBooking is a paid mutator transaction binding the contract method 0x9f60c50c.
//
// Solidity: function extendBooking(uint64 index, uint256 secs) returns()
func (_Broker *BrokerTransactor) ExtendBooking(opts *bind.TransactOpts, index uint64, secs *big.Int) (*types.Transaction, error) {
	return _Broker.contract.Transact(opts, "extendBooking", index, secs)
}

// ExtendBooking is a paid mutator transaction binding the contract method 0x9f60c50c.
//
// Solidity: function extendBooking(uint64 index, uint256 secs) returns()
func (_Broker *BrokerSession) ExtendBooking(index uint64, secs *big.Int) (*types.Transaction, error) {
	return _Broker.Contract.ExtendBooking(&_Broker.TransactOpts, index, secs)
}

// ExtendBooking is a paid mutator transaction binding the contract method 0x9f60c50c.
//
// Solidity: function extendBooking(uint64 index, uint256 secs) returns()
func (_Broker *BrokerTransactorSession) ExtendBooking(index uint64, secs *big.Int) (*types.Transaction, error) {
	return _Broker.Contract.ExtendBooking(&_Broker.TransactOpts, index, secs)
}

// SetCommunityContract is a paid mutator transaction binding the contract method 0xa35310d4.
//
// Solidity: function setCommunityContract(address newCommunityAddress) returns(bool)
func (_Broker *BrokerTransactor) SetCommunityContract(opts *bind.TransactOpts, newCommunityAddress common.Address) (*types.Transaction, error) {
	return _Broker.contract.Transact(opts, "setCommunityContract", newCommunityAddress)
}

// SetCommunityContract is a paid mutator transaction binding the contract method 0xa35310d4.
//
// Solidity: function setCommunityContract(address newCommunityAddress) returns(bool)
func (_Broker *BrokerSession) SetCommunityContract(newCommunityAddress common.Address) (*types.Transaction, error) {
	return _Broker.Contract.SetCommunityContract(&_Broker.TransactOpts, newCommunityAddress)
}

// SetCommunityContract is a paid mutator transaction binding the contract method 0xa35310d4.
//
// Solidity: function setCommunityContract(address newCommunityAddress) returns(bool)
func (_Broker *BrokerTransactorSession) SetCommunityContract(newCommunityAddress common.Address) (*types.Transaction, error) {
	return _Broker.Contract.SetCommunityContract(&_Broker.TransactOpts, newCommunityAddress)
}

// SetCommunityFee is a paid mutator transaction binding the contract method 0x0d12bbdb.
//
// Solidity: function setCommunityFee(uint256 fee) returns(bool)
func (_Broker *BrokerTransactor) SetCommunityFee(opts *bind.TransactOpts, fee *big.Int) (*types.Transaction, error) {
	return _Broker.contract.Transact(opts, "setCommunityFee", fee)
}

// SetCommunityFee is a paid mutator transaction binding the contract method 0x0d12bbdb.
//
// Solidity: function setCommunityFee(uint256 fee) returns(bool)
func (_Broker *BrokerSession) SetCommunityFee(fee *big.Int) (*types.Transaction, error) {
	return _Broker.Contract.SetCommunityFee(&_Broker.TransactOpts, fee)
}

// SetCommunityFee is a paid mutator transaction binding the contract method 0x0d12bbdb.
//
// Solidity: function setCommunityFee(uint256 fee) returns(bool)
func (_Broker *BrokerTransactorSession) SetCommunityFee(fee *big.Int) (*types.Transaction, error) {
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

// SetStablecoinAddress is a paid mutator transaction binding the contract method 0x9ecd0500.
//
// Solidity: function setStablecoinAddress(address t) returns(bool)
func (_Broker *BrokerTransactor) SetStablecoinAddress(opts *bind.TransactOpts, t common.Address) (*types.Transaction, error) {
	return _Broker.contract.Transact(opts, "setStablecoinAddress", t)
}

// SetStablecoinAddress is a paid mutator transaction binding the contract method 0x9ecd0500.
//
// Solidity: function setStablecoinAddress(address t) returns(bool)
func (_Broker *BrokerSession) SetStablecoinAddress(t common.Address) (*types.Transaction, error) {
	return _Broker.Contract.SetStablecoinAddress(&_Broker.TransactOpts, t)
}

// SetStablecoinAddress is a paid mutator transaction binding the contract method 0x9ecd0500.
//
// Solidity: function setStablecoinAddress(address t) returns(bool)
func (_Broker *BrokerTransactorSession) SetStablecoinAddress(t common.Address) (*types.Transaction, error) {
	return _Broker.Contract.SetStablecoinAddress(&_Broker.TransactOpts, t)
}

// WithdrawCoin is a paid mutator transaction binding the contract method 0x4eb264e8.
//
// Solidity: function withdrawCoin(uint256 numTokens) returns(bool)
func (_Broker *BrokerTransactor) WithdrawCoin(opts *bind.TransactOpts, numTokens *big.Int) (*types.Transaction, error) {
	return _Broker.contract.Transact(opts, "withdrawCoin", numTokens)
}

// WithdrawCoin is a paid mutator transaction binding the contract method 0x4eb264e8.
//
// Solidity: function withdrawCoin(uint256 numTokens) returns(bool)
func (_Broker *BrokerSession) WithdrawCoin(numTokens *big.Int) (*types.Transaction, error) {
	return _Broker.Contract.WithdrawCoin(&_Broker.TransactOpts, numTokens)
}

// WithdrawCoin is a paid mutator transaction binding the contract method 0x4eb264e8.
//
// Solidity: function withdrawCoin(uint256 numTokens) returns(bool)
func (_Broker *BrokerTransactorSession) WithdrawCoin(numTokens *big.Int) (*types.Transaction, error) {
	return _Broker.Contract.WithdrawCoin(&_Broker.TransactOpts, numTokens)
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
	MinerPayout *big.Int
	Index       *big.Int
	Miner       common.Address
	User        common.Address
	TimeUsed    *big.Int
	VmTypeId    *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBookingClaimed is a free log retrieval operation binding the contract event 0x75d2566632c43a31ab3bc2a82efa0c6d7797ad569fb798ae857716022614d90c.
//
// Solidity: event BookingClaimed(uint256 minerPayout, uint256 index, address indexed miner, address indexed user, uint256 timeUsed, uint256 vmTypeId)
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

// WatchBookingClaimed is a free log subscription operation binding the contract event 0x75d2566632c43a31ab3bc2a82efa0c6d7797ad569fb798ae857716022614d90c.
//
// Solidity: event BookingClaimed(uint256 minerPayout, uint256 index, address indexed miner, address indexed user, uint256 timeUsed, uint256 vmTypeId)
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

// ParseBookingClaimed is a log parse operation binding the contract event 0x75d2566632c43a31ab3bc2a82efa0c6d7797ad569fb798ae857716022614d90c.
//
// Solidity: event BookingClaimed(uint256 minerPayout, uint256 index, address indexed miner, address indexed user, uint256 timeUsed, uint256 vmTypeId)
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
	Index    *big.Int
	Miner    common.Address
	User     common.Address
	VmTypeId *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBookingExtended is a free log retrieval operation binding the contract event 0x075b18a02d51e5a00f1bf8c8adbeba55d6fc957987d4285794bcd90298147468.
//
// Solidity: event BookingExtended(uint256 index, address indexed miner, address indexed user, uint256 vmTypeId)
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

// WatchBookingExtended is a free log subscription operation binding the contract event 0x075b18a02d51e5a00f1bf8c8adbeba55d6fc957987d4285794bcd90298147468.
//
// Solidity: event BookingExtended(uint256 index, address indexed miner, address indexed user, uint256 vmTypeId)
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

// ParseBookingExtended is a log parse operation binding the contract event 0x075b18a02d51e5a00f1bf8c8adbeba55d6fc957987d4285794bcd90298147468.
//
// Solidity: event BookingExtended(uint256 index, address indexed miner, address indexed user, uint256 vmTypeId)
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
	MinerPayout *big.Int
	Index       *big.Int
	Miner       common.Address
	User        common.Address
	TimeUsed    *big.Int
	VmTypeId    *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBookingReported is a free log retrieval operation binding the contract event 0x26c93dbbeb58948b01dbae3d96c9cbe902a2f92082c8d91ed5f12e8579b4cf41.
//
// Solidity: event BookingReported(uint256 minerPayout, uint256 index, address indexed miner, address indexed user, uint256 timeUsed, uint256 vmTypeId)
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

// WatchBookingReported is a free log subscription operation binding the contract event 0x26c93dbbeb58948b01dbae3d96c9cbe902a2f92082c8d91ed5f12e8579b4cf41.
//
// Solidity: event BookingReported(uint256 minerPayout, uint256 index, address indexed miner, address indexed user, uint256 timeUsed, uint256 vmTypeId)
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

// ParseBookingReported is a log parse operation binding the contract event 0x26c93dbbeb58948b01dbae3d96c9cbe902a2f92082c8d91ed5f12e8579b4cf41.
//
// Solidity: event BookingReported(uint256 minerPayout, uint256 index, address indexed miner, address indexed user, uint256 timeUsed, uint256 vmTypeId)
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
	Index    *big.Int
	Miner    common.Address
	User     common.Address
	VmTypeId *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBookingStarted is a free log retrieval operation binding the contract event 0x246cbcaa18d1e49adb9a7db733d35a6e5a2690726bc6b2326bd3c53fa602f1b6.
//
// Solidity: event BookingStarted(uint256 index, address indexed miner, address indexed user, uint256 vmTypeId)
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

// WatchBookingStarted is a free log subscription operation binding the contract event 0x246cbcaa18d1e49adb9a7db733d35a6e5a2690726bc6b2326bd3c53fa602f1b6.
//
// Solidity: event BookingStarted(uint256 index, address indexed miner, address indexed user, uint256 vmTypeId)
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

// ParseBookingStarted is a log parse operation binding the contract event 0x246cbcaa18d1e49adb9a7db733d35a6e5a2690726bc6b2326bd3c53fa602f1b6.
//
// Solidity: event BookingStarted(uint256 index, address indexed miner, address indexed user, uint256 vmTypeId)
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
	MinerPayout *big.Int
	Index       *big.Int
	Miner       common.Address
	User        common.Address
	TimeUsed    *big.Int
	VmTypeId    *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBookingStopped is a free log retrieval operation binding the contract event 0x232ba2cbdb135f5ee58e4023da7094472434ff90d2d63f3c10232a45ee4c00d1.
//
// Solidity: event BookingStopped(uint256 minerPayout, uint256 index, address indexed miner, address indexed user, uint256 timeUsed, uint256 vmTypeId)
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

// WatchBookingStopped is a free log subscription operation binding the contract event 0x232ba2cbdb135f5ee58e4023da7094472434ff90d2d63f3c10232a45ee4c00d1.
//
// Solidity: event BookingStopped(uint256 minerPayout, uint256 index, address indexed miner, address indexed user, uint256 timeUsed, uint256 vmTypeId)
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

// ParseBookingStopped is a log parse operation binding the contract event 0x232ba2cbdb135f5ee58e4023da7094472434ff90d2d63f3c10232a45ee4c00d1.
//
// Solidity: event BookingStopped(uint256 minerPayout, uint256 index, address indexed miner, address indexed user, uint256 timeUsed, uint256 vmTypeId)
func (_Broker *BrokerFilterer) ParseBookingStopped(log types.Log) (*BrokerBookingStopped, error) {
	event := new(BrokerBookingStopped)
	if err := _Broker.contract.UnpackLog(event, "BookingStopped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
