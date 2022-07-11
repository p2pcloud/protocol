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

// FiatTokenMetaData contains all meta data concerning the FiatToken contract.
var FiatTokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"value\",\"type\":\"uint64\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"value\",\"type\":\"uint64\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"numTokens\",\"type\":\"uint64\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenOwner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"numTokens\",\"type\":\"uint64\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"numTokens\",\"type\":\"uint64\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040526001600260006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555034801561003a57600080fd5b50600260009054906101000a900467ffffffffffffffff166000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550610f2c806100c56000396000f3fe608060405234801561001057600080fd5b50600436106100935760003560e01c8063313ce56711610066578063313ce567146101345780635d359fbd1461015257806370a082311461018257806395d89b41146101b2578063dd62ed3e146101d057610093565b806306fdde03146100985780631086a9aa146100b657806318160ddd146100e65780632ea0dfe114610104575b600080fd5b6100a0610200565b6040516100ad9190610bf9565b60405180910390f35b6100d060048036038101906100cb9190610cbe565b610239565b6040516100dd9190610d19565b60405180910390f35b6100ee61034d565b6040516100fb9190610d43565b60405180910390f35b61011e60048036038101906101199190610d5e565b61036b565b60405161012b9190610d19565b60405180910390f35b61013c6107c3565b6040516101499190610dcd565b60405180910390f35b61016c60048036038101906101679190610cbe565b6107c8565b6040516101799190610d19565b60405180910390f35b61019c60048036038101906101979190610de8565b610a30565b6040516101a99190610d43565b60405180910390f35b6101ba610a8c565b6040516101c79190610bf9565b60405180910390f35b6101ea60048036038101906101e59190610e15565b610ac5565b6040516101f79190610d43565b60405180910390f35b6040518060400160405280600e81526020017f55534420737461626c65636f696e00000000000000000000000000000000000081525081565b600081600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f16304dfea7f3fbabcf59225f0629cb307fecb8d5652b069080aa9be2c765d7d28460405161033b9190610d43565b60405180910390a36001905092915050565b6000600260009054906101000a900467ffffffffffffffff16905090565b60008060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900467ffffffffffffffff1667ffffffffffffffff168267ffffffffffffffff1611156103e057600080fd5b600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900467ffffffffffffffff1667ffffffffffffffff168267ffffffffffffffff16111561049157600080fd5b816000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900467ffffffffffffffff166104ef9190610e84565b6000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555081600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900467ffffffffffffffff166105ef9190610e84565b600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900467ffffffffffffffff166106ef9190610eb8565b6000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f831ac82b07fb396dafef0077cea6e002235d88e63f35cbd5df2c065107f1e74a846040516107b09190610d43565b60405180910390a3600190509392505050565b600681565b60008060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900467ffffffffffffffff1667ffffffffffffffff168267ffffffffffffffff16111561083d57600080fd5b816000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900467ffffffffffffffff1661089b9190610e84565b6000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900467ffffffffffffffff1661095d9190610eb8565b6000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f831ac82b07fb396dafef0077cea6e002235d88e63f35cbd5df2c065107f1e74a84604051610a1e9190610d43565b60405180910390a36001905092915050565b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900467ffffffffffffffff169050919050565b6040518060400160405280600481526020017f555344430000000000000000000000000000000000000000000000000000000081525081565b6000600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900467ffffffffffffffff16905092915050565b600081519050919050565b600082825260208201905092915050565b60005b83811015610b9a578082015181840152602081019050610b7f565b83811115610ba9576000848401525b50505050565b6000601f19601f8301169050919050565b6000610bcb82610b60565b610bd58185610b6b565b9350610be5818560208601610b7c565b610bee81610baf565b840191505092915050565b60006020820190508181036000830152610c138184610bc0565b905092915050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610c4b82610c20565b9050919050565b610c5b81610c40565b8114610c6657600080fd5b50565b600081359050610c7881610c52565b92915050565b600067ffffffffffffffff82169050919050565b610c9b81610c7e565b8114610ca657600080fd5b50565b600081359050610cb881610c92565b92915050565b60008060408385031215610cd557610cd4610c1b565b5b6000610ce385828601610c69565b9250506020610cf485828601610ca9565b9150509250929050565b60008115159050919050565b610d1381610cfe565b82525050565b6000602082019050610d2e6000830184610d0a565b92915050565b610d3d81610c7e565b82525050565b6000602082019050610d586000830184610d34565b92915050565b600080600060608486031215610d7757610d76610c1b565b5b6000610d8586828701610c69565b9350506020610d9686828701610c69565b9250506040610da786828701610ca9565b9150509250925092565b600060ff82169050919050565b610dc781610db1565b82525050565b6000602082019050610de26000830184610dbe565b92915050565b600060208284031215610dfe57610dfd610c1b565b5b6000610e0c84828501610c69565b91505092915050565b60008060408385031215610e2c57610e2b610c1b565b5b6000610e3a85828601610c69565b9250506020610e4b85828601610c69565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000610e8f82610c7e565b9150610e9a83610c7e565b925082821015610ead57610eac610e55565b5b828203905092915050565b6000610ec382610c7e565b9150610ece83610c7e565b92508267ffffffffffffffff03821115610eeb57610eea610e55565b5b82820190509291505056fea264697066735822122045fa59504baf34a4ae5010c8ad97929a7361586a8581454e4a04d430332d2cad64736f6c634300080f0033",
}

// FiatTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use FiatTokenMetaData.ABI instead.
var FiatTokenABI = FiatTokenMetaData.ABI

// FiatTokenBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use FiatTokenMetaData.Bin instead.
var FiatTokenBin = FiatTokenMetaData.Bin

// DeployFiatToken deploys a new Ethereum contract, binding an instance of FiatToken to it.
func DeployFiatToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *FiatToken, error) {
	parsed, err := FiatTokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(FiatTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FiatToken{FiatTokenCaller: FiatTokenCaller{contract: contract}, FiatTokenTransactor: FiatTokenTransactor{contract: contract}, FiatTokenFilterer: FiatTokenFilterer{contract: contract}}, nil
}

// FiatToken is an auto generated Go binding around an Ethereum contract.
type FiatToken struct {
	FiatTokenCaller     // Read-only binding to the contract
	FiatTokenTransactor // Write-only binding to the contract
	FiatTokenFilterer   // Log filterer for contract events
}

// FiatTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type FiatTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FiatTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FiatTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FiatTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FiatTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FiatTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FiatTokenSession struct {
	Contract     *FiatToken        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FiatTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FiatTokenCallerSession struct {
	Contract *FiatTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// FiatTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FiatTokenTransactorSession struct {
	Contract     *FiatTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// FiatTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type FiatTokenRaw struct {
	Contract *FiatToken // Generic contract binding to access the raw methods on
}

// FiatTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FiatTokenCallerRaw struct {
	Contract *FiatTokenCaller // Generic read-only contract binding to access the raw methods on
}

// FiatTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FiatTokenTransactorRaw struct {
	Contract *FiatTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFiatToken creates a new instance of FiatToken, bound to a specific deployed contract.
func NewFiatToken(address common.Address, backend bind.ContractBackend) (*FiatToken, error) {
	contract, err := bindFiatToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FiatToken{FiatTokenCaller: FiatTokenCaller{contract: contract}, FiatTokenTransactor: FiatTokenTransactor{contract: contract}, FiatTokenFilterer: FiatTokenFilterer{contract: contract}}, nil
}

// NewFiatTokenCaller creates a new read-only instance of FiatToken, bound to a specific deployed contract.
func NewFiatTokenCaller(address common.Address, caller bind.ContractCaller) (*FiatTokenCaller, error) {
	contract, err := bindFiatToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FiatTokenCaller{contract: contract}, nil
}

// NewFiatTokenTransactor creates a new write-only instance of FiatToken, bound to a specific deployed contract.
func NewFiatTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*FiatTokenTransactor, error) {
	contract, err := bindFiatToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FiatTokenTransactor{contract: contract}, nil
}

// NewFiatTokenFilterer creates a new log filterer instance of FiatToken, bound to a specific deployed contract.
func NewFiatTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*FiatTokenFilterer, error) {
	contract, err := bindFiatToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FiatTokenFilterer{contract: contract}, nil
}

// bindFiatToken binds a generic wrapper to an already deployed contract.
func bindFiatToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FiatTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FiatToken *FiatTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FiatToken.Contract.FiatTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FiatToken *FiatTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FiatToken.Contract.FiatTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FiatToken *FiatTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FiatToken.Contract.FiatTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FiatToken *FiatTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FiatToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FiatToken *FiatTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FiatToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FiatToken *FiatTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FiatToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address delegate) view returns(uint64)
func (_FiatToken *FiatTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, delegate common.Address) (uint64, error) {
	var out []interface{}
	err := _FiatToken.contract.Call(opts, &out, "allowance", owner, delegate)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address delegate) view returns(uint64)
func (_FiatToken *FiatTokenSession) Allowance(owner common.Address, delegate common.Address) (uint64, error) {
	return _FiatToken.Contract.Allowance(&_FiatToken.CallOpts, owner, delegate)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address delegate) view returns(uint64)
func (_FiatToken *FiatTokenCallerSession) Allowance(owner common.Address, delegate common.Address) (uint64, error) {
	return _FiatToken.Contract.Allowance(&_FiatToken.CallOpts, owner, delegate)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address tokenOwner) view returns(uint64)
func (_FiatToken *FiatTokenCaller) BalanceOf(opts *bind.CallOpts, tokenOwner common.Address) (uint64, error) {
	var out []interface{}
	err := _FiatToken.contract.Call(opts, &out, "balanceOf", tokenOwner)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address tokenOwner) view returns(uint64)
func (_FiatToken *FiatTokenSession) BalanceOf(tokenOwner common.Address) (uint64, error) {
	return _FiatToken.Contract.BalanceOf(&_FiatToken.CallOpts, tokenOwner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address tokenOwner) view returns(uint64)
func (_FiatToken *FiatTokenCallerSession) BalanceOf(tokenOwner common.Address) (uint64, error) {
	return _FiatToken.Contract.BalanceOf(&_FiatToken.CallOpts, tokenOwner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_FiatToken *FiatTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _FiatToken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_FiatToken *FiatTokenSession) Decimals() (uint8, error) {
	return _FiatToken.Contract.Decimals(&_FiatToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_FiatToken *FiatTokenCallerSession) Decimals() (uint8, error) {
	return _FiatToken.Contract.Decimals(&_FiatToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_FiatToken *FiatTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _FiatToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_FiatToken *FiatTokenSession) Name() (string, error) {
	return _FiatToken.Contract.Name(&_FiatToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_FiatToken *FiatTokenCallerSession) Name() (string, error) {
	return _FiatToken.Contract.Name(&_FiatToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_FiatToken *FiatTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _FiatToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_FiatToken *FiatTokenSession) Symbol() (string, error) {
	return _FiatToken.Contract.Symbol(&_FiatToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_FiatToken *FiatTokenCallerSession) Symbol() (string, error) {
	return _FiatToken.Contract.Symbol(&_FiatToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint64)
func (_FiatToken *FiatTokenCaller) TotalSupply(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _FiatToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint64)
func (_FiatToken *FiatTokenSession) TotalSupply() (uint64, error) {
	return _FiatToken.Contract.TotalSupply(&_FiatToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint64)
func (_FiatToken *FiatTokenCallerSession) TotalSupply() (uint64, error) {
	return _FiatToken.Contract.TotalSupply(&_FiatToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x1086a9aa.
//
// Solidity: function approve(address delegate, uint64 numTokens) returns(bool)
func (_FiatToken *FiatTokenTransactor) Approve(opts *bind.TransactOpts, delegate common.Address, numTokens uint64) (*types.Transaction, error) {
	return _FiatToken.contract.Transact(opts, "approve", delegate, numTokens)
}

// Approve is a paid mutator transaction binding the contract method 0x1086a9aa.
//
// Solidity: function approve(address delegate, uint64 numTokens) returns(bool)
func (_FiatToken *FiatTokenSession) Approve(delegate common.Address, numTokens uint64) (*types.Transaction, error) {
	return _FiatToken.Contract.Approve(&_FiatToken.TransactOpts, delegate, numTokens)
}

// Approve is a paid mutator transaction binding the contract method 0x1086a9aa.
//
// Solidity: function approve(address delegate, uint64 numTokens) returns(bool)
func (_FiatToken *FiatTokenTransactorSession) Approve(delegate common.Address, numTokens uint64) (*types.Transaction, error) {
	return _FiatToken.Contract.Approve(&_FiatToken.TransactOpts, delegate, numTokens)
}

// Transfer is a paid mutator transaction binding the contract method 0x5d359fbd.
//
// Solidity: function transfer(address receiver, uint64 numTokens) returns(bool)
func (_FiatToken *FiatTokenTransactor) Transfer(opts *bind.TransactOpts, receiver common.Address, numTokens uint64) (*types.Transaction, error) {
	return _FiatToken.contract.Transact(opts, "transfer", receiver, numTokens)
}

// Transfer is a paid mutator transaction binding the contract method 0x5d359fbd.
//
// Solidity: function transfer(address receiver, uint64 numTokens) returns(bool)
func (_FiatToken *FiatTokenSession) Transfer(receiver common.Address, numTokens uint64) (*types.Transaction, error) {
	return _FiatToken.Contract.Transfer(&_FiatToken.TransactOpts, receiver, numTokens)
}

// Transfer is a paid mutator transaction binding the contract method 0x5d359fbd.
//
// Solidity: function transfer(address receiver, uint64 numTokens) returns(bool)
func (_FiatToken *FiatTokenTransactorSession) Transfer(receiver common.Address, numTokens uint64) (*types.Transaction, error) {
	return _FiatToken.Contract.Transfer(&_FiatToken.TransactOpts, receiver, numTokens)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x2ea0dfe1.
//
// Solidity: function transferFrom(address owner, address buyer, uint64 numTokens) returns(bool)
func (_FiatToken *FiatTokenTransactor) TransferFrom(opts *bind.TransactOpts, owner common.Address, buyer common.Address, numTokens uint64) (*types.Transaction, error) {
	return _FiatToken.contract.Transact(opts, "transferFrom", owner, buyer, numTokens)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x2ea0dfe1.
//
// Solidity: function transferFrom(address owner, address buyer, uint64 numTokens) returns(bool)
func (_FiatToken *FiatTokenSession) TransferFrom(owner common.Address, buyer common.Address, numTokens uint64) (*types.Transaction, error) {
	return _FiatToken.Contract.TransferFrom(&_FiatToken.TransactOpts, owner, buyer, numTokens)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x2ea0dfe1.
//
// Solidity: function transferFrom(address owner, address buyer, uint64 numTokens) returns(bool)
func (_FiatToken *FiatTokenTransactorSession) TransferFrom(owner common.Address, buyer common.Address, numTokens uint64) (*types.Transaction, error) {
	return _FiatToken.Contract.TransferFrom(&_FiatToken.TransactOpts, owner, buyer, numTokens)
}

// FiatTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the FiatToken contract.
type FiatTokenApprovalIterator struct {
	Event *FiatTokenApproval // Event containing the contract specifics and raw log

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
func (it *FiatTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FiatTokenApproval)
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
		it.Event = new(FiatTokenApproval)
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
func (it *FiatTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FiatTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FiatTokenApproval represents a Approval event raised by the FiatToken contract.
type FiatTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x16304dfea7f3fbabcf59225f0629cb307fecb8d5652b069080aa9be2c765d7d2.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint64 value)
func (_FiatToken *FiatTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*FiatTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _FiatToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &FiatTokenApprovalIterator{contract: _FiatToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x16304dfea7f3fbabcf59225f0629cb307fecb8d5652b069080aa9be2c765d7d2.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint64 value)
func (_FiatToken *FiatTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *FiatTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _FiatToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FiatTokenApproval)
				if err := _FiatToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x16304dfea7f3fbabcf59225f0629cb307fecb8d5652b069080aa9be2c765d7d2.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint64 value)
func (_FiatToken *FiatTokenFilterer) ParseApproval(log types.Log) (*FiatTokenApproval, error) {
	event := new(FiatTokenApproval)
	if err := _FiatToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FiatTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the FiatToken contract.
type FiatTokenTransferIterator struct {
	Event *FiatTokenTransfer // Event containing the contract specifics and raw log

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
func (it *FiatTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FiatTokenTransfer)
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
		it.Event = new(FiatTokenTransfer)
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
func (it *FiatTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FiatTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FiatTokenTransfer represents a Transfer event raised by the FiatToken contract.
type FiatTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value uint64
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0x831ac82b07fb396dafef0077cea6e002235d88e63f35cbd5df2c065107f1e74a.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint64 value)
func (_FiatToken *FiatTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*FiatTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _FiatToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &FiatTokenTransferIterator{contract: _FiatToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0x831ac82b07fb396dafef0077cea6e002235d88e63f35cbd5df2c065107f1e74a.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint64 value)
func (_FiatToken *FiatTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *FiatTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _FiatToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FiatTokenTransfer)
				if err := _FiatToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0x831ac82b07fb396dafef0077cea6e002235d88e63f35cbd5df2c065107f1e74a.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint64 value)
func (_FiatToken *FiatTokenFilterer) ParseTransfer(log types.Log) (*FiatTokenTransfer, error) {
	event := new(FiatTokenTransfer)
	if err := _FiatToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
