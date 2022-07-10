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
	ABI: "[{\"inputs\":[],\"name\":\"GetTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes20\",\"name\":\"_signature\",\"type\":\"bytes20\"}],\"name\":\"SetMtlsHash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pricePerSecond\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vmTypeId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"machinesAvailable\",\"type\":\"uint256\"}],\"name\":\"addOffer\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offerIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"secs\",\"type\":\"uint256\"}],\"name\":\"bookVM\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"communityFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"communityWallet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_miner\",\"type\":\"address\"}],\"name\":\"findBookingsByMiner\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vmTypeId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pricePerSecond\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bookedTill\",\"type\":\"uint256\"}],\"internalType\":\"structBroker.Booking[]\",\"name\":\"filteredBookings\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"findBookingsByUser\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vmTypeId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pricePerSecond\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bookedTill\",\"type\":\"uint256\"}],\"internalType\":\"structBroker.Booking[]\",\"name\":\"filteredBookings\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_vmTypeId\",\"type\":\"uint256\"}],\"name\":\"getAvailableOffers\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pricePerSecond\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"machinesAvailable\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vmTypeId\",\"type\":\"uint256\"}],\"internalType\":\"structBroker.VMOffer[]\",\"name\":\"filteredOffers\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"}],\"name\":\"getBooking\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vmTypeId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pricePerSecond\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bookedTill\",\"type\":\"uint256\"}],\"internalType\":\"structBroker.Booking\",\"name\":\"booking\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCommunityFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCommunityWallet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getMinerUrl\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"}],\"name\":\"getMinersOffers\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pricePerSecond\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"machinesAvailable\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vmTypeId\",\"type\":\"uint256\"}],\"internalType\":\"structBroker.VMOffer[]\",\"name\":\"filteredOffers\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getMtlsHash\",\"outputs\":[{\"internalType\":\"bytes20\",\"name\":\"\",\"type\":\"bytes20\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"setCommunityFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"}],\"name\":\"setCommunityWallet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"url\",\"type\":\"bytes32\"}],\"name\":\"setMunerUrl\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offerIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pricePerSecond\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vmTypeId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"machinesAvailable\",\"type\":\"uint256\"}],\"name\":\"updateOffer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040526101f46000556000600355600060055534801561002057600080fd5b50612286806100306000396000f3fe608060405234801561001057600080fd5b506004361061012c5760003560e01c806385d6bb81116100ad578063c757483911610071578063c75748391461036b578063cbb6bfbd14610389578063e0988013146103b9578063ea5f3984146103d5578063f44f11cd146104055761012c565b806385d6bb81146102b35780638961be6b146102cf57806390ad688b146102ed578063ab49ec451461030b578063c700ff0a1461033b5761012c565b8063301bcaae116100f4578063301bcaae146101d55780633b3e328b1461020557806354fd4d50146102355780635a610b57146102535780637ee080c6146102835761012c565b80630d12bbdb14610131578063117298e41461014d57806318c5e5021461016957806326e82d71146101875780632dc8a2b9146101a5575b600080fd5b61014b600480360381019061014691906117d6565b610421565b005b6101676004803603810190610162919061185b565b6104bb565b005b610171610529565b60405161017e9190611897565b60405180910390f35b61018f610531565b60405161019c91906118f3565b60405180910390f35b6101bf60048036038101906101ba919061194e565b61055b565b6040516101cc9190611a14565b60405180910390f35b6101ef60048036038101906101ea9190611a5b565b610666565b6040516101fc9190611bb2565b60405180910390f35b61021f600480360381019061021a9190611bd4565b610926565b60405161022c9190611897565b60405180910390f35b61023d610a1e565b60405161024a9190611897565b60405180910390f35b61026d60048036038101906102689190611a5b565b610a23565b60405161027a9190611bb2565b60405180910390f35b61029d600480360381019061029891906117d6565b610ce3565b6040516102aa9190611d3e565b60405180910390f35b6102cd60048036038101906102c89190611a5b565b610f24565b005b6102d761104f565b6040516102e49190611897565b60405180910390f35b6102f5611055565b6040516103029190611897565b60405180910390f35b61032560048036038101906103209190611a5b565b61105e565b6040516103329190611d79565b60405180910390f35b61035560048036038101906103509190611a5b565b6110a7565b6040516103629190611d3e565b60405180910390f35b610373611311565b60405161038091906118f3565b60405180910390f35b6103a3600480360381019061039e9190611d94565b611337565b6040516103b09190611897565b60405180910390f35b6103d360048036038101906103ce9190611dd4565b61155c565b005b6103ef60048036038101906103ea9190611a5b565b611657565b6040516103fc9190611e4a565b60405180910390f35b61041f600480360381019061041a9190611e91565b6116ad565b005b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146104b1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104a890611f41565b60405180910390fd5b8060008190555050565b80600760003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908360601c021790555050565b600042905090565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6105636116f4565b600460008367ffffffffffffffff1681526020019081526020016000206040518060c001604052908160008201548152602001600182015481526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016003820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600482015481526020016005820154815250509050919050565b6060600060055467ffffffffffffffff81111561068657610685611f61565b5b6040519080825280602002602001820160405280156106bf57816020015b6106ac6116f4565b8152602001906001900390816106a45790505b509050600080600090505b600554811015610870578473ffffffffffffffffffffffffffffffffffffffff166004600083815260200190815260200160002060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff160361085d57600460008281526020019081526020016000206040518060c001604052908160008201548152602001600182015481526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016003820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016004820154815260200160058201548152505083838151811061084257610841611f90565b5b602002602001018190525060018261085a9190611fee565b91505b808061086890612044565b9150506106ca565b508067ffffffffffffffff81111561088b5761088a611f61565b5b6040519080825280602002602001820160405280156108c457816020015b6108b16116f4565b8152602001906001900390816108a95790505b50925060005b8181101561091e578281815181106108e5576108e4611f90565b5b6020026020010151848281518110610900576108ff611f90565b5b6020026020010181905250808061091690612044565b9150506108ca565b505050919050565b60006040518060a0016040528060035481526020013373ffffffffffffffffffffffffffffffffffffffff168152602001858152602001838152602001848152506002600060035481526020019081526020016000206000820151816000015560208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160020155606082015181600301556080820151816004015590505060036000815480929190610a0190612044565b91905055506001600354610a15919061208c565b90509392505050565b600181565b6060600060055467ffffffffffffffff811115610a4357610a42611f61565b5b604051908082528060200260200182016040528015610a7c57816020015b610a696116f4565b815260200190600190039081610a615790505b509050600080600090505b600554811015610c2d578473ffffffffffffffffffffffffffffffffffffffff166004600083815260200190815260200160002060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603610c1a57600460008281526020019081526020016000206040518060c001604052908160008201548152602001600182015481526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016003820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160048201548152602001600582015481525050838381518110610bff57610bfe611f90565b5b6020026020010181905250600182610c179190611fee565b91505b8080610c2590612044565b915050610a87565b508067ffffffffffffffff811115610c4857610c47611f61565b5b604051908082528060200260200182016040528015610c8157816020015b610c6e6116f4565b815260200190600190039081610c665790505b50925060005b81811015610cdb57828181518110610ca257610ca1611f90565b5b6020026020010151848281518110610cbd57610cbc611f90565b5b60200260200101819052508080610cd390612044565b915050610c87565b505050919050565b6060600060035467ffffffffffffffff811115610d0357610d02611f61565b5b604051908082528060200260200182016040528015610d3c57816020015b610d29611756565b815260200190600190039081610d215790505b509050600080600090505b600354811015610e6e57846002600083815260200190815260200160002060040154148015610d8c575060006002600083815260200190815260200160002060030154115b15610e5b57600260008281526020019081526020016000206040518060a0016040529081600082015481526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820154815260200160038201548152602001600482015481525050838381518110610e4057610e3f611f90565b5b6020026020010181905250600182610e589190611fee565b91505b8080610e6690612044565b915050610d47565b508067ffffffffffffffff811115610e8957610e88611f61565b5b604051908082528060200260200182016040528015610ec257816020015b610eaf611756565b815260200190600190039081610ea75790505b50925060005b81811015610f1c57828181518110610ee357610ee2611f90565b5b6020026020010151848281518110610efe57610efd611f90565b5b60200260200101819052508080610f1490612044565b915050610ec8565b505050919050565b600073ffffffffffffffffffffffffffffffffffffffff16600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461100b57600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461100a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161100190612132565b60405180910390fd5b5b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60005481565b60008054905090565b6000600660008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b6060600060035467ffffffffffffffff8111156110c7576110c6611f61565b5b60405190808252806020026020018201604052801561110057816020015b6110ed611756565b8152602001906001900390816110e55790505b509050600080600090505b60035481101561125b578473ffffffffffffffffffffffffffffffffffffffff166002600083815260200190815260200160002060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff160361124857600260008281526020019081526020016000206040518060a0016040529081600082015481526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600282015481526020016003820154815260200160048201548152505083838151811061122d5761122c611f90565b5b60200260200101819052506001826112459190611fee565b91505b808061125390612044565b91505061110b565b508067ffffffffffffffff81111561127657611275611f61565b5b6040519080825280602002602001820160405280156112af57816020015b61129c611756565b8152602001906001900390816112945790505b50925060005b81811015611309578281815181106112d0576112cf611f90565b5b60200260200101518482815181106112eb576112ea611f90565b5b6020026020010181905250808061130190612044565b9150506112b5565b505050919050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600080600260008581526020019081526020016000206003015411611391576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016113889061219e565b60405180910390fd5b60006040518060c001604052806005548152602001600260008781526020019081526020016000206004015481526020016002600087815260200190815260200160002060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020013373ffffffffffffffffffffffffffffffffffffffff1681526020016002600087815260200190815260200160002060020154815260200184426114589190611fee565b815250905080600460006005548152602001908152602001600020600082015181600001556020820151816001015560408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506080820151816004015560a082015181600501559050506005600081548092919061153f90612044565b91905055506001600554611553919061208c565b91505092915050565b3373ffffffffffffffffffffffffffffffffffffffff166002600086815260200190815260200160002060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614611600576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016115f790612230565b60405180910390fd5b82600260008681526020019081526020016000206002018190555080600260008681526020019081526020016000206003018190555081600260008681526020019081526020016000206004018190555050505050565b6000600760008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460601b9050919050565b80600660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555050565b6040518060c001604052806000815260200160008152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff16815260200160008152602001600081525090565b6040518060a0016040528060008152602001600073ffffffffffffffffffffffffffffffffffffffff1681526020016000815260200160008152602001600081525090565b600080fd5b6000819050919050565b6117b3816117a0565b81146117be57600080fd5b50565b6000813590506117d0816117aa565b92915050565b6000602082840312156117ec576117eb61179b565b5b60006117fa848285016117c1565b91505092915050565b60007fffffffffffffffffffffffffffffffffffffffff00000000000000000000000082169050919050565b61183881611803565b811461184357600080fd5b50565b6000813590506118558161182f565b92915050565b6000602082840312156118715761187061179b565b5b600061187f84828501611846565b91505092915050565b611891816117a0565b82525050565b60006020820190506118ac6000830184611888565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006118dd826118b2565b9050919050565b6118ed816118d2565b82525050565b600060208201905061190860008301846118e4565b92915050565b600067ffffffffffffffff82169050919050565b61192b8161190e565b811461193657600080fd5b50565b60008135905061194881611922565b92915050565b6000602082840312156119645761196361179b565b5b600061197284828501611939565b91505092915050565b611984816117a0565b82525050565b611993816118d2565b82525050565b60c0820160008201516119af600085018261197b565b5060208201516119c2602085018261197b565b5060408201516119d5604085018261198a565b5060608201516119e8606085018261198a565b5060808201516119fb608085018261197b565b5060a0820151611a0e60a085018261197b565b50505050565b600060c082019050611a296000830184611999565b92915050565b611a38816118d2565b8114611a4357600080fd5b50565b600081359050611a5581611a2f565b92915050565b600060208284031215611a7157611a7061179b565b5b6000611a7f84828501611a46565b91505092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b60c082016000820151611aca600085018261197b565b506020820151611add602085018261197b565b506040820151611af0604085018261198a565b506060820151611b03606085018261198a565b506080820151611b16608085018261197b565b5060a0820151611b2960a085018261197b565b50505050565b6000611b3b8383611ab4565b60c08301905092915050565b6000602082019050919050565b6000611b5f82611a88565b611b698185611a93565b9350611b7483611aa4565b8060005b83811015611ba5578151611b8c8882611b2f565b9750611b9783611b47565b925050600181019050611b78565b5085935050505092915050565b60006020820190508181036000830152611bcc8184611b54565b905092915050565b600080600060608486031215611bed57611bec61179b565b5b6000611bfb868287016117c1565b9350506020611c0c868287016117c1565b9250506040611c1d868287016117c1565b9150509250925092565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b60a082016000820151611c69600085018261197b565b506020820151611c7c602085018261198a565b506040820151611c8f604085018261197b565b506060820151611ca2606085018261197b565b506080820151611cb5608085018261197b565b50505050565b6000611cc78383611c53565b60a08301905092915050565b6000602082019050919050565b6000611ceb82611c27565b611cf58185611c32565b9350611d0083611c43565b8060005b83811015611d31578151611d188882611cbb565b9750611d2383611cd3565b925050600181019050611d04565b5085935050505092915050565b60006020820190508181036000830152611d588184611ce0565b905092915050565b6000819050919050565b611d7381611d60565b82525050565b6000602082019050611d8e6000830184611d6a565b92915050565b60008060408385031215611dab57611daa61179b565b5b6000611db9858286016117c1565b9250506020611dca858286016117c1565b9150509250929050565b60008060008060808587031215611dee57611ded61179b565b5b6000611dfc878288016117c1565b9450506020611e0d878288016117c1565b9350506040611e1e878288016117c1565b9250506060611e2f878288016117c1565b91505092959194509250565b611e4481611803565b82525050565b6000602082019050611e5f6000830184611e3b565b92915050565b611e6e81611d60565b8114611e7957600080fd5b50565b600081359050611e8b81611e65565b92915050565b600060208284031215611ea757611ea661179b565b5b6000611eb584828501611e7c565b91505092915050565b600082825260208201905092915050565b7f4f6e6c792074686520436f6d6d756e697479206f776e65722063616e2075706460008201527f61746520616e2046656500000000000000000000000000000000000000000000602082015250565b6000611f2b602a83611ebe565b9150611f3682611ecf565b604082019050919050565b60006020820190508181036000830152611f5a81611f1e565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000611ff9826117a0565b9150612004836117a0565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0382111561203957612038611fbf565b5b828201905092915050565b600061204f826117a0565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361208157612080611fbf565b5b600182019050919050565b6000612097826117a0565b91506120a2836117a0565b9250828210156120b5576120b4611fbf565b5b828203905092915050565b7f4f6e6c792074686520436f6d6d756e697479206f776e65722063616e2075706460008201527f61746520616e2077616c6c657400000000000000000000000000000000000000602082015250565b600061211c602d83611ebe565b9150612127826120c0565b604082019050919050565b6000602082019050818103600083015261214b8161210f565b9050919050565b7f4e6f206d616368696e657320617661696c61626c650000000000000000000000600082015250565b6000612188601583611ebe565b915061219382612152565b602082019050919050565b600060208201905081810360008301526121b78161217b565b9050919050565b7f4f6e6c7920746865206f776e65722063616e2075706461746520616e206f666660008201527f6572000000000000000000000000000000000000000000000000000000000000602082015250565b600061221a602283611ebe565b9150612225826121be565b604082019050919050565b600060208201905081810360008301526122498161220d565b905091905056fea2646970667358221220416200f5714ec6d2c19d1f67ac911ddce149f358abe45e96a9e1d84e1c3ac30364736f6c634300080f0033",
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

// CommunityFee is a free data retrieval call binding the contract method 0x8961be6b.
//
// Solidity: function communityFee() view returns(uint256)
func (_Broker *BrokerCaller) CommunityFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Broker.contract.Call(opts, &out, "communityFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CommunityFee is a free data retrieval call binding the contract method 0x8961be6b.
//
// Solidity: function communityFee() view returns(uint256)
func (_Broker *BrokerSession) CommunityFee() (*big.Int, error) {
	return _Broker.Contract.CommunityFee(&_Broker.CallOpts)
}

// CommunityFee is a free data retrieval call binding the contract method 0x8961be6b.
//
// Solidity: function communityFee() view returns(uint256)
func (_Broker *BrokerCallerSession) CommunityFee() (*big.Int, error) {
	return _Broker.Contract.CommunityFee(&_Broker.CallOpts)
}

// CommunityWallet is a free data retrieval call binding the contract method 0xc7574839.
//
// Solidity: function communityWallet() view returns(address)
func (_Broker *BrokerCaller) CommunityWallet(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Broker.contract.Call(opts, &out, "communityWallet")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CommunityWallet is a free data retrieval call binding the contract method 0xc7574839.
//
// Solidity: function communityWallet() view returns(address)
func (_Broker *BrokerSession) CommunityWallet() (common.Address, error) {
	return _Broker.Contract.CommunityWallet(&_Broker.CallOpts)
}

// CommunityWallet is a free data retrieval call binding the contract method 0xc7574839.
//
// Solidity: function communityWallet() view returns(address)
func (_Broker *BrokerCallerSession) CommunityWallet() (common.Address, error) {
	return _Broker.Contract.CommunityWallet(&_Broker.CallOpts)
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

// GetCommunityWallet is a free data retrieval call binding the contract method 0x26e82d71.
//
// Solidity: function getCommunityWallet() view returns(address)
func (_Broker *BrokerCaller) GetCommunityWallet(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Broker.contract.Call(opts, &out, "getCommunityWallet")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetCommunityWallet is a free data retrieval call binding the contract method 0x26e82d71.
//
// Solidity: function getCommunityWallet() view returns(address)
func (_Broker *BrokerSession) GetCommunityWallet() (common.Address, error) {
	return _Broker.Contract.GetCommunityWallet(&_Broker.CallOpts)
}

// GetCommunityWallet is a free data retrieval call binding the contract method 0x26e82d71.
//
// Solidity: function getCommunityWallet() view returns(address)
func (_Broker *BrokerCallerSession) GetCommunityWallet() (common.Address, error) {
	return _Broker.Contract.GetCommunityWallet(&_Broker.CallOpts)
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

// SetCommunityFee is a paid mutator transaction binding the contract method 0x0d12bbdb.
//
// Solidity: function setCommunityFee(uint256 fee) returns()
func (_Broker *BrokerTransactor) SetCommunityFee(opts *bind.TransactOpts, fee *big.Int) (*types.Transaction, error) {
	return _Broker.contract.Transact(opts, "setCommunityFee", fee)
}

// SetCommunityFee is a paid mutator transaction binding the contract method 0x0d12bbdb.
//
// Solidity: function setCommunityFee(uint256 fee) returns()
func (_Broker *BrokerSession) SetCommunityFee(fee *big.Int) (*types.Transaction, error) {
	return _Broker.Contract.SetCommunityFee(&_Broker.TransactOpts, fee)
}

// SetCommunityFee is a paid mutator transaction binding the contract method 0x0d12bbdb.
//
// Solidity: function setCommunityFee(uint256 fee) returns()
func (_Broker *BrokerTransactorSession) SetCommunityFee(fee *big.Int) (*types.Transaction, error) {
	return _Broker.Contract.SetCommunityFee(&_Broker.TransactOpts, fee)
}

// SetCommunityWallet is a paid mutator transaction binding the contract method 0x85d6bb81.
//
// Solidity: function setCommunityWallet(address wallet) returns()
func (_Broker *BrokerTransactor) SetCommunityWallet(opts *bind.TransactOpts, wallet common.Address) (*types.Transaction, error) {
	return _Broker.contract.Transact(opts, "setCommunityWallet", wallet)
}

// SetCommunityWallet is a paid mutator transaction binding the contract method 0x85d6bb81.
//
// Solidity: function setCommunityWallet(address wallet) returns()
func (_Broker *BrokerSession) SetCommunityWallet(wallet common.Address) (*types.Transaction, error) {
	return _Broker.Contract.SetCommunityWallet(&_Broker.TransactOpts, wallet)
}

// SetCommunityWallet is a paid mutator transaction binding the contract method 0x85d6bb81.
//
// Solidity: function setCommunityWallet(address wallet) returns()
func (_Broker *BrokerTransactorSession) SetCommunityWallet(wallet common.Address) (*types.Transaction, error) {
	return _Broker.Contract.SetCommunityWallet(&_Broker.TransactOpts, wallet)
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
