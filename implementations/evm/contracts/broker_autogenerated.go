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
	ABI: "[{\"inputs\":[],\"name\":\"GetTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes20\",\"name\":\"_signature\",\"type\":\"bytes20\"}],\"name\":\"SetMtlsHash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pricePerSecond\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vmTypeId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"machinesAvailable\",\"type\":\"uint256\"}],\"name\":\"addOffer\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offerIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"secs\",\"type\":\"uint256\"}],\"name\":\"bookVM\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"numTokens\",\"type\":\"uint256\"}],\"name\":\"depositCoin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_miner\",\"type\":\"address\"}],\"name\":\"findBookingsByMiner\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vmTypeId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pricePerSecond\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bookedTill\",\"type\":\"uint256\"}],\"internalType\":\"structBroker.Booking[]\",\"name\":\"filteredBookings\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"findBookingsByUser\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vmTypeId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pricePerSecond\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bookedTill\",\"type\":\"uint256\"}],\"internalType\":\"structBroker.Booking[]\",\"name\":\"filteredBookings\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_vmTypeId\",\"type\":\"uint256\"}],\"name\":\"getAvailableOffers\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pricePerSecond\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"machinesAvailable\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vmTypeId\",\"type\":\"uint256\"}],\"internalType\":\"structBroker.VMOffer[]\",\"name\":\"filteredOffers\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"}],\"name\":\"getBooking\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vmTypeId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pricePerSecond\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bookedTill\",\"type\":\"uint256\"}],\"internalType\":\"structBroker.Booking\",\"name\":\"booking\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCoinDecimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getMinerUrl\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"}],\"name\":\"getMinersOffers\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pricePerSecond\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"machinesAvailable\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vmTypeId\",\"type\":\"uint256\"}],\"internalType\":\"structBroker.VMOffer[]\",\"name\":\"filteredOffers\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getMtlsHash\",\"outputs\":[{\"internalType\":\"bytes20\",\"name\":\"\",\"type\":\"bytes20\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStablecoinAddress\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"url\",\"type\":\"bytes32\"}],\"name\":\"setMunerUrl\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"t\",\"type\":\"address\"}],\"name\":\"setStablecoinAddress\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offerIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pricePerSecond\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vmTypeId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"machinesAvailable\",\"type\":\"uint256\"}],\"name\":\"updateOffer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"userAllowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"userBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"userTokenBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"numTokens\",\"type\":\"uint256\"}],\"name\":\"withdrawCoin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040526000600155600060035534801561001a57600080fd5b506128b78061002a6000396000f3fe608060405234801561001057600080fd5b50600436106101425760003560e01c80637ee080c6116100b8578063ca3b3b4b1161007c578063ca3b3b4b146103c7578063cbb6bfbd146103f7578063cf696d1414610427578063e098801314610445578063ea5f398414610461578063f44f11cd1461049157610142565b80637ee080c6146102e95780639ecd050014610319578063ab49ec4514610349578063bf15276514610379578063c700ff0a1461039757610142565b8063301bcaae1161010a578063301bcaae146101ed5780633b3e328b1461021d5780634eb264e81461024d57806354fd4d501461027d5780635a610b571461029b5780637bef4f54146102cb57610142565b8063117298e41461014757806318c5e5021461016357806320255c7e146101815780632d0d6c331461019f5780632dc8a2b9146101bd575b600080fd5b610161600480360381019061015c9190611c12565b6104ad565b005b61016b61051b565b6040516101789190611c58565b60405180910390f35b610189610523565b6040516101969190611c58565b60405180910390f35b6101a76105c8565b6040516101b49190611c58565b60405180910390f35b6101d760048036038101906101d29190611cb3565b61066b565b6040516101e49190611dab565b60405180910390f35b61020760048036038101906102029190611df2565b610776565b6040516102149190611f49565b60405180910390f35b61023760048036038101906102329190611f97565b610a36565b6040516102449190611c58565b60405180910390f35b61026760048036038101906102629190611fea565b610b2c565b6040516102749190612032565b60405180910390f35b610285610cbd565b6040516102929190611c58565b60405180910390f35b6102b560048036038101906102b09190611df2565b610cc2565b6040516102c29190611f49565b60405180910390f35b6102d3610f82565b6040516102e091906120ac565b60405180910390f35b61030360048036038101906102fe9190611fea565b610fac565b60405161031091906121de565b60405180910390f35b610333600480360381019061032e919061223e565b6111ea565b6040516103409190612032565b60405180910390f35b610363600480360381019061035e9190611df2565b611236565b6040516103709190612284565b60405180910390f35b61038161127f565b60405161038e9190611c58565b60405180910390f35b6103b160048036038101906103ac9190611df2565b611310565b6040516103be91906121de565b60405180910390f35b6103e160048036038101906103dc9190611fea565b611578565b6040516103ee9190612032565b60405180910390f35b610411600480360381019061040c919061229f565b6116c1565b60405161041e9190611c58565b60405180910390f35b61042f6118e2565b60405161043c91906122fb565b60405180910390f35b61045f600480360381019061045a9190612316565b61197a565b005b61047b60048036038101906104769190611df2565b611a71565b604051610488919061238c565b60405180910390f35b6104ab60048036038101906104a691906123d3565b611ac7565b005b80600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908360601c021790555050565b600042905090565b6000600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663dd62ed3e33306040518363ffffffff1660e01b815260040161058292919061240f565b602060405180830381865afa15801561059f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105c3919061244d565b905090565b6000600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231336040518263ffffffff1660e01b8152600401610625919061247a565b602060405180830381865afa158015610642573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610666919061244d565b905090565b610673611b0e565b600260008367ffffffffffffffff1681526020019081526020016000206040518060c001604052908160008201548152602001600182015481526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016003820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600482015481526020016005820154815250509050919050565b6060600060035467ffffffffffffffff81111561079657610795612495565b5b6040519080825280602002602001820160405280156107cf57816020015b6107bc611b0e565b8152602001906001900390816107b45790505b509050600080600090505b600354811015610980578473ffffffffffffffffffffffffffffffffffffffff166002600083815260200190815260200160002060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff160361096d57600260008281526020019081526020016000206040518060c001604052908160008201548152602001600182015481526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016003820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160048201548152602001600582015481525050838381518110610952576109516124c4565b5b602002602001018190525060018261096a9190612522565b91505b808061097890612578565b9150506107da565b508067ffffffffffffffff81111561099b5761099a612495565b5b6040519080825280602002602001820160405280156109d457816020015b6109c1611b0e565b8152602001906001900390816109b95790505b50925060005b81811015610a2e578281815181106109f5576109f46124c4565b5b6020026020010151848281518110610a1057610a0f6124c4565b5b60200260200101819052508080610a2690612578565b9150506109da565b505050919050565b60006040518060a0016040528060015481526020013373ffffffffffffffffffffffffffffffffffffffff1681526020018581526020018381526020018481525060008060015481526020019081526020016000206000820151816000015560208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160020155606082015181600301556080820151816004015590505060016000815480929190610b1090612578565b919050555060018054610b2391906125c0565b90509392505050565b600081610b3761127f565b1015610b78576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b6f90612651565b60405180910390fd5b600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb33846040518363ffffffff1660e01b8152600401610bd5929190612671565b6020604051808303816000875af1158015610bf4573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c1891906126c6565b610c255760009050610cb8565b81600660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054610c7091906125c0565b600660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550600190505b919050565b600181565b6060600060035467ffffffffffffffff811115610ce257610ce1612495565b5b604051908082528060200260200182016040528015610d1b57816020015b610d08611b0e565b815260200190600190039081610d005790505b509050600080600090505b600354811015610ecc578473ffffffffffffffffffffffffffffffffffffffff166002600083815260200190815260200160002060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603610eb957600260008281526020019081526020016000206040518060c001604052908160008201548152602001600182015481526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016003820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160048201548152602001600582015481525050838381518110610e9e57610e9d6124c4565b5b6020026020010181905250600182610eb69190612522565b91505b8080610ec490612578565b915050610d26565b508067ffffffffffffffff811115610ee757610ee6612495565b5b604051908082528060200260200182016040528015610f2057816020015b610f0d611b0e565b815260200190600190039081610f055790505b50925060005b81811015610f7a57828181518110610f4157610f406124c4565b5b6020026020010151848281518110610f5c57610f5b6124c4565b5b60200260200101819052508080610f7290612578565b915050610f26565b505050919050565b6000600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6060600060015467ffffffffffffffff811115610fcc57610fcb612495565b5b60405190808252806020026020018201604052801561100557816020015b610ff2611b70565b815260200190600190039081610fea5790505b509050600080600090505b6001548110156111345784600080838152602001908152602001600020600401541480156110535750600080600083815260200190815260200160002060030154115b15611121576000808281526020019081526020016000206040518060a0016040529081600082015481526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820154815260200160038201548152602001600482015481525050838381518110611106576111056124c4565b5b602002602001018190525060018261111e9190612522565b91505b808061112c90612578565b915050611010565b508067ffffffffffffffff81111561114f5761114e612495565b5b60405190808252806020026020018201604052801561118857816020015b611175611b70565b81526020019060019003908161116d5790505b50925060005b818110156111e2578281815181106111a9576111a86124c4565b5b60200260200101518482815181106111c4576111c36124c4565b5b602002602001018190525080806111da90612578565b91505061118e565b505050919050565b600081600860006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060009050919050565b6000600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b6000600760003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054600660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461130b91906125c0565b905090565b6060600060015467ffffffffffffffff8111156113305761132f612495565b5b60405190808252806020026020018201604052801561136957816020015b611356611b70565b81526020019060019003908161134e5790505b509050600080600090505b6001548110156114c2578473ffffffffffffffffffffffffffffffffffffffff1660008083815260200190815260200160002060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16036114af576000808281526020019081526020016000206040518060a0016040529081600082015481526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820154815260200160038201548152602001600482015481525050838381518110611494576114936124c4565b5b60200260200101819052506001826114ac9190612522565b91505b80806114ba90612578565b915050611374565b508067ffffffffffffffff8111156114dd576114dc612495565b5b60405190808252806020026020018201604052801561151657816020015b611503611b70565b8152602001906001900390816114fb5790505b50925060005b8181101561157057828181518110611537576115366124c4565b5b6020026020010151848281518110611552576115516124c4565b5b6020026020010181905250808061156890612578565b91505061151c565b505050919050565b6000600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd3330856040518463ffffffff1660e01b81526004016115d9939291906126f3565b6020604051808303816000875af11580156115f8573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061161c91906126c6565b61162957600090506116bc565b81600660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546116749190612522565b600660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550600190505b919050565b600080600080858152602001908152602001600020600301541161171a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161171190612776565b60405180910390fd5b60006040518060c00160405280600354815260200160008087815260200190815260200160002060040154815260200160008087815260200190815260200160002060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020013373ffffffffffffffffffffffffffffffffffffffff16815260200160008087815260200190815260200160002060020154815260200184426117de9190612522565b815250905080600260006003548152602001908152602001600020600082015181600001556020820151816001015560408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506080820151816004015560a08201518160050155905050600360008154809291906118c590612578565b919050555060016003546118d991906125c0565b91505092915050565b6000600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663313ce5676040518163ffffffff1660e01b8152600401602060405180830381865afa158015611951573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061197591906127c2565b905090565b3373ffffffffffffffffffffffffffffffffffffffff1660008086815260200190815260200160002060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614611a1d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611a1490612861565b60405180910390fd5b82600080868152602001908152602001600020600201819055508060008086815260200190815260200160002060030181905550816000808681526020019081526020016000206004018190555050505050565b6000600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460601b9050919050565b80600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555050565b6040518060c001604052806000815260200160008152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff16815260200160008152602001600081525090565b6040518060a0016040528060008152602001600073ffffffffffffffffffffffffffffffffffffffff1681526020016000815260200160008152602001600081525090565b600080fd5b60007fffffffffffffffffffffffffffffffffffffffff00000000000000000000000082169050919050565b611bef81611bba565b8114611bfa57600080fd5b50565b600081359050611c0c81611be6565b92915050565b600060208284031215611c2857611c27611bb5565b5b6000611c3684828501611bfd565b91505092915050565b6000819050919050565b611c5281611c3f565b82525050565b6000602082019050611c6d6000830184611c49565b92915050565b600067ffffffffffffffff82169050919050565b611c9081611c73565b8114611c9b57600080fd5b50565b600081359050611cad81611c87565b92915050565b600060208284031215611cc957611cc8611bb5565b5b6000611cd784828501611c9e565b91505092915050565b611ce981611c3f565b82525050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000611d1a82611cef565b9050919050565b611d2a81611d0f565b82525050565b60c082016000820151611d466000850182611ce0565b506020820151611d596020850182611ce0565b506040820151611d6c6040850182611d21565b506060820151611d7f6060850182611d21565b506080820151611d926080850182611ce0565b5060a0820151611da560a0850182611ce0565b50505050565b600060c082019050611dc06000830184611d30565b92915050565b611dcf81611d0f565b8114611dda57600080fd5b50565b600081359050611dec81611dc6565b92915050565b600060208284031215611e0857611e07611bb5565b5b6000611e1684828501611ddd565b91505092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b60c082016000820151611e616000850182611ce0565b506020820151611e746020850182611ce0565b506040820151611e876040850182611d21565b506060820151611e9a6060850182611d21565b506080820151611ead6080850182611ce0565b5060a0820151611ec060a0850182611ce0565b50505050565b6000611ed28383611e4b565b60c08301905092915050565b6000602082019050919050565b6000611ef682611e1f565b611f008185611e2a565b9350611f0b83611e3b565b8060005b83811015611f3c578151611f238882611ec6565b9750611f2e83611ede565b925050600181019050611f0f565b5085935050505092915050565b60006020820190508181036000830152611f638184611eeb565b905092915050565b611f7481611c3f565b8114611f7f57600080fd5b50565b600081359050611f9181611f6b565b92915050565b600080600060608486031215611fb057611faf611bb5565b5b6000611fbe86828701611f82565b9350506020611fcf86828701611f82565b9250506040611fe086828701611f82565b9150509250925092565b60006020828403121561200057611fff611bb5565b5b600061200e84828501611f82565b91505092915050565b60008115159050919050565b61202c81612017565b82525050565b60006020820190506120476000830184612023565b92915050565b6000819050919050565b600061207261206d61206884611cef565b61204d565b611cef565b9050919050565b600061208482612057565b9050919050565b600061209682612079565b9050919050565b6120a68161208b565b82525050565b60006020820190506120c1600083018461209d565b92915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b60a0820160008201516121096000850182611ce0565b50602082015161211c6020850182611d21565b50604082015161212f6040850182611ce0565b5060608201516121426060850182611ce0565b5060808201516121556080850182611ce0565b50505050565b600061216783836120f3565b60a08301905092915050565b6000602082019050919050565b600061218b826120c7565b61219581856120d2565b93506121a0836120e3565b8060005b838110156121d15781516121b8888261215b565b97506121c383612173565b9250506001810190506121a4565b5085935050505092915050565b600060208201905081810360008301526121f88184612180565b905092915050565b600061220b82611d0f565b9050919050565b61221b81612200565b811461222657600080fd5b50565b60008135905061223881612212565b92915050565b60006020828403121561225457612253611bb5565b5b600061226284828501612229565b91505092915050565b6000819050919050565b61227e8161226b565b82525050565b60006020820190506122996000830184612275565b92915050565b600080604083850312156122b6576122b5611bb5565b5b60006122c485828601611f82565b92505060206122d585828601611f82565b9150509250929050565b600060ff82169050919050565b6122f5816122df565b82525050565b600060208201905061231060008301846122ec565b92915050565b600080600080608085870312156123305761232f611bb5565b5b600061233e87828801611f82565b945050602061234f87828801611f82565b935050604061236087828801611f82565b925050606061237187828801611f82565b91505092959194509250565b61238681611bba565b82525050565b60006020820190506123a1600083018461237d565b92915050565b6123b08161226b565b81146123bb57600080fd5b50565b6000813590506123cd816123a7565b92915050565b6000602082840312156123e9576123e8611bb5565b5b60006123f7848285016123be565b91505092915050565b61240981611d0f565b82525050565b60006040820190506124246000830185612400565b6124316020830184612400565b9392505050565b60008151905061244781611f6b565b92915050565b60006020828403121561246357612462611bb5565b5b600061247184828501612438565b91505092915050565b600060208201905061248f6000830184612400565b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061252d82611c3f565b915061253883611c3f565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0382111561256d5761256c6124f3565b5b828201905092915050565b600061258382611c3f565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036125b5576125b46124f3565b5b600182019050919050565b60006125cb82611c3f565b91506125d683611c3f565b9250828210156125e9576125e86124f3565b5b828203905092915050565b600082825260208201905092915050565b7f696e73756666696369656e742066756e647320746f2077697468647261770000600082015250565b600061263b601e836125f4565b915061264682612605565b602082019050919050565b6000602082019050818103600083015261266a8161262e565b9050919050565b60006040820190506126866000830185612400565b6126936020830184611c49565b9392505050565b6126a381612017565b81146126ae57600080fd5b50565b6000815190506126c08161269a565b92915050565b6000602082840312156126dc576126db611bb5565b5b60006126ea848285016126b1565b91505092915050565b60006060820190506127086000830186612400565b6127156020830185612400565b6127226040830184611c49565b949350505050565b7f4e6f206d616368696e657320617661696c61626c650000000000000000000000600082015250565b60006127606015836125f4565b915061276b8261272a565b602082019050919050565b6000602082019050818103600083015261278f81612753565b9050919050565b61279f816122df565b81146127aa57600080fd5b50565b6000815190506127bc81612796565b92915050565b6000602082840312156127d8576127d7611bb5565b5b60006127e6848285016127ad565b91505092915050565b7f4f6e6c7920746865206f776e65722063616e2075706461746520616e206f666660008201527f6572000000000000000000000000000000000000000000000000000000000000602082015250565b600061284b6022836125f4565b9150612856826127ef565b604082019050919050565b6000602082019050818103600083015261287a8161283e565b905091905056fea264697066735822122013395acb4d424923b55506ee0f6b677d86c2d86d286d011147dcc16389be920664736f6c634300080f0033",
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
