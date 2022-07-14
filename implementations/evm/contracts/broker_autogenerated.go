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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"GetTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes20\",\"name\":\"_signature\",\"type\":\"bytes20\"}],\"name\":\"SetMtlsHash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pricePerSecond\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vmTypeId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"machinesAvailable\",\"type\":\"uint256\"}],\"name\":\"addOffer\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offerIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"secs\",\"type\":\"uint256\"}],\"name\":\"bookVM\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"numTokens\",\"type\":\"uint256\"}],\"name\":\"depositCoin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_miner\",\"type\":\"address\"}],\"name\":\"findBookingsByMiner\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vmTypeId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pricePerSecond\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bookedTill\",\"type\":\"uint256\"}],\"internalType\":\"structBroker.Booking[]\",\"name\":\"filteredBookings\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"findBookingsByUser\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vmTypeId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pricePerSecond\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bookedTill\",\"type\":\"uint256\"}],\"internalType\":\"structBroker.Booking[]\",\"name\":\"filteredBookings\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_vmTypeId\",\"type\":\"uint256\"}],\"name\":\"getAvailableOffers\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pricePerSecond\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"machinesAvailable\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vmTypeId\",\"type\":\"uint256\"}],\"internalType\":\"structBroker.VMOffer[]\",\"name\":\"filteredOffers\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"}],\"name\":\"getBooking\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vmTypeId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pricePerSecond\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bookedTill\",\"type\":\"uint256\"}],\"internalType\":\"structBroker.Booking\",\"name\":\"booking\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getMinerUrl\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"}],\"name\":\"getMinersOffers\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pricePerSecond\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"machinesAvailable\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vmTypeId\",\"type\":\"uint256\"}],\"internalType\":\"structBroker.VMOffer[]\",\"name\":\"filteredOffers\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getMtlsHash\",\"outputs\":[{\"internalType\":\"bytes20\",\"name\":\"\",\"type\":\"bytes20\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"url\",\"type\":\"bytes32\"}],\"name\":\"setMunerUrl\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"testApprove\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offerIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pricePerSecond\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vmTypeId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"machinesAvailable\",\"type\":\"uint256\"}],\"name\":\"updateOffer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"userAllowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"userBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"userDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"userTokenBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"numTokens\",\"type\":\"uint256\"}],\"name\":\"withdrawCoin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040526000600155600060035534801561001a57600080fd5b5073772c89b7b306b88a7b72c66535506e5811ba30da600860006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506127378061007f6000396000f3fe608060405234801561001057600080fd5b50600436106101375760003560e01c80635a610b57116100b8578063ca3b3b4b1161007c578063ca3b3b4b1461039e578063cbb6bfbd146103ce578063e0988013146103fe578063ea5f39841461041a578063f44f11cd1461044a578063fb44e83a1461046657610137565b80635a610b57146102c05780637ee080c6146102f0578063ab49ec4514610320578063bf15276514610350578063c700ff0a1461036e57610137565b8063301bcaae116100ff578063301bcaae146101f45780633b3e328b1461022457806348146113146102545780634eb264e81461027257806354fd4d50146102a257610137565b80630c6b2cbf1461013c578063117298e41461016c57806318c5e502146101885780632d0d6c33146101a65780632dc8a2b9146101c4575b600080fd5b61015660048036038101906101519190611bcd565b610482565b6040516101639190611c13565b60405180910390f35b61018660048036038101906101819190611c86565b610529565b005b610190610597565b60405161019d9190611c13565b60405180910390f35b6101ae61059f565b6040516101bb9190611c13565b60405180910390f35b6101de60048036038101906101d99190611cf3565b610642565b6040516101eb9190611db9565b60405180910390f35b61020e60048036038101906102099190611bcd565b61074d565b60405161021b9190611efe565b60405180910390f35b61023e60048036038101906102399190611f4c565b610a0d565b60405161024b9190611c13565b60405180910390f35b61025c610b03565b6040516102699190611c13565b60405180910390f35b61028c60048036038101906102879190611f9f565b610b4a565b6040516102999190611fe7565b60405180910390f35b6102aa610cdb565b6040516102b79190611c13565b60405180910390f35b6102da60048036038101906102d59190611bcd565b610ce0565b6040516102e79190611efe565b60405180910390f35b61030a60048036038101906103059190611f9f565b610fa0565b6040516103179190612119565b60405180910390f35b61033a60048036038101906103359190611bcd565b6111de565b6040516103479190612154565b60405180910390f35b610358611227565b6040516103659190611c13565b60405180910390f35b61038860048036038101906103839190611bcd565b6112b8565b6040516103959190612119565b60405180910390f35b6103b860048036038101906103b39190611f9f565b611520565b6040516103c59190611fe7565b60405180910390f35b6103e860048036038101906103e3919061216f565b611669565b6040516103f59190611c13565b60405180910390f35b610418600480360381019061041391906121af565b61188a565b005b610434600480360381019061042f9190611bcd565b611981565b6040516104419190612225565b60405180910390f35b610464600480360381019061045f919061226c565b6119d7565b005b610480600480360381019061047b9190612299565b611a1e565b005b6000600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663dd62ed3e33846040518363ffffffff1660e01b81526004016104e19291906122e8565b602060405180830381865afa1580156104fe573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105229190612326565b9050919050565b80600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908360601c021790555050565b600042905090565b6000600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231336040518263ffffffff1660e01b81526004016105fc9190612353565b602060405180830381865afa158015610619573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061063d9190612326565b905090565b61064a611ac3565b600260008367ffffffffffffffff1681526020019081526020016000206040518060c001604052908160008201548152602001600182015481526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016003820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600482015481526020016005820154815250509050919050565b6060600060035467ffffffffffffffff81111561076d5761076c61236e565b5b6040519080825280602002602001820160405280156107a657816020015b610793611ac3565b81526020019060019003908161078b5790505b509050600080600090505b600354811015610957578473ffffffffffffffffffffffffffffffffffffffff166002600083815260200190815260200160002060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff160361094457600260008281526020019081526020016000206040518060c001604052908160008201548152602001600182015481526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016003820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600482015481526020016005820154815250508383815181106109295761092861239d565b5b602002602001018190525060018261094191906123fb565b91505b808061094f90612451565b9150506107b1565b508067ffffffffffffffff8111156109725761097161236e565b5b6040519080825280602002602001820160405280156109ab57816020015b610998611ac3565b8152602001906001900390816109905790505b50925060005b81811015610a05578281815181106109cc576109cb61239d565b5b60200260200101518482815181106109e7576109e661239d565b5b602002602001018190525080806109fd90612451565b9150506109b1565b505050919050565b60006040518060a0016040528060015481526020013373ffffffffffffffffffffffffffffffffffffffff1681526020018581526020018381526020018481525060008060015481526020019081526020016000206000820151816000015560208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160020155606082015181600301556080820151816004015590505060016000815480929190610ae790612451565b919050555060018054610afa9190612499565b90509392505050565b6000600660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905090565b600081610b55611227565b1015610b96576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b8d9061252a565b60405180910390fd5b600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb33846040518363ffffffff1660e01b8152600401610bf392919061254a565b6020604051808303816000875af1158015610c12573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c36919061259f565b610c435760009050610cd6565b81600660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054610c8e9190612499565b600660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550600190505b919050565b600181565b6060600060035467ffffffffffffffff811115610d0057610cff61236e565b5b604051908082528060200260200182016040528015610d3957816020015b610d26611ac3565b815260200190600190039081610d1e5790505b509050600080600090505b600354811015610eea578473ffffffffffffffffffffffffffffffffffffffff166002600083815260200190815260200160002060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603610ed757600260008281526020019081526020016000206040518060c001604052908160008201548152602001600182015481526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016003820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160048201548152602001600582015481525050838381518110610ebc57610ebb61239d565b5b6020026020010181905250600182610ed491906123fb565b91505b8080610ee290612451565b915050610d44565b508067ffffffffffffffff811115610f0557610f0461236e565b5b604051908082528060200260200182016040528015610f3e57816020015b610f2b611ac3565b815260200190600190039081610f235790505b50925060005b81811015610f9857828181518110610f5f57610f5e61239d565b5b6020026020010151848281518110610f7a57610f7961239d565b5b60200260200101819052508080610f9090612451565b915050610f44565b505050919050565b6060600060015467ffffffffffffffff811115610fc057610fbf61236e565b5b604051908082528060200260200182016040528015610ff957816020015b610fe6611b25565b815260200190600190039081610fde5790505b509050600080600090505b6001548110156111285784600080838152602001908152602001600020600401541480156110475750600080600083815260200190815260200160002060030154115b15611115576000808281526020019081526020016000206040518060a0016040529081600082015481526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160028201548152602001600382015481526020016004820154815250508383815181106110fa576110f961239d565b5b602002602001018190525060018261111291906123fb565b91505b808061112090612451565b915050611004565b508067ffffffffffffffff8111156111435761114261236e565b5b60405190808252806020026020018201604052801561117c57816020015b611169611b25565b8152602001906001900390816111615790505b50925060005b818110156111d65782818151811061119d5761119c61239d565b5b60200260200101518482815181106111b8576111b761239d565b5b602002602001018190525080806111ce90612451565b915050611182565b505050919050565b6000600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b6000600760003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054600660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546112b39190612499565b905090565b6060600060015467ffffffffffffffff8111156112d8576112d761236e565b5b60405190808252806020026020018201604052801561131157816020015b6112fe611b25565b8152602001906001900390816112f65790505b509050600080600090505b60015481101561146a578473ffffffffffffffffffffffffffffffffffffffff1660008083815260200190815260200160002060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603611457576000808281526020019081526020016000206040518060a0016040529081600082015481526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600282015481526020016003820154815260200160048201548152505083838151811061143c5761143b61239d565b5b602002602001018190525060018261145491906123fb565b91505b808061146290612451565b91505061131c565b508067ffffffffffffffff8111156114855761148461236e565b5b6040519080825280602002602001820160405280156114be57816020015b6114ab611b25565b8152602001906001900390816114a35790505b50925060005b81811015611518578281815181106114df576114de61239d565b5b60200260200101518482815181106114fa576114f961239d565b5b6020026020010181905250808061151090612451565b9150506114c4565b505050919050565b6000600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd3330856040518463ffffffff1660e01b8152600401611581939291906125cc565b6020604051808303816000875af11580156115a0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115c4919061259f565b6115d15760009050611664565b81600660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461161c91906123fb565b600660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550600190505b919050565b60008060008085815260200190815260200160002060030154116116c2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016116b99061264f565b60405180910390fd5b60006040518060c00160405280600354815260200160008087815260200190815260200160002060040154815260200160008087815260200190815260200160002060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020013373ffffffffffffffffffffffffffffffffffffffff168152602001600080878152602001908152602001600020600201548152602001844261178691906123fb565b815250905080600260006003548152602001908152602001600020600082015181600001556020820151816001015560408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506080820151816004015560a082015181600501559050506003600081548092919061186d90612451565b919050555060016003546118819190612499565b91505092915050565b3373ffffffffffffffffffffffffffffffffffffffff1660008086815260200190815260200160002060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461192d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611924906126e1565b60405180910390fd5b82600080868152602001908152602001600020600201819055508060008086815260200190815260200160002060030181905550816000808681526020019081526020016000206004018190555050505050565b6000600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460601b9050919050565b80600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555050565b600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663095ea7b383836040518363ffffffff1660e01b8152600401611a7b92919061254a565b6020604051808303816000875af1158015611a9a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611abe919061259f565b505050565b6040518060c001604052806000815260200160008152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff16815260200160008152602001600081525090565b6040518060a0016040528060008152602001600073ffffffffffffffffffffffffffffffffffffffff1681526020016000815260200160008152602001600081525090565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000611b9a82611b6f565b9050919050565b611baa81611b8f565b8114611bb557600080fd5b50565b600081359050611bc781611ba1565b92915050565b600060208284031215611be357611be2611b6a565b5b6000611bf184828501611bb8565b91505092915050565b6000819050919050565b611c0d81611bfa565b82525050565b6000602082019050611c286000830184611c04565b92915050565b60007fffffffffffffffffffffffffffffffffffffffff00000000000000000000000082169050919050565b611c6381611c2e565b8114611c6e57600080fd5b50565b600081359050611c8081611c5a565b92915050565b600060208284031215611c9c57611c9b611b6a565b5b6000611caa84828501611c71565b91505092915050565b600067ffffffffffffffff82169050919050565b611cd081611cb3565b8114611cdb57600080fd5b50565b600081359050611ced81611cc7565b92915050565b600060208284031215611d0957611d08611b6a565b5b6000611d1784828501611cde565b91505092915050565b611d2981611bfa565b82525050565b611d3881611b8f565b82525050565b60c082016000820151611d546000850182611d20565b506020820151611d676020850182611d20565b506040820151611d7a6040850182611d2f565b506060820151611d8d6060850182611d2f565b506080820151611da06080850182611d20565b5060a0820151611db360a0850182611d20565b50505050565b600060c082019050611dce6000830184611d3e565b92915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b60c082016000820151611e166000850182611d20565b506020820151611e296020850182611d20565b506040820151611e3c6040850182611d2f565b506060820151611e4f6060850182611d2f565b506080820151611e626080850182611d20565b5060a0820151611e7560a0850182611d20565b50505050565b6000611e878383611e00565b60c08301905092915050565b6000602082019050919050565b6000611eab82611dd4565b611eb58185611ddf565b9350611ec083611df0565b8060005b83811015611ef1578151611ed88882611e7b565b9750611ee383611e93565b925050600181019050611ec4565b5085935050505092915050565b60006020820190508181036000830152611f188184611ea0565b905092915050565b611f2981611bfa565b8114611f3457600080fd5b50565b600081359050611f4681611f20565b92915050565b600080600060608486031215611f6557611f64611b6a565b5b6000611f7386828701611f37565b9350506020611f8486828701611f37565b9250506040611f9586828701611f37565b9150509250925092565b600060208284031215611fb557611fb4611b6a565b5b6000611fc384828501611f37565b91505092915050565b60008115159050919050565b611fe181611fcc565b82525050565b6000602082019050611ffc6000830184611fd8565b92915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b60a0820160008201516120446000850182611d20565b5060208201516120576020850182611d2f565b50604082015161206a6040850182611d20565b50606082015161207d6060850182611d20565b5060808201516120906080850182611d20565b50505050565b60006120a2838361202e565b60a08301905092915050565b6000602082019050919050565b60006120c682612002565b6120d0818561200d565b93506120db8361201e565b8060005b8381101561210c5781516120f38882612096565b97506120fe836120ae565b9250506001810190506120df565b5085935050505092915050565b6000602082019050818103600083015261213381846120bb565b905092915050565b6000819050919050565b61214e8161213b565b82525050565b60006020820190506121696000830184612145565b92915050565b6000806040838503121561218657612185611b6a565b5b600061219485828601611f37565b92505060206121a585828601611f37565b9150509250929050565b600080600080608085870312156121c9576121c8611b6a565b5b60006121d787828801611f37565b94505060206121e887828801611f37565b93505060406121f987828801611f37565b925050606061220a87828801611f37565b91505092959194509250565b61221f81611c2e565b82525050565b600060208201905061223a6000830184612216565b92915050565b6122498161213b565b811461225457600080fd5b50565b60008135905061226681612240565b92915050565b60006020828403121561228257612281611b6a565b5b600061229084828501612257565b91505092915050565b600080604083850312156122b0576122af611b6a565b5b60006122be85828601611bb8565b92505060206122cf85828601611f37565b9150509250929050565b6122e281611b8f565b82525050565b60006040820190506122fd60008301856122d9565b61230a60208301846122d9565b9392505050565b60008151905061232081611f20565b92915050565b60006020828403121561233c5761233b611b6a565b5b600061234a84828501612311565b91505092915050565b600060208201905061236860008301846122d9565b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061240682611bfa565b915061241183611bfa565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115612446576124456123cc565b5b828201905092915050565b600061245c82611bfa565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361248e5761248d6123cc565b5b600182019050919050565b60006124a482611bfa565b91506124af83611bfa565b9250828210156124c2576124c16123cc565b5b828203905092915050565b600082825260208201905092915050565b7f696e73756666696369656e742066756e647320746f2077697468647261770000600082015250565b6000612514601e836124cd565b915061251f826124de565b602082019050919050565b6000602082019050818103600083015261254381612507565b9050919050565b600060408201905061255f60008301856122d9565b61256c6020830184611c04565b9392505050565b61257c81611fcc565b811461258757600080fd5b50565b60008151905061259981612573565b92915050565b6000602082840312156125b5576125b4611b6a565b5b60006125c38482850161258a565b91505092915050565b60006060820190506125e160008301866122d9565b6125ee60208301856122d9565b6125fb6040830184611c04565b949350505050565b7f4e6f206d616368696e657320617661696c61626c650000000000000000000000600082015250565b60006126396015836124cd565b915061264482612603565b602082019050919050565b600060208201905081810360008301526126688161262c565b9050919050565b7f4f6e6c7920746865206f776e65722063616e2075706461746520616e206f666660008201527f6572000000000000000000000000000000000000000000000000000000000000602082015250565b60006126cb6022836124cd565b91506126d68261266f565b604082019050919050565b600060208201905081810360008301526126fa816126be565b905091905056fea2646970667358221220ef6b310ec41cbd053317d15c0064333696d52d54c7eaed20e8b0480cae60e6d564736f6c634300080f0033",
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

// TestApprove is a paid mutator transaction binding the contract method 0xfb44e83a.
//
// Solidity: function testApprove(address to, uint256 amount) returns()
func (_Broker *BrokerTransactor) TestApprove(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Broker.contract.Transact(opts, "testApprove", to, amount)
}

// TestApprove is a paid mutator transaction binding the contract method 0xfb44e83a.
//
// Solidity: function testApprove(address to, uint256 amount) returns()
func (_Broker *BrokerSession) TestApprove(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Broker.Contract.TestApprove(&_Broker.TransactOpts, to, amount)
}

// TestApprove is a paid mutator transaction binding the contract method 0xfb44e83a.
//
// Solidity: function testApprove(address to, uint256 amount) returns()
func (_Broker *BrokerTransactorSession) TestApprove(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Broker.Contract.TestApprove(&_Broker.TransactOpts, to, amount)
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

// UserAllowance is a paid mutator transaction binding the contract method 0x0c6b2cbf.
//
// Solidity: function userAllowance(address to) returns(uint256)
func (_Broker *BrokerTransactor) UserAllowance(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _Broker.contract.Transact(opts, "userAllowance", to)
}

// UserAllowance is a paid mutator transaction binding the contract method 0x0c6b2cbf.
//
// Solidity: function userAllowance(address to) returns(uint256)
func (_Broker *BrokerSession) UserAllowance(to common.Address) (*types.Transaction, error) {
	return _Broker.Contract.UserAllowance(&_Broker.TransactOpts, to)
}

// UserAllowance is a paid mutator transaction binding the contract method 0x0c6b2cbf.
//
// Solidity: function userAllowance(address to) returns(uint256)
func (_Broker *BrokerTransactorSession) UserAllowance(to common.Address) (*types.Transaction, error) {
	return _Broker.Contract.UserAllowance(&_Broker.TransactOpts, to)
}

// UserBalance is a paid mutator transaction binding the contract method 0xbf152765.
//
// Solidity: function userBalance() returns(uint256)
func (_Broker *BrokerTransactor) UserBalance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Broker.contract.Transact(opts, "userBalance")
}

// UserBalance is a paid mutator transaction binding the contract method 0xbf152765.
//
// Solidity: function userBalance() returns(uint256)
func (_Broker *BrokerSession) UserBalance() (*types.Transaction, error) {
	return _Broker.Contract.UserBalance(&_Broker.TransactOpts)
}

// UserBalance is a paid mutator transaction binding the contract method 0xbf152765.
//
// Solidity: function userBalance() returns(uint256)
func (_Broker *BrokerTransactorSession) UserBalance() (*types.Transaction, error) {
	return _Broker.Contract.UserBalance(&_Broker.TransactOpts)
}

// UserDeposit is a paid mutator transaction binding the contract method 0x48146113.
//
// Solidity: function userDeposit() returns(uint256)
func (_Broker *BrokerTransactor) UserDeposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Broker.contract.Transact(opts, "userDeposit")
}

// UserDeposit is a paid mutator transaction binding the contract method 0x48146113.
//
// Solidity: function userDeposit() returns(uint256)
func (_Broker *BrokerSession) UserDeposit() (*types.Transaction, error) {
	return _Broker.Contract.UserDeposit(&_Broker.TransactOpts)
}

// UserDeposit is a paid mutator transaction binding the contract method 0x48146113.
//
// Solidity: function userDeposit() returns(uint256)
func (_Broker *BrokerTransactorSession) UserDeposit() (*types.Transaction, error) {
	return _Broker.Contract.UserDeposit(&_Broker.TransactOpts)
}

// UserTokenBalance is a paid mutator transaction binding the contract method 0x2d0d6c33.
//
// Solidity: function userTokenBalance() returns(uint256)
func (_Broker *BrokerTransactor) UserTokenBalance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Broker.contract.Transact(opts, "userTokenBalance")
}

// UserTokenBalance is a paid mutator transaction binding the contract method 0x2d0d6c33.
//
// Solidity: function userTokenBalance() returns(uint256)
func (_Broker *BrokerSession) UserTokenBalance() (*types.Transaction, error) {
	return _Broker.Contract.UserTokenBalance(&_Broker.TransactOpts)
}

// UserTokenBalance is a paid mutator transaction binding the contract method 0x2d0d6c33.
//
// Solidity: function userTokenBalance() returns(uint256)
func (_Broker *BrokerTransactorSession) UserTokenBalance() (*types.Transaction, error) {
	return _Broker.Contract.UserTokenBalance(&_Broker.TransactOpts)
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
