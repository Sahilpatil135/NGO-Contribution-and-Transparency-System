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
	_ = abi.ConvertType
)

// CauseMilestoneEscrowMetaData contains all meta data concerning the CauseMilestoneEscrow contract.
var CauseMilestoneEscrowMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"CauseAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CauseNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBeneficiary\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidGoal\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroDonation\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes16\",\"name\":\"causeId\",\"type\":\"bytes16\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"goal\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"}],\"name\":\"CauseRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes16\",\"name\":\"causeId\",\"type\":\"bytes16\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"donor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newCollected\",\"type\":\"uint256\"}],\"name\":\"DonationReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes16\",\"name\":\"causeId\",\"type\":\"bytes16\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"milestone\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"}],\"name\":\"MilestoneReleased\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes16\",\"name\":\"causeId\",\"type\":\"bytes16\"}],\"name\":\"donate\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes16\",\"name\":\"causeId\",\"type\":\"bytes16\"}],\"name\":\"getCause\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"goal\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"released\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"milestonesPaid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes16\",\"name\":\"causeId\",\"type\":\"bytes16\"},{\"internalType\":\"uint256\",\"name\":\"goal\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"beneficiary\",\"type\":\"address\"}],\"name\":\"registerCause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// CauseMilestoneEscrowABI is the input ABI used to generate the binding from.
// Deprecated: Use CauseMilestoneEscrowMetaData.ABI instead.
var CauseMilestoneEscrowABI = CauseMilestoneEscrowMetaData.ABI

// CauseMilestoneEscrow is an auto generated Go binding around an Ethereum contract.
type CauseMilestoneEscrow struct {
	CauseMilestoneEscrowCaller     // Read-only binding to the contract
	CauseMilestoneEscrowTransactor // Write-only binding to the contract
	CauseMilestoneEscrowFilterer   // Log filterer for contract events
}

// CauseMilestoneEscrowCaller is an auto generated read-only Go binding around an Ethereum contract.
type CauseMilestoneEscrowCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CauseMilestoneEscrowTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CauseMilestoneEscrowTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CauseMilestoneEscrowFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CauseMilestoneEscrowFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CauseMilestoneEscrowSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CauseMilestoneEscrowSession struct {
	Contract     *CauseMilestoneEscrow // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// CauseMilestoneEscrowCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CauseMilestoneEscrowCallerSession struct {
	Contract *CauseMilestoneEscrowCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// CauseMilestoneEscrowTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CauseMilestoneEscrowTransactorSession struct {
	Contract     *CauseMilestoneEscrowTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// CauseMilestoneEscrowRaw is an auto generated low-level Go binding around an Ethereum contract.
type CauseMilestoneEscrowRaw struct {
	Contract *CauseMilestoneEscrow // Generic contract binding to access the raw methods on
}

// CauseMilestoneEscrowCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CauseMilestoneEscrowCallerRaw struct {
	Contract *CauseMilestoneEscrowCaller // Generic read-only contract binding to access the raw methods on
}

// CauseMilestoneEscrowTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CauseMilestoneEscrowTransactorRaw struct {
	Contract *CauseMilestoneEscrowTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCauseMilestoneEscrow creates a new instance of CauseMilestoneEscrow, bound to a specific deployed contract.
func NewCauseMilestoneEscrow(address common.Address, backend bind.ContractBackend) (*CauseMilestoneEscrow, error) {
	contract, err := bindCauseMilestoneEscrow(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CauseMilestoneEscrow{CauseMilestoneEscrowCaller: CauseMilestoneEscrowCaller{contract: contract}, CauseMilestoneEscrowTransactor: CauseMilestoneEscrowTransactor{contract: contract}, CauseMilestoneEscrowFilterer: CauseMilestoneEscrowFilterer{contract: contract}}, nil
}

// NewCauseMilestoneEscrowCaller creates a new read-only instance of CauseMilestoneEscrow, bound to a specific deployed contract.
func NewCauseMilestoneEscrowCaller(address common.Address, caller bind.ContractCaller) (*CauseMilestoneEscrowCaller, error) {
	contract, err := bindCauseMilestoneEscrow(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CauseMilestoneEscrowCaller{contract: contract}, nil
}

// NewCauseMilestoneEscrowTransactor creates a new write-only instance of CauseMilestoneEscrow, bound to a specific deployed contract.
func NewCauseMilestoneEscrowTransactor(address common.Address, transactor bind.ContractTransactor) (*CauseMilestoneEscrowTransactor, error) {
	contract, err := bindCauseMilestoneEscrow(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CauseMilestoneEscrowTransactor{contract: contract}, nil
}

// NewCauseMilestoneEscrowFilterer creates a new log filterer instance of CauseMilestoneEscrow, bound to a specific deployed contract.
func NewCauseMilestoneEscrowFilterer(address common.Address, filterer bind.ContractFilterer) (*CauseMilestoneEscrowFilterer, error) {
	contract, err := bindCauseMilestoneEscrow(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CauseMilestoneEscrowFilterer{contract: contract}, nil
}

// bindCauseMilestoneEscrow binds a generic wrapper to an already deployed contract.
func bindCauseMilestoneEscrow(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CauseMilestoneEscrowMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CauseMilestoneEscrow *CauseMilestoneEscrowRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CauseMilestoneEscrow.Contract.CauseMilestoneEscrowCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CauseMilestoneEscrow *CauseMilestoneEscrowRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CauseMilestoneEscrow.Contract.CauseMilestoneEscrowTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CauseMilestoneEscrow *CauseMilestoneEscrowRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CauseMilestoneEscrow.Contract.CauseMilestoneEscrowTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CauseMilestoneEscrow *CauseMilestoneEscrowCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CauseMilestoneEscrow.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CauseMilestoneEscrow *CauseMilestoneEscrowTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CauseMilestoneEscrow.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CauseMilestoneEscrow *CauseMilestoneEscrowTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CauseMilestoneEscrow.Contract.contract.Transact(opts, method, params...)
}

// GetCause is a free data retrieval call binding the contract method 0x331435a7.
//
// Solidity: function getCause(bytes16 causeId) view returns(uint256 goal, uint256 collected, uint256 released, uint256 milestonesPaid, address beneficiary, bool exists)
func (_CauseMilestoneEscrow *CauseMilestoneEscrowCaller) GetCause(opts *bind.CallOpts, causeId [16]byte) (struct {
	Goal           *big.Int
	Collected      *big.Int
	Released       *big.Int
	MilestonesPaid *big.Int
	Beneficiary    common.Address
	Exists         bool
}, error) {
	var out []interface{}
	err := _CauseMilestoneEscrow.contract.Call(opts, &out, "getCause", causeId)

	outstruct := new(struct {
		Goal           *big.Int
		Collected      *big.Int
		Released       *big.Int
		MilestonesPaid *big.Int
		Beneficiary    common.Address
		Exists         bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Goal = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Collected = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Released = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.MilestonesPaid = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Beneficiary = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Exists = *abi.ConvertType(out[5], new(bool)).(*bool)

	return *outstruct, err

}

// GetCause is a free data retrieval call binding the contract method 0x331435a7.
//
// Solidity: function getCause(bytes16 causeId) view returns(uint256 goal, uint256 collected, uint256 released, uint256 milestonesPaid, address beneficiary, bool exists)
func (_CauseMilestoneEscrow *CauseMilestoneEscrowSession) GetCause(causeId [16]byte) (struct {
	Goal           *big.Int
	Collected      *big.Int
	Released       *big.Int
	MilestonesPaid *big.Int
	Beneficiary    common.Address
	Exists         bool
}, error) {
	return _CauseMilestoneEscrow.Contract.GetCause(&_CauseMilestoneEscrow.CallOpts, causeId)
}

// GetCause is a free data retrieval call binding the contract method 0x331435a7.
//
// Solidity: function getCause(bytes16 causeId) view returns(uint256 goal, uint256 collected, uint256 released, uint256 milestonesPaid, address beneficiary, bool exists)
func (_CauseMilestoneEscrow *CauseMilestoneEscrowCallerSession) GetCause(causeId [16]byte) (struct {
	Goal           *big.Int
	Collected      *big.Int
	Released       *big.Int
	MilestonesPaid *big.Int
	Beneficiary    common.Address
	Exists         bool
}, error) {
	return _CauseMilestoneEscrow.Contract.GetCause(&_CauseMilestoneEscrow.CallOpts, causeId)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CauseMilestoneEscrow *CauseMilestoneEscrowCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CauseMilestoneEscrow.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CauseMilestoneEscrow *CauseMilestoneEscrowSession) Owner() (common.Address, error) {
	return _CauseMilestoneEscrow.Contract.Owner(&_CauseMilestoneEscrow.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CauseMilestoneEscrow *CauseMilestoneEscrowCallerSession) Owner() (common.Address, error) {
	return _CauseMilestoneEscrow.Contract.Owner(&_CauseMilestoneEscrow.CallOpts)
}

// Donate is a paid mutator transaction binding the contract method 0xba6bbaa4.
//
// Solidity: function donate(bytes16 causeId) payable returns()
func (_CauseMilestoneEscrow *CauseMilestoneEscrowTransactor) Donate(opts *bind.TransactOpts, causeId [16]byte) (*types.Transaction, error) {
	return _CauseMilestoneEscrow.contract.Transact(opts, "donate", causeId)
}

// Donate is a paid mutator transaction binding the contract method 0xba6bbaa4.
//
// Solidity: function donate(bytes16 causeId) payable returns()
func (_CauseMilestoneEscrow *CauseMilestoneEscrowSession) Donate(causeId [16]byte) (*types.Transaction, error) {
	return _CauseMilestoneEscrow.Contract.Donate(&_CauseMilestoneEscrow.TransactOpts, causeId)
}

// Donate is a paid mutator transaction binding the contract method 0xba6bbaa4.
//
// Solidity: function donate(bytes16 causeId) payable returns()
func (_CauseMilestoneEscrow *CauseMilestoneEscrowTransactorSession) Donate(causeId [16]byte) (*types.Transaction, error) {
	return _CauseMilestoneEscrow.Contract.Donate(&_CauseMilestoneEscrow.TransactOpts, causeId)
}

// RegisterCause is a paid mutator transaction binding the contract method 0x2611e606.
//
// Solidity: function registerCause(bytes16 causeId, uint256 goal, address beneficiary) returns()
func (_CauseMilestoneEscrow *CauseMilestoneEscrowTransactor) RegisterCause(opts *bind.TransactOpts, causeId [16]byte, goal *big.Int, beneficiary common.Address) (*types.Transaction, error) {
	return _CauseMilestoneEscrow.contract.Transact(opts, "registerCause", causeId, goal, beneficiary)
}

// RegisterCause is a paid mutator transaction binding the contract method 0x2611e606.
//
// Solidity: function registerCause(bytes16 causeId, uint256 goal, address beneficiary) returns()
func (_CauseMilestoneEscrow *CauseMilestoneEscrowSession) RegisterCause(causeId [16]byte, goal *big.Int, beneficiary common.Address) (*types.Transaction, error) {
	return _CauseMilestoneEscrow.Contract.RegisterCause(&_CauseMilestoneEscrow.TransactOpts, causeId, goal, beneficiary)
}

// RegisterCause is a paid mutator transaction binding the contract method 0x2611e606.
//
// Solidity: function registerCause(bytes16 causeId, uint256 goal, address beneficiary) returns()
func (_CauseMilestoneEscrow *CauseMilestoneEscrowTransactorSession) RegisterCause(causeId [16]byte, goal *big.Int, beneficiary common.Address) (*types.Transaction, error) {
	return _CauseMilestoneEscrow.Contract.RegisterCause(&_CauseMilestoneEscrow.TransactOpts, causeId, goal, beneficiary)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CauseMilestoneEscrow *CauseMilestoneEscrowTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _CauseMilestoneEscrow.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CauseMilestoneEscrow *CauseMilestoneEscrowSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CauseMilestoneEscrow.Contract.TransferOwnership(&_CauseMilestoneEscrow.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CauseMilestoneEscrow *CauseMilestoneEscrowTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CauseMilestoneEscrow.Contract.TransferOwnership(&_CauseMilestoneEscrow.TransactOpts, newOwner)
}

// CauseMilestoneEscrowCauseRegisteredIterator is returned from FilterCauseRegistered and is used to iterate over the raw logs and unpacked data for CauseRegistered events raised by the CauseMilestoneEscrow contract.
type CauseMilestoneEscrowCauseRegisteredIterator struct {
	Event *CauseMilestoneEscrowCauseRegistered // Event containing the contract specifics and raw log

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
func (it *CauseMilestoneEscrowCauseRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CauseMilestoneEscrowCauseRegistered)
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
		it.Event = new(CauseMilestoneEscrowCauseRegistered)
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
func (it *CauseMilestoneEscrowCauseRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CauseMilestoneEscrowCauseRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CauseMilestoneEscrowCauseRegistered represents a CauseRegistered event raised by the CauseMilestoneEscrow contract.
type CauseMilestoneEscrowCauseRegistered struct {
	CauseId     [16]byte
	Goal        *big.Int
	Beneficiary common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCauseRegistered is a free log retrieval operation binding the contract event 0x43feb3d961f67df78f76056c0b1215b820eee97ac617448ed7fa169648fedd07.
//
// Solidity: event CauseRegistered(bytes16 indexed causeId, uint256 goal, address beneficiary)
func (_CauseMilestoneEscrow *CauseMilestoneEscrowFilterer) FilterCauseRegistered(opts *bind.FilterOpts, causeId [][16]byte) (*CauseMilestoneEscrowCauseRegisteredIterator, error) {

	var causeIdRule []interface{}
	for _, causeIdItem := range causeId {
		causeIdRule = append(causeIdRule, causeIdItem)
	}

	logs, sub, err := _CauseMilestoneEscrow.contract.FilterLogs(opts, "CauseRegistered", causeIdRule)
	if err != nil {
		return nil, err
	}
	return &CauseMilestoneEscrowCauseRegisteredIterator{contract: _CauseMilestoneEscrow.contract, event: "CauseRegistered", logs: logs, sub: sub}, nil
}

// WatchCauseRegistered is a free log subscription operation binding the contract event 0x43feb3d961f67df78f76056c0b1215b820eee97ac617448ed7fa169648fedd07.
//
// Solidity: event CauseRegistered(bytes16 indexed causeId, uint256 goal, address beneficiary)
func (_CauseMilestoneEscrow *CauseMilestoneEscrowFilterer) WatchCauseRegistered(opts *bind.WatchOpts, sink chan<- *CauseMilestoneEscrowCauseRegistered, causeId [][16]byte) (event.Subscription, error) {

	var causeIdRule []interface{}
	for _, causeIdItem := range causeId {
		causeIdRule = append(causeIdRule, causeIdItem)
	}

	logs, sub, err := _CauseMilestoneEscrow.contract.WatchLogs(opts, "CauseRegistered", causeIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CauseMilestoneEscrowCauseRegistered)
				if err := _CauseMilestoneEscrow.contract.UnpackLog(event, "CauseRegistered", log); err != nil {
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

// ParseCauseRegistered is a log parse operation binding the contract event 0x43feb3d961f67df78f76056c0b1215b820eee97ac617448ed7fa169648fedd07.
//
// Solidity: event CauseRegistered(bytes16 indexed causeId, uint256 goal, address beneficiary)
func (_CauseMilestoneEscrow *CauseMilestoneEscrowFilterer) ParseCauseRegistered(log types.Log) (*CauseMilestoneEscrowCauseRegistered, error) {
	event := new(CauseMilestoneEscrowCauseRegistered)
	if err := _CauseMilestoneEscrow.contract.UnpackLog(event, "CauseRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CauseMilestoneEscrowDonationReceivedIterator is returned from FilterDonationReceived and is used to iterate over the raw logs and unpacked data for DonationReceived events raised by the CauseMilestoneEscrow contract.
type CauseMilestoneEscrowDonationReceivedIterator struct {
	Event *CauseMilestoneEscrowDonationReceived // Event containing the contract specifics and raw log

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
func (it *CauseMilestoneEscrowDonationReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CauseMilestoneEscrowDonationReceived)
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
		it.Event = new(CauseMilestoneEscrowDonationReceived)
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
func (it *CauseMilestoneEscrowDonationReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CauseMilestoneEscrowDonationReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CauseMilestoneEscrowDonationReceived represents a DonationReceived event raised by the CauseMilestoneEscrow contract.
type CauseMilestoneEscrowDonationReceived struct {
	CauseId      [16]byte
	Donor        common.Address
	Amount       *big.Int
	NewCollected *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDonationReceived is a free log retrieval operation binding the contract event 0x4a0ed74de833312cb350c60682bb15380307793f06f35c3e3a21c53db5031490.
//
// Solidity: event DonationReceived(bytes16 indexed causeId, address indexed donor, uint256 amount, uint256 newCollected)
func (_CauseMilestoneEscrow *CauseMilestoneEscrowFilterer) FilterDonationReceived(opts *bind.FilterOpts, causeId [][16]byte, donor []common.Address) (*CauseMilestoneEscrowDonationReceivedIterator, error) {

	var causeIdRule []interface{}
	for _, causeIdItem := range causeId {
		causeIdRule = append(causeIdRule, causeIdItem)
	}
	var donorRule []interface{}
	for _, donorItem := range donor {
		donorRule = append(donorRule, donorItem)
	}

	logs, sub, err := _CauseMilestoneEscrow.contract.FilterLogs(opts, "DonationReceived", causeIdRule, donorRule)
	if err != nil {
		return nil, err
	}
	return &CauseMilestoneEscrowDonationReceivedIterator{contract: _CauseMilestoneEscrow.contract, event: "DonationReceived", logs: logs, sub: sub}, nil
}

// WatchDonationReceived is a free log subscription operation binding the contract event 0x4a0ed74de833312cb350c60682bb15380307793f06f35c3e3a21c53db5031490.
//
// Solidity: event DonationReceived(bytes16 indexed causeId, address indexed donor, uint256 amount, uint256 newCollected)
func (_CauseMilestoneEscrow *CauseMilestoneEscrowFilterer) WatchDonationReceived(opts *bind.WatchOpts, sink chan<- *CauseMilestoneEscrowDonationReceived, causeId [][16]byte, donor []common.Address) (event.Subscription, error) {

	var causeIdRule []interface{}
	for _, causeIdItem := range causeId {
		causeIdRule = append(causeIdRule, causeIdItem)
	}
	var donorRule []interface{}
	for _, donorItem := range donor {
		donorRule = append(donorRule, donorItem)
	}

	logs, sub, err := _CauseMilestoneEscrow.contract.WatchLogs(opts, "DonationReceived", causeIdRule, donorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CauseMilestoneEscrowDonationReceived)
				if err := _CauseMilestoneEscrow.contract.UnpackLog(event, "DonationReceived", log); err != nil {
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

// ParseDonationReceived is a log parse operation binding the contract event 0x4a0ed74de833312cb350c60682bb15380307793f06f35c3e3a21c53db5031490.
//
// Solidity: event DonationReceived(bytes16 indexed causeId, address indexed donor, uint256 amount, uint256 newCollected)
func (_CauseMilestoneEscrow *CauseMilestoneEscrowFilterer) ParseDonationReceived(log types.Log) (*CauseMilestoneEscrowDonationReceived, error) {
	event := new(CauseMilestoneEscrowDonationReceived)
	if err := _CauseMilestoneEscrow.contract.UnpackLog(event, "DonationReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CauseMilestoneEscrowMilestoneReleasedIterator is returned from FilterMilestoneReleased and is used to iterate over the raw logs and unpacked data for MilestoneReleased events raised by the CauseMilestoneEscrow contract.
type CauseMilestoneEscrowMilestoneReleasedIterator struct {
	Event *CauseMilestoneEscrowMilestoneReleased // Event containing the contract specifics and raw log

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
func (it *CauseMilestoneEscrowMilestoneReleasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CauseMilestoneEscrowMilestoneReleased)
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
		it.Event = new(CauseMilestoneEscrowMilestoneReleased)
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
func (it *CauseMilestoneEscrowMilestoneReleasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CauseMilestoneEscrowMilestoneReleasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CauseMilestoneEscrowMilestoneReleased represents a MilestoneReleased event raised by the CauseMilestoneEscrow contract.
type CauseMilestoneEscrowMilestoneReleased struct {
	CauseId     [16]byte
	Milestone   uint8
	Amount      *big.Int
	Beneficiary common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMilestoneReleased is a free log retrieval operation binding the contract event 0x9d89fab7f9b64557f2f5c34e46b58fb977989a4cc728af6c48f1b3be6c725700.
//
// Solidity: event MilestoneReleased(bytes16 indexed causeId, uint8 milestone, uint256 amount, address indexed beneficiary)
func (_CauseMilestoneEscrow *CauseMilestoneEscrowFilterer) FilterMilestoneReleased(opts *bind.FilterOpts, causeId [][16]byte, beneficiary []common.Address) (*CauseMilestoneEscrowMilestoneReleasedIterator, error) {

	var causeIdRule []interface{}
	for _, causeIdItem := range causeId {
		causeIdRule = append(causeIdRule, causeIdItem)
	}

	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _CauseMilestoneEscrow.contract.FilterLogs(opts, "MilestoneReleased", causeIdRule, beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return &CauseMilestoneEscrowMilestoneReleasedIterator{contract: _CauseMilestoneEscrow.contract, event: "MilestoneReleased", logs: logs, sub: sub}, nil
}

// WatchMilestoneReleased is a free log subscription operation binding the contract event 0x9d89fab7f9b64557f2f5c34e46b58fb977989a4cc728af6c48f1b3be6c725700.
//
// Solidity: event MilestoneReleased(bytes16 indexed causeId, uint8 milestone, uint256 amount, address indexed beneficiary)
func (_CauseMilestoneEscrow *CauseMilestoneEscrowFilterer) WatchMilestoneReleased(opts *bind.WatchOpts, sink chan<- *CauseMilestoneEscrowMilestoneReleased, causeId [][16]byte, beneficiary []common.Address) (event.Subscription, error) {

	var causeIdRule []interface{}
	for _, causeIdItem := range causeId {
		causeIdRule = append(causeIdRule, causeIdItem)
	}

	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _CauseMilestoneEscrow.contract.WatchLogs(opts, "MilestoneReleased", causeIdRule, beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CauseMilestoneEscrowMilestoneReleased)
				if err := _CauseMilestoneEscrow.contract.UnpackLog(event, "MilestoneReleased", log); err != nil {
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

// ParseMilestoneReleased is a log parse operation binding the contract event 0x9d89fab7f9b64557f2f5c34e46b58fb977989a4cc728af6c48f1b3be6c725700.
//
// Solidity: event MilestoneReleased(bytes16 indexed causeId, uint8 milestone, uint256 amount, address indexed beneficiary)
func (_CauseMilestoneEscrow *CauseMilestoneEscrowFilterer) ParseMilestoneReleased(log types.Log) (*CauseMilestoneEscrowMilestoneReleased, error) {
	event := new(CauseMilestoneEscrowMilestoneReleased)
	if err := _CauseMilestoneEscrow.contract.UnpackLog(event, "MilestoneReleased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
