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

// DonationLedgerDonation is an auto generated low-level Go binding around an user-defined struct.
type DonationLedgerDonation struct {
	CauseId    [16]byte
	DonorId    [16]byte
	Amount     *big.Int
	Timestamp  *big.Int
	PaymentRef string
}

// DonationLedgerMetaData contains all meta data concerning the DonationLedger contract.
var DonationLedgerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes16\",\"name\":\"\",\"type\":\"bytes16\"}],\"name\":\"donations\",\"outputs\":[{\"internalType\":\"bytes16\",\"name\":\"causeId\",\"type\":\"bytes16\"},{\"internalType\":\"bytes16\",\"name\":\"donorId\",\"type\":\"bytes16\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"paymentRef\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes16\",\"name\":\"\",\"type\":\"bytes16\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"donationsByCause\",\"outputs\":[{\"internalType\":\"bytes16\",\"name\":\"\",\"type\":\"bytes16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes16\",\"name\":\"donationId\",\"type\":\"bytes16\"}],\"name\":\"getDonation\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes16\",\"name\":\"causeId\",\"type\":\"bytes16\"},{\"internalType\":\"bytes16\",\"name\":\"donorId\",\"type\":\"bytes16\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"paymentRef\",\"type\":\"string\"}],\"internalType\":\"structDonationLedger.Donation\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes16\",\"name\":\"causeId\",\"type\":\"bytes16\"}],\"name\":\"getDonationsByCause\",\"outputs\":[{\"internalType\":\"bytes16[]\",\"name\":\"\",\"type\":\"bytes16[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes16\",\"name\":\"donationId\",\"type\":\"bytes16\"},{\"internalType\":\"bytes16\",\"name\":\"causeId\",\"type\":\"bytes16\"},{\"internalType\":\"bytes16\",\"name\":\"donorId\",\"type\":\"bytes16\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"paymentRef\",\"type\":\"string\"}],\"name\":\"recordDonation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// DonationLedgerABI is the input ABI used to generate the binding from.
// Deprecated: Use DonationLedgerMetaData.ABI instead.
var DonationLedgerABI = DonationLedgerMetaData.ABI

// DonationLedger is an auto generated Go binding around an Ethereum contract.
type DonationLedger struct {
	DonationLedgerCaller     // Read-only binding to the contract
	DonationLedgerTransactor // Write-only binding to the contract
	DonationLedgerFilterer   // Log filterer for contract events
}

// DonationLedgerCaller is an auto generated read-only Go binding around an Ethereum contract.
type DonationLedgerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DonationLedgerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DonationLedgerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DonationLedgerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DonationLedgerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DonationLedgerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DonationLedgerSession struct {
	Contract     *DonationLedger   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DonationLedgerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DonationLedgerCallerSession struct {
	Contract *DonationLedgerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// DonationLedgerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DonationLedgerTransactorSession struct {
	Contract     *DonationLedgerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// DonationLedgerRaw is an auto generated low-level Go binding around an Ethereum contract.
type DonationLedgerRaw struct {
	Contract *DonationLedger // Generic contract binding to access the raw methods on
}

// DonationLedgerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DonationLedgerCallerRaw struct {
	Contract *DonationLedgerCaller // Generic read-only contract binding to access the raw methods on
}

// DonationLedgerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DonationLedgerTransactorRaw struct {
	Contract *DonationLedgerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDonationLedger creates a new instance of DonationLedger, bound to a specific deployed contract.
func NewDonationLedger(address common.Address, backend bind.ContractBackend) (*DonationLedger, error) {
	contract, err := bindDonationLedger(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DonationLedger{DonationLedgerCaller: DonationLedgerCaller{contract: contract}, DonationLedgerTransactor: DonationLedgerTransactor{contract: contract}, DonationLedgerFilterer: DonationLedgerFilterer{contract: contract}}, nil
}

// NewDonationLedgerCaller creates a new read-only instance of DonationLedger, bound to a specific deployed contract.
func NewDonationLedgerCaller(address common.Address, caller bind.ContractCaller) (*DonationLedgerCaller, error) {
	contract, err := bindDonationLedger(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DonationLedgerCaller{contract: contract}, nil
}

// NewDonationLedgerTransactor creates a new write-only instance of DonationLedger, bound to a specific deployed contract.
func NewDonationLedgerTransactor(address common.Address, transactor bind.ContractTransactor) (*DonationLedgerTransactor, error) {
	contract, err := bindDonationLedger(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DonationLedgerTransactor{contract: contract}, nil
}

// NewDonationLedgerFilterer creates a new log filterer instance of DonationLedger, bound to a specific deployed contract.
func NewDonationLedgerFilterer(address common.Address, filterer bind.ContractFilterer) (*DonationLedgerFilterer, error) {
	contract, err := bindDonationLedger(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DonationLedgerFilterer{contract: contract}, nil
}

// bindDonationLedger binds a generic wrapper to an already deployed contract.
func bindDonationLedger(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DonationLedgerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DonationLedger *DonationLedgerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DonationLedger.Contract.DonationLedgerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DonationLedger *DonationLedgerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DonationLedger.Contract.DonationLedgerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DonationLedger *DonationLedgerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DonationLedger.Contract.DonationLedgerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DonationLedger *DonationLedgerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DonationLedger.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DonationLedger *DonationLedgerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DonationLedger.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DonationLedger *DonationLedgerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DonationLedger.Contract.contract.Transact(opts, method, params...)
}

// Donations is a free data retrieval call binding the contract method 0x618e9ee2.
//
// Solidity: function donations(bytes16 ) view returns(bytes16 causeId, bytes16 donorId, uint256 amount, uint256 timestamp, string paymentRef)
func (_DonationLedger *DonationLedgerCaller) Donations(opts *bind.CallOpts, arg0 [16]byte) (struct {
	CauseId    [16]byte
	DonorId    [16]byte
	Amount     *big.Int
	Timestamp  *big.Int
	PaymentRef string
}, error) {
	var out []interface{}
	err := _DonationLedger.contract.Call(opts, &out, "donations", arg0)

	outstruct := new(struct {
		CauseId    [16]byte
		DonorId    [16]byte
		Amount     *big.Int
		Timestamp  *big.Int
		PaymentRef string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.CauseId = *abi.ConvertType(out[0], new([16]byte)).(*[16]byte)
	outstruct.DonorId = *abi.ConvertType(out[1], new([16]byte)).(*[16]byte)
	outstruct.Amount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Timestamp = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.PaymentRef = *abi.ConvertType(out[4], new(string)).(*string)

	return *outstruct, err

}

// Donations is a free data retrieval call binding the contract method 0x618e9ee2.
//
// Solidity: function donations(bytes16 ) view returns(bytes16 causeId, bytes16 donorId, uint256 amount, uint256 timestamp, string paymentRef)
func (_DonationLedger *DonationLedgerSession) Donations(arg0 [16]byte) (struct {
	CauseId    [16]byte
	DonorId    [16]byte
	Amount     *big.Int
	Timestamp  *big.Int
	PaymentRef string
}, error) {
	return _DonationLedger.Contract.Donations(&_DonationLedger.CallOpts, arg0)
}

// Donations is a free data retrieval call binding the contract method 0x618e9ee2.
//
// Solidity: function donations(bytes16 ) view returns(bytes16 causeId, bytes16 donorId, uint256 amount, uint256 timestamp, string paymentRef)
func (_DonationLedger *DonationLedgerCallerSession) Donations(arg0 [16]byte) (struct {
	CauseId    [16]byte
	DonorId    [16]byte
	Amount     *big.Int
	Timestamp  *big.Int
	PaymentRef string
}, error) {
	return _DonationLedger.Contract.Donations(&_DonationLedger.CallOpts, arg0)
}

// DonationsByCause is a free data retrieval call binding the contract method 0x31130faf.
//
// Solidity: function donationsByCause(bytes16 , uint256 ) view returns(bytes16)
func (_DonationLedger *DonationLedgerCaller) DonationsByCause(opts *bind.CallOpts, arg0 [16]byte, arg1 *big.Int) ([16]byte, error) {
	var out []interface{}
	err := _DonationLedger.contract.Call(opts, &out, "donationsByCause", arg0, arg1)

	if err != nil {
		return *new([16]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([16]byte)).(*[16]byte)

	return out0, err

}

// DonationsByCause is a free data retrieval call binding the contract method 0x31130faf.
//
// Solidity: function donationsByCause(bytes16 , uint256 ) view returns(bytes16)
func (_DonationLedger *DonationLedgerSession) DonationsByCause(arg0 [16]byte, arg1 *big.Int) ([16]byte, error) {
	return _DonationLedger.Contract.DonationsByCause(&_DonationLedger.CallOpts, arg0, arg1)
}

// DonationsByCause is a free data retrieval call binding the contract method 0x31130faf.
//
// Solidity: function donationsByCause(bytes16 , uint256 ) view returns(bytes16)
func (_DonationLedger *DonationLedgerCallerSession) DonationsByCause(arg0 [16]byte, arg1 *big.Int) ([16]byte, error) {
	return _DonationLedger.Contract.DonationsByCause(&_DonationLedger.CallOpts, arg0, arg1)
}

// GetDonation is a free data retrieval call binding the contract method 0x829d4487.
//
// Solidity: function getDonation(bytes16 donationId) view returns((bytes16,bytes16,uint256,uint256,string))
func (_DonationLedger *DonationLedgerCaller) GetDonation(opts *bind.CallOpts, donationId [16]byte) (DonationLedgerDonation, error) {
	var out []interface{}
	err := _DonationLedger.contract.Call(opts, &out, "getDonation", donationId)

	if err != nil {
		return *new(DonationLedgerDonation), err
	}

	out0 := *abi.ConvertType(out[0], new(DonationLedgerDonation)).(*DonationLedgerDonation)

	return out0, err

}

// GetDonation is a free data retrieval call binding the contract method 0x829d4487.
//
// Solidity: function getDonation(bytes16 donationId) view returns((bytes16,bytes16,uint256,uint256,string))
func (_DonationLedger *DonationLedgerSession) GetDonation(donationId [16]byte) (DonationLedgerDonation, error) {
	return _DonationLedger.Contract.GetDonation(&_DonationLedger.CallOpts, donationId)
}

// GetDonation is a free data retrieval call binding the contract method 0x829d4487.
//
// Solidity: function getDonation(bytes16 donationId) view returns((bytes16,bytes16,uint256,uint256,string))
func (_DonationLedger *DonationLedgerCallerSession) GetDonation(donationId [16]byte) (DonationLedgerDonation, error) {
	return _DonationLedger.Contract.GetDonation(&_DonationLedger.CallOpts, donationId)
}

// GetDonationsByCause is a free data retrieval call binding the contract method 0x4f7c89f8.
//
// Solidity: function getDonationsByCause(bytes16 causeId) view returns(bytes16[])
func (_DonationLedger *DonationLedgerCaller) GetDonationsByCause(opts *bind.CallOpts, causeId [16]byte) ([][16]byte, error) {
	var out []interface{}
	err := _DonationLedger.contract.Call(opts, &out, "getDonationsByCause", causeId)

	if err != nil {
		return *new([][16]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][16]byte)).(*[][16]byte)

	return out0, err

}

// GetDonationsByCause is a free data retrieval call binding the contract method 0x4f7c89f8.
//
// Solidity: function getDonationsByCause(bytes16 causeId) view returns(bytes16[])
func (_DonationLedger *DonationLedgerSession) GetDonationsByCause(causeId [16]byte) ([][16]byte, error) {
	return _DonationLedger.Contract.GetDonationsByCause(&_DonationLedger.CallOpts, causeId)
}

// GetDonationsByCause is a free data retrieval call binding the contract method 0x4f7c89f8.
//
// Solidity: function getDonationsByCause(bytes16 causeId) view returns(bytes16[])
func (_DonationLedger *DonationLedgerCallerSession) GetDonationsByCause(causeId [16]byte) ([][16]byte, error) {
	return _DonationLedger.Contract.GetDonationsByCause(&_DonationLedger.CallOpts, causeId)
}

// RecordDonation is a paid mutator transaction binding the contract method 0x77c61acd.
//
// Solidity: function recordDonation(bytes16 donationId, bytes16 causeId, bytes16 donorId, uint256 amount, string paymentRef) returns()
func (_DonationLedger *DonationLedgerTransactor) RecordDonation(opts *bind.TransactOpts, donationId [16]byte, causeId [16]byte, donorId [16]byte, amount *big.Int, paymentRef string) (*types.Transaction, error) {
	return _DonationLedger.contract.Transact(opts, "recordDonation", donationId, causeId, donorId, amount, paymentRef)
}

// RecordDonation is a paid mutator transaction binding the contract method 0x77c61acd.
//
// Solidity: function recordDonation(bytes16 donationId, bytes16 causeId, bytes16 donorId, uint256 amount, string paymentRef) returns()
func (_DonationLedger *DonationLedgerSession) RecordDonation(donationId [16]byte, causeId [16]byte, donorId [16]byte, amount *big.Int, paymentRef string) (*types.Transaction, error) {
	return _DonationLedger.Contract.RecordDonation(&_DonationLedger.TransactOpts, donationId, causeId, donorId, amount, paymentRef)
}

// RecordDonation is a paid mutator transaction binding the contract method 0x77c61acd.
//
// Solidity: function recordDonation(bytes16 donationId, bytes16 causeId, bytes16 donorId, uint256 amount, string paymentRef) returns()
func (_DonationLedger *DonationLedgerTransactorSession) RecordDonation(donationId [16]byte, causeId [16]byte, donorId [16]byte, amount *big.Int, paymentRef string) (*types.Transaction, error) {
	return _DonationLedger.Contract.RecordDonation(&_DonationLedger.TransactOpts, donationId, causeId, donorId, amount, paymentRef)
}
