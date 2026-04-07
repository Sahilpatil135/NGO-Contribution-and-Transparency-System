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

// MilestoneTrackerMetaData contains all meta data concerning the MilestoneTracker contract.
var MilestoneTrackerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes16\",\"name\":\"causeId\",\"type\":\"bytes16\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"goal\",\"type\":\"uint256\"}],\"name\":\"CauseRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes16\",\"name\":\"causeId\",\"type\":\"bytes16\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newCollected\",\"type\":\"uint256\"}],\"name\":\"DonationRecorded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes16\",\"name\":\"causeId\",\"type\":\"bytes16\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"milestone\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountToDiburse\",\"type\":\"uint256\"}],\"name\":\"MilestoneReached\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes16\",\"name\":\"\",\"type\":\"bytes16\"}],\"name\":\"causes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"goal\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collected\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"milestonesPaid\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes16\",\"name\":\"causeId\",\"type\":\"bytes16\"}],\"name\":\"getCause\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"goal\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collected\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"milestonesPaid\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes16\",\"name\":\"causeId\",\"type\":\"bytes16\"},{\"internalType\":\"uint8\",\"name\":\"milestone\",\"type\":\"uint8\"}],\"name\":\"isMilestoneReached\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes16\",\"name\":\"causeId\",\"type\":\"bytes16\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"recordDonation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes16\",\"name\":\"causeId\",\"type\":\"bytes16\"},{\"internalType\":\"uint256\",\"name\":\"goal\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"initialCollected\",\"type\":\"uint256\"}],\"name\":\"registerOrUpdateCause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600f57600080fd5b50610d568061001f6000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c8063331435a71461005c5780634299071d1461008f57806345192280146100ab57806353553331146100c757806381553084146100fa575b600080fd5b610076600480360381019061007191906107b0565b61012a565b604051610086949392919061082d565b60405180910390f35b6100a960048036038101906100a4919061089e565b6101aa565b005b6100c560048036038101906100c091906108de565b61032c565b005b6100e160048036038101906100dc91906107b0565b610555565b6040516100f1949392919061082d565b60405180910390f35b610114600480360381019061010f919061095d565b61059f565b604051610121919061099d565b60405180910390f35b6000806000806000806000876fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff191681526020019081526020016000209050806000015481600101548260020160009054906101000a900460ff168360020160019054906101000a900460ff169450945094509450509193509193565b600080836fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff1916815260200190815260200160002060020160019054906101000a900460ff16610232576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161022990610a15565b60405180910390fd5b60008111610275576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161026c90610a81565b60405180910390fd5b6000806000846fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff191681526020019081526020016000209050818160010160008282546102c59190610ad0565b92505081905550826fffffffffffffffffffffffffffffffff19167f44b218be164f6bac27fc56af0a7b202413c77f42782416d6aa7a003d9d1b8fee838360010154604051610315929190610b04565b60405180910390a26103278382610652565b505050565b6000821161036f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161036690610b79565b60405180910390fd5b600080846fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff1916815260200190815260200160002060020160019054906101000a900460ff1661050f576040518060800160405280838152602001828152602001600060ff16815260200160011515815250600080856fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff19168152602001908152602001600020600082015181600001556020820151816001015560408201518160020160006101000a81548160ff021916908360ff16021790555060608201518160020160016101000a81548160ff021916908315150217905550905050826fffffffffffffffffffffffffffffffff19167f145d9a2cdce81cbe08debbe8acfe1c129851e6c5f083000408ed2dba5045f7b6836040516104b79190610b99565b60405180910390a2600081111561050a5761050983600080866fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff19168152602001908152602001600020610652565b5b610550565b81600080856fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff19168152602001908152602001600020600001819055505b505050565b60006020528060005260406000206000915090508060000154908060010154908060020160009054906101000a900460ff16908060020160019054906101000a900460ff16905084565b600060018260ff16101580156105b9575060048260ff1611155b6105f8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105ef90610c00565b60405180910390fd5b8160ff16600080856fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff1916815260200190815260200160002060020160009054906101000a900460ff1660ff161015905092915050565b5b60048160020160009054906101000a900460ff1660ff16101561074f57600060018260020160009054906101000a900460ff166106909190610c20565b90506000600483600001548360ff166106a99190610c55565b6106b39190610cc6565b905080836001015410156106c857505061074f565b6000600484600001546106db9190610cc6565b9050828460020160006101000a81548160ff021916908360ff160217905550846fffffffffffffffffffffffffffffffff19167f2ea29736068cc2fb33847eb950da5cf65c733582dce27abd14f43115c4d9ed67848360405161073f929190610cf7565b60405180910390a2505050610653565b5050565b600080fd5b60007fffffffffffffffffffffffffffffffff0000000000000000000000000000000082169050919050565b61078d81610758565b811461079857600080fd5b50565b6000813590506107aa81610784565b92915050565b6000602082840312156107c6576107c5610753565b5b60006107d48482850161079b565b91505092915050565b6000819050919050565b6107f0816107dd565b82525050565b600060ff82169050919050565b61080c816107f6565b82525050565b60008115159050919050565b61082781610812565b82525050565b600060808201905061084260008301876107e7565b61084f60208301866107e7565b61085c6040830185610803565b610869606083018461081e565b95945050505050565b61087b816107dd565b811461088657600080fd5b50565b60008135905061089881610872565b92915050565b600080604083850312156108b5576108b4610753565b5b60006108c38582860161079b565b92505060206108d485828601610889565b9150509250929050565b6000806000606084860312156108f7576108f6610753565b5b60006109058682870161079b565b935050602061091686828701610889565b925050604061092786828701610889565b9150509250925092565b61093a816107f6565b811461094557600080fd5b50565b60008135905061095781610931565b92915050565b6000806040838503121561097457610973610753565b5b60006109828582860161079b565b925050602061099385828601610948565b9150509250929050565b60006020820190506109b2600083018461081e565b92915050565b600082825260208201905092915050565b7f4361757365206e6f742072656769737465726564000000000000000000000000600082015250565b60006109ff6014836109b8565b9150610a0a826109c9565b602082019050919050565b60006020820190508181036000830152610a2e816109f2565b9050919050565b7f416d6f756e74206d7573742062652067726561746572207468616e207a65726f600082015250565b6000610a6b6020836109b8565b9150610a7682610a35565b602082019050919050565b60006020820190508181036000830152610a9a81610a5e565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000610adb826107dd565b9150610ae6836107dd565b9250828201905080821115610afe57610afd610aa1565b5b92915050565b6000604082019050610b1960008301856107e7565b610b2660208301846107e7565b9392505050565b7f476f616c206d7573742062652067726561746572207468616e207a65726f0000600082015250565b6000610b63601e836109b8565b9150610b6e82610b2d565b602082019050919050565b60006020820190508181036000830152610b9281610b56565b9050919050565b6000602082019050610bae60008301846107e7565b92915050565b7f4d696c6573746f6e65206d75737420626520312d340000000000000000000000600082015250565b6000610bea6015836109b8565b9150610bf582610bb4565b602082019050919050565b60006020820190508181036000830152610c1981610bdd565b9050919050565b6000610c2b826107f6565b9150610c36836107f6565b9250828201905060ff811115610c4f57610c4e610aa1565b5b92915050565b6000610c60826107dd565b9150610c6b836107dd565b9250828202610c79816107dd565b91508282048414831517610c9057610c8f610aa1565b5b5092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b6000610cd1826107dd565b9150610cdc836107dd565b925082610cec57610ceb610c97565b5b828204905092915050565b6000604082019050610d0c6000830185610803565b610d1960208301846107e7565b939250505056fea2646970667358221220053a9c0dfd1950338aa82b82bc5a7bceb5b828b2f8ff6132a4164e4fe208ce6e64736f6c634300081c0033",
}

// MilestoneTrackerABI is the input ABI used to generate the binding from.
// Deprecated: Use MilestoneTrackerMetaData.ABI instead.
var MilestoneTrackerABI = MilestoneTrackerMetaData.ABI

// MilestoneTrackerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MilestoneTrackerMetaData.Bin instead.
var MilestoneTrackerBin = MilestoneTrackerMetaData.Bin

// DeployMilestoneTracker deploys a new Ethereum contract, binding an instance of MilestoneTracker to it.
func DeployMilestoneTracker(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MilestoneTracker, error) {
	parsed, err := MilestoneTrackerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MilestoneTrackerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MilestoneTracker{MilestoneTrackerCaller: MilestoneTrackerCaller{contract: contract}, MilestoneTrackerTransactor: MilestoneTrackerTransactor{contract: contract}, MilestoneTrackerFilterer: MilestoneTrackerFilterer{contract: contract}}, nil
}

// MilestoneTracker is an auto generated Go binding around an Ethereum contract.
type MilestoneTracker struct {
	MilestoneTrackerCaller     // Read-only binding to the contract
	MilestoneTrackerTransactor // Write-only binding to the contract
	MilestoneTrackerFilterer   // Log filterer for contract events
}

// MilestoneTrackerCaller is an auto generated read-only Go binding around an Ethereum contract.
type MilestoneTrackerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MilestoneTrackerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MilestoneTrackerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MilestoneTrackerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MilestoneTrackerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MilestoneTrackerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MilestoneTrackerSession struct {
	Contract     *MilestoneTracker // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MilestoneTrackerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MilestoneTrackerCallerSession struct {
	Contract *MilestoneTrackerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// MilestoneTrackerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MilestoneTrackerTransactorSession struct {
	Contract     *MilestoneTrackerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// MilestoneTrackerRaw is an auto generated low-level Go binding around an Ethereum contract.
type MilestoneTrackerRaw struct {
	Contract *MilestoneTracker // Generic contract binding to access the raw methods on
}

// MilestoneTrackerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MilestoneTrackerCallerRaw struct {
	Contract *MilestoneTrackerCaller // Generic read-only contract binding to access the raw methods on
}

// MilestoneTrackerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MilestoneTrackerTransactorRaw struct {
	Contract *MilestoneTrackerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMilestoneTracker creates a new instance of MilestoneTracker, bound to a specific deployed contract.
func NewMilestoneTracker(address common.Address, backend bind.ContractBackend) (*MilestoneTracker, error) {
	contract, err := bindMilestoneTracker(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MilestoneTracker{MilestoneTrackerCaller: MilestoneTrackerCaller{contract: contract}, MilestoneTrackerTransactor: MilestoneTrackerTransactor{contract: contract}, MilestoneTrackerFilterer: MilestoneTrackerFilterer{contract: contract}}, nil
}

// NewMilestoneTrackerCaller creates a new read-only instance of MilestoneTracker, bound to a specific deployed contract.
func NewMilestoneTrackerCaller(address common.Address, caller bind.ContractCaller) (*MilestoneTrackerCaller, error) {
	contract, err := bindMilestoneTracker(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MilestoneTrackerCaller{contract: contract}, nil
}

// NewMilestoneTrackerTransactor creates a new write-only instance of MilestoneTracker, bound to a specific deployed contract.
func NewMilestoneTrackerTransactor(address common.Address, transactor bind.ContractTransactor) (*MilestoneTrackerTransactor, error) {
	contract, err := bindMilestoneTracker(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MilestoneTrackerTransactor{contract: contract}, nil
}

// NewMilestoneTrackerFilterer creates a new log filterer instance of MilestoneTracker, bound to a specific deployed contract.
func NewMilestoneTrackerFilterer(address common.Address, filterer bind.ContractFilterer) (*MilestoneTrackerFilterer, error) {
	contract, err := bindMilestoneTracker(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MilestoneTrackerFilterer{contract: contract}, nil
}

// bindMilestoneTracker binds a generic wrapper to an already deployed contract.
func bindMilestoneTracker(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MilestoneTrackerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MilestoneTracker *MilestoneTrackerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MilestoneTracker.Contract.MilestoneTrackerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MilestoneTracker *MilestoneTrackerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MilestoneTracker.Contract.MilestoneTrackerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MilestoneTracker *MilestoneTrackerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MilestoneTracker.Contract.MilestoneTrackerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MilestoneTracker *MilestoneTrackerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MilestoneTracker.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MilestoneTracker *MilestoneTrackerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MilestoneTracker.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MilestoneTracker *MilestoneTrackerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MilestoneTracker.Contract.contract.Transact(opts, method, params...)
}

// Causes is a free data retrieval call binding the contract method 0x53553331.
//
// Solidity: function causes(bytes16 ) view returns(uint256 goal, uint256 collected, uint8 milestonesPaid, bool exists)
func (_MilestoneTracker *MilestoneTrackerCaller) Causes(opts *bind.CallOpts, arg0 [16]byte) (struct {
	Goal           *big.Int
	Collected      *big.Int
	MilestonesPaid uint8
	Exists         bool
}, error) {
	var out []interface{}
	err := _MilestoneTracker.contract.Call(opts, &out, "causes", arg0)

	outstruct := new(struct {
		Goal           *big.Int
		Collected      *big.Int
		MilestonesPaid uint8
		Exists         bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Goal = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Collected = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.MilestonesPaid = *abi.ConvertType(out[2], new(uint8)).(*uint8)
	outstruct.Exists = *abi.ConvertType(out[3], new(bool)).(*bool)

	return *outstruct, err

}

// Causes is a free data retrieval call binding the contract method 0x53553331.
//
// Solidity: function causes(bytes16 ) view returns(uint256 goal, uint256 collected, uint8 milestonesPaid, bool exists)
func (_MilestoneTracker *MilestoneTrackerSession) Causes(arg0 [16]byte) (struct {
	Goal           *big.Int
	Collected      *big.Int
	MilestonesPaid uint8
	Exists         bool
}, error) {
	return _MilestoneTracker.Contract.Causes(&_MilestoneTracker.CallOpts, arg0)
}

// Causes is a free data retrieval call binding the contract method 0x53553331.
//
// Solidity: function causes(bytes16 ) view returns(uint256 goal, uint256 collected, uint8 milestonesPaid, bool exists)
func (_MilestoneTracker *MilestoneTrackerCallerSession) Causes(arg0 [16]byte) (struct {
	Goal           *big.Int
	Collected      *big.Int
	MilestonesPaid uint8
	Exists         bool
}, error) {
	return _MilestoneTracker.Contract.Causes(&_MilestoneTracker.CallOpts, arg0)
}

// GetCause is a free data retrieval call binding the contract method 0x331435a7.
//
// Solidity: function getCause(bytes16 causeId) view returns(uint256 goal, uint256 collected, uint8 milestonesPaid, bool exists)
func (_MilestoneTracker *MilestoneTrackerCaller) GetCause(opts *bind.CallOpts, causeId [16]byte) (struct {
	Goal           *big.Int
	Collected      *big.Int
	MilestonesPaid uint8
	Exists         bool
}, error) {
	var out []interface{}
	err := _MilestoneTracker.contract.Call(opts, &out, "getCause", causeId)

	outstruct := new(struct {
		Goal           *big.Int
		Collected      *big.Int
		MilestonesPaid uint8
		Exists         bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Goal = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Collected = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.MilestonesPaid = *abi.ConvertType(out[2], new(uint8)).(*uint8)
	outstruct.Exists = *abi.ConvertType(out[3], new(bool)).(*bool)

	return *outstruct, err

}

// GetCause is a free data retrieval call binding the contract method 0x331435a7.
//
// Solidity: function getCause(bytes16 causeId) view returns(uint256 goal, uint256 collected, uint8 milestonesPaid, bool exists)
func (_MilestoneTracker *MilestoneTrackerSession) GetCause(causeId [16]byte) (struct {
	Goal           *big.Int
	Collected      *big.Int
	MilestonesPaid uint8
	Exists         bool
}, error) {
	return _MilestoneTracker.Contract.GetCause(&_MilestoneTracker.CallOpts, causeId)
}

// GetCause is a free data retrieval call binding the contract method 0x331435a7.
//
// Solidity: function getCause(bytes16 causeId) view returns(uint256 goal, uint256 collected, uint8 milestonesPaid, bool exists)
func (_MilestoneTracker *MilestoneTrackerCallerSession) GetCause(causeId [16]byte) (struct {
	Goal           *big.Int
	Collected      *big.Int
	MilestonesPaid uint8
	Exists         bool
}, error) {
	return _MilestoneTracker.Contract.GetCause(&_MilestoneTracker.CallOpts, causeId)
}

// IsMilestoneReached is a free data retrieval call binding the contract method 0x81553084.
//
// Solidity: function isMilestoneReached(bytes16 causeId, uint8 milestone) view returns(bool)
func (_MilestoneTracker *MilestoneTrackerCaller) IsMilestoneReached(opts *bind.CallOpts, causeId [16]byte, milestone uint8) (bool, error) {
	var out []interface{}
	err := _MilestoneTracker.contract.Call(opts, &out, "isMilestoneReached", causeId, milestone)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMilestoneReached is a free data retrieval call binding the contract method 0x81553084.
//
// Solidity: function isMilestoneReached(bytes16 causeId, uint8 milestone) view returns(bool)
func (_MilestoneTracker *MilestoneTrackerSession) IsMilestoneReached(causeId [16]byte, milestone uint8) (bool, error) {
	return _MilestoneTracker.Contract.IsMilestoneReached(&_MilestoneTracker.CallOpts, causeId, milestone)
}

// IsMilestoneReached is a free data retrieval call binding the contract method 0x81553084.
//
// Solidity: function isMilestoneReached(bytes16 causeId, uint8 milestone) view returns(bool)
func (_MilestoneTracker *MilestoneTrackerCallerSession) IsMilestoneReached(causeId [16]byte, milestone uint8) (bool, error) {
	return _MilestoneTracker.Contract.IsMilestoneReached(&_MilestoneTracker.CallOpts, causeId, milestone)
}

// RecordDonation is a paid mutator transaction binding the contract method 0x4299071d.
//
// Solidity: function recordDonation(bytes16 causeId, uint256 amount) returns()
func (_MilestoneTracker *MilestoneTrackerTransactor) RecordDonation(opts *bind.TransactOpts, causeId [16]byte, amount *big.Int) (*types.Transaction, error) {
	return _MilestoneTracker.contract.Transact(opts, "recordDonation", causeId, amount)
}

// RecordDonation is a paid mutator transaction binding the contract method 0x4299071d.
//
// Solidity: function recordDonation(bytes16 causeId, uint256 amount) returns()
func (_MilestoneTracker *MilestoneTrackerSession) RecordDonation(causeId [16]byte, amount *big.Int) (*types.Transaction, error) {
	return _MilestoneTracker.Contract.RecordDonation(&_MilestoneTracker.TransactOpts, causeId, amount)
}

// RecordDonation is a paid mutator transaction binding the contract method 0x4299071d.
//
// Solidity: function recordDonation(bytes16 causeId, uint256 amount) returns()
func (_MilestoneTracker *MilestoneTrackerTransactorSession) RecordDonation(causeId [16]byte, amount *big.Int) (*types.Transaction, error) {
	return _MilestoneTracker.Contract.RecordDonation(&_MilestoneTracker.TransactOpts, causeId, amount)
}

// RegisterOrUpdateCause is a paid mutator transaction binding the contract method 0x45192280.
//
// Solidity: function registerOrUpdateCause(bytes16 causeId, uint256 goal, uint256 initialCollected) returns()
func (_MilestoneTracker *MilestoneTrackerTransactor) RegisterOrUpdateCause(opts *bind.TransactOpts, causeId [16]byte, goal *big.Int, initialCollected *big.Int) (*types.Transaction, error) {
	return _MilestoneTracker.contract.Transact(opts, "registerOrUpdateCause", causeId, goal, initialCollected)
}

// RegisterOrUpdateCause is a paid mutator transaction binding the contract method 0x45192280.
//
// Solidity: function registerOrUpdateCause(bytes16 causeId, uint256 goal, uint256 initialCollected) returns()
func (_MilestoneTracker *MilestoneTrackerSession) RegisterOrUpdateCause(causeId [16]byte, goal *big.Int, initialCollected *big.Int) (*types.Transaction, error) {
	return _MilestoneTracker.Contract.RegisterOrUpdateCause(&_MilestoneTracker.TransactOpts, causeId, goal, initialCollected)
}

// RegisterOrUpdateCause is a paid mutator transaction binding the contract method 0x45192280.
//
// Solidity: function registerOrUpdateCause(bytes16 causeId, uint256 goal, uint256 initialCollected) returns()
func (_MilestoneTracker *MilestoneTrackerTransactorSession) RegisterOrUpdateCause(causeId [16]byte, goal *big.Int, initialCollected *big.Int) (*types.Transaction, error) {
	return _MilestoneTracker.Contract.RegisterOrUpdateCause(&_MilestoneTracker.TransactOpts, causeId, goal, initialCollected)
}

// MilestoneTrackerCauseRegisteredIterator is returned from FilterCauseRegistered and is used to iterate over the raw logs and unpacked data for CauseRegistered events raised by the MilestoneTracker contract.
type MilestoneTrackerCauseRegisteredIterator struct {
	Event *MilestoneTrackerCauseRegistered // Event containing the contract specifics and raw log

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
func (it *MilestoneTrackerCauseRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MilestoneTrackerCauseRegistered)
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
		it.Event = new(MilestoneTrackerCauseRegistered)
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
func (it *MilestoneTrackerCauseRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MilestoneTrackerCauseRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MilestoneTrackerCauseRegistered represents a CauseRegistered event raised by the MilestoneTracker contract.
type MilestoneTrackerCauseRegistered struct {
	CauseId [16]byte
	Goal    *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterCauseRegistered is a free log retrieval operation binding the contract event 0x145d9a2cdce81cbe08debbe8acfe1c129851e6c5f083000408ed2dba5045f7b6.
//
// Solidity: event CauseRegistered(bytes16 indexed causeId, uint256 goal)
func (_MilestoneTracker *MilestoneTrackerFilterer) FilterCauseRegistered(opts *bind.FilterOpts, causeId [][16]byte) (*MilestoneTrackerCauseRegisteredIterator, error) {

	var causeIdRule []interface{}
	for _, causeIdItem := range causeId {
		causeIdRule = append(causeIdRule, causeIdItem)
	}

	logs, sub, err := _MilestoneTracker.contract.FilterLogs(opts, "CauseRegistered", causeIdRule)
	if err != nil {
		return nil, err
	}
	return &MilestoneTrackerCauseRegisteredIterator{contract: _MilestoneTracker.contract, event: "CauseRegistered", logs: logs, sub: sub}, nil
}

// WatchCauseRegistered is a free log subscription operation binding the contract event 0x145d9a2cdce81cbe08debbe8acfe1c129851e6c5f083000408ed2dba5045f7b6.
//
// Solidity: event CauseRegistered(bytes16 indexed causeId, uint256 goal)
func (_MilestoneTracker *MilestoneTrackerFilterer) WatchCauseRegistered(opts *bind.WatchOpts, sink chan<- *MilestoneTrackerCauseRegistered, causeId [][16]byte) (event.Subscription, error) {

	var causeIdRule []interface{}
	for _, causeIdItem := range causeId {
		causeIdRule = append(causeIdRule, causeIdItem)
	}

	logs, sub, err := _MilestoneTracker.contract.WatchLogs(opts, "CauseRegistered", causeIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MilestoneTrackerCauseRegistered)
				if err := _MilestoneTracker.contract.UnpackLog(event, "CauseRegistered", log); err != nil {
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

// ParseCauseRegistered is a log parse operation binding the contract event 0x145d9a2cdce81cbe08debbe8acfe1c129851e6c5f083000408ed2dba5045f7b6.
//
// Solidity: event CauseRegistered(bytes16 indexed causeId, uint256 goal)
func (_MilestoneTracker *MilestoneTrackerFilterer) ParseCauseRegistered(log types.Log) (*MilestoneTrackerCauseRegistered, error) {
	event := new(MilestoneTrackerCauseRegistered)
	if err := _MilestoneTracker.contract.UnpackLog(event, "CauseRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MilestoneTrackerDonationRecordedIterator is returned from FilterDonationRecorded and is used to iterate over the raw logs and unpacked data for DonationRecorded events raised by the MilestoneTracker contract.
type MilestoneTrackerDonationRecordedIterator struct {
	Event *MilestoneTrackerDonationRecorded // Event containing the contract specifics and raw log

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
func (it *MilestoneTrackerDonationRecordedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MilestoneTrackerDonationRecorded)
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
		it.Event = new(MilestoneTrackerDonationRecorded)
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
func (it *MilestoneTrackerDonationRecordedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MilestoneTrackerDonationRecordedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MilestoneTrackerDonationRecorded represents a DonationRecorded event raised by the MilestoneTracker contract.
type MilestoneTrackerDonationRecorded struct {
	CauseId      [16]byte
	Amount       *big.Int
	NewCollected *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDonationRecorded is a free log retrieval operation binding the contract event 0x44b218be164f6bac27fc56af0a7b202413c77f42782416d6aa7a003d9d1b8fee.
//
// Solidity: event DonationRecorded(bytes16 indexed causeId, uint256 amount, uint256 newCollected)
func (_MilestoneTracker *MilestoneTrackerFilterer) FilterDonationRecorded(opts *bind.FilterOpts, causeId [][16]byte) (*MilestoneTrackerDonationRecordedIterator, error) {

	var causeIdRule []interface{}
	for _, causeIdItem := range causeId {
		causeIdRule = append(causeIdRule, causeIdItem)
	}

	logs, sub, err := _MilestoneTracker.contract.FilterLogs(opts, "DonationRecorded", causeIdRule)
	if err != nil {
		return nil, err
	}
	return &MilestoneTrackerDonationRecordedIterator{contract: _MilestoneTracker.contract, event: "DonationRecorded", logs: logs, sub: sub}, nil
}

// WatchDonationRecorded is a free log subscription operation binding the contract event 0x44b218be164f6bac27fc56af0a7b202413c77f42782416d6aa7a003d9d1b8fee.
//
// Solidity: event DonationRecorded(bytes16 indexed causeId, uint256 amount, uint256 newCollected)
func (_MilestoneTracker *MilestoneTrackerFilterer) WatchDonationRecorded(opts *bind.WatchOpts, sink chan<- *MilestoneTrackerDonationRecorded, causeId [][16]byte) (event.Subscription, error) {

	var causeIdRule []interface{}
	for _, causeIdItem := range causeId {
		causeIdRule = append(causeIdRule, causeIdItem)
	}

	logs, sub, err := _MilestoneTracker.contract.WatchLogs(opts, "DonationRecorded", causeIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MilestoneTrackerDonationRecorded)
				if err := _MilestoneTracker.contract.UnpackLog(event, "DonationRecorded", log); err != nil {
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

// ParseDonationRecorded is a log parse operation binding the contract event 0x44b218be164f6bac27fc56af0a7b202413c77f42782416d6aa7a003d9d1b8fee.
//
// Solidity: event DonationRecorded(bytes16 indexed causeId, uint256 amount, uint256 newCollected)
func (_MilestoneTracker *MilestoneTrackerFilterer) ParseDonationRecorded(log types.Log) (*MilestoneTrackerDonationRecorded, error) {
	event := new(MilestoneTrackerDonationRecorded)
	if err := _MilestoneTracker.contract.UnpackLog(event, "DonationRecorded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MilestoneTrackerMilestoneReachedIterator is returned from FilterMilestoneReached and is used to iterate over the raw logs and unpacked data for MilestoneReached events raised by the MilestoneTracker contract.
type MilestoneTrackerMilestoneReachedIterator struct {
	Event *MilestoneTrackerMilestoneReached // Event containing the contract specifics and raw log

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
func (it *MilestoneTrackerMilestoneReachedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MilestoneTrackerMilestoneReached)
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
		it.Event = new(MilestoneTrackerMilestoneReached)
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
func (it *MilestoneTrackerMilestoneReachedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MilestoneTrackerMilestoneReachedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MilestoneTrackerMilestoneReached represents a MilestoneReached event raised by the MilestoneTracker contract.
type MilestoneTrackerMilestoneReached struct {
	CauseId         [16]byte
	Milestone       uint8
	AmountToDiburse *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterMilestoneReached is a free log retrieval operation binding the contract event 0x2ea29736068cc2fb33847eb950da5cf65c733582dce27abd14f43115c4d9ed67.
//
// Solidity: event MilestoneReached(bytes16 indexed causeId, uint8 milestone, uint256 amountToDiburse)
func (_MilestoneTracker *MilestoneTrackerFilterer) FilterMilestoneReached(opts *bind.FilterOpts, causeId [][16]byte) (*MilestoneTrackerMilestoneReachedIterator, error) {

	var causeIdRule []interface{}
	for _, causeIdItem := range causeId {
		causeIdRule = append(causeIdRule, causeIdItem)
	}

	logs, sub, err := _MilestoneTracker.contract.FilterLogs(opts, "MilestoneReached", causeIdRule)
	if err != nil {
		return nil, err
	}
	return &MilestoneTrackerMilestoneReachedIterator{contract: _MilestoneTracker.contract, event: "MilestoneReached", logs: logs, sub: sub}, nil
}

// WatchMilestoneReached is a free log subscription operation binding the contract event 0x2ea29736068cc2fb33847eb950da5cf65c733582dce27abd14f43115c4d9ed67.
//
// Solidity: event MilestoneReached(bytes16 indexed causeId, uint8 milestone, uint256 amountToDiburse)
func (_MilestoneTracker *MilestoneTrackerFilterer) WatchMilestoneReached(opts *bind.WatchOpts, sink chan<- *MilestoneTrackerMilestoneReached, causeId [][16]byte) (event.Subscription, error) {

	var causeIdRule []interface{}
	for _, causeIdItem := range causeId {
		causeIdRule = append(causeIdRule, causeIdItem)
	}

	logs, sub, err := _MilestoneTracker.contract.WatchLogs(opts, "MilestoneReached", causeIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MilestoneTrackerMilestoneReached)
				if err := _MilestoneTracker.contract.UnpackLog(event, "MilestoneReached", log); err != nil {
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

// ParseMilestoneReached is a log parse operation binding the contract event 0x2ea29736068cc2fb33847eb950da5cf65c733582dce27abd14f43115c4d9ed67.
//
// Solidity: event MilestoneReached(bytes16 indexed causeId, uint8 milestone, uint256 amountToDiburse)
func (_MilestoneTracker *MilestoneTrackerFilterer) ParseMilestoneReached(log types.Log) (*MilestoneTrackerMilestoneReached, error) {
	event := new(MilestoneTrackerMilestoneReached)
	if err := _MilestoneTracker.contract.UnpackLog(event, "MilestoneReached", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
