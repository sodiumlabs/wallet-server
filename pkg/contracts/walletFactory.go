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

// WalletFactoryMetaData contains all meta data concerning the WalletFactory contract.
var WalletFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_walletImplementation\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_guardianStorage\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_refundAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_manager\",\"type\":\"address\"}],\"name\":\"ManagerAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_manager\",\"type\":\"address\"}],\"name\":\"ManagerRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"OwnerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"RefundAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"guardian\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"refundToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"refundAmount\",\"type\":\"uint256\"}],\"name\":\"WalletCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_manager\",\"type\":\"address\"}],\"name\":\"addManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"guardianStorage\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"managers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"refundAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"walletImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"revokeManager\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"_modules\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"_guardian\",\"type\":\"address\"},{\"internalType\":\"bytes20\",\"name\":\"_salt\",\"type\":\"bytes20\"},{\"internalType\":\"uint256\",\"name\":\"_refundAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_refundToken\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_ownerSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_managerSignature\",\"type\":\"bytes\"}],\"name\":\"createCounterfactualWallet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_wallet\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_modules\",\"type\":\"address[]\"},{\"internalType\":\"bytes20\",\"name\":\"_salt\",\"type\":\"bytes20\"}],\"name\":\"getAddressForCounterfactualWallet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_wallet\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_refundAddress\",\"type\":\"address\"}],\"name\":\"changeRefundAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractBaseWallet\",\"name\":\"_wallet\",\"type\":\"address\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
}

// WalletFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use WalletFactoryMetaData.ABI instead.
var WalletFactoryABI = WalletFactoryMetaData.ABI

// WalletFactory is an auto generated Go binding around an Ethereum contract.
type WalletFactory struct {
	WalletFactoryCaller     // Read-only binding to the contract
	WalletFactoryTransactor // Write-only binding to the contract
	WalletFactoryFilterer   // Log filterer for contract events
}

// WalletFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type WalletFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WalletFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WalletFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WalletFactorySession struct {
	Contract     *WalletFactory    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WalletFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WalletFactoryCallerSession struct {
	Contract *WalletFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// WalletFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WalletFactoryTransactorSession struct {
	Contract     *WalletFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// WalletFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type WalletFactoryRaw struct {
	Contract *WalletFactory // Generic contract binding to access the raw methods on
}

// WalletFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WalletFactoryCallerRaw struct {
	Contract *WalletFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// WalletFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WalletFactoryTransactorRaw struct {
	Contract *WalletFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWalletFactory creates a new instance of WalletFactory, bound to a specific deployed contract.
func NewWalletFactory(address common.Address, backend bind.ContractBackend) (*WalletFactory, error) {
	contract, err := bindWalletFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WalletFactory{WalletFactoryCaller: WalletFactoryCaller{contract: contract}, WalletFactoryTransactor: WalletFactoryTransactor{contract: contract}, WalletFactoryFilterer: WalletFactoryFilterer{contract: contract}}, nil
}

// NewWalletFactoryCaller creates a new read-only instance of WalletFactory, bound to a specific deployed contract.
func NewWalletFactoryCaller(address common.Address, caller bind.ContractCaller) (*WalletFactoryCaller, error) {
	contract, err := bindWalletFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WalletFactoryCaller{contract: contract}, nil
}

// NewWalletFactoryTransactor creates a new write-only instance of WalletFactory, bound to a specific deployed contract.
func NewWalletFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*WalletFactoryTransactor, error) {
	contract, err := bindWalletFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WalletFactoryTransactor{contract: contract}, nil
}

// NewWalletFactoryFilterer creates a new log filterer instance of WalletFactory, bound to a specific deployed contract.
func NewWalletFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*WalletFactoryFilterer, error) {
	contract, err := bindWalletFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WalletFactoryFilterer{contract: contract}, nil
}

// bindWalletFactory binds a generic wrapper to an already deployed contract.
func bindWalletFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WalletFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WalletFactory *WalletFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WalletFactory.Contract.WalletFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WalletFactory *WalletFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WalletFactory.Contract.WalletFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WalletFactory *WalletFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WalletFactory.Contract.WalletFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WalletFactory *WalletFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WalletFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WalletFactory *WalletFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WalletFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WalletFactory *WalletFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WalletFactory.Contract.contract.Transact(opts, method, params...)
}

// GetAddressForCounterfactualWallet is a free data retrieval call binding the contract method 0x64b9c7a1.
//
// Solidity: function getAddressForCounterfactualWallet(address[] _modules, bytes20 _salt) view returns(address _wallet)
func (_WalletFactory *WalletFactoryCaller) GetAddressForCounterfactualWallet(opts *bind.CallOpts, _modules []common.Address, _salt [20]byte) (common.Address, error) {
	var out []interface{}
	err := _WalletFactory.contract.Call(opts, &out, "getAddressForCounterfactualWallet", _modules, _salt)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddressForCounterfactualWallet is a free data retrieval call binding the contract method 0x64b9c7a1.
//
// Solidity: function getAddressForCounterfactualWallet(address[] _modules, bytes20 _salt) view returns(address _wallet)
func (_WalletFactory *WalletFactorySession) GetAddressForCounterfactualWallet(_modules []common.Address, _salt [20]byte) (common.Address, error) {
	return _WalletFactory.Contract.GetAddressForCounterfactualWallet(&_WalletFactory.CallOpts, _modules, _salt)
}

// GetAddressForCounterfactualWallet is a free data retrieval call binding the contract method 0x64b9c7a1.
//
// Solidity: function getAddressForCounterfactualWallet(address[] _modules, bytes20 _salt) view returns(address _wallet)
func (_WalletFactory *WalletFactoryCallerSession) GetAddressForCounterfactualWallet(_modules []common.Address, _salt [20]byte) (common.Address, error) {
	return _WalletFactory.Contract.GetAddressForCounterfactualWallet(&_WalletFactory.CallOpts, _modules, _salt)
}

// GuardianStorage is a free data retrieval call binding the contract method 0xd89784fc.
//
// Solidity: function guardianStorage() view returns(address)
func (_WalletFactory *WalletFactoryCaller) GuardianStorage(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WalletFactory.contract.Call(opts, &out, "guardianStorage")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GuardianStorage is a free data retrieval call binding the contract method 0xd89784fc.
//
// Solidity: function guardianStorage() view returns(address)
func (_WalletFactory *WalletFactorySession) GuardianStorage() (common.Address, error) {
	return _WalletFactory.Contract.GuardianStorage(&_WalletFactory.CallOpts)
}

// GuardianStorage is a free data retrieval call binding the contract method 0xd89784fc.
//
// Solidity: function guardianStorage() view returns(address)
func (_WalletFactory *WalletFactoryCallerSession) GuardianStorage() (common.Address, error) {
	return _WalletFactory.Contract.GuardianStorage(&_WalletFactory.CallOpts)
}

// Init is a free data retrieval call binding the contract method 0x19ab453c.
//
// Solidity: function init(address _wallet) pure returns()
func (_WalletFactory *WalletFactoryCaller) Init(opts *bind.CallOpts, _wallet common.Address) error {
	var out []interface{}
	err := _WalletFactory.contract.Call(opts, &out, "init", _wallet)

	if err != nil {
		return err
	}

	return err

}

// Init is a free data retrieval call binding the contract method 0x19ab453c.
//
// Solidity: function init(address _wallet) pure returns()
func (_WalletFactory *WalletFactorySession) Init(_wallet common.Address) error {
	return _WalletFactory.Contract.Init(&_WalletFactory.CallOpts, _wallet)
}

// Init is a free data retrieval call binding the contract method 0x19ab453c.
//
// Solidity: function init(address _wallet) pure returns()
func (_WalletFactory *WalletFactoryCallerSession) Init(_wallet common.Address) error {
	return _WalletFactory.Contract.Init(&_WalletFactory.CallOpts, _wallet)
}

// Managers is a free data retrieval call binding the contract method 0xfdff9b4d.
//
// Solidity: function managers(address ) view returns(bool)
func (_WalletFactory *WalletFactoryCaller) Managers(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _WalletFactory.contract.Call(opts, &out, "managers", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Managers is a free data retrieval call binding the contract method 0xfdff9b4d.
//
// Solidity: function managers(address ) view returns(bool)
func (_WalletFactory *WalletFactorySession) Managers(arg0 common.Address) (bool, error) {
	return _WalletFactory.Contract.Managers(&_WalletFactory.CallOpts, arg0)
}

// Managers is a free data retrieval call binding the contract method 0xfdff9b4d.
//
// Solidity: function managers(address ) view returns(bool)
func (_WalletFactory *WalletFactoryCallerSession) Managers(arg0 common.Address) (bool, error) {
	return _WalletFactory.Contract.Managers(&_WalletFactory.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WalletFactory *WalletFactoryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WalletFactory.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WalletFactory *WalletFactorySession) Owner() (common.Address, error) {
	return _WalletFactory.Contract.Owner(&_WalletFactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WalletFactory *WalletFactoryCallerSession) Owner() (common.Address, error) {
	return _WalletFactory.Contract.Owner(&_WalletFactory.CallOpts)
}

// RefundAddress is a free data retrieval call binding the contract method 0x0cb61f6c.
//
// Solidity: function refundAddress() view returns(address)
func (_WalletFactory *WalletFactoryCaller) RefundAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WalletFactory.contract.Call(opts, &out, "refundAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RefundAddress is a free data retrieval call binding the contract method 0x0cb61f6c.
//
// Solidity: function refundAddress() view returns(address)
func (_WalletFactory *WalletFactorySession) RefundAddress() (common.Address, error) {
	return _WalletFactory.Contract.RefundAddress(&_WalletFactory.CallOpts)
}

// RefundAddress is a free data retrieval call binding the contract method 0x0cb61f6c.
//
// Solidity: function refundAddress() view returns(address)
func (_WalletFactory *WalletFactoryCallerSession) RefundAddress() (common.Address, error) {
	return _WalletFactory.Contract.RefundAddress(&_WalletFactory.CallOpts)
}

// RevokeManager is a free data retrieval call binding the contract method 0x377e32e6.
//
// Solidity: function revokeManager(address ) pure returns()
func (_WalletFactory *WalletFactoryCaller) RevokeManager(opts *bind.CallOpts, arg0 common.Address) error {
	var out []interface{}
	err := _WalletFactory.contract.Call(opts, &out, "revokeManager", arg0)

	if err != nil {
		return err
	}

	return err

}

// RevokeManager is a free data retrieval call binding the contract method 0x377e32e6.
//
// Solidity: function revokeManager(address ) pure returns()
func (_WalletFactory *WalletFactorySession) RevokeManager(arg0 common.Address) error {
	return _WalletFactory.Contract.RevokeManager(&_WalletFactory.CallOpts, arg0)
}

// RevokeManager is a free data retrieval call binding the contract method 0x377e32e6.
//
// Solidity: function revokeManager(address ) pure returns()
func (_WalletFactory *WalletFactoryCallerSession) RevokeManager(arg0 common.Address) error {
	return _WalletFactory.Contract.RevokeManager(&_WalletFactory.CallOpts, arg0)
}

// WalletImplementation is a free data retrieval call binding the contract method 0x8117abc1.
//
// Solidity: function walletImplementation() view returns(address)
func (_WalletFactory *WalletFactoryCaller) WalletImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WalletFactory.contract.Call(opts, &out, "walletImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WalletImplementation is a free data retrieval call binding the contract method 0x8117abc1.
//
// Solidity: function walletImplementation() view returns(address)
func (_WalletFactory *WalletFactorySession) WalletImplementation() (common.Address, error) {
	return _WalletFactory.Contract.WalletImplementation(&_WalletFactory.CallOpts)
}

// WalletImplementation is a free data retrieval call binding the contract method 0x8117abc1.
//
// Solidity: function walletImplementation() view returns(address)
func (_WalletFactory *WalletFactoryCallerSession) WalletImplementation() (common.Address, error) {
	return _WalletFactory.Contract.WalletImplementation(&_WalletFactory.CallOpts)
}

// AddManager is a paid mutator transaction binding the contract method 0x2d06177a.
//
// Solidity: function addManager(address _manager) returns()
func (_WalletFactory *WalletFactoryTransactor) AddManager(opts *bind.TransactOpts, _manager common.Address) (*types.Transaction, error) {
	return _WalletFactory.contract.Transact(opts, "addManager", _manager)
}

// AddManager is a paid mutator transaction binding the contract method 0x2d06177a.
//
// Solidity: function addManager(address _manager) returns()
func (_WalletFactory *WalletFactorySession) AddManager(_manager common.Address) (*types.Transaction, error) {
	return _WalletFactory.Contract.AddManager(&_WalletFactory.TransactOpts, _manager)
}

// AddManager is a paid mutator transaction binding the contract method 0x2d06177a.
//
// Solidity: function addManager(address _manager) returns()
func (_WalletFactory *WalletFactoryTransactorSession) AddManager(_manager common.Address) (*types.Transaction, error) {
	return _WalletFactory.Contract.AddManager(&_WalletFactory.TransactOpts, _manager)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address _newOwner) returns()
func (_WalletFactory *WalletFactoryTransactor) ChangeOwner(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _WalletFactory.contract.Transact(opts, "changeOwner", _newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address _newOwner) returns()
func (_WalletFactory *WalletFactorySession) ChangeOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _WalletFactory.Contract.ChangeOwner(&_WalletFactory.TransactOpts, _newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address _newOwner) returns()
func (_WalletFactory *WalletFactoryTransactorSession) ChangeOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _WalletFactory.Contract.ChangeOwner(&_WalletFactory.TransactOpts, _newOwner)
}

// ChangeRefundAddress is a paid mutator transaction binding the contract method 0x2c828525.
//
// Solidity: function changeRefundAddress(address _refundAddress) returns()
func (_WalletFactory *WalletFactoryTransactor) ChangeRefundAddress(opts *bind.TransactOpts, _refundAddress common.Address) (*types.Transaction, error) {
	return _WalletFactory.contract.Transact(opts, "changeRefundAddress", _refundAddress)
}

// ChangeRefundAddress is a paid mutator transaction binding the contract method 0x2c828525.
//
// Solidity: function changeRefundAddress(address _refundAddress) returns()
func (_WalletFactory *WalletFactorySession) ChangeRefundAddress(_refundAddress common.Address) (*types.Transaction, error) {
	return _WalletFactory.Contract.ChangeRefundAddress(&_WalletFactory.TransactOpts, _refundAddress)
}

// ChangeRefundAddress is a paid mutator transaction binding the contract method 0x2c828525.
//
// Solidity: function changeRefundAddress(address _refundAddress) returns()
func (_WalletFactory *WalletFactoryTransactorSession) ChangeRefundAddress(_refundAddress common.Address) (*types.Transaction, error) {
	return _WalletFactory.Contract.ChangeRefundAddress(&_WalletFactory.TransactOpts, _refundAddress)
}

// CreateCounterfactualWallet is a paid mutator transaction binding the contract method 0xa80dd9df.
//
// Solidity: function createCounterfactualWallet(address _owner, address[] _modules, address _guardian, bytes20 _salt, uint256 _refundAmount, address _refundToken, bytes _ownerSignature, bytes _managerSignature) returns(address _wallet)
func (_WalletFactory *WalletFactoryTransactor) CreateCounterfactualWallet(opts *bind.TransactOpts, _owner common.Address, _modules []common.Address, _guardian common.Address, _salt [20]byte, _refundAmount *big.Int, _refundToken common.Address, _ownerSignature []byte, _managerSignature []byte) (*types.Transaction, error) {
	return _WalletFactory.contract.Transact(opts, "createCounterfactualWallet", _owner, _modules, _guardian, _salt, _refundAmount, _refundToken, _ownerSignature, _managerSignature)
}

// CreateCounterfactualWallet is a paid mutator transaction binding the contract method 0xa80dd9df.
//
// Solidity: function createCounterfactualWallet(address _owner, address[] _modules, address _guardian, bytes20 _salt, uint256 _refundAmount, address _refundToken, bytes _ownerSignature, bytes _managerSignature) returns(address _wallet)
func (_WalletFactory *WalletFactorySession) CreateCounterfactualWallet(_owner common.Address, _modules []common.Address, _guardian common.Address, _salt [20]byte, _refundAmount *big.Int, _refundToken common.Address, _ownerSignature []byte, _managerSignature []byte) (*types.Transaction, error) {
	return _WalletFactory.Contract.CreateCounterfactualWallet(&_WalletFactory.TransactOpts, _owner, _modules, _guardian, _salt, _refundAmount, _refundToken, _ownerSignature, _managerSignature)
}

// CreateCounterfactualWallet is a paid mutator transaction binding the contract method 0xa80dd9df.
//
// Solidity: function createCounterfactualWallet(address _owner, address[] _modules, address _guardian, bytes20 _salt, uint256 _refundAmount, address _refundToken, bytes _ownerSignature, bytes _managerSignature) returns(address _wallet)
func (_WalletFactory *WalletFactoryTransactorSession) CreateCounterfactualWallet(_owner common.Address, _modules []common.Address, _guardian common.Address, _salt [20]byte, _refundAmount *big.Int, _refundToken common.Address, _ownerSignature []byte, _managerSignature []byte) (*types.Transaction, error) {
	return _WalletFactory.Contract.CreateCounterfactualWallet(&_WalletFactory.TransactOpts, _owner, _modules, _guardian, _salt, _refundAmount, _refundToken, _ownerSignature, _managerSignature)
}

// WalletFactoryManagerAddedIterator is returned from FilterManagerAdded and is used to iterate over the raw logs and unpacked data for ManagerAdded events raised by the WalletFactory contract.
type WalletFactoryManagerAddedIterator struct {
	Event *WalletFactoryManagerAdded // Event containing the contract specifics and raw log

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
func (it *WalletFactoryManagerAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletFactoryManagerAdded)
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
		it.Event = new(WalletFactoryManagerAdded)
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
func (it *WalletFactoryManagerAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletFactoryManagerAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletFactoryManagerAdded represents a ManagerAdded event raised by the WalletFactory contract.
type WalletFactoryManagerAdded struct {
	Manager common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterManagerAdded is a free log retrieval operation binding the contract event 0x3b4a40cccf2058c593542587329dd385be4f0b588db5471fbd9598e56dd7093a.
//
// Solidity: event ManagerAdded(address indexed _manager)
func (_WalletFactory *WalletFactoryFilterer) FilterManagerAdded(opts *bind.FilterOpts, _manager []common.Address) (*WalletFactoryManagerAddedIterator, error) {

	var _managerRule []interface{}
	for _, _managerItem := range _manager {
		_managerRule = append(_managerRule, _managerItem)
	}

	logs, sub, err := _WalletFactory.contract.FilterLogs(opts, "ManagerAdded", _managerRule)
	if err != nil {
		return nil, err
	}
	return &WalletFactoryManagerAddedIterator{contract: _WalletFactory.contract, event: "ManagerAdded", logs: logs, sub: sub}, nil
}

// WatchManagerAdded is a free log subscription operation binding the contract event 0x3b4a40cccf2058c593542587329dd385be4f0b588db5471fbd9598e56dd7093a.
//
// Solidity: event ManagerAdded(address indexed _manager)
func (_WalletFactory *WalletFactoryFilterer) WatchManagerAdded(opts *bind.WatchOpts, sink chan<- *WalletFactoryManagerAdded, _manager []common.Address) (event.Subscription, error) {

	var _managerRule []interface{}
	for _, _managerItem := range _manager {
		_managerRule = append(_managerRule, _managerItem)
	}

	logs, sub, err := _WalletFactory.contract.WatchLogs(opts, "ManagerAdded", _managerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletFactoryManagerAdded)
				if err := _WalletFactory.contract.UnpackLog(event, "ManagerAdded", log); err != nil {
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

// ParseManagerAdded is a log parse operation binding the contract event 0x3b4a40cccf2058c593542587329dd385be4f0b588db5471fbd9598e56dd7093a.
//
// Solidity: event ManagerAdded(address indexed _manager)
func (_WalletFactory *WalletFactoryFilterer) ParseManagerAdded(log types.Log) (*WalletFactoryManagerAdded, error) {
	event := new(WalletFactoryManagerAdded)
	if err := _WalletFactory.contract.UnpackLog(event, "ManagerAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletFactoryManagerRevokedIterator is returned from FilterManagerRevoked and is used to iterate over the raw logs and unpacked data for ManagerRevoked events raised by the WalletFactory contract.
type WalletFactoryManagerRevokedIterator struct {
	Event *WalletFactoryManagerRevoked // Event containing the contract specifics and raw log

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
func (it *WalletFactoryManagerRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletFactoryManagerRevoked)
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
		it.Event = new(WalletFactoryManagerRevoked)
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
func (it *WalletFactoryManagerRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletFactoryManagerRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletFactoryManagerRevoked represents a ManagerRevoked event raised by the WalletFactory contract.
type WalletFactoryManagerRevoked struct {
	Manager common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterManagerRevoked is a free log retrieval operation binding the contract event 0xe5def11e0516f317f9c37b8835aec29fc01db4d4b6d6fecaca339d3596a29bc1.
//
// Solidity: event ManagerRevoked(address indexed _manager)
func (_WalletFactory *WalletFactoryFilterer) FilterManagerRevoked(opts *bind.FilterOpts, _manager []common.Address) (*WalletFactoryManagerRevokedIterator, error) {

	var _managerRule []interface{}
	for _, _managerItem := range _manager {
		_managerRule = append(_managerRule, _managerItem)
	}

	logs, sub, err := _WalletFactory.contract.FilterLogs(opts, "ManagerRevoked", _managerRule)
	if err != nil {
		return nil, err
	}
	return &WalletFactoryManagerRevokedIterator{contract: _WalletFactory.contract, event: "ManagerRevoked", logs: logs, sub: sub}, nil
}

// WatchManagerRevoked is a free log subscription operation binding the contract event 0xe5def11e0516f317f9c37b8835aec29fc01db4d4b6d6fecaca339d3596a29bc1.
//
// Solidity: event ManagerRevoked(address indexed _manager)
func (_WalletFactory *WalletFactoryFilterer) WatchManagerRevoked(opts *bind.WatchOpts, sink chan<- *WalletFactoryManagerRevoked, _manager []common.Address) (event.Subscription, error) {

	var _managerRule []interface{}
	for _, _managerItem := range _manager {
		_managerRule = append(_managerRule, _managerItem)
	}

	logs, sub, err := _WalletFactory.contract.WatchLogs(opts, "ManagerRevoked", _managerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletFactoryManagerRevoked)
				if err := _WalletFactory.contract.UnpackLog(event, "ManagerRevoked", log); err != nil {
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

// ParseManagerRevoked is a log parse operation binding the contract event 0xe5def11e0516f317f9c37b8835aec29fc01db4d4b6d6fecaca339d3596a29bc1.
//
// Solidity: event ManagerRevoked(address indexed _manager)
func (_WalletFactory *WalletFactoryFilterer) ParseManagerRevoked(log types.Log) (*WalletFactoryManagerRevoked, error) {
	event := new(WalletFactoryManagerRevoked)
	if err := _WalletFactory.contract.UnpackLog(event, "ManagerRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletFactoryOwnerChangedIterator is returned from FilterOwnerChanged and is used to iterate over the raw logs and unpacked data for OwnerChanged events raised by the WalletFactory contract.
type WalletFactoryOwnerChangedIterator struct {
	Event *WalletFactoryOwnerChanged // Event containing the contract specifics and raw log

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
func (it *WalletFactoryOwnerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletFactoryOwnerChanged)
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
		it.Event = new(WalletFactoryOwnerChanged)
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
func (it *WalletFactoryOwnerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletFactoryOwnerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletFactoryOwnerChanged represents a OwnerChanged event raised by the WalletFactory contract.
type WalletFactoryOwnerChanged struct {
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerChanged is a free log retrieval operation binding the contract event 0xa2ea9883a321a3e97b8266c2b078bfeec6d50c711ed71f874a90d500ae2eaf36.
//
// Solidity: event OwnerChanged(address indexed _newOwner)
func (_WalletFactory *WalletFactoryFilterer) FilterOwnerChanged(opts *bind.FilterOpts, _newOwner []common.Address) (*WalletFactoryOwnerChangedIterator, error) {

	var _newOwnerRule []interface{}
	for _, _newOwnerItem := range _newOwner {
		_newOwnerRule = append(_newOwnerRule, _newOwnerItem)
	}

	logs, sub, err := _WalletFactory.contract.FilterLogs(opts, "OwnerChanged", _newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &WalletFactoryOwnerChangedIterator{contract: _WalletFactory.contract, event: "OwnerChanged", logs: logs, sub: sub}, nil
}

// WatchOwnerChanged is a free log subscription operation binding the contract event 0xa2ea9883a321a3e97b8266c2b078bfeec6d50c711ed71f874a90d500ae2eaf36.
//
// Solidity: event OwnerChanged(address indexed _newOwner)
func (_WalletFactory *WalletFactoryFilterer) WatchOwnerChanged(opts *bind.WatchOpts, sink chan<- *WalletFactoryOwnerChanged, _newOwner []common.Address) (event.Subscription, error) {

	var _newOwnerRule []interface{}
	for _, _newOwnerItem := range _newOwner {
		_newOwnerRule = append(_newOwnerRule, _newOwnerItem)
	}

	logs, sub, err := _WalletFactory.contract.WatchLogs(opts, "OwnerChanged", _newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletFactoryOwnerChanged)
				if err := _WalletFactory.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
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

// ParseOwnerChanged is a log parse operation binding the contract event 0xa2ea9883a321a3e97b8266c2b078bfeec6d50c711ed71f874a90d500ae2eaf36.
//
// Solidity: event OwnerChanged(address indexed _newOwner)
func (_WalletFactory *WalletFactoryFilterer) ParseOwnerChanged(log types.Log) (*WalletFactoryOwnerChanged, error) {
	event := new(WalletFactoryOwnerChanged)
	if err := _WalletFactory.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletFactoryRefundAddressChangedIterator is returned from FilterRefundAddressChanged and is used to iterate over the raw logs and unpacked data for RefundAddressChanged events raised by the WalletFactory contract.
type WalletFactoryRefundAddressChangedIterator struct {
	Event *WalletFactoryRefundAddressChanged // Event containing the contract specifics and raw log

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
func (it *WalletFactoryRefundAddressChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletFactoryRefundAddressChanged)
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
		it.Event = new(WalletFactoryRefundAddressChanged)
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
func (it *WalletFactoryRefundAddressChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletFactoryRefundAddressChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletFactoryRefundAddressChanged represents a RefundAddressChanged event raised by the WalletFactory contract.
type WalletFactoryRefundAddressChanged struct {
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterRefundAddressChanged is a free log retrieval operation binding the contract event 0xa7cb165192538768851363c5aa55b1ade75d692a51063730feccdd57d002a6ed.
//
// Solidity: event RefundAddressChanged(address addr)
func (_WalletFactory *WalletFactoryFilterer) FilterRefundAddressChanged(opts *bind.FilterOpts) (*WalletFactoryRefundAddressChangedIterator, error) {

	logs, sub, err := _WalletFactory.contract.FilterLogs(opts, "RefundAddressChanged")
	if err != nil {
		return nil, err
	}
	return &WalletFactoryRefundAddressChangedIterator{contract: _WalletFactory.contract, event: "RefundAddressChanged", logs: logs, sub: sub}, nil
}

// WatchRefundAddressChanged is a free log subscription operation binding the contract event 0xa7cb165192538768851363c5aa55b1ade75d692a51063730feccdd57d002a6ed.
//
// Solidity: event RefundAddressChanged(address addr)
func (_WalletFactory *WalletFactoryFilterer) WatchRefundAddressChanged(opts *bind.WatchOpts, sink chan<- *WalletFactoryRefundAddressChanged) (event.Subscription, error) {

	logs, sub, err := _WalletFactory.contract.WatchLogs(opts, "RefundAddressChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletFactoryRefundAddressChanged)
				if err := _WalletFactory.contract.UnpackLog(event, "RefundAddressChanged", log); err != nil {
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

// ParseRefundAddressChanged is a log parse operation binding the contract event 0xa7cb165192538768851363c5aa55b1ade75d692a51063730feccdd57d002a6ed.
//
// Solidity: event RefundAddressChanged(address addr)
func (_WalletFactory *WalletFactoryFilterer) ParseRefundAddressChanged(log types.Log) (*WalletFactoryRefundAddressChanged, error) {
	event := new(WalletFactoryRefundAddressChanged)
	if err := _WalletFactory.contract.UnpackLog(event, "RefundAddressChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletFactoryWalletCreatedIterator is returned from FilterWalletCreated and is used to iterate over the raw logs and unpacked data for WalletCreated events raised by the WalletFactory contract.
type WalletFactoryWalletCreatedIterator struct {
	Event *WalletFactoryWalletCreated // Event containing the contract specifics and raw log

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
func (it *WalletFactoryWalletCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletFactoryWalletCreated)
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
		it.Event = new(WalletFactoryWalletCreated)
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
func (it *WalletFactoryWalletCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletFactoryWalletCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletFactoryWalletCreated represents a WalletCreated event raised by the WalletFactory contract.
type WalletFactoryWalletCreated struct {
	Wallet       common.Address
	Owner        common.Address
	Guardian     common.Address
	RefundToken  common.Address
	RefundAmount *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterWalletCreated is a free log retrieval operation binding the contract event 0xacd43e061b8ee386537b57919358cfa44933c4e47ccd4b2e1916e54133cec748.
//
// Solidity: event WalletCreated(address indexed wallet, address indexed owner, address indexed guardian, address refundToken, uint256 refundAmount)
func (_WalletFactory *WalletFactoryFilterer) FilterWalletCreated(opts *bind.FilterOpts, wallet []common.Address, owner []common.Address, guardian []common.Address) (*WalletFactoryWalletCreatedIterator, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var guardianRule []interface{}
	for _, guardianItem := range guardian {
		guardianRule = append(guardianRule, guardianItem)
	}

	logs, sub, err := _WalletFactory.contract.FilterLogs(opts, "WalletCreated", walletRule, ownerRule, guardianRule)
	if err != nil {
		return nil, err
	}
	return &WalletFactoryWalletCreatedIterator{contract: _WalletFactory.contract, event: "WalletCreated", logs: logs, sub: sub}, nil
}

// WatchWalletCreated is a free log subscription operation binding the contract event 0xacd43e061b8ee386537b57919358cfa44933c4e47ccd4b2e1916e54133cec748.
//
// Solidity: event WalletCreated(address indexed wallet, address indexed owner, address indexed guardian, address refundToken, uint256 refundAmount)
func (_WalletFactory *WalletFactoryFilterer) WatchWalletCreated(opts *bind.WatchOpts, sink chan<- *WalletFactoryWalletCreated, wallet []common.Address, owner []common.Address, guardian []common.Address) (event.Subscription, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var guardianRule []interface{}
	for _, guardianItem := range guardian {
		guardianRule = append(guardianRule, guardianItem)
	}

	logs, sub, err := _WalletFactory.contract.WatchLogs(opts, "WalletCreated", walletRule, ownerRule, guardianRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletFactoryWalletCreated)
				if err := _WalletFactory.contract.UnpackLog(event, "WalletCreated", log); err != nil {
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

// ParseWalletCreated is a log parse operation binding the contract event 0xacd43e061b8ee386537b57919358cfa44933c4e47ccd4b2e1916e54133cec748.
//
// Solidity: event WalletCreated(address indexed wallet, address indexed owner, address indexed guardian, address refundToken, uint256 refundAmount)
func (_WalletFactory *WalletFactoryFilterer) ParseWalletCreated(log types.Log) (*WalletFactoryWalletCreated, error) {
	event := new(WalletFactoryWalletCreated)
	if err := _WalletFactory.contract.UnpackLog(event, "WalletCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
