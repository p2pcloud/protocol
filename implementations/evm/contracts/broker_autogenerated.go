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
	ABI: "[{\"inputs\":[],\"name\":\"GetTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes20\",\"name\":\"_signature\",\"type\":\"bytes20\"}],\"name\":\"SetMtlsHash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pricePerSecond\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vmTypeId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"machinesAvailable\",\"type\":\"uint256\"}],\"name\":\"addOffer\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offerIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"secs\",\"type\":\"uint256\"}],\"name\":\"bookVM\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"numTokens\",\"type\":\"uint256\"}],\"name\":\"depositCoin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_miner\",\"type\":\"address\"}],\"name\":\"findBookingsByMiner\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vmTypeId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pricePerSecond\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bookedTill\",\"type\":\"uint256\"}],\"internalType\":\"structBroker.Booking[]\",\"name\":\"filteredBookings\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"findBookingsByUser\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vmTypeId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pricePerSecond\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bookedTill\",\"type\":\"uint256\"}],\"internalType\":\"structBroker.Booking[]\",\"name\":\"filteredBookings\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_vmTypeId\",\"type\":\"uint256\"}],\"name\":\"getAvailableOffers\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pricePerSecond\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"machinesAvailable\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vmTypeId\",\"type\":\"uint256\"}],\"internalType\":\"structBroker.VMOffer[]\",\"name\":\"filteredOffers\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"}],\"name\":\"getBooking\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vmTypeId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pricePerSecond\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bookedTill\",\"type\":\"uint256\"}],\"internalType\":\"structBroker.Booking\",\"name\":\"booking\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCoinDecimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCommunityContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCommunityFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getMinerUrl\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"}],\"name\":\"getMinersOffers\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pricePerSecond\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"machinesAvailable\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vmTypeId\",\"type\":\"uint256\"}],\"internalType\":\"structBroker.VMOffer[]\",\"name\":\"filteredOffers\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getMtlsHash\",\"outputs\":[{\"internalType\":\"bytes20\",\"name\":\"\",\"type\":\"bytes20\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStablecoinAddress\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"communityAddress\",\"type\":\"address\"}],\"name\":\"setCommunityContract\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"setCommunityFee\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"url\",\"type\":\"bytes32\"}],\"name\":\"setMunerUrl\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"t\",\"type\":\"address\"}],\"name\":\"setStablecoinAddress\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offerIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pricePerSecond\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vmTypeId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"machinesAvailable\",\"type\":\"uint256\"}],\"name\":\"updateOffer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"userAllowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"userBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"userTokenBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"numTokens\",\"type\":\"uint256\"}],\"name\":\"withdrawCoin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040526000600155600060035534801561001a57600080fd5b50612d558061002a6000396000f3fe608060405234801561001057600080fd5b506004361061018e5760003560e01c80637ee080c6116100de578063c700ff0a11610097578063cf696d1411610071578063cf696d141461050f578063e09880131461052d578063ea5f398414610549578063f44f11cd146105795761018e565b8063c700ff0a1461047f578063ca3b3b4b146104af578063cbb6bfbd146104df5761018e565b80637ee080c61461038357806390ad688b146103b35780639ecd0500146103d1578063a35310d414610401578063ab49ec4514610431578063bf152765146104615761018e565b8063301bcaae1161014b57806354fd4d501161012557806354fd4d50146102f9578063569e2e0f146103175780635a610b57146103355780637bef4f54146103655761018e565b8063301bcaae146102695780633b3e328b146102995780634eb264e8146102c95761018e565b80630d12bbdb14610193578063117298e4146101c357806318c5e502146101df57806320255c7e146101fd5780632d0d6c331461021b5780632dc8a2b914610239575b600080fd5b6101ad60048036038101906101a89190611ed8565b610595565b6040516101ba9190611f20565b60405180910390f35b6101dd60048036038101906101d89190611f93565b610685565b005b6101e76106f3565b6040516101f49190611fcf565b60405180910390f35b6102056106fb565b6040516102129190611fcf565b60405180910390f35b6102236107a0565b6040516102309190611fcf565b60405180910390f35b610253600480360381019061024e919061202a565b610843565b6040516102609190612122565b60405180910390f35b610283600480360381019061027e9190612169565b61094e565b60405161029091906122c0565b60405180910390f35b6102b360048036038101906102ae91906122e2565b610c0e565b6040516102c09190611fcf565b60405180910390f35b6102e360048036038101906102de9190611ed8565b610d04565b6040516102f09190611f20565b60405180910390f35b610301610e95565b60405161030e9190611fcf565b60405180910390f35b61031f610e9a565b60405161032c9190612344565b60405180910390f35b61034f600480360381019061034a9190612169565b610ec4565b60405161035c91906122c0565b60405180910390f35b61036d611184565b60405161037a91906123be565b60405180910390f35b61039d60048036038101906103989190611ed8565b6111ae565b6040516103aa91906124f0565b60405180910390f35b6103bb6113ec565b6040516103c89190611fcf565b60405180910390f35b6103eb60048036038101906103e69190612550565b6113f6565b6040516103f89190611f20565b60405180910390f35b61041b60048036038101906104169190612169565b6114d2565b6040516104289190611f20565b60405180910390f35b61044b60048036038101906104469190612169565b61151e565b6040516104589190612596565b60405180910390f35b610469611567565b6040516104769190611fcf565b60405180910390f35b61049960048036038101906104949190612169565b6115f8565b6040516104a691906124f0565b60405180910390f35b6104c960048036038101906104c49190611ed8565b611860565b6040516104d69190611f20565b60405180910390f35b6104f960048036038101906104f491906125b1565b6119a9565b6040516105069190611fcf565b60405180910390f35b610517611bca565b604051610524919061260d565b60405180910390f35b61054760048036038101906105429190612628565b611c62565b005b610563600480360381019061055e9190612169565b611d59565b604051610570919061269e565b60405180910390f35b610593600480360381019061058e91906126e5565b611daf565b005b600080821180156105a65750606582105b6105e5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105dc90612795565b60405180910390fd5b600960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610675576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161066c90612827565b60405180910390fd5b81600a8190555060009050919050565b80600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908360601c021790555050565b600042905090565b6000600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663dd62ed3e33306040518363ffffffff1660e01b815260040161075a929190612847565b602060405180830381865afa158015610777573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061079b9190612885565b905090565b6000600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231336040518263ffffffff1660e01b81526004016107fd9190612344565b602060405180830381865afa15801561081a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061083e9190612885565b905090565b61084b611df6565b600260008367ffffffffffffffff1681526020019081526020016000206040518060c001604052908160008201548152602001600182015481526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016003820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600482015481526020016005820154815250509050919050565b6060600060035467ffffffffffffffff81111561096e5761096d6128b2565b5b6040519080825280602002602001820160405280156109a757816020015b610994611df6565b81526020019060019003908161098c5790505b509050600080600090505b600354811015610b58578473ffffffffffffffffffffffffffffffffffffffff166002600083815260200190815260200160002060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603610b4557600260008281526020019081526020016000206040518060c001604052908160008201548152602001600182015481526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016003820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160048201548152602001600582015481525050838381518110610b2a57610b296128e1565b5b6020026020010181905250600182610b42919061293f565b91505b8080610b5090612995565b9150506109b2565b508067ffffffffffffffff811115610b7357610b726128b2565b5b604051908082528060200260200182016040528015610bac57816020015b610b99611df6565b815260200190600190039081610b915790505b50925060005b81811015610c0657828181518110610bcd57610bcc6128e1565b5b6020026020010151848281518110610be857610be76128e1565b5b60200260200101819052508080610bfe90612995565b915050610bb2565b505050919050565b60006040518060a0016040528060015481526020013373ffffffffffffffffffffffffffffffffffffffff1681526020018581526020018381526020018481525060008060015481526020019081526020016000206000820151816000015560208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160020155606082015181600301556080820151816004015590505060016000815480929190610ce890612995565b919050555060018054610cfb91906129dd565b90509392505050565b600081610d0f611567565b1015610d50576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d4790612a5d565b60405180910390fd5b600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb33846040518363ffffffff1660e01b8152600401610dad929190612a7d565b6020604051808303816000875af1158015610dcc573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610df09190612ad2565b610dfd5760009050610e90565b81600660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054610e4891906129dd565b600660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550600190505b919050565b600181565b6000600960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6060600060035467ffffffffffffffff811115610ee457610ee36128b2565b5b604051908082528060200260200182016040528015610f1d57816020015b610f0a611df6565b815260200190600190039081610f025790505b509050600080600090505b6003548110156110ce578473ffffffffffffffffffffffffffffffffffffffff166002600083815260200190815260200160002060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16036110bb57600260008281526020019081526020016000206040518060c001604052908160008201548152602001600182015481526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016003820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600482015481526020016005820154815250508383815181106110a05761109f6128e1565b5b60200260200101819052506001826110b8919061293f565b91505b80806110c690612995565b915050610f28565b508067ffffffffffffffff8111156110e9576110e86128b2565b5b60405190808252806020026020018201604052801561112257816020015b61110f611df6565b8152602001906001900390816111075790505b50925060005b8181101561117c57828181518110611143576111426128e1565b5b602002602001015184828151811061115e5761115d6128e1565b5b6020026020010181905250808061117490612995565b915050611128565b505050919050565b6000600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6060600060015467ffffffffffffffff8111156111ce576111cd6128b2565b5b60405190808252806020026020018201604052801561120757816020015b6111f4611e58565b8152602001906001900390816111ec5790505b509050600080600090505b6001548110156113365784600080838152602001908152602001600020600401541480156112555750600080600083815260200190815260200160002060030154115b15611323576000808281526020019081526020016000206040518060a0016040529081600082015481526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820154815260200160038201548152602001600482015481525050838381518110611308576113076128e1565b5b6020026020010181905250600182611320919061293f565b91505b808061132e90612995565b915050611212565b508067ffffffffffffffff811115611351576113506128b2565b5b60405190808252806020026020018201604052801561138a57816020015b611377611e58565b81526020019060019003908161136f5790505b50925060005b818110156113e4578281815181106113ab576113aa6128e1565b5b60200260200101518482815181106113c6576113c56128e1565b5b602002602001018190525080806113dc90612995565b915050611390565b505050919050565b6000600a54905090565b6000600960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611488576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161147f90612b71565b60405180910390fd5b81600860006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060009050919050565b600081600960006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060009050919050565b6000600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b6000600760003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054600660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546115f391906129dd565b905090565b6060600060015467ffffffffffffffff811115611618576116176128b2565b5b60405190808252806020026020018201604052801561165157816020015b61163e611e58565b8152602001906001900390816116365790505b509050600080600090505b6001548110156117aa578473ffffffffffffffffffffffffffffffffffffffff1660008083815260200190815260200160002060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603611797576000808281526020019081526020016000206040518060a0016040529081600082015481526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600282015481526020016003820154815260200160048201548152505083838151811061177c5761177b6128e1565b5b6020026020010181905250600182611794919061293f565b91505b80806117a290612995565b91505061165c565b508067ffffffffffffffff8111156117c5576117c46128b2565b5b6040519080825280602002602001820160405280156117fe57816020015b6117eb611e58565b8152602001906001900390816117e35790505b50925060005b818110156118585782818151811061181f5761181e6128e1565b5b602002602001015184828151811061183a576118396128e1565b5b6020026020010181905250808061185090612995565b915050611804565b505050919050565b6000600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd3330856040518463ffffffff1660e01b81526004016118c193929190612b91565b6020604051808303816000875af11580156118e0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906119049190612ad2565b61191157600090506119a4565b81600660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461195c919061293f565b600660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550600190505b919050565b6000806000808581526020019081526020016000206003015411611a02576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016119f990612c14565b60405180910390fd5b60006040518060c00160405280600354815260200160008087815260200190815260200160002060040154815260200160008087815260200190815260200160002060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020013373ffffffffffffffffffffffffffffffffffffffff1681526020016000808781526020019081526020016000206002015481526020018442611ac6919061293f565b815250905080600260006003548152602001908152602001600020600082015181600001556020820151816001015560408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506080820151816004015560a0820151816005015590505060036000815480929190611bad90612995565b91905055506001600354611bc191906129dd565b91505092915050565b6000600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663313ce5676040518163ffffffff1660e01b8152600401602060405180830381865afa158015611c39573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611c5d9190612c60565b905090565b3373ffffffffffffffffffffffffffffffffffffffff1660008086815260200190815260200160002060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614611d05576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611cfc90612cff565b60405180910390fd5b82600080868152602001908152602001600020600201819055508060008086815260200190815260200160002060030181905550816000808681526020019081526020016000206004018190555050505050565b6000600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460601b9050919050565b80600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555050565b6040518060c001604052806000815260200160008152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff16815260200160008152602001600081525090565b6040518060a0016040528060008152602001600073ffffffffffffffffffffffffffffffffffffffff1681526020016000815260200160008152602001600081525090565b600080fd5b6000819050919050565b611eb581611ea2565b8114611ec057600080fd5b50565b600081359050611ed281611eac565b92915050565b600060208284031215611eee57611eed611e9d565b5b6000611efc84828501611ec3565b91505092915050565b60008115159050919050565b611f1a81611f05565b82525050565b6000602082019050611f356000830184611f11565b92915050565b60007fffffffffffffffffffffffffffffffffffffffff00000000000000000000000082169050919050565b611f7081611f3b565b8114611f7b57600080fd5b50565b600081359050611f8d81611f67565b92915050565b600060208284031215611fa957611fa8611e9d565b5b6000611fb784828501611f7e565b91505092915050565b611fc981611ea2565b82525050565b6000602082019050611fe46000830184611fc0565b92915050565b600067ffffffffffffffff82169050919050565b61200781611fea565b811461201257600080fd5b50565b60008135905061202481611ffe565b92915050565b6000602082840312156120405761203f611e9d565b5b600061204e84828501612015565b91505092915050565b61206081611ea2565b82525050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061209182612066565b9050919050565b6120a181612086565b82525050565b60c0820160008201516120bd6000850182612057565b5060208201516120d06020850182612057565b5060408201516120e36040850182612098565b5060608201516120f66060850182612098565b5060808201516121096080850182612057565b5060a082015161211c60a0850182612057565b50505050565b600060c08201905061213760008301846120a7565b92915050565b61214681612086565b811461215157600080fd5b50565b6000813590506121638161213d565b92915050565b60006020828403121561217f5761217e611e9d565b5b600061218d84828501612154565b91505092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b60c0820160008201516121d86000850182612057565b5060208201516121eb6020850182612057565b5060408201516121fe6040850182612098565b5060608201516122116060850182612098565b5060808201516122246080850182612057565b5060a082015161223760a0850182612057565b50505050565b600061224983836121c2565b60c08301905092915050565b6000602082019050919050565b600061226d82612196565b61227781856121a1565b9350612282836121b2565b8060005b838110156122b357815161229a888261223d565b97506122a583612255565b925050600181019050612286565b5085935050505092915050565b600060208201905081810360008301526122da8184612262565b905092915050565b6000806000606084860312156122fb576122fa611e9d565b5b600061230986828701611ec3565b935050602061231a86828701611ec3565b925050604061232b86828701611ec3565b9150509250925092565b61233e81612086565b82525050565b60006020820190506123596000830184612335565b92915050565b6000819050919050565b600061238461237f61237a84612066565b61235f565b612066565b9050919050565b600061239682612369565b9050919050565b60006123a88261238b565b9050919050565b6123b88161239d565b82525050565b60006020820190506123d360008301846123af565b92915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b60a08201600082015161241b6000850182612057565b50602082015161242e6020850182612098565b5060408201516124416040850182612057565b5060608201516124546060850182612057565b5060808201516124676080850182612057565b50505050565b60006124798383612405565b60a08301905092915050565b6000602082019050919050565b600061249d826123d9565b6124a781856123e4565b93506124b2836123f5565b8060005b838110156124e35781516124ca888261246d565b97506124d583612485565b9250506001810190506124b6565b5085935050505092915050565b6000602082019050818103600083015261250a8184612492565b905092915050565b600061251d82612086565b9050919050565b61252d81612512565b811461253857600080fd5b50565b60008135905061254a81612524565b92915050565b60006020828403121561256657612565611e9d565b5b60006125748482850161253b565b91505092915050565b6000819050919050565b6125908161257d565b82525050565b60006020820190506125ab6000830184612587565b92915050565b600080604083850312156125c8576125c7611e9d565b5b60006125d685828601611ec3565b92505060206125e785828601611ec3565b9150509250929050565b600060ff82169050919050565b612607816125f1565b82525050565b600060208201905061262260008301846125fe565b92915050565b6000806000806080858703121561264257612641611e9d565b5b600061265087828801611ec3565b945050602061266187828801611ec3565b935050604061267287828801611ec3565b925050606061268387828801611ec3565b91505092959194509250565b61269881611f3b565b82525050565b60006020820190506126b3600083018461268f565b92915050565b6126c28161257d565b81146126cd57600080fd5b50565b6000813590506126df816126b9565b92915050565b6000602082840312156126fb576126fa611e9d565b5b6000612709848285016126d0565b91505092915050565b600082825260208201905092915050565b7f636f6d6d756e697479206665652073686f756c6420626520696e2072616e676560008201527f206f66203120746f203130300000000000000000000000000000000000000000602082015250565b600061277f602c83612712565b915061278a82612723565b604082019050919050565b600060208201905081810360008301526127ae81612772565b9050919050565b7f6f6e6c7920636f6d6d756e69747920636f6e74726163742063616e207365742060008201527f6665650000000000000000000000000000000000000000000000000000000000602082015250565b6000612811602383612712565b915061281c826127b5565b604082019050919050565b6000602082019050818103600083015261284081612804565b9050919050565b600060408201905061285c6000830185612335565b6128696020830184612335565b9392505050565b60008151905061287f81611eac565b92915050565b60006020828403121561289b5761289a611e9d565b5b60006128a984828501612870565b91505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061294a82611ea2565b915061295583611ea2565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0382111561298a57612989612910565b5b828201905092915050565b60006129a082611ea2565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036129d2576129d1612910565b5b600182019050919050565b60006129e882611ea2565b91506129f383611ea2565b925082821015612a0657612a05612910565b5b828203905092915050565b7f696e73756666696369656e742066756e647320746f2077697468647261770000600082015250565b6000612a47601e83612712565b9150612a5282612a11565b602082019050919050565b60006020820190508181036000830152612a7681612a3a565b9050919050565b6000604082019050612a926000830185612335565b612a9f6020830184611fc0565b9392505050565b612aaf81611f05565b8114612aba57600080fd5b50565b600081519050612acc81612aa6565b92915050565b600060208284031215612ae857612ae7611e9d565b5b6000612af684828501612abd565b91505092915050565b7f6f6e6c7920636f6d6d756e69747920636f6e74726163742063616e207365742060008201527f737461626c65636f696e00000000000000000000000000000000000000000000602082015250565b6000612b5b602a83612712565b9150612b6682612aff565b604082019050919050565b60006020820190508181036000830152612b8a81612b4e565b9050919050565b6000606082019050612ba66000830186612335565b612bb36020830185612335565b612bc06040830184611fc0565b949350505050565b7f4e6f206d616368696e657320617661696c61626c650000000000000000000000600082015250565b6000612bfe601583612712565b9150612c0982612bc8565b602082019050919050565b60006020820190508181036000830152612c2d81612bf1565b9050919050565b612c3d816125f1565b8114612c4857600080fd5b50565b600081519050612c5a81612c34565b92915050565b600060208284031215612c7657612c75611e9d565b5b6000612c8484828501612c4b565b91505092915050565b7f4f6e6c7920746865206f776e65722063616e2075706461746520616e206f666660008201527f6572000000000000000000000000000000000000000000000000000000000000602082015250565b6000612ce9602283612712565b9150612cf482612c8d565b604082019050919050565b60006020820190508181036000830152612d1881612cdc565b905091905056fea2646970667358221220ac493e38ab2950e39a42d0e4ad98449d8ac659e6bd41c9c2c7197003d474f36064736f6c634300080f0033",
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
// Solidity: function findBookingsByMiner(address _miner) view returns((uint256,uint256,address,address,uint256,uint256)[] filteredBookings)
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
// Solidity: function findBookingsByMiner(address _miner) view returns((uint256,uint256,address,address,uint256,uint256)[] filteredBookings)
func (_Broker *BrokerSession) FindBookingsByMiner(_miner common.Address) ([]BrokerBooking, error) {
	return _Broker.Contract.FindBookingsByMiner(&_Broker.CallOpts, _miner)
}

// FindBookingsByMiner is a free data retrieval call binding the contract method 0x301bcaae.
//
// Solidity: function findBookingsByMiner(address _miner) view returns((uint256,uint256,address,address,uint256,uint256)[] filteredBookings)
func (_Broker *BrokerCallerSession) FindBookingsByMiner(_miner common.Address) ([]BrokerBooking, error) {
	return _Broker.Contract.FindBookingsByMiner(&_Broker.CallOpts, _miner)
}

// FindBookingsByUser is a free data retrieval call binding the contract method 0x5a610b57.
//
// Solidity: function findBookingsByUser(address _owner) view returns((uint256,uint256,address,address,uint256,uint256)[] filteredBookings)
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
// Solidity: function findBookingsByUser(address _owner) view returns((uint256,uint256,address,address,uint256,uint256)[] filteredBookings)
func (_Broker *BrokerSession) FindBookingsByUser(_owner common.Address) ([]BrokerBooking, error) {
	return _Broker.Contract.FindBookingsByUser(&_Broker.CallOpts, _owner)
}

// FindBookingsByUser is a free data retrieval call binding the contract method 0x5a610b57.
//
// Solidity: function findBookingsByUser(address _owner) view returns((uint256,uint256,address,address,uint256,uint256)[] filteredBookings)
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
// Solidity: function getBooking(uint64 index) view returns((uint256,uint256,address,address,uint256,uint256) booking)
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
// Solidity: function getBooking(uint64 index) view returns((uint256,uint256,address,address,uint256,uint256) booking)
func (_Broker *BrokerSession) GetBooking(index uint64) (BrokerBooking, error) {
	return _Broker.Contract.GetBooking(&_Broker.CallOpts, index)
}

// GetBooking is a free data retrieval call binding the contract method 0x2dc8a2b9.
//
// Solidity: function getBooking(uint64 index) view returns((uint256,uint256,address,address,uint256,uint256) booking)
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

// GetMtlsHash is a free data retrieval call binding the contract method 0xea5f3984.
//
// Solidity: function getMtlsHash(address _user) view returns(bytes20)
func (_Broker *BrokerCaller) GetMtlsHash(opts *bind.CallOpts, _user common.Address) ([20]byte, error) {
	var out []interface{}
	err := _Broker.contract.Call(opts, &out, "getMtlsHash", _user)

	if err != nil {
		return *new([20]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([20]byte)).(*[20]byte)

	return out0, err

}

// GetMtlsHash is a free data retrieval call binding the contract method 0xea5f3984.
//
// Solidity: function getMtlsHash(address _user) view returns(bytes20)
func (_Broker *BrokerSession) GetMtlsHash(_user common.Address) ([20]byte, error) {
	return _Broker.Contract.GetMtlsHash(&_Broker.CallOpts, _user)
}

// GetMtlsHash is a free data retrieval call binding the contract method 0xea5f3984.
//
// Solidity: function getMtlsHash(address _user) view returns(bytes20)
func (_Broker *BrokerCallerSession) GetMtlsHash(_user common.Address) ([20]byte, error) {
	return _Broker.Contract.GetMtlsHash(&_Broker.CallOpts, _user)
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

// SetMtlsHash is a paid mutator transaction binding the contract method 0x117298e4.
//
// Solidity: function SetMtlsHash(bytes20 _signature) returns()
func (_Broker *BrokerTransactor) SetMtlsHash(opts *bind.TransactOpts, _signature [20]byte) (*types.Transaction, error) {
	return _Broker.contract.Transact(opts, "SetMtlsHash", _signature)
}

// SetMtlsHash is a paid mutator transaction binding the contract method 0x117298e4.
//
// Solidity: function SetMtlsHash(bytes20 _signature) returns()
func (_Broker *BrokerSession) SetMtlsHash(_signature [20]byte) (*types.Transaction, error) {
	return _Broker.Contract.SetMtlsHash(&_Broker.TransactOpts, _signature)
}

// SetMtlsHash is a paid mutator transaction binding the contract method 0x117298e4.
//
// Solidity: function SetMtlsHash(bytes20 _signature) returns()
func (_Broker *BrokerTransactorSession) SetMtlsHash(_signature [20]byte) (*types.Transaction, error) {
	return _Broker.Contract.SetMtlsHash(&_Broker.TransactOpts, _signature)
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

// SetCommunityContract is a paid mutator transaction binding the contract method 0xa35310d4.
//
// Solidity: function setCommunityContract(address communityAddress) returns(bool)
func (_Broker *BrokerTransactor) SetCommunityContract(opts *bind.TransactOpts, communityAddress common.Address) (*types.Transaction, error) {
	return _Broker.contract.Transact(opts, "setCommunityContract", communityAddress)
}

// SetCommunityContract is a paid mutator transaction binding the contract method 0xa35310d4.
//
// Solidity: function setCommunityContract(address communityAddress) returns(bool)
func (_Broker *BrokerSession) SetCommunityContract(communityAddress common.Address) (*types.Transaction, error) {
	return _Broker.Contract.SetCommunityContract(&_Broker.TransactOpts, communityAddress)
}

// SetCommunityContract is a paid mutator transaction binding the contract method 0xa35310d4.
//
// Solidity: function setCommunityContract(address communityAddress) returns(bool)
func (_Broker *BrokerTransactorSession) SetCommunityContract(communityAddress common.Address) (*types.Transaction, error) {
	return _Broker.Contract.SetCommunityContract(&_Broker.TransactOpts, communityAddress)
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

// UpdateOffer is a paid mutator transaction binding the contract method 0xe0988013.
//
// Solidity: function updateOffer(uint256 offerIndex, uint256 pricePerSecond, uint256 vmTypeId, uint256 machinesAvailable) returns()
func (_Broker *BrokerTransactor) UpdateOffer(opts *bind.TransactOpts, offerIndex *big.Int, pricePerSecond *big.Int, vmTypeId *big.Int, machinesAvailable *big.Int) (*types.Transaction, error) {
	return _Broker.contract.Transact(opts, "updateOffer", offerIndex, pricePerSecond, vmTypeId, machinesAvailable)
}

// UpdateOffer is a paid mutator transaction binding the contract method 0xe0988013.
//
// Solidity: function updateOffer(uint256 offerIndex, uint256 pricePerSecond, uint256 vmTypeId, uint256 machinesAvailable) returns()
func (_Broker *BrokerSession) UpdateOffer(offerIndex *big.Int, pricePerSecond *big.Int, vmTypeId *big.Int, machinesAvailable *big.Int) (*types.Transaction, error) {
	return _Broker.Contract.UpdateOffer(&_Broker.TransactOpts, offerIndex, pricePerSecond, vmTypeId, machinesAvailable)
}

// UpdateOffer is a paid mutator transaction binding the contract method 0xe0988013.
//
// Solidity: function updateOffer(uint256 offerIndex, uint256 pricePerSecond, uint256 vmTypeId, uint256 machinesAvailable) returns()
func (_Broker *BrokerTransactorSession) UpdateOffer(offerIndex *big.Int, pricePerSecond *big.Int, vmTypeId *big.Int, machinesAvailable *big.Int) (*types.Transaction, error) {
	return _Broker.Contract.UpdateOffer(&_Broker.TransactOpts, offerIndex, pricePerSecond, vmTypeId, machinesAvailable)
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
