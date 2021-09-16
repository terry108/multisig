// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dms_abi

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

// AddressABI is the input ABI used to generate the binding from.
const AddressABI = "[]"

// AddressBin is the compiled bytecode used for deploying new contracts.
var AddressBin = "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220596fa053f51153fbf511adb167c54f1b2fddcc75eea1f1e89fa6e40613ff6d0b64736f6c63430008000033"

// DeployAddress deploys a new Ethereum contract, binding an instance of Address to it.
func DeployAddress(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Address, error) {
	parsed, err := abi.JSON(strings.NewReader(AddressABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AddressBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Address{AddressCaller: AddressCaller{contract: contract}, AddressTransactor: AddressTransactor{contract: contract}, AddressFilterer: AddressFilterer{contract: contract}}, nil
}

// Address is an auto generated Go binding around an Ethereum contract.
type Address struct {
	AddressCaller     // Read-only binding to the contract
	AddressTransactor // Write-only binding to the contract
	AddressFilterer   // Log filterer for contract events
}

// AddressCaller is an auto generated read-only Go binding around an Ethereum contract.
type AddressCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AddressTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AddressFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AddressSession struct {
	Contract     *Address          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AddressCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AddressCallerSession struct {
	Contract *AddressCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// AddressTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AddressTransactorSession struct {
	Contract     *AddressTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// AddressRaw is an auto generated low-level Go binding around an Ethereum contract.
type AddressRaw struct {
	Contract *Address // Generic contract binding to access the raw methods on
}

// AddressCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AddressCallerRaw struct {
	Contract *AddressCaller // Generic read-only contract binding to access the raw methods on
}

// AddressTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AddressTransactorRaw struct {
	Contract *AddressTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAddress creates a new instance of Address, bound to a specific deployed contract.
func NewAddress(address common.Address, backend bind.ContractBackend) (*Address, error) {
	contract, err := bindAddress(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Address{AddressCaller: AddressCaller{contract: contract}, AddressTransactor: AddressTransactor{contract: contract}, AddressFilterer: AddressFilterer{contract: contract}}, nil
}

// NewAddressCaller creates a new read-only instance of Address, bound to a specific deployed contract.
func NewAddressCaller(address common.Address, caller bind.ContractCaller) (*AddressCaller, error) {
	contract, err := bindAddress(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AddressCaller{contract: contract}, nil
}

// NewAddressTransactor creates a new write-only instance of Address, bound to a specific deployed contract.
func NewAddressTransactor(address common.Address, transactor bind.ContractTransactor) (*AddressTransactor, error) {
	contract, err := bindAddress(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AddressTransactor{contract: contract}, nil
}

// NewAddressFilterer creates a new log filterer instance of Address, bound to a specific deployed contract.
func NewAddressFilterer(address common.Address, filterer bind.ContractFilterer) (*AddressFilterer, error) {
	contract, err := bindAddress(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AddressFilterer{contract: contract}, nil
}

// bindAddress binds a generic wrapper to an already deployed contract.
func bindAddress(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AddressABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Address *AddressRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Address.Contract.AddressCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Address *AddressRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Address.Contract.AddressTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Address *AddressRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Address.Contract.AddressTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Address *AddressCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Address.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Address *AddressTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Address.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Address *AddressTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Address.Contract.contract.Transact(opts, method, params...)
}

// BytesLibABI is the input ABI used to generate the binding from.
const BytesLibABI = "[]"

// BytesLibBin is the compiled bytecode used for deploying new contracts.
var BytesLibBin = "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212209c583abbc354b4d8568a91c7cb0f159ede02f41b066a3367999de2f55fc59b7264736f6c63430008000033"

// DeployBytesLib deploys a new Ethereum contract, binding an instance of BytesLib to it.
func DeployBytesLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BytesLib, error) {
	parsed, err := abi.JSON(strings.NewReader(BytesLibABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BytesLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BytesLib{BytesLibCaller: BytesLibCaller{contract: contract}, BytesLibTransactor: BytesLibTransactor{contract: contract}, BytesLibFilterer: BytesLibFilterer{contract: contract}}, nil
}

// BytesLib is an auto generated Go binding around an Ethereum contract.
type BytesLib struct {
	BytesLibCaller     // Read-only binding to the contract
	BytesLibTransactor // Write-only binding to the contract
	BytesLibFilterer   // Log filterer for contract events
}

// BytesLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type BytesLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BytesLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BytesLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BytesLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BytesLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BytesLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BytesLibSession struct {
	Contract     *BytesLib         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BytesLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BytesLibCallerSession struct {
	Contract *BytesLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// BytesLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BytesLibTransactorSession struct {
	Contract     *BytesLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// BytesLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type BytesLibRaw struct {
	Contract *BytesLib // Generic contract binding to access the raw methods on
}

// BytesLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BytesLibCallerRaw struct {
	Contract *BytesLibCaller // Generic read-only contract binding to access the raw methods on
}

// BytesLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BytesLibTransactorRaw struct {
	Contract *BytesLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBytesLib creates a new instance of BytesLib, bound to a specific deployed contract.
func NewBytesLib(address common.Address, backend bind.ContractBackend) (*BytesLib, error) {
	contract, err := bindBytesLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BytesLib{BytesLibCaller: BytesLibCaller{contract: contract}, BytesLibTransactor: BytesLibTransactor{contract: contract}, BytesLibFilterer: BytesLibFilterer{contract: contract}}, nil
}

// NewBytesLibCaller creates a new read-only instance of BytesLib, bound to a specific deployed contract.
func NewBytesLibCaller(address common.Address, caller bind.ContractCaller) (*BytesLibCaller, error) {
	contract, err := bindBytesLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BytesLibCaller{contract: contract}, nil
}

// NewBytesLibTransactor creates a new write-only instance of BytesLib, bound to a specific deployed contract.
func NewBytesLibTransactor(address common.Address, transactor bind.ContractTransactor) (*BytesLibTransactor, error) {
	contract, err := bindBytesLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BytesLibTransactor{contract: contract}, nil
}

// NewBytesLibFilterer creates a new log filterer instance of BytesLib, bound to a specific deployed contract.
func NewBytesLibFilterer(address common.Address, filterer bind.ContractFilterer) (*BytesLibFilterer, error) {
	contract, err := bindBytesLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BytesLibFilterer{contract: contract}, nil
}

// bindBytesLib binds a generic wrapper to an already deployed contract.
func bindBytesLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BytesLibABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BytesLib *BytesLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BytesLib.Contract.BytesLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BytesLib *BytesLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BytesLib.Contract.BytesLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BytesLib *BytesLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BytesLib.Contract.BytesLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BytesLib *BytesLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BytesLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BytesLib *BytesLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BytesLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BytesLib *BytesLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BytesLib.Contract.contract.Transact(opts, method, params...)
}

// DynamicMultiSigABI is the input ABI used to generate the binding from.
const DynamicMultiSigABI = "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_managers\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"to\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"ERC20\",\"type\":\"address\"}],\"name\":\"CrossOutFunds\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"to\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"ERC721\",\"type\":\"address\"}],\"name\":\"CrossOutNFT\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DepositFunds\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TransferFunds\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"txKey\",\"type\":\"string\"}],\"name\":\"TxManagerChangeCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"txKey\",\"type\":\"string\"}],\"name\":\"TxUpgradeCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"txKey\",\"type\":\"string\"}],\"name\":\"TxWithdrawCompleted\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"allManagers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"closeUpgrade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"txKey\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"adds\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"removes\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"count\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"signatures\",\"type\":\"bytes\"}],\"name\":\"createOrSignManagerChange\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"txKey\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"upgradeContract\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signatures\",\"type\":\"bytes\"}],\"name\":\"createOrSignUpgrade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"txKey\",\"type\":\"string\"},{\"internalType\":\"addresspayable\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isERC20\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"ERC20\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signatures\",\"type\":\"bytes\"}],\"name\":\"createOrSignWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"txKey\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"ERC721\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"signatures\",\"type\":\"bytes\"}],\"name\":\"createOrSignWithdrawNFT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"to\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"ERC20\",\"type\":\"address\"}],\"name\":\"crossOut\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"to\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"ERC721\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"crossOutNFT\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"current_min_signatures\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_manager\",\"type\":\"address\"}],\"name\":\"ifManager\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"txKey\",\"type\":\"string\"}],\"name\":\"isCompletedTx\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ERC20\",\"type\":\"address\"}],\"name\":\"isMinterERC20\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ERC721\",\"type\":\"address\"}],\"name\":\"isMinterERC721\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"max_managers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"min_managers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ERC20\",\"type\":\"address\"}],\"name\":\"registerMinterERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ERC721\",\"type\":\"address\"}],\"name\":\"registerMinterERC721\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"signatureLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ERC20\",\"type\":\"address\"}],\"name\":\"unregisterMinterERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ERC721\",\"type\":\"address\"}],\"name\":\"unregisterMinterERC721\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"upgrade\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"upgradeContractAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"upgradeContractS1\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ERC20\",\"type\":\"address\"}],\"name\":\"upgradeContractS2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// DynamicMultiSigFuncSigs maps the 4-byte function signature to its string representation.
var DynamicMultiSigFuncSigs = map[string]string{
	"9c30b35e": "allManagers()",
	"d4cacbaa": "closeUpgrade()",
	"00719226": "createOrSignManagerChange(string,address[],address[],uint8,bytes)",
	"408e8b7a": "createOrSignUpgrade(string,address,bytes)",
	"ab6c2b10": "createOrSignWithdraw(string,address,uint256,bool,address,bytes)",
	"1c702f35": "createOrSignWithdrawNFT(string,address,address,uint256,string,bytes)",
	"0889d1f0": "crossOut(string,uint256,address)",
	"c76bbdc5": "crossOutNFT(string,address,uint256)",
	"e079cee9": "current_min_signatures()",
	"75173b70": "ifManager(address)",
	"a5e399b3": "isCompletedTx(string)",
	"6a7142e1": "isMinterERC20(address)",
	"279fb6e0": "isMinterERC721(address)",
	"f7f2ff74": "max_managers()",
	"ad4b61a8": "min_managers()",
	"150b7a02": "onERC721Received(address,address,uint256,bytes)",
	"8da5cb5b": "owner()",
	"2c4e722e": "rate()",
	"34774b71": "registerMinterERC20(address)",
	"7de24876": "registerMinterERC721(address)",
	"1b9a9323": "signatureLength()",
	"9dcdc978": "unregisterMinterERC20(address)",
	"bc8870ed": "unregisterMinterERC721(address)",
	"d55ec697": "upgrade()",
	"30b2d84d": "upgradeContractAddress()",
	"1dda9c05": "upgradeContractS1()",
	"5bda3fcf": "upgradeContractS2(address)",
}

// DynamicMultiSigBin is the compiled bytecode used for deploying new contracts.
var DynamicMultiSigBin = "0x6080604052600080546001600160a81b0319169055600f600155600360028190556042905560416004553480156200003657600080fd5b506040516200449138038062004491833981016040819052620000599162000391565b60015481511115620000885760405162461bcd60e51b81526004016200007f90620004dd565b60405180910390fd5b60025481511015620000ae5760405162461bcd60e51b81526004016200007f906200045f565b60058054610100600160a81b03191633610100021790558051620000da906009906020840190620002f8565b5060005b60095460ff82161015620002215760016008600060098460ff16815481106200011757634e487b7160e01b600052603260045260246000fd5b6000918252602080832091909101546001600160a01b031683528201929092526040018120805460ff191660ff9384161790556009805460019360069392919086169081106200017757634e487b7160e01b600052603260045260246000fd5b6000918252602080832091909101546001600160a01b031683528201929092526040019020805460ff191660ff928316179055600980546007928416908110620001d157634e487b7160e01b600052603260045260246000fd5b60009182526020808320909101548354600181018555938352912090910180546001600160a01b0319166001600160a01b03909216919091179055806200021881620005e2565b915050620000de565b5060055461010090046001600160a01b031660009081526008602052604090205460ff1615620002655760405162461bcd60e51b81526004016200007f9062000524565b60095462000273906200028e565b6005805460ff191660ff929092169190911790555062000631565b6000808211620002b25760405162461bcd60e51b81526004016200007f90620004a6565b60006001606484600354620002c89190620005a6565b620002d491906200056a565b620002e09190620005c8565b9050620002ef60648262000585565b9150505b919050565b82805482825590600052602060002090810192821562000350579160200282015b828111156200035057825182546001600160a01b0319166001600160a01b0390911617825560209092019160019091019062000319565b506200035e92915062000362565b5090565b5b808211156200035e576000815560010162000363565b80516001600160a01b0381168114620002f357600080fd5b60006020808385031215620003a4578182fd5b82516001600160401b0380821115620003bb578384fd5b818501915085601f830112620003cf578384fd5b815181811115620003e457620003e46200061b565b838102604051858282010181811085821117156200040657620004066200061b565b604052828152858101935084860182860187018a101562000425578788fd5b8795505b8386101562000452576200043d8162000379565b85526001959095019493860193860162000429565b5098975050505050505050565b60208082526027908201527f4e6f74207265616368696e6720746865206d696e206e756d626572206f66206d604082015266616e616765727360c81b606082015260800190565b60208082526014908201527f4d616e616765722043616e277420656d7074792e000000000000000000000000604082015260600190565b60208082526027908201527f457863656564656420746865206d6178696d756d206e756d626572206f66206d604082015266616e616765727360c81b606082015260800190565b60208082526026908201527f436f6e74726163742063726561746f722063616e6e6f7420616374206173206d60408201526530b730b3b2b960d11b606082015260800190565b6000821982111562000580576200058062000605565b500190565b600082620005a157634e487b7160e01b81526012600452602481fd5b500490565b6000816000190483118215151615620005c357620005c362000605565b500290565b600082821015620005dd57620005dd62000605565b500390565b600060ff821660ff811415620005fc57620005fc62000605565b60010192915050565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052604160045260246000fd5b613e5080620006416000396000f3fe60806040526004361061019f5760003560e01c806375173b70116100ec578063ad4b61a81161008a578063d4cacbaa11610064578063d4cacbaa146104b1578063d55ec697146104c6578063e079cee9146104db578063f7f2ff74146104fd576101df565b8063ad4b61a81461045c578063bc8870ed14610471578063c76bbdc514610491576101df565b80639c30b35e116100c65780639c30b35e146103da5780639dcdc978146103fc578063a5e399b31461041c578063ab6c2b101461043c576101df565b806375173b70146103855780637de24876146103a55780638da5cb5b146103c5576101df565b8063279fb6e01161015957806334774b711161013357806334774b7114610305578063408e8b7a146103255780635bda3fcf146103455780636a7142e114610365576101df565b8063279fb6e0146102ae5780632c4e722e146102ce57806330b2d84d146102e3576101df565b8062719226146101e15780630889d1f014610201578063150b7a021461022a5780631b9a9323146102575780631c702f35146102795780631dda9c0514610299576101df565b366101df577fd241e73300212f6df233a8e6d3146b88a9d4964e06621d54b5ff6afeba7b1b8833346040516101d592919061304b565b60405180910390a1005b005b3480156101ed57600080fd5b506101df6101fc366004612e79565b610512565b61021461020f366004612f3b565b6106ec565b6040516102219190613114565b60405180910390f35b34801561023657600080fd5b5061024a610245366004612b64565b6109a3565b604051610221919061313d565b34801561026357600080fd5b5061026c6109b3565b6040516102219190613c90565b34801561028557600080fd5b506101df610294366004612cc4565b6109b9565b3480156102a557600080fd5b506101df610b9a565b3480156102ba57600080fd5b506102146102c9366004612b2c565b610c5c565b3480156102da57600080fd5b5061026c610c80565b3480156102ef57600080fd5b506102f8610c86565b6040516102219190613037565b34801561031157600080fd5b506101df610320366004612b2c565b610c9a565b34801561033157600080fd5b506101df610340366004612dac565b610d6a565b34801561035157600080fd5b506101df610360366004612b2c565b610f49565b34801561037157600080fd5b50610214610380366004612b2c565b61113f565b34801561039157600080fd5b506102146103a0366004612b2c565b61115f565b3480156103b157600080fd5b506101df6103c0366004612b2c565b611180565b3480156103d157600080fd5b506102f8611250565b3480156103e657600080fd5b506103ef611264565b6040516102219190613101565b34801561040857600080fd5b506101df610417366004612b2c565b6112c6565b34801561042857600080fd5b50610214610437366004612bea565b61133b565b34801561044857600080fd5b506101df610457366004612c1d565b611369565b34801561046857600080fd5b5061026c6115f7565b34801561047d57600080fd5b506101df61048c366004612b2c565b6115fd565b34801561049d57600080fd5b506102146104ac366004612e21565b611672565b3480156104bd57600080fd5b506101df6117e0565b3480156104d257600080fd5b5061021461183d565b3480156104e757600080fd5b506104f0611846565b6040516102219190613c99565b34801561050957600080fd5b5061026c61184f565b3360009081526008602052604090205460ff1660011461054d5760405162461bcd60e51b81526004016105449061358d565b60405180910390fd5b845160401461056e5760405162461bcd60e51b815260040161054490613a68565b60008451118061057f575060008351115b61059b5760405162461bcd60e51b8152600401610544906139bb565b600b856040516105ab919061301b565b9081526040519081900360200190205460ff16156105db5760405162461bcd60e51b81526004016105449061346d565b6105e58484611855565b6000858584866002604051602001610601959493929190613247565b60408051601f1981840301815291815281516020928301206000818152600a90935291205490915060ff16156106495760405162461bcd60e51b815260040161054490613399565b6106538183611a98565b61066f5760405162461bcd60e51b81526004016105449061398c565b61067884611adb565b61068185611cd6565b60095461068d90611dbc565b6005805460ff191660ff929092169190911790556106ad86826001611e12565b7fac9b82db4e104d515319a481096bfd91a4f40ee10837d5a2c8d51b9a03dc48ae866040516106dc9190613152565b60405180910390a1505050505050565b6000338361070c5760405162461bcd60e51b8152600401610544906136cc565b6001600160a01b0383161561093a5734156107395760405162461bcd60e51b815260040161054490613521565b61074b836001600160a01b0316611e5f565b6107675760405162461bcd60e51b815260040161054490613c4b565b604051636eb1769f60e11b815283906000906001600160a01b0383169063dd62ed3e9061079a9086903090600401613064565b60206040518083038186803b1580156107b257600080fd5b505afa1580156107c6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107ea9190612f94565b90508581101561080c5760405162461bcd60e51b81526004016105449061383b565b6040516370a0823160e01b81526000906001600160a01b038416906370a082319061083b908790600401613037565b60206040518083038186803b15801561085357600080fd5b505afa158015610867573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061088b9190612f94565b9050868110156108ad5760405162461bcd60e51b815260040161054490613955565b6108c26001600160a01b03841685308a611e65565b6108cb8661113f565b1561093257604051630852cd8d60e31b815286906001600160a01b038216906342966c68906108fe908b90600401613c90565b600060405180830381600087803b15801561091857600080fd5b505af115801561092c573d6000803e3d6000fd5b50505050505b505050610959565b8334146109595760405162461bcd60e51b815260040161054490613a03565b7f5ddf9724d8fe5d9e12499be2867f93d41a582733dcd65f74a486ad7e306671468186868660405161098e94939291906130c6565b60405180910390a160019150505b9392505050565b630a85bd0160e11b949350505050565b60045481565b3360009081526008602052604090205460ff166001146109eb5760405162461bcd60e51b81526004016105449061358d565b8651604014610a0c5760405162461bcd60e51b815260040161054490613a68565b6001600160a01b038516610a325760405162461bcd60e51b815260040161054490613353565b600b87604051610a42919061301b565b9081526040519081900360200190205460ff1615610a725760405162461bcd60e51b81526004016105449061346d565b610a7d868686611ebd565b60008787878787876002604051602001610a9d97969594939291906131ac565b60408051601f1981840301815291815281516020928301206000818152600a90935291205490915060ff1615610ae55760405162461bcd60e51b815260040161054490613399565b610aef8183611a98565b610b0b5760405162461bcd60e51b81526004016105449061398c565b610b4d87878787878080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250611fff92505050565b610b5988826001611e12565b7f8ed8b1f0dd3babfdf1477ba2b27a5b0d2f1c9148448fd22cf2c75e658293c7b188604051610b889190613152565b60405180910390a15050505050505050565b60055461010090046001600160a01b03163314610bc95760405162461bcd60e51b815260040161054490613b01565b60005460ff16610beb5760405162461bcd60e51b815260040161054490613769565b60005461010090046001600160a01b0316610c185760405162461bcd60e51b81526004016105449061387d565b600080546040516001600160a01b0361010090920491909116914780156108fc02929091818181858888f19350505050158015610c59573d6000803e3d6000fd5b50565b6001600160a01b0381166000908152600d602052604090205460ff1615155b919050565b60035481565b60005461010090046001600160a01b031681565b60055461010090046001600160a01b03163314610cc95760405162461bcd60e51b815260040161054490613b01565b306001600160a01b0382161415610cf25760405162461bcd60e51b81526004016105449061369c565b610d04816001600160a01b0316611e5f565b610d205760405162461bcd60e51b815260040161054490613c4b565b610d298161113f565b15610d465760405162461bcd60e51b8152600401610544906132df565b6001600160a01b03166000908152600c60205260409020805460ff19166001179055565b3360009081526008602052604090205460ff16600114610d9c5760405162461bcd60e51b81526004016105449061358d565b8251604014610dbd5760405162461bcd60e51b815260040161054490613a68565b600b83604051610dcd919061301b565b9081526040519081900360200190205460ff1615610dfd5760405162461bcd60e51b81526004016105449061346d565b60005460ff1615610e205760405162461bcd60e51b815260040161054490613a3a565b610e32826001600160a01b0316611e5f565b610e4e5760405162461bcd60e51b815260040161054490613c4b565b600083836002604051602001610e6693929190613219565b60408051601f1981840301815291815281516020928301206000818152600a90935291205490915060ff1615610eae5760405162461bcd60e51b815260040161054490613399565b610eb88183611a98565b610ed45760405162461bcd60e51b81526004016105449061398c565b60008054600160ff199091168117610100600160a81b0319166101006001600160a01b0387160217909155610f0c9085908390611e12565b7f5e06c4b22547d430736ce834764dbfee08f1c4cf7ae3d53178aa56effa593ed084604051610f3b9190613152565b60405180910390a150505050565b60055461010090046001600160a01b03163314610f785760405162461bcd60e51b815260040161054490613b01565b60005460ff16610f9a5760405162461bcd60e51b815260040161054490613769565b60005461010090046001600160a01b0316610fc75760405162461bcd60e51b81526004016105449061387d565b306001600160a01b0382161415610ff05760405162461bcd60e51b81526004016105449061369c565b611002816001600160a01b0316611e5f565b61101e5760405162461bcd60e51b815260040161054490613c4b565b6040516370a0823160e01b815281906000906001600160a01b038316906370a082319061104f903090600401613037565b60206040518083038186803b15801561106757600080fd5b505afa15801561107b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061109f9190612f94565b90506000546110c0906001600160a01b0384811691610100900416836120e3565b6110c98361113f565b1561113a576000546040516301fc6bd160e21b815284916001600160a01b03808416926307f1af4492611106926101009091041690600401613037565b600060405180830381600087803b15801561112057600080fd5b505af1158015611134573d6000803e3d6000fd5b50505050505b505050565b6001600160a01b03166000908152600c602052604090205460ff16151590565b6001600160a01b031660009081526008602052604090205460ff1660011490565b60055461010090046001600160a01b031633146111af5760405162461bcd60e51b815260040161054490613b01565b306001600160a01b03821614156111d85760405162461bcd60e51b81526004016105449061369c565b6111ea816001600160a01b0316611e5f565b6112065760405162461bcd60e51b815260040161054490613c4b565b61120f81610c5c565b1561122c5760405162461bcd60e51b8152600401610544906132df565b6001600160a01b03166000908152600d60205260409020805460ff19166001179055565b60055461010090046001600160a01b031681565b606060098054806020026020016040519081016040528092919081815260200182805480156112bc57602002820191906000526020600020905b81546001600160a01b0316815260019091019060200180831161129e575b5050505050905090565b60055461010090046001600160a01b031633146112f55760405162461bcd60e51b815260040161054490613b01565b6112fe8161113f565b61131a5760405162461bcd60e51b81526004016105449061391e565b6001600160a01b03166000908152600c60205260409020805460ff19169055565b600080600b8360405161134e919061301b565b9081526040519081900360200190205460ff16119050919050565b3360009081526008602052604090205460ff1660011461139b5760405162461bcd60e51b81526004016105449061358d565b85516040146113bc5760405162461bcd60e51b815260040161054490613a68565b6001600160a01b0385166113e25760405162461bcd60e51b815260040161054490613353565b600084116114025760405162461bcd60e51b8152600401610544906137f3565b600b86604051611412919061301b565b9081526040519081900360200190205460ff16156114425760405162461bcd60e51b81526004016105449061346d565b821561145857611453828686612102565b611478565b834710156114785760405162461bcd60e51b8152600401610544906135c4565b60008686868686600260405160200161149696959493929190613165565b60408051601f1981840301815291815281516020928301206000818152600a90935291205490915060ff16156114de5760405162461bcd60e51b815260040161054490613399565b6114e88183611a98565b6115045760405162461bcd60e51b81526004016105449061398c565b831561151a57611515838787612235565b6115ab565b8447101561153a5760405162461bcd60e51b8152600401610544906135c4565b6040516001600160a01b0387169086156108fc029087906000818181858888f19350505050158015611570573d6000803e3d6000fd5b507fc95f8b91b103304386b955ef73fadac189f8ad66b33369b6c34a17a60db7bd0a86866040516115a292919061304b565b60405180910390a15b6115b787826001611e12565b7f8ed8b1f0dd3babfdf1477ba2b27a5b0d2f1c9148448fd22cf2c75e658293c7b1876040516115e69190613152565b60405180910390a150505050505050565b60025481565b60055461010090046001600160a01b0316331461162c5760405162461bcd60e51b815260040161054490613b01565b61163581610c5c565b6116515760405162461bcd60e51b81526004016105449061391e565b6001600160a01b03166000908152600d60205260409020805460ff19169055565b6000336116876001600160a01b038516611e5f565b6116a35760405162461bcd60e51b815260040161054490613c4b565b600083116116c35760405162461bcd60e51b8152600401610544906136cc565b604051632142170760e11b815284906001600160a01b038216906342842e0e906116f59085903090899060040161307e565b600060405180830381600087803b15801561170f57600080fd5b505af1158015611723573d6000803e3d6000fd5b5050505061173085610c5c565b1561179757604051630852cd8d60e31b815285906001600160a01b038216906342966c6890611763908890600401613c90565b600060405180830381600087803b15801561177d57600080fd5b505af1158015611791573d6000803e3d6000fd5b50505050505b7f06caac0500ba0f6fbe6852539c648ab228f15265797d52ca06096c9e74bb17df828786886040516117cc94939291906130c6565b60405180910390a150600195945050505050565b60055461010090046001600160a01b0316331461180f5760405162461bcd60e51b815260040161054490613b01565b60005460ff166118315760405162461bcd60e51b815260040161054490613769565b6000805460ff19169055565b60005460ff1681565b60055460ff1681565b60015481565b815160005b8181101561190d57600084828151811061188457634e487b7160e01b600052603260045260246000fd5b6020026020010151905060006001600160a01b0316816001600160a01b031614156118c15760405162461bcd60e51b815260040161054490613658565b6001600160a01b03811660009081526008602052604090205460ff16156118fa5760405162461bcd60e51b8152600401610544906138c0565b508061190581613d90565b91505061185a565b5061191783612362565b6119335760405162461bcd60e51b8152600401610544906133c5565b60055461194e9061010090046001600160a01b03168461247a565b61196a5760405162461bcd60e51b815260040161054490613c05565b61197382612362565b61198f5760405162461bcd60e51b815260040161054490613b38565b815160005b81811015611a555760008482815181106119be57634e487b7160e01b600052603260045260246000fd5b6020908102919091018101516001600160a01b0381166000908152600690925260409091205490915060ff1615611a075760405162461bcd60e51b815260040161054490613b84565b6001600160a01b03811660009081526008602052604090205460ff16600114611a425760405162461bcd60e51b815260040161054490613789565b5080611a4d81613d90565b915050611994565b5060015483518551600954611a6a9190613cd1565b611a749190613d4d565b1115611a925760405162461bcd60e51b815260040161054490613722565b50505050565b60006103cf82511115611abd5760405162461bcd60e51b815260040161054490613621565b6000611ac98484612515565b60055460ff1611159150505b92915050565b8051611ae657610c59565b60005b8151811015611b535760086000838381518110611b1657634e487b7160e01b600052603260045260246000fd5b6020908102919091018101516001600160a01b03168252810191909152604001600020805460ff1916905580611b4b81613d90565b915050611ae9565b5060005b600954811015611bff576008600060098381548110611b8657634e487b7160e01b600052603260045260246000fd5b60009182526020808320909101546001600160a01b0316835282019290925260400190205460ff16611bed5760098181548110611bd357634e487b7160e01b600052603260045260246000fd5b600091825260209091200180546001600160a01b03191690555b80611bf781613d90565b915050611b57565b50601060005b60095481101561113a57600060098281548110611c3257634e487b7160e01b600052603260045260246000fd5b6000918252602090912001546001600160a01b0316905080611c61578260101415611c5b578192505b50611cc4565b82601014611cc2578060098481548110611c8b57634e487b7160e01b600052603260045260246000fd5b600091825260209091200180546001600160a01b0319166001600160a01b039290921691909117905582611cbe81613d90565b9350505b505b80611cce81613d90565b915050611c05565b8051611ce157610c59565b60005b8151811015611db8576000828281518110611d0f57634e487b7160e01b600052603260045260246000fd5b6020908102919091018101516001600160a01b0381166000908152600890925260409091205490915060ff16611da5576001600160a01b0381166000818152600860205260408120805460ff191660019081179091556009805491820181559091527f6e1540171b6c0c960b71a7020d9f60077f6af931a8bbf590da0223dacf75c7af0180546001600160a01b03191690911790555b5080611db081613d90565b915050611ce4565b5050565b6000808211611ddd5760405162461bcd60e51b815260040161054490613411565b60006001606484600354611df19190613d2e565b611dfb9190613cd1565b611e059190613d4d565b905061099c606482613d0e565b80600b84604051611e23919061301b565b9081526040805160209281900383019020805460ff1990811660ff958616179091556000958652600a9092529093208054909316911617905550565b3b151590565b611a92846323b872dd60e01b858585604051602401611e869392919061307e565b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b0319909316929092179091526126c1565b6001600160a01b038216611ee35760405162461bcd60e51b81526004016105449061329c565b306001600160a01b0384161415611f0c5760405162461bcd60e51b81526004016105449061369c565b611f1e836001600160a01b0316611e5f565b611f3a5760405162461bcd60e51b815260040161054490613c4b565b611f4383610c5c565b15611f4d5761113a565b6040516331a9108f60e11b815283906000906001600160a01b03831690636352211e90611f7e908690600401613c90565b60206040518083038186803b158015611f9657600080fd5b505afa158015611faa573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611fce9190612b48565b90506001600160a01b0381163014611ff85760405162461bcd60e51b815260040161054490613327565b5050505050565b61200884610c5c565b156120755760405163d0def52160e01b815284906001600160a01b0382169063d0def5219061203d90879086906004016130a2565b600060405180830381600087803b15801561205757600080fd5b505af115801561206b573d6000803e3d6000fd5b5050505050611a92565b604051632142170760e11b8152309085906001600160a01b038216906342842e0e906120a99085908990899060040161307e565b600060405180830381600087803b1580156120c357600080fd5b505af11580156120d7573d6000803e3d6000fd5b50505050505050505050565b61113a8363a9059cbb60e01b8484604051602401611e8692919061304b565b6001600160a01b0382166121285760405162461bcd60e51b81526004016105449061329c565b306001600160a01b03841614156121515760405162461bcd60e51b81526004016105449061369c565b612163836001600160a01b0316611e5f565b61217f5760405162461bcd60e51b815260040161054490613c4b565b6121888361113f565b156121925761113a565b6040516370a0823160e01b815283906000906001600160a01b038316906370a08231906121c3903090600401613037565b60206040518083038186803b1580156121db57600080fd5b505afa1580156121ef573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906122139190612f94565b905082811015611ff85760405162461bcd60e51b8152600401610544906134a4565b61223e8361113f565b156122ab576040516340c10f1960e01b815283906001600160a01b038216906340c10f1990612273908690869060040161304b565b600060405180830381600087803b15801561228d57600080fd5b505af11580156122a1573d6000803e3d6000fd5b505050505061113a565b6040516370a0823160e01b815283906000906001600160a01b038316906370a08231906122dc903090600401613037565b60206040518083038186803b1580156122f457600080fd5b505afa158015612308573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061232c9190612f94565b90508281101561234e5760405162461bcd60e51b8152600401610544906134a4565b611ff86001600160a01b03831685856120e3565b6000805b825181101561247157600083828151811061239157634e487b7160e01b600052603260045260246000fd5b6020026020010151905060006001600160a01b0316816001600160a01b031614156123bc5750612471565b60006123c9836001613cd1565b90505b845181101561245c5760008582815181106123f757634e487b7160e01b600052603260045260246000fd5b6020026020010151905060006001600160a01b0316816001600160a01b03161415612422575061245c565b806001600160a01b0316836001600160a01b03161415612449576000945050505050610c7b565b508061245481613d90565b9150506123cc565b5050808061246990613d90565b915050612366565b50600192915050565b60008060005b835181101561250a578381815181106124a957634e487b7160e01b600052603260045260246000fd5b6020026020010151915060006001600160a01b0316826001600160a01b031614156124d35761250a565b846001600160a01b0316826001600160a01b031614156124f857600092505050611ad5565b8061250281613d90565b915050612480565b506001949350505050565b60045481516000918291829161252b9190612750565b905060008167ffffffffffffffff81111561255657634e487b7160e01b600052604160045260246000fd5b60405190808252806020026020018201604052801561257f578160200160208202803683370190505b50905060008060005b848110156126855760006125a9846004548b61275c9092919063ffffffff16565b905060006125b78b83612824565b90506001600160a01b0381166125df5760405162461bcd60e51b8152600401610544906136f8565b6001600160a01b03811660009081526008602052604090205460ff1660011415612661578761260d81613d90565b9850508086858061261d90613dab565b965060ff168151811061264057634e487b7160e01b600052603260045260246000fd5b60200260200101906001600160a01b031690816001600160a01b0316815250505b60045461266e9086613cd1565b94505050808061267d90613d90565b915050612588565b50600061269184612362565b905060609350806126b45760405162461bcd60e51b81526004016105449061343f565b5093979650505050505050565b6000612716826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166129269092919063ffffffff16565b80519091501561113a57808060200190518101906127349190612bce565b61113a5760405162461bcd60e51b815260040161054490613bbb565b600061099c8284613d0e565b60608161276a81601f613cd1565b10156127885760405162461bcd60e51b815260040161054490613565565b6127928284613cd1565b845110156127b25760405162461bcd60e51b815260040161054490613ad6565b6060821580156127d1576040519150600082526020820160405261281b565b6040519150601f8416801560200281840101858101878315602002848b0101015b8183101561280a5780518352602092830192016127f2565b5050858452601f01601f1916604052505b50949350505050565b6000806000806004548551146128405760009350505050611ad5565b50505060208201516040830151606084015160001a7f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08211156128895760009350505050611ad5565b601b8160ff1610156128a3576128a0601b82613ce9565b90505b8060ff16601b141580156128bb57508060ff16601c14155b156128cc5760009350505050611ad5565b600186828585604051600081526020016040526040516128ef949392919061311f565b6020604051602081039080840390855afa158015612911573d6000803e3d6000fd5b5050604051601f190151979650505050505050565b6060612935848460008561293d565b949350505050565b60608247101561295f5760405162461bcd60e51b8152600401610544906134db565b61296885611e5f565b6129845760405162461bcd60e51b815260040161054490613a9f565b600080866001600160a01b031685876040516129a0919061301b565b60006040518083038185875af1925050503d80600081146129dd576040519150601f19603f3d011682016040523d82523d6000602084013e6129e2565b606091505b50915091506129f28282866129fd565b979650505050505050565b60608315612a0c57508161099c565b825115612a1c5782518084602001fd5b8160405162461bcd60e51b81526004016105449190613152565b600082601f830112612a46578081fd5b8135602067ffffffffffffffff821115612a6257612a62613de1565b808202612a70828201613ca7565b838152828101908684018388018501891015612a8a578687fd5b8693505b85841015612ab5578035612aa181613df7565b835260019390930192918401918401612a8e565b50979650505050505050565b600082601f830112612ad1578081fd5b813567ffffffffffffffff811115612aeb57612aeb613de1565b612afe601f8201601f1916602001613ca7565b818152846020838601011115612b12578283fd5b816020850160208301379081016020019190915292915050565b600060208284031215612b3d578081fd5b813561099c81613df7565b600060208284031215612b59578081fd5b815161099c81613df7565b60008060008060808587031215612b79578283fd5b8435612b8481613df7565b93506020850135612b9481613df7565b925060408501359150606085013567ffffffffffffffff811115612bb6578182fd5b612bc287828801612ac1565b91505092959194509250565b600060208284031215612bdf578081fd5b815161099c81613e0c565b600060208284031215612bfb578081fd5b813567ffffffffffffffff811115612c11578182fd5b61293584828501612ac1565b60008060008060008060c08789031215612c35578182fd5b863567ffffffffffffffff80821115612c4c578384fd5b612c588a838b01612ac1565b975060208901359150612c6a82613df7565b90955060408801359450606088013590612c8382613e0c565b909350608088013590612c9582613df7565b90925060a08801359080821115612caa578283fd5b50612cb789828a01612ac1565b9150509295509295509295565b600080600080600080600060c0888a031215612cde578081fd5b873567ffffffffffffffff80821115612cf5578283fd5b612d018b838c01612ac1565b985060208a01359150612d1382613df7565b909650604089013590612d2582613df7565b9095506060890135945060808901359080821115612d41578283fd5b818a0191508a601f830112612d54578283fd5b813581811115612d62578384fd5b8b6020828501011115612d73578384fd5b6020830195508094505060a08a0135915080821115612d90578283fd5b50612d9d8a828b01612ac1565b91505092959891949750929550565b600080600060608486031215612dc0578283fd5b833567ffffffffffffffff80821115612dd7578485fd5b612de387838801612ac1565b945060208601359150612df582613df7565b90925060408501359080821115612e0a578283fd5b50612e1786828701612ac1565b9150509250925092565b600080600060608486031215612e35578081fd5b833567ffffffffffffffff811115612e4b578182fd5b612e5786828701612ac1565b9350506020840135612e6881613df7565b929592945050506040919091013590565b600080600080600060a08688031215612e90578283fd5b853567ffffffffffffffff80821115612ea7578485fd5b612eb389838a01612ac1565b96506020880135915080821115612ec8578485fd5b612ed489838a01612a36565b95506040880135915080821115612ee9578485fd5b612ef589838a01612a36565b94506060880135915060ff82168214612f0c578283fd5b90925060808701359080821115612f21578283fd5b50612f2e88828901612ac1565b9150509295509295909350565b600080600060608486031215612f4f578081fd5b833567ffffffffffffffff811115612f65578182fd5b612f7186828701612ac1565b935050602084013591506040840135612f8981613df7565b809150509250925092565b600060208284031215612fa5578081fd5b5051919050565b6000815180845260208085019450808401835b83811015612fe45781516001600160a01b031687529582019590820190600101612fbf565b509495945050505050565b60008151808452613007816020860160208601613d64565b601f01601f19169290920160200192915050565b6000825161302d818460208701613d64565b9190910192915050565b6001600160a01b0391909116815260200190565b6001600160a01b03929092168252602082015260400190565b6001600160a01b0392831681529116602082015260400190565b6001600160a01b039384168152919092166020820152604081019190915260600190565b6001600160a01b038316815260406020820181905260009061293590830184612fef565b600060018060a01b038087168352608060208401526130e86080840187612fef565b6040840195909552929092166060909101525092915050565b60006020825261099c6020830184612fac565b901515815260200190565b93845260ff9290921660208401526040830152606082015260800190565b6001600160e01b031991909116815260200190565b60006020825261099c6020830184612fef565b600060c0825261317860c0830189612fef565b6001600160a01b03978816602084015260408301969096525092151560608401529316608082015260a00191909152919050565b600060c082526131bf60c083018a612fef565b60018060a01b03808a16602085015280891660408501525086606084015282810360808401528481528486602083013781602086830101526020601f19601f8701168201019150508260a083015298975050505050505050565b60006060825261322c6060830186612fef565b6001600160a01b039490941660208301525060400152919050565b600060a0825261325a60a0830188612fef565b828103602084015261326c8188612fac565b905060ff8616604084015282810360608401526132898186612fac565b9150508260808301529695505050505050565b60208082526023908201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260408201526265737360e81b606082015260800190565b60208082526028908201527f5468697320616464726573732068617320616c7265616479206265656e20726560408201526719da5cdd195c995960c21b606082015260800190565b6020808252601290820152712737ba1037bbb732b91037b3103a37b5b2b760711b604082015260600190565b60208082526026908201527f57697468647261773a207472616e7366657220746f20746865207a65726f206160408201526564647265737360d01b606082015260800190565b602080825260129082015271496e76616c6964207369676e61747572657360701b604082015260600190565b6020808252602c908201527f4475706c696361746520706172616d657465727320666f72207468652061646460408201526b3932b9b9903a37903537b4b760a11b606082015260800190565b60208082526014908201527326b0b730b3b2b91021b0b713ba1032b6b83a3c9760611b604082015260600190565b6020808252601490820152735369676e617475726573206475706c696361746560601b604082015260600190565b6020808252601e908201527f5472616e73616374696f6e20686173206265656e20636f6d706c657465640000604082015260600190565b6020808252601a908201527f4e6f20656e6f7567682062616c616e6365206f6620746f6b656e000000000000604082015260600190565b60208082526026908201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6040820152651c8818d85b1b60d21b606082015260800190565b60208082526024908201527f45524332303a20446f6573206e6f742061636365707420457468657265756d2060408201526321b7b4b760e11b606082015260800190565b6020808252600e908201526d736c6963655f6f766572666c6f7760901b604082015260600190565b6020808252601b908201527f4f6e6c79206d616e616765722063616e20657865637574652069740000000000604082015260600190565b6020808252603f908201527f5468697320636f6e7472616374206164647265737320646f6573206e6f74206860408201527f6176652073756666696369656e742062616c616e6365206f6620657468657200606082015260800190565b6020808252601d908201527f4d6178206c656e677468206f66207369676e6174757265733a20393735000000604082015260600190565b60208082526024908201527f4552524f523a204465746563746564207a65726f206164647265737320696e206040820152636164647360e01b606082015260800190565b6020808252601690820152752237903737ba3434b73390313c903cb7bab939b2b63360511b604082015260600190565b60208082526012908201527111549493d48e8816995c9bc8185b5bdd5b9d60721b604082015260600190565b60208082526010908201526f29b4b3b730ba3ab932b99032b93937b960811b604082015260600190565b60208082526027908201527f457863656564656420746865206d6178696d756d206e756d626572206f66206d604082015266616e616765727360c81b606082015260800190565b60208082526006908201526511195b9a595960d21b604082015260600190565b60208082526044908201527f5468657265206172652061646472657373657320696e2074686520657869746960408201527f6e672061646472657373206c697374207468617420617265206e6f74206d616e60608201526330b3b2b960e11b608082015260a00190565b60208082526028908201527f5769746864726177616c20616d6f756e74206d75737420626520677265617465604082015267072207468616e20360c41b606082015260800190565b60208082526022908201527f4e6f20656e6f75676820616d6f756e7420666f7220617574686f72697a61746960408201526137b760f11b606082015260800190565b60208082526023908201527f4552524f523a207472616e7366657220746f20746865207a65726f206164647260408201526265737360e81b606082015260800190565b602080825260409082018190527f5468652061646472657373206c6973742074686174206973206265696e672061908201527f6464656420616c7265616479206578697374732061732061206d616e61676572606082015260800190565b6020808252601e908201527f546869732061646472657373206973206e6f7420726567697374657265640000604082015260600190565b6020808252601e908201527f4e6f20656e6f7567682062616c616e6365206f662074686520746f6b656e0000604082015260600190565b60208082526015908201527415985b1a59081cda59db985d1d5c995cc819985a5b605a1b604082015260600190565b60208082526028908201527f546865726520617265206e6f206d616e6167657273206a6f696e696e67206f726040820152672065786974696e6760c01b606082015260800190565b6020808252601d908201527f496e636f6e73697374656e637920457468657265756d20616d6f756e74000000604082015260600190565b602080825260149082015273125d081a185cc81899595b881d5c19dc9859195960621b604082015260600190565b60208082526019908201527f4669786564206c656e677468206f662074784b65793a20363400000000000000604082015260600190565b6020808252601d908201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604082015260600190565b602080825260119082015270736c6963655f6f75744f66426f756e647360781b604082015260600190565b60208082526019908201527f4f6e6c79206f776e65722063616e206578656375746520697400000000000000604082015260600190565b6020808252602c908201527f4475706c696361746520706172616d657465727320666f72207468652061646460408201526b1c995cdcc81d1bc8195e1a5d60a21b606082015260800190565b60208082526017908201527f43616e277420657869742073656564206d616e61676572000000000000000000604082015260600190565b6020808252602a908201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6040820152691bdd081cdd58d8d9595960b21b606082015260800190565b60208082526026908201527f436f6e74726163742063726561746f722063616e6e6f7420616374206173206d60408201526530b730b3b2b960d11b606082015260800190565b60208082526025908201527f5468652061646472657373206973206e6f74206120636f6e7472616374206164604082015264647265737360d81b606082015260800190565b90815260200190565b60ff91909116815260200190565b60405181810167ffffffffffffffff81118282101715613cc957613cc9613de1565b604052919050565b60008219821115613ce457613ce4613dcb565b500190565b600060ff821660ff84168060ff03821115613d0657613d06613dcb565b019392505050565b600082613d2957634e487b7160e01b81526012600452602481fd5b500490565b6000816000190483118215151615613d4857613d48613dcb565b500290565b600082821015613d5f57613d5f613dcb565b500390565b60005b83811015613d7f578181015183820152602001613d67565b83811115611a925750506000910152565b6000600019821415613da457613da4613dcb565b5060010190565b600060ff821660ff811415613dc257613dc2613dcb565b60010192915050565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052604160045260246000fd5b6001600160a01b0381168114610c5957600080fd5b8015158114610c5957600080fdfea264697066735822122019c3ee65ea169e892d8d72506341418ed96c5e7077b17586068df7271de9fefc64736f6c63430008000033"

// DeployDynamicMultiSig deploys a new Ethereum contract, binding an instance of DynamicMultiSig to it.
func DeployDynamicMultiSig(auth *bind.TransactOpts, backend bind.ContractBackend, _managers []common.Address) (common.Address, *types.Transaction, *DynamicMultiSig, error) {
	parsed, err := abi.JSON(strings.NewReader(DynamicMultiSigABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DynamicMultiSigBin), backend, _managers)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DynamicMultiSig{DynamicMultiSigCaller: DynamicMultiSigCaller{contract: contract}, DynamicMultiSigTransactor: DynamicMultiSigTransactor{contract: contract}, DynamicMultiSigFilterer: DynamicMultiSigFilterer{contract: contract}}, nil
}

// DynamicMultiSig is an auto generated Go binding around an Ethereum contract.
type DynamicMultiSig struct {
	DynamicMultiSigCaller     // Read-only binding to the contract
	DynamicMultiSigTransactor // Write-only binding to the contract
	DynamicMultiSigFilterer   // Log filterer for contract events
}

// DynamicMultiSigCaller is an auto generated read-only Go binding around an Ethereum contract.
type DynamicMultiSigCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DynamicMultiSigTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DynamicMultiSigTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DynamicMultiSigFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DynamicMultiSigFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DynamicMultiSigSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DynamicMultiSigSession struct {
	Contract     *DynamicMultiSig  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DynamicMultiSigCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DynamicMultiSigCallerSession struct {
	Contract *DynamicMultiSigCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// DynamicMultiSigTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DynamicMultiSigTransactorSession struct {
	Contract     *DynamicMultiSigTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// DynamicMultiSigRaw is an auto generated low-level Go binding around an Ethereum contract.
type DynamicMultiSigRaw struct {
	Contract *DynamicMultiSig // Generic contract binding to access the raw methods on
}

// DynamicMultiSigCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DynamicMultiSigCallerRaw struct {
	Contract *DynamicMultiSigCaller // Generic read-only contract binding to access the raw methods on
}

// DynamicMultiSigTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DynamicMultiSigTransactorRaw struct {
	Contract *DynamicMultiSigTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDynamicMultiSig creates a new instance of DynamicMultiSig, bound to a specific deployed contract.
func NewDynamicMultiSig(address common.Address, backend bind.ContractBackend) (*DynamicMultiSig, error) {
	contract, err := bindDynamicMultiSig(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DynamicMultiSig{DynamicMultiSigCaller: DynamicMultiSigCaller{contract: contract}, DynamicMultiSigTransactor: DynamicMultiSigTransactor{contract: contract}, DynamicMultiSigFilterer: DynamicMultiSigFilterer{contract: contract}}, nil
}

// NewDynamicMultiSigCaller creates a new read-only instance of DynamicMultiSig, bound to a specific deployed contract.
func NewDynamicMultiSigCaller(address common.Address, caller bind.ContractCaller) (*DynamicMultiSigCaller, error) {
	contract, err := bindDynamicMultiSig(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DynamicMultiSigCaller{contract: contract}, nil
}

// NewDynamicMultiSigTransactor creates a new write-only instance of DynamicMultiSig, bound to a specific deployed contract.
func NewDynamicMultiSigTransactor(address common.Address, transactor bind.ContractTransactor) (*DynamicMultiSigTransactor, error) {
	contract, err := bindDynamicMultiSig(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DynamicMultiSigTransactor{contract: contract}, nil
}

// NewDynamicMultiSigFilterer creates a new log filterer instance of DynamicMultiSig, bound to a specific deployed contract.
func NewDynamicMultiSigFilterer(address common.Address, filterer bind.ContractFilterer) (*DynamicMultiSigFilterer, error) {
	contract, err := bindDynamicMultiSig(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DynamicMultiSigFilterer{contract: contract}, nil
}

// bindDynamicMultiSig binds a generic wrapper to an already deployed contract.
func bindDynamicMultiSig(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DynamicMultiSigABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DynamicMultiSig *DynamicMultiSigRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DynamicMultiSig.Contract.DynamicMultiSigCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DynamicMultiSig *DynamicMultiSigRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DynamicMultiSig.Contract.DynamicMultiSigTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DynamicMultiSig *DynamicMultiSigRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DynamicMultiSig.Contract.DynamicMultiSigTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DynamicMultiSig *DynamicMultiSigCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DynamicMultiSig.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DynamicMultiSig *DynamicMultiSigTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DynamicMultiSig.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DynamicMultiSig *DynamicMultiSigTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DynamicMultiSig.Contract.contract.Transact(opts, method, params...)
}

// AllManagers is a free data retrieval call binding the contract method 0x9c30b35e.
//
// Solidity: function allManagers() view returns(address[])
func (_DynamicMultiSig *DynamicMultiSigCaller) AllManagers(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _DynamicMultiSig.contract.Call(opts, &out, "allManagers")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// AllManagers is a free data retrieval call binding the contract method 0x9c30b35e.
//
// Solidity: function allManagers() view returns(address[])
func (_DynamicMultiSig *DynamicMultiSigSession) AllManagers() ([]common.Address, error) {
	return _DynamicMultiSig.Contract.AllManagers(&_DynamicMultiSig.CallOpts)
}

// AllManagers is a free data retrieval call binding the contract method 0x9c30b35e.
//
// Solidity: function allManagers() view returns(address[])
func (_DynamicMultiSig *DynamicMultiSigCallerSession) AllManagers() ([]common.Address, error) {
	return _DynamicMultiSig.Contract.AllManagers(&_DynamicMultiSig.CallOpts)
}

// CurrentMinSignatures is a free data retrieval call binding the contract method 0xe079cee9.
//
// Solidity: function current_min_signatures() view returns(uint8)
func (_DynamicMultiSig *DynamicMultiSigCaller) CurrentMinSignatures(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _DynamicMultiSig.contract.Call(opts, &out, "current_min_signatures")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// CurrentMinSignatures is a free data retrieval call binding the contract method 0xe079cee9.
//
// Solidity: function current_min_signatures() view returns(uint8)
func (_DynamicMultiSig *DynamicMultiSigSession) CurrentMinSignatures() (uint8, error) {
	return _DynamicMultiSig.Contract.CurrentMinSignatures(&_DynamicMultiSig.CallOpts)
}

// CurrentMinSignatures is a free data retrieval call binding the contract method 0xe079cee9.
//
// Solidity: function current_min_signatures() view returns(uint8)
func (_DynamicMultiSig *DynamicMultiSigCallerSession) CurrentMinSignatures() (uint8, error) {
	return _DynamicMultiSig.Contract.CurrentMinSignatures(&_DynamicMultiSig.CallOpts)
}

// IfManager is a free data retrieval call binding the contract method 0x75173b70.
//
// Solidity: function ifManager(address _manager) view returns(bool)
func (_DynamicMultiSig *DynamicMultiSigCaller) IfManager(opts *bind.CallOpts, _manager common.Address) (bool, error) {
	var out []interface{}
	err := _DynamicMultiSig.contract.Call(opts, &out, "ifManager", _manager)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IfManager is a free data retrieval call binding the contract method 0x75173b70.
//
// Solidity: function ifManager(address _manager) view returns(bool)
func (_DynamicMultiSig *DynamicMultiSigSession) IfManager(_manager common.Address) (bool, error) {
	return _DynamicMultiSig.Contract.IfManager(&_DynamicMultiSig.CallOpts, _manager)
}

// IfManager is a free data retrieval call binding the contract method 0x75173b70.
//
// Solidity: function ifManager(address _manager) view returns(bool)
func (_DynamicMultiSig *DynamicMultiSigCallerSession) IfManager(_manager common.Address) (bool, error) {
	return _DynamicMultiSig.Contract.IfManager(&_DynamicMultiSig.CallOpts, _manager)
}

// IsCompletedTx is a free data retrieval call binding the contract method 0xa5e399b3.
//
// Solidity: function isCompletedTx(string txKey) view returns(bool)
func (_DynamicMultiSig *DynamicMultiSigCaller) IsCompletedTx(opts *bind.CallOpts, txKey string) (bool, error) {
	var out []interface{}
	err := _DynamicMultiSig.contract.Call(opts, &out, "isCompletedTx", txKey)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCompletedTx is a free data retrieval call binding the contract method 0xa5e399b3.
//
// Solidity: function isCompletedTx(string txKey) view returns(bool)
func (_DynamicMultiSig *DynamicMultiSigSession) IsCompletedTx(txKey string) (bool, error) {
	return _DynamicMultiSig.Contract.IsCompletedTx(&_DynamicMultiSig.CallOpts, txKey)
}

// IsCompletedTx is a free data retrieval call binding the contract method 0xa5e399b3.
//
// Solidity: function isCompletedTx(string txKey) view returns(bool)
func (_DynamicMultiSig *DynamicMultiSigCallerSession) IsCompletedTx(txKey string) (bool, error) {
	return _DynamicMultiSig.Contract.IsCompletedTx(&_DynamicMultiSig.CallOpts, txKey)
}

// IsMinterERC20 is a free data retrieval call binding the contract method 0x6a7142e1.
//
// Solidity: function isMinterERC20(address ERC20) view returns(bool)
func (_DynamicMultiSig *DynamicMultiSigCaller) IsMinterERC20(opts *bind.CallOpts, ERC20 common.Address) (bool, error) {
	var out []interface{}
	err := _DynamicMultiSig.contract.Call(opts, &out, "isMinterERC20", ERC20)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMinterERC20 is a free data retrieval call binding the contract method 0x6a7142e1.
//
// Solidity: function isMinterERC20(address ERC20) view returns(bool)
func (_DynamicMultiSig *DynamicMultiSigSession) IsMinterERC20(ERC20 common.Address) (bool, error) {
	return _DynamicMultiSig.Contract.IsMinterERC20(&_DynamicMultiSig.CallOpts, ERC20)
}

// IsMinterERC20 is a free data retrieval call binding the contract method 0x6a7142e1.
//
// Solidity: function isMinterERC20(address ERC20) view returns(bool)
func (_DynamicMultiSig *DynamicMultiSigCallerSession) IsMinterERC20(ERC20 common.Address) (bool, error) {
	return _DynamicMultiSig.Contract.IsMinterERC20(&_DynamicMultiSig.CallOpts, ERC20)
}

// IsMinterERC721 is a free data retrieval call binding the contract method 0x279fb6e0.
//
// Solidity: function isMinterERC721(address ERC721) view returns(bool)
func (_DynamicMultiSig *DynamicMultiSigCaller) IsMinterERC721(opts *bind.CallOpts, ERC721 common.Address) (bool, error) {
	var out []interface{}
	err := _DynamicMultiSig.contract.Call(opts, &out, "isMinterERC721", ERC721)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMinterERC721 is a free data retrieval call binding the contract method 0x279fb6e0.
//
// Solidity: function isMinterERC721(address ERC721) view returns(bool)
func (_DynamicMultiSig *DynamicMultiSigSession) IsMinterERC721(ERC721 common.Address) (bool, error) {
	return _DynamicMultiSig.Contract.IsMinterERC721(&_DynamicMultiSig.CallOpts, ERC721)
}

// IsMinterERC721 is a free data retrieval call binding the contract method 0x279fb6e0.
//
// Solidity: function isMinterERC721(address ERC721) view returns(bool)
func (_DynamicMultiSig *DynamicMultiSigCallerSession) IsMinterERC721(ERC721 common.Address) (bool, error) {
	return _DynamicMultiSig.Contract.IsMinterERC721(&_DynamicMultiSig.CallOpts, ERC721)
}

// MaxManagers is a free data retrieval call binding the contract method 0xf7f2ff74.
//
// Solidity: function max_managers() view returns(uint256)
func (_DynamicMultiSig *DynamicMultiSigCaller) MaxManagers(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DynamicMultiSig.contract.Call(opts, &out, "max_managers")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxManagers is a free data retrieval call binding the contract method 0xf7f2ff74.
//
// Solidity: function max_managers() view returns(uint256)
func (_DynamicMultiSig *DynamicMultiSigSession) MaxManagers() (*big.Int, error) {
	return _DynamicMultiSig.Contract.MaxManagers(&_DynamicMultiSig.CallOpts)
}

// MaxManagers is a free data retrieval call binding the contract method 0xf7f2ff74.
//
// Solidity: function max_managers() view returns(uint256)
func (_DynamicMultiSig *DynamicMultiSigCallerSession) MaxManagers() (*big.Int, error) {
	return _DynamicMultiSig.Contract.MaxManagers(&_DynamicMultiSig.CallOpts)
}

// MinManagers is a free data retrieval call binding the contract method 0xad4b61a8.
//
// Solidity: function min_managers() view returns(uint256)
func (_DynamicMultiSig *DynamicMultiSigCaller) MinManagers(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DynamicMultiSig.contract.Call(opts, &out, "min_managers")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinManagers is a free data retrieval call binding the contract method 0xad4b61a8.
//
// Solidity: function min_managers() view returns(uint256)
func (_DynamicMultiSig *DynamicMultiSigSession) MinManagers() (*big.Int, error) {
	return _DynamicMultiSig.Contract.MinManagers(&_DynamicMultiSig.CallOpts)
}

// MinManagers is a free data retrieval call binding the contract method 0xad4b61a8.
//
// Solidity: function min_managers() view returns(uint256)
func (_DynamicMultiSig *DynamicMultiSigCallerSession) MinManagers() (*big.Int, error) {
	return _DynamicMultiSig.Contract.MinManagers(&_DynamicMultiSig.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DynamicMultiSig *DynamicMultiSigCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DynamicMultiSig.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DynamicMultiSig *DynamicMultiSigSession) Owner() (common.Address, error) {
	return _DynamicMultiSig.Contract.Owner(&_DynamicMultiSig.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DynamicMultiSig *DynamicMultiSigCallerSession) Owner() (common.Address, error) {
	return _DynamicMultiSig.Contract.Owner(&_DynamicMultiSig.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_DynamicMultiSig *DynamicMultiSigCaller) Rate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DynamicMultiSig.contract.Call(opts, &out, "rate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_DynamicMultiSig *DynamicMultiSigSession) Rate() (*big.Int, error) {
	return _DynamicMultiSig.Contract.Rate(&_DynamicMultiSig.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_DynamicMultiSig *DynamicMultiSigCallerSession) Rate() (*big.Int, error) {
	return _DynamicMultiSig.Contract.Rate(&_DynamicMultiSig.CallOpts)
}

// SignatureLength is a free data retrieval call binding the contract method 0x1b9a9323.
//
// Solidity: function signatureLength() view returns(uint256)
func (_DynamicMultiSig *DynamicMultiSigCaller) SignatureLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DynamicMultiSig.contract.Call(opts, &out, "signatureLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SignatureLength is a free data retrieval call binding the contract method 0x1b9a9323.
//
// Solidity: function signatureLength() view returns(uint256)
func (_DynamicMultiSig *DynamicMultiSigSession) SignatureLength() (*big.Int, error) {
	return _DynamicMultiSig.Contract.SignatureLength(&_DynamicMultiSig.CallOpts)
}

// SignatureLength is a free data retrieval call binding the contract method 0x1b9a9323.
//
// Solidity: function signatureLength() view returns(uint256)
func (_DynamicMultiSig *DynamicMultiSigCallerSession) SignatureLength() (*big.Int, error) {
	return _DynamicMultiSig.Contract.SignatureLength(&_DynamicMultiSig.CallOpts)
}

// Upgrade is a free data retrieval call binding the contract method 0xd55ec697.
//
// Solidity: function upgrade() view returns(bool)
func (_DynamicMultiSig *DynamicMultiSigCaller) Upgrade(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _DynamicMultiSig.contract.Call(opts, &out, "upgrade")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Upgrade is a free data retrieval call binding the contract method 0xd55ec697.
//
// Solidity: function upgrade() view returns(bool)
func (_DynamicMultiSig *DynamicMultiSigSession) Upgrade() (bool, error) {
	return _DynamicMultiSig.Contract.Upgrade(&_DynamicMultiSig.CallOpts)
}

// Upgrade is a free data retrieval call binding the contract method 0xd55ec697.
//
// Solidity: function upgrade() view returns(bool)
func (_DynamicMultiSig *DynamicMultiSigCallerSession) Upgrade() (bool, error) {
	return _DynamicMultiSig.Contract.Upgrade(&_DynamicMultiSig.CallOpts)
}

// UpgradeContractAddress is a free data retrieval call binding the contract method 0x30b2d84d.
//
// Solidity: function upgradeContractAddress() view returns(address)
func (_DynamicMultiSig *DynamicMultiSigCaller) UpgradeContractAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DynamicMultiSig.contract.Call(opts, &out, "upgradeContractAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UpgradeContractAddress is a free data retrieval call binding the contract method 0x30b2d84d.
//
// Solidity: function upgradeContractAddress() view returns(address)
func (_DynamicMultiSig *DynamicMultiSigSession) UpgradeContractAddress() (common.Address, error) {
	return _DynamicMultiSig.Contract.UpgradeContractAddress(&_DynamicMultiSig.CallOpts)
}

// UpgradeContractAddress is a free data retrieval call binding the contract method 0x30b2d84d.
//
// Solidity: function upgradeContractAddress() view returns(address)
func (_DynamicMultiSig *DynamicMultiSigCallerSession) UpgradeContractAddress() (common.Address, error) {
	return _DynamicMultiSig.Contract.UpgradeContractAddress(&_DynamicMultiSig.CallOpts)
}

// CloseUpgrade is a paid mutator transaction binding the contract method 0xd4cacbaa.
//
// Solidity: function closeUpgrade() returns()
func (_DynamicMultiSig *DynamicMultiSigTransactor) CloseUpgrade(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DynamicMultiSig.contract.Transact(opts, "closeUpgrade")
}

// CloseUpgrade is a paid mutator transaction binding the contract method 0xd4cacbaa.
//
// Solidity: function closeUpgrade() returns()
func (_DynamicMultiSig *DynamicMultiSigSession) CloseUpgrade() (*types.Transaction, error) {
	return _DynamicMultiSig.Contract.CloseUpgrade(&_DynamicMultiSig.TransactOpts)
}

// CloseUpgrade is a paid mutator transaction binding the contract method 0xd4cacbaa.
//
// Solidity: function closeUpgrade() returns()
func (_DynamicMultiSig *DynamicMultiSigTransactorSession) CloseUpgrade() (*types.Transaction, error) {
	return _DynamicMultiSig.Contract.CloseUpgrade(&_DynamicMultiSig.TransactOpts)
}

// CreateOrSignManagerChange is a paid mutator transaction binding the contract method 0x00719226.
//
// Solidity: function createOrSignManagerChange(string txKey, address[] adds, address[] removes, uint8 count, bytes signatures) returns()
func (_DynamicMultiSig *DynamicMultiSigTransactor) CreateOrSignManagerChange(opts *bind.TransactOpts, txKey string, adds []common.Address, removes []common.Address, count uint8, signatures []byte) (*types.Transaction, error) {
	return _DynamicMultiSig.contract.Transact(opts, "createOrSignManagerChange", txKey, adds, removes, count, signatures)
}

// CreateOrSignManagerChange is a paid mutator transaction binding the contract method 0x00719226.
//
// Solidity: function createOrSignManagerChange(string txKey, address[] adds, address[] removes, uint8 count, bytes signatures) returns()
func (_DynamicMultiSig *DynamicMultiSigSession) CreateOrSignManagerChange(txKey string, adds []common.Address, removes []common.Address, count uint8, signatures []byte) (*types.Transaction, error) {
	return _DynamicMultiSig.Contract.CreateOrSignManagerChange(&_DynamicMultiSig.TransactOpts, txKey, adds, removes, count, signatures)
}

// CreateOrSignManagerChange is a paid mutator transaction binding the contract method 0x00719226.
//
// Solidity: function createOrSignManagerChange(string txKey, address[] adds, address[] removes, uint8 count, bytes signatures) returns()
func (_DynamicMultiSig *DynamicMultiSigTransactorSession) CreateOrSignManagerChange(txKey string, adds []common.Address, removes []common.Address, count uint8, signatures []byte) (*types.Transaction, error) {
	return _DynamicMultiSig.Contract.CreateOrSignManagerChange(&_DynamicMultiSig.TransactOpts, txKey, adds, removes, count, signatures)
}

// CreateOrSignUpgrade is a paid mutator transaction binding the contract method 0x408e8b7a.
//
// Solidity: function createOrSignUpgrade(string txKey, address upgradeContract, bytes signatures) returns()
func (_DynamicMultiSig *DynamicMultiSigTransactor) CreateOrSignUpgrade(opts *bind.TransactOpts, txKey string, upgradeContract common.Address, signatures []byte) (*types.Transaction, error) {
	return _DynamicMultiSig.contract.Transact(opts, "createOrSignUpgrade", txKey, upgradeContract, signatures)
}

// CreateOrSignUpgrade is a paid mutator transaction binding the contract method 0x408e8b7a.
//
// Solidity: function createOrSignUpgrade(string txKey, address upgradeContract, bytes signatures) returns()
func (_DynamicMultiSig *DynamicMultiSigSession) CreateOrSignUpgrade(txKey string, upgradeContract common.Address, signatures []byte) (*types.Transaction, error) {
	return _DynamicMultiSig.Contract.CreateOrSignUpgrade(&_DynamicMultiSig.TransactOpts, txKey, upgradeContract, signatures)
}

// CreateOrSignUpgrade is a paid mutator transaction binding the contract method 0x408e8b7a.
//
// Solidity: function createOrSignUpgrade(string txKey, address upgradeContract, bytes signatures) returns()
func (_DynamicMultiSig *DynamicMultiSigTransactorSession) CreateOrSignUpgrade(txKey string, upgradeContract common.Address, signatures []byte) (*types.Transaction, error) {
	return _DynamicMultiSig.Contract.CreateOrSignUpgrade(&_DynamicMultiSig.TransactOpts, txKey, upgradeContract, signatures)
}

// CreateOrSignWithdraw is a paid mutator transaction binding the contract method 0xab6c2b10.
//
// Solidity: function createOrSignWithdraw(string txKey, address to, uint256 amount, bool isERC20, address ERC20, bytes signatures) returns()
func (_DynamicMultiSig *DynamicMultiSigTransactor) CreateOrSignWithdraw(opts *bind.TransactOpts, txKey string, to common.Address, amount *big.Int, isERC20 bool, ERC20 common.Address, signatures []byte) (*types.Transaction, error) {
	return _DynamicMultiSig.contract.Transact(opts, "createOrSignWithdraw", txKey, to, amount, isERC20, ERC20, signatures)
}

// CreateOrSignWithdraw is a paid mutator transaction binding the contract method 0xab6c2b10.
//
// Solidity: function createOrSignWithdraw(string txKey, address to, uint256 amount, bool isERC20, address ERC20, bytes signatures) returns()
func (_DynamicMultiSig *DynamicMultiSigSession) CreateOrSignWithdraw(txKey string, to common.Address, amount *big.Int, isERC20 bool, ERC20 common.Address, signatures []byte) (*types.Transaction, error) {
	return _DynamicMultiSig.Contract.CreateOrSignWithdraw(&_DynamicMultiSig.TransactOpts, txKey, to, amount, isERC20, ERC20, signatures)
}

// CreateOrSignWithdraw is a paid mutator transaction binding the contract method 0xab6c2b10.
//
// Solidity: function createOrSignWithdraw(string txKey, address to, uint256 amount, bool isERC20, address ERC20, bytes signatures) returns()
func (_DynamicMultiSig *DynamicMultiSigTransactorSession) CreateOrSignWithdraw(txKey string, to common.Address, amount *big.Int, isERC20 bool, ERC20 common.Address, signatures []byte) (*types.Transaction, error) {
	return _DynamicMultiSig.Contract.CreateOrSignWithdraw(&_DynamicMultiSig.TransactOpts, txKey, to, amount, isERC20, ERC20, signatures)
}

// CreateOrSignWithdrawNFT is a paid mutator transaction binding the contract method 0x1c702f35.
//
// Solidity: function createOrSignWithdrawNFT(string txKey, address ERC721, address to, uint256 tokenId, string tokenURI, bytes signatures) returns()
func (_DynamicMultiSig *DynamicMultiSigTransactor) CreateOrSignWithdrawNFT(opts *bind.TransactOpts, txKey string, ERC721 common.Address, to common.Address, tokenId *big.Int, tokenURI string, signatures []byte) (*types.Transaction, error) {
	return _DynamicMultiSig.contract.Transact(opts, "createOrSignWithdrawNFT", txKey, ERC721, to, tokenId, tokenURI, signatures)
}

// CreateOrSignWithdrawNFT is a paid mutator transaction binding the contract method 0x1c702f35.
//
// Solidity: function createOrSignWithdrawNFT(string txKey, address ERC721, address to, uint256 tokenId, string tokenURI, bytes signatures) returns()
func (_DynamicMultiSig *DynamicMultiSigSession) CreateOrSignWithdrawNFT(txKey string, ERC721 common.Address, to common.Address, tokenId *big.Int, tokenURI string, signatures []byte) (*types.Transaction, error) {
	return _DynamicMultiSig.Contract.CreateOrSignWithdrawNFT(&_DynamicMultiSig.TransactOpts, txKey, ERC721, to, tokenId, tokenURI, signatures)
}

// CreateOrSignWithdrawNFT is a paid mutator transaction binding the contract method 0x1c702f35.
//
// Solidity: function createOrSignWithdrawNFT(string txKey, address ERC721, address to, uint256 tokenId, string tokenURI, bytes signatures) returns()
func (_DynamicMultiSig *DynamicMultiSigTransactorSession) CreateOrSignWithdrawNFT(txKey string, ERC721 common.Address, to common.Address, tokenId *big.Int, tokenURI string, signatures []byte) (*types.Transaction, error) {
	return _DynamicMultiSig.Contract.CreateOrSignWithdrawNFT(&_DynamicMultiSig.TransactOpts, txKey, ERC721, to, tokenId, tokenURI, signatures)
}

// CrossOut is a paid mutator transaction binding the contract method 0x0889d1f0.
//
// Solidity: function crossOut(string to, uint256 amount, address ERC20) payable returns(bool)
func (_DynamicMultiSig *DynamicMultiSigTransactor) CrossOut(opts *bind.TransactOpts, to string, amount *big.Int, ERC20 common.Address) (*types.Transaction, error) {
	return _DynamicMultiSig.contract.Transact(opts, "crossOut", to, amount, ERC20)
}

// CrossOut is a paid mutator transaction binding the contract method 0x0889d1f0.
//
// Solidity: function crossOut(string to, uint256 amount, address ERC20) payable returns(bool)
func (_DynamicMultiSig *DynamicMultiSigSession) CrossOut(to string, amount *big.Int, ERC20 common.Address) (*types.Transaction, error) {
	return _DynamicMultiSig.Contract.CrossOut(&_DynamicMultiSig.TransactOpts, to, amount, ERC20)
}

// CrossOut is a paid mutator transaction binding the contract method 0x0889d1f0.
//
// Solidity: function crossOut(string to, uint256 amount, address ERC20) payable returns(bool)
func (_DynamicMultiSig *DynamicMultiSigTransactorSession) CrossOut(to string, amount *big.Int, ERC20 common.Address) (*types.Transaction, error) {
	return _DynamicMultiSig.Contract.CrossOut(&_DynamicMultiSig.TransactOpts, to, amount, ERC20)
}

// CrossOutNFT is a paid mutator transaction binding the contract method 0xc76bbdc5.
//
// Solidity: function crossOutNFT(string to, address ERC721, uint256 tokenId) returns(bool)
func (_DynamicMultiSig *DynamicMultiSigTransactor) CrossOutNFT(opts *bind.TransactOpts, to string, ERC721 common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _DynamicMultiSig.contract.Transact(opts, "crossOutNFT", to, ERC721, tokenId)
}

// CrossOutNFT is a paid mutator transaction binding the contract method 0xc76bbdc5.
//
// Solidity: function crossOutNFT(string to, address ERC721, uint256 tokenId) returns(bool)
func (_DynamicMultiSig *DynamicMultiSigSession) CrossOutNFT(to string, ERC721 common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _DynamicMultiSig.Contract.CrossOutNFT(&_DynamicMultiSig.TransactOpts, to, ERC721, tokenId)
}

// CrossOutNFT is a paid mutator transaction binding the contract method 0xc76bbdc5.
//
// Solidity: function crossOutNFT(string to, address ERC721, uint256 tokenId) returns(bool)
func (_DynamicMultiSig *DynamicMultiSigTransactorSession) CrossOutNFT(to string, ERC721 common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _DynamicMultiSig.Contract.CrossOutNFT(&_DynamicMultiSig.TransactOpts, to, ERC721, tokenId)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_DynamicMultiSig *DynamicMultiSigTransactor) OnERC721Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _DynamicMultiSig.contract.Transact(opts, "onERC721Received", arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_DynamicMultiSig *DynamicMultiSigSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _DynamicMultiSig.Contract.OnERC721Received(&_DynamicMultiSig.TransactOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_DynamicMultiSig *DynamicMultiSigTransactorSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _DynamicMultiSig.Contract.OnERC721Received(&_DynamicMultiSig.TransactOpts, arg0, arg1, arg2, arg3)
}

// RegisterMinterERC20 is a paid mutator transaction binding the contract method 0x34774b71.
//
// Solidity: function registerMinterERC20(address ERC20) returns()
func (_DynamicMultiSig *DynamicMultiSigTransactor) RegisterMinterERC20(opts *bind.TransactOpts, ERC20 common.Address) (*types.Transaction, error) {
	return _DynamicMultiSig.contract.Transact(opts, "registerMinterERC20", ERC20)
}

// RegisterMinterERC20 is a paid mutator transaction binding the contract method 0x34774b71.
//
// Solidity: function registerMinterERC20(address ERC20) returns()
func (_DynamicMultiSig *DynamicMultiSigSession) RegisterMinterERC20(ERC20 common.Address) (*types.Transaction, error) {
	return _DynamicMultiSig.Contract.RegisterMinterERC20(&_DynamicMultiSig.TransactOpts, ERC20)
}

// RegisterMinterERC20 is a paid mutator transaction binding the contract method 0x34774b71.
//
// Solidity: function registerMinterERC20(address ERC20) returns()
func (_DynamicMultiSig *DynamicMultiSigTransactorSession) RegisterMinterERC20(ERC20 common.Address) (*types.Transaction, error) {
	return _DynamicMultiSig.Contract.RegisterMinterERC20(&_DynamicMultiSig.TransactOpts, ERC20)
}

// RegisterMinterERC721 is a paid mutator transaction binding the contract method 0x7de24876.
//
// Solidity: function registerMinterERC721(address ERC721) returns()
func (_DynamicMultiSig *DynamicMultiSigTransactor) RegisterMinterERC721(opts *bind.TransactOpts, ERC721 common.Address) (*types.Transaction, error) {
	return _DynamicMultiSig.contract.Transact(opts, "registerMinterERC721", ERC721)
}

// RegisterMinterERC721 is a paid mutator transaction binding the contract method 0x7de24876.
//
// Solidity: function registerMinterERC721(address ERC721) returns()
func (_DynamicMultiSig *DynamicMultiSigSession) RegisterMinterERC721(ERC721 common.Address) (*types.Transaction, error) {
	return _DynamicMultiSig.Contract.RegisterMinterERC721(&_DynamicMultiSig.TransactOpts, ERC721)
}

// RegisterMinterERC721 is a paid mutator transaction binding the contract method 0x7de24876.
//
// Solidity: function registerMinterERC721(address ERC721) returns()
func (_DynamicMultiSig *DynamicMultiSigTransactorSession) RegisterMinterERC721(ERC721 common.Address) (*types.Transaction, error) {
	return _DynamicMultiSig.Contract.RegisterMinterERC721(&_DynamicMultiSig.TransactOpts, ERC721)
}

// UnregisterMinterERC20 is a paid mutator transaction binding the contract method 0x9dcdc978.
//
// Solidity: function unregisterMinterERC20(address ERC20) returns()
func (_DynamicMultiSig *DynamicMultiSigTransactor) UnregisterMinterERC20(opts *bind.TransactOpts, ERC20 common.Address) (*types.Transaction, error) {
	return _DynamicMultiSig.contract.Transact(opts, "unregisterMinterERC20", ERC20)
}

// UnregisterMinterERC20 is a paid mutator transaction binding the contract method 0x9dcdc978.
//
// Solidity: function unregisterMinterERC20(address ERC20) returns()
func (_DynamicMultiSig *DynamicMultiSigSession) UnregisterMinterERC20(ERC20 common.Address) (*types.Transaction, error) {
	return _DynamicMultiSig.Contract.UnregisterMinterERC20(&_DynamicMultiSig.TransactOpts, ERC20)
}

// UnregisterMinterERC20 is a paid mutator transaction binding the contract method 0x9dcdc978.
//
// Solidity: function unregisterMinterERC20(address ERC20) returns()
func (_DynamicMultiSig *DynamicMultiSigTransactorSession) UnregisterMinterERC20(ERC20 common.Address) (*types.Transaction, error) {
	return _DynamicMultiSig.Contract.UnregisterMinterERC20(&_DynamicMultiSig.TransactOpts, ERC20)
}

// UnregisterMinterERC721 is a paid mutator transaction binding the contract method 0xbc8870ed.
//
// Solidity: function unregisterMinterERC721(address ERC721) returns()
func (_DynamicMultiSig *DynamicMultiSigTransactor) UnregisterMinterERC721(opts *bind.TransactOpts, ERC721 common.Address) (*types.Transaction, error) {
	return _DynamicMultiSig.contract.Transact(opts, "unregisterMinterERC721", ERC721)
}

// UnregisterMinterERC721 is a paid mutator transaction binding the contract method 0xbc8870ed.
//
// Solidity: function unregisterMinterERC721(address ERC721) returns()
func (_DynamicMultiSig *DynamicMultiSigSession) UnregisterMinterERC721(ERC721 common.Address) (*types.Transaction, error) {
	return _DynamicMultiSig.Contract.UnregisterMinterERC721(&_DynamicMultiSig.TransactOpts, ERC721)
}

// UnregisterMinterERC721 is a paid mutator transaction binding the contract method 0xbc8870ed.
//
// Solidity: function unregisterMinterERC721(address ERC721) returns()
func (_DynamicMultiSig *DynamicMultiSigTransactorSession) UnregisterMinterERC721(ERC721 common.Address) (*types.Transaction, error) {
	return _DynamicMultiSig.Contract.UnregisterMinterERC721(&_DynamicMultiSig.TransactOpts, ERC721)
}

// UpgradeContractS1 is a paid mutator transaction binding the contract method 0x1dda9c05.
//
// Solidity: function upgradeContractS1() returns()
func (_DynamicMultiSig *DynamicMultiSigTransactor) UpgradeContractS1(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DynamicMultiSig.contract.Transact(opts, "upgradeContractS1")
}

// UpgradeContractS1 is a paid mutator transaction binding the contract method 0x1dda9c05.
//
// Solidity: function upgradeContractS1() returns()
func (_DynamicMultiSig *DynamicMultiSigSession) UpgradeContractS1() (*types.Transaction, error) {
	return _DynamicMultiSig.Contract.UpgradeContractS1(&_DynamicMultiSig.TransactOpts)
}

// UpgradeContractS1 is a paid mutator transaction binding the contract method 0x1dda9c05.
//
// Solidity: function upgradeContractS1() returns()
func (_DynamicMultiSig *DynamicMultiSigTransactorSession) UpgradeContractS1() (*types.Transaction, error) {
	return _DynamicMultiSig.Contract.UpgradeContractS1(&_DynamicMultiSig.TransactOpts)
}

// UpgradeContractS2 is a paid mutator transaction binding the contract method 0x5bda3fcf.
//
// Solidity: function upgradeContractS2(address ERC20) returns()
func (_DynamicMultiSig *DynamicMultiSigTransactor) UpgradeContractS2(opts *bind.TransactOpts, ERC20 common.Address) (*types.Transaction, error) {
	return _DynamicMultiSig.contract.Transact(opts, "upgradeContractS2", ERC20)
}

// UpgradeContractS2 is a paid mutator transaction binding the contract method 0x5bda3fcf.
//
// Solidity: function upgradeContractS2(address ERC20) returns()
func (_DynamicMultiSig *DynamicMultiSigSession) UpgradeContractS2(ERC20 common.Address) (*types.Transaction, error) {
	return _DynamicMultiSig.Contract.UpgradeContractS2(&_DynamicMultiSig.TransactOpts, ERC20)
}

// UpgradeContractS2 is a paid mutator transaction binding the contract method 0x5bda3fcf.
//
// Solidity: function upgradeContractS2(address ERC20) returns()
func (_DynamicMultiSig *DynamicMultiSigTransactorSession) UpgradeContractS2(ERC20 common.Address) (*types.Transaction, error) {
	return _DynamicMultiSig.Contract.UpgradeContractS2(&_DynamicMultiSig.TransactOpts, ERC20)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_DynamicMultiSig *DynamicMultiSigTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _DynamicMultiSig.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_DynamicMultiSig *DynamicMultiSigSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _DynamicMultiSig.Contract.Fallback(&_DynamicMultiSig.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_DynamicMultiSig *DynamicMultiSigTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _DynamicMultiSig.Contract.Fallback(&_DynamicMultiSig.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_DynamicMultiSig *DynamicMultiSigTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DynamicMultiSig.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_DynamicMultiSig *DynamicMultiSigSession) Receive() (*types.Transaction, error) {
	return _DynamicMultiSig.Contract.Receive(&_DynamicMultiSig.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_DynamicMultiSig *DynamicMultiSigTransactorSession) Receive() (*types.Transaction, error) {
	return _DynamicMultiSig.Contract.Receive(&_DynamicMultiSig.TransactOpts)
}

// DynamicMultiSigCrossOutFundsIterator is returned from FilterCrossOutFunds and is used to iterate over the raw logs and unpacked data for CrossOutFunds events raised by the DynamicMultiSig contract.
type DynamicMultiSigCrossOutFundsIterator struct {
	Event *DynamicMultiSigCrossOutFunds // Event containing the contract specifics and raw log

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
func (it *DynamicMultiSigCrossOutFundsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DynamicMultiSigCrossOutFunds)
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
		it.Event = new(DynamicMultiSigCrossOutFunds)
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
func (it *DynamicMultiSigCrossOutFundsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DynamicMultiSigCrossOutFundsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DynamicMultiSigCrossOutFunds represents a CrossOutFunds event raised by the DynamicMultiSig contract.
type DynamicMultiSigCrossOutFunds struct {
	From   common.Address
	To     string
	Amount *big.Int
	ERC20  common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCrossOutFunds is a free log retrieval operation binding the contract event 0x5ddf9724d8fe5d9e12499be2867f93d41a582733dcd65f74a486ad7e30667146.
//
// Solidity: event CrossOutFunds(address from, string to, uint256 amount, address ERC20)
func (_DynamicMultiSig *DynamicMultiSigFilterer) FilterCrossOutFunds(opts *bind.FilterOpts) (*DynamicMultiSigCrossOutFundsIterator, error) {

	logs, sub, err := _DynamicMultiSig.contract.FilterLogs(opts, "CrossOutFunds")
	if err != nil {
		return nil, err
	}
	return &DynamicMultiSigCrossOutFundsIterator{contract: _DynamicMultiSig.contract, event: "CrossOutFunds", logs: logs, sub: sub}, nil
}

// WatchCrossOutFunds is a free log subscription operation binding the contract event 0x5ddf9724d8fe5d9e12499be2867f93d41a582733dcd65f74a486ad7e30667146.
//
// Solidity: event CrossOutFunds(address from, string to, uint256 amount, address ERC20)
func (_DynamicMultiSig *DynamicMultiSigFilterer) WatchCrossOutFunds(opts *bind.WatchOpts, sink chan<- *DynamicMultiSigCrossOutFunds) (event.Subscription, error) {

	logs, sub, err := _DynamicMultiSig.contract.WatchLogs(opts, "CrossOutFunds")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DynamicMultiSigCrossOutFunds)
				if err := _DynamicMultiSig.contract.UnpackLog(event, "CrossOutFunds", log); err != nil {
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

// ParseCrossOutFunds is a log parse operation binding the contract event 0x5ddf9724d8fe5d9e12499be2867f93d41a582733dcd65f74a486ad7e30667146.
//
// Solidity: event CrossOutFunds(address from, string to, uint256 amount, address ERC20)
func (_DynamicMultiSig *DynamicMultiSigFilterer) ParseCrossOutFunds(log types.Log) (*DynamicMultiSigCrossOutFunds, error) {
	event := new(DynamicMultiSigCrossOutFunds)
	if err := _DynamicMultiSig.contract.UnpackLog(event, "CrossOutFunds", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DynamicMultiSigCrossOutNFTIterator is returned from FilterCrossOutNFT and is used to iterate over the raw logs and unpacked data for CrossOutNFT events raised by the DynamicMultiSig contract.
type DynamicMultiSigCrossOutNFTIterator struct {
	Event *DynamicMultiSigCrossOutNFT // Event containing the contract specifics and raw log

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
func (it *DynamicMultiSigCrossOutNFTIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DynamicMultiSigCrossOutNFT)
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
		it.Event = new(DynamicMultiSigCrossOutNFT)
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
func (it *DynamicMultiSigCrossOutNFTIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DynamicMultiSigCrossOutNFTIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DynamicMultiSigCrossOutNFT represents a CrossOutNFT event raised by the DynamicMultiSig contract.
type DynamicMultiSigCrossOutNFT struct {
	From    common.Address
	To      string
	TokenId *big.Int
	ERC721  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterCrossOutNFT is a free log retrieval operation binding the contract event 0x06caac0500ba0f6fbe6852539c648ab228f15265797d52ca06096c9e74bb17df.
//
// Solidity: event CrossOutNFT(address from, string to, uint256 tokenId, address ERC721)
func (_DynamicMultiSig *DynamicMultiSigFilterer) FilterCrossOutNFT(opts *bind.FilterOpts) (*DynamicMultiSigCrossOutNFTIterator, error) {

	logs, sub, err := _DynamicMultiSig.contract.FilterLogs(opts, "CrossOutNFT")
	if err != nil {
		return nil, err
	}
	return &DynamicMultiSigCrossOutNFTIterator{contract: _DynamicMultiSig.contract, event: "CrossOutNFT", logs: logs, sub: sub}, nil
}

// WatchCrossOutNFT is a free log subscription operation binding the contract event 0x06caac0500ba0f6fbe6852539c648ab228f15265797d52ca06096c9e74bb17df.
//
// Solidity: event CrossOutNFT(address from, string to, uint256 tokenId, address ERC721)
func (_DynamicMultiSig *DynamicMultiSigFilterer) WatchCrossOutNFT(opts *bind.WatchOpts, sink chan<- *DynamicMultiSigCrossOutNFT) (event.Subscription, error) {

	logs, sub, err := _DynamicMultiSig.contract.WatchLogs(opts, "CrossOutNFT")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DynamicMultiSigCrossOutNFT)
				if err := _DynamicMultiSig.contract.UnpackLog(event, "CrossOutNFT", log); err != nil {
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

// ParseCrossOutNFT is a log parse operation binding the contract event 0x06caac0500ba0f6fbe6852539c648ab228f15265797d52ca06096c9e74bb17df.
//
// Solidity: event CrossOutNFT(address from, string to, uint256 tokenId, address ERC721)
func (_DynamicMultiSig *DynamicMultiSigFilterer) ParseCrossOutNFT(log types.Log) (*DynamicMultiSigCrossOutNFT, error) {
	event := new(DynamicMultiSigCrossOutNFT)
	if err := _DynamicMultiSig.contract.UnpackLog(event, "CrossOutNFT", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DynamicMultiSigDepositFundsIterator is returned from FilterDepositFunds and is used to iterate over the raw logs and unpacked data for DepositFunds events raised by the DynamicMultiSig contract.
type DynamicMultiSigDepositFundsIterator struct {
	Event *DynamicMultiSigDepositFunds // Event containing the contract specifics and raw log

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
func (it *DynamicMultiSigDepositFundsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DynamicMultiSigDepositFunds)
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
		it.Event = new(DynamicMultiSigDepositFunds)
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
func (it *DynamicMultiSigDepositFundsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DynamicMultiSigDepositFundsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DynamicMultiSigDepositFunds represents a DepositFunds event raised by the DynamicMultiSig contract.
type DynamicMultiSigDepositFunds struct {
	From   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDepositFunds is a free log retrieval operation binding the contract event 0xd241e73300212f6df233a8e6d3146b88a9d4964e06621d54b5ff6afeba7b1b88.
//
// Solidity: event DepositFunds(address from, uint256 amount)
func (_DynamicMultiSig *DynamicMultiSigFilterer) FilterDepositFunds(opts *bind.FilterOpts) (*DynamicMultiSigDepositFundsIterator, error) {

	logs, sub, err := _DynamicMultiSig.contract.FilterLogs(opts, "DepositFunds")
	if err != nil {
		return nil, err
	}
	return &DynamicMultiSigDepositFundsIterator{contract: _DynamicMultiSig.contract, event: "DepositFunds", logs: logs, sub: sub}, nil
}

// WatchDepositFunds is a free log subscription operation binding the contract event 0xd241e73300212f6df233a8e6d3146b88a9d4964e06621d54b5ff6afeba7b1b88.
//
// Solidity: event DepositFunds(address from, uint256 amount)
func (_DynamicMultiSig *DynamicMultiSigFilterer) WatchDepositFunds(opts *bind.WatchOpts, sink chan<- *DynamicMultiSigDepositFunds) (event.Subscription, error) {

	logs, sub, err := _DynamicMultiSig.contract.WatchLogs(opts, "DepositFunds")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DynamicMultiSigDepositFunds)
				if err := _DynamicMultiSig.contract.UnpackLog(event, "DepositFunds", log); err != nil {
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

// ParseDepositFunds is a log parse operation binding the contract event 0xd241e73300212f6df233a8e6d3146b88a9d4964e06621d54b5ff6afeba7b1b88.
//
// Solidity: event DepositFunds(address from, uint256 amount)
func (_DynamicMultiSig *DynamicMultiSigFilterer) ParseDepositFunds(log types.Log) (*DynamicMultiSigDepositFunds, error) {
	event := new(DynamicMultiSigDepositFunds)
	if err := _DynamicMultiSig.contract.UnpackLog(event, "DepositFunds", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DynamicMultiSigTransferFundsIterator is returned from FilterTransferFunds and is used to iterate over the raw logs and unpacked data for TransferFunds events raised by the DynamicMultiSig contract.
type DynamicMultiSigTransferFundsIterator struct {
	Event *DynamicMultiSigTransferFunds // Event containing the contract specifics and raw log

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
func (it *DynamicMultiSigTransferFundsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DynamicMultiSigTransferFunds)
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
		it.Event = new(DynamicMultiSigTransferFunds)
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
func (it *DynamicMultiSigTransferFundsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DynamicMultiSigTransferFundsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DynamicMultiSigTransferFunds represents a TransferFunds event raised by the DynamicMultiSig contract.
type DynamicMultiSigTransferFunds struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransferFunds is a free log retrieval operation binding the contract event 0xc95f8b91b103304386b955ef73fadac189f8ad66b33369b6c34a17a60db7bd0a.
//
// Solidity: event TransferFunds(address to, uint256 amount)
func (_DynamicMultiSig *DynamicMultiSigFilterer) FilterTransferFunds(opts *bind.FilterOpts) (*DynamicMultiSigTransferFundsIterator, error) {

	logs, sub, err := _DynamicMultiSig.contract.FilterLogs(opts, "TransferFunds")
	if err != nil {
		return nil, err
	}
	return &DynamicMultiSigTransferFundsIterator{contract: _DynamicMultiSig.contract, event: "TransferFunds", logs: logs, sub: sub}, nil
}

// WatchTransferFunds is a free log subscription operation binding the contract event 0xc95f8b91b103304386b955ef73fadac189f8ad66b33369b6c34a17a60db7bd0a.
//
// Solidity: event TransferFunds(address to, uint256 amount)
func (_DynamicMultiSig *DynamicMultiSigFilterer) WatchTransferFunds(opts *bind.WatchOpts, sink chan<- *DynamicMultiSigTransferFunds) (event.Subscription, error) {

	logs, sub, err := _DynamicMultiSig.contract.WatchLogs(opts, "TransferFunds")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DynamicMultiSigTransferFunds)
				if err := _DynamicMultiSig.contract.UnpackLog(event, "TransferFunds", log); err != nil {
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

// ParseTransferFunds is a log parse operation binding the contract event 0xc95f8b91b103304386b955ef73fadac189f8ad66b33369b6c34a17a60db7bd0a.
//
// Solidity: event TransferFunds(address to, uint256 amount)
func (_DynamicMultiSig *DynamicMultiSigFilterer) ParseTransferFunds(log types.Log) (*DynamicMultiSigTransferFunds, error) {
	event := new(DynamicMultiSigTransferFunds)
	if err := _DynamicMultiSig.contract.UnpackLog(event, "TransferFunds", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DynamicMultiSigTxManagerChangeCompletedIterator is returned from FilterTxManagerChangeCompleted and is used to iterate over the raw logs and unpacked data for TxManagerChangeCompleted events raised by the DynamicMultiSig contract.
type DynamicMultiSigTxManagerChangeCompletedIterator struct {
	Event *DynamicMultiSigTxManagerChangeCompleted // Event containing the contract specifics and raw log

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
func (it *DynamicMultiSigTxManagerChangeCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DynamicMultiSigTxManagerChangeCompleted)
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
		it.Event = new(DynamicMultiSigTxManagerChangeCompleted)
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
func (it *DynamicMultiSigTxManagerChangeCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DynamicMultiSigTxManagerChangeCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DynamicMultiSigTxManagerChangeCompleted represents a TxManagerChangeCompleted event raised by the DynamicMultiSig contract.
type DynamicMultiSigTxManagerChangeCompleted struct {
	TxKey string
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTxManagerChangeCompleted is a free log retrieval operation binding the contract event 0xac9b82db4e104d515319a481096bfd91a4f40ee10837d5a2c8d51b9a03dc48ae.
//
// Solidity: event TxManagerChangeCompleted(string txKey)
func (_DynamicMultiSig *DynamicMultiSigFilterer) FilterTxManagerChangeCompleted(opts *bind.FilterOpts) (*DynamicMultiSigTxManagerChangeCompletedIterator, error) {

	logs, sub, err := _DynamicMultiSig.contract.FilterLogs(opts, "TxManagerChangeCompleted")
	if err != nil {
		return nil, err
	}
	return &DynamicMultiSigTxManagerChangeCompletedIterator{contract: _DynamicMultiSig.contract, event: "TxManagerChangeCompleted", logs: logs, sub: sub}, nil
}

// WatchTxManagerChangeCompleted is a free log subscription operation binding the contract event 0xac9b82db4e104d515319a481096bfd91a4f40ee10837d5a2c8d51b9a03dc48ae.
//
// Solidity: event TxManagerChangeCompleted(string txKey)
func (_DynamicMultiSig *DynamicMultiSigFilterer) WatchTxManagerChangeCompleted(opts *bind.WatchOpts, sink chan<- *DynamicMultiSigTxManagerChangeCompleted) (event.Subscription, error) {

	logs, sub, err := _DynamicMultiSig.contract.WatchLogs(opts, "TxManagerChangeCompleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DynamicMultiSigTxManagerChangeCompleted)
				if err := _DynamicMultiSig.contract.UnpackLog(event, "TxManagerChangeCompleted", log); err != nil {
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

// ParseTxManagerChangeCompleted is a log parse operation binding the contract event 0xac9b82db4e104d515319a481096bfd91a4f40ee10837d5a2c8d51b9a03dc48ae.
//
// Solidity: event TxManagerChangeCompleted(string txKey)
func (_DynamicMultiSig *DynamicMultiSigFilterer) ParseTxManagerChangeCompleted(log types.Log) (*DynamicMultiSigTxManagerChangeCompleted, error) {
	event := new(DynamicMultiSigTxManagerChangeCompleted)
	if err := _DynamicMultiSig.contract.UnpackLog(event, "TxManagerChangeCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DynamicMultiSigTxUpgradeCompletedIterator is returned from FilterTxUpgradeCompleted and is used to iterate over the raw logs and unpacked data for TxUpgradeCompleted events raised by the DynamicMultiSig contract.
type DynamicMultiSigTxUpgradeCompletedIterator struct {
	Event *DynamicMultiSigTxUpgradeCompleted // Event containing the contract specifics and raw log

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
func (it *DynamicMultiSigTxUpgradeCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DynamicMultiSigTxUpgradeCompleted)
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
		it.Event = new(DynamicMultiSigTxUpgradeCompleted)
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
func (it *DynamicMultiSigTxUpgradeCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DynamicMultiSigTxUpgradeCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DynamicMultiSigTxUpgradeCompleted represents a TxUpgradeCompleted event raised by the DynamicMultiSig contract.
type DynamicMultiSigTxUpgradeCompleted struct {
	TxKey string
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTxUpgradeCompleted is a free log retrieval operation binding the contract event 0x5e06c4b22547d430736ce834764dbfee08f1c4cf7ae3d53178aa56effa593ed0.
//
// Solidity: event TxUpgradeCompleted(string txKey)
func (_DynamicMultiSig *DynamicMultiSigFilterer) FilterTxUpgradeCompleted(opts *bind.FilterOpts) (*DynamicMultiSigTxUpgradeCompletedIterator, error) {

	logs, sub, err := _DynamicMultiSig.contract.FilterLogs(opts, "TxUpgradeCompleted")
	if err != nil {
		return nil, err
	}
	return &DynamicMultiSigTxUpgradeCompletedIterator{contract: _DynamicMultiSig.contract, event: "TxUpgradeCompleted", logs: logs, sub: sub}, nil
}

// WatchTxUpgradeCompleted is a free log subscription operation binding the contract event 0x5e06c4b22547d430736ce834764dbfee08f1c4cf7ae3d53178aa56effa593ed0.
//
// Solidity: event TxUpgradeCompleted(string txKey)
func (_DynamicMultiSig *DynamicMultiSigFilterer) WatchTxUpgradeCompleted(opts *bind.WatchOpts, sink chan<- *DynamicMultiSigTxUpgradeCompleted) (event.Subscription, error) {

	logs, sub, err := _DynamicMultiSig.contract.WatchLogs(opts, "TxUpgradeCompleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DynamicMultiSigTxUpgradeCompleted)
				if err := _DynamicMultiSig.contract.UnpackLog(event, "TxUpgradeCompleted", log); err != nil {
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

// ParseTxUpgradeCompleted is a log parse operation binding the contract event 0x5e06c4b22547d430736ce834764dbfee08f1c4cf7ae3d53178aa56effa593ed0.
//
// Solidity: event TxUpgradeCompleted(string txKey)
func (_DynamicMultiSig *DynamicMultiSigFilterer) ParseTxUpgradeCompleted(log types.Log) (*DynamicMultiSigTxUpgradeCompleted, error) {
	event := new(DynamicMultiSigTxUpgradeCompleted)
	if err := _DynamicMultiSig.contract.UnpackLog(event, "TxUpgradeCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DynamicMultiSigTxWithdrawCompletedIterator is returned from FilterTxWithdrawCompleted and is used to iterate over the raw logs and unpacked data for TxWithdrawCompleted events raised by the DynamicMultiSig contract.
type DynamicMultiSigTxWithdrawCompletedIterator struct {
	Event *DynamicMultiSigTxWithdrawCompleted // Event containing the contract specifics and raw log

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
func (it *DynamicMultiSigTxWithdrawCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DynamicMultiSigTxWithdrawCompleted)
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
		it.Event = new(DynamicMultiSigTxWithdrawCompleted)
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
func (it *DynamicMultiSigTxWithdrawCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DynamicMultiSigTxWithdrawCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DynamicMultiSigTxWithdrawCompleted represents a TxWithdrawCompleted event raised by the DynamicMultiSig contract.
type DynamicMultiSigTxWithdrawCompleted struct {
	TxKey string
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTxWithdrawCompleted is a free log retrieval operation binding the contract event 0x8ed8b1f0dd3babfdf1477ba2b27a5b0d2f1c9148448fd22cf2c75e658293c7b1.
//
// Solidity: event TxWithdrawCompleted(string txKey)
func (_DynamicMultiSig *DynamicMultiSigFilterer) FilterTxWithdrawCompleted(opts *bind.FilterOpts) (*DynamicMultiSigTxWithdrawCompletedIterator, error) {

	logs, sub, err := _DynamicMultiSig.contract.FilterLogs(opts, "TxWithdrawCompleted")
	if err != nil {
		return nil, err
	}
	return &DynamicMultiSigTxWithdrawCompletedIterator{contract: _DynamicMultiSig.contract, event: "TxWithdrawCompleted", logs: logs, sub: sub}, nil
}

// WatchTxWithdrawCompleted is a free log subscription operation binding the contract event 0x8ed8b1f0dd3babfdf1477ba2b27a5b0d2f1c9148448fd22cf2c75e658293c7b1.
//
// Solidity: event TxWithdrawCompleted(string txKey)
func (_DynamicMultiSig *DynamicMultiSigFilterer) WatchTxWithdrawCompleted(opts *bind.WatchOpts, sink chan<- *DynamicMultiSigTxWithdrawCompleted) (event.Subscription, error) {

	logs, sub, err := _DynamicMultiSig.contract.WatchLogs(opts, "TxWithdrawCompleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DynamicMultiSigTxWithdrawCompleted)
				if err := _DynamicMultiSig.contract.UnpackLog(event, "TxWithdrawCompleted", log); err != nil {
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

// ParseTxWithdrawCompleted is a log parse operation binding the contract event 0x8ed8b1f0dd3babfdf1477ba2b27a5b0d2f1c9148448fd22cf2c75e658293c7b1.
//
// Solidity: event TxWithdrawCompleted(string txKey)
func (_DynamicMultiSig *DynamicMultiSigFilterer) ParseTxWithdrawCompleted(log types.Log) (*DynamicMultiSigTxWithdrawCompleted, error) {
	event := new(DynamicMultiSigTxWithdrawCompleted)
	if err := _DynamicMultiSig.contract.UnpackLog(event, "TxWithdrawCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC165ABI is the input ABI used to generate the binding from.
const IERC165ABI = "[{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IERC165FuncSigs maps the 4-byte function signature to its string representation.
var IERC165FuncSigs = map[string]string{
	"01ffc9a7": "supportsInterface(bytes4)",
}

// IERC165 is an auto generated Go binding around an Ethereum contract.
type IERC165 struct {
	IERC165Caller     // Read-only binding to the contract
	IERC165Transactor // Write-only binding to the contract
	IERC165Filterer   // Log filterer for contract events
}

// IERC165Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC165Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC165Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC165Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC165Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC165Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC165Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC165Session struct {
	Contract     *IERC165          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC165CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC165CallerSession struct {
	Contract *IERC165Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// IERC165TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC165TransactorSession struct {
	Contract     *IERC165Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IERC165Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC165Raw struct {
	Contract *IERC165 // Generic contract binding to access the raw methods on
}

// IERC165CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC165CallerRaw struct {
	Contract *IERC165Caller // Generic read-only contract binding to access the raw methods on
}

// IERC165TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC165TransactorRaw struct {
	Contract *IERC165Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC165 creates a new instance of IERC165, bound to a specific deployed contract.
func NewIERC165(address common.Address, backend bind.ContractBackend) (*IERC165, error) {
	contract, err := bindIERC165(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC165{IERC165Caller: IERC165Caller{contract: contract}, IERC165Transactor: IERC165Transactor{contract: contract}, IERC165Filterer: IERC165Filterer{contract: contract}}, nil
}

// NewIERC165Caller creates a new read-only instance of IERC165, bound to a specific deployed contract.
func NewIERC165Caller(address common.Address, caller bind.ContractCaller) (*IERC165Caller, error) {
	contract, err := bindIERC165(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC165Caller{contract: contract}, nil
}

// NewIERC165Transactor creates a new write-only instance of IERC165, bound to a specific deployed contract.
func NewIERC165Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC165Transactor, error) {
	contract, err := bindIERC165(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC165Transactor{contract: contract}, nil
}

// NewIERC165Filterer creates a new log filterer instance of IERC165, bound to a specific deployed contract.
func NewIERC165Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC165Filterer, error) {
	contract, err := bindIERC165(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC165Filterer{contract: contract}, nil
}

// bindIERC165 binds a generic wrapper to an already deployed contract.
func bindIERC165(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC165ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC165 *IERC165Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC165.Contract.IERC165Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC165 *IERC165Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC165.Contract.IERC165Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC165 *IERC165Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC165.Contract.IERC165Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC165 *IERC165CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC165.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC165 *IERC165TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC165.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC165 *IERC165TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC165.Contract.contract.Transact(opts, method, params...)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC165 *IERC165Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _IERC165.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC165 *IERC165Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IERC165.Contract.SupportsInterface(&_IERC165.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC165 *IERC165CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IERC165.Contract.SupportsInterface(&_IERC165.CallOpts, interfaceId)
}

// IERC20ABI is the input ABI used to generate the binding from.
const IERC20ABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IERC20FuncSigs maps the 4-byte function signature to its string representation.
var IERC20FuncSigs = map[string]string{
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// IERC20 is an auto generated Go binding around an Ethereum contract.
type IERC20 struct {
	IERC20Caller     // Read-only binding to the contract
	IERC20Transactor // Write-only binding to the contract
	IERC20Filterer   // Log filterer for contract events
}

// IERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20Session struct {
	Contract     *IERC20           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20CallerSession struct {
	Contract *IERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20TransactorSession struct {
	Contract     *IERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20Raw struct {
	Contract *IERC20 // Generic contract binding to access the raw methods on
}

// IERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20CallerRaw struct {
	Contract *IERC20Caller // Generic read-only contract binding to access the raw methods on
}

// IERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20TransactorRaw struct {
	Contract *IERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20 creates a new instance of IERC20, bound to a specific deployed contract.
func NewIERC20(address common.Address, backend bind.ContractBackend) (*IERC20, error) {
	contract, err := bindIERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20{IERC20Caller: IERC20Caller{contract: contract}, IERC20Transactor: IERC20Transactor{contract: contract}, IERC20Filterer: IERC20Filterer{contract: contract}}, nil
}

// NewIERC20Caller creates a new read-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Caller(address common.Address, caller bind.ContractCaller) (*IERC20Caller, error) {
	contract, err := bindIERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Caller{contract: contract}, nil
}

// NewIERC20Transactor creates a new write-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC20Transactor, error) {
	contract, err := bindIERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Transactor{contract: contract}, nil
}

// NewIERC20Filterer creates a new log filterer instance of IERC20, bound to a specific deployed contract.
func NewIERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC20Filterer, error) {
	contract, err := bindIERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20Filterer{contract: contract}, nil
}

// bindIERC20 binds a generic wrapper to an already deployed contract.
func bindIERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.IERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Session) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, sender, recipient, amount)
}

// IERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC20 contract.
type IERC20ApprovalIterator struct {
	Event *IERC20Approval // Event containing the contract specifics and raw log

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
func (it *IERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Approval)
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
		it.Event = new(IERC20Approval)
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
func (it *IERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Approval represents a Approval event raised by the IERC20 contract.
type IERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IERC20ApprovalIterator{contract: _IERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Approval)
				if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) ParseApproval(log types.Log) (*IERC20Approval, error) {
	event := new(IERC20Approval)
	if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC20 contract.
type IERC20TransferIterator struct {
	Event *IERC20Transfer // Event containing the contract specifics and raw log

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
func (it *IERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Transfer)
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
		it.Event = new(IERC20Transfer)
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
func (it *IERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Transfer represents a Transfer event raised by the IERC20 contract.
type IERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IERC20TransferIterator{contract: _IERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Transfer)
				if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) ParseTransfer(log types.Log) (*IERC20Transfer, error) {
	event := new(IERC20Transfer)
	if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20MinterABI is the input ABI used to generate the binding from.
const IERC20MinterABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newMinter\",\"type\":\"address\"}],\"name\":\"replaceMinter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IERC20MinterFuncSigs maps the 4-byte function signature to its string representation.
var IERC20MinterFuncSigs = map[string]string{
	"42966c68": "burn(uint256)",
	"40c10f19": "mint(address,uint256)",
	"07f1af44": "replaceMinter(address)",
}

// IERC20Minter is an auto generated Go binding around an Ethereum contract.
type IERC20Minter struct {
	IERC20MinterCaller     // Read-only binding to the contract
	IERC20MinterTransactor // Write-only binding to the contract
	IERC20MinterFilterer   // Log filterer for contract events
}

// IERC20MinterCaller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20MinterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20MinterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20MinterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20MinterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20MinterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20MinterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20MinterSession struct {
	Contract     *IERC20Minter     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20MinterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20MinterCallerSession struct {
	Contract *IERC20MinterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// IERC20MinterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20MinterTransactorSession struct {
	Contract     *IERC20MinterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IERC20MinterRaw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20MinterRaw struct {
	Contract *IERC20Minter // Generic contract binding to access the raw methods on
}

// IERC20MinterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20MinterCallerRaw struct {
	Contract *IERC20MinterCaller // Generic read-only contract binding to access the raw methods on
}

// IERC20MinterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20MinterTransactorRaw struct {
	Contract *IERC20MinterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20Minter creates a new instance of IERC20Minter, bound to a specific deployed contract.
func NewIERC20Minter(address common.Address, backend bind.ContractBackend) (*IERC20Minter, error) {
	contract, err := bindIERC20Minter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20Minter{IERC20MinterCaller: IERC20MinterCaller{contract: contract}, IERC20MinterTransactor: IERC20MinterTransactor{contract: contract}, IERC20MinterFilterer: IERC20MinterFilterer{contract: contract}}, nil
}

// NewIERC20MinterCaller creates a new read-only instance of IERC20Minter, bound to a specific deployed contract.
func NewIERC20MinterCaller(address common.Address, caller bind.ContractCaller) (*IERC20MinterCaller, error) {
	contract, err := bindIERC20Minter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20MinterCaller{contract: contract}, nil
}

// NewIERC20MinterTransactor creates a new write-only instance of IERC20Minter, bound to a specific deployed contract.
func NewIERC20MinterTransactor(address common.Address, transactor bind.ContractTransactor) (*IERC20MinterTransactor, error) {
	contract, err := bindIERC20Minter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20MinterTransactor{contract: contract}, nil
}

// NewIERC20MinterFilterer creates a new log filterer instance of IERC20Minter, bound to a specific deployed contract.
func NewIERC20MinterFilterer(address common.Address, filterer bind.ContractFilterer) (*IERC20MinterFilterer, error) {
	contract, err := bindIERC20Minter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20MinterFilterer{contract: contract}, nil
}

// bindIERC20Minter binds a generic wrapper to an already deployed contract.
func bindIERC20Minter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC20MinterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20Minter *IERC20MinterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20Minter.Contract.IERC20MinterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20Minter *IERC20MinterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20Minter.Contract.IERC20MinterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20Minter *IERC20MinterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20Minter.Contract.IERC20MinterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20Minter *IERC20MinterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20Minter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20Minter *IERC20MinterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20Minter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20Minter *IERC20MinterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20Minter.Contract.contract.Transact(opts, method, params...)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_IERC20Minter *IERC20MinterTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Minter.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_IERC20Minter *IERC20MinterSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _IERC20Minter.Contract.Burn(&_IERC20Minter.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_IERC20Minter *IERC20MinterTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _IERC20Minter.Contract.Burn(&_IERC20Minter.TransactOpts, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_IERC20Minter *IERC20MinterTransactor) Mint(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Minter.contract.Transact(opts, "mint", to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_IERC20Minter *IERC20MinterSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Minter.Contract.Mint(&_IERC20Minter.TransactOpts, to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_IERC20Minter *IERC20MinterTransactorSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Minter.Contract.Mint(&_IERC20Minter.TransactOpts, to, amount)
}

// ReplaceMinter is a paid mutator transaction binding the contract method 0x07f1af44.
//
// Solidity: function replaceMinter(address newMinter) returns()
func (_IERC20Minter *IERC20MinterTransactor) ReplaceMinter(opts *bind.TransactOpts, newMinter common.Address) (*types.Transaction, error) {
	return _IERC20Minter.contract.Transact(opts, "replaceMinter", newMinter)
}

// ReplaceMinter is a paid mutator transaction binding the contract method 0x07f1af44.
//
// Solidity: function replaceMinter(address newMinter) returns()
func (_IERC20Minter *IERC20MinterSession) ReplaceMinter(newMinter common.Address) (*types.Transaction, error) {
	return _IERC20Minter.Contract.ReplaceMinter(&_IERC20Minter.TransactOpts, newMinter)
}

// ReplaceMinter is a paid mutator transaction binding the contract method 0x07f1af44.
//
// Solidity: function replaceMinter(address newMinter) returns()
func (_IERC20Minter *IERC20MinterTransactorSession) ReplaceMinter(newMinter common.Address) (*types.Transaction, error) {
	return _IERC20Minter.Contract.ReplaceMinter(&_IERC20Minter.TransactOpts, newMinter)
}

// IERC721ABI is the input ABI used to generate the binding from.
const IERC721ABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IERC721FuncSigs maps the 4-byte function signature to its string representation.
var IERC721FuncSigs = map[string]string{
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"081812fc": "getApproved(uint256)",
	"e985e9c5": "isApprovedForAll(address,address)",
	"6352211e": "ownerOf(uint256)",
	"42842e0e": "safeTransferFrom(address,address,uint256)",
	"b88d4fde": "safeTransferFrom(address,address,uint256,bytes)",
	"a22cb465": "setApprovalForAll(address,bool)",
	"01ffc9a7": "supportsInterface(bytes4)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// IERC721 is an auto generated Go binding around an Ethereum contract.
type IERC721 struct {
	IERC721Caller     // Read-only binding to the contract
	IERC721Transactor // Write-only binding to the contract
	IERC721Filterer   // Log filterer for contract events
}

// IERC721Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC721Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC721Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC721Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC721Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC721Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC721Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC721Session struct {
	Contract     *IERC721          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC721CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC721CallerSession struct {
	Contract *IERC721Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// IERC721TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC721TransactorSession struct {
	Contract     *IERC721Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IERC721Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC721Raw struct {
	Contract *IERC721 // Generic contract binding to access the raw methods on
}

// IERC721CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC721CallerRaw struct {
	Contract *IERC721Caller // Generic read-only contract binding to access the raw methods on
}

// IERC721TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC721TransactorRaw struct {
	Contract *IERC721Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC721 creates a new instance of IERC721, bound to a specific deployed contract.
func NewIERC721(address common.Address, backend bind.ContractBackend) (*IERC721, error) {
	contract, err := bindIERC721(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC721{IERC721Caller: IERC721Caller{contract: contract}, IERC721Transactor: IERC721Transactor{contract: contract}, IERC721Filterer: IERC721Filterer{contract: contract}}, nil
}

// NewIERC721Caller creates a new read-only instance of IERC721, bound to a specific deployed contract.
func NewIERC721Caller(address common.Address, caller bind.ContractCaller) (*IERC721Caller, error) {
	contract, err := bindIERC721(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC721Caller{contract: contract}, nil
}

// NewIERC721Transactor creates a new write-only instance of IERC721, bound to a specific deployed contract.
func NewIERC721Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC721Transactor, error) {
	contract, err := bindIERC721(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC721Transactor{contract: contract}, nil
}

// NewIERC721Filterer creates a new log filterer instance of IERC721, bound to a specific deployed contract.
func NewIERC721Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC721Filterer, error) {
	contract, err := bindIERC721(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC721Filterer{contract: contract}, nil
}

// bindIERC721 binds a generic wrapper to an already deployed contract.
func bindIERC721(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC721ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC721 *IERC721Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC721.Contract.IERC721Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC721 *IERC721Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC721.Contract.IERC721Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC721 *IERC721Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC721.Contract.IERC721Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC721 *IERC721CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC721.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC721 *IERC721TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC721.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC721 *IERC721TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC721.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256 balance)
func (_IERC721 *IERC721Caller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC721.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256 balance)
func (_IERC721 *IERC721Session) BalanceOf(owner common.Address) (*big.Int, error) {
	return _IERC721.Contract.BalanceOf(&_IERC721.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256 balance)
func (_IERC721 *IERC721CallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _IERC721.Contract.BalanceOf(&_IERC721.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address operator)
func (_IERC721 *IERC721Caller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _IERC721.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address operator)
func (_IERC721 *IERC721Session) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _IERC721.Contract.GetApproved(&_IERC721.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address operator)
func (_IERC721 *IERC721CallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _IERC721.Contract.GetApproved(&_IERC721.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_IERC721 *IERC721Caller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _IERC721.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_IERC721 *IERC721Session) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _IERC721.Contract.IsApprovedForAll(&_IERC721.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_IERC721 *IERC721CallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _IERC721.Contract.IsApprovedForAll(&_IERC721.CallOpts, owner, operator)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address owner)
func (_IERC721 *IERC721Caller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _IERC721.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address owner)
func (_IERC721 *IERC721Session) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _IERC721.Contract.OwnerOf(&_IERC721.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address owner)
func (_IERC721 *IERC721CallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _IERC721.Contract.OwnerOf(&_IERC721.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC721 *IERC721Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _IERC721.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC721 *IERC721Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IERC721.Contract.SupportsInterface(&_IERC721.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC721 *IERC721CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IERC721.Contract.SupportsInterface(&_IERC721.CallOpts, interfaceId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_IERC721 *IERC721Transactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_IERC721 *IERC721Session) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721.Contract.Approve(&_IERC721.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_IERC721 *IERC721TransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721.Contract.Approve(&_IERC721.TransactOpts, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721 *IERC721Transactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721 *IERC721Session) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721.Contract.SafeTransferFrom(&_IERC721.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721 *IERC721TransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721.Contract.SafeTransferFrom(&_IERC721.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_IERC721 *IERC721Transactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC721.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_IERC721 *IERC721Session) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC721.Contract.SafeTransferFrom0(&_IERC721.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_IERC721 *IERC721TransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC721.Contract.SafeTransferFrom0(&_IERC721.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool _approved) returns()
func (_IERC721 *IERC721Transactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, _approved bool) (*types.Transaction, error) {
	return _IERC721.contract.Transact(opts, "setApprovalForAll", operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool _approved) returns()
func (_IERC721 *IERC721Session) SetApprovalForAll(operator common.Address, _approved bool) (*types.Transaction, error) {
	return _IERC721.Contract.SetApprovalForAll(&_IERC721.TransactOpts, operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool _approved) returns()
func (_IERC721 *IERC721TransactorSession) SetApprovalForAll(operator common.Address, _approved bool) (*types.Transaction, error) {
	return _IERC721.Contract.SetApprovalForAll(&_IERC721.TransactOpts, operator, _approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721 *IERC721Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721 *IERC721Session) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721.Contract.TransferFrom(&_IERC721.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721 *IERC721TransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721.Contract.TransferFrom(&_IERC721.TransactOpts, from, to, tokenId)
}

// IERC721ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC721 contract.
type IERC721ApprovalIterator struct {
	Event *IERC721Approval // Event containing the contract specifics and raw log

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
func (it *IERC721ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC721Approval)
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
		it.Event = new(IERC721Approval)
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
func (it *IERC721ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC721ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC721Approval represents a Approval event raised by the IERC721 contract.
type IERC721Approval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_IERC721 *IERC721Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*IERC721ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _IERC721.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &IERC721ApprovalIterator{contract: _IERC721.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_IERC721 *IERC721Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC721Approval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _IERC721.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC721Approval)
				if err := _IERC721.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_IERC721 *IERC721Filterer) ParseApproval(log types.Log) (*IERC721Approval, error) {
	event := new(IERC721Approval)
	if err := _IERC721.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC721ApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the IERC721 contract.
type IERC721ApprovalForAllIterator struct {
	Event *IERC721ApprovalForAll // Event containing the contract specifics and raw log

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
func (it *IERC721ApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC721ApprovalForAll)
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
		it.Event = new(IERC721ApprovalForAll)
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
func (it *IERC721ApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC721ApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC721ApprovalForAll represents a ApprovalForAll event raised by the IERC721 contract.
type IERC721ApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_IERC721 *IERC721Filterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*IERC721ApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _IERC721.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &IERC721ApprovalForAllIterator{contract: _IERC721.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_IERC721 *IERC721Filterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *IERC721ApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _IERC721.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC721ApprovalForAll)
				if err := _IERC721.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_IERC721 *IERC721Filterer) ParseApprovalForAll(log types.Log) (*IERC721ApprovalForAll, error) {
	event := new(IERC721ApprovalForAll)
	if err := _IERC721.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC721TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC721 contract.
type IERC721TransferIterator struct {
	Event *IERC721Transfer // Event containing the contract specifics and raw log

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
func (it *IERC721TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC721Transfer)
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
		it.Event = new(IERC721Transfer)
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
func (it *IERC721TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC721TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC721Transfer represents a Transfer event raised by the IERC721 contract.
type IERC721Transfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_IERC721 *IERC721Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*IERC721TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _IERC721.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &IERC721TransferIterator{contract: _IERC721.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_IERC721 *IERC721Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC721Transfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _IERC721.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC721Transfer)
				if err := _IERC721.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_IERC721 *IERC721Filterer) ParseTransfer(log types.Log) (*IERC721Transfer, error) {
	event := new(IERC721Transfer)
	if err := _IERC721.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC721MinterABI is the input ABI used to generate the binding from.
const IERC721MinterABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newMinter\",\"type\":\"address\"}],\"name\":\"replaceMinter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IERC721MinterFuncSigs maps the 4-byte function signature to its string representation.
var IERC721MinterFuncSigs = map[string]string{
	"42966c68": "burn(uint256)",
	"d0def521": "mint(address,string)",
	"07f1af44": "replaceMinter(address)",
}

// IERC721Minter is an auto generated Go binding around an Ethereum contract.
type IERC721Minter struct {
	IERC721MinterCaller     // Read-only binding to the contract
	IERC721MinterTransactor // Write-only binding to the contract
	IERC721MinterFilterer   // Log filterer for contract events
}

// IERC721MinterCaller is an auto generated read-only Go binding around an Ethereum contract.
type IERC721MinterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC721MinterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC721MinterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC721MinterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC721MinterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC721MinterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC721MinterSession struct {
	Contract     *IERC721Minter    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC721MinterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC721MinterCallerSession struct {
	Contract *IERC721MinterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// IERC721MinterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC721MinterTransactorSession struct {
	Contract     *IERC721MinterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// IERC721MinterRaw is an auto generated low-level Go binding around an Ethereum contract.
type IERC721MinterRaw struct {
	Contract *IERC721Minter // Generic contract binding to access the raw methods on
}

// IERC721MinterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC721MinterCallerRaw struct {
	Contract *IERC721MinterCaller // Generic read-only contract binding to access the raw methods on
}

// IERC721MinterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC721MinterTransactorRaw struct {
	Contract *IERC721MinterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC721Minter creates a new instance of IERC721Minter, bound to a specific deployed contract.
func NewIERC721Minter(address common.Address, backend bind.ContractBackend) (*IERC721Minter, error) {
	contract, err := bindIERC721Minter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC721Minter{IERC721MinterCaller: IERC721MinterCaller{contract: contract}, IERC721MinterTransactor: IERC721MinterTransactor{contract: contract}, IERC721MinterFilterer: IERC721MinterFilterer{contract: contract}}, nil
}

// NewIERC721MinterCaller creates a new read-only instance of IERC721Minter, bound to a specific deployed contract.
func NewIERC721MinterCaller(address common.Address, caller bind.ContractCaller) (*IERC721MinterCaller, error) {
	contract, err := bindIERC721Minter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC721MinterCaller{contract: contract}, nil
}

// NewIERC721MinterTransactor creates a new write-only instance of IERC721Minter, bound to a specific deployed contract.
func NewIERC721MinterTransactor(address common.Address, transactor bind.ContractTransactor) (*IERC721MinterTransactor, error) {
	contract, err := bindIERC721Minter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC721MinterTransactor{contract: contract}, nil
}

// NewIERC721MinterFilterer creates a new log filterer instance of IERC721Minter, bound to a specific deployed contract.
func NewIERC721MinterFilterer(address common.Address, filterer bind.ContractFilterer) (*IERC721MinterFilterer, error) {
	contract, err := bindIERC721Minter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC721MinterFilterer{contract: contract}, nil
}

// bindIERC721Minter binds a generic wrapper to an already deployed contract.
func bindIERC721Minter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC721MinterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC721Minter *IERC721MinterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC721Minter.Contract.IERC721MinterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC721Minter *IERC721MinterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC721Minter.Contract.IERC721MinterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC721Minter *IERC721MinterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC721Minter.Contract.IERC721MinterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC721Minter *IERC721MinterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC721Minter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC721Minter *IERC721MinterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC721Minter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC721Minter *IERC721MinterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC721Minter.Contract.contract.Transact(opts, method, params...)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_IERC721Minter *IERC721MinterTransactor) Burn(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721Minter.contract.Transact(opts, "burn", tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_IERC721Minter *IERC721MinterSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721Minter.Contract.Burn(&_IERC721Minter.TransactOpts, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_IERC721Minter *IERC721MinterTransactorSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721Minter.Contract.Burn(&_IERC721Minter.TransactOpts, tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0xd0def521.
//
// Solidity: function mint(address to, string tokenURI) returns()
func (_IERC721Minter *IERC721MinterTransactor) Mint(opts *bind.TransactOpts, to common.Address, tokenURI string) (*types.Transaction, error) {
	return _IERC721Minter.contract.Transact(opts, "mint", to, tokenURI)
}

// Mint is a paid mutator transaction binding the contract method 0xd0def521.
//
// Solidity: function mint(address to, string tokenURI) returns()
func (_IERC721Minter *IERC721MinterSession) Mint(to common.Address, tokenURI string) (*types.Transaction, error) {
	return _IERC721Minter.Contract.Mint(&_IERC721Minter.TransactOpts, to, tokenURI)
}

// Mint is a paid mutator transaction binding the contract method 0xd0def521.
//
// Solidity: function mint(address to, string tokenURI) returns()
func (_IERC721Minter *IERC721MinterTransactorSession) Mint(to common.Address, tokenURI string) (*types.Transaction, error) {
	return _IERC721Minter.Contract.Mint(&_IERC721Minter.TransactOpts, to, tokenURI)
}

// ReplaceMinter is a paid mutator transaction binding the contract method 0x07f1af44.
//
// Solidity: function replaceMinter(address newMinter) returns()
func (_IERC721Minter *IERC721MinterTransactor) ReplaceMinter(opts *bind.TransactOpts, newMinter common.Address) (*types.Transaction, error) {
	return _IERC721Minter.contract.Transact(opts, "replaceMinter", newMinter)
}

// ReplaceMinter is a paid mutator transaction binding the contract method 0x07f1af44.
//
// Solidity: function replaceMinter(address newMinter) returns()
func (_IERC721Minter *IERC721MinterSession) ReplaceMinter(newMinter common.Address) (*types.Transaction, error) {
	return _IERC721Minter.Contract.ReplaceMinter(&_IERC721Minter.TransactOpts, newMinter)
}

// ReplaceMinter is a paid mutator transaction binding the contract method 0x07f1af44.
//
// Solidity: function replaceMinter(address newMinter) returns()
func (_IERC721Minter *IERC721MinterTransactorSession) ReplaceMinter(newMinter common.Address) (*types.Transaction, error) {
	return _IERC721Minter.Contract.ReplaceMinter(&_IERC721Minter.TransactOpts, newMinter)
}

// IERC721ReceiverABI is the input ABI used to generate the binding from.
const IERC721ReceiverABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IERC721ReceiverFuncSigs maps the 4-byte function signature to its string representation.
var IERC721ReceiverFuncSigs = map[string]string{
	"150b7a02": "onERC721Received(address,address,uint256,bytes)",
}

// IERC721Receiver is an auto generated Go binding around an Ethereum contract.
type IERC721Receiver struct {
	IERC721ReceiverCaller     // Read-only binding to the contract
	IERC721ReceiverTransactor // Write-only binding to the contract
	IERC721ReceiverFilterer   // Log filterer for contract events
}

// IERC721ReceiverCaller is an auto generated read-only Go binding around an Ethereum contract.
type IERC721ReceiverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC721ReceiverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC721ReceiverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC721ReceiverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC721ReceiverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC721ReceiverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC721ReceiverSession struct {
	Contract     *IERC721Receiver  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC721ReceiverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC721ReceiverCallerSession struct {
	Contract *IERC721ReceiverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// IERC721ReceiverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC721ReceiverTransactorSession struct {
	Contract     *IERC721ReceiverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// IERC721ReceiverRaw is an auto generated low-level Go binding around an Ethereum contract.
type IERC721ReceiverRaw struct {
	Contract *IERC721Receiver // Generic contract binding to access the raw methods on
}

// IERC721ReceiverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC721ReceiverCallerRaw struct {
	Contract *IERC721ReceiverCaller // Generic read-only contract binding to access the raw methods on
}

// IERC721ReceiverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC721ReceiverTransactorRaw struct {
	Contract *IERC721ReceiverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC721Receiver creates a new instance of IERC721Receiver, bound to a specific deployed contract.
func NewIERC721Receiver(address common.Address, backend bind.ContractBackend) (*IERC721Receiver, error) {
	contract, err := bindIERC721Receiver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC721Receiver{IERC721ReceiverCaller: IERC721ReceiverCaller{contract: contract}, IERC721ReceiverTransactor: IERC721ReceiverTransactor{contract: contract}, IERC721ReceiverFilterer: IERC721ReceiverFilterer{contract: contract}}, nil
}

// NewIERC721ReceiverCaller creates a new read-only instance of IERC721Receiver, bound to a specific deployed contract.
func NewIERC721ReceiverCaller(address common.Address, caller bind.ContractCaller) (*IERC721ReceiverCaller, error) {
	contract, err := bindIERC721Receiver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC721ReceiverCaller{contract: contract}, nil
}

// NewIERC721ReceiverTransactor creates a new write-only instance of IERC721Receiver, bound to a specific deployed contract.
func NewIERC721ReceiverTransactor(address common.Address, transactor bind.ContractTransactor) (*IERC721ReceiverTransactor, error) {
	contract, err := bindIERC721Receiver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC721ReceiverTransactor{contract: contract}, nil
}

// NewIERC721ReceiverFilterer creates a new log filterer instance of IERC721Receiver, bound to a specific deployed contract.
func NewIERC721ReceiverFilterer(address common.Address, filterer bind.ContractFilterer) (*IERC721ReceiverFilterer, error) {
	contract, err := bindIERC721Receiver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC721ReceiverFilterer{contract: contract}, nil
}

// bindIERC721Receiver binds a generic wrapper to an already deployed contract.
func bindIERC721Receiver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC721ReceiverABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC721Receiver *IERC721ReceiverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC721Receiver.Contract.IERC721ReceiverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC721Receiver *IERC721ReceiverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC721Receiver.Contract.IERC721ReceiverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC721Receiver *IERC721ReceiverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC721Receiver.Contract.IERC721ReceiverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC721Receiver *IERC721ReceiverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC721Receiver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC721Receiver *IERC721ReceiverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC721Receiver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC721Receiver *IERC721ReceiverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC721Receiver.Contract.contract.Transact(opts, method, params...)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) returns(bytes4)
func (_IERC721Receiver *IERC721ReceiverTransactor) OnERC721Received(opts *bind.TransactOpts, operator common.Address, from common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC721Receiver.contract.Transact(opts, "onERC721Received", operator, from, tokenId, data)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) returns(bytes4)
func (_IERC721Receiver *IERC721ReceiverSession) OnERC721Received(operator common.Address, from common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC721Receiver.Contract.OnERC721Received(&_IERC721Receiver.TransactOpts, operator, from, tokenId, data)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) returns(bytes4)
func (_IERC721Receiver *IERC721ReceiverTransactorSession) OnERC721Received(operator common.Address, from common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC721Receiver.Contract.OnERC721Received(&_IERC721Receiver.TransactOpts, operator, from, tokenId, data)
}

// SafeERC20ABI is the input ABI used to generate the binding from.
const SafeERC20ABI = "[]"

// SafeERC20Bin is the compiled bytecode used for deploying new contracts.
var SafeERC20Bin = "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220ee4a73ae75c03460fafd50254e2613850e729153ea6cca65acf9a65fd1cd192764736f6c63430008000033"

// DeploySafeERC20 deploys a new Ethereum contract, binding an instance of SafeERC20 to it.
func DeploySafeERC20(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeERC20, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeERC20ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeERC20Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeERC20{SafeERC20Caller: SafeERC20Caller{contract: contract}, SafeERC20Transactor: SafeERC20Transactor{contract: contract}, SafeERC20Filterer: SafeERC20Filterer{contract: contract}}, nil
}

// SafeERC20 is an auto generated Go binding around an Ethereum contract.
type SafeERC20 struct {
	SafeERC20Caller     // Read-only binding to the contract
	SafeERC20Transactor // Write-only binding to the contract
	SafeERC20Filterer   // Log filterer for contract events
}

// SafeERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type SafeERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeERC20Session struct {
	Contract     *SafeERC20        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeERC20CallerSession struct {
	Contract *SafeERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// SafeERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeERC20TransactorSession struct {
	Contract     *SafeERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// SafeERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type SafeERC20Raw struct {
	Contract *SafeERC20 // Generic contract binding to access the raw methods on
}

// SafeERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeERC20CallerRaw struct {
	Contract *SafeERC20Caller // Generic read-only contract binding to access the raw methods on
}

// SafeERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeERC20TransactorRaw struct {
	Contract *SafeERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeERC20 creates a new instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20(address common.Address, backend bind.ContractBackend) (*SafeERC20, error) {
	contract, err := bindSafeERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeERC20{SafeERC20Caller: SafeERC20Caller{contract: contract}, SafeERC20Transactor: SafeERC20Transactor{contract: contract}, SafeERC20Filterer: SafeERC20Filterer{contract: contract}}, nil
}

// NewSafeERC20Caller creates a new read-only instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20Caller(address common.Address, caller bind.ContractCaller) (*SafeERC20Caller, error) {
	contract, err := bindSafeERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeERC20Caller{contract: contract}, nil
}

// NewSafeERC20Transactor creates a new write-only instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*SafeERC20Transactor, error) {
	contract, err := bindSafeERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeERC20Transactor{contract: contract}, nil
}

// NewSafeERC20Filterer creates a new log filterer instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*SafeERC20Filterer, error) {
	contract, err := bindSafeERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeERC20Filterer{contract: contract}, nil
}

// bindSafeERC20 binds a generic wrapper to an already deployed contract.
func bindSafeERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeERC20 *SafeERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeERC20.Contract.SafeERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeERC20 *SafeERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeERC20.Contract.SafeERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeERC20 *SafeERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeERC20.Contract.SafeERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeERC20 *SafeERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeERC20 *SafeERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeERC20 *SafeERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeERC20.Contract.contract.Transact(opts, method, params...)
}

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
var SafeMathBin = "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220206bf8204981c8e5d105650dfef004a61bf28c44fdc2c0709ccc6c0180b98ddd64736f6c63430008000033"

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}
