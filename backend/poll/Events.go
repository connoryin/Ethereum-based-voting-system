// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package poll

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

// EventspollInfo is an auto generated low-level Go binding around an user-defined struct.
type EventspollInfo struct {
	Owner   *big.Int
	EventId *big.Int
	Addr    common.Address
	Exists  bool
	IsEnd   bool
}

// EventsMetaData contains all meta data concerning the Events contract.
var EventsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"adminID\",\"type\":\"uint256\"}],\"name\":\"getPollInfoByAdminID\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"owner\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"eventId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isEnd\",\"type\":\"bool\"}],\"internalType\":\"structEvents.pollInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"invCode\",\"type\":\"bytes32\"}],\"name\":\"getPollInfoByInvCode\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"invCodes\",\"type\":\"bytes32[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"owner\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"eventId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isEnd\",\"type\":\"bool\"}],\"internalType\":\"structEvents.pollInfo\",\"name\":\"info\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"adminID\",\"type\":\"uint256\"}],\"name\":\"uploadPollInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610c05806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80632df5b5a8146100465780635889504714610079578063904176cf146100a9575b600080fd5b610060600480360381019061005b919061052f565b6100c5565b60405161007094939291906105d1565b60405180910390f35b610093600480360381019061008e9190610642565b6101c1565b6040516100a091906107b3565b60405180910390f35b6100c360048036038101906100be9190610a17565b6102d3565b005b60008060008060008086815260200190815260200160002060020160149054906101000a900460ff1661012d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161012490610ae3565b60405180910390fd5b600080868152602001908152602001600020600001546000808781526020019081526020016000206001015460008088815260200190815260200160002060020160159054906101000a900460ff1660008089815260200190815260200160002060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1693509350935093509193509193565b606060016000838152602001908152602001600020805480602002602001604051908101604052809291908181526020016000905b828210156102c857838290600052602060002090600302016040518060a001604052908160008201548152602001600182015481526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820160149054906101000a900460ff161515151581526020016002820160159054906101000a900460ff161515151581525050815260200190600101906101f6565b505050509050919050565b60005b835181101561040557826000808684815181106102f6576102f5610b03565b5b60200260200101518152602001908152602001600020600082015181600001556020820151816001015560408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160020160146101000a81548160ff02191690831515021790555060808201518160020160156101000a81548160ff02191690831515021790555090505060016000808684815181106103c2576103c1610b03565b5b6020026020010151815260200190815260200160002060020160146101000a81548160ff02191690831515021790555080806103fd90610b61565b9150506102d6565b5060016000828152602001908152602001600020829080600181540180825580915050600190039060005260206000209060030201600090919091909150600082015181600001556020820151816001015560408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160020160146101000a81548160ff02191690831515021790555060808201518160020160156101000a81548160ff0219169083151502179055505050505050565b6000604051905090565b600080fd5b600080fd5b6000819050919050565b61050c816104f9565b811461051757600080fd5b50565b60008135905061052981610503565b92915050565b600060208284031215610545576105446104ef565b5b60006105538482850161051a565b91505092915050565b6000819050919050565b61056f8161055c565b82525050565b60008115159050919050565b61058a81610575565b82525050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006105bb82610590565b9050919050565b6105cb816105b0565b82525050565b60006080820190506105e66000830187610566565b6105f36020830186610566565b6106006040830185610581565b61060d60608301846105c2565b95945050505050565b61061f8161055c565b811461062a57600080fd5b50565b60008135905061063c81610616565b92915050565b600060208284031215610658576106576104ef565b5b60006106668482850161062d565b91505092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b6106a48161055c565b82525050565b6106b3816105b0565b82525050565b6106c281610575565b82525050565b60a0820160008201516106de600085018261069b565b5060208201516106f1602085018261069b565b50604082015161070460408501826106aa565b50606082015161071760608501826106b9565b50608082015161072a60808501826106b9565b50505050565b600061073c83836106c8565b60a08301905092915050565b6000602082019050919050565b60006107608261066f565b61076a818561067a565b93506107758361068b565b8060005b838110156107a657815161078d8882610730565b975061079883610748565b925050600181019050610779565b5085935050505092915050565b600060208201905081810360008301526107cd8184610755565b905092915050565b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610823826107da565b810181811067ffffffffffffffff82111715610842576108416107eb565b5b80604052505050565b60006108556104e5565b9050610861828261081a565b919050565b600067ffffffffffffffff821115610881576108806107eb565b5b602082029050602081019050919050565b600080fd5b60006108aa6108a584610866565b61084b565b905080838252602082019050602084028301858111156108cd576108cc610892565b5b835b818110156108f657806108e2888261051a565b8452602084019350506020810190506108cf565b5050509392505050565b600082601f830112610915576109146107d5565b5b8135610925848260208601610897565b91505092915050565b600080fd5b61093c816105b0565b811461094757600080fd5b50565b60008135905061095981610933565b92915050565b61096881610575565b811461097357600080fd5b50565b6000813590506109858161095f565b92915050565b600060a082840312156109a1576109a061092e565b5b6109ab60a061084b565b905060006109bb8482850161062d565b60008301525060206109cf8482850161062d565b60208301525060406109e38482850161094a565b60408301525060606109f784828501610976565b6060830152506080610a0b84828501610976565b60808301525092915050565b600080600060e08486031215610a3057610a2f6104ef565b5b600084013567ffffffffffffffff811115610a4e57610a4d6104f4565b5b610a5a86828701610900565b9350506020610a6b8682870161098b565b92505060c0610a7c8682870161062d565b9150509250925092565b600082825260208201905092915050565b7f4576656e7420646f6573206e6f74206578697374000000000000000000000000600082015250565b6000610acd601483610a86565b9150610ad882610a97565b602082019050919050565b60006020820190508181036000830152610afc81610ac0565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000610b6c8261055c565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610b9e57610b9d610b32565b5b60018201905091905056fea2646970667358221220b74c19147bb42f6e2d60525437d0291e3f15094de906c4590a27a0644d9b6d4b64736f6c637828302e382e31342d646576656c6f702e323032322e342e31352b636f6d6d69742e35353931373430350059",
}

// EventsABI is the input ABI used to generate the binding from.
// Deprecated: Use EventsMetaData.ABI instead.
var EventsABI = EventsMetaData.ABI

// EventsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EventsMetaData.Bin instead.
var EventsBin = EventsMetaData.Bin

// DeployEvents deploys a new Ethereum contract, binding an instance of Events to it.
func DeployEvents(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Events, error) {
	parsed, err := EventsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EventsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Events{EventsCaller: EventsCaller{contract: contract}, EventsTransactor: EventsTransactor{contract: contract}, EventsFilterer: EventsFilterer{contract: contract}}, nil
}

// Events is an auto generated Go binding around an Ethereum contract.
type Events struct {
	EventsCaller     // Read-only binding to the contract
	EventsTransactor // Write-only binding to the contract
	EventsFilterer   // Log filterer for contract events
}

// EventsCaller is an auto generated read-only Go binding around an Ethereum contract.
type EventsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EventsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EventsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EventsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EventsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EventsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EventsSession struct {
	Contract     *Events           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EventsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EventsCallerSession struct {
	Contract *EventsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// EventsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EventsTransactorSession struct {
	Contract     *EventsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EventsRaw is an auto generated low-level Go binding around an Ethereum contract.
type EventsRaw struct {
	Contract *Events // Generic contract binding to access the raw methods on
}

// EventsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EventsCallerRaw struct {
	Contract *EventsCaller // Generic read-only contract binding to access the raw methods on
}

// EventsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EventsTransactorRaw struct {
	Contract *EventsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEvents creates a new instance of Events, bound to a specific deployed contract.
func NewEvents(address common.Address, backend bind.ContractBackend) (*Events, error) {
	contract, err := bindEvents(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Events{EventsCaller: EventsCaller{contract: contract}, EventsTransactor: EventsTransactor{contract: contract}, EventsFilterer: EventsFilterer{contract: contract}}, nil
}

// NewEventsCaller creates a new read-only instance of Events, bound to a specific deployed contract.
func NewEventsCaller(address common.Address, caller bind.ContractCaller) (*EventsCaller, error) {
	contract, err := bindEvents(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EventsCaller{contract: contract}, nil
}

// NewEventsTransactor creates a new write-only instance of Events, bound to a specific deployed contract.
func NewEventsTransactor(address common.Address, transactor bind.ContractTransactor) (*EventsTransactor, error) {
	contract, err := bindEvents(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EventsTransactor{contract: contract}, nil
}

// NewEventsFilterer creates a new log filterer instance of Events, bound to a specific deployed contract.
func NewEventsFilterer(address common.Address, filterer bind.ContractFilterer) (*EventsFilterer, error) {
	contract, err := bindEvents(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EventsFilterer{contract: contract}, nil
}

// bindEvents binds a generic wrapper to an already deployed contract.
func bindEvents(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EventsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Events *EventsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Events.Contract.EventsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Events *EventsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Events.Contract.EventsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Events *EventsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Events.Contract.EventsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Events *EventsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Events.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Events *EventsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Events.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Events *EventsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Events.Contract.contract.Transact(opts, method, params...)
}

// GetPollInfoByAdminID is a free data retrieval call binding the contract method 0x58895047.
//
// Solidity: function getPollInfoByAdminID(uint256 adminID) view returns((uint256,uint256,address,bool,bool)[])
func (_Events *EventsCaller) GetPollInfoByAdminID(opts *bind.CallOpts, adminID *big.Int) ([]EventspollInfo, error) {
	var out []interface{}
	err := _Events.contract.Call(opts, &out, "getPollInfoByAdminID", adminID)

	if err != nil {
		return *new([]EventspollInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]EventspollInfo)).(*[]EventspollInfo)

	return out0, err

}

// GetPollInfoByAdminID is a free data retrieval call binding the contract method 0x58895047.
//
// Solidity: function getPollInfoByAdminID(uint256 adminID) view returns((uint256,uint256,address,bool,bool)[])
func (_Events *EventsSession) GetPollInfoByAdminID(adminID *big.Int) ([]EventspollInfo, error) {
	return _Events.Contract.GetPollInfoByAdminID(&_Events.CallOpts, adminID)
}

// GetPollInfoByAdminID is a free data retrieval call binding the contract method 0x58895047.
//
// Solidity: function getPollInfoByAdminID(uint256 adminID) view returns((uint256,uint256,address,bool,bool)[])
func (_Events *EventsCallerSession) GetPollInfoByAdminID(adminID *big.Int) ([]EventspollInfo, error) {
	return _Events.Contract.GetPollInfoByAdminID(&_Events.CallOpts, adminID)
}

// GetPollInfoByInvCode is a free data retrieval call binding the contract method 0x2df5b5a8.
//
// Solidity: function getPollInfoByInvCode(bytes32 invCode) view returns(uint256, uint256, bool, address)
func (_Events *EventsCaller) GetPollInfoByInvCode(opts *bind.CallOpts, invCode [32]byte) (*big.Int, *big.Int, bool, common.Address, error) {
	var out []interface{}
	err := _Events.contract.Call(opts, &out, "getPollInfoByInvCode", invCode)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(bool), *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(bool)).(*bool)
	out3 := *abi.ConvertType(out[3], new(common.Address)).(*common.Address)

	return out0, out1, out2, out3, err

}

// GetPollInfoByInvCode is a free data retrieval call binding the contract method 0x2df5b5a8.
//
// Solidity: function getPollInfoByInvCode(bytes32 invCode) view returns(uint256, uint256, bool, address)
func (_Events *EventsSession) GetPollInfoByInvCode(invCode [32]byte) (*big.Int, *big.Int, bool, common.Address, error) {
	return _Events.Contract.GetPollInfoByInvCode(&_Events.CallOpts, invCode)
}

// GetPollInfoByInvCode is a free data retrieval call binding the contract method 0x2df5b5a8.
//
// Solidity: function getPollInfoByInvCode(bytes32 invCode) view returns(uint256, uint256, bool, address)
func (_Events *EventsCallerSession) GetPollInfoByInvCode(invCode [32]byte) (*big.Int, *big.Int, bool, common.Address, error) {
	return _Events.Contract.GetPollInfoByInvCode(&_Events.CallOpts, invCode)
}

// UploadPollInfo is a paid mutator transaction binding the contract method 0x904176cf.
//
// Solidity: function uploadPollInfo(bytes32[] invCodes, (uint256,uint256,address,bool,bool) info, uint256 adminID) returns()
func (_Events *EventsTransactor) UploadPollInfo(opts *bind.TransactOpts, invCodes [][32]byte, info EventspollInfo, adminID *big.Int) (*types.Transaction, error) {
	return _Events.contract.Transact(opts, "uploadPollInfo", invCodes, info, adminID)
}

// UploadPollInfo is a paid mutator transaction binding the contract method 0x904176cf.
//
// Solidity: function uploadPollInfo(bytes32[] invCodes, (uint256,uint256,address,bool,bool) info, uint256 adminID) returns()
func (_Events *EventsSession) UploadPollInfo(invCodes [][32]byte, info EventspollInfo, adminID *big.Int) (*types.Transaction, error) {
	return _Events.Contract.UploadPollInfo(&_Events.TransactOpts, invCodes, info, adminID)
}

// UploadPollInfo is a paid mutator transaction binding the contract method 0x904176cf.
//
// Solidity: function uploadPollInfo(bytes32[] invCodes, (uint256,uint256,address,bool,bool) info, uint256 adminID) returns()
func (_Events *EventsTransactorSession) UploadPollInfo(invCodes [][32]byte, info EventspollInfo, adminID *big.Int) (*types.Transaction, error) {
	return _Events.Contract.UploadPollInfo(&_Events.TransactOpts, invCodes, info, adminID)
}
