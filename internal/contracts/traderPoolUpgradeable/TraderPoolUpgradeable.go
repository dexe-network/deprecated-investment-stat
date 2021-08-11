// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package traderPoolUpgradeable

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
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// TokenABI is the input ABI used to generate the binding from.
const TokenABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"who\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountBT\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int128\",\"name\":\"price\",\"type\":\"int128\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"fromAsset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"toAsset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fromAmt\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"toAmt\",\"type\":\"uint256\"}],\"name\":\"Exchanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Loss\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Profit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"who\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountBT\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"commision\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TRADER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"assetTokenAddresses\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"availableCap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"basicToken\",\"outputs\":[{\"internalType\":\"contractIERC20Token\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"commissions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"depositETHTo\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"depositTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"deposits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"int128\",\"name\":\"price\",\"type\":\"int128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dexeCommissionAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dexeCommissionBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"insuranceContractAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"investorWhitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isActualOn\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isInvestorsWhitelistEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxDepositedTokenPrice\",\"outputs\":[{\"internalType\":\"int128\",\"name\":\"\",\"type\":\"int128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"plt\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"storageVersion\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"traderCommissionAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"traderCommissionBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"traderLiquidityBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"traderWhitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"withdrawETHTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"withdrawTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[9]\",\"name\":\"iaddr\",\"type\":\"address[9]\"},{\"internalType\":\"uint256\",\"name\":\"_commissions\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_actual\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"_investorRestricted\",\"type\":\"bool\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_traderAddress\",\"type\":\"address\"}],\"name\":\"addTraderAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_traderAddress\",\"type\":\"address\"}],\"name\":\"removeTraderAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_investors\",\"type\":\"address[]\"}],\"name\":\"addInvestorAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_investors\",\"type\":\"address[]\"}],\"name\":\"removeInvestorAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"fromAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"toAsset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fromAmt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_calldata\",\"type\":\"bytes\"}],\"name\":\"initiateExchangeOperation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"positionsLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_index\",\"type\":\"uint16\"}],\"name\":\"positionAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"positionFor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawTraderCommission\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawDexeCommission\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_traderCommissionAddress\",\"type\":\"address\"}],\"name\":\"setTraderCommissionAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_type\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"_nom\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"_denom\",\"type\":\"uint16\"}],\"name\":\"setCommission\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMaxPositionOpenAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"}],\"name\":\"getUserData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"int128\",\"name\":\"\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalValueLocked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_type\",\"type\":\"uint256\"}],\"name\":\"getCommission\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Token is an auto generated Go binding around an Ethereum contract.
type Token struct {
	TokenCaller     // Read-only binding to the contract
	TokenTransactor // Write-only binding to the contract
	TokenFilterer   // Log filterer for contract events
}

// TokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenSession struct {
	Contract     *Token            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenCallerSession struct {
	Contract *TokenCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenTransactorSession struct {
	Contract     *TokenTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenRaw struct {
	Contract *Token // Generic contract binding to access the raw methods on
}

// TokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenCallerRaw struct {
	Contract *TokenCaller // Generic read-only contract binding to access the raw methods on
}

// TokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenTransactorRaw struct {
	Contract *TokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewToken creates a new instance of Token, bound to a specific deployed contract.
func NewToken(address common.Address, backend bind.ContractBackend) (*Token, error) {
	contract, err := bindToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Token{TokenCaller: TokenCaller{contract: contract}, TokenTransactor: TokenTransactor{contract: contract}, TokenFilterer: TokenFilterer{contract: contract}}, nil
}

// NewTokenCaller creates a new read-only instance of Token, bound to a specific deployed contract.
func NewTokenCaller(address common.Address, caller bind.ContractCaller) (*TokenCaller, error) {
	contract, err := bindToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenCaller{contract: contract}, nil
}

// NewTokenTransactor creates a new write-only instance of Token, bound to a specific deployed contract.
func NewTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenTransactor, error) {
	contract, err := bindToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenTransactor{contract: contract}, nil
}

// NewTokenFilterer creates a new log filterer instance of Token, bound to a specific deployed contract.
func NewTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenFilterer, error) {
	contract, err := bindToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenFilterer{contract: contract}, nil
}

// bindToken binds a generic wrapper to an already deployed contract.
func bindToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Token *TokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Token.Contract.TokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Token *TokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.Contract.TokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Token *TokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Token.Contract.TokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Token *TokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Token.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Token *TokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Token *TokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Token.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Token *TokenCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Token *TokenSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Token.Contract.DEFAULTADMINROLE(&_Token.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Token *TokenCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Token.Contract.DEFAULTADMINROLE(&_Token.CallOpts)
}

// TRADERROLE is a free data retrieval call binding the contract method 0xf0a56fc8.
//
// Solidity: function TRADER_ROLE() view returns(bytes32)
func (_Token *TokenCaller) TRADERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "TRADER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TRADERROLE is a free data retrieval call binding the contract method 0xf0a56fc8.
//
// Solidity: function TRADER_ROLE() view returns(bytes32)
func (_Token *TokenSession) TRADERROLE() ([32]byte, error) {
	return _Token.Contract.TRADERROLE(&_Token.CallOpts)
}

// TRADERROLE is a free data retrieval call binding the contract method 0xf0a56fc8.
//
// Solidity: function TRADER_ROLE() view returns(bytes32)
func (_Token *TokenCallerSession) TRADERROLE() ([32]byte, error) {
	return _Token.Contract.TRADERROLE(&_Token.CallOpts)
}

// AssetTokenAddresses is a free data retrieval call binding the contract method 0xe92ca572.
//
// Solidity: function assetTokenAddresses(uint256 ) view returns(address)
func (_Token *TokenCaller) AssetTokenAddresses(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "assetTokenAddresses", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AssetTokenAddresses is a free data retrieval call binding the contract method 0xe92ca572.
//
// Solidity: function assetTokenAddresses(uint256 ) view returns(address)
func (_Token *TokenSession) AssetTokenAddresses(arg0 *big.Int) (common.Address, error) {
	return _Token.Contract.AssetTokenAddresses(&_Token.CallOpts, arg0)
}

// AssetTokenAddresses is a free data retrieval call binding the contract method 0xe92ca572.
//
// Solidity: function assetTokenAddresses(uint256 ) view returns(address)
func (_Token *TokenCallerSession) AssetTokenAddresses(arg0 *big.Int) (common.Address, error) {
	return _Token.Contract.AssetTokenAddresses(&_Token.CallOpts, arg0)
}

// AvailableCap is a free data retrieval call binding the contract method 0x17630ded.
//
// Solidity: function availableCap() view returns(uint256)
func (_Token *TokenCaller) AvailableCap(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "availableCap")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AvailableCap is a free data retrieval call binding the contract method 0x17630ded.
//
// Solidity: function availableCap() view returns(uint256)
func (_Token *TokenSession) AvailableCap() (*big.Int, error) {
	return _Token.Contract.AvailableCap(&_Token.CallOpts)
}

// AvailableCap is a free data retrieval call binding the contract method 0x17630ded.
//
// Solidity: function availableCap() view returns(uint256)
func (_Token *TokenCallerSession) AvailableCap() (*big.Int, error) {
	return _Token.Contract.AvailableCap(&_Token.CallOpts)
}

// BasicToken is a free data retrieval call binding the contract method 0x190b8dc5.
//
// Solidity: function basicToken() view returns(address)
func (_Token *TokenCaller) BasicToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "basicToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BasicToken is a free data retrieval call binding the contract method 0x190b8dc5.
//
// Solidity: function basicToken() view returns(address)
func (_Token *TokenSession) BasicToken() (common.Address, error) {
	return _Token.Contract.BasicToken(&_Token.CallOpts)
}

// BasicToken is a free data retrieval call binding the contract method 0x190b8dc5.
//
// Solidity: function basicToken() view returns(address)
func (_Token *TokenCallerSession) BasicToken() (common.Address, error) {
	return _Token.Contract.BasicToken(&_Token.CallOpts)
}

// Commissions is a free data retrieval call binding the contract method 0x790aed17.
//
// Solidity: function commissions() view returns(uint256)
func (_Token *TokenCaller) Commissions(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "commissions")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Commissions is a free data retrieval call binding the contract method 0x790aed17.
//
// Solidity: function commissions() view returns(uint256)
func (_Token *TokenSession) Commissions() (*big.Int, error) {
	return _Token.Contract.Commissions(&_Token.CallOpts)
}

// Commissions is a free data retrieval call binding the contract method 0x790aed17.
//
// Solidity: function commissions() view returns(uint256)
func (_Token *TokenCallerSession) Commissions() (*big.Int, error) {
	return _Token.Contract.Commissions(&_Token.CallOpts)
}

// Deposits is a free data retrieval call binding the contract method 0xfc7e286d.
//
// Solidity: function deposits(address ) view returns(uint256 amount, int128 price)
func (_Token *TokenCaller) Deposits(opts *bind.CallOpts, arg0 common.Address) (struct {
	Amount *big.Int
	Price  *big.Int
}, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "deposits", arg0)

	outstruct := new(struct {
		Amount *big.Int
		Price  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Price = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Deposits is a free data retrieval call binding the contract method 0xfc7e286d.
//
// Solidity: function deposits(address ) view returns(uint256 amount, int128 price)
func (_Token *TokenSession) Deposits(arg0 common.Address) (struct {
	Amount *big.Int
	Price  *big.Int
}, error) {
	return _Token.Contract.Deposits(&_Token.CallOpts, arg0)
}

// Deposits is a free data retrieval call binding the contract method 0xfc7e286d.
//
// Solidity: function deposits(address ) view returns(uint256 amount, int128 price)
func (_Token *TokenCallerSession) Deposits(arg0 common.Address) (struct {
	Amount *big.Int
	Price  *big.Int
}, error) {
	return _Token.Contract.Deposits(&_Token.CallOpts, arg0)
}

// DexeCommissionAddress is a free data retrieval call binding the contract method 0x324e449d.
//
// Solidity: function dexeCommissionAddress() view returns(address)
func (_Token *TokenCaller) DexeCommissionAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "dexeCommissionAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DexeCommissionAddress is a free data retrieval call binding the contract method 0x324e449d.
//
// Solidity: function dexeCommissionAddress() view returns(address)
func (_Token *TokenSession) DexeCommissionAddress() (common.Address, error) {
	return _Token.Contract.DexeCommissionAddress(&_Token.CallOpts)
}

// DexeCommissionAddress is a free data retrieval call binding the contract method 0x324e449d.
//
// Solidity: function dexeCommissionAddress() view returns(address)
func (_Token *TokenCallerSession) DexeCommissionAddress() (common.Address, error) {
	return _Token.Contract.DexeCommissionAddress(&_Token.CallOpts)
}

// DexeCommissionBalance is a free data retrieval call binding the contract method 0x7108921f.
//
// Solidity: function dexeCommissionBalance() view returns(uint256)
func (_Token *TokenCaller) DexeCommissionBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "dexeCommissionBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DexeCommissionBalance is a free data retrieval call binding the contract method 0x7108921f.
//
// Solidity: function dexeCommissionBalance() view returns(uint256)
func (_Token *TokenSession) DexeCommissionBalance() (*big.Int, error) {
	return _Token.Contract.DexeCommissionBalance(&_Token.CallOpts)
}

// DexeCommissionBalance is a free data retrieval call binding the contract method 0x7108921f.
//
// Solidity: function dexeCommissionBalance() view returns(uint256)
func (_Token *TokenCallerSession) DexeCommissionBalance() (*big.Int, error) {
	return _Token.Contract.DexeCommissionBalance(&_Token.CallOpts)
}

// GetCommission is a free data retrieval call binding the contract method 0x3c8bccd9.
//
// Solidity: function getCommission(uint256 _type) view returns(uint16, uint16)
func (_Token *TokenCaller) GetCommission(opts *bind.CallOpts, _type *big.Int) (uint16, uint16, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "getCommission", _type)

	if err != nil {
		return *new(uint16), *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)
	out1 := *abi.ConvertType(out[1], new(uint16)).(*uint16)

	return out0, out1, err

}

// GetCommission is a free data retrieval call binding the contract method 0x3c8bccd9.
//
// Solidity: function getCommission(uint256 _type) view returns(uint16, uint16)
func (_Token *TokenSession) GetCommission(_type *big.Int) (uint16, uint16, error) {
	return _Token.Contract.GetCommission(&_Token.CallOpts, _type)
}

// GetCommission is a free data retrieval call binding the contract method 0x3c8bccd9.
//
// Solidity: function getCommission(uint256 _type) view returns(uint16, uint16)
func (_Token *TokenCallerSession) GetCommission(_type *big.Int) (uint16, uint16, error) {
	return _Token.Contract.GetCommission(&_Token.CallOpts, _type)
}

// GetMaxPositionOpenAmount is a free data retrieval call binding the contract method 0x0f9d04a1.
//
// Solidity: function getMaxPositionOpenAmount() view returns(uint256)
func (_Token *TokenCaller) GetMaxPositionOpenAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "getMaxPositionOpenAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaxPositionOpenAmount is a free data retrieval call binding the contract method 0x0f9d04a1.
//
// Solidity: function getMaxPositionOpenAmount() view returns(uint256)
func (_Token *TokenSession) GetMaxPositionOpenAmount() (*big.Int, error) {
	return _Token.Contract.GetMaxPositionOpenAmount(&_Token.CallOpts)
}

// GetMaxPositionOpenAmount is a free data retrieval call binding the contract method 0x0f9d04a1.
//
// Solidity: function getMaxPositionOpenAmount() view returns(uint256)
func (_Token *TokenCallerSession) GetMaxPositionOpenAmount() (*big.Int, error) {
	return _Token.Contract.GetMaxPositionOpenAmount(&_Token.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Token *TokenCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Token *TokenSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Token.Contract.GetRoleAdmin(&_Token.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Token *TokenCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Token.Contract.GetRoleAdmin(&_Token.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Token *TokenCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Token *TokenSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Token.Contract.GetRoleMember(&_Token.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Token *TokenCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Token.Contract.GetRoleMember(&_Token.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Token *TokenCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Token *TokenSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Token.Contract.GetRoleMemberCount(&_Token.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Token *TokenCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Token.Contract.GetRoleMemberCount(&_Token.CallOpts, role)
}

// GetTotalValueLocked is a free data retrieval call binding the contract method 0xb26025aa.
//
// Solidity: function getTotalValueLocked() view returns(uint256, uint256)
func (_Token *TokenCaller) GetTotalValueLocked(opts *bind.CallOpts) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "getTotalValueLocked")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetTotalValueLocked is a free data retrieval call binding the contract method 0xb26025aa.
//
// Solidity: function getTotalValueLocked() view returns(uint256, uint256)
func (_Token *TokenSession) GetTotalValueLocked() (*big.Int, *big.Int, error) {
	return _Token.Contract.GetTotalValueLocked(&_Token.CallOpts)
}

// GetTotalValueLocked is a free data retrieval call binding the contract method 0xb26025aa.
//
// Solidity: function getTotalValueLocked() view returns(uint256, uint256)
func (_Token *TokenCallerSession) GetTotalValueLocked() (*big.Int, *big.Int, error) {
	return _Token.Contract.GetTotalValueLocked(&_Token.CallOpts)
}

// GetUserData is a free data retrieval call binding the contract method 0xffc9896b.
//
// Solidity: function getUserData(address holder) view returns(uint256, int128, uint256)
func (_Token *TokenCaller) GetUserData(opts *bind.CallOpts, holder common.Address) (*big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "getUserData", holder)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// GetUserData is a free data retrieval call binding the contract method 0xffc9896b.
//
// Solidity: function getUserData(address holder) view returns(uint256, int128, uint256)
func (_Token *TokenSession) GetUserData(holder common.Address) (*big.Int, *big.Int, *big.Int, error) {
	return _Token.Contract.GetUserData(&_Token.CallOpts, holder)
}

// GetUserData is a free data retrieval call binding the contract method 0xffc9896b.
//
// Solidity: function getUserData(address holder) view returns(uint256, int128, uint256)
func (_Token *TokenCallerSession) GetUserData(holder common.Address) (*big.Int, *big.Int, *big.Int, error) {
	return _Token.Contract.GetUserData(&_Token.CallOpts, holder)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Token *TokenCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Token *TokenSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Token.Contract.HasRole(&_Token.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Token *TokenCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Token.Contract.HasRole(&_Token.CallOpts, role, account)
}

// InsuranceContractAddress is a free data retrieval call binding the contract method 0x14c179c0.
//
// Solidity: function insuranceContractAddress() view returns(address)
func (_Token *TokenCaller) InsuranceContractAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "insuranceContractAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// InsuranceContractAddress is a free data retrieval call binding the contract method 0x14c179c0.
//
// Solidity: function insuranceContractAddress() view returns(address)
func (_Token *TokenSession) InsuranceContractAddress() (common.Address, error) {
	return _Token.Contract.InsuranceContractAddress(&_Token.CallOpts)
}

// InsuranceContractAddress is a free data retrieval call binding the contract method 0x14c179c0.
//
// Solidity: function insuranceContractAddress() view returns(address)
func (_Token *TokenCallerSession) InsuranceContractAddress() (common.Address, error) {
	return _Token.Contract.InsuranceContractAddress(&_Token.CallOpts)
}

// InvestorWhitelist is a free data retrieval call binding the contract method 0x1bf8060c.
//
// Solidity: function investorWhitelist(address ) view returns(bool)
func (_Token *TokenCaller) InvestorWhitelist(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "investorWhitelist", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// InvestorWhitelist is a free data retrieval call binding the contract method 0x1bf8060c.
//
// Solidity: function investorWhitelist(address ) view returns(bool)
func (_Token *TokenSession) InvestorWhitelist(arg0 common.Address) (bool, error) {
	return _Token.Contract.InvestorWhitelist(&_Token.CallOpts, arg0)
}

// InvestorWhitelist is a free data retrieval call binding the contract method 0x1bf8060c.
//
// Solidity: function investorWhitelist(address ) view returns(bool)
func (_Token *TokenCallerSession) InvestorWhitelist(arg0 common.Address) (bool, error) {
	return _Token.Contract.InvestorWhitelist(&_Token.CallOpts, arg0)
}

// IsActualOn is a free data retrieval call binding the contract method 0xa1f1084f.
//
// Solidity: function isActualOn() view returns(bool)
func (_Token *TokenCaller) IsActualOn(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "isActualOn")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsActualOn is a free data retrieval call binding the contract method 0xa1f1084f.
//
// Solidity: function isActualOn() view returns(bool)
func (_Token *TokenSession) IsActualOn() (bool, error) {
	return _Token.Contract.IsActualOn(&_Token.CallOpts)
}

// IsActualOn is a free data retrieval call binding the contract method 0xa1f1084f.
//
// Solidity: function isActualOn() view returns(bool)
func (_Token *TokenCallerSession) IsActualOn() (bool, error) {
	return _Token.Contract.IsActualOn(&_Token.CallOpts)
}

// IsInvestorsWhitelistEnabled is a free data retrieval call binding the contract method 0xe9773ab8.
//
// Solidity: function isInvestorsWhitelistEnabled() view returns(bool)
func (_Token *TokenCaller) IsInvestorsWhitelistEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "isInvestorsWhitelistEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsInvestorsWhitelistEnabled is a free data retrieval call binding the contract method 0xe9773ab8.
//
// Solidity: function isInvestorsWhitelistEnabled() view returns(bool)
func (_Token *TokenSession) IsInvestorsWhitelistEnabled() (bool, error) {
	return _Token.Contract.IsInvestorsWhitelistEnabled(&_Token.CallOpts)
}

// IsInvestorsWhitelistEnabled is a free data retrieval call binding the contract method 0xe9773ab8.
//
// Solidity: function isInvestorsWhitelistEnabled() view returns(bool)
func (_Token *TokenCallerSession) IsInvestorsWhitelistEnabled() (bool, error) {
	return _Token.Contract.IsInvestorsWhitelistEnabled(&_Token.CallOpts)
}

// MaxDepositedTokenPrice is a free data retrieval call binding the contract method 0xdbb5ecd3.
//
// Solidity: function maxDepositedTokenPrice() view returns(int128)
func (_Token *TokenCaller) MaxDepositedTokenPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "maxDepositedTokenPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxDepositedTokenPrice is a free data retrieval call binding the contract method 0xdbb5ecd3.
//
// Solidity: function maxDepositedTokenPrice() view returns(int128)
func (_Token *TokenSession) MaxDepositedTokenPrice() (*big.Int, error) {
	return _Token.Contract.MaxDepositedTokenPrice(&_Token.CallOpts)
}

// MaxDepositedTokenPrice is a free data retrieval call binding the contract method 0xdbb5ecd3.
//
// Solidity: function maxDepositedTokenPrice() view returns(int128)
func (_Token *TokenCallerSession) MaxDepositedTokenPrice() (*big.Int, error) {
	return _Token.Contract.MaxDepositedTokenPrice(&_Token.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Token *TokenCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Token *TokenSession) Paused() (bool, error) {
	return _Token.Contract.Paused(&_Token.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Token *TokenCallerSession) Paused() (bool, error) {
	return _Token.Contract.Paused(&_Token.CallOpts)
}

// Plt is a free data retrieval call binding the contract method 0xbebdef25.
//
// Solidity: function plt() view returns(address)
func (_Token *TokenCaller) Plt(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "plt")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Plt is a free data retrieval call binding the contract method 0xbebdef25.
//
// Solidity: function plt() view returns(address)
func (_Token *TokenSession) Plt() (common.Address, error) {
	return _Token.Contract.Plt(&_Token.CallOpts)
}

// Plt is a free data retrieval call binding the contract method 0xbebdef25.
//
// Solidity: function plt() view returns(address)
func (_Token *TokenCallerSession) Plt() (common.Address, error) {
	return _Token.Contract.Plt(&_Token.CallOpts)
}

// PositionAt is a free data retrieval call binding the contract method 0xfd93b171.
//
// Solidity: function positionAt(uint16 _index) view returns(uint256, uint256, address)
func (_Token *TokenCaller) PositionAt(opts *bind.CallOpts, _index uint16) (*big.Int, *big.Int, common.Address, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "positionAt", _index)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(common.Address)).(*common.Address)

	return out0, out1, out2, err

}

// PositionAt is a free data retrieval call binding the contract method 0xfd93b171.
//
// Solidity: function positionAt(uint16 _index) view returns(uint256, uint256, address)
func (_Token *TokenSession) PositionAt(_index uint16) (*big.Int, *big.Int, common.Address, error) {
	return _Token.Contract.PositionAt(&_Token.CallOpts, _index)
}

// PositionAt is a free data retrieval call binding the contract method 0xfd93b171.
//
// Solidity: function positionAt(uint16 _index) view returns(uint256, uint256, address)
func (_Token *TokenCallerSession) PositionAt(_index uint16) (*big.Int, *big.Int, common.Address, error) {
	return _Token.Contract.PositionAt(&_Token.CallOpts, _index)
}

// PositionFor is a free data retrieval call binding the contract method 0x9d6bc8fa.
//
// Solidity: function positionFor(address asset) view returns(uint256, uint256, address)
func (_Token *TokenCaller) PositionFor(opts *bind.CallOpts, asset common.Address) (*big.Int, *big.Int, common.Address, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "positionFor", asset)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(common.Address)).(*common.Address)

	return out0, out1, out2, err

}

// PositionFor is a free data retrieval call binding the contract method 0x9d6bc8fa.
//
// Solidity: function positionFor(address asset) view returns(uint256, uint256, address)
func (_Token *TokenSession) PositionFor(asset common.Address) (*big.Int, *big.Int, common.Address, error) {
	return _Token.Contract.PositionFor(&_Token.CallOpts, asset)
}

// PositionFor is a free data retrieval call binding the contract method 0x9d6bc8fa.
//
// Solidity: function positionFor(address asset) view returns(uint256, uint256, address)
func (_Token *TokenCallerSession) PositionFor(asset common.Address) (*big.Int, *big.Int, common.Address, error) {
	return _Token.Contract.PositionFor(&_Token.CallOpts, asset)
}

// PositionsLength is a free data retrieval call binding the contract method 0xd6887bfa.
//
// Solidity: function positionsLength() view returns(uint256)
func (_Token *TokenCaller) PositionsLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "positionsLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PositionsLength is a free data retrieval call binding the contract method 0xd6887bfa.
//
// Solidity: function positionsLength() view returns(uint256)
func (_Token *TokenSession) PositionsLength() (*big.Int, error) {
	return _Token.Contract.PositionsLength(&_Token.CallOpts)
}

// PositionsLength is a free data retrieval call binding the contract method 0xd6887bfa.
//
// Solidity: function positionsLength() view returns(uint256)
func (_Token *TokenCallerSession) PositionsLength() (*big.Int, error) {
	return _Token.Contract.PositionsLength(&_Token.CallOpts)
}

// StorageVersion is a free data retrieval call binding the contract method 0x403ebd03.
//
// Solidity: function storageVersion() view returns(uint128)
func (_Token *TokenCaller) StorageVersion(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "storageVersion")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StorageVersion is a free data retrieval call binding the contract method 0x403ebd03.
//
// Solidity: function storageVersion() view returns(uint128)
func (_Token *TokenSession) StorageVersion() (*big.Int, error) {
	return _Token.Contract.StorageVersion(&_Token.CallOpts)
}

// StorageVersion is a free data retrieval call binding the contract method 0x403ebd03.
//
// Solidity: function storageVersion() view returns(uint128)
func (_Token *TokenCallerSession) StorageVersion() (*big.Int, error) {
	return _Token.Contract.StorageVersion(&_Token.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Token *TokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Token *TokenSession) TotalSupply() (*big.Int, error) {
	return _Token.Contract.TotalSupply(&_Token.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Token *TokenCallerSession) TotalSupply() (*big.Int, error) {
	return _Token.Contract.TotalSupply(&_Token.CallOpts)
}

// TraderCommissionAddress is a free data retrieval call binding the contract method 0x29228fb4.
//
// Solidity: function traderCommissionAddress() view returns(address)
func (_Token *TokenCaller) TraderCommissionAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "traderCommissionAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TraderCommissionAddress is a free data retrieval call binding the contract method 0x29228fb4.
//
// Solidity: function traderCommissionAddress() view returns(address)
func (_Token *TokenSession) TraderCommissionAddress() (common.Address, error) {
	return _Token.Contract.TraderCommissionAddress(&_Token.CallOpts)
}

// TraderCommissionAddress is a free data retrieval call binding the contract method 0x29228fb4.
//
// Solidity: function traderCommissionAddress() view returns(address)
func (_Token *TokenCallerSession) TraderCommissionAddress() (common.Address, error) {
	return _Token.Contract.TraderCommissionAddress(&_Token.CallOpts)
}

// TraderCommissionBalance is a free data retrieval call binding the contract method 0x24538660.
//
// Solidity: function traderCommissionBalance() view returns(uint256)
func (_Token *TokenCaller) TraderCommissionBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "traderCommissionBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TraderCommissionBalance is a free data retrieval call binding the contract method 0x24538660.
//
// Solidity: function traderCommissionBalance() view returns(uint256)
func (_Token *TokenSession) TraderCommissionBalance() (*big.Int, error) {
	return _Token.Contract.TraderCommissionBalance(&_Token.CallOpts)
}

// TraderCommissionBalance is a free data retrieval call binding the contract method 0x24538660.
//
// Solidity: function traderCommissionBalance() view returns(uint256)
func (_Token *TokenCallerSession) TraderCommissionBalance() (*big.Int, error) {
	return _Token.Contract.TraderCommissionBalance(&_Token.CallOpts)
}

// TraderLiquidityBalance is a free data retrieval call binding the contract method 0x3eb0b312.
//
// Solidity: function traderLiquidityBalance() view returns(uint256)
func (_Token *TokenCaller) TraderLiquidityBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "traderLiquidityBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TraderLiquidityBalance is a free data retrieval call binding the contract method 0x3eb0b312.
//
// Solidity: function traderLiquidityBalance() view returns(uint256)
func (_Token *TokenSession) TraderLiquidityBalance() (*big.Int, error) {
	return _Token.Contract.TraderLiquidityBalance(&_Token.CallOpts)
}

// TraderLiquidityBalance is a free data retrieval call binding the contract method 0x3eb0b312.
//
// Solidity: function traderLiquidityBalance() view returns(uint256)
func (_Token *TokenCallerSession) TraderLiquidityBalance() (*big.Int, error) {
	return _Token.Contract.TraderLiquidityBalance(&_Token.CallOpts)
}

// TraderWhitelist is a free data retrieval call binding the contract method 0x6c1c4734.
//
// Solidity: function traderWhitelist(address ) view returns(bool)
func (_Token *TokenCaller) TraderWhitelist(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "traderWhitelist", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TraderWhitelist is a free data retrieval call binding the contract method 0x6c1c4734.
//
// Solidity: function traderWhitelist(address ) view returns(bool)
func (_Token *TokenSession) TraderWhitelist(arg0 common.Address) (bool, error) {
	return _Token.Contract.TraderWhitelist(&_Token.CallOpts, arg0)
}

// TraderWhitelist is a free data retrieval call binding the contract method 0x6c1c4734.
//
// Solidity: function traderWhitelist(address ) view returns(bool)
func (_Token *TokenCallerSession) TraderWhitelist(arg0 common.Address) (bool, error) {
	return _Token.Contract.TraderWhitelist(&_Token.CallOpts, arg0)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint32)
func (_Token *TokenCaller) Version(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint32)
func (_Token *TokenSession) Version() (uint32, error) {
	return _Token.Contract.Version(&_Token.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint32)
func (_Token *TokenCallerSession) Version() (uint32, error) {
	return _Token.Contract.Version(&_Token.CallOpts)
}

// AddInvestorAddress is a paid mutator transaction binding the contract method 0x97776f64.
//
// Solidity: function addInvestorAddress(address[] _investors) returns()
func (_Token *TokenTransactor) AddInvestorAddress(opts *bind.TransactOpts, _investors []common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "addInvestorAddress", _investors)
}

// AddInvestorAddress is a paid mutator transaction binding the contract method 0x97776f64.
//
// Solidity: function addInvestorAddress(address[] _investors) returns()
func (_Token *TokenSession) AddInvestorAddress(_investors []common.Address) (*types.Transaction, error) {
	return _Token.Contract.AddInvestorAddress(&_Token.TransactOpts, _investors)
}

// AddInvestorAddress is a paid mutator transaction binding the contract method 0x97776f64.
//
// Solidity: function addInvestorAddress(address[] _investors) returns()
func (_Token *TokenTransactorSession) AddInvestorAddress(_investors []common.Address) (*types.Transaction, error) {
	return _Token.Contract.AddInvestorAddress(&_Token.TransactOpts, _investors)
}

// AddTraderAddress is a paid mutator transaction binding the contract method 0xc6bd0fad.
//
// Solidity: function addTraderAddress(address _traderAddress) returns()
func (_Token *TokenTransactor) AddTraderAddress(opts *bind.TransactOpts, _traderAddress common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "addTraderAddress", _traderAddress)
}

// AddTraderAddress is a paid mutator transaction binding the contract method 0xc6bd0fad.
//
// Solidity: function addTraderAddress(address _traderAddress) returns()
func (_Token *TokenSession) AddTraderAddress(_traderAddress common.Address) (*types.Transaction, error) {
	return _Token.Contract.AddTraderAddress(&_Token.TransactOpts, _traderAddress)
}

// AddTraderAddress is a paid mutator transaction binding the contract method 0xc6bd0fad.
//
// Solidity: function addTraderAddress(address _traderAddress) returns()
func (_Token *TokenTransactorSession) AddTraderAddress(_traderAddress common.Address) (*types.Transaction, error) {
	return _Token.Contract.AddTraderAddress(&_Token.TransactOpts, _traderAddress)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_Token *TokenTransactor) Deposit(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "deposit", amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_Token *TokenSession) Deposit(amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Deposit(&_Token.TransactOpts, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_Token *TokenTransactorSession) Deposit(amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Deposit(&_Token.TransactOpts, amount)
}

// DepositETHTo is a paid mutator transaction binding the contract method 0xa66d8fad.
//
// Solidity: function depositETHTo(address to) payable returns()
func (_Token *TokenTransactor) DepositETHTo(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "depositETHTo", to)
}

// DepositETHTo is a paid mutator transaction binding the contract method 0xa66d8fad.
//
// Solidity: function depositETHTo(address to) payable returns()
func (_Token *TokenSession) DepositETHTo(to common.Address) (*types.Transaction, error) {
	return _Token.Contract.DepositETHTo(&_Token.TransactOpts, to)
}

// DepositETHTo is a paid mutator transaction binding the contract method 0xa66d8fad.
//
// Solidity: function depositETHTo(address to) payable returns()
func (_Token *TokenTransactorSession) DepositETHTo(to common.Address) (*types.Transaction, error) {
	return _Token.Contract.DepositETHTo(&_Token.TransactOpts, to)
}

// DepositTo is a paid mutator transaction binding the contract method 0x70aff70f.
//
// Solidity: function depositTo(uint256 amount, address to) returns()
func (_Token *TokenTransactor) DepositTo(opts *bind.TransactOpts, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "depositTo", amount, to)
}

// DepositTo is a paid mutator transaction binding the contract method 0x70aff70f.
//
// Solidity: function depositTo(uint256 amount, address to) returns()
func (_Token *TokenSession) DepositTo(amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _Token.Contract.DepositTo(&_Token.TransactOpts, amount, to)
}

// DepositTo is a paid mutator transaction binding the contract method 0x70aff70f.
//
// Solidity: function depositTo(uint256 amount, address to) returns()
func (_Token *TokenTransactorSession) DepositTo(amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _Token.Contract.DepositTo(&_Token.TransactOpts, amount, to)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Token *TokenTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Token *TokenSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Token.Contract.GrantRole(&_Token.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Token *TokenTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Token.Contract.GrantRole(&_Token.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x00cfb42c.
//
// Solidity: function initialize(address[9] iaddr, uint256 _commissions, bool _actual, bool _investorRestricted) returns()
func (_Token *TokenTransactor) Initialize(opts *bind.TransactOpts, iaddr [9]common.Address, _commissions *big.Int, _actual bool, _investorRestricted bool) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "initialize", iaddr, _commissions, _actual, _investorRestricted)
}

// Initialize is a paid mutator transaction binding the contract method 0x00cfb42c.
//
// Solidity: function initialize(address[9] iaddr, uint256 _commissions, bool _actual, bool _investorRestricted) returns()
func (_Token *TokenSession) Initialize(iaddr [9]common.Address, _commissions *big.Int, _actual bool, _investorRestricted bool) (*types.Transaction, error) {
	return _Token.Contract.Initialize(&_Token.TransactOpts, iaddr, _commissions, _actual, _investorRestricted)
}

// Initialize is a paid mutator transaction binding the contract method 0x00cfb42c.
//
// Solidity: function initialize(address[9] iaddr, uint256 _commissions, bool _actual, bool _investorRestricted) returns()
func (_Token *TokenTransactorSession) Initialize(iaddr [9]common.Address, _commissions *big.Int, _actual bool, _investorRestricted bool) (*types.Transaction, error) {
	return _Token.Contract.Initialize(&_Token.TransactOpts, iaddr, _commissions, _actual, _investorRestricted)
}

// InitiateExchangeOperation is a paid mutator transaction binding the contract method 0x4f083e8a.
//
// Solidity: function initiateExchangeOperation(address fromAsset, address toAsset, uint256 fromAmt, address caller, bytes _calldata) returns()
func (_Token *TokenTransactor) InitiateExchangeOperation(opts *bind.TransactOpts, fromAsset common.Address, toAsset common.Address, fromAmt *big.Int, caller common.Address, _calldata []byte) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "initiateExchangeOperation", fromAsset, toAsset, fromAmt, caller, _calldata)
}

// InitiateExchangeOperation is a paid mutator transaction binding the contract method 0x4f083e8a.
//
// Solidity: function initiateExchangeOperation(address fromAsset, address toAsset, uint256 fromAmt, address caller, bytes _calldata) returns()
func (_Token *TokenSession) InitiateExchangeOperation(fromAsset common.Address, toAsset common.Address, fromAmt *big.Int, caller common.Address, _calldata []byte) (*types.Transaction, error) {
	return _Token.Contract.InitiateExchangeOperation(&_Token.TransactOpts, fromAsset, toAsset, fromAmt, caller, _calldata)
}

// InitiateExchangeOperation is a paid mutator transaction binding the contract method 0x4f083e8a.
//
// Solidity: function initiateExchangeOperation(address fromAsset, address toAsset, uint256 fromAmt, address caller, bytes _calldata) returns()
func (_Token *TokenTransactorSession) InitiateExchangeOperation(fromAsset common.Address, toAsset common.Address, fromAmt *big.Int, caller common.Address, _calldata []byte) (*types.Transaction, error) {
	return _Token.Contract.InitiateExchangeOperation(&_Token.TransactOpts, fromAsset, toAsset, fromAmt, caller, _calldata)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Token *TokenTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Token *TokenSession) Pause() (*types.Transaction, error) {
	return _Token.Contract.Pause(&_Token.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Token *TokenTransactorSession) Pause() (*types.Transaction, error) {
	return _Token.Contract.Pause(&_Token.TransactOpts)
}

// RemoveInvestorAddress is a paid mutator transaction binding the contract method 0x619ced0f.
//
// Solidity: function removeInvestorAddress(address[] _investors) returns()
func (_Token *TokenTransactor) RemoveInvestorAddress(opts *bind.TransactOpts, _investors []common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "removeInvestorAddress", _investors)
}

// RemoveInvestorAddress is a paid mutator transaction binding the contract method 0x619ced0f.
//
// Solidity: function removeInvestorAddress(address[] _investors) returns()
func (_Token *TokenSession) RemoveInvestorAddress(_investors []common.Address) (*types.Transaction, error) {
	return _Token.Contract.RemoveInvestorAddress(&_Token.TransactOpts, _investors)
}

// RemoveInvestorAddress is a paid mutator transaction binding the contract method 0x619ced0f.
//
// Solidity: function removeInvestorAddress(address[] _investors) returns()
func (_Token *TokenTransactorSession) RemoveInvestorAddress(_investors []common.Address) (*types.Transaction, error) {
	return _Token.Contract.RemoveInvestorAddress(&_Token.TransactOpts, _investors)
}

// RemoveTraderAddress is a paid mutator transaction binding the contract method 0xcc429528.
//
// Solidity: function removeTraderAddress(address _traderAddress) returns()
func (_Token *TokenTransactor) RemoveTraderAddress(opts *bind.TransactOpts, _traderAddress common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "removeTraderAddress", _traderAddress)
}

// RemoveTraderAddress is a paid mutator transaction binding the contract method 0xcc429528.
//
// Solidity: function removeTraderAddress(address _traderAddress) returns()
func (_Token *TokenSession) RemoveTraderAddress(_traderAddress common.Address) (*types.Transaction, error) {
	return _Token.Contract.RemoveTraderAddress(&_Token.TransactOpts, _traderAddress)
}

// RemoveTraderAddress is a paid mutator transaction binding the contract method 0xcc429528.
//
// Solidity: function removeTraderAddress(address _traderAddress) returns()
func (_Token *TokenTransactorSession) RemoveTraderAddress(_traderAddress common.Address) (*types.Transaction, error) {
	return _Token.Contract.RemoveTraderAddress(&_Token.TransactOpts, _traderAddress)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Token *TokenTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Token *TokenSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Token.Contract.RenounceRole(&_Token.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Token *TokenTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Token.Contract.RenounceRole(&_Token.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Token *TokenTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Token *TokenSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Token.Contract.RevokeRole(&_Token.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Token *TokenTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Token.Contract.RevokeRole(&_Token.TransactOpts, role, account)
}

// SetCommission is a paid mutator transaction binding the contract method 0x5d252d82.
//
// Solidity: function setCommission(uint256 _type, uint16 _nom, uint16 _denom) returns()
func (_Token *TokenTransactor) SetCommission(opts *bind.TransactOpts, _type *big.Int, _nom uint16, _denom uint16) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "setCommission", _type, _nom, _denom)
}

// SetCommission is a paid mutator transaction binding the contract method 0x5d252d82.
//
// Solidity: function setCommission(uint256 _type, uint16 _nom, uint16 _denom) returns()
func (_Token *TokenSession) SetCommission(_type *big.Int, _nom uint16, _denom uint16) (*types.Transaction, error) {
	return _Token.Contract.SetCommission(&_Token.TransactOpts, _type, _nom, _denom)
}

// SetCommission is a paid mutator transaction binding the contract method 0x5d252d82.
//
// Solidity: function setCommission(uint256 _type, uint16 _nom, uint16 _denom) returns()
func (_Token *TokenTransactorSession) SetCommission(_type *big.Int, _nom uint16, _denom uint16) (*types.Transaction, error) {
	return _Token.Contract.SetCommission(&_Token.TransactOpts, _type, _nom, _denom)
}

// SetTraderCommissionAddress is a paid mutator transaction binding the contract method 0x90f9cb08.
//
// Solidity: function setTraderCommissionAddress(address _traderCommissionAddress) returns()
func (_Token *TokenTransactor) SetTraderCommissionAddress(opts *bind.TransactOpts, _traderCommissionAddress common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "setTraderCommissionAddress", _traderCommissionAddress)
}

// SetTraderCommissionAddress is a paid mutator transaction binding the contract method 0x90f9cb08.
//
// Solidity: function setTraderCommissionAddress(address _traderCommissionAddress) returns()
func (_Token *TokenSession) SetTraderCommissionAddress(_traderCommissionAddress common.Address) (*types.Transaction, error) {
	return _Token.Contract.SetTraderCommissionAddress(&_Token.TransactOpts, _traderCommissionAddress)
}

// SetTraderCommissionAddress is a paid mutator transaction binding the contract method 0x90f9cb08.
//
// Solidity: function setTraderCommissionAddress(address _traderCommissionAddress) returns()
func (_Token *TokenTransactorSession) SetTraderCommissionAddress(_traderCommissionAddress common.Address) (*types.Transaction, error) {
	return _Token.Contract.SetTraderCommissionAddress(&_Token.TransactOpts, _traderCommissionAddress)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Token *TokenTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Token *TokenSession) Unpause() (*types.Transaction, error) {
	return _Token.Contract.Unpause(&_Token.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Token *TokenTransactorSession) Unpause() (*types.Transaction, error) {
	return _Token.Contract.Unpause(&_Token.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Token *TokenTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Token *TokenSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Withdraw(&_Token.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Token *TokenTransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Withdraw(&_Token.TransactOpts, amount)
}

// WithdrawDexeCommission is a paid mutator transaction binding the contract method 0xad3793fb.
//
// Solidity: function withdrawDexeCommission(uint256 amount) returns()
func (_Token *TokenTransactor) WithdrawDexeCommission(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "withdrawDexeCommission", amount)
}

// WithdrawDexeCommission is a paid mutator transaction binding the contract method 0xad3793fb.
//
// Solidity: function withdrawDexeCommission(uint256 amount) returns()
func (_Token *TokenSession) WithdrawDexeCommission(amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.WithdrawDexeCommission(&_Token.TransactOpts, amount)
}

// WithdrawDexeCommission is a paid mutator transaction binding the contract method 0xad3793fb.
//
// Solidity: function withdrawDexeCommission(uint256 amount) returns()
func (_Token *TokenTransactorSession) WithdrawDexeCommission(amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.WithdrawDexeCommission(&_Token.TransactOpts, amount)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0xf14210a6.
//
// Solidity: function withdrawETH(uint256 amount) returns()
func (_Token *TokenTransactor) WithdrawETH(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "withdrawETH", amount)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0xf14210a6.
//
// Solidity: function withdrawETH(uint256 amount) returns()
func (_Token *TokenSession) WithdrawETH(amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.WithdrawETH(&_Token.TransactOpts, amount)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0xf14210a6.
//
// Solidity: function withdrawETH(uint256 amount) returns()
func (_Token *TokenTransactorSession) WithdrawETH(amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.WithdrawETH(&_Token.TransactOpts, amount)
}

// WithdrawETHTo is a paid mutator transaction binding the contract method 0xdca3a041.
//
// Solidity: function withdrawETHTo(uint256 amount, address to) returns()
func (_Token *TokenTransactor) WithdrawETHTo(opts *bind.TransactOpts, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "withdrawETHTo", amount, to)
}

// WithdrawETHTo is a paid mutator transaction binding the contract method 0xdca3a041.
//
// Solidity: function withdrawETHTo(uint256 amount, address to) returns()
func (_Token *TokenSession) WithdrawETHTo(amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _Token.Contract.WithdrawETHTo(&_Token.TransactOpts, amount, to)
}

// WithdrawETHTo is a paid mutator transaction binding the contract method 0xdca3a041.
//
// Solidity: function withdrawETHTo(uint256 amount, address to) returns()
func (_Token *TokenTransactorSession) WithdrawETHTo(amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _Token.Contract.WithdrawETHTo(&_Token.TransactOpts, amount, to)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0xc86283c8.
//
// Solidity: function withdrawTo(uint256 amount, address to) returns()
func (_Token *TokenTransactor) WithdrawTo(opts *bind.TransactOpts, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "withdrawTo", amount, to)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0xc86283c8.
//
// Solidity: function withdrawTo(uint256 amount, address to) returns()
func (_Token *TokenSession) WithdrawTo(amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _Token.Contract.WithdrawTo(&_Token.TransactOpts, amount, to)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0xc86283c8.
//
// Solidity: function withdrawTo(uint256 amount, address to) returns()
func (_Token *TokenTransactorSession) WithdrawTo(amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _Token.Contract.WithdrawTo(&_Token.TransactOpts, amount, to)
}

// WithdrawTraderCommission is a paid mutator transaction binding the contract method 0x0f36e00b.
//
// Solidity: function withdrawTraderCommission(uint256 amount) returns()
func (_Token *TokenTransactor) WithdrawTraderCommission(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "withdrawTraderCommission", amount)
}

// WithdrawTraderCommission is a paid mutator transaction binding the contract method 0x0f36e00b.
//
// Solidity: function withdrawTraderCommission(uint256 amount) returns()
func (_Token *TokenSession) WithdrawTraderCommission(amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.WithdrawTraderCommission(&_Token.TransactOpts, amount)
}

// WithdrawTraderCommission is a paid mutator transaction binding the contract method 0x0f36e00b.
//
// Solidity: function withdrawTraderCommission(uint256 amount) returns()
func (_Token *TokenTransactorSession) WithdrawTraderCommission(amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.WithdrawTraderCommission(&_Token.TransactOpts, amount)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Token *TokenTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Token.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Token *TokenSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Token.Contract.Fallback(&_Token.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Token *TokenTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Token.Contract.Fallback(&_Token.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Token *TokenTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Token *TokenSession) Receive() (*types.Transaction, error) {
	return _Token.Contract.Receive(&_Token.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Token *TokenTransactorSession) Receive() (*types.Transaction, error) {
	return _Token.Contract.Receive(&_Token.TransactOpts)
}

// TokenDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Token contract.
type TokenDepositIterator struct {
	Event *TokenDeposit // Event containing the contract specifics and raw log

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
func (it *TokenDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenDeposit)
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
		it.Event = new(TokenDeposit)
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
func (it *TokenDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenDeposit represents a Deposit event raised by the Token contract.
type TokenDeposit struct {
	Who       common.Address
	AmountBT  *big.Int
	Liquidity *big.Int
	Price     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x70d4722f7c84cf1ff653224133fa245bde8c50ee9ce56636429f4ee954445a11.
//
// Solidity: event Deposit(address indexed who, uint256 amountBT, uint256 liquidity, int128 price)
func (_Token *TokenFilterer) FilterDeposit(opts *bind.FilterOpts, who []common.Address) (*TokenDepositIterator, error) {

	var whoRule []interface{}
	for _, whoItem := range who {
		whoRule = append(whoRule, whoItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "Deposit", whoRule)
	if err != nil {
		return nil, err
	}
	return &TokenDepositIterator{contract: _Token.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x70d4722f7c84cf1ff653224133fa245bde8c50ee9ce56636429f4ee954445a11.
//
// Solidity: event Deposit(address indexed who, uint256 amountBT, uint256 liquidity, int128 price)
func (_Token *TokenFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *TokenDeposit, who []common.Address) (event.Subscription, error) {

	var whoRule []interface{}
	for _, whoItem := range who {
		whoRule = append(whoRule, whoItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "Deposit", whoRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenDeposit)
				if err := _Token.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0x70d4722f7c84cf1ff653224133fa245bde8c50ee9ce56636429f4ee954445a11.
//
// Solidity: event Deposit(address indexed who, uint256 amountBT, uint256 liquidity, int128 price)
func (_Token *TokenFilterer) ParseDeposit(log types.Log) (*TokenDeposit, error) {
	event := new(TokenDeposit)
	if err := _Token.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenExchangedIterator is returned from FilterExchanged and is used to iterate over the raw logs and unpacked data for Exchanged events raised by the Token contract.
type TokenExchangedIterator struct {
	Event *TokenExchanged // Event containing the contract specifics and raw log

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
func (it *TokenExchangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenExchanged)
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
		it.Event = new(TokenExchanged)
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
func (it *TokenExchangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenExchangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenExchanged represents a Exchanged event raised by the Token contract.
type TokenExchanged struct {
	FromAsset common.Address
	ToAsset   common.Address
	FromAmt   *big.Int
	ToAmt     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterExchanged is a free log retrieval operation binding the contract event 0x9a406bc63ac2307f5ee52002fb44c825e70ce5c177ca9cd148425d32053ed68d.
//
// Solidity: event Exchanged(address fromAsset, address toAsset, uint256 fromAmt, uint256 toAmt)
func (_Token *TokenFilterer) FilterExchanged(opts *bind.FilterOpts) (*TokenExchangedIterator, error) {

	logs, sub, err := _Token.contract.FilterLogs(opts, "Exchanged")
	if err != nil {
		return nil, err
	}
	return &TokenExchangedIterator{contract: _Token.contract, event: "Exchanged", logs: logs, sub: sub}, nil
}

// WatchExchanged is a free log subscription operation binding the contract event 0x9a406bc63ac2307f5ee52002fb44c825e70ce5c177ca9cd148425d32053ed68d.
//
// Solidity: event Exchanged(address fromAsset, address toAsset, uint256 fromAmt, uint256 toAmt)
func (_Token *TokenFilterer) WatchExchanged(opts *bind.WatchOpts, sink chan<- *TokenExchanged) (event.Subscription, error) {

	logs, sub, err := _Token.contract.WatchLogs(opts, "Exchanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenExchanged)
				if err := _Token.contract.UnpackLog(event, "Exchanged", log); err != nil {
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

// ParseExchanged is a log parse operation binding the contract event 0x9a406bc63ac2307f5ee52002fb44c825e70ce5c177ca9cd148425d32053ed68d.
//
// Solidity: event Exchanged(address fromAsset, address toAsset, uint256 fromAmt, uint256 toAmt)
func (_Token *TokenFilterer) ParseExchanged(log types.Log) (*TokenExchanged, error) {
	event := new(TokenExchanged)
	if err := _Token.contract.UnpackLog(event, "Exchanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenLossIterator is returned from FilterLoss and is used to iterate over the raw logs and unpacked data for Loss events raised by the Token contract.
type TokenLossIterator struct {
	Event *TokenLoss // Event containing the contract specifics and raw log

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
func (it *TokenLossIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenLoss)
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
		it.Event = new(TokenLoss)
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
func (it *TokenLossIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenLossIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenLoss represents a Loss event raised by the Token contract.
type TokenLoss struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLoss is a free log retrieval operation binding the contract event 0xb65b2e085d7d040c313f7d4e1ac90f5937026fee497e0e24a7eff16a55e1c5ea.
//
// Solidity: event Loss(uint256 amount)
func (_Token *TokenFilterer) FilterLoss(opts *bind.FilterOpts) (*TokenLossIterator, error) {

	logs, sub, err := _Token.contract.FilterLogs(opts, "Loss")
	if err != nil {
		return nil, err
	}
	return &TokenLossIterator{contract: _Token.contract, event: "Loss", logs: logs, sub: sub}, nil
}

// WatchLoss is a free log subscription operation binding the contract event 0xb65b2e085d7d040c313f7d4e1ac90f5937026fee497e0e24a7eff16a55e1c5ea.
//
// Solidity: event Loss(uint256 amount)
func (_Token *TokenFilterer) WatchLoss(opts *bind.WatchOpts, sink chan<- *TokenLoss) (event.Subscription, error) {

	logs, sub, err := _Token.contract.WatchLogs(opts, "Loss")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenLoss)
				if err := _Token.contract.UnpackLog(event, "Loss", log); err != nil {
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

// ParseLoss is a log parse operation binding the contract event 0xb65b2e085d7d040c313f7d4e1ac90f5937026fee497e0e24a7eff16a55e1c5ea.
//
// Solidity: event Loss(uint256 amount)
func (_Token *TokenFilterer) ParseLoss(log types.Log) (*TokenLoss, error) {
	event := new(TokenLoss)
	if err := _Token.contract.UnpackLog(event, "Loss", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Token contract.
type TokenPausedIterator struct {
	Event *TokenPaused // Event containing the contract specifics and raw log

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
func (it *TokenPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenPaused)
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
		it.Event = new(TokenPaused)
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
func (it *TokenPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenPaused represents a Paused event raised by the Token contract.
type TokenPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Token *TokenFilterer) FilterPaused(opts *bind.FilterOpts) (*TokenPausedIterator, error) {

	logs, sub, err := _Token.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &TokenPausedIterator{contract: _Token.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Token *TokenFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *TokenPaused) (event.Subscription, error) {

	logs, sub, err := _Token.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenPaused)
				if err := _Token.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Token *TokenFilterer) ParsePaused(log types.Log) (*TokenPaused, error) {
	event := new(TokenPaused)
	if err := _Token.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenProfitIterator is returned from FilterProfit and is used to iterate over the raw logs and unpacked data for Profit events raised by the Token contract.
type TokenProfitIterator struct {
	Event *TokenProfit // Event containing the contract specifics and raw log

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
func (it *TokenProfitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenProfit)
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
		it.Event = new(TokenProfit)
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
func (it *TokenProfitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenProfitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenProfit represents a Profit event raised by the Token contract.
type TokenProfit struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterProfit is a free log retrieval operation binding the contract event 0x357d905f1831209797df4d55d79c5c5bf1d9f7311c976afd05e13d881eab9bc8.
//
// Solidity: event Profit(uint256 amount)
func (_Token *TokenFilterer) FilterProfit(opts *bind.FilterOpts) (*TokenProfitIterator, error) {

	logs, sub, err := _Token.contract.FilterLogs(opts, "Profit")
	if err != nil {
		return nil, err
	}
	return &TokenProfitIterator{contract: _Token.contract, event: "Profit", logs: logs, sub: sub}, nil
}

// WatchProfit is a free log subscription operation binding the contract event 0x357d905f1831209797df4d55d79c5c5bf1d9f7311c976afd05e13d881eab9bc8.
//
// Solidity: event Profit(uint256 amount)
func (_Token *TokenFilterer) WatchProfit(opts *bind.WatchOpts, sink chan<- *TokenProfit) (event.Subscription, error) {

	logs, sub, err := _Token.contract.WatchLogs(opts, "Profit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenProfit)
				if err := _Token.contract.UnpackLog(event, "Profit", log); err != nil {
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

// ParseProfit is a log parse operation binding the contract event 0x357d905f1831209797df4d55d79c5c5bf1d9f7311c976afd05e13d881eab9bc8.
//
// Solidity: event Profit(uint256 amount)
func (_Token *TokenFilterer) ParseProfit(log types.Log) (*TokenProfit, error) {
	event := new(TokenProfit)
	if err := _Token.contract.UnpackLog(event, "Profit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Token contract.
type TokenRoleAdminChangedIterator struct {
	Event *TokenRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *TokenRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenRoleAdminChanged)
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
		it.Event = new(TokenRoleAdminChanged)
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
func (it *TokenRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenRoleAdminChanged represents a RoleAdminChanged event raised by the Token contract.
type TokenRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Token *TokenFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*TokenRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &TokenRoleAdminChangedIterator{contract: _Token.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Token *TokenFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *TokenRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenRoleAdminChanged)
				if err := _Token.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Token *TokenFilterer) ParseRoleAdminChanged(log types.Log) (*TokenRoleAdminChanged, error) {
	event := new(TokenRoleAdminChanged)
	if err := _Token.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Token contract.
type TokenRoleGrantedIterator struct {
	Event *TokenRoleGranted // Event containing the contract specifics and raw log

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
func (it *TokenRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenRoleGranted)
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
		it.Event = new(TokenRoleGranted)
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
func (it *TokenRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenRoleGranted represents a RoleGranted event raised by the Token contract.
type TokenRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Token *TokenFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*TokenRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &TokenRoleGrantedIterator{contract: _Token.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Token *TokenFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *TokenRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenRoleGranted)
				if err := _Token.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Token *TokenFilterer) ParseRoleGranted(log types.Log) (*TokenRoleGranted, error) {
	event := new(TokenRoleGranted)
	if err := _Token.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Token contract.
type TokenRoleRevokedIterator struct {
	Event *TokenRoleRevoked // Event containing the contract specifics and raw log

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
func (it *TokenRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenRoleRevoked)
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
		it.Event = new(TokenRoleRevoked)
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
func (it *TokenRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenRoleRevoked represents a RoleRevoked event raised by the Token contract.
type TokenRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Token *TokenFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*TokenRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &TokenRoleRevokedIterator{contract: _Token.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Token *TokenFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *TokenRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenRoleRevoked)
				if err := _Token.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Token *TokenFilterer) ParseRoleRevoked(log types.Log) (*TokenRoleRevoked, error) {
	event := new(TokenRoleRevoked)
	if err := _Token.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Token contract.
type TokenUnpausedIterator struct {
	Event *TokenUnpaused // Event containing the contract specifics and raw log

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
func (it *TokenUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenUnpaused)
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
		it.Event = new(TokenUnpaused)
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
func (it *TokenUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenUnpaused represents a Unpaused event raised by the Token contract.
type TokenUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Token *TokenFilterer) FilterUnpaused(opts *bind.FilterOpts) (*TokenUnpausedIterator, error) {

	logs, sub, err := _Token.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &TokenUnpausedIterator{contract: _Token.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Token *TokenFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *TokenUnpaused) (event.Subscription, error) {

	logs, sub, err := _Token.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenUnpaused)
				if err := _Token.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Token *TokenFilterer) ParseUnpaused(log types.Log) (*TokenUnpaused, error) {
	event := new(TokenUnpaused)
	if err := _Token.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Token contract.
type TokenWithdrawIterator struct {
	Event *TokenWithdraw // Event containing the contract specifics and raw log

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
func (it *TokenWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenWithdraw)
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
		it.Event = new(TokenWithdraw)
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
func (it *TokenWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenWithdraw represents a Withdraw event raised by the Token contract.
type TokenWithdraw struct {
	Who       common.Address
	AmountBT  *big.Int
	Liquidity *big.Int
	Commision *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x02f25270a4d87bea75db541cdfe559334a275b4a233520ed6c0a2429667cca94.
//
// Solidity: event Withdraw(address indexed who, uint256 amountBT, uint256 liquidity, uint256 commision)
func (_Token *TokenFilterer) FilterWithdraw(opts *bind.FilterOpts, who []common.Address) (*TokenWithdrawIterator, error) {

	var whoRule []interface{}
	for _, whoItem := range who {
		whoRule = append(whoRule, whoItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "Withdraw", whoRule)
	if err != nil {
		return nil, err
	}
	return &TokenWithdrawIterator{contract: _Token.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x02f25270a4d87bea75db541cdfe559334a275b4a233520ed6c0a2429667cca94.
//
// Solidity: event Withdraw(address indexed who, uint256 amountBT, uint256 liquidity, uint256 commision)
func (_Token *TokenFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *TokenWithdraw, who []common.Address) (event.Subscription, error) {

	var whoRule []interface{}
	for _, whoItem := range who {
		whoRule = append(whoRule, whoItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "Withdraw", whoRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenWithdraw)
				if err := _Token.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x02f25270a4d87bea75db541cdfe559334a275b4a233520ed6c0a2429667cca94.
//
// Solidity: event Withdraw(address indexed who, uint256 amountBT, uint256 liquidity, uint256 commision)
func (_Token *TokenFilterer) ParseWithdraw(log types.Log) (*TokenWithdraw, error) {
	event := new(TokenWithdraw)
	if err := _Token.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
