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

// TransactionManagerCall is an auto generated low-level Go binding around an user-defined struct.
type TransactionManagerCall struct {
	To    common.Address
	Value *big.Int
	Data  []byte
}

// ArgentModuleMetaData contains all meta data concerning the ArgentModule contract.
var ArgentModuleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIModuleRegistry\",\"name\":\"_registry\",\"type\":\"address\"},{\"internalType\":\"contractIGuardianStorage\",\"name\":\"_guardianStorage\",\"type\":\"address\"},{\"internalType\":\"contractITransferStorage\",\"name\":\"_userWhitelist\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_uniswapRouter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_securityPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_securityWindow\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_recoveryPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_lockPeriod\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"whitelistAfter\",\"type\":\"uint64\"}],\"name\":\"AddedToWhitelist\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"guardian\",\"type\":\"address\"}],\"name\":\"GuardianAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"guardian\",\"type\":\"address\"}],\"name\":\"GuardianAdditionCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"guardian\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"executeAfter\",\"type\":\"uint256\"}],\"name\":\"GuardianAdditionRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"guardian\",\"type\":\"address\"}],\"name\":\"GuardianRevokationCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"guardian\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"executeAfter\",\"type\":\"uint256\"}],\"name\":\"GuardianRevokationRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"guardian\",\"type\":\"address\"}],\"name\":\"GuardianRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"releaseAfter\",\"type\":\"uint64\"}],\"name\":\"Locked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"}],\"name\":\"ModuleCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransfered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_recovery\",\"type\":\"address\"}],\"name\":\"RecoveryCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_recovery\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"executeAfter\",\"type\":\"uint64\"}],\"name\":\"RecoveryExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_recovery\",\"type\":\"address\"}],\"name\":\"RecoveryFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"refundAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"refundToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"refundAmount\",\"type\":\"uint256\"}],\"name\":\"Refund\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"RemovedFromWhitelist\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sessionKey\",\"type\":\"address\"}],\"name\":\"SessionCleared\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sessionKey\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"expires\",\"type\":\"uint64\"}],\"name\":\"SessionCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"signedHash\",\"type\":\"bytes32\"}],\"name\":\"TransactionExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"}],\"name\":\"Unlocked\",\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_wallet\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_guardian\",\"type\":\"address\"}],\"name\":\"addGuardian\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_wallet\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"}],\"name\":\"addToWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_wallet\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_guardian\",\"type\":\"address\"}],\"name\":\"cancelGuardianAddition\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_wallet\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_guardian\",\"type\":\"address\"}],\"name\":\"cancelGuardianRevokation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_wallet\",\"type\":\"address\"}],\"name\":\"cancelRecovery\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_wallet\",\"type\":\"address\"}],\"name\":\"clearSession\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_wallet\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_guardian\",\"type\":\"address\"}],\"name\":\"confirmGuardianAddition\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_wallet\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_guardian\",\"type\":\"address\"}],\"name\":\"confirmGuardianRevokation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_wallet\",\"type\":\"address\"}],\"name\":\"enableERC1155TokenReceiver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_wallet\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_signatures\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_gasPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_refundToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_refundAddress\",\"type\":\"address\"}],\"name\":\"execute\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_wallet\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_recovery\",\"type\":\"address\"}],\"name\":\"executeRecovery\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_wallet\",\"type\":\"address\"}],\"name\":\"finalizeRecovery\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_wallet\",\"type\":\"address\"}],\"name\":\"getGuardians\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_guardians\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_wallet\",\"type\":\"address\"}],\"name\":\"getLock\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"_releaseAfter\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_wallet\",\"type\":\"address\"}],\"name\":\"getNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_wallet\",\"type\":\"address\"}],\"name\":\"getRecovery\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_executeAfter\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"_guardianCount\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_wallet\",\"type\":\"address\"}],\"name\":\"getSession\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"key\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"expires\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_refundToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_refundAddress\",\"type\":\"address\"}],\"name\":\"getSignHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_wallet\",\"type\":\"address\"}],\"name\":\"guardianCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_count\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_wallet\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_signHash\",\"type\":\"bytes32\"}],\"name\":\"isExecutedTx\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"executed\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_wallet\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_guardian\",\"type\":\"address\"}],\"name\":\"isGuardian\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_isGuardian\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_wallet\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_guardian\",\"type\":\"address\"}],\"name\":\"isGuardianOrGuardianSigner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_isGuardian\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_wallet\",\"type\":\"address\"}],\"name\":\"isLocked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_msgHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"isValidSignature\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_wallet\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"}],\"name\":\"isWhitelisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_isWhitelisted\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_wallet\",\"type\":\"address\"}],\"name\":\"lock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_wallet\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structTransactionManager.Call[]\",\"name\":\"_transactions\",\"type\":\"tuple[]\"}],\"name\":\"multiCall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_wallet\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structTransactionManager.Call[]\",\"name\":\"_transactions\",\"type\":\"tuple[]\"}],\"name\":\"multiCallWithGuardians\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"recoverToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_wallet\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"}],\"name\":\"removeFromWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_wallet\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_guardian\",\"type\":\"address\"}],\"name\":\"revokeGuardian\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_interfaceID\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_methodId\",\"type\":\"bytes4\"}],\"name\":\"supportsStaticCall\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_isSupported\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_wallet\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_wallet\",\"type\":\"address\"}],\"name\":\"unlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_wallet\",\"type\":\"address\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_wallet\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_module\",\"type\":\"address\"}],\"name\":\"addModule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_wallet\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"getRequiredSignatures\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"enumBaseModule.OwnerSignature\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true}]",
}

// ArgentModuleABI is the input ABI used to generate the binding from.
// Deprecated: Use ArgentModuleMetaData.ABI instead.
var ArgentModuleABI = ArgentModuleMetaData.ABI

// ArgentModule is an auto generated Go binding around an Ethereum contract.
type ArgentModule struct {
	ArgentModuleCaller     // Read-only binding to the contract
	ArgentModuleTransactor // Write-only binding to the contract
	ArgentModuleFilterer   // Log filterer for contract events
}

// ArgentModuleCaller is an auto generated read-only Go binding around an Ethereum contract.
type ArgentModuleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArgentModuleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ArgentModuleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArgentModuleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ArgentModuleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArgentModuleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ArgentModuleSession struct {
	Contract     *ArgentModule     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ArgentModuleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ArgentModuleCallerSession struct {
	Contract *ArgentModuleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ArgentModuleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ArgentModuleTransactorSession struct {
	Contract     *ArgentModuleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ArgentModuleRaw is an auto generated low-level Go binding around an Ethereum contract.
type ArgentModuleRaw struct {
	Contract *ArgentModule // Generic contract binding to access the raw methods on
}

// ArgentModuleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ArgentModuleCallerRaw struct {
	Contract *ArgentModuleCaller // Generic read-only contract binding to access the raw methods on
}

// ArgentModuleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ArgentModuleTransactorRaw struct {
	Contract *ArgentModuleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewArgentModule creates a new instance of ArgentModule, bound to a specific deployed contract.
func NewArgentModule(address common.Address, backend bind.ContractBackend) (*ArgentModule, error) {
	contract, err := bindArgentModule(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ArgentModule{ArgentModuleCaller: ArgentModuleCaller{contract: contract}, ArgentModuleTransactor: ArgentModuleTransactor{contract: contract}, ArgentModuleFilterer: ArgentModuleFilterer{contract: contract}}, nil
}

// NewArgentModuleCaller creates a new read-only instance of ArgentModule, bound to a specific deployed contract.
func NewArgentModuleCaller(address common.Address, caller bind.ContractCaller) (*ArgentModuleCaller, error) {
	contract, err := bindArgentModule(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ArgentModuleCaller{contract: contract}, nil
}

// NewArgentModuleTransactor creates a new write-only instance of ArgentModule, bound to a specific deployed contract.
func NewArgentModuleTransactor(address common.Address, transactor bind.ContractTransactor) (*ArgentModuleTransactor, error) {
	contract, err := bindArgentModule(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ArgentModuleTransactor{contract: contract}, nil
}

// NewArgentModuleFilterer creates a new log filterer instance of ArgentModule, bound to a specific deployed contract.
func NewArgentModuleFilterer(address common.Address, filterer bind.ContractFilterer) (*ArgentModuleFilterer, error) {
	contract, err := bindArgentModule(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ArgentModuleFilterer{contract: contract}, nil
}

// bindArgentModule binds a generic wrapper to an already deployed contract.
func bindArgentModule(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ArgentModuleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ArgentModule *ArgentModuleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ArgentModule.Contract.ArgentModuleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ArgentModule *ArgentModuleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArgentModule.Contract.ArgentModuleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ArgentModule *ArgentModuleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArgentModule.Contract.ArgentModuleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ArgentModule *ArgentModuleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ArgentModule.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ArgentModule *ArgentModuleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArgentModule.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ArgentModule *ArgentModuleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArgentModule.Contract.contract.Transact(opts, method, params...)
}

// NAME is a free data retrieval call binding the contract method 0xa3f4df7e.
//
// Solidity: function NAME() view returns(bytes32)
func (_ArgentModule *ArgentModuleCaller) NAME(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ArgentModule.contract.Call(opts, &out, "NAME")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// NAME is a free data retrieval call binding the contract method 0xa3f4df7e.
//
// Solidity: function NAME() view returns(bytes32)
func (_ArgentModule *ArgentModuleSession) NAME() ([32]byte, error) {
	return _ArgentModule.Contract.NAME(&_ArgentModule.CallOpts)
}

// NAME is a free data retrieval call binding the contract method 0xa3f4df7e.
//
// Solidity: function NAME() view returns(bytes32)
func (_ArgentModule *ArgentModuleCallerSession) NAME() ([32]byte, error) {
	return _ArgentModule.Contract.NAME(&_ArgentModule.CallOpts)
}

// GetGuardians is a free data retrieval call binding the contract method 0xf18858ab.
//
// Solidity: function getGuardians(address _wallet) view returns(address[] _guardians)
func (_ArgentModule *ArgentModuleCaller) GetGuardians(opts *bind.CallOpts, _wallet common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _ArgentModule.contract.Call(opts, &out, "getGuardians", _wallet)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetGuardians is a free data retrieval call binding the contract method 0xf18858ab.
//
// Solidity: function getGuardians(address _wallet) view returns(address[] _guardians)
func (_ArgentModule *ArgentModuleSession) GetGuardians(_wallet common.Address) ([]common.Address, error) {
	return _ArgentModule.Contract.GetGuardians(&_ArgentModule.CallOpts, _wallet)
}

// GetGuardians is a free data retrieval call binding the contract method 0xf18858ab.
//
// Solidity: function getGuardians(address _wallet) view returns(address[] _guardians)
func (_ArgentModule *ArgentModuleCallerSession) GetGuardians(_wallet common.Address) ([]common.Address, error) {
	return _ArgentModule.Contract.GetGuardians(&_ArgentModule.CallOpts, _wallet)
}

// GetLock is a free data retrieval call binding the contract method 0x6b9db4e6.
//
// Solidity: function getLock(address _wallet) view returns(uint64 _releaseAfter)
func (_ArgentModule *ArgentModuleCaller) GetLock(opts *bind.CallOpts, _wallet common.Address) (uint64, error) {
	var out []interface{}
	err := _ArgentModule.contract.Call(opts, &out, "getLock", _wallet)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetLock is a free data retrieval call binding the contract method 0x6b9db4e6.
//
// Solidity: function getLock(address _wallet) view returns(uint64 _releaseAfter)
func (_ArgentModule *ArgentModuleSession) GetLock(_wallet common.Address) (uint64, error) {
	return _ArgentModule.Contract.GetLock(&_ArgentModule.CallOpts, _wallet)
}

// GetLock is a free data retrieval call binding the contract method 0x6b9db4e6.
//
// Solidity: function getLock(address _wallet) view returns(uint64 _releaseAfter)
func (_ArgentModule *ArgentModuleCallerSession) GetLock(_wallet common.Address) (uint64, error) {
	return _ArgentModule.Contract.GetLock(&_ArgentModule.CallOpts, _wallet)
}

// GetNonce is a free data retrieval call binding the contract method 0x2d0335ab.
//
// Solidity: function getNonce(address _wallet) view returns(uint256 nonce)
func (_ArgentModule *ArgentModuleCaller) GetNonce(opts *bind.CallOpts, _wallet common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ArgentModule.contract.Call(opts, &out, "getNonce", _wallet)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNonce is a free data retrieval call binding the contract method 0x2d0335ab.
//
// Solidity: function getNonce(address _wallet) view returns(uint256 nonce)
func (_ArgentModule *ArgentModuleSession) GetNonce(_wallet common.Address) (*big.Int, error) {
	return _ArgentModule.Contract.GetNonce(&_ArgentModule.CallOpts, _wallet)
}

// GetNonce is a free data retrieval call binding the contract method 0x2d0335ab.
//
// Solidity: function getNonce(address _wallet) view returns(uint256 nonce)
func (_ArgentModule *ArgentModuleCallerSession) GetNonce(_wallet common.Address) (*big.Int, error) {
	return _ArgentModule.Contract.GetNonce(&_ArgentModule.CallOpts, _wallet)
}

// GetRecovery is a free data retrieval call binding the contract method 0x9769c3fe.
//
// Solidity: function getRecovery(address _wallet) view returns(address _address, uint64 _executeAfter, uint32 _guardianCount)
func (_ArgentModule *ArgentModuleCaller) GetRecovery(opts *bind.CallOpts, _wallet common.Address) (struct {
	Address       common.Address
	ExecuteAfter  uint64
	GuardianCount uint32
}, error) {
	var out []interface{}
	err := _ArgentModule.contract.Call(opts, &out, "getRecovery", _wallet)

	outstruct := new(struct {
		Address       common.Address
		ExecuteAfter  uint64
		GuardianCount uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Address = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.ExecuteAfter = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.GuardianCount = *abi.ConvertType(out[2], new(uint32)).(*uint32)

	return *outstruct, err

}

// GetRecovery is a free data retrieval call binding the contract method 0x9769c3fe.
//
// Solidity: function getRecovery(address _wallet) view returns(address _address, uint64 _executeAfter, uint32 _guardianCount)
func (_ArgentModule *ArgentModuleSession) GetRecovery(_wallet common.Address) (struct {
	Address       common.Address
	ExecuteAfter  uint64
	GuardianCount uint32
}, error) {
	return _ArgentModule.Contract.GetRecovery(&_ArgentModule.CallOpts, _wallet)
}

// GetRecovery is a free data retrieval call binding the contract method 0x9769c3fe.
//
// Solidity: function getRecovery(address _wallet) view returns(address _address, uint64 _executeAfter, uint32 _guardianCount)
func (_ArgentModule *ArgentModuleCallerSession) GetRecovery(_wallet common.Address) (struct {
	Address       common.Address
	ExecuteAfter  uint64
	GuardianCount uint32
}, error) {
	return _ArgentModule.Contract.GetRecovery(&_ArgentModule.CallOpts, _wallet)
}

// GetRequiredSignatures is a free data retrieval call binding the contract method 0x3b73d67f.
//
// Solidity: function getRequiredSignatures(address _wallet, bytes _data) view returns(uint256, uint8)
func (_ArgentModule *ArgentModuleCaller) GetRequiredSignatures(opts *bind.CallOpts, _wallet common.Address, _data []byte) (*big.Int, uint8, error) {
	var out []interface{}
	err := _ArgentModule.contract.Call(opts, &out, "getRequiredSignatures", _wallet, _data)

	if err != nil {
		return *new(*big.Int), *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(uint8)).(*uint8)

	return out0, out1, err

}

// GetRequiredSignatures is a free data retrieval call binding the contract method 0x3b73d67f.
//
// Solidity: function getRequiredSignatures(address _wallet, bytes _data) view returns(uint256, uint8)
func (_ArgentModule *ArgentModuleSession) GetRequiredSignatures(_wallet common.Address, _data []byte) (*big.Int, uint8, error) {
	return _ArgentModule.Contract.GetRequiredSignatures(&_ArgentModule.CallOpts, _wallet, _data)
}

// GetRequiredSignatures is a free data retrieval call binding the contract method 0x3b73d67f.
//
// Solidity: function getRequiredSignatures(address _wallet, bytes _data) view returns(uint256, uint8)
func (_ArgentModule *ArgentModuleCallerSession) GetRequiredSignatures(_wallet common.Address, _data []byte) (*big.Int, uint8, error) {
	return _ArgentModule.Contract.GetRequiredSignatures(&_ArgentModule.CallOpts, _wallet, _data)
}

// GetSession is a free data retrieval call binding the contract method 0x8c8e13b9.
//
// Solidity: function getSession(address _wallet) view returns(address key, uint64 expires)
func (_ArgentModule *ArgentModuleCaller) GetSession(opts *bind.CallOpts, _wallet common.Address) (struct {
	Key     common.Address
	Expires uint64
}, error) {
	var out []interface{}
	err := _ArgentModule.contract.Call(opts, &out, "getSession", _wallet)

	outstruct := new(struct {
		Key     common.Address
		Expires uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Key = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Expires = *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return *outstruct, err

}

// GetSession is a free data retrieval call binding the contract method 0x8c8e13b9.
//
// Solidity: function getSession(address _wallet) view returns(address key, uint64 expires)
func (_ArgentModule *ArgentModuleSession) GetSession(_wallet common.Address) (struct {
	Key     common.Address
	Expires uint64
}, error) {
	return _ArgentModule.Contract.GetSession(&_ArgentModule.CallOpts, _wallet)
}

// GetSession is a free data retrieval call binding the contract method 0x8c8e13b9.
//
// Solidity: function getSession(address _wallet) view returns(address key, uint64 expires)
func (_ArgentModule *ArgentModuleCallerSession) GetSession(_wallet common.Address) (struct {
	Key     common.Address
	Expires uint64
}, error) {
	return _ArgentModule.Contract.GetSession(&_ArgentModule.CallOpts, _wallet)
}

// GetSignHash is a free data retrieval call binding the contract method 0x620ccb5c.
//
// Solidity: function getSignHash(address _from, uint256 _value, bytes _data, uint256 _nonce, uint256 _gasPrice, uint256 _gasLimit, address _refundToken, address _refundAddress) view returns(bytes32, bytes)
func (_ArgentModule *ArgentModuleCaller) GetSignHash(opts *bind.CallOpts, _from common.Address, _value *big.Int, _data []byte, _nonce *big.Int, _gasPrice *big.Int, _gasLimit *big.Int, _refundToken common.Address, _refundAddress common.Address) ([32]byte, []byte, error) {
	var out []interface{}
	err := _ArgentModule.contract.Call(opts, &out, "getSignHash", _from, _value, _data, _nonce, _gasPrice, _gasLimit, _refundToken, _refundAddress)

	if err != nil {
		return *new([32]byte), *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	out1 := *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return out0, out1, err

}

// GetSignHash is a free data retrieval call binding the contract method 0x620ccb5c.
//
// Solidity: function getSignHash(address _from, uint256 _value, bytes _data, uint256 _nonce, uint256 _gasPrice, uint256 _gasLimit, address _refundToken, address _refundAddress) view returns(bytes32, bytes)
func (_ArgentModule *ArgentModuleSession) GetSignHash(_from common.Address, _value *big.Int, _data []byte, _nonce *big.Int, _gasPrice *big.Int, _gasLimit *big.Int, _refundToken common.Address, _refundAddress common.Address) ([32]byte, []byte, error) {
	return _ArgentModule.Contract.GetSignHash(&_ArgentModule.CallOpts, _from, _value, _data, _nonce, _gasPrice, _gasLimit, _refundToken, _refundAddress)
}

// GetSignHash is a free data retrieval call binding the contract method 0x620ccb5c.
//
// Solidity: function getSignHash(address _from, uint256 _value, bytes _data, uint256 _nonce, uint256 _gasPrice, uint256 _gasLimit, address _refundToken, address _refundAddress) view returns(bytes32, bytes)
func (_ArgentModule *ArgentModuleCallerSession) GetSignHash(_from common.Address, _value *big.Int, _data []byte, _nonce *big.Int, _gasPrice *big.Int, _gasLimit *big.Int, _refundToken common.Address, _refundAddress common.Address) ([32]byte, []byte, error) {
	return _ArgentModule.Contract.GetSignHash(&_ArgentModule.CallOpts, _from, _value, _data, _nonce, _gasPrice, _gasLimit, _refundToken, _refundAddress)
}

// GuardianCount is a free data retrieval call binding the contract method 0x5040fb76.
//
// Solidity: function guardianCount(address _wallet) view returns(uint256 _count)
func (_ArgentModule *ArgentModuleCaller) GuardianCount(opts *bind.CallOpts, _wallet common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ArgentModule.contract.Call(opts, &out, "guardianCount", _wallet)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GuardianCount is a free data retrieval call binding the contract method 0x5040fb76.
//
// Solidity: function guardianCount(address _wallet) view returns(uint256 _count)
func (_ArgentModule *ArgentModuleSession) GuardianCount(_wallet common.Address) (*big.Int, error) {
	return _ArgentModule.Contract.GuardianCount(&_ArgentModule.CallOpts, _wallet)
}

// GuardianCount is a free data retrieval call binding the contract method 0x5040fb76.
//
// Solidity: function guardianCount(address _wallet) view returns(uint256 _count)
func (_ArgentModule *ArgentModuleCallerSession) GuardianCount(_wallet common.Address) (*big.Int, error) {
	return _ArgentModule.Contract.GuardianCount(&_ArgentModule.CallOpts, _wallet)
}

// IsExecutedTx is a free data retrieval call binding the contract method 0x60c0fdc0.
//
// Solidity: function isExecutedTx(address _wallet, bytes32 _signHash) view returns(bool executed)
func (_ArgentModule *ArgentModuleCaller) IsExecutedTx(opts *bind.CallOpts, _wallet common.Address, _signHash [32]byte) (bool, error) {
	var out []interface{}
	err := _ArgentModule.contract.Call(opts, &out, "isExecutedTx", _wallet, _signHash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsExecutedTx is a free data retrieval call binding the contract method 0x60c0fdc0.
//
// Solidity: function isExecutedTx(address _wallet, bytes32 _signHash) view returns(bool executed)
func (_ArgentModule *ArgentModuleSession) IsExecutedTx(_wallet common.Address, _signHash [32]byte) (bool, error) {
	return _ArgentModule.Contract.IsExecutedTx(&_ArgentModule.CallOpts, _wallet, _signHash)
}

// IsExecutedTx is a free data retrieval call binding the contract method 0x60c0fdc0.
//
// Solidity: function isExecutedTx(address _wallet, bytes32 _signHash) view returns(bool executed)
func (_ArgentModule *ArgentModuleCallerSession) IsExecutedTx(_wallet common.Address, _signHash [32]byte) (bool, error) {
	return _ArgentModule.Contract.IsExecutedTx(&_ArgentModule.CallOpts, _wallet, _signHash)
}

// IsGuardian is a free data retrieval call binding the contract method 0xd4ee9734.
//
// Solidity: function isGuardian(address _wallet, address _guardian) view returns(bool _isGuardian)
func (_ArgentModule *ArgentModuleCaller) IsGuardian(opts *bind.CallOpts, _wallet common.Address, _guardian common.Address) (bool, error) {
	var out []interface{}
	err := _ArgentModule.contract.Call(opts, &out, "isGuardian", _wallet, _guardian)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsGuardian is a free data retrieval call binding the contract method 0xd4ee9734.
//
// Solidity: function isGuardian(address _wallet, address _guardian) view returns(bool _isGuardian)
func (_ArgentModule *ArgentModuleSession) IsGuardian(_wallet common.Address, _guardian common.Address) (bool, error) {
	return _ArgentModule.Contract.IsGuardian(&_ArgentModule.CallOpts, _wallet, _guardian)
}

// IsGuardian is a free data retrieval call binding the contract method 0xd4ee9734.
//
// Solidity: function isGuardian(address _wallet, address _guardian) view returns(bool _isGuardian)
func (_ArgentModule *ArgentModuleCallerSession) IsGuardian(_wallet common.Address, _guardian common.Address) (bool, error) {
	return _ArgentModule.Contract.IsGuardian(&_ArgentModule.CallOpts, _wallet, _guardian)
}

// IsGuardianOrGuardianSigner is a free data retrieval call binding the contract method 0xeac01e47.
//
// Solidity: function isGuardianOrGuardianSigner(address _wallet, address _guardian) view returns(bool _isGuardian)
func (_ArgentModule *ArgentModuleCaller) IsGuardianOrGuardianSigner(opts *bind.CallOpts, _wallet common.Address, _guardian common.Address) (bool, error) {
	var out []interface{}
	err := _ArgentModule.contract.Call(opts, &out, "isGuardianOrGuardianSigner", _wallet, _guardian)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsGuardianOrGuardianSigner is a free data retrieval call binding the contract method 0xeac01e47.
//
// Solidity: function isGuardianOrGuardianSigner(address _wallet, address _guardian) view returns(bool _isGuardian)
func (_ArgentModule *ArgentModuleSession) IsGuardianOrGuardianSigner(_wallet common.Address, _guardian common.Address) (bool, error) {
	return _ArgentModule.Contract.IsGuardianOrGuardianSigner(&_ArgentModule.CallOpts, _wallet, _guardian)
}

// IsGuardianOrGuardianSigner is a free data retrieval call binding the contract method 0xeac01e47.
//
// Solidity: function isGuardianOrGuardianSigner(address _wallet, address _guardian) view returns(bool _isGuardian)
func (_ArgentModule *ArgentModuleCallerSession) IsGuardianOrGuardianSigner(_wallet common.Address, _guardian common.Address) (bool, error) {
	return _ArgentModule.Contract.IsGuardianOrGuardianSigner(&_ArgentModule.CallOpts, _wallet, _guardian)
}

// IsLocked is a free data retrieval call binding the contract method 0x4a4fbeec.
//
// Solidity: function isLocked(address _wallet) view returns(bool)
func (_ArgentModule *ArgentModuleCaller) IsLocked(opts *bind.CallOpts, _wallet common.Address) (bool, error) {
	var out []interface{}
	err := _ArgentModule.contract.Call(opts, &out, "isLocked", _wallet)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsLocked is a free data retrieval call binding the contract method 0x4a4fbeec.
//
// Solidity: function isLocked(address _wallet) view returns(bool)
func (_ArgentModule *ArgentModuleSession) IsLocked(_wallet common.Address) (bool, error) {
	return _ArgentModule.Contract.IsLocked(&_ArgentModule.CallOpts, _wallet)
}

// IsLocked is a free data retrieval call binding the contract method 0x4a4fbeec.
//
// Solidity: function isLocked(address _wallet) view returns(bool)
func (_ArgentModule *ArgentModuleCallerSession) IsLocked(_wallet common.Address) (bool, error) {
	return _ArgentModule.Contract.IsLocked(&_ArgentModule.CallOpts, _wallet)
}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 _msgHash, bytes _signature) view returns(bytes4)
func (_ArgentModule *ArgentModuleCaller) IsValidSignature(opts *bind.CallOpts, _msgHash [32]byte, _signature []byte) ([4]byte, error) {
	var out []interface{}
	err := _ArgentModule.contract.Call(opts, &out, "isValidSignature", _msgHash, _signature)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 _msgHash, bytes _signature) view returns(bytes4)
func (_ArgentModule *ArgentModuleSession) IsValidSignature(_msgHash [32]byte, _signature []byte) ([4]byte, error) {
	return _ArgentModule.Contract.IsValidSignature(&_ArgentModule.CallOpts, _msgHash, _signature)
}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 _msgHash, bytes _signature) view returns(bytes4)
func (_ArgentModule *ArgentModuleCallerSession) IsValidSignature(_msgHash [32]byte, _signature []byte) ([4]byte, error) {
	return _ArgentModule.Contract.IsValidSignature(&_ArgentModule.CallOpts, _msgHash, _signature)
}

// IsWhitelisted is a free data retrieval call binding the contract method 0xb6b35272.
//
// Solidity: function isWhitelisted(address _wallet, address _target) view returns(bool _isWhitelisted)
func (_ArgentModule *ArgentModuleCaller) IsWhitelisted(opts *bind.CallOpts, _wallet common.Address, _target common.Address) (bool, error) {
	var out []interface{}
	err := _ArgentModule.contract.Call(opts, &out, "isWhitelisted", _wallet, _target)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWhitelisted is a free data retrieval call binding the contract method 0xb6b35272.
//
// Solidity: function isWhitelisted(address _wallet, address _target) view returns(bool _isWhitelisted)
func (_ArgentModule *ArgentModuleSession) IsWhitelisted(_wallet common.Address, _target common.Address) (bool, error) {
	return _ArgentModule.Contract.IsWhitelisted(&_ArgentModule.CallOpts, _wallet, _target)
}

// IsWhitelisted is a free data retrieval call binding the contract method 0xb6b35272.
//
// Solidity: function isWhitelisted(address _wallet, address _target) view returns(bool _isWhitelisted)
func (_ArgentModule *ArgentModuleCallerSession) IsWhitelisted(_wallet common.Address, _target common.Address) (bool, error) {
	return _ArgentModule.Contract.IsWhitelisted(&_ArgentModule.CallOpts, _wallet, _target)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceID) pure returns(bool)
func (_ArgentModule *ArgentModuleCaller) SupportsInterface(opts *bind.CallOpts, _interfaceID [4]byte) (bool, error) {
	var out []interface{}
	err := _ArgentModule.contract.Call(opts, &out, "supportsInterface", _interfaceID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceID) pure returns(bool)
func (_ArgentModule *ArgentModuleSession) SupportsInterface(_interfaceID [4]byte) (bool, error) {
	return _ArgentModule.Contract.SupportsInterface(&_ArgentModule.CallOpts, _interfaceID)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceID) pure returns(bool)
func (_ArgentModule *ArgentModuleCallerSession) SupportsInterface(_interfaceID [4]byte) (bool, error) {
	return _ArgentModule.Contract.SupportsInterface(&_ArgentModule.CallOpts, _interfaceID)
}

// SupportsStaticCall is a free data retrieval call binding the contract method 0x25b50934.
//
// Solidity: function supportsStaticCall(bytes4 _methodId) pure returns(bool _isSupported)
func (_ArgentModule *ArgentModuleCaller) SupportsStaticCall(opts *bind.CallOpts, _methodId [4]byte) (bool, error) {
	var out []interface{}
	err := _ArgentModule.contract.Call(opts, &out, "supportsStaticCall", _methodId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsStaticCall is a free data retrieval call binding the contract method 0x25b50934.
//
// Solidity: function supportsStaticCall(bytes4 _methodId) pure returns(bool _isSupported)
func (_ArgentModule *ArgentModuleSession) SupportsStaticCall(_methodId [4]byte) (bool, error) {
	return _ArgentModule.Contract.SupportsStaticCall(&_ArgentModule.CallOpts, _methodId)
}

// SupportsStaticCall is a free data retrieval call binding the contract method 0x25b50934.
//
// Solidity: function supportsStaticCall(bytes4 _methodId) pure returns(bool _isSupported)
func (_ArgentModule *ArgentModuleCallerSession) SupportsStaticCall(_methodId [4]byte) (bool, error) {
	return _ArgentModule.Contract.SupportsStaticCall(&_ArgentModule.CallOpts, _methodId)
}

// AddGuardian is a paid mutator transaction binding the contract method 0xc6845210.
//
// Solidity: function addGuardian(address _wallet, address _guardian) returns()
func (_ArgentModule *ArgentModuleTransactor) AddGuardian(opts *bind.TransactOpts, _wallet common.Address, _guardian common.Address) (*types.Transaction, error) {
	return _ArgentModule.contract.Transact(opts, "addGuardian", _wallet, _guardian)
}

// AddGuardian is a paid mutator transaction binding the contract method 0xc6845210.
//
// Solidity: function addGuardian(address _wallet, address _guardian) returns()
func (_ArgentModule *ArgentModuleSession) AddGuardian(_wallet common.Address, _guardian common.Address) (*types.Transaction, error) {
	return _ArgentModule.Contract.AddGuardian(&_ArgentModule.TransactOpts, _wallet, _guardian)
}

// AddGuardian is a paid mutator transaction binding the contract method 0xc6845210.
//
// Solidity: function addGuardian(address _wallet, address _guardian) returns()
func (_ArgentModule *ArgentModuleTransactorSession) AddGuardian(_wallet common.Address, _guardian common.Address) (*types.Transaction, error) {
	return _ArgentModule.Contract.AddGuardian(&_ArgentModule.TransactOpts, _wallet, _guardian)
}

// AddModule is a paid mutator transaction binding the contract method 0x5a1db8c4.
//
// Solidity: function addModule(address _wallet, address _module) returns()
func (_ArgentModule *ArgentModuleTransactor) AddModule(opts *bind.TransactOpts, _wallet common.Address, _module common.Address) (*types.Transaction, error) {
	return _ArgentModule.contract.Transact(opts, "addModule", _wallet, _module)
}

// AddModule is a paid mutator transaction binding the contract method 0x5a1db8c4.
//
// Solidity: function addModule(address _wallet, address _module) returns()
func (_ArgentModule *ArgentModuleSession) AddModule(_wallet common.Address, _module common.Address) (*types.Transaction, error) {
	return _ArgentModule.Contract.AddModule(&_ArgentModule.TransactOpts, _wallet, _module)
}

// AddModule is a paid mutator transaction binding the contract method 0x5a1db8c4.
//
// Solidity: function addModule(address _wallet, address _module) returns()
func (_ArgentModule *ArgentModuleTransactorSession) AddModule(_wallet common.Address, _module common.Address) (*types.Transaction, error) {
	return _ArgentModule.Contract.AddModule(&_ArgentModule.TransactOpts, _wallet, _module)
}

// AddToWhitelist is a paid mutator transaction binding the contract method 0x57518243.
//
// Solidity: function addToWhitelist(address _wallet, address _target) returns()
func (_ArgentModule *ArgentModuleTransactor) AddToWhitelist(opts *bind.TransactOpts, _wallet common.Address, _target common.Address) (*types.Transaction, error) {
	return _ArgentModule.contract.Transact(opts, "addToWhitelist", _wallet, _target)
}

// AddToWhitelist is a paid mutator transaction binding the contract method 0x57518243.
//
// Solidity: function addToWhitelist(address _wallet, address _target) returns()
func (_ArgentModule *ArgentModuleSession) AddToWhitelist(_wallet common.Address, _target common.Address) (*types.Transaction, error) {
	return _ArgentModule.Contract.AddToWhitelist(&_ArgentModule.TransactOpts, _wallet, _target)
}

// AddToWhitelist is a paid mutator transaction binding the contract method 0x57518243.
//
// Solidity: function addToWhitelist(address _wallet, address _target) returns()
func (_ArgentModule *ArgentModuleTransactorSession) AddToWhitelist(_wallet common.Address, _target common.Address) (*types.Transaction, error) {
	return _ArgentModule.Contract.AddToWhitelist(&_ArgentModule.TransactOpts, _wallet, _target)
}

// CancelGuardianAddition is a paid mutator transaction binding the contract method 0xa6eb0690.
//
// Solidity: function cancelGuardianAddition(address _wallet, address _guardian) returns()
func (_ArgentModule *ArgentModuleTransactor) CancelGuardianAddition(opts *bind.TransactOpts, _wallet common.Address, _guardian common.Address) (*types.Transaction, error) {
	return _ArgentModule.contract.Transact(opts, "cancelGuardianAddition", _wallet, _guardian)
}

// CancelGuardianAddition is a paid mutator transaction binding the contract method 0xa6eb0690.
//
// Solidity: function cancelGuardianAddition(address _wallet, address _guardian) returns()
func (_ArgentModule *ArgentModuleSession) CancelGuardianAddition(_wallet common.Address, _guardian common.Address) (*types.Transaction, error) {
	return _ArgentModule.Contract.CancelGuardianAddition(&_ArgentModule.TransactOpts, _wallet, _guardian)
}

// CancelGuardianAddition is a paid mutator transaction binding the contract method 0xa6eb0690.
//
// Solidity: function cancelGuardianAddition(address _wallet, address _guardian) returns()
func (_ArgentModule *ArgentModuleTransactorSession) CancelGuardianAddition(_wallet common.Address, _guardian common.Address) (*types.Transaction, error) {
	return _ArgentModule.Contract.CancelGuardianAddition(&_ArgentModule.TransactOpts, _wallet, _guardian)
}

// CancelGuardianRevokation is a paid mutator transaction binding the contract method 0x2960739b.
//
// Solidity: function cancelGuardianRevokation(address _wallet, address _guardian) returns()
func (_ArgentModule *ArgentModuleTransactor) CancelGuardianRevokation(opts *bind.TransactOpts, _wallet common.Address, _guardian common.Address) (*types.Transaction, error) {
	return _ArgentModule.contract.Transact(opts, "cancelGuardianRevokation", _wallet, _guardian)
}

// CancelGuardianRevokation is a paid mutator transaction binding the contract method 0x2960739b.
//
// Solidity: function cancelGuardianRevokation(address _wallet, address _guardian) returns()
func (_ArgentModule *ArgentModuleSession) CancelGuardianRevokation(_wallet common.Address, _guardian common.Address) (*types.Transaction, error) {
	return _ArgentModule.Contract.CancelGuardianRevokation(&_ArgentModule.TransactOpts, _wallet, _guardian)
}

// CancelGuardianRevokation is a paid mutator transaction binding the contract method 0x2960739b.
//
// Solidity: function cancelGuardianRevokation(address _wallet, address _guardian) returns()
func (_ArgentModule *ArgentModuleTransactorSession) CancelGuardianRevokation(_wallet common.Address, _guardian common.Address) (*types.Transaction, error) {
	return _ArgentModule.Contract.CancelGuardianRevokation(&_ArgentModule.TransactOpts, _wallet, _guardian)
}

// CancelRecovery is a paid mutator transaction binding the contract method 0xc90db447.
//
// Solidity: function cancelRecovery(address _wallet) returns()
func (_ArgentModule *ArgentModuleTransactor) CancelRecovery(opts *bind.TransactOpts, _wallet common.Address) (*types.Transaction, error) {
	return _ArgentModule.contract.Transact(opts, "cancelRecovery", _wallet)
}

// CancelRecovery is a paid mutator transaction binding the contract method 0xc90db447.
//
// Solidity: function cancelRecovery(address _wallet) returns()
func (_ArgentModule *ArgentModuleSession) CancelRecovery(_wallet common.Address) (*types.Transaction, error) {
	return _ArgentModule.Contract.CancelRecovery(&_ArgentModule.TransactOpts, _wallet)
}

// CancelRecovery is a paid mutator transaction binding the contract method 0xc90db447.
//
// Solidity: function cancelRecovery(address _wallet) returns()
func (_ArgentModule *ArgentModuleTransactorSession) CancelRecovery(_wallet common.Address) (*types.Transaction, error) {
	return _ArgentModule.Contract.CancelRecovery(&_ArgentModule.TransactOpts, _wallet)
}

// ClearSession is a paid mutator transaction binding the contract method 0xba821088.
//
// Solidity: function clearSession(address _wallet) returns()
func (_ArgentModule *ArgentModuleTransactor) ClearSession(opts *bind.TransactOpts, _wallet common.Address) (*types.Transaction, error) {
	return _ArgentModule.contract.Transact(opts, "clearSession", _wallet)
}

// ClearSession is a paid mutator transaction binding the contract method 0xba821088.
//
// Solidity: function clearSession(address _wallet) returns()
func (_ArgentModule *ArgentModuleSession) ClearSession(_wallet common.Address) (*types.Transaction, error) {
	return _ArgentModule.Contract.ClearSession(&_ArgentModule.TransactOpts, _wallet)
}

// ClearSession is a paid mutator transaction binding the contract method 0xba821088.
//
// Solidity: function clearSession(address _wallet) returns()
func (_ArgentModule *ArgentModuleTransactorSession) ClearSession(_wallet common.Address) (*types.Transaction, error) {
	return _ArgentModule.Contract.ClearSession(&_ArgentModule.TransactOpts, _wallet)
}

// ConfirmGuardianAddition is a paid mutator transaction binding the contract method 0x70135f52.
//
// Solidity: function confirmGuardianAddition(address _wallet, address _guardian) returns()
func (_ArgentModule *ArgentModuleTransactor) ConfirmGuardianAddition(opts *bind.TransactOpts, _wallet common.Address, _guardian common.Address) (*types.Transaction, error) {
	return _ArgentModule.contract.Transact(opts, "confirmGuardianAddition", _wallet, _guardian)
}

// ConfirmGuardianAddition is a paid mutator transaction binding the contract method 0x70135f52.
//
// Solidity: function confirmGuardianAddition(address _wallet, address _guardian) returns()
func (_ArgentModule *ArgentModuleSession) ConfirmGuardianAddition(_wallet common.Address, _guardian common.Address) (*types.Transaction, error) {
	return _ArgentModule.Contract.ConfirmGuardianAddition(&_ArgentModule.TransactOpts, _wallet, _guardian)
}

// ConfirmGuardianAddition is a paid mutator transaction binding the contract method 0x70135f52.
//
// Solidity: function confirmGuardianAddition(address _wallet, address _guardian) returns()
func (_ArgentModule *ArgentModuleTransactorSession) ConfirmGuardianAddition(_wallet common.Address, _guardian common.Address) (*types.Transaction, error) {
	return _ArgentModule.Contract.ConfirmGuardianAddition(&_ArgentModule.TransactOpts, _wallet, _guardian)
}

// ConfirmGuardianRevokation is a paid mutator transaction binding the contract method 0x4b3ef054.
//
// Solidity: function confirmGuardianRevokation(address _wallet, address _guardian) returns()
func (_ArgentModule *ArgentModuleTransactor) ConfirmGuardianRevokation(opts *bind.TransactOpts, _wallet common.Address, _guardian common.Address) (*types.Transaction, error) {
	return _ArgentModule.contract.Transact(opts, "confirmGuardianRevokation", _wallet, _guardian)
}

// ConfirmGuardianRevokation is a paid mutator transaction binding the contract method 0x4b3ef054.
//
// Solidity: function confirmGuardianRevokation(address _wallet, address _guardian) returns()
func (_ArgentModule *ArgentModuleSession) ConfirmGuardianRevokation(_wallet common.Address, _guardian common.Address) (*types.Transaction, error) {
	return _ArgentModule.Contract.ConfirmGuardianRevokation(&_ArgentModule.TransactOpts, _wallet, _guardian)
}

// ConfirmGuardianRevokation is a paid mutator transaction binding the contract method 0x4b3ef054.
//
// Solidity: function confirmGuardianRevokation(address _wallet, address _guardian) returns()
func (_ArgentModule *ArgentModuleTransactorSession) ConfirmGuardianRevokation(_wallet common.Address, _guardian common.Address) (*types.Transaction, error) {
	return _ArgentModule.Contract.ConfirmGuardianRevokation(&_ArgentModule.TransactOpts, _wallet, _guardian)
}

// EnableERC1155TokenReceiver is a paid mutator transaction binding the contract method 0x59b4958a.
//
// Solidity: function enableERC1155TokenReceiver(address _wallet) returns()
func (_ArgentModule *ArgentModuleTransactor) EnableERC1155TokenReceiver(opts *bind.TransactOpts, _wallet common.Address) (*types.Transaction, error) {
	return _ArgentModule.contract.Transact(opts, "enableERC1155TokenReceiver", _wallet)
}

// EnableERC1155TokenReceiver is a paid mutator transaction binding the contract method 0x59b4958a.
//
// Solidity: function enableERC1155TokenReceiver(address _wallet) returns()
func (_ArgentModule *ArgentModuleSession) EnableERC1155TokenReceiver(_wallet common.Address) (*types.Transaction, error) {
	return _ArgentModule.Contract.EnableERC1155TokenReceiver(&_ArgentModule.TransactOpts, _wallet)
}

// EnableERC1155TokenReceiver is a paid mutator transaction binding the contract method 0x59b4958a.
//
// Solidity: function enableERC1155TokenReceiver(address _wallet) returns()
func (_ArgentModule *ArgentModuleTransactorSession) EnableERC1155TokenReceiver(_wallet common.Address) (*types.Transaction, error) {
	return _ArgentModule.Contract.EnableERC1155TokenReceiver(&_ArgentModule.TransactOpts, _wallet)
}

// Execute is a paid mutator transaction binding the contract method 0xe0724b6e.
//
// Solidity: function execute(address _wallet, bytes _data, uint256 _nonce, bytes _signatures, uint256 _gasPrice, uint256 _gasLimit, address _refundToken, address _refundAddress) returns(bool)
func (_ArgentModule *ArgentModuleTransactor) Execute(opts *bind.TransactOpts, _wallet common.Address, _data []byte, _nonce *big.Int, _signatures []byte, _gasPrice *big.Int, _gasLimit *big.Int, _refundToken common.Address, _refundAddress common.Address) (*types.Transaction, error) {
	return _ArgentModule.contract.Transact(opts, "execute", _wallet, _data, _nonce, _signatures, _gasPrice, _gasLimit, _refundToken, _refundAddress)
}

// Execute is a paid mutator transaction binding the contract method 0xe0724b6e.
//
// Solidity: function execute(address _wallet, bytes _data, uint256 _nonce, bytes _signatures, uint256 _gasPrice, uint256 _gasLimit, address _refundToken, address _refundAddress) returns(bool)
func (_ArgentModule *ArgentModuleSession) Execute(_wallet common.Address, _data []byte, _nonce *big.Int, _signatures []byte, _gasPrice *big.Int, _gasLimit *big.Int, _refundToken common.Address, _refundAddress common.Address) (*types.Transaction, error) {
	return _ArgentModule.Contract.Execute(&_ArgentModule.TransactOpts, _wallet, _data, _nonce, _signatures, _gasPrice, _gasLimit, _refundToken, _refundAddress)
}

// Execute is a paid mutator transaction binding the contract method 0xe0724b6e.
//
// Solidity: function execute(address _wallet, bytes _data, uint256 _nonce, bytes _signatures, uint256 _gasPrice, uint256 _gasLimit, address _refundToken, address _refundAddress) returns(bool)
func (_ArgentModule *ArgentModuleTransactorSession) Execute(_wallet common.Address, _data []byte, _nonce *big.Int, _signatures []byte, _gasPrice *big.Int, _gasLimit *big.Int, _refundToken common.Address, _refundAddress common.Address) (*types.Transaction, error) {
	return _ArgentModule.Contract.Execute(&_ArgentModule.TransactOpts, _wallet, _data, _nonce, _signatures, _gasPrice, _gasLimit, _refundToken, _refundAddress)
}

// ExecuteRecovery is a paid mutator transaction binding the contract method 0xb0ba4da0.
//
// Solidity: function executeRecovery(address _wallet, address _recovery) returns()
func (_ArgentModule *ArgentModuleTransactor) ExecuteRecovery(opts *bind.TransactOpts, _wallet common.Address, _recovery common.Address) (*types.Transaction, error) {
	return _ArgentModule.contract.Transact(opts, "executeRecovery", _wallet, _recovery)
}

// ExecuteRecovery is a paid mutator transaction binding the contract method 0xb0ba4da0.
//
// Solidity: function executeRecovery(address _wallet, address _recovery) returns()
func (_ArgentModule *ArgentModuleSession) ExecuteRecovery(_wallet common.Address, _recovery common.Address) (*types.Transaction, error) {
	return _ArgentModule.Contract.ExecuteRecovery(&_ArgentModule.TransactOpts, _wallet, _recovery)
}

// ExecuteRecovery is a paid mutator transaction binding the contract method 0xb0ba4da0.
//
// Solidity: function executeRecovery(address _wallet, address _recovery) returns()
func (_ArgentModule *ArgentModuleTransactorSession) ExecuteRecovery(_wallet common.Address, _recovery common.Address) (*types.Transaction, error) {
	return _ArgentModule.Contract.ExecuteRecovery(&_ArgentModule.TransactOpts, _wallet, _recovery)
}

// FinalizeRecovery is a paid mutator transaction binding the contract method 0x315a7af3.
//
// Solidity: function finalizeRecovery(address _wallet) returns()
func (_ArgentModule *ArgentModuleTransactor) FinalizeRecovery(opts *bind.TransactOpts, _wallet common.Address) (*types.Transaction, error) {
	return _ArgentModule.contract.Transact(opts, "finalizeRecovery", _wallet)
}

// FinalizeRecovery is a paid mutator transaction binding the contract method 0x315a7af3.
//
// Solidity: function finalizeRecovery(address _wallet) returns()
func (_ArgentModule *ArgentModuleSession) FinalizeRecovery(_wallet common.Address) (*types.Transaction, error) {
	return _ArgentModule.Contract.FinalizeRecovery(&_ArgentModule.TransactOpts, _wallet)
}

// FinalizeRecovery is a paid mutator transaction binding the contract method 0x315a7af3.
//
// Solidity: function finalizeRecovery(address _wallet) returns()
func (_ArgentModule *ArgentModuleTransactorSession) FinalizeRecovery(_wallet common.Address) (*types.Transaction, error) {
	return _ArgentModule.Contract.FinalizeRecovery(&_ArgentModule.TransactOpts, _wallet)
}

// Init is a paid mutator transaction binding the contract method 0x19ab453c.
//
// Solidity: function init(address _wallet) returns()
func (_ArgentModule *ArgentModuleTransactor) Init(opts *bind.TransactOpts, _wallet common.Address) (*types.Transaction, error) {
	return _ArgentModule.contract.Transact(opts, "init", _wallet)
}

// Init is a paid mutator transaction binding the contract method 0x19ab453c.
//
// Solidity: function init(address _wallet) returns()
func (_ArgentModule *ArgentModuleSession) Init(_wallet common.Address) (*types.Transaction, error) {
	return _ArgentModule.Contract.Init(&_ArgentModule.TransactOpts, _wallet)
}

// Init is a paid mutator transaction binding the contract method 0x19ab453c.
//
// Solidity: function init(address _wallet) returns()
func (_ArgentModule *ArgentModuleTransactorSession) Init(_wallet common.Address) (*types.Transaction, error) {
	return _ArgentModule.Contract.Init(&_ArgentModule.TransactOpts, _wallet)
}

// Lock is a paid mutator transaction binding the contract method 0xf435f5a7.
//
// Solidity: function lock(address _wallet) returns()
func (_ArgentModule *ArgentModuleTransactor) Lock(opts *bind.TransactOpts, _wallet common.Address) (*types.Transaction, error) {
	return _ArgentModule.contract.Transact(opts, "lock", _wallet)
}

// Lock is a paid mutator transaction binding the contract method 0xf435f5a7.
//
// Solidity: function lock(address _wallet) returns()
func (_ArgentModule *ArgentModuleSession) Lock(_wallet common.Address) (*types.Transaction, error) {
	return _ArgentModule.Contract.Lock(&_ArgentModule.TransactOpts, _wallet)
}

// Lock is a paid mutator transaction binding the contract method 0xf435f5a7.
//
// Solidity: function lock(address _wallet) returns()
func (_ArgentModule *ArgentModuleTransactorSession) Lock(_wallet common.Address) (*types.Transaction, error) {
	return _ArgentModule.Contract.Lock(&_ArgentModule.TransactOpts, _wallet)
}

// MultiCall is a paid mutator transaction binding the contract method 0xa5efb235.
//
// Solidity: function multiCall(address _wallet, (address,uint256,bytes)[] _transactions) returns(bytes[])
func (_ArgentModule *ArgentModuleTransactor) MultiCall(opts *bind.TransactOpts, _wallet common.Address, _transactions []TransactionManagerCall) (*types.Transaction, error) {
	return _ArgentModule.contract.Transact(opts, "multiCall", _wallet, _transactions)
}

// MultiCall is a paid mutator transaction binding the contract method 0xa5efb235.
//
// Solidity: function multiCall(address _wallet, (address,uint256,bytes)[] _transactions) returns(bytes[])
func (_ArgentModule *ArgentModuleSession) MultiCall(_wallet common.Address, _transactions []TransactionManagerCall) (*types.Transaction, error) {
	return _ArgentModule.Contract.MultiCall(&_ArgentModule.TransactOpts, _wallet, _transactions)
}

// MultiCall is a paid mutator transaction binding the contract method 0xa5efb235.
//
// Solidity: function multiCall(address _wallet, (address,uint256,bytes)[] _transactions) returns(bytes[])
func (_ArgentModule *ArgentModuleTransactorSession) MultiCall(_wallet common.Address, _transactions []TransactionManagerCall) (*types.Transaction, error) {
	return _ArgentModule.Contract.MultiCall(&_ArgentModule.TransactOpts, _wallet, _transactions)
}

// MultiCallWithGuardians is a paid mutator transaction binding the contract method 0xf143ddba.
//
// Solidity: function multiCallWithGuardians(address _wallet, (address,uint256,bytes)[] _transactions) returns(bytes[])
func (_ArgentModule *ArgentModuleTransactor) MultiCallWithGuardians(opts *bind.TransactOpts, _wallet common.Address, _transactions []TransactionManagerCall) (*types.Transaction, error) {
	return _ArgentModule.contract.Transact(opts, "multiCallWithGuardians", _wallet, _transactions)
}

// MultiCallWithGuardians is a paid mutator transaction binding the contract method 0xf143ddba.
//
// Solidity: function multiCallWithGuardians(address _wallet, (address,uint256,bytes)[] _transactions) returns(bytes[])
func (_ArgentModule *ArgentModuleSession) MultiCallWithGuardians(_wallet common.Address, _transactions []TransactionManagerCall) (*types.Transaction, error) {
	return _ArgentModule.Contract.MultiCallWithGuardians(&_ArgentModule.TransactOpts, _wallet, _transactions)
}

// MultiCallWithGuardians is a paid mutator transaction binding the contract method 0xf143ddba.
//
// Solidity: function multiCallWithGuardians(address _wallet, (address,uint256,bytes)[] _transactions) returns(bytes[])
func (_ArgentModule *ArgentModuleTransactorSession) MultiCallWithGuardians(_wallet common.Address, _transactions []TransactionManagerCall) (*types.Transaction, error) {
	return _ArgentModule.Contract.MultiCallWithGuardians(&_ArgentModule.TransactOpts, _wallet, _transactions)
}

// RecoverToken is a paid mutator transaction binding the contract method 0x9be65a60.
//
// Solidity: function recoverToken(address _token) returns()
func (_ArgentModule *ArgentModuleTransactor) RecoverToken(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _ArgentModule.contract.Transact(opts, "recoverToken", _token)
}

// RecoverToken is a paid mutator transaction binding the contract method 0x9be65a60.
//
// Solidity: function recoverToken(address _token) returns()
func (_ArgentModule *ArgentModuleSession) RecoverToken(_token common.Address) (*types.Transaction, error) {
	return _ArgentModule.Contract.RecoverToken(&_ArgentModule.TransactOpts, _token)
}

// RecoverToken is a paid mutator transaction binding the contract method 0x9be65a60.
//
// Solidity: function recoverToken(address _token) returns()
func (_ArgentModule *ArgentModuleTransactorSession) RecoverToken(_token common.Address) (*types.Transaction, error) {
	return _ArgentModule.Contract.RecoverToken(&_ArgentModule.TransactOpts, _token)
}

// RemoveFromWhitelist is a paid mutator transaction binding the contract method 0xf8d3277d.
//
// Solidity: function removeFromWhitelist(address _wallet, address _target) returns()
func (_ArgentModule *ArgentModuleTransactor) RemoveFromWhitelist(opts *bind.TransactOpts, _wallet common.Address, _target common.Address) (*types.Transaction, error) {
	return _ArgentModule.contract.Transact(opts, "removeFromWhitelist", _wallet, _target)
}

// RemoveFromWhitelist is a paid mutator transaction binding the contract method 0xf8d3277d.
//
// Solidity: function removeFromWhitelist(address _wallet, address _target) returns()
func (_ArgentModule *ArgentModuleSession) RemoveFromWhitelist(_wallet common.Address, _target common.Address) (*types.Transaction, error) {
	return _ArgentModule.Contract.RemoveFromWhitelist(&_ArgentModule.TransactOpts, _wallet, _target)
}

// RemoveFromWhitelist is a paid mutator transaction binding the contract method 0xf8d3277d.
//
// Solidity: function removeFromWhitelist(address _wallet, address _target) returns()
func (_ArgentModule *ArgentModuleTransactorSession) RemoveFromWhitelist(_wallet common.Address, _target common.Address) (*types.Transaction, error) {
	return _ArgentModule.Contract.RemoveFromWhitelist(&_ArgentModule.TransactOpts, _wallet, _target)
}

// RevokeGuardian is a paid mutator transaction binding the contract method 0x1d97d8cc.
//
// Solidity: function revokeGuardian(address _wallet, address _guardian) returns()
func (_ArgentModule *ArgentModuleTransactor) RevokeGuardian(opts *bind.TransactOpts, _wallet common.Address, _guardian common.Address) (*types.Transaction, error) {
	return _ArgentModule.contract.Transact(opts, "revokeGuardian", _wallet, _guardian)
}

// RevokeGuardian is a paid mutator transaction binding the contract method 0x1d97d8cc.
//
// Solidity: function revokeGuardian(address _wallet, address _guardian) returns()
func (_ArgentModule *ArgentModuleSession) RevokeGuardian(_wallet common.Address, _guardian common.Address) (*types.Transaction, error) {
	return _ArgentModule.Contract.RevokeGuardian(&_ArgentModule.TransactOpts, _wallet, _guardian)
}

// RevokeGuardian is a paid mutator transaction binding the contract method 0x1d97d8cc.
//
// Solidity: function revokeGuardian(address _wallet, address _guardian) returns()
func (_ArgentModule *ArgentModuleTransactorSession) RevokeGuardian(_wallet common.Address, _guardian common.Address) (*types.Transaction, error) {
	return _ArgentModule.Contract.RevokeGuardian(&_ArgentModule.TransactOpts, _wallet, _guardian)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0x6d435421.
//
// Solidity: function transferOwnership(address _wallet, address _newOwner) returns()
func (_ArgentModule *ArgentModuleTransactor) TransferOwnership(opts *bind.TransactOpts, _wallet common.Address, _newOwner common.Address) (*types.Transaction, error) {
	return _ArgentModule.contract.Transact(opts, "transferOwnership", _wallet, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0x6d435421.
//
// Solidity: function transferOwnership(address _wallet, address _newOwner) returns()
func (_ArgentModule *ArgentModuleSession) TransferOwnership(_wallet common.Address, _newOwner common.Address) (*types.Transaction, error) {
	return _ArgentModule.Contract.TransferOwnership(&_ArgentModule.TransactOpts, _wallet, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0x6d435421.
//
// Solidity: function transferOwnership(address _wallet, address _newOwner) returns()
func (_ArgentModule *ArgentModuleTransactorSession) TransferOwnership(_wallet common.Address, _newOwner common.Address) (*types.Transaction, error) {
	return _ArgentModule.Contract.TransferOwnership(&_ArgentModule.TransactOpts, _wallet, _newOwner)
}

// Unlock is a paid mutator transaction binding the contract method 0x2f6c493c.
//
// Solidity: function unlock(address _wallet) returns()
func (_ArgentModule *ArgentModuleTransactor) Unlock(opts *bind.TransactOpts, _wallet common.Address) (*types.Transaction, error) {
	return _ArgentModule.contract.Transact(opts, "unlock", _wallet)
}

// Unlock is a paid mutator transaction binding the contract method 0x2f6c493c.
//
// Solidity: function unlock(address _wallet) returns()
func (_ArgentModule *ArgentModuleSession) Unlock(_wallet common.Address) (*types.Transaction, error) {
	return _ArgentModule.Contract.Unlock(&_ArgentModule.TransactOpts, _wallet)
}

// Unlock is a paid mutator transaction binding the contract method 0x2f6c493c.
//
// Solidity: function unlock(address _wallet) returns()
func (_ArgentModule *ArgentModuleTransactorSession) Unlock(_wallet common.Address) (*types.Transaction, error) {
	return _ArgentModule.Contract.Unlock(&_ArgentModule.TransactOpts, _wallet)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_ArgentModule *ArgentModuleTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _ArgentModule.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_ArgentModule *ArgentModuleSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _ArgentModule.Contract.Fallback(&_ArgentModule.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_ArgentModule *ArgentModuleTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _ArgentModule.Contract.Fallback(&_ArgentModule.TransactOpts, calldata)
}

// ArgentModuleAddedToWhitelistIterator is returned from FilterAddedToWhitelist and is used to iterate over the raw logs and unpacked data for AddedToWhitelist events raised by the ArgentModule contract.
type ArgentModuleAddedToWhitelistIterator struct {
	Event *ArgentModuleAddedToWhitelist // Event containing the contract specifics and raw log

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
func (it *ArgentModuleAddedToWhitelistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArgentModuleAddedToWhitelist)
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
		it.Event = new(ArgentModuleAddedToWhitelist)
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
func (it *ArgentModuleAddedToWhitelistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArgentModuleAddedToWhitelistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArgentModuleAddedToWhitelist represents a AddedToWhitelist event raised by the ArgentModule contract.
type ArgentModuleAddedToWhitelist struct {
	Wallet         common.Address
	Target         common.Address
	WhitelistAfter uint64
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterAddedToWhitelist is a free log retrieval operation binding the contract event 0x1f57f9641d3e8733ed672fef5ac85464bd7215ef2f21e83428e8408248b13dcd.
//
// Solidity: event AddedToWhitelist(address indexed wallet, address indexed target, uint64 whitelistAfter)
func (_ArgentModule *ArgentModuleFilterer) FilterAddedToWhitelist(opts *bind.FilterOpts, wallet []common.Address, target []common.Address) (*ArgentModuleAddedToWhitelistIterator, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _ArgentModule.contract.FilterLogs(opts, "AddedToWhitelist", walletRule, targetRule)
	if err != nil {
		return nil, err
	}
	return &ArgentModuleAddedToWhitelistIterator{contract: _ArgentModule.contract, event: "AddedToWhitelist", logs: logs, sub: sub}, nil
}

// WatchAddedToWhitelist is a free log subscription operation binding the contract event 0x1f57f9641d3e8733ed672fef5ac85464bd7215ef2f21e83428e8408248b13dcd.
//
// Solidity: event AddedToWhitelist(address indexed wallet, address indexed target, uint64 whitelistAfter)
func (_ArgentModule *ArgentModuleFilterer) WatchAddedToWhitelist(opts *bind.WatchOpts, sink chan<- *ArgentModuleAddedToWhitelist, wallet []common.Address, target []common.Address) (event.Subscription, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _ArgentModule.contract.WatchLogs(opts, "AddedToWhitelist", walletRule, targetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArgentModuleAddedToWhitelist)
				if err := _ArgentModule.contract.UnpackLog(event, "AddedToWhitelist", log); err != nil {
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

// ParseAddedToWhitelist is a log parse operation binding the contract event 0x1f57f9641d3e8733ed672fef5ac85464bd7215ef2f21e83428e8408248b13dcd.
//
// Solidity: event AddedToWhitelist(address indexed wallet, address indexed target, uint64 whitelistAfter)
func (_ArgentModule *ArgentModuleFilterer) ParseAddedToWhitelist(log types.Log) (*ArgentModuleAddedToWhitelist, error) {
	event := new(ArgentModuleAddedToWhitelist)
	if err := _ArgentModule.contract.UnpackLog(event, "AddedToWhitelist", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArgentModuleGuardianAddedIterator is returned from FilterGuardianAdded and is used to iterate over the raw logs and unpacked data for GuardianAdded events raised by the ArgentModule contract.
type ArgentModuleGuardianAddedIterator struct {
	Event *ArgentModuleGuardianAdded // Event containing the contract specifics and raw log

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
func (it *ArgentModuleGuardianAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArgentModuleGuardianAdded)
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
		it.Event = new(ArgentModuleGuardianAdded)
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
func (it *ArgentModuleGuardianAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArgentModuleGuardianAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArgentModuleGuardianAdded represents a GuardianAdded event raised by the ArgentModule contract.
type ArgentModuleGuardianAdded struct {
	Wallet   common.Address
	Guardian common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterGuardianAdded is a free log retrieval operation binding the contract event 0xbc3292102fa77e083913064b282926717cdfaede4d35f553d66366c0a3da755a.
//
// Solidity: event GuardianAdded(address indexed wallet, address indexed guardian)
func (_ArgentModule *ArgentModuleFilterer) FilterGuardianAdded(opts *bind.FilterOpts, wallet []common.Address, guardian []common.Address) (*ArgentModuleGuardianAddedIterator, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}
	var guardianRule []interface{}
	for _, guardianItem := range guardian {
		guardianRule = append(guardianRule, guardianItem)
	}

	logs, sub, err := _ArgentModule.contract.FilterLogs(opts, "GuardianAdded", walletRule, guardianRule)
	if err != nil {
		return nil, err
	}
	return &ArgentModuleGuardianAddedIterator{contract: _ArgentModule.contract, event: "GuardianAdded", logs: logs, sub: sub}, nil
}

// WatchGuardianAdded is a free log subscription operation binding the contract event 0xbc3292102fa77e083913064b282926717cdfaede4d35f553d66366c0a3da755a.
//
// Solidity: event GuardianAdded(address indexed wallet, address indexed guardian)
func (_ArgentModule *ArgentModuleFilterer) WatchGuardianAdded(opts *bind.WatchOpts, sink chan<- *ArgentModuleGuardianAdded, wallet []common.Address, guardian []common.Address) (event.Subscription, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}
	var guardianRule []interface{}
	for _, guardianItem := range guardian {
		guardianRule = append(guardianRule, guardianItem)
	}

	logs, sub, err := _ArgentModule.contract.WatchLogs(opts, "GuardianAdded", walletRule, guardianRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArgentModuleGuardianAdded)
				if err := _ArgentModule.contract.UnpackLog(event, "GuardianAdded", log); err != nil {
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

// ParseGuardianAdded is a log parse operation binding the contract event 0xbc3292102fa77e083913064b282926717cdfaede4d35f553d66366c0a3da755a.
//
// Solidity: event GuardianAdded(address indexed wallet, address indexed guardian)
func (_ArgentModule *ArgentModuleFilterer) ParseGuardianAdded(log types.Log) (*ArgentModuleGuardianAdded, error) {
	event := new(ArgentModuleGuardianAdded)
	if err := _ArgentModule.contract.UnpackLog(event, "GuardianAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArgentModuleGuardianAdditionCancelledIterator is returned from FilterGuardianAdditionCancelled and is used to iterate over the raw logs and unpacked data for GuardianAdditionCancelled events raised by the ArgentModule contract.
type ArgentModuleGuardianAdditionCancelledIterator struct {
	Event *ArgentModuleGuardianAdditionCancelled // Event containing the contract specifics and raw log

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
func (it *ArgentModuleGuardianAdditionCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArgentModuleGuardianAdditionCancelled)
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
		it.Event = new(ArgentModuleGuardianAdditionCancelled)
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
func (it *ArgentModuleGuardianAdditionCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArgentModuleGuardianAdditionCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArgentModuleGuardianAdditionCancelled represents a GuardianAdditionCancelled event raised by the ArgentModule contract.
type ArgentModuleGuardianAdditionCancelled struct {
	Wallet   common.Address
	Guardian common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterGuardianAdditionCancelled is a free log retrieval operation binding the contract event 0xaa13b27c23e9e3f3d5f3861a53b7a2931e019170a6a19ed64942e26a1dd5987a.
//
// Solidity: event GuardianAdditionCancelled(address indexed wallet, address indexed guardian)
func (_ArgentModule *ArgentModuleFilterer) FilterGuardianAdditionCancelled(opts *bind.FilterOpts, wallet []common.Address, guardian []common.Address) (*ArgentModuleGuardianAdditionCancelledIterator, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}
	var guardianRule []interface{}
	for _, guardianItem := range guardian {
		guardianRule = append(guardianRule, guardianItem)
	}

	logs, sub, err := _ArgentModule.contract.FilterLogs(opts, "GuardianAdditionCancelled", walletRule, guardianRule)
	if err != nil {
		return nil, err
	}
	return &ArgentModuleGuardianAdditionCancelledIterator{contract: _ArgentModule.contract, event: "GuardianAdditionCancelled", logs: logs, sub: sub}, nil
}

// WatchGuardianAdditionCancelled is a free log subscription operation binding the contract event 0xaa13b27c23e9e3f3d5f3861a53b7a2931e019170a6a19ed64942e26a1dd5987a.
//
// Solidity: event GuardianAdditionCancelled(address indexed wallet, address indexed guardian)
func (_ArgentModule *ArgentModuleFilterer) WatchGuardianAdditionCancelled(opts *bind.WatchOpts, sink chan<- *ArgentModuleGuardianAdditionCancelled, wallet []common.Address, guardian []common.Address) (event.Subscription, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}
	var guardianRule []interface{}
	for _, guardianItem := range guardian {
		guardianRule = append(guardianRule, guardianItem)
	}

	logs, sub, err := _ArgentModule.contract.WatchLogs(opts, "GuardianAdditionCancelled", walletRule, guardianRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArgentModuleGuardianAdditionCancelled)
				if err := _ArgentModule.contract.UnpackLog(event, "GuardianAdditionCancelled", log); err != nil {
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

// ParseGuardianAdditionCancelled is a log parse operation binding the contract event 0xaa13b27c23e9e3f3d5f3861a53b7a2931e019170a6a19ed64942e26a1dd5987a.
//
// Solidity: event GuardianAdditionCancelled(address indexed wallet, address indexed guardian)
func (_ArgentModule *ArgentModuleFilterer) ParseGuardianAdditionCancelled(log types.Log) (*ArgentModuleGuardianAdditionCancelled, error) {
	event := new(ArgentModuleGuardianAdditionCancelled)
	if err := _ArgentModule.contract.UnpackLog(event, "GuardianAdditionCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArgentModuleGuardianAdditionRequestedIterator is returned from FilterGuardianAdditionRequested and is used to iterate over the raw logs and unpacked data for GuardianAdditionRequested events raised by the ArgentModule contract.
type ArgentModuleGuardianAdditionRequestedIterator struct {
	Event *ArgentModuleGuardianAdditionRequested // Event containing the contract specifics and raw log

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
func (it *ArgentModuleGuardianAdditionRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArgentModuleGuardianAdditionRequested)
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
		it.Event = new(ArgentModuleGuardianAdditionRequested)
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
func (it *ArgentModuleGuardianAdditionRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArgentModuleGuardianAdditionRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArgentModuleGuardianAdditionRequested represents a GuardianAdditionRequested event raised by the ArgentModule contract.
type ArgentModuleGuardianAdditionRequested struct {
	Wallet       common.Address
	Guardian     common.Address
	ExecuteAfter *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterGuardianAdditionRequested is a free log retrieval operation binding the contract event 0xe4166e4bc55a182bd13d933553241bb3441b91d15fbc74c5c752f96965563bde.
//
// Solidity: event GuardianAdditionRequested(address indexed wallet, address indexed guardian, uint256 executeAfter)
func (_ArgentModule *ArgentModuleFilterer) FilterGuardianAdditionRequested(opts *bind.FilterOpts, wallet []common.Address, guardian []common.Address) (*ArgentModuleGuardianAdditionRequestedIterator, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}
	var guardianRule []interface{}
	for _, guardianItem := range guardian {
		guardianRule = append(guardianRule, guardianItem)
	}

	logs, sub, err := _ArgentModule.contract.FilterLogs(opts, "GuardianAdditionRequested", walletRule, guardianRule)
	if err != nil {
		return nil, err
	}
	return &ArgentModuleGuardianAdditionRequestedIterator{contract: _ArgentModule.contract, event: "GuardianAdditionRequested", logs: logs, sub: sub}, nil
}

// WatchGuardianAdditionRequested is a free log subscription operation binding the contract event 0xe4166e4bc55a182bd13d933553241bb3441b91d15fbc74c5c752f96965563bde.
//
// Solidity: event GuardianAdditionRequested(address indexed wallet, address indexed guardian, uint256 executeAfter)
func (_ArgentModule *ArgentModuleFilterer) WatchGuardianAdditionRequested(opts *bind.WatchOpts, sink chan<- *ArgentModuleGuardianAdditionRequested, wallet []common.Address, guardian []common.Address) (event.Subscription, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}
	var guardianRule []interface{}
	for _, guardianItem := range guardian {
		guardianRule = append(guardianRule, guardianItem)
	}

	logs, sub, err := _ArgentModule.contract.WatchLogs(opts, "GuardianAdditionRequested", walletRule, guardianRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArgentModuleGuardianAdditionRequested)
				if err := _ArgentModule.contract.UnpackLog(event, "GuardianAdditionRequested", log); err != nil {
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

// ParseGuardianAdditionRequested is a log parse operation binding the contract event 0xe4166e4bc55a182bd13d933553241bb3441b91d15fbc74c5c752f96965563bde.
//
// Solidity: event GuardianAdditionRequested(address indexed wallet, address indexed guardian, uint256 executeAfter)
func (_ArgentModule *ArgentModuleFilterer) ParseGuardianAdditionRequested(log types.Log) (*ArgentModuleGuardianAdditionRequested, error) {
	event := new(ArgentModuleGuardianAdditionRequested)
	if err := _ArgentModule.contract.UnpackLog(event, "GuardianAdditionRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArgentModuleGuardianRevokationCancelledIterator is returned from FilterGuardianRevokationCancelled and is used to iterate over the raw logs and unpacked data for GuardianRevokationCancelled events raised by the ArgentModule contract.
type ArgentModuleGuardianRevokationCancelledIterator struct {
	Event *ArgentModuleGuardianRevokationCancelled // Event containing the contract specifics and raw log

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
func (it *ArgentModuleGuardianRevokationCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArgentModuleGuardianRevokationCancelled)
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
		it.Event = new(ArgentModuleGuardianRevokationCancelled)
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
func (it *ArgentModuleGuardianRevokationCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArgentModuleGuardianRevokationCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArgentModuleGuardianRevokationCancelled represents a GuardianRevokationCancelled event raised by the ArgentModule contract.
type ArgentModuleGuardianRevokationCancelled struct {
	Wallet   common.Address
	Guardian common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterGuardianRevokationCancelled is a free log retrieval operation binding the contract event 0xc0b205956d5e27c296695de329b5a014584a4f51824b1725a0eefc1174d6dbd5.
//
// Solidity: event GuardianRevokationCancelled(address indexed wallet, address indexed guardian)
func (_ArgentModule *ArgentModuleFilterer) FilterGuardianRevokationCancelled(opts *bind.FilterOpts, wallet []common.Address, guardian []common.Address) (*ArgentModuleGuardianRevokationCancelledIterator, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}
	var guardianRule []interface{}
	for _, guardianItem := range guardian {
		guardianRule = append(guardianRule, guardianItem)
	}

	logs, sub, err := _ArgentModule.contract.FilterLogs(opts, "GuardianRevokationCancelled", walletRule, guardianRule)
	if err != nil {
		return nil, err
	}
	return &ArgentModuleGuardianRevokationCancelledIterator{contract: _ArgentModule.contract, event: "GuardianRevokationCancelled", logs: logs, sub: sub}, nil
}

// WatchGuardianRevokationCancelled is a free log subscription operation binding the contract event 0xc0b205956d5e27c296695de329b5a014584a4f51824b1725a0eefc1174d6dbd5.
//
// Solidity: event GuardianRevokationCancelled(address indexed wallet, address indexed guardian)
func (_ArgentModule *ArgentModuleFilterer) WatchGuardianRevokationCancelled(opts *bind.WatchOpts, sink chan<- *ArgentModuleGuardianRevokationCancelled, wallet []common.Address, guardian []common.Address) (event.Subscription, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}
	var guardianRule []interface{}
	for _, guardianItem := range guardian {
		guardianRule = append(guardianRule, guardianItem)
	}

	logs, sub, err := _ArgentModule.contract.WatchLogs(opts, "GuardianRevokationCancelled", walletRule, guardianRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArgentModuleGuardianRevokationCancelled)
				if err := _ArgentModule.contract.UnpackLog(event, "GuardianRevokationCancelled", log); err != nil {
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

// ParseGuardianRevokationCancelled is a log parse operation binding the contract event 0xc0b205956d5e27c296695de329b5a014584a4f51824b1725a0eefc1174d6dbd5.
//
// Solidity: event GuardianRevokationCancelled(address indexed wallet, address indexed guardian)
func (_ArgentModule *ArgentModuleFilterer) ParseGuardianRevokationCancelled(log types.Log) (*ArgentModuleGuardianRevokationCancelled, error) {
	event := new(ArgentModuleGuardianRevokationCancelled)
	if err := _ArgentModule.contract.UnpackLog(event, "GuardianRevokationCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArgentModuleGuardianRevokationRequestedIterator is returned from FilterGuardianRevokationRequested and is used to iterate over the raw logs and unpacked data for GuardianRevokationRequested events raised by the ArgentModule contract.
type ArgentModuleGuardianRevokationRequestedIterator struct {
	Event *ArgentModuleGuardianRevokationRequested // Event containing the contract specifics and raw log

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
func (it *ArgentModuleGuardianRevokationRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArgentModuleGuardianRevokationRequested)
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
		it.Event = new(ArgentModuleGuardianRevokationRequested)
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
func (it *ArgentModuleGuardianRevokationRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArgentModuleGuardianRevokationRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArgentModuleGuardianRevokationRequested represents a GuardianRevokationRequested event raised by the ArgentModule contract.
type ArgentModuleGuardianRevokationRequested struct {
	Wallet       common.Address
	Guardian     common.Address
	ExecuteAfter *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterGuardianRevokationRequested is a free log retrieval operation binding the contract event 0x9746f6868f544595794833da53250bd19e72334733336cfd5dd6fbc5f6a6ac42.
//
// Solidity: event GuardianRevokationRequested(address indexed wallet, address indexed guardian, uint256 executeAfter)
func (_ArgentModule *ArgentModuleFilterer) FilterGuardianRevokationRequested(opts *bind.FilterOpts, wallet []common.Address, guardian []common.Address) (*ArgentModuleGuardianRevokationRequestedIterator, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}
	var guardianRule []interface{}
	for _, guardianItem := range guardian {
		guardianRule = append(guardianRule, guardianItem)
	}

	logs, sub, err := _ArgentModule.contract.FilterLogs(opts, "GuardianRevokationRequested", walletRule, guardianRule)
	if err != nil {
		return nil, err
	}
	return &ArgentModuleGuardianRevokationRequestedIterator{contract: _ArgentModule.contract, event: "GuardianRevokationRequested", logs: logs, sub: sub}, nil
}

// WatchGuardianRevokationRequested is a free log subscription operation binding the contract event 0x9746f6868f544595794833da53250bd19e72334733336cfd5dd6fbc5f6a6ac42.
//
// Solidity: event GuardianRevokationRequested(address indexed wallet, address indexed guardian, uint256 executeAfter)
func (_ArgentModule *ArgentModuleFilterer) WatchGuardianRevokationRequested(opts *bind.WatchOpts, sink chan<- *ArgentModuleGuardianRevokationRequested, wallet []common.Address, guardian []common.Address) (event.Subscription, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}
	var guardianRule []interface{}
	for _, guardianItem := range guardian {
		guardianRule = append(guardianRule, guardianItem)
	}

	logs, sub, err := _ArgentModule.contract.WatchLogs(opts, "GuardianRevokationRequested", walletRule, guardianRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArgentModuleGuardianRevokationRequested)
				if err := _ArgentModule.contract.UnpackLog(event, "GuardianRevokationRequested", log); err != nil {
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

// ParseGuardianRevokationRequested is a log parse operation binding the contract event 0x9746f6868f544595794833da53250bd19e72334733336cfd5dd6fbc5f6a6ac42.
//
// Solidity: event GuardianRevokationRequested(address indexed wallet, address indexed guardian, uint256 executeAfter)
func (_ArgentModule *ArgentModuleFilterer) ParseGuardianRevokationRequested(log types.Log) (*ArgentModuleGuardianRevokationRequested, error) {
	event := new(ArgentModuleGuardianRevokationRequested)
	if err := _ArgentModule.contract.UnpackLog(event, "GuardianRevokationRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArgentModuleGuardianRevokedIterator is returned from FilterGuardianRevoked and is used to iterate over the raw logs and unpacked data for GuardianRevoked events raised by the ArgentModule contract.
type ArgentModuleGuardianRevokedIterator struct {
	Event *ArgentModuleGuardianRevoked // Event containing the contract specifics and raw log

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
func (it *ArgentModuleGuardianRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArgentModuleGuardianRevoked)
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
		it.Event = new(ArgentModuleGuardianRevoked)
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
func (it *ArgentModuleGuardianRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArgentModuleGuardianRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArgentModuleGuardianRevoked represents a GuardianRevoked event raised by the ArgentModule contract.
type ArgentModuleGuardianRevoked struct {
	Wallet   common.Address
	Guardian common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterGuardianRevoked is a free log retrieval operation binding the contract event 0x548f10dcba266544123ad8cf8284f25c4baa659cba25dbdf16a06ea11235de9b.
//
// Solidity: event GuardianRevoked(address indexed wallet, address indexed guardian)
func (_ArgentModule *ArgentModuleFilterer) FilterGuardianRevoked(opts *bind.FilterOpts, wallet []common.Address, guardian []common.Address) (*ArgentModuleGuardianRevokedIterator, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}
	var guardianRule []interface{}
	for _, guardianItem := range guardian {
		guardianRule = append(guardianRule, guardianItem)
	}

	logs, sub, err := _ArgentModule.contract.FilterLogs(opts, "GuardianRevoked", walletRule, guardianRule)
	if err != nil {
		return nil, err
	}
	return &ArgentModuleGuardianRevokedIterator{contract: _ArgentModule.contract, event: "GuardianRevoked", logs: logs, sub: sub}, nil
}

// WatchGuardianRevoked is a free log subscription operation binding the contract event 0x548f10dcba266544123ad8cf8284f25c4baa659cba25dbdf16a06ea11235de9b.
//
// Solidity: event GuardianRevoked(address indexed wallet, address indexed guardian)
func (_ArgentModule *ArgentModuleFilterer) WatchGuardianRevoked(opts *bind.WatchOpts, sink chan<- *ArgentModuleGuardianRevoked, wallet []common.Address, guardian []common.Address) (event.Subscription, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}
	var guardianRule []interface{}
	for _, guardianItem := range guardian {
		guardianRule = append(guardianRule, guardianItem)
	}

	logs, sub, err := _ArgentModule.contract.WatchLogs(opts, "GuardianRevoked", walletRule, guardianRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArgentModuleGuardianRevoked)
				if err := _ArgentModule.contract.UnpackLog(event, "GuardianRevoked", log); err != nil {
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

// ParseGuardianRevoked is a log parse operation binding the contract event 0x548f10dcba266544123ad8cf8284f25c4baa659cba25dbdf16a06ea11235de9b.
//
// Solidity: event GuardianRevoked(address indexed wallet, address indexed guardian)
func (_ArgentModule *ArgentModuleFilterer) ParseGuardianRevoked(log types.Log) (*ArgentModuleGuardianRevoked, error) {
	event := new(ArgentModuleGuardianRevoked)
	if err := _ArgentModule.contract.UnpackLog(event, "GuardianRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArgentModuleLockedIterator is returned from FilterLocked and is used to iterate over the raw logs and unpacked data for Locked events raised by the ArgentModule contract.
type ArgentModuleLockedIterator struct {
	Event *ArgentModuleLocked // Event containing the contract specifics and raw log

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
func (it *ArgentModuleLockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArgentModuleLocked)
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
		it.Event = new(ArgentModuleLocked)
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
func (it *ArgentModuleLockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArgentModuleLockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArgentModuleLocked represents a Locked event raised by the ArgentModule contract.
type ArgentModuleLocked struct {
	Wallet       common.Address
	ReleaseAfter uint64
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterLocked is a free log retrieval operation binding the contract event 0x6395bace6e0acbe4f22761b149d3cc2e88c7dde6bf4d8481825eef404cf989a1.
//
// Solidity: event Locked(address indexed wallet, uint64 releaseAfter)
func (_ArgentModule *ArgentModuleFilterer) FilterLocked(opts *bind.FilterOpts, wallet []common.Address) (*ArgentModuleLockedIterator, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _ArgentModule.contract.FilterLogs(opts, "Locked", walletRule)
	if err != nil {
		return nil, err
	}
	return &ArgentModuleLockedIterator{contract: _ArgentModule.contract, event: "Locked", logs: logs, sub: sub}, nil
}

// WatchLocked is a free log subscription operation binding the contract event 0x6395bace6e0acbe4f22761b149d3cc2e88c7dde6bf4d8481825eef404cf989a1.
//
// Solidity: event Locked(address indexed wallet, uint64 releaseAfter)
func (_ArgentModule *ArgentModuleFilterer) WatchLocked(opts *bind.WatchOpts, sink chan<- *ArgentModuleLocked, wallet []common.Address) (event.Subscription, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _ArgentModule.contract.WatchLogs(opts, "Locked", walletRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArgentModuleLocked)
				if err := _ArgentModule.contract.UnpackLog(event, "Locked", log); err != nil {
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

// ParseLocked is a log parse operation binding the contract event 0x6395bace6e0acbe4f22761b149d3cc2e88c7dde6bf4d8481825eef404cf989a1.
//
// Solidity: event Locked(address indexed wallet, uint64 releaseAfter)
func (_ArgentModule *ArgentModuleFilterer) ParseLocked(log types.Log) (*ArgentModuleLocked, error) {
	event := new(ArgentModuleLocked)
	if err := _ArgentModule.contract.UnpackLog(event, "Locked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArgentModuleModuleCreatedIterator is returned from FilterModuleCreated and is used to iterate over the raw logs and unpacked data for ModuleCreated events raised by the ArgentModule contract.
type ArgentModuleModuleCreatedIterator struct {
	Event *ArgentModuleModuleCreated // Event containing the contract specifics and raw log

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
func (it *ArgentModuleModuleCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArgentModuleModuleCreated)
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
		it.Event = new(ArgentModuleModuleCreated)
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
func (it *ArgentModuleModuleCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArgentModuleModuleCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArgentModuleModuleCreated represents a ModuleCreated event raised by the ArgentModule contract.
type ArgentModuleModuleCreated struct {
	Name [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterModuleCreated is a free log retrieval operation binding the contract event 0x3019c8fc80239e3dff8f781212ae2004839c2cb61d6c70acd279ac65392145df.
//
// Solidity: event ModuleCreated(bytes32 name)
func (_ArgentModule *ArgentModuleFilterer) FilterModuleCreated(opts *bind.FilterOpts) (*ArgentModuleModuleCreatedIterator, error) {

	logs, sub, err := _ArgentModule.contract.FilterLogs(opts, "ModuleCreated")
	if err != nil {
		return nil, err
	}
	return &ArgentModuleModuleCreatedIterator{contract: _ArgentModule.contract, event: "ModuleCreated", logs: logs, sub: sub}, nil
}

// WatchModuleCreated is a free log subscription operation binding the contract event 0x3019c8fc80239e3dff8f781212ae2004839c2cb61d6c70acd279ac65392145df.
//
// Solidity: event ModuleCreated(bytes32 name)
func (_ArgentModule *ArgentModuleFilterer) WatchModuleCreated(opts *bind.WatchOpts, sink chan<- *ArgentModuleModuleCreated) (event.Subscription, error) {

	logs, sub, err := _ArgentModule.contract.WatchLogs(opts, "ModuleCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArgentModuleModuleCreated)
				if err := _ArgentModule.contract.UnpackLog(event, "ModuleCreated", log); err != nil {
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

// ParseModuleCreated is a log parse operation binding the contract event 0x3019c8fc80239e3dff8f781212ae2004839c2cb61d6c70acd279ac65392145df.
//
// Solidity: event ModuleCreated(bytes32 name)
func (_ArgentModule *ArgentModuleFilterer) ParseModuleCreated(log types.Log) (*ArgentModuleModuleCreated, error) {
	event := new(ArgentModuleModuleCreated)
	if err := _ArgentModule.contract.UnpackLog(event, "ModuleCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArgentModuleOwnershipTransferedIterator is returned from FilterOwnershipTransfered and is used to iterate over the raw logs and unpacked data for OwnershipTransfered events raised by the ArgentModule contract.
type ArgentModuleOwnershipTransferedIterator struct {
	Event *ArgentModuleOwnershipTransfered // Event containing the contract specifics and raw log

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
func (it *ArgentModuleOwnershipTransferedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArgentModuleOwnershipTransfered)
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
		it.Event = new(ArgentModuleOwnershipTransfered)
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
func (it *ArgentModuleOwnershipTransferedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArgentModuleOwnershipTransferedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArgentModuleOwnershipTransfered represents a OwnershipTransfered event raised by the ArgentModule contract.
type ArgentModuleOwnershipTransfered struct {
	Wallet   common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransfered is a free log retrieval operation binding the contract event 0x0d18b5fd22306e373229b9439188228edca81207d1667f604daf6cef8aa3ee67.
//
// Solidity: event OwnershipTransfered(address indexed wallet, address indexed _newOwner)
func (_ArgentModule *ArgentModuleFilterer) FilterOwnershipTransfered(opts *bind.FilterOpts, wallet []common.Address, _newOwner []common.Address) (*ArgentModuleOwnershipTransferedIterator, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}
	var _newOwnerRule []interface{}
	for _, _newOwnerItem := range _newOwner {
		_newOwnerRule = append(_newOwnerRule, _newOwnerItem)
	}

	logs, sub, err := _ArgentModule.contract.FilterLogs(opts, "OwnershipTransfered", walletRule, _newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ArgentModuleOwnershipTransferedIterator{contract: _ArgentModule.contract, event: "OwnershipTransfered", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransfered is a free log subscription operation binding the contract event 0x0d18b5fd22306e373229b9439188228edca81207d1667f604daf6cef8aa3ee67.
//
// Solidity: event OwnershipTransfered(address indexed wallet, address indexed _newOwner)
func (_ArgentModule *ArgentModuleFilterer) WatchOwnershipTransfered(opts *bind.WatchOpts, sink chan<- *ArgentModuleOwnershipTransfered, wallet []common.Address, _newOwner []common.Address) (event.Subscription, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}
	var _newOwnerRule []interface{}
	for _, _newOwnerItem := range _newOwner {
		_newOwnerRule = append(_newOwnerRule, _newOwnerItem)
	}

	logs, sub, err := _ArgentModule.contract.WatchLogs(opts, "OwnershipTransfered", walletRule, _newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArgentModuleOwnershipTransfered)
				if err := _ArgentModule.contract.UnpackLog(event, "OwnershipTransfered", log); err != nil {
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

// ParseOwnershipTransfered is a log parse operation binding the contract event 0x0d18b5fd22306e373229b9439188228edca81207d1667f604daf6cef8aa3ee67.
//
// Solidity: event OwnershipTransfered(address indexed wallet, address indexed _newOwner)
func (_ArgentModule *ArgentModuleFilterer) ParseOwnershipTransfered(log types.Log) (*ArgentModuleOwnershipTransfered, error) {
	event := new(ArgentModuleOwnershipTransfered)
	if err := _ArgentModule.contract.UnpackLog(event, "OwnershipTransfered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArgentModuleRecoveryCanceledIterator is returned from FilterRecoveryCanceled and is used to iterate over the raw logs and unpacked data for RecoveryCanceled events raised by the ArgentModule contract.
type ArgentModuleRecoveryCanceledIterator struct {
	Event *ArgentModuleRecoveryCanceled // Event containing the contract specifics and raw log

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
func (it *ArgentModuleRecoveryCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArgentModuleRecoveryCanceled)
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
		it.Event = new(ArgentModuleRecoveryCanceled)
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
func (it *ArgentModuleRecoveryCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArgentModuleRecoveryCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArgentModuleRecoveryCanceled represents a RecoveryCanceled event raised by the ArgentModule contract.
type ArgentModuleRecoveryCanceled struct {
	Wallet   common.Address
	Recovery common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRecoveryCanceled is a free log retrieval operation binding the contract event 0xc45926607303da71dbeffd2ed5c6b00f581982586b697655d19ae4c4d558f259.
//
// Solidity: event RecoveryCanceled(address indexed wallet, address indexed _recovery)
func (_ArgentModule *ArgentModuleFilterer) FilterRecoveryCanceled(opts *bind.FilterOpts, wallet []common.Address, _recovery []common.Address) (*ArgentModuleRecoveryCanceledIterator, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}
	var _recoveryRule []interface{}
	for _, _recoveryItem := range _recovery {
		_recoveryRule = append(_recoveryRule, _recoveryItem)
	}

	logs, sub, err := _ArgentModule.contract.FilterLogs(opts, "RecoveryCanceled", walletRule, _recoveryRule)
	if err != nil {
		return nil, err
	}
	return &ArgentModuleRecoveryCanceledIterator{contract: _ArgentModule.contract, event: "RecoveryCanceled", logs: logs, sub: sub}, nil
}

// WatchRecoveryCanceled is a free log subscription operation binding the contract event 0xc45926607303da71dbeffd2ed5c6b00f581982586b697655d19ae4c4d558f259.
//
// Solidity: event RecoveryCanceled(address indexed wallet, address indexed _recovery)
func (_ArgentModule *ArgentModuleFilterer) WatchRecoveryCanceled(opts *bind.WatchOpts, sink chan<- *ArgentModuleRecoveryCanceled, wallet []common.Address, _recovery []common.Address) (event.Subscription, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}
	var _recoveryRule []interface{}
	for _, _recoveryItem := range _recovery {
		_recoveryRule = append(_recoveryRule, _recoveryItem)
	}

	logs, sub, err := _ArgentModule.contract.WatchLogs(opts, "RecoveryCanceled", walletRule, _recoveryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArgentModuleRecoveryCanceled)
				if err := _ArgentModule.contract.UnpackLog(event, "RecoveryCanceled", log); err != nil {
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

// ParseRecoveryCanceled is a log parse operation binding the contract event 0xc45926607303da71dbeffd2ed5c6b00f581982586b697655d19ae4c4d558f259.
//
// Solidity: event RecoveryCanceled(address indexed wallet, address indexed _recovery)
func (_ArgentModule *ArgentModuleFilterer) ParseRecoveryCanceled(log types.Log) (*ArgentModuleRecoveryCanceled, error) {
	event := new(ArgentModuleRecoveryCanceled)
	if err := _ArgentModule.contract.UnpackLog(event, "RecoveryCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArgentModuleRecoveryExecutedIterator is returned from FilterRecoveryExecuted and is used to iterate over the raw logs and unpacked data for RecoveryExecuted events raised by the ArgentModule contract.
type ArgentModuleRecoveryExecutedIterator struct {
	Event *ArgentModuleRecoveryExecuted // Event containing the contract specifics and raw log

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
func (it *ArgentModuleRecoveryExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArgentModuleRecoveryExecuted)
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
		it.Event = new(ArgentModuleRecoveryExecuted)
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
func (it *ArgentModuleRecoveryExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArgentModuleRecoveryExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArgentModuleRecoveryExecuted represents a RecoveryExecuted event raised by the ArgentModule contract.
type ArgentModuleRecoveryExecuted struct {
	Wallet       common.Address
	Recovery     common.Address
	ExecuteAfter uint64
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRecoveryExecuted is a free log retrieval operation binding the contract event 0x5f59bfd9baba55ae30bb440923cbbe30987d50e12a4e9134ffac3fd9afc3526d.
//
// Solidity: event RecoveryExecuted(address indexed wallet, address indexed _recovery, uint64 executeAfter)
func (_ArgentModule *ArgentModuleFilterer) FilterRecoveryExecuted(opts *bind.FilterOpts, wallet []common.Address, _recovery []common.Address) (*ArgentModuleRecoveryExecutedIterator, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}
	var _recoveryRule []interface{}
	for _, _recoveryItem := range _recovery {
		_recoveryRule = append(_recoveryRule, _recoveryItem)
	}

	logs, sub, err := _ArgentModule.contract.FilterLogs(opts, "RecoveryExecuted", walletRule, _recoveryRule)
	if err != nil {
		return nil, err
	}
	return &ArgentModuleRecoveryExecutedIterator{contract: _ArgentModule.contract, event: "RecoveryExecuted", logs: logs, sub: sub}, nil
}

// WatchRecoveryExecuted is a free log subscription operation binding the contract event 0x5f59bfd9baba55ae30bb440923cbbe30987d50e12a4e9134ffac3fd9afc3526d.
//
// Solidity: event RecoveryExecuted(address indexed wallet, address indexed _recovery, uint64 executeAfter)
func (_ArgentModule *ArgentModuleFilterer) WatchRecoveryExecuted(opts *bind.WatchOpts, sink chan<- *ArgentModuleRecoveryExecuted, wallet []common.Address, _recovery []common.Address) (event.Subscription, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}
	var _recoveryRule []interface{}
	for _, _recoveryItem := range _recovery {
		_recoveryRule = append(_recoveryRule, _recoveryItem)
	}

	logs, sub, err := _ArgentModule.contract.WatchLogs(opts, "RecoveryExecuted", walletRule, _recoveryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArgentModuleRecoveryExecuted)
				if err := _ArgentModule.contract.UnpackLog(event, "RecoveryExecuted", log); err != nil {
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

// ParseRecoveryExecuted is a log parse operation binding the contract event 0x5f59bfd9baba55ae30bb440923cbbe30987d50e12a4e9134ffac3fd9afc3526d.
//
// Solidity: event RecoveryExecuted(address indexed wallet, address indexed _recovery, uint64 executeAfter)
func (_ArgentModule *ArgentModuleFilterer) ParseRecoveryExecuted(log types.Log) (*ArgentModuleRecoveryExecuted, error) {
	event := new(ArgentModuleRecoveryExecuted)
	if err := _ArgentModule.contract.UnpackLog(event, "RecoveryExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArgentModuleRecoveryFinalizedIterator is returned from FilterRecoveryFinalized and is used to iterate over the raw logs and unpacked data for RecoveryFinalized events raised by the ArgentModule contract.
type ArgentModuleRecoveryFinalizedIterator struct {
	Event *ArgentModuleRecoveryFinalized // Event containing the contract specifics and raw log

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
func (it *ArgentModuleRecoveryFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArgentModuleRecoveryFinalized)
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
		it.Event = new(ArgentModuleRecoveryFinalized)
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
func (it *ArgentModuleRecoveryFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArgentModuleRecoveryFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArgentModuleRecoveryFinalized represents a RecoveryFinalized event raised by the ArgentModule contract.
type ArgentModuleRecoveryFinalized struct {
	Wallet   common.Address
	Recovery common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRecoveryFinalized is a free log retrieval operation binding the contract event 0xd8667de85dae2d56d76e700d16de53d21ac2ce4d5549cb0bf51c55fdc37f0bc1.
//
// Solidity: event RecoveryFinalized(address indexed wallet, address indexed _recovery)
func (_ArgentModule *ArgentModuleFilterer) FilterRecoveryFinalized(opts *bind.FilterOpts, wallet []common.Address, _recovery []common.Address) (*ArgentModuleRecoveryFinalizedIterator, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}
	var _recoveryRule []interface{}
	for _, _recoveryItem := range _recovery {
		_recoveryRule = append(_recoveryRule, _recoveryItem)
	}

	logs, sub, err := _ArgentModule.contract.FilterLogs(opts, "RecoveryFinalized", walletRule, _recoveryRule)
	if err != nil {
		return nil, err
	}
	return &ArgentModuleRecoveryFinalizedIterator{contract: _ArgentModule.contract, event: "RecoveryFinalized", logs: logs, sub: sub}, nil
}

// WatchRecoveryFinalized is a free log subscription operation binding the contract event 0xd8667de85dae2d56d76e700d16de53d21ac2ce4d5549cb0bf51c55fdc37f0bc1.
//
// Solidity: event RecoveryFinalized(address indexed wallet, address indexed _recovery)
func (_ArgentModule *ArgentModuleFilterer) WatchRecoveryFinalized(opts *bind.WatchOpts, sink chan<- *ArgentModuleRecoveryFinalized, wallet []common.Address, _recovery []common.Address) (event.Subscription, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}
	var _recoveryRule []interface{}
	for _, _recoveryItem := range _recovery {
		_recoveryRule = append(_recoveryRule, _recoveryItem)
	}

	logs, sub, err := _ArgentModule.contract.WatchLogs(opts, "RecoveryFinalized", walletRule, _recoveryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArgentModuleRecoveryFinalized)
				if err := _ArgentModule.contract.UnpackLog(event, "RecoveryFinalized", log); err != nil {
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

// ParseRecoveryFinalized is a log parse operation binding the contract event 0xd8667de85dae2d56d76e700d16de53d21ac2ce4d5549cb0bf51c55fdc37f0bc1.
//
// Solidity: event RecoveryFinalized(address indexed wallet, address indexed _recovery)
func (_ArgentModule *ArgentModuleFilterer) ParseRecoveryFinalized(log types.Log) (*ArgentModuleRecoveryFinalized, error) {
	event := new(ArgentModuleRecoveryFinalized)
	if err := _ArgentModule.contract.UnpackLog(event, "RecoveryFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArgentModuleRefundIterator is returned from FilterRefund and is used to iterate over the raw logs and unpacked data for Refund events raised by the ArgentModule contract.
type ArgentModuleRefundIterator struct {
	Event *ArgentModuleRefund // Event containing the contract specifics and raw log

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
func (it *ArgentModuleRefundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArgentModuleRefund)
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
		it.Event = new(ArgentModuleRefund)
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
func (it *ArgentModuleRefundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArgentModuleRefundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArgentModuleRefund represents a Refund event raised by the ArgentModule contract.
type ArgentModuleRefund struct {
	Wallet        common.Address
	RefundAddress common.Address
	RefundToken   common.Address
	RefundAmount  *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRefund is a free log retrieval operation binding the contract event 0x22edd2bbb0b0afbdcf90d91da8a5e2100f8d8f67cdc766dee1742e9a36d6add3.
//
// Solidity: event Refund(address indexed wallet, address indexed refundAddress, address refundToken, uint256 refundAmount)
func (_ArgentModule *ArgentModuleFilterer) FilterRefund(opts *bind.FilterOpts, wallet []common.Address, refundAddress []common.Address) (*ArgentModuleRefundIterator, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}
	var refundAddressRule []interface{}
	for _, refundAddressItem := range refundAddress {
		refundAddressRule = append(refundAddressRule, refundAddressItem)
	}

	logs, sub, err := _ArgentModule.contract.FilterLogs(opts, "Refund", walletRule, refundAddressRule)
	if err != nil {
		return nil, err
	}
	return &ArgentModuleRefundIterator{contract: _ArgentModule.contract, event: "Refund", logs: logs, sub: sub}, nil
}

// WatchRefund is a free log subscription operation binding the contract event 0x22edd2bbb0b0afbdcf90d91da8a5e2100f8d8f67cdc766dee1742e9a36d6add3.
//
// Solidity: event Refund(address indexed wallet, address indexed refundAddress, address refundToken, uint256 refundAmount)
func (_ArgentModule *ArgentModuleFilterer) WatchRefund(opts *bind.WatchOpts, sink chan<- *ArgentModuleRefund, wallet []common.Address, refundAddress []common.Address) (event.Subscription, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}
	var refundAddressRule []interface{}
	for _, refundAddressItem := range refundAddress {
		refundAddressRule = append(refundAddressRule, refundAddressItem)
	}

	logs, sub, err := _ArgentModule.contract.WatchLogs(opts, "Refund", walletRule, refundAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArgentModuleRefund)
				if err := _ArgentModule.contract.UnpackLog(event, "Refund", log); err != nil {
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

// ParseRefund is a log parse operation binding the contract event 0x22edd2bbb0b0afbdcf90d91da8a5e2100f8d8f67cdc766dee1742e9a36d6add3.
//
// Solidity: event Refund(address indexed wallet, address indexed refundAddress, address refundToken, uint256 refundAmount)
func (_ArgentModule *ArgentModuleFilterer) ParseRefund(log types.Log) (*ArgentModuleRefund, error) {
	event := new(ArgentModuleRefund)
	if err := _ArgentModule.contract.UnpackLog(event, "Refund", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArgentModuleRemovedFromWhitelistIterator is returned from FilterRemovedFromWhitelist and is used to iterate over the raw logs and unpacked data for RemovedFromWhitelist events raised by the ArgentModule contract.
type ArgentModuleRemovedFromWhitelistIterator struct {
	Event *ArgentModuleRemovedFromWhitelist // Event containing the contract specifics and raw log

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
func (it *ArgentModuleRemovedFromWhitelistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArgentModuleRemovedFromWhitelist)
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
		it.Event = new(ArgentModuleRemovedFromWhitelist)
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
func (it *ArgentModuleRemovedFromWhitelistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArgentModuleRemovedFromWhitelistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArgentModuleRemovedFromWhitelist represents a RemovedFromWhitelist event raised by the ArgentModule contract.
type ArgentModuleRemovedFromWhitelist struct {
	Wallet common.Address
	Target common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRemovedFromWhitelist is a free log retrieval operation binding the contract event 0xd288ab5da2e1f37cf384a1565a3f905ad289b092fbdd31950dbbfef148c04f88.
//
// Solidity: event RemovedFromWhitelist(address indexed wallet, address indexed target)
func (_ArgentModule *ArgentModuleFilterer) FilterRemovedFromWhitelist(opts *bind.FilterOpts, wallet []common.Address, target []common.Address) (*ArgentModuleRemovedFromWhitelistIterator, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _ArgentModule.contract.FilterLogs(opts, "RemovedFromWhitelist", walletRule, targetRule)
	if err != nil {
		return nil, err
	}
	return &ArgentModuleRemovedFromWhitelistIterator{contract: _ArgentModule.contract, event: "RemovedFromWhitelist", logs: logs, sub: sub}, nil
}

// WatchRemovedFromWhitelist is a free log subscription operation binding the contract event 0xd288ab5da2e1f37cf384a1565a3f905ad289b092fbdd31950dbbfef148c04f88.
//
// Solidity: event RemovedFromWhitelist(address indexed wallet, address indexed target)
func (_ArgentModule *ArgentModuleFilterer) WatchRemovedFromWhitelist(opts *bind.WatchOpts, sink chan<- *ArgentModuleRemovedFromWhitelist, wallet []common.Address, target []common.Address) (event.Subscription, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _ArgentModule.contract.WatchLogs(opts, "RemovedFromWhitelist", walletRule, targetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArgentModuleRemovedFromWhitelist)
				if err := _ArgentModule.contract.UnpackLog(event, "RemovedFromWhitelist", log); err != nil {
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

// ParseRemovedFromWhitelist is a log parse operation binding the contract event 0xd288ab5da2e1f37cf384a1565a3f905ad289b092fbdd31950dbbfef148c04f88.
//
// Solidity: event RemovedFromWhitelist(address indexed wallet, address indexed target)
func (_ArgentModule *ArgentModuleFilterer) ParseRemovedFromWhitelist(log types.Log) (*ArgentModuleRemovedFromWhitelist, error) {
	event := new(ArgentModuleRemovedFromWhitelist)
	if err := _ArgentModule.contract.UnpackLog(event, "RemovedFromWhitelist", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArgentModuleSessionClearedIterator is returned from FilterSessionCleared and is used to iterate over the raw logs and unpacked data for SessionCleared events raised by the ArgentModule contract.
type ArgentModuleSessionClearedIterator struct {
	Event *ArgentModuleSessionCleared // Event containing the contract specifics and raw log

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
func (it *ArgentModuleSessionClearedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArgentModuleSessionCleared)
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
		it.Event = new(ArgentModuleSessionCleared)
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
func (it *ArgentModuleSessionClearedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArgentModuleSessionClearedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArgentModuleSessionCleared represents a SessionCleared event raised by the ArgentModule contract.
type ArgentModuleSessionCleared struct {
	Wallet     common.Address
	SessionKey common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSessionCleared is a free log retrieval operation binding the contract event 0xeb290a597820eccc6b8b31f942bd97c633d5138f4d849751f770f3cb3900e57a.
//
// Solidity: event SessionCleared(address indexed wallet, address sessionKey)
func (_ArgentModule *ArgentModuleFilterer) FilterSessionCleared(opts *bind.FilterOpts, wallet []common.Address) (*ArgentModuleSessionClearedIterator, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _ArgentModule.contract.FilterLogs(opts, "SessionCleared", walletRule)
	if err != nil {
		return nil, err
	}
	return &ArgentModuleSessionClearedIterator{contract: _ArgentModule.contract, event: "SessionCleared", logs: logs, sub: sub}, nil
}

// WatchSessionCleared is a free log subscription operation binding the contract event 0xeb290a597820eccc6b8b31f942bd97c633d5138f4d849751f770f3cb3900e57a.
//
// Solidity: event SessionCleared(address indexed wallet, address sessionKey)
func (_ArgentModule *ArgentModuleFilterer) WatchSessionCleared(opts *bind.WatchOpts, sink chan<- *ArgentModuleSessionCleared, wallet []common.Address) (event.Subscription, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _ArgentModule.contract.WatchLogs(opts, "SessionCleared", walletRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArgentModuleSessionCleared)
				if err := _ArgentModule.contract.UnpackLog(event, "SessionCleared", log); err != nil {
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

// ParseSessionCleared is a log parse operation binding the contract event 0xeb290a597820eccc6b8b31f942bd97c633d5138f4d849751f770f3cb3900e57a.
//
// Solidity: event SessionCleared(address indexed wallet, address sessionKey)
func (_ArgentModule *ArgentModuleFilterer) ParseSessionCleared(log types.Log) (*ArgentModuleSessionCleared, error) {
	event := new(ArgentModuleSessionCleared)
	if err := _ArgentModule.contract.UnpackLog(event, "SessionCleared", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArgentModuleSessionCreatedIterator is returned from FilterSessionCreated and is used to iterate over the raw logs and unpacked data for SessionCreated events raised by the ArgentModule contract.
type ArgentModuleSessionCreatedIterator struct {
	Event *ArgentModuleSessionCreated // Event containing the contract specifics and raw log

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
func (it *ArgentModuleSessionCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArgentModuleSessionCreated)
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
		it.Event = new(ArgentModuleSessionCreated)
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
func (it *ArgentModuleSessionCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArgentModuleSessionCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArgentModuleSessionCreated represents a SessionCreated event raised by the ArgentModule contract.
type ArgentModuleSessionCreated struct {
	Wallet     common.Address
	SessionKey common.Address
	Expires    uint64
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSessionCreated is a free log retrieval operation binding the contract event 0x2ecea11087d1dc1431b517cbb5a559a9e33e58a1afeaac288f782c1c8bed8b8a.
//
// Solidity: event SessionCreated(address indexed wallet, address sessionKey, uint64 expires)
func (_ArgentModule *ArgentModuleFilterer) FilterSessionCreated(opts *bind.FilterOpts, wallet []common.Address) (*ArgentModuleSessionCreatedIterator, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _ArgentModule.contract.FilterLogs(opts, "SessionCreated", walletRule)
	if err != nil {
		return nil, err
	}
	return &ArgentModuleSessionCreatedIterator{contract: _ArgentModule.contract, event: "SessionCreated", logs: logs, sub: sub}, nil
}

// WatchSessionCreated is a free log subscription operation binding the contract event 0x2ecea11087d1dc1431b517cbb5a559a9e33e58a1afeaac288f782c1c8bed8b8a.
//
// Solidity: event SessionCreated(address indexed wallet, address sessionKey, uint64 expires)
func (_ArgentModule *ArgentModuleFilterer) WatchSessionCreated(opts *bind.WatchOpts, sink chan<- *ArgentModuleSessionCreated, wallet []common.Address) (event.Subscription, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _ArgentModule.contract.WatchLogs(opts, "SessionCreated", walletRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArgentModuleSessionCreated)
				if err := _ArgentModule.contract.UnpackLog(event, "SessionCreated", log); err != nil {
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

// ParseSessionCreated is a log parse operation binding the contract event 0x2ecea11087d1dc1431b517cbb5a559a9e33e58a1afeaac288f782c1c8bed8b8a.
//
// Solidity: event SessionCreated(address indexed wallet, address sessionKey, uint64 expires)
func (_ArgentModule *ArgentModuleFilterer) ParseSessionCreated(log types.Log) (*ArgentModuleSessionCreated, error) {
	event := new(ArgentModuleSessionCreated)
	if err := _ArgentModule.contract.UnpackLog(event, "SessionCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArgentModuleTransactionExecutedIterator is returned from FilterTransactionExecuted and is used to iterate over the raw logs and unpacked data for TransactionExecuted events raised by the ArgentModule contract.
type ArgentModuleTransactionExecutedIterator struct {
	Event *ArgentModuleTransactionExecuted // Event containing the contract specifics and raw log

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
func (it *ArgentModuleTransactionExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArgentModuleTransactionExecuted)
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
		it.Event = new(ArgentModuleTransactionExecuted)
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
func (it *ArgentModuleTransactionExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArgentModuleTransactionExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArgentModuleTransactionExecuted represents a TransactionExecuted event raised by the ArgentModule contract.
type ArgentModuleTransactionExecuted struct {
	Wallet     common.Address
	Success    bool
	ReturnData []byte
	SignedHash [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTransactionExecuted is a free log retrieval operation binding the contract event 0x7da4525a280527268ba2e963ee6c1b18f43c9507bcb1d2560f652ab17c76e90a.
//
// Solidity: event TransactionExecuted(address indexed wallet, bool indexed success, bytes returnData, bytes32 signedHash)
func (_ArgentModule *ArgentModuleFilterer) FilterTransactionExecuted(opts *bind.FilterOpts, wallet []common.Address, success []bool) (*ArgentModuleTransactionExecutedIterator, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}
	var successRule []interface{}
	for _, successItem := range success {
		successRule = append(successRule, successItem)
	}

	logs, sub, err := _ArgentModule.contract.FilterLogs(opts, "TransactionExecuted", walletRule, successRule)
	if err != nil {
		return nil, err
	}
	return &ArgentModuleTransactionExecutedIterator{contract: _ArgentModule.contract, event: "TransactionExecuted", logs: logs, sub: sub}, nil
}

// WatchTransactionExecuted is a free log subscription operation binding the contract event 0x7da4525a280527268ba2e963ee6c1b18f43c9507bcb1d2560f652ab17c76e90a.
//
// Solidity: event TransactionExecuted(address indexed wallet, bool indexed success, bytes returnData, bytes32 signedHash)
func (_ArgentModule *ArgentModuleFilterer) WatchTransactionExecuted(opts *bind.WatchOpts, sink chan<- *ArgentModuleTransactionExecuted, wallet []common.Address, success []bool) (event.Subscription, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}
	var successRule []interface{}
	for _, successItem := range success {
		successRule = append(successRule, successItem)
	}

	logs, sub, err := _ArgentModule.contract.WatchLogs(opts, "TransactionExecuted", walletRule, successRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArgentModuleTransactionExecuted)
				if err := _ArgentModule.contract.UnpackLog(event, "TransactionExecuted", log); err != nil {
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

// ParseTransactionExecuted is a log parse operation binding the contract event 0x7da4525a280527268ba2e963ee6c1b18f43c9507bcb1d2560f652ab17c76e90a.
//
// Solidity: event TransactionExecuted(address indexed wallet, bool indexed success, bytes returnData, bytes32 signedHash)
func (_ArgentModule *ArgentModuleFilterer) ParseTransactionExecuted(log types.Log) (*ArgentModuleTransactionExecuted, error) {
	event := new(ArgentModuleTransactionExecuted)
	if err := _ArgentModule.contract.UnpackLog(event, "TransactionExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArgentModuleUnlockedIterator is returned from FilterUnlocked and is used to iterate over the raw logs and unpacked data for Unlocked events raised by the ArgentModule contract.
type ArgentModuleUnlockedIterator struct {
	Event *ArgentModuleUnlocked // Event containing the contract specifics and raw log

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
func (it *ArgentModuleUnlockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArgentModuleUnlocked)
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
		it.Event = new(ArgentModuleUnlocked)
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
func (it *ArgentModuleUnlockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArgentModuleUnlockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArgentModuleUnlocked represents a Unlocked event raised by the ArgentModule contract.
type ArgentModuleUnlocked struct {
	Wallet common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUnlocked is a free log retrieval operation binding the contract event 0x7e6adfec7e3f286831a0200a754127c171a2da564078722cb97704741bbdb0ea.
//
// Solidity: event Unlocked(address indexed wallet)
func (_ArgentModule *ArgentModuleFilterer) FilterUnlocked(opts *bind.FilterOpts, wallet []common.Address) (*ArgentModuleUnlockedIterator, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _ArgentModule.contract.FilterLogs(opts, "Unlocked", walletRule)
	if err != nil {
		return nil, err
	}
	return &ArgentModuleUnlockedIterator{contract: _ArgentModule.contract, event: "Unlocked", logs: logs, sub: sub}, nil
}

// WatchUnlocked is a free log subscription operation binding the contract event 0x7e6adfec7e3f286831a0200a754127c171a2da564078722cb97704741bbdb0ea.
//
// Solidity: event Unlocked(address indexed wallet)
func (_ArgentModule *ArgentModuleFilterer) WatchUnlocked(opts *bind.WatchOpts, sink chan<- *ArgentModuleUnlocked, wallet []common.Address) (event.Subscription, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _ArgentModule.contract.WatchLogs(opts, "Unlocked", walletRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArgentModuleUnlocked)
				if err := _ArgentModule.contract.UnpackLog(event, "Unlocked", log); err != nil {
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

// ParseUnlocked is a log parse operation binding the contract event 0x7e6adfec7e3f286831a0200a754127c171a2da564078722cb97704741bbdb0ea.
//
// Solidity: event Unlocked(address indexed wallet)
func (_ArgentModule *ArgentModuleFilterer) ParseUnlocked(log types.Log) (*ArgentModuleUnlocked, error) {
	event := new(ArgentModuleUnlocked)
	if err := _ArgentModule.contract.UnpackLog(event, "Unlocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
