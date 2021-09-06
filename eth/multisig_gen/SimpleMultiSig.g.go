// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package multisig_gen

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

// SimpleMultiSigABI is the input ABI used to generate the binding from.
const SimpleMultiSigABI = "[{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"nonceNum_\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"threshold_\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"owners_\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"_confirmAddrs\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_destination\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"Execute\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"bucketIdx\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"expireTime\",\"type\":\"uint256\"},{\"internalType\":\"uint8[]\",\"name\":\"sigV\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"sigR\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"sigS\",\"type\":\"bytes32[]\"},{\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBucketLength\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOwersLength\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"nonceBucket\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ownersArr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"threshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// SimpleMultiSigFuncSigs maps the 4-byte function signature to its string representation.
var SimpleMultiSigFuncSigs = map[string]string{
	"163947e5": "execute(uint16,uint256,uint8[],bytes32[],bytes32[],address,uint256,bytes,address,uint256)",
	"c683a2db": "getBucketLength()",
	"ca7541ee": "getOwersLength()",
	"0d8e6e2c": "getVersion()",
	"c3142f4a": "nonceBucket(uint256)",
	"aa5df9e2": "ownersArr(uint256)",
	"42cde4e8": "threshold()",
}

// SimpleMultiSigBin is the compiled bytecode used for deploying new contracts.
var SimpleMultiSigBin = "0x60806040523480156200001157600080fd5b5060405162001103380380620011038339810160408190526200003491620003f0565b600a82511115801562000048575081518311155b8015620000555750600083115b6200007d5760405162461bcd60e51b8152600401620000749062000553565b60405180910390fd5b60008461ffff161180156200009757506101008461ffff16105b620000b65760405162461bcd60e51b815260040162000074906200057a565b6000805b8351811015620001c757816001600160a01b0316848281518110620000ef57634e487b7160e01b600052603260045260246000fd5b60200260200101516001600160a01b031611620001205760405162461bcd60e51b815260040162000074906200051c565b6001600260008684815181106200014757634e487b7160e01b600052603260045260246000fd5b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002060006101000a81548160ff021916908315150217905550838181518110620001a757634e487b7160e01b600052603260045260246000fd5b602002602001015191508080620001be90620005a1565b915050620000ba565b508251620001dd90600390602086019062000314565b50600184905561ffff85166001600160401b038111156200020e57634e487b7160e01b600052604160045260246000fd5b60405190808252806020026020018201604052801562000238578160200160208202803683370190505b5080516200024f916000916020909101906200037e565b50604051620002ed907fd87cd6ef79d4e2b95e15ce8abf732db51ec771f1ca2edccf22a46c729ac56472907fb7a0bfa1b79f2443f4d73ebb9259cddbcd510b18be6fc4da7d1aa7b1786e73e6907fc89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc690869030907f251543af6a222378665a76fe38dbceae4871a070b7fdaf5c6c30cf758dc33cc090602001620004eb565b604051602081830303815290604052805190602001206004819055505050505050620005df565b8280548282559060005260206000209081019282156200036c579160200282015b828111156200036c57825182546001600160a01b0319166001600160a01b0390911617825560209092019160019091019062000335565b506200037a929150620003bc565b5090565b8280548282559060005260206000209081019282156200036c579160200282015b828111156200036c5782518255916020019190600101906200039f565b5b808211156200037a5760008155600101620003bd565b80516001600160a01b0381168114620003eb57600080fd5b919050565b6000806000806080858703121562000406578384fd5b845161ffff8116811462000418578485fd5b60208681015160408801519296509450906001600160401b03808211156200043e578485fd5b818801915088601f83011262000452578485fd5b815181811115620004675762000467620005c9565b83810260405185828201018181108582111715620004895762000489620005c9565b604052828152858101935084860182860187018d1015620004a8578889fd5b8895505b83861015620004d557620004c081620003d3565b855260019590950194938601938601620004ac565b5060609a909a0151989b979a5050505050505050565b9586526020860194909452604085019290925260608401526001600160a01b0316608083015260a082015260c00190565b60208082526017908201527f72657065617465645f6f776e65722f756e736f72746564000000000000000000604082015260600190565b6020808252600d908201526c18985917dd1a1c995cda1bdb19609a1b604082015260600190565b6020808252600d908201526c6261645f6e6f6e63655f6e756d60981b604082015260600190565b6000600019821415620005c257634e487b7160e01b81526011600452602481fd5b5060010190565b634e487b7160e01b600052604160045260246000fd5b610b1480620005ef6000396000f3fe6080604052600436106100745760003560e01c8063aa5df9e21161004e578063aa5df9e21461012d578063c3142f4a1461015a578063c683a2db1461017a578063ca7541ee1461019c5761007b565b80630d8e6e2c146100be578063163947e5146100e957806342cde4e81461010b5761007b565b3661007b57005b336001600160a01b03167fe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c346040516100b49190610a28565b60405180910390a2005b3480156100ca57600080fd5b506100d36101b1565b6040516100e091906108bd565b60405180910390f35b3480156100f557600080fd5b5061010961010436600461071d565b6101d0565b005b34801561011757600080fd5b50610120610553565b6040516100e09190610a28565b34801561013957600080fd5b5061014d610148366004610816565b610559565b6040516100e09190610849565b34801561016657600080fd5b50610120610175366004610816565b610583565b34801561018657600080fd5b5061018f6105a4565b6040516100e09190610a19565b3480156101a857600080fd5b5061018f6105aa565b6040805180820190915260058152640312e302e360dc1b602082015290565b6001548751146101fb5760405162461bcd60e51b81526004016101f2906109af565b60405180910390fd5b8551875114801561020d575087518751145b6102295760405162461bcd60e51b81526004016101f290610988565b6001600160a01b03821633148061024757506001600160a01b038216155b6102635760405162461bcd60e51b81526004016101f290610935565b8842106102825760405162461bcd60e51b81526004016101f2906109d2565b60005461ffff8b16106102a75760405162461bcd60e51b81526004016101f29061095b565b6000808b61ffff16815481106102cd57634e487b7160e01b600052603260045260246000fd5b600091825260208083209091015486518783012060405191945061031f927f3ee892349ae4bbe61dce18f95115b5dc02daf49204cc602458cd4c1f540d56d7928f928c928c9289918c918c910161085d565b60405160208183030381529060405280519060200120905060006004548260405160200161034e92919061082e565b6040516020818303038152906040528051906020012090506000805b6001548110156104a55760006001848f848151811061039957634e487b7160e01b600052603260045260246000fd5b60200260200101518f85815181106103c157634e487b7160e01b600052603260045260246000fd5b60200260200101518f86815181106103e957634e487b7160e01b600052603260045260246000fd5b60200260200101516040516000815260200160405260405161040e949392919061089f565b6020604051602081039080840390855afa158015610430573d6000803e3d6000fd5b505050602060405103519050826001600160a01b0316816001600160a01b031611801561047557506001600160a01b03811660009081526002602052604090205460ff165b6104915760405162461bcd60e51b81526004016101f2906109f8565b91508061049d81610a97565b91505061036a565b5060008e61ffff16815481106104cb57634e487b7160e01b600052603260045260246000fd5b906000526020600020015460016104e29190610a7f565b60008f61ffff168154811061050757634e487b7160e01b600052603260045260246000fd5b906000526020600020018190555060008080895160208b018c8e8bf19050806105425760405162461bcd60e51b81526004016101f290610910565b505050505050505050505050505050565b60015481565b6003818154811061056957600080fd5b6000918252602090912001546001600160a01b0316905081565b6000818154811061059357600080fd5b600091825260209091200154905081565b60005490565b60035490565b80356001600160a01b03811681146105c757600080fd5b919050565b600082601f8301126105dc578081fd5b813560206105f16105ec83610a5b565b610a31565b828152818101908583018385028701840188101561060d578586fd5b855b8581101561062b5781358452928401929084019060010161060f565b5090979650505050505050565b600082601f830112610648578081fd5b813560206106586105ec83610a5b565b8281528181019085830183850287018401881015610674578586fd5b855b8581101561062b57813560ff8116811461068e578788fd5b84529284019290840190600101610676565b600082601f8301126106b0578081fd5b813567ffffffffffffffff8111156106ca576106ca610ac8565b6106dd601f8201601f1916602001610a31565b8181528460208386010111156106f1578283fd5b816020850160208301379081016020019190915292915050565b803561ffff811681146105c757600080fd5b6000806000806000806000806000806101408b8d03121561073c578586fd5b6107458b61070b565b995060208b0135985060408b013567ffffffffffffffff80821115610768578788fd5b6107748e838f01610638565b995060608d0135915080821115610789578788fd5b6107958e838f016105cc565b985060808d01359150808211156107aa578788fd5b6107b68e838f016105cc565b97506107c460a08e016105b0565b965060c08d0135955060e08d01359150808211156107e0578485fd5b506107ed8d828e016106a0565b9350506107fd6101008c016105b0565b91506101208b013590509295989b9194979a5092959850565b600060208284031215610827578081fd5b5035919050565b61190160f01b81526002810192909252602282015260420190565b6001600160a01b0391909116815260200190565b97885260208801969096526001600160a01b0394851660408801526060870193909352608086019190915260a08501521660c083015260e08201526101000190565b93845260ff9290921660208401526040830152606082015260800190565b6000602080835283518082850152825b818110156108e9578581018301518582016040015282016108cd565b818111156108fa5783604083870101525b50601f01601f1916929092016040019392505050565b6020808252600b908201526a18d85b1b17d9985a5b195960aa1b604082015260600190565b6020808252600c908201526b3130b22fb2bc32b1baba37b960a11b604082015260600190565b6020808252601390820152726275636b65745f6f75745f6f665f72616e676560681b604082015260600190565b6020808252600d908201526c3130b22fb632b72fb917b997bb60991b604082015260600190565b6020808252600990820152683130b22fb92fb632b760b91b604082015260600190565b6020808252600c908201526b1d1a5b5957d95e1c1a5c995960a21b604082015260600190565b6020808252600790820152666261645f73696760c81b604082015260600190565b61ffff91909116815260200190565b90815260200190565b60405181810167ffffffffffffffff81118282101715610a5357610a53610ac8565b604052919050565b600067ffffffffffffffff821115610a7557610a75610ac8565b5060209081020190565b60008219821115610a9257610a92610ab2565b500190565b6000600019821415610aab57610aab610ab2565b5060010190565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052604160045260246000fdfea2646970667358221220bc7822549a298a53adccab9f5f3128f4dd468697557fae9aa39264fabf4ed6c664736f6c63430008000033"

// DeploySimpleMultiSig deploys a new Ethereum contract, binding an instance of SimpleMultiSig to it.
func DeploySimpleMultiSig(auth *bind.TransactOpts, backend bind.ContractBackend, nonceNum_ uint16, threshold_ *big.Int, owners_ []common.Address, chainId *big.Int) (common.Address, *types.Transaction, *SimpleMultiSig, error) {
	parsed, err := abi.JSON(strings.NewReader(SimpleMultiSigABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SimpleMultiSigBin), backend, nonceNum_, threshold_, owners_, chainId)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SimpleMultiSig{SimpleMultiSigCaller: SimpleMultiSigCaller{contract: contract}, SimpleMultiSigTransactor: SimpleMultiSigTransactor{contract: contract}, SimpleMultiSigFilterer: SimpleMultiSigFilterer{contract: contract}}, nil
}

// SimpleMultiSig is an auto generated Go binding around an Ethereum contract.
type SimpleMultiSig struct {
	SimpleMultiSigCaller     // Read-only binding to the contract
	SimpleMultiSigTransactor // Write-only binding to the contract
	SimpleMultiSigFilterer   // Log filterer for contract events
}

// SimpleMultiSigCaller is an auto generated read-only Go binding around an Ethereum contract.
type SimpleMultiSigCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleMultiSigTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SimpleMultiSigTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleMultiSigFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SimpleMultiSigFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleMultiSigSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SimpleMultiSigSession struct {
	Contract     *SimpleMultiSig   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SimpleMultiSigCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SimpleMultiSigCallerSession struct {
	Contract *SimpleMultiSigCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// SimpleMultiSigTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SimpleMultiSigTransactorSession struct {
	Contract     *SimpleMultiSigTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// SimpleMultiSigRaw is an auto generated low-level Go binding around an Ethereum contract.
type SimpleMultiSigRaw struct {
	Contract *SimpleMultiSig // Generic contract binding to access the raw methods on
}

// SimpleMultiSigCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SimpleMultiSigCallerRaw struct {
	Contract *SimpleMultiSigCaller // Generic read-only contract binding to access the raw methods on
}

// SimpleMultiSigTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SimpleMultiSigTransactorRaw struct {
	Contract *SimpleMultiSigTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSimpleMultiSig creates a new instance of SimpleMultiSig, bound to a specific deployed contract.
func NewSimpleMultiSig(address common.Address, backend bind.ContractBackend) (*SimpleMultiSig, error) {
	contract, err := bindSimpleMultiSig(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SimpleMultiSig{SimpleMultiSigCaller: SimpleMultiSigCaller{contract: contract}, SimpleMultiSigTransactor: SimpleMultiSigTransactor{contract: contract}, SimpleMultiSigFilterer: SimpleMultiSigFilterer{contract: contract}}, nil
}

// NewSimpleMultiSigCaller creates a new read-only instance of SimpleMultiSig, bound to a specific deployed contract.
func NewSimpleMultiSigCaller(address common.Address, caller bind.ContractCaller) (*SimpleMultiSigCaller, error) {
	contract, err := bindSimpleMultiSig(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SimpleMultiSigCaller{contract: contract}, nil
}

// NewSimpleMultiSigTransactor creates a new write-only instance of SimpleMultiSig, bound to a specific deployed contract.
func NewSimpleMultiSigTransactor(address common.Address, transactor bind.ContractTransactor) (*SimpleMultiSigTransactor, error) {
	contract, err := bindSimpleMultiSig(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SimpleMultiSigTransactor{contract: contract}, nil
}

// NewSimpleMultiSigFilterer creates a new log filterer instance of SimpleMultiSig, bound to a specific deployed contract.
func NewSimpleMultiSigFilterer(address common.Address, filterer bind.ContractFilterer) (*SimpleMultiSigFilterer, error) {
	contract, err := bindSimpleMultiSig(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SimpleMultiSigFilterer{contract: contract}, nil
}

// bindSimpleMultiSig binds a generic wrapper to an already deployed contract.
func bindSimpleMultiSig(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SimpleMultiSigABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SimpleMultiSig *SimpleMultiSigRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SimpleMultiSig.Contract.SimpleMultiSigCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SimpleMultiSig *SimpleMultiSigRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleMultiSig.Contract.SimpleMultiSigTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SimpleMultiSig *SimpleMultiSigRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SimpleMultiSig.Contract.SimpleMultiSigTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SimpleMultiSig *SimpleMultiSigCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SimpleMultiSig.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SimpleMultiSig *SimpleMultiSigTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleMultiSig.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SimpleMultiSig *SimpleMultiSigTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SimpleMultiSig.Contract.contract.Transact(opts, method, params...)
}

// GetBucketLength is a free data retrieval call binding the contract method 0xc683a2db.
//
// Solidity: function getBucketLength() view returns(uint16)
func (_SimpleMultiSig *SimpleMultiSigCaller) GetBucketLength(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _SimpleMultiSig.contract.Call(opts, &out, "getBucketLength")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// GetBucketLength is a free data retrieval call binding the contract method 0xc683a2db.
//
// Solidity: function getBucketLength() view returns(uint16)
func (_SimpleMultiSig *SimpleMultiSigSession) GetBucketLength() (uint16, error) {
	return _SimpleMultiSig.Contract.GetBucketLength(&_SimpleMultiSig.CallOpts)
}

// GetBucketLength is a free data retrieval call binding the contract method 0xc683a2db.
//
// Solidity: function getBucketLength() view returns(uint16)
func (_SimpleMultiSig *SimpleMultiSigCallerSession) GetBucketLength() (uint16, error) {
	return _SimpleMultiSig.Contract.GetBucketLength(&_SimpleMultiSig.CallOpts)
}

// GetOwersLength is a free data retrieval call binding the contract method 0xca7541ee.
//
// Solidity: function getOwersLength() view returns(uint16)
func (_SimpleMultiSig *SimpleMultiSigCaller) GetOwersLength(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _SimpleMultiSig.contract.Call(opts, &out, "getOwersLength")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// GetOwersLength is a free data retrieval call binding the contract method 0xca7541ee.
//
// Solidity: function getOwersLength() view returns(uint16)
func (_SimpleMultiSig *SimpleMultiSigSession) GetOwersLength() (uint16, error) {
	return _SimpleMultiSig.Contract.GetOwersLength(&_SimpleMultiSig.CallOpts)
}

// GetOwersLength is a free data retrieval call binding the contract method 0xca7541ee.
//
// Solidity: function getOwersLength() view returns(uint16)
func (_SimpleMultiSig *SimpleMultiSigCallerSession) GetOwersLength() (uint16, error) {
	return _SimpleMultiSig.Contract.GetOwersLength(&_SimpleMultiSig.CallOpts)
}

// GetVersion is a free data retrieval call binding the contract method 0x0d8e6e2c.
//
// Solidity: function getVersion() pure returns(string version)
func (_SimpleMultiSig *SimpleMultiSigCaller) GetVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SimpleMultiSig.contract.Call(opts, &out, "getVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetVersion is a free data retrieval call binding the contract method 0x0d8e6e2c.
//
// Solidity: function getVersion() pure returns(string version)
func (_SimpleMultiSig *SimpleMultiSigSession) GetVersion() (string, error) {
	return _SimpleMultiSig.Contract.GetVersion(&_SimpleMultiSig.CallOpts)
}

// GetVersion is a free data retrieval call binding the contract method 0x0d8e6e2c.
//
// Solidity: function getVersion() pure returns(string version)
func (_SimpleMultiSig *SimpleMultiSigCallerSession) GetVersion() (string, error) {
	return _SimpleMultiSig.Contract.GetVersion(&_SimpleMultiSig.CallOpts)
}

// NonceBucket is a free data retrieval call binding the contract method 0xc3142f4a.
//
// Solidity: function nonceBucket(uint256 ) view returns(uint256)
func (_SimpleMultiSig *SimpleMultiSigCaller) NonceBucket(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SimpleMultiSig.contract.Call(opts, &out, "nonceBucket", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NonceBucket is a free data retrieval call binding the contract method 0xc3142f4a.
//
// Solidity: function nonceBucket(uint256 ) view returns(uint256)
func (_SimpleMultiSig *SimpleMultiSigSession) NonceBucket(arg0 *big.Int) (*big.Int, error) {
	return _SimpleMultiSig.Contract.NonceBucket(&_SimpleMultiSig.CallOpts, arg0)
}

// NonceBucket is a free data retrieval call binding the contract method 0xc3142f4a.
//
// Solidity: function nonceBucket(uint256 ) view returns(uint256)
func (_SimpleMultiSig *SimpleMultiSigCallerSession) NonceBucket(arg0 *big.Int) (*big.Int, error) {
	return _SimpleMultiSig.Contract.NonceBucket(&_SimpleMultiSig.CallOpts, arg0)
}

// OwnersArr is a free data retrieval call binding the contract method 0xaa5df9e2.
//
// Solidity: function ownersArr(uint256 ) view returns(address)
func (_SimpleMultiSig *SimpleMultiSigCaller) OwnersArr(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _SimpleMultiSig.contract.Call(opts, &out, "ownersArr", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnersArr is a free data retrieval call binding the contract method 0xaa5df9e2.
//
// Solidity: function ownersArr(uint256 ) view returns(address)
func (_SimpleMultiSig *SimpleMultiSigSession) OwnersArr(arg0 *big.Int) (common.Address, error) {
	return _SimpleMultiSig.Contract.OwnersArr(&_SimpleMultiSig.CallOpts, arg0)
}

// OwnersArr is a free data retrieval call binding the contract method 0xaa5df9e2.
//
// Solidity: function ownersArr(uint256 ) view returns(address)
func (_SimpleMultiSig *SimpleMultiSigCallerSession) OwnersArr(arg0 *big.Int) (common.Address, error) {
	return _SimpleMultiSig.Contract.OwnersArr(&_SimpleMultiSig.CallOpts, arg0)
}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint256)
func (_SimpleMultiSig *SimpleMultiSigCaller) Threshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SimpleMultiSig.contract.Call(opts, &out, "threshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint256)
func (_SimpleMultiSig *SimpleMultiSigSession) Threshold() (*big.Int, error) {
	return _SimpleMultiSig.Contract.Threshold(&_SimpleMultiSig.CallOpts)
}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint256)
func (_SimpleMultiSig *SimpleMultiSigCallerSession) Threshold() (*big.Int, error) {
	return _SimpleMultiSig.Contract.Threshold(&_SimpleMultiSig.CallOpts)
}

// Execute is a paid mutator transaction binding the contract method 0x163947e5.
//
// Solidity: function execute(uint16 bucketIdx, uint256 expireTime, uint8[] sigV, bytes32[] sigR, bytes32[] sigS, address destination, uint256 value, bytes data, address executor, uint256 gasLimit) returns()
func (_SimpleMultiSig *SimpleMultiSigTransactor) Execute(opts *bind.TransactOpts, bucketIdx uint16, expireTime *big.Int, sigV []uint8, sigR [][32]byte, sigS [][32]byte, destination common.Address, value *big.Int, data []byte, executor common.Address, gasLimit *big.Int) (*types.Transaction, error) {
	return _SimpleMultiSig.contract.Transact(opts, "execute", bucketIdx, expireTime, sigV, sigR, sigS, destination, value, data, executor, gasLimit)
}

// Execute is a paid mutator transaction binding the contract method 0x163947e5.
//
// Solidity: function execute(uint16 bucketIdx, uint256 expireTime, uint8[] sigV, bytes32[] sigR, bytes32[] sigS, address destination, uint256 value, bytes data, address executor, uint256 gasLimit) returns()
func (_SimpleMultiSig *SimpleMultiSigSession) Execute(bucketIdx uint16, expireTime *big.Int, sigV []uint8, sigR [][32]byte, sigS [][32]byte, destination common.Address, value *big.Int, data []byte, executor common.Address, gasLimit *big.Int) (*types.Transaction, error) {
	return _SimpleMultiSig.Contract.Execute(&_SimpleMultiSig.TransactOpts, bucketIdx, expireTime, sigV, sigR, sigS, destination, value, data, executor, gasLimit)
}

// Execute is a paid mutator transaction binding the contract method 0x163947e5.
//
// Solidity: function execute(uint16 bucketIdx, uint256 expireTime, uint8[] sigV, bytes32[] sigR, bytes32[] sigS, address destination, uint256 value, bytes data, address executor, uint256 gasLimit) returns()
func (_SimpleMultiSig *SimpleMultiSigTransactorSession) Execute(bucketIdx uint16, expireTime *big.Int, sigV []uint8, sigR [][32]byte, sigS [][32]byte, destination common.Address, value *big.Int, data []byte, executor common.Address, gasLimit *big.Int) (*types.Transaction, error) {
	return _SimpleMultiSig.Contract.Execute(&_SimpleMultiSig.TransactOpts, bucketIdx, expireTime, sigV, sigR, sigS, destination, value, data, executor, gasLimit)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_SimpleMultiSig *SimpleMultiSigTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _SimpleMultiSig.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_SimpleMultiSig *SimpleMultiSigSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _SimpleMultiSig.Contract.Fallback(&_SimpleMultiSig.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_SimpleMultiSig *SimpleMultiSigTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _SimpleMultiSig.Contract.Fallback(&_SimpleMultiSig.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SimpleMultiSig *SimpleMultiSigTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleMultiSig.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SimpleMultiSig *SimpleMultiSigSession) Receive() (*types.Transaction, error) {
	return _SimpleMultiSig.Contract.Receive(&_SimpleMultiSig.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SimpleMultiSig *SimpleMultiSigTransactorSession) Receive() (*types.Transaction, error) {
	return _SimpleMultiSig.Contract.Receive(&_SimpleMultiSig.TransactOpts)
}

// SimpleMultiSigDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the SimpleMultiSig contract.
type SimpleMultiSigDepositIterator struct {
	Event *SimpleMultiSigDeposit // Event containing the contract specifics and raw log

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
func (it *SimpleMultiSigDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleMultiSigDeposit)
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
		it.Event = new(SimpleMultiSigDeposit)
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
func (it *SimpleMultiSigDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleMultiSigDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleMultiSigDeposit represents a Deposit event raised by the SimpleMultiSig contract.
type SimpleMultiSigDeposit struct {
	From  common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed _from, uint256 _value)
func (_SimpleMultiSig *SimpleMultiSigFilterer) FilterDeposit(opts *bind.FilterOpts, _from []common.Address) (*SimpleMultiSigDepositIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _SimpleMultiSig.contract.FilterLogs(opts, "Deposit", _fromRule)
	if err != nil {
		return nil, err
	}
	return &SimpleMultiSigDepositIterator{contract: _SimpleMultiSig.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed _from, uint256 _value)
func (_SimpleMultiSig *SimpleMultiSigFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *SimpleMultiSigDeposit, _from []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _SimpleMultiSig.contract.WatchLogs(opts, "Deposit", _fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleMultiSigDeposit)
				if err := _SimpleMultiSig.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed _from, uint256 _value)
func (_SimpleMultiSig *SimpleMultiSigFilterer) ParseDeposit(log types.Log) (*SimpleMultiSigDeposit, error) {
	event := new(SimpleMultiSigDeposit)
	if err := _SimpleMultiSig.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SimpleMultiSigExecuteIterator is returned from FilterExecute and is used to iterate over the raw logs and unpacked data for Execute events raised by the SimpleMultiSig contract.
type SimpleMultiSigExecuteIterator struct {
	Event *SimpleMultiSigExecute // Event containing the contract specifics and raw log

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
func (it *SimpleMultiSigExecuteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleMultiSigExecute)
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
		it.Event = new(SimpleMultiSigExecute)
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
func (it *SimpleMultiSigExecuteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleMultiSigExecuteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleMultiSigExecute represents a Execute event raised by the SimpleMultiSig contract.
type SimpleMultiSigExecute struct {
	ConfirmAddrs []common.Address
	Destination  common.Address
	Value        *big.Int
	Data         []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterExecute is a free log retrieval operation binding the contract event 0x07f4110a9f6788eae6a0b088d9aca06ec3cd9e2c6eae12a1d393d6d041d18c30.
//
// Solidity: event Execute(address[] _confirmAddrs, address _destination, uint256 _value, bytes data)
func (_SimpleMultiSig *SimpleMultiSigFilterer) FilterExecute(opts *bind.FilterOpts) (*SimpleMultiSigExecuteIterator, error) {

	logs, sub, err := _SimpleMultiSig.contract.FilterLogs(opts, "Execute")
	if err != nil {
		return nil, err
	}
	return &SimpleMultiSigExecuteIterator{contract: _SimpleMultiSig.contract, event: "Execute", logs: logs, sub: sub}, nil
}

// WatchExecute is a free log subscription operation binding the contract event 0x07f4110a9f6788eae6a0b088d9aca06ec3cd9e2c6eae12a1d393d6d041d18c30.
//
// Solidity: event Execute(address[] _confirmAddrs, address _destination, uint256 _value, bytes data)
func (_SimpleMultiSig *SimpleMultiSigFilterer) WatchExecute(opts *bind.WatchOpts, sink chan<- *SimpleMultiSigExecute) (event.Subscription, error) {

	logs, sub, err := _SimpleMultiSig.contract.WatchLogs(opts, "Execute")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleMultiSigExecute)
				if err := _SimpleMultiSig.contract.UnpackLog(event, "Execute", log); err != nil {
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

// ParseExecute is a log parse operation binding the contract event 0x07f4110a9f6788eae6a0b088d9aca06ec3cd9e2c6eae12a1d393d6d041d18c30.
//
// Solidity: event Execute(address[] _confirmAddrs, address _destination, uint256 _value, bytes data)
func (_SimpleMultiSig *SimpleMultiSigFilterer) ParseExecute(log types.Log) (*SimpleMultiSigExecute, error) {
	event := new(SimpleMultiSigExecute)
	if err := _SimpleMultiSig.contract.UnpackLog(event, "Execute", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
