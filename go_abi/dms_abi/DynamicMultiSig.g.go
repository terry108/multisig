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
var AddressBin = "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122047faa52051c0b8541707bb3eebe07317de5cde29bca1c98bf5fc016de74f666864736f6c63430008000033"

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
var BytesLibBin = "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122075f1b525fc8adabd47a241a876ad48f9a5e6eb2a3111c448eaf279f2b5a0ad1d64736f6c63430008000033"

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
const DynamicMultiSigABI = "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_managers\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"to\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"ERC20\",\"type\":\"address\"}],\"name\":\"CrossOutFunds\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DepositFunds\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TransferFunds\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"txKey\",\"type\":\"string\"}],\"name\":\"TxManagerChangeCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"txKey\",\"type\":\"string\"}],\"name\":\"TxUpgradeCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"txKey\",\"type\":\"string\"}],\"name\":\"TxWithdrawCompleted\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"allManagers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"closeUpgrade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"txKey\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"adds\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"removes\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"count\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"signatures\",\"type\":\"bytes\"}],\"name\":\"createOrSignManagerChange\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"txKey\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"upgradeContract\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signatures\",\"type\":\"bytes\"}],\"name\":\"createOrSignUpgrade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"txKey\",\"type\":\"string\"},{\"internalType\":\"addresspayable\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isERC20\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"ERC20\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signatures\",\"type\":\"bytes\"}],\"name\":\"createOrSignWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"to\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"ERC20\",\"type\":\"address\"}],\"name\":\"crossOut\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"to\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"ERC721\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"crossOutNFT\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"current_min_signatures\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_manager\",\"type\":\"address\"}],\"name\":\"ifManager\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"txKey\",\"type\":\"string\"}],\"name\":\"isCompletedTx\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ERC20\",\"type\":\"address\"}],\"name\":\"isMinterERC20\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ERC721\",\"type\":\"address\"}],\"name\":\"isMinterERC721\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"max_managers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"min_managers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ERC20\",\"type\":\"address\"}],\"name\":\"registerMinterERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ERC721\",\"type\":\"address\"}],\"name\":\"registerMinterERC721\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"signatureLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ERC20\",\"type\":\"address\"}],\"name\":\"unregisterMinterERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ERC721\",\"type\":\"address\"}],\"name\":\"unregisterMinterERC721\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"upgrade\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"upgradeContractAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"upgradeContractS1\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ERC20\",\"type\":\"address\"}],\"name\":\"upgradeContractS2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// DynamicMultiSigFuncSigs maps the 4-byte function signature to its string representation.
var DynamicMultiSigFuncSigs = map[string]string{
	"9c30b35e": "allManagers()",
	"d4cacbaa": "closeUpgrade()",
	"00719226": "createOrSignManagerChange(string,address[],address[],uint8,bytes)",
	"408e8b7a": "createOrSignUpgrade(string,address,bytes)",
	"ab6c2b10": "createOrSignWithdraw(string,address,uint256,bool,address,bytes)",
	"0889d1f0": "crossOut(string,uint256,address)",
	"c76bbdc5": "crossOutNFT(string,address,uint256)",
	"e079cee9": "current_min_signatures()",
	"75173b70": "ifManager(address)",
	"a5e399b3": "isCompletedTx(string)",
	"6a7142e1": "isMinterERC20(address)",
	"279fb6e0": "isMinterERC721(address)",
	"f7f2ff74": "max_managers()",
	"ad4b61a8": "min_managers()",
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
var DynamicMultiSigBin = "0x6080604052600080546001600160a81b0319169055600f600155600360028190556042905560416004553480156200003657600080fd5b5060405162003dbe38038062003dbe833981016040819052620000599162000391565b60015481511115620000885760405162461bcd60e51b81526004016200007f90620004dd565b60405180910390fd5b60025481511015620000ae5760405162461bcd60e51b81526004016200007f906200045f565b60058054610100600160a81b03191633610100021790558051620000da906009906020840190620002f8565b5060005b60095460ff82161015620002215760016008600060098460ff16815481106200011757634e487b7160e01b600052603260045260246000fd5b6000918252602080832091909101546001600160a01b031683528201929092526040018120805460ff191660ff9384161790556009805460019360069392919086169081106200017757634e487b7160e01b600052603260045260246000fd5b6000918252602080832091909101546001600160a01b031683528201929092526040019020805460ff191660ff928316179055600980546007928416908110620001d157634e487b7160e01b600052603260045260246000fd5b60009182526020808320909101548354600181018555938352912090910180546001600160a01b0319166001600160a01b03909216919091179055806200021881620005e2565b915050620000de565b5060055461010090046001600160a01b031660009081526008602052604090205460ff1615620002655760405162461bcd60e51b81526004016200007f9062000524565b60095462000273906200028e565b6005805460ff191660ff929092169190911790555062000631565b6000808211620002b25760405162461bcd60e51b81526004016200007f90620004a6565b60006001606484600354620002c89190620005a6565b620002d491906200056a565b620002e09190620005c8565b9050620002ef60648262000585565b9150505b919050565b82805482825590600052602060002090810192821562000350579160200282015b828111156200035057825182546001600160a01b0319166001600160a01b0390911617825560209092019160019091019062000319565b506200035e92915062000362565b5090565b5b808211156200035e576000815560010162000363565b80516001600160a01b0381168114620002f357600080fd5b60006020808385031215620003a4578182fd5b82516001600160401b0380821115620003bb578384fd5b818501915085601f830112620003cf578384fd5b815181811115620003e457620003e46200061b565b838102604051858282010181811085821117156200040657620004066200061b565b604052828152858101935084860182860187018a101562000425578788fd5b8795505b8386101562000452576200043d8162000379565b85526001959095019493860193860162000429565b5098975050505050505050565b60208082526027908201527f4e6f74207265616368696e6720746865206d696e206e756d626572206f66206d604082015266616e616765727360c81b606082015260800190565b60208082526014908201527f4d616e616765722043616e277420656d7074792e000000000000000000000000604082015260600190565b60208082526027908201527f457863656564656420746865206d6178696d756d206e756d626572206f66206d604082015266616e616765727360c81b606082015260800190565b60208082526026908201527f436f6e74726163742063726561746f722063616e6e6f7420616374206173206d60408201526530b730b3b2b960d11b606082015260800190565b6000821982111562000580576200058062000605565b500190565b600082620005a157634e487b7160e01b81526012600452602481fd5b500490565b6000816000190483118215151615620005c357620005c362000605565b500290565b600082821015620005dd57620005dd62000605565b500390565b600060ff821660ff811415620005fc57620005fc62000605565b60010192915050565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052604160045260246000fd5b61377d80620006416000396000f3fe6080604052600436106101695760003560e01c80637de24876116100d1578063ad4b61a81161008a578063d4cacbaa11610064578063d4cacbaa1461042e578063d55ec69714610443578063e079cee914610458578063f7f2ff741461047a576101a9565b8063ad4b61a8146103d9578063bc8870ed146103ee578063c76bbdc51461040e576101a9565b80637de24876146103225780638da5cb5b146103425780639c30b35e146103575780639dcdc97814610379578063a5e399b314610399578063ab6c2b10146103b9576101a9565b806330b2d84d1161012357806330b2d84d1461026057806334774b7114610282578063408e8b7a146102a25780635bda3fcf146102c25780636a7142e1146102e257806375173b7014610302576101a9565b8062719226146101ab5780630889d1f0146101cb5780631b9a9323146101f45780631dda9c0514610216578063279fb6e01461022b5780632c4e722e1461024b576101a9565b366101a9577fd241e73300212f6df233a8e6d3146b88a9d4964e06621d54b5ff6afeba7b1b88333460405161019f929190612a4a565b60405180910390a1005b005b3480156101b757600080fd5b506101a96101c6366004612878565b61048f565b6101de6101d936600461293a565b610669565b6040516101eb9190612aef565b60405180910390f35b34801561020057600080fd5b50610209610920565b6040516101eb91906135bd565b34801561022257600080fd5b506101a9610926565b34801561023757600080fd5b506101de610246366004612699565b6109e8565b34801561025757600080fd5b50610209610a0c565b34801561026c57600080fd5b50610275610a12565b6040516101eb9190612a36565b34801561028e57600080fd5b506101a961029d366004612699565b610a26565b3480156102ae57600080fd5b506101a96102bd3660046127ab565b610af6565b3480156102ce57600080fd5b506101a96102dd366004612699565b610cd5565b3480156102ee57600080fd5b506101de6102fd366004612699565b610ecb565b34801561030e57600080fd5b506101de61031d366004612699565b610eeb565b34801561032e57600080fd5b506101a961033d366004612699565b610f0c565b34801561034e57600080fd5b50610275610fdc565b34801561036357600080fd5b5061036c610ff0565b6040516101eb9190612adc565b34801561038557600080fd5b506101a9610394366004612699565b611052565b3480156103a557600080fd5b506101de6103b43660046126d1565b6110c7565b3480156103c557600080fd5b506101a96103d4366004612704565b6110f5565b3480156103e557600080fd5b50610209611383565b3480156103fa57600080fd5b506101a9610409366004612699565b611389565b34801561041a57600080fd5b506101de610429366004612820565b6113fe565b34801561043a57600080fd5b506101a961156c565b34801561044f57600080fd5b506101de6115c9565b34801561046457600080fd5b5061046d6115d2565b6040516101eb91906135c6565b34801561048657600080fd5b506102096115db565b3360009081526008602052604090205460ff166001146104ca5760405162461bcd60e51b81526004016104c190612eba565b60405180910390fd5b84516040146104eb5760405162461bcd60e51b81526004016104c190613395565b6000845111806104fc575060008351115b6105185760405162461bcd60e51b81526004016104c1906132e8565b600b856040516105289190612a1a565b9081526040519081900360200190205460ff16156105585760405162461bcd60e51b81526004016104c190612d9a565b61056284846115e1565b600085858486600260405160200161057e959493929190612ba0565b60408051601f1981840301815291815281516020928301206000818152600a90935291205490915060ff16156105c65760405162461bcd60e51b81526004016104c190612cc6565b6105d08183611824565b6105ec5760405162461bcd60e51b81526004016104c1906132b9565b6105f584611867565b6105fe85611a62565b60095461060a90611b48565b6005805460ff191660ff9290921691909117905561062a86826001611b9e565b7fac9b82db4e104d515319a481096bfd91a4f40ee10837d5a2c8d51b9a03dc48ae866040516106599190612b18565b60405180910390a1505050505050565b600033836106895760405162461bcd60e51b81526004016104c190612ff9565b6001600160a01b038316156108b75734156106b65760405162461bcd60e51b81526004016104c190612e4e565b6106c8836001600160a01b0316611beb565b6106e45760405162461bcd60e51b81526004016104c190613578565b604051636eb1769f60e11b815283906000906001600160a01b0383169063dd62ed3e906107179086903090600401612a63565b60206040518083038186803b15801561072f57600080fd5b505afa158015610743573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107679190612993565b9050858110156107895760405162461bcd60e51b81526004016104c190613168565b6040516370a0823160e01b81526000906001600160a01b038416906370a08231906107b8908790600401612a36565b60206040518083038186803b1580156107d057600080fd5b505afa1580156107e4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108089190612993565b90508681101561082a5760405162461bcd60e51b81526004016104c190613282565b61083f6001600160a01b03841685308a611bf1565b61084886610ecb565b156108af57604051630852cd8d60e31b815286906001600160a01b038216906342966c689061087b908b906004016135bd565b600060405180830381600087803b15801561089557600080fd5b505af11580156108a9573d6000803e3d6000fd5b50505050505b5050506108d6565b8334146108d65760405162461bcd60e51b81526004016104c190613330565b7f5ddf9724d8fe5d9e12499be2867f93d41a582733dcd65f74a486ad7e306671468186868660405161090b9493929190612aa1565b60405180910390a160019150505b9392505050565b60045481565b60055461010090046001600160a01b031633146109555760405162461bcd60e51b81526004016104c19061342e565b60005460ff166109775760405162461bcd60e51b81526004016104c190613096565b60005461010090046001600160a01b03166109a45760405162461bcd60e51b81526004016104c1906131aa565b600080546040516001600160a01b0361010090920491909116914780156108fc02929091818181858888f193505050501580156109e5573d6000803e3d6000fd5b50565b6001600160a01b0381166000908152600d602052604090205460ff1615155b919050565b60035481565b60005461010090046001600160a01b031681565b60055461010090046001600160a01b03163314610a555760405162461bcd60e51b81526004016104c19061342e565b306001600160a01b0382161415610a7e5760405162461bcd60e51b81526004016104c190612fc9565b610a90816001600160a01b0316611beb565b610aac5760405162461bcd60e51b81526004016104c190613578565b610ab581610ecb565b15610ad25760405162461bcd60e51b81526004016104c190612c38565b6001600160a01b03166000908152600c60205260409020805460ff19166001179055565b3360009081526008602052604090205460ff16600114610b285760405162461bcd60e51b81526004016104c190612eba565b8251604014610b495760405162461bcd60e51b81526004016104c190613395565b600b83604051610b599190612a1a565b9081526040519081900360200190205460ff1615610b895760405162461bcd60e51b81526004016104c190612d9a565b60005460ff1615610bac5760405162461bcd60e51b81526004016104c190613367565b610bbe826001600160a01b0316611beb565b610bda5760405162461bcd60e51b81526004016104c190613578565b600083836002604051602001610bf293929190612b72565b60408051601f1981840301815291815281516020928301206000818152600a90935291205490915060ff1615610c3a5760405162461bcd60e51b81526004016104c190612cc6565b610c448183611824565b610c605760405162461bcd60e51b81526004016104c1906132b9565b60008054600160ff199091168117610100600160a81b0319166101006001600160a01b0387160217909155610c989085908390611b9e565b7f5e06c4b22547d430736ce834764dbfee08f1c4cf7ae3d53178aa56effa593ed084604051610cc79190612b18565b60405180910390a150505050565b60055461010090046001600160a01b03163314610d045760405162461bcd60e51b81526004016104c19061342e565b60005460ff16610d265760405162461bcd60e51b81526004016104c190613096565b60005461010090046001600160a01b0316610d535760405162461bcd60e51b81526004016104c1906131aa565b306001600160a01b0382161415610d7c5760405162461bcd60e51b81526004016104c190612fc9565b610d8e816001600160a01b0316611beb565b610daa5760405162461bcd60e51b81526004016104c190613578565b6040516370a0823160e01b815281906000906001600160a01b038316906370a0823190610ddb903090600401612a36565b60206040518083038186803b158015610df357600080fd5b505afa158015610e07573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e2b9190612993565b9050600054610e4c906001600160a01b038481169161010090041683611c49565b610e5583610ecb565b15610ec6576000546040516301fc6bd160e21b815284916001600160a01b03808416926307f1af4492610e92926101009091041690600401612a36565b600060405180830381600087803b158015610eac57600080fd5b505af1158015610ec0573d6000803e3d6000fd5b50505050505b505050565b6001600160a01b03166000908152600c602052604090205460ff16151590565b6001600160a01b031660009081526008602052604090205460ff1660011490565b60055461010090046001600160a01b03163314610f3b5760405162461bcd60e51b81526004016104c19061342e565b306001600160a01b0382161415610f645760405162461bcd60e51b81526004016104c190612fc9565b610f76816001600160a01b0316611beb565b610f925760405162461bcd60e51b81526004016104c190613578565b610f9b816109e8565b15610fb85760405162461bcd60e51b81526004016104c190612c38565b6001600160a01b03166000908152600d60205260409020805460ff19166001179055565b60055461010090046001600160a01b031681565b6060600980548060200260200160405190810160405280929190818152602001828054801561104857602002820191906000526020600020905b81546001600160a01b0316815260019091019060200180831161102a575b5050505050905090565b60055461010090046001600160a01b031633146110815760405162461bcd60e51b81526004016104c19061342e565b61108a81610ecb565b6110a65760405162461bcd60e51b81526004016104c19061324b565b6001600160a01b03166000908152600c60205260409020805460ff19169055565b600080600b836040516110da9190612a1a565b9081526040519081900360200190205460ff16119050919050565b3360009081526008602052604090205460ff166001146111275760405162461bcd60e51b81526004016104c190612eba565b85516040146111485760405162461bcd60e51b81526004016104c190613395565b6001600160a01b03851661116e5760405162461bcd60e51b81526004016104c190612c80565b6000841161118e5760405162461bcd60e51b81526004016104c190613120565b600b8660405161119e9190612a1a565b9081526040519081900360200190205460ff16156111ce5760405162461bcd60e51b81526004016104c190612d9a565b82156111e4576111df828686611c68565b611204565b834710156112045760405162461bcd60e51b81526004016104c190612ef1565b60008686868686600260405160200161122296959493929190612b2b565b60408051601f1981840301815291815281516020928301206000818152600a90935291205490915060ff161561126a5760405162461bcd60e51b81526004016104c190612cc6565b6112748183611824565b6112905760405162461bcd60e51b81526004016104c1906132b9565b83156112a6576112a1838787611da2565b611337565b844710156112c65760405162461bcd60e51b81526004016104c190612ef1565b6040516001600160a01b0387169086156108fc029087906000818181858888f193505050501580156112fc573d6000803e3d6000fd5b507fc95f8b91b103304386b955ef73fadac189f8ad66b33369b6c34a17a60db7bd0a868660405161132e929190612a4a565b60405180910390a15b61134387826001611b9e565b7f8ed8b1f0dd3babfdf1477ba2b27a5b0d2f1c9148448fd22cf2c75e658293c7b1876040516113729190612b18565b60405180910390a150505050505050565b60025481565b60055461010090046001600160a01b031633146113b85760405162461bcd60e51b81526004016104c19061342e565b6113c1816109e8565b6113dd5760405162461bcd60e51b81526004016104c19061324b565b6001600160a01b03166000908152600d60205260409020805460ff19169055565b6000336114136001600160a01b038516611beb565b61142f5760405162461bcd60e51b81526004016104c190613578565b6000831161144f5760405162461bcd60e51b81526004016104c190612ff9565b604051632142170760e11b815284906001600160a01b038216906342842e0e9061148190859030908990600401612a7d565b600060405180830381600087803b15801561149b57600080fd5b505af11580156114af573d6000803e3d6000fd5b505050506114bc856109e8565b1561152357604051630852cd8d60e31b815285906001600160a01b038216906342966c68906114ef9088906004016135bd565b600060405180830381600087803b15801561150957600080fd5b505af115801561151d573d6000803e3d6000fd5b50505050505b7f5ddf9724d8fe5d9e12499be2867f93d41a582733dcd65f74a486ad7e30667146828786886040516115589493929190612aa1565b60405180910390a150600195945050505050565b60055461010090046001600160a01b0316331461159b5760405162461bcd60e51b81526004016104c19061342e565b60005460ff166115bd5760405162461bcd60e51b81526004016104c190613096565b6000805460ff19169055565b60005460ff1681565b60055460ff1681565b60015481565b815160005b8181101561169957600084828151811061161057634e487b7160e01b600052603260045260246000fd5b6020026020010151905060006001600160a01b0316816001600160a01b0316141561164d5760405162461bcd60e51b81526004016104c190612f85565b6001600160a01b03811660009081526008602052604090205460ff16156116865760405162461bcd60e51b81526004016104c1906131ed565b5080611691816136bd565b9150506115e6565b506116a383611ecf565b6116bf5760405162461bcd60e51b81526004016104c190612cf2565b6005546116da9061010090046001600160a01b031684611fe7565b6116f65760405162461bcd60e51b81526004016104c190613532565b6116ff82611ecf565b61171b5760405162461bcd60e51b81526004016104c190613465565b815160005b818110156117e157600084828151811061174a57634e487b7160e01b600052603260045260246000fd5b6020908102919091018101516001600160a01b0381166000908152600690925260409091205490915060ff16156117935760405162461bcd60e51b81526004016104c1906134b1565b6001600160a01b03811660009081526008602052604090205460ff166001146117ce5760405162461bcd60e51b81526004016104c1906130b6565b50806117d9816136bd565b915050611720565b50600154835185516009546117f691906135fe565b611800919061367a565b111561181e5760405162461bcd60e51b81526004016104c19061304f565b50505050565b60006103cf825111156118495760405162461bcd60e51b81526004016104c190612f4e565b60006118558484612082565b60055460ff1611159150505b92915050565b8051611872576109e5565b60005b81518110156118df57600860008383815181106118a257634e487b7160e01b600052603260045260246000fd5b6020908102919091018101516001600160a01b03168252810191909152604001600020805460ff19169055806118d7816136bd565b915050611875565b5060005b60095481101561198b57600860006009838154811061191257634e487b7160e01b600052603260045260246000fd5b60009182526020808320909101546001600160a01b0316835282019290925260400190205460ff16611979576009818154811061195f57634e487b7160e01b600052603260045260246000fd5b600091825260209091200180546001600160a01b03191690555b80611983816136bd565b9150506118e3565b50601060005b600954811015610ec6576000600982815481106119be57634e487b7160e01b600052603260045260246000fd5b6000918252602090912001546001600160a01b03169050806119ed5782601014156119e7578192505b50611a50565b82601014611a4e578060098481548110611a1757634e487b7160e01b600052603260045260246000fd5b600091825260209091200180546001600160a01b0319166001600160a01b039290921691909117905582611a4a816136bd565b9350505b505b80611a5a816136bd565b915050611991565b8051611a6d576109e5565b60005b8151811015611b44576000828281518110611a9b57634e487b7160e01b600052603260045260246000fd5b6020908102919091018101516001600160a01b0381166000908152600890925260409091205490915060ff16611b31576001600160a01b0381166000818152600860205260408120805460ff191660019081179091556009805491820181559091527f6e1540171b6c0c960b71a7020d9f60077f6af931a8bbf590da0223dacf75c7af0180546001600160a01b03191690911790555b5080611b3c816136bd565b915050611a70565b5050565b6000808211611b695760405162461bcd60e51b81526004016104c190612d3e565b60006001606484600354611b7d919061365b565b611b8791906135fe565b611b91919061367a565b905061091960648261363b565b80600b84604051611baf9190612a1a565b9081526040805160209281900383019020805460ff1990811660ff958616179091556000958652600a9092529093208054909316911617905550565b3b151590565b61181e846323b872dd60e01b858585604051602401611c1293929190612a7d565b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b03199093169290921790915261222e565b610ec68363a9059cbb60e01b8484604051602401611c12929190612a4a565b6001600160a01b038216611c8e5760405162461bcd60e51b81526004016104c190612bf5565b306001600160a01b0384161415611cb75760405162461bcd60e51b81526004016104c190612fc9565b611cc9836001600160a01b0316611beb565b611ce55760405162461bcd60e51b81526004016104c190613578565b611cee83610ecb565b15611cf857610ec6565b6040516370a0823160e01b815283906000906001600160a01b038316906370a0823190611d29903090600401612a36565b60206040518083038186803b158015611d4157600080fd5b505afa158015611d55573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611d799190612993565b905082811015611d9b5760405162461bcd60e51b81526004016104c190612dd1565b5050505050565b611dab83610ecb565b15611e18576040516340c10f1960e01b815283906001600160a01b038216906340c10f1990611de09086908690600401612a4a565b600060405180830381600087803b158015611dfa57600080fd5b505af1158015611e0e573d6000803e3d6000fd5b5050505050610ec6565b6040516370a0823160e01b815283906000906001600160a01b038316906370a0823190611e49903090600401612a36565b60206040518083038186803b158015611e6157600080fd5b505afa158015611e75573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611e999190612993565b905082811015611ebb5760405162461bcd60e51b81526004016104c190612dd1565b611d9b6001600160a01b0383168585611c49565b6000805b8251811015611fde576000838281518110611efe57634e487b7160e01b600052603260045260246000fd5b6020026020010151905060006001600160a01b0316816001600160a01b03161415611f295750611fde565b6000611f368360016135fe565b90505b8451811015611fc9576000858281518110611f6457634e487b7160e01b600052603260045260246000fd5b6020026020010151905060006001600160a01b0316816001600160a01b03161415611f8f5750611fc9565b806001600160a01b0316836001600160a01b03161415611fb6576000945050505050610a07565b5080611fc1816136bd565b915050611f39565b50508080611fd6906136bd565b915050611ed3565b50600192915050565b60008060005b83518110156120775783818151811061201657634e487b7160e01b600052603260045260246000fd5b6020026020010151915060006001600160a01b0316826001600160a01b0316141561204057612077565b846001600160a01b0316826001600160a01b0316141561206557600092505050611861565b8061206f816136bd565b915050611fed565b506001949350505050565b60045481516000918291829161209891906122bd565b905060008167ffffffffffffffff8111156120c357634e487b7160e01b600052604160045260246000fd5b6040519080825280602002602001820160405280156120ec578160200160208202803683370190505b50905060008060005b848110156121f2576000612116846004548b6122c99092919063ffffffff16565b905060006121248b83612391565b90506001600160a01b03811661214c5760405162461bcd60e51b81526004016104c190613025565b6001600160a01b03811660009081526008602052604090205460ff16600114156121ce578761217a816136bd565b9850508086858061218a906136d8565b965060ff16815181106121ad57634e487b7160e01b600052603260045260246000fd5b60200260200101906001600160a01b031690816001600160a01b0316815250505b6004546121db90866135fe565b9450505080806121ea906136bd565b9150506120f5565b5060006121fe84611ecf565b905060609350806122215760405162461bcd60e51b81526004016104c190612d6c565b5093979650505050505050565b6000612283826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166124939092919063ffffffff16565b805190915015610ec657808060200190518101906122a191906126b5565b610ec65760405162461bcd60e51b81526004016104c1906134e8565b6000610919828461363b565b6060816122d781601f6135fe565b10156122f55760405162461bcd60e51b81526004016104c190612e92565b6122ff82846135fe565b8451101561231f5760405162461bcd60e51b81526004016104c190613403565b60608215801561233e5760405191506000825260208201604052612388565b6040519150601f8416801560200281840101858101878315602002848b0101015b8183101561237757805183526020928301920161235f565b5050858452601f01601f1916604052505b50949350505050565b6000806000806004548551146123ad5760009350505050611861565b50505060208201516040830151606084015160001a7f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08211156123f65760009350505050611861565b601b8160ff1610156124105761240d601b82613616565b90505b8060ff16601b1415801561242857508060ff16601c14155b156124395760009350505050611861565b6001868285856040516000815260200160405260405161245c9493929190612afa565b6020604051602081039080840390855afa15801561247e573d6000803e3d6000fd5b5050604051601f190151979650505050505050565b60606124a284846000856124aa565b949350505050565b6060824710156124cc5760405162461bcd60e51b81526004016104c190612e08565b6124d585611beb565b6124f15760405162461bcd60e51b81526004016104c1906133cc565b600080866001600160a01b0316858760405161250d9190612a1a565b60006040518083038185875af1925050503d806000811461254a576040519150601f19603f3d011682016040523d82523d6000602084013e61254f565b606091505b509150915061255f82828661256a565b979650505050505050565b60608315612579575081610919565b8251156125895782518084602001fd5b8160405162461bcd60e51b81526004016104c19190612b18565b600082601f8301126125b3578081fd5b8135602067ffffffffffffffff8211156125cf576125cf61370e565b8082026125dd8282016135d4565b8381528281019086840183880185018910156125f7578687fd5b8693505b8584101561262257803561260e81613724565b8352600193909301929184019184016125fb565b50979650505050505050565b600082601f83011261263e578081fd5b813567ffffffffffffffff8111156126585761265861370e565b61266b601f8201601f19166020016135d4565b81815284602083860101111561267f578283fd5b816020850160208301379081016020019190915292915050565b6000602082840312156126aa578081fd5b813561091981613724565b6000602082840312156126c6578081fd5b815161091981613739565b6000602082840312156126e2578081fd5b813567ffffffffffffffff8111156126f8578182fd5b6124a28482850161262e565b60008060008060008060c0878903121561271c578182fd5b863567ffffffffffffffff80821115612733578384fd5b61273f8a838b0161262e565b97506020890135915061275182613724565b9095506040880135945060608801359061276a82613739565b90935060808801359061277c82613724565b90925060a08801359080821115612791578283fd5b5061279e89828a0161262e565b9150509295509295509295565b6000806000606084860312156127bf578283fd5b833567ffffffffffffffff808211156127d6578485fd5b6127e28783880161262e565b9450602086013591506127f482613724565b90925060408501359080821115612809578283fd5b506128168682870161262e565b9150509250925092565b600080600060608486031215612834578283fd5b833567ffffffffffffffff81111561284a578384fd5b6128568682870161262e565b935050602084013561286781613724565b929592945050506040919091013590565b600080600080600060a0868803121561288f578081fd5b853567ffffffffffffffff808211156128a6578283fd5b6128b289838a0161262e565b965060208801359150808211156128c7578283fd5b6128d389838a016125a3565b955060408801359150808211156128e8578283fd5b6128f489838a016125a3565b94506060880135915060ff8216821461290b578283fd5b90925060808701359080821115612920578283fd5b5061292d8882890161262e565b9150509295509295909350565b60008060006060848603121561294e578283fd5b833567ffffffffffffffff811115612964578384fd5b6129708682870161262e565b93505060208401359150604084013561298881613724565b809150509250925092565b6000602082840312156129a4578081fd5b5051919050565b6000815180845260208085019450808401835b838110156129e35781516001600160a01b0316875295820195908201906001016129be565b509495945050505050565b60008151808452612a06816020860160208601613691565b601f01601f19169290920160200192915050565b60008251612a2c818460208701613691565b9190910192915050565b6001600160a01b0391909116815260200190565b6001600160a01b03929092168252602082015260400190565b6001600160a01b0392831681529116602082015260400190565b6001600160a01b039384168152919092166020820152604081019190915260600190565b600060018060a01b03808716835260806020840152612ac360808401876129ee565b6040840195909552929092166060909101525092915050565b60006020825261091960208301846129ab565b901515815260200190565b93845260ff9290921660208401526040830152606082015260800190565b60006020825261091960208301846129ee565b600060c08252612b3e60c08301896129ee565b6001600160a01b03978816602084015260408301969096525092151560608401529316608082015260a00191909152919050565b600060608252612b8560608301866129ee565b6001600160a01b039490941660208301525060400152919050565b600060a08252612bb360a08301886129ee565b8281036020840152612bc581886129ab565b905060ff861660408401528281036060840152612be281866129ab565b9150508260808301529695505050505050565b60208082526023908201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260408201526265737360e81b606082015260800190565b60208082526028908201527f5468697320616464726573732068617320616c7265616479206265656e20726560408201526719da5cdd195c995960c21b606082015260800190565b60208082526026908201527f57697468647261773a207472616e7366657220746f20746865207a65726f206160408201526564647265737360d01b606082015260800190565b602080825260129082015271496e76616c6964207369676e61747572657360701b604082015260600190565b6020808252602c908201527f4475706c696361746520706172616d657465727320666f72207468652061646460408201526b3932b9b9903a37903537b4b760a11b606082015260800190565b60208082526014908201527326b0b730b3b2b91021b0b713ba1032b6b83a3c9760611b604082015260600190565b6020808252601490820152735369676e617475726573206475706c696361746560601b604082015260600190565b6020808252601e908201527f5472616e73616374696f6e20686173206265656e20636f6d706c657465640000604082015260600190565b6020808252601a908201527f4e6f20656e6f7567682062616c616e6365206f6620746f6b656e000000000000604082015260600190565b60208082526026908201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6040820152651c8818d85b1b60d21b606082015260800190565b60208082526024908201527f45524332303a20446f6573206e6f742061636365707420457468657265756d2060408201526321b7b4b760e11b606082015260800190565b6020808252600e908201526d736c6963655f6f766572666c6f7760901b604082015260600190565b6020808252601b908201527f4f6e6c79206d616e616765722063616e20657865637574652069740000000000604082015260600190565b6020808252603f908201527f5468697320636f6e7472616374206164647265737320646f6573206e6f74206860408201527f6176652073756666696369656e742062616c616e6365206f6620657468657200606082015260800190565b6020808252601d908201527f4d6178206c656e677468206f66207369676e6174757265733a20393735000000604082015260600190565b60208082526024908201527f4552524f523a204465746563746564207a65726f206164647265737320696e206040820152636164647360e01b606082015260800190565b6020808252601690820152752237903737ba3434b73390313c903cb7bab939b2b63360511b604082015260600190565b60208082526012908201527111549493d48e8816995c9bc8185b5bdd5b9d60721b604082015260600190565b60208082526010908201526f29b4b3b730ba3ab932b99032b93937b960811b604082015260600190565b60208082526027908201527f457863656564656420746865206d6178696d756d206e756d626572206f66206d604082015266616e616765727360c81b606082015260800190565b60208082526006908201526511195b9a595960d21b604082015260600190565b60208082526044908201527f5468657265206172652061646472657373657320696e2074686520657869746960408201527f6e672061646472657373206c697374207468617420617265206e6f74206d616e60608201526330b3b2b960e11b608082015260a00190565b60208082526028908201527f5769746864726177616c20616d6f756e74206d75737420626520677265617465604082015267072207468616e20360c41b606082015260800190565b60208082526022908201527f4e6f20656e6f75676820616d6f756e7420666f7220617574686f72697a61746960408201526137b760f11b606082015260800190565b60208082526023908201527f4552524f523a207472616e7366657220746f20746865207a65726f206164647260408201526265737360e81b606082015260800190565b602080825260409082018190527f5468652061646472657373206c6973742074686174206973206265696e672061908201527f6464656420616c7265616479206578697374732061732061206d616e61676572606082015260800190565b6020808252601e908201527f546869732061646472657373206973206e6f7420726567697374657265640000604082015260600190565b6020808252601e908201527f4e6f20656e6f7567682062616c616e6365206f662074686520746f6b656e0000604082015260600190565b60208082526015908201527415985b1a59081cda59db985d1d5c995cc819985a5b605a1b604082015260600190565b60208082526028908201527f546865726520617265206e6f206d616e6167657273206a6f696e696e67206f726040820152672065786974696e6760c01b606082015260800190565b6020808252601d908201527f496e636f6e73697374656e637920457468657265756d20616d6f756e74000000604082015260600190565b602080825260149082015273125d081a185cc81899595b881d5c19dc9859195960621b604082015260600190565b60208082526019908201527f4669786564206c656e677468206f662074784b65793a20363400000000000000604082015260600190565b6020808252601d908201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604082015260600190565b602080825260119082015270736c6963655f6f75744f66426f756e647360781b604082015260600190565b60208082526019908201527f4f6e6c79206f776e65722063616e206578656375746520697400000000000000604082015260600190565b6020808252602c908201527f4475706c696361746520706172616d657465727320666f72207468652061646460408201526b1c995cdcc81d1bc8195e1a5d60a21b606082015260800190565b60208082526017908201527f43616e277420657869742073656564206d616e61676572000000000000000000604082015260600190565b6020808252602a908201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6040820152691bdd081cdd58d8d9595960b21b606082015260800190565b60208082526026908201527f436f6e74726163742063726561746f722063616e6e6f7420616374206173206d60408201526530b730b3b2b960d11b606082015260800190565b60208082526025908201527f5468652061646472657373206973206e6f74206120636f6e7472616374206164604082015264647265737360d81b606082015260800190565b90815260200190565b60ff91909116815260200190565b60405181810167ffffffffffffffff811182821017156135f6576135f661370e565b604052919050565b60008219821115613611576136116136f8565b500190565b600060ff821660ff84168060ff03821115613633576136336136f8565b019392505050565b60008261365657634e487b7160e01b81526012600452602481fd5b500490565b6000816000190483118215151615613675576136756136f8565b500290565b60008282101561368c5761368c6136f8565b500390565b60005b838110156136ac578181015183820152602001613694565b8381111561181e5750506000910152565b60006000198214156136d1576136d16136f8565b5060010190565b600060ff821660ff8114156136ef576136ef6136f8565b60010192915050565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052604160045260246000fd5b6001600160a01b03811681146109e557600080fd5b80151581146109e557600080fdfea26469706673582212203317eb13a9fdcac0935bc45f920d7333ea321ba986550dc0f309fc68b016ef7e64736f6c63430008000033"

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
const IERC721MinterABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"player\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newMinter\",\"type\":\"address\"}],\"name\":\"replaceMinter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

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
// Solidity: function mint(address player, string tokenURI) returns()
func (_IERC721Minter *IERC721MinterTransactor) Mint(opts *bind.TransactOpts, player common.Address, tokenURI string) (*types.Transaction, error) {
	return _IERC721Minter.contract.Transact(opts, "mint", player, tokenURI)
}

// Mint is a paid mutator transaction binding the contract method 0xd0def521.
//
// Solidity: function mint(address player, string tokenURI) returns()
func (_IERC721Minter *IERC721MinterSession) Mint(player common.Address, tokenURI string) (*types.Transaction, error) {
	return _IERC721Minter.Contract.Mint(&_IERC721Minter.TransactOpts, player, tokenURI)
}

// Mint is a paid mutator transaction binding the contract method 0xd0def521.
//
// Solidity: function mint(address player, string tokenURI) returns()
func (_IERC721Minter *IERC721MinterTransactorSession) Mint(player common.Address, tokenURI string) (*types.Transaction, error) {
	return _IERC721Minter.Contract.Mint(&_IERC721Minter.TransactOpts, player, tokenURI)
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

// SafeERC20ABI is the input ABI used to generate the binding from.
const SafeERC20ABI = "[]"

// SafeERC20Bin is the compiled bytecode used for deploying new contracts.
var SafeERC20Bin = "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220558d774c264d2469421acd9e907681298104acf38d1dd7e67ce40d6f9f0913f064736f6c63430008000033"

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
var SafeMathBin = "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220b69c67131b8ed22605cf1e6ac92be4ec2e268ac5825a698a5430aa2cc981c3f164736f6c63430008000033"

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
