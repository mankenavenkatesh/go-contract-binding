// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// StorageABI is the input ABI used to generate the binding from.
const StorageABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"newval\",\"type\":\"uint256\"}],\"name\":\"setVal\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"d\",\"type\":\"uint256\"},{\"name\":\"a\",\"type\":\"address\"},{\"name\":\"b\",\"type\":\"bytes32\"},{\"name\":\"b1\",\"type\":\"bytes\"}],\"name\":\"s\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getVal\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"vote\",\"type\":\"bytes\"},{\"name\":\"sigs\",\"type\":\"bytes\"},{\"name\":\"txData\",\"type\":\"bytes\"}],\"name\":\"submitHeaderBlock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Storage is an auto generated Go binding around an Ethereum contract.
type Storage struct {
	StorageCaller     // Read-only binding to the contract
	StorageTransactor // Write-only binding to the contract
	StorageFilterer   // Log filterer for contract events
}

// StorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type StorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StorageSession struct {
	Contract     *Storage          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StorageCallerSession struct {
	Contract *StorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// StorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StorageTransactorSession struct {
	Contract     *StorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// StorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type StorageRaw struct {
	Contract *Storage // Generic contract binding to access the raw methods on
}

// StorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StorageCallerRaw struct {
	Contract *StorageCaller // Generic read-only contract binding to access the raw methods on
}

// StorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StorageTransactorRaw struct {
	Contract *StorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStorage creates a new instance of Storage, bound to a specific deployed contract.
func NewStorage(address common.Address, backend bind.ContractBackend) (*Storage, error) {
	contract, err := bindStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Storage{StorageCaller: StorageCaller{contract: contract}, StorageTransactor: StorageTransactor{contract: contract}, StorageFilterer: StorageFilterer{contract: contract}}, nil
}

// NewStorageCaller creates a new read-only instance of Storage, bound to a specific deployed contract.
func NewStorageCaller(address common.Address, caller bind.ContractCaller) (*StorageCaller, error) {
	contract, err := bindStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StorageCaller{contract: contract}, nil
}

// NewStorageTransactor creates a new write-only instance of Storage, bound to a specific deployed contract.
func NewStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*StorageTransactor, error) {
	contract, err := bindStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StorageTransactor{contract: contract}, nil
}

// NewStorageFilterer creates a new log filterer instance of Storage, bound to a specific deployed contract.
func NewStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*StorageFilterer, error) {
	contract, err := bindStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StorageFilterer{contract: contract}, nil
}

// bindStorage binds a generic wrapper to an already deployed contract.
func bindStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StorageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Storage *StorageRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Storage.Contract.StorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Storage *StorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Storage.Contract.StorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Storage *StorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Storage.Contract.StorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Storage *StorageCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Storage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Storage *StorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Storage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Storage *StorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Storage.Contract.contract.Transact(opts, method, params...)
}

// GetVal is a free data retrieval call binding the contract method 0xe1cb0e52.
//
// Solidity: function getVal() constant returns(uint256)
func (_Storage *StorageCaller) GetVal(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Storage.contract.Call(opts, out, "getVal")
	return *ret0, err
}

// GetVal is a free data retrieval call binding the contract method 0xe1cb0e52.
//
// Solidity: function getVal() constant returns(uint256)
func (_Storage *StorageSession) GetVal() (*big.Int, error) {
	return _Storage.Contract.GetVal(&_Storage.CallOpts)
}

// GetVal is a free data retrieval call binding the contract method 0xe1cb0e52.
//
// Solidity: function getVal() constant returns(uint256)
func (_Storage *StorageCallerSession) GetVal() (*big.Int, error) {
	return _Storage.Contract.GetVal(&_Storage.CallOpts)
}

// S is a paid mutator transaction binding the contract method 0x618fb5ea.
//
// Solidity: function s(uint256 d, address a, bytes32 b, bytes b1) returns()
func (_Storage *StorageTransactor) S(opts *bind.TransactOpts, d *big.Int, a common.Address, b [32]byte, b1 []byte) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "s", d, a, b, b1)
}

// S is a paid mutator transaction binding the contract method 0x618fb5ea.
//
// Solidity: function s(uint256 d, address a, bytes32 b, bytes b1) returns()
func (_Storage *StorageSession) S(d *big.Int, a common.Address, b [32]byte, b1 []byte) (*types.Transaction, error) {
	return _Storage.Contract.S(&_Storage.TransactOpts, d, a, b, b1)
}

// S is a paid mutator transaction binding the contract method 0x618fb5ea.
//
// Solidity: function s(uint256 d, address a, bytes32 b, bytes b1) returns()
func (_Storage *StorageTransactorSession) S(d *big.Int, a common.Address, b [32]byte, b1 []byte) (*types.Transaction, error) {
	return _Storage.Contract.S(&_Storage.TransactOpts, d, a, b, b1)
}

// SetVal is a paid mutator transaction binding the contract method 0x3d4197f0.
//
// Solidity: function setVal(uint256 newval) returns()
func (_Storage *StorageTransactor) SetVal(opts *bind.TransactOpts, newval *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "setVal", newval)
}

// SetVal is a paid mutator transaction binding the contract method 0x3d4197f0.
//
// Solidity: function setVal(uint256 newval) returns()
func (_Storage *StorageSession) SetVal(newval *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.SetVal(&_Storage.TransactOpts, newval)
}

// SetVal is a paid mutator transaction binding the contract method 0x3d4197f0.
//
// Solidity: function setVal(uint256 newval) returns()
func (_Storage *StorageTransactorSession) SetVal(newval *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.SetVal(&_Storage.TransactOpts, newval)
}

// SubmitHeaderBlock is a paid mutator transaction binding the contract method 0xec83d3ba.
//
// Solidity: function submitHeaderBlock(bytes vote, bytes sigs, bytes txData) returns()
func (_Storage *StorageTransactor) SubmitHeaderBlock(opts *bind.TransactOpts, vote []byte, sigs []byte, txData []byte) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "submitHeaderBlock", vote, sigs, txData)
}

// SubmitHeaderBlock is a paid mutator transaction binding the contract method 0xec83d3ba.
//
// Solidity: function submitHeaderBlock(bytes vote, bytes sigs, bytes txData) returns()
func (_Storage *StorageSession) SubmitHeaderBlock(vote []byte, sigs []byte, txData []byte) (*types.Transaction, error) {
	return _Storage.Contract.SubmitHeaderBlock(&_Storage.TransactOpts, vote, sigs, txData)
}

// SubmitHeaderBlock is a paid mutator transaction binding the contract method 0xec83d3ba.
//
// Solidity: function submitHeaderBlock(bytes vote, bytes sigs, bytes txData) returns()
func (_Storage *StorageTransactorSession) SubmitHeaderBlock(vote []byte, sigs []byte, txData []byte) (*types.Transaction, error) {
	return _Storage.Contract.SubmitHeaderBlock(&_Storage.TransactOpts, vote, sigs, txData)
}
