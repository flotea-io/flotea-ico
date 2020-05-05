// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package floico

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
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// FloicoABI is the input ABI used to generate the binding from.
const FloicoABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddr\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"_participates\",\"type\":\"address[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"names\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"BoughtTokens\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"buyOverBackend\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"res\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"buyTokens\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_description\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_durationHours\",\"type\":\"uint256\"},{\"internalType\":\"enumVotingIco.ActionType\",\"name\":\"_actionType\",\"type\":\"uint8\"},{\"internalType\":\"addresspayable\",\"name\":\"_actionAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_name\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"createProposal\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalIndex\",\"type\":\"uint256\"}],\"name\":\"finishProposal\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"description\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"result\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalIndex\",\"type\":\"uint256\"}],\"name\":\"FinishVoting\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"transport\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"description\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumVotingIco.ActionType\",\"name\":\"actionType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"actionAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalIndex\",\"type\":\"uint256\"}],\"name\":\"ProposalCreated\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"discount\",\"type\":\"uint256\"}],\"name\":\"setDiscount\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_enabled\",\"type\":\"bool\"}],\"name\":\"setEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"tokenFallback\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"description\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"vote\",\"type\":\"uint8\"}],\"name\":\"Vote\",\"type\":\"event\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"vote\",\"type\":\"uint8\"}],\"name\":\"voteInProposal\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"enumVotingIco.ActionType\",\"name\":\"_actionType\",\"type\":\"uint8\"},{\"internalType\":\"addresspayable\",\"name\":\"_actionAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_senderAddress\",\"type\":\"address\"}],\"name\":\"beforeCreateProposal\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalIndex\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"}],\"name\":\"beforeFinishProposal\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"error\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"votedYes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"votedNo\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalIndex\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"}],\"name\":\"beforeVoteInProposal\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"enabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getPhase\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"res\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"getPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"weiAmount\",\"type\":\"uint256\"}],\"name\":\"getTokensForWei\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"goalReached\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"info\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"initialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"initialTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"leftInActualPhase\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"res\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"participants\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"participantsLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"phase\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"phaseInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"priceDiscount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"proposals\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"description\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"result\",\"type\":\"uint8\"},{\"internalType\":\"enumVotingIco.ActionType\",\"name\":\"actionType\",\"type\":\"uint8\"},{\"internalType\":\"addresspayable\",\"name\":\"actionAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"proposalsLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"raisedAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"startTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"statusOfProposal\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint8[]\",\"name\":\"\",\"type\":\"uint8[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokensAvailable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"weiAailable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Floico is an auto generated Go binding around an Ethereum contract.
type Floico struct {
	FloicoCaller     // Read-only binding to the contract
	FloicoTransactor // Write-only binding to the contract
	FloicoFilterer   // Log filterer for contract events
}

// FloicoCaller is an auto generated read-only Go binding around an Ethereum contract.
type FloicoCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FloicoTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FloicoTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FloicoFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FloicoFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FloicoSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FloicoSession struct {
	Contract     *Floico           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FloicoCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FloicoCallerSession struct {
	Contract *FloicoCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// FloicoTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FloicoTransactorSession struct {
	Contract     *FloicoTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FloicoRaw is an auto generated low-level Go binding around an Ethereum contract.
type FloicoRaw struct {
	Contract *Floico // Generic contract binding to access the raw methods on
}

// FloicoCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FloicoCallerRaw struct {
	Contract *FloicoCaller // Generic read-only contract binding to access the raw methods on
}

// FloicoTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FloicoTransactorRaw struct {
	Contract *FloicoTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFloico creates a new instance of Floico, bound to a specific deployed contract.
func NewFloico(address common.Address, backend bind.ContractBackend) (*Floico, error) {
	contract, err := bindFloico(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Floico{FloicoCaller: FloicoCaller{contract: contract}, FloicoTransactor: FloicoTransactor{contract: contract}, FloicoFilterer: FloicoFilterer{contract: contract}}, nil
}

// NewFloicoCaller creates a new read-only instance of Floico, bound to a specific deployed contract.
func NewFloicoCaller(address common.Address, caller bind.ContractCaller) (*FloicoCaller, error) {
	contract, err := bindFloico(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FloicoCaller{contract: contract}, nil
}

// NewFloicoTransactor creates a new write-only instance of Floico, bound to a specific deployed contract.
func NewFloicoTransactor(address common.Address, transactor bind.ContractTransactor) (*FloicoTransactor, error) {
	contract, err := bindFloico(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FloicoTransactor{contract: contract}, nil
}

// NewFloicoFilterer creates a new log filterer instance of Floico, bound to a specific deployed contract.
func NewFloicoFilterer(address common.Address, filterer bind.ContractFilterer) (*FloicoFilterer, error) {
	contract, err := bindFloico(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FloicoFilterer{contract: contract}, nil
}

// bindFloico binds a generic wrapper to an already deployed contract.
func bindFloico(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FloicoABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Floico *FloicoRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Floico.Contract.FloicoCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Floico *FloicoRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Floico.Contract.FloicoTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Floico *FloicoRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Floico.Contract.FloicoTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Floico *FloicoCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Floico.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Floico *FloicoTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Floico.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Floico *FloicoTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Floico.Contract.contract.Transact(opts, method, params...)
}

// BeforeCreateProposal is a free data retrieval call binding the contract method 0x4c9807bb.
//
// Solidity: function beforeCreateProposal(uint8 _actionType, address _actionAddress, address _senderAddress) constant returns(bool, string)
func (_Floico *FloicoCaller) BeforeCreateProposal(opts *bind.CallOpts, _actionType uint8, _actionAddress common.Address, _senderAddress common.Address) (bool, string, error) {
	var (
		ret0 = new(bool)
		ret1 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Floico.contract.Call(opts, out, "beforeCreateProposal", _actionType, _actionAddress, _senderAddress)
	return *ret0, *ret1, err
}

// BeforeCreateProposal is a free data retrieval call binding the contract method 0x4c9807bb.
//
// Solidity: function beforeCreateProposal(uint8 _actionType, address _actionAddress, address _senderAddress) constant returns(bool, string)
func (_Floico *FloicoSession) BeforeCreateProposal(_actionType uint8, _actionAddress common.Address, _senderAddress common.Address) (bool, string, error) {
	return _Floico.Contract.BeforeCreateProposal(&_Floico.CallOpts, _actionType, _actionAddress, _senderAddress)
}

// BeforeCreateProposal is a free data retrieval call binding the contract method 0x4c9807bb.
//
// Solidity: function beforeCreateProposal(uint8 _actionType, address _actionAddress, address _senderAddress) constant returns(bool, string)
func (_Floico *FloicoCallerSession) BeforeCreateProposal(_actionType uint8, _actionAddress common.Address, _senderAddress common.Address) (bool, string, error) {
	return _Floico.Contract.BeforeCreateProposal(&_Floico.CallOpts, _actionType, _actionAddress, _senderAddress)
}

// BeforeFinishProposal is a free data retrieval call binding the contract method 0x6ca19679.
//
// Solidity: function beforeFinishProposal(uint256 proposalIndex, address senderAddress) constant returns(bool error, string message, uint256 votedYes, uint256 votedNo)
func (_Floico *FloicoCaller) BeforeFinishProposal(opts *bind.CallOpts, proposalIndex *big.Int, senderAddress common.Address) (struct {
	Error    bool
	Message  string
	VotedYes *big.Int
	VotedNo  *big.Int
}, error) {
	ret := new(struct {
		Error    bool
		Message  string
		VotedYes *big.Int
		VotedNo  *big.Int
	})
	out := ret
	err := _Floico.contract.Call(opts, out, "beforeFinishProposal", proposalIndex, senderAddress)
	return *ret, err
}

// BeforeFinishProposal is a free data retrieval call binding the contract method 0x6ca19679.
//
// Solidity: function beforeFinishProposal(uint256 proposalIndex, address senderAddress) constant returns(bool error, string message, uint256 votedYes, uint256 votedNo)
func (_Floico *FloicoSession) BeforeFinishProposal(proposalIndex *big.Int, senderAddress common.Address) (struct {
	Error    bool
	Message  string
	VotedYes *big.Int
	VotedNo  *big.Int
}, error) {
	return _Floico.Contract.BeforeFinishProposal(&_Floico.CallOpts, proposalIndex, senderAddress)
}

// BeforeFinishProposal is a free data retrieval call binding the contract method 0x6ca19679.
//
// Solidity: function beforeFinishProposal(uint256 proposalIndex, address senderAddress) constant returns(bool error, string message, uint256 votedYes, uint256 votedNo)
func (_Floico *FloicoCallerSession) BeforeFinishProposal(proposalIndex *big.Int, senderAddress common.Address) (struct {
	Error    bool
	Message  string
	VotedYes *big.Int
	VotedNo  *big.Int
}, error) {
	return _Floico.Contract.BeforeFinishProposal(&_Floico.CallOpts, proposalIndex, senderAddress)
}

// BeforeVoteInProposal is a free data retrieval call binding the contract method 0xc4330507.
//
// Solidity: function beforeVoteInProposal(uint256 proposalIndex, address senderAddress) constant returns(bool, string)
func (_Floico *FloicoCaller) BeforeVoteInProposal(opts *bind.CallOpts, proposalIndex *big.Int, senderAddress common.Address) (bool, string, error) {
	var (
		ret0 = new(bool)
		ret1 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Floico.contract.Call(opts, out, "beforeVoteInProposal", proposalIndex, senderAddress)
	return *ret0, *ret1, err
}

// BeforeVoteInProposal is a free data retrieval call binding the contract method 0xc4330507.
//
// Solidity: function beforeVoteInProposal(uint256 proposalIndex, address senderAddress) constant returns(bool, string)
func (_Floico *FloicoSession) BeforeVoteInProposal(proposalIndex *big.Int, senderAddress common.Address) (bool, string, error) {
	return _Floico.Contract.BeforeVoteInProposal(&_Floico.CallOpts, proposalIndex, senderAddress)
}

// BeforeVoteInProposal is a free data retrieval call binding the contract method 0xc4330507.
//
// Solidity: function beforeVoteInProposal(uint256 proposalIndex, address senderAddress) constant returns(bool, string)
func (_Floico *FloicoCallerSession) BeforeVoteInProposal(proposalIndex *big.Int, senderAddress common.Address) (bool, string, error) {
	return _Floico.Contract.BeforeVoteInProposal(&_Floico.CallOpts, proposalIndex, senderAddress)
}

// Enabled is a free data retrieval call binding the contract method 0x238dafe0.
//
// Solidity: function enabled() constant returns(bool)
func (_Floico *FloicoCaller) Enabled(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Floico.contract.Call(opts, out, "enabled")
	return *ret0, err
}

// Enabled is a free data retrieval call binding the contract method 0x238dafe0.
//
// Solidity: function enabled() constant returns(bool)
func (_Floico *FloicoSession) Enabled() (bool, error) {
	return _Floico.Contract.Enabled(&_Floico.CallOpts)
}

// Enabled is a free data retrieval call binding the contract method 0x238dafe0.
//
// Solidity: function enabled() constant returns(bool)
func (_Floico *FloicoCallerSession) Enabled() (bool, error) {
	return _Floico.Contract.Enabled(&_Floico.CallOpts)
}

// GetPhase is a free data retrieval call binding the contract method 0xeced0280.
//
// Solidity: function getPhase() constant returns(uint256 res)
func (_Floico *FloicoCaller) GetPhase(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Floico.contract.Call(opts, out, "getPhase")
	return *ret0, err
}

// GetPhase is a free data retrieval call binding the contract method 0xeced0280.
//
// Solidity: function getPhase() constant returns(uint256 res)
func (_Floico *FloicoSession) GetPhase() (*big.Int, error) {
	return _Floico.Contract.GetPhase(&_Floico.CallOpts)
}

// GetPhase is a free data retrieval call binding the contract method 0xeced0280.
//
// Solidity: function getPhase() constant returns(uint256 res)
func (_Floico *FloicoCallerSession) GetPhase() (*big.Int, error) {
	return _Floico.Contract.GetPhase(&_Floico.CallOpts)
}

// GetPrice is a free data retrieval call binding the contract method 0xe7572230.
//
// Solidity: function getPrice(uint256 amount) constant returns(uint256)
func (_Floico *FloicoCaller) GetPrice(opts *bind.CallOpts, amount *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Floico.contract.Call(opts, out, "getPrice", amount)
	return *ret0, err
}

// GetPrice is a free data retrieval call binding the contract method 0xe7572230.
//
// Solidity: function getPrice(uint256 amount) constant returns(uint256)
func (_Floico *FloicoSession) GetPrice(amount *big.Int) (*big.Int, error) {
	return _Floico.Contract.GetPrice(&_Floico.CallOpts, amount)
}

// GetPrice is a free data retrieval call binding the contract method 0xe7572230.
//
// Solidity: function getPrice(uint256 amount) constant returns(uint256)
func (_Floico *FloicoCallerSession) GetPrice(amount *big.Int) (*big.Int, error) {
	return _Floico.Contract.GetPrice(&_Floico.CallOpts, amount)
}

// GetTokensForWei is a free data retrieval call binding the contract method 0xa90a3d34.
//
// Solidity: function getTokensForWei(uint256 weiAmount) constant returns(uint256)
func (_Floico *FloicoCaller) GetTokensForWei(opts *bind.CallOpts, weiAmount *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Floico.contract.Call(opts, out, "getTokensForWei", weiAmount)
	return *ret0, err
}

// GetTokensForWei is a free data retrieval call binding the contract method 0xa90a3d34.
//
// Solidity: function getTokensForWei(uint256 weiAmount) constant returns(uint256)
func (_Floico *FloicoSession) GetTokensForWei(weiAmount *big.Int) (*big.Int, error) {
	return _Floico.Contract.GetTokensForWei(&_Floico.CallOpts, weiAmount)
}

// GetTokensForWei is a free data retrieval call binding the contract method 0xa90a3d34.
//
// Solidity: function getTokensForWei(uint256 weiAmount) constant returns(uint256)
func (_Floico *FloicoCallerSession) GetTokensForWei(weiAmount *big.Int) (*big.Int, error) {
	return _Floico.Contract.GetTokensForWei(&_Floico.CallOpts, weiAmount)
}

// GoalReached is a free data retrieval call binding the contract method 0x7d3d6522.
//
// Solidity: function goalReached() constant returns(bool)
func (_Floico *FloicoCaller) GoalReached(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Floico.contract.Call(opts, out, "goalReached")
	return *ret0, err
}

// GoalReached is a free data retrieval call binding the contract method 0x7d3d6522.
//
// Solidity: function goalReached() constant returns(bool)
func (_Floico *FloicoSession) GoalReached() (bool, error) {
	return _Floico.Contract.GoalReached(&_Floico.CallOpts)
}

// GoalReached is a free data retrieval call binding the contract method 0x7d3d6522.
//
// Solidity: function goalReached() constant returns(bool)
func (_Floico *FloicoCallerSession) GoalReached() (bool, error) {
	return _Floico.Contract.GoalReached(&_Floico.CallOpts)
}

// Info is a free data retrieval call binding the contract method 0x370158ea.
//
// Solidity: function info() constant returns(uint256, uint256, uint256, uint256, bool, bool, uint256[], uint256[], uint256)
func (_Floico *FloicoCaller) Info(opts *bind.CallOpts) (*big.Int, *big.Int, *big.Int, *big.Int, bool, bool, []*big.Int, []*big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
		ret2 = new(*big.Int)
		ret3 = new(*big.Int)
		ret4 = new(bool)
		ret5 = new(bool)
		ret6 = new([]*big.Int)
		ret7 = new([]*big.Int)
		ret8 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
		ret5,
		ret6,
		ret7,
		ret8,
	}
	err := _Floico.contract.Call(opts, out, "info")
	return *ret0, *ret1, *ret2, *ret3, *ret4, *ret5, *ret6, *ret7, *ret8, err
}

// Info is a free data retrieval call binding the contract method 0x370158ea.
//
// Solidity: function info() constant returns(uint256, uint256, uint256, uint256, bool, bool, uint256[], uint256[], uint256)
func (_Floico *FloicoSession) Info() (*big.Int, *big.Int, *big.Int, *big.Int, bool, bool, []*big.Int, []*big.Int, *big.Int, error) {
	return _Floico.Contract.Info(&_Floico.CallOpts)
}

// Info is a free data retrieval call binding the contract method 0x370158ea.
//
// Solidity: function info() constant returns(uint256, uint256, uint256, uint256, bool, bool, uint256[], uint256[], uint256)
func (_Floico *FloicoCallerSession) Info() (*big.Int, *big.Int, *big.Int, *big.Int, bool, bool, []*big.Int, []*big.Int, *big.Int, error) {
	return _Floico.Contract.Info(&_Floico.CallOpts)
}

// InitialTokens is a free data retrieval call binding the contract method 0x50bfeadc.
//
// Solidity: function initialTokens() constant returns(uint256)
func (_Floico *FloicoCaller) InitialTokens(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Floico.contract.Call(opts, out, "initialTokens")
	return *ret0, err
}

// InitialTokens is a free data retrieval call binding the contract method 0x50bfeadc.
//
// Solidity: function initialTokens() constant returns(uint256)
func (_Floico *FloicoSession) InitialTokens() (*big.Int, error) {
	return _Floico.Contract.InitialTokens(&_Floico.CallOpts)
}

// InitialTokens is a free data retrieval call binding the contract method 0x50bfeadc.
//
// Solidity: function initialTokens() constant returns(uint256)
func (_Floico *FloicoCallerSession) InitialTokens() (*big.Int, error) {
	return _Floico.Contract.InitialTokens(&_Floico.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() constant returns(bool)
func (_Floico *FloicoCaller) Initialized(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Floico.contract.Call(opts, out, "initialized")
	return *ret0, err
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() constant returns(bool)
func (_Floico *FloicoSession) Initialized() (bool, error) {
	return _Floico.Contract.Initialized(&_Floico.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() constant returns(bool)
func (_Floico *FloicoCallerSession) Initialized() (bool, error) {
	return _Floico.Contract.Initialized(&_Floico.CallOpts)
}

// IsActive is a free data retrieval call binding the contract method 0x22f3e2d4.
//
// Solidity: function isActive() constant returns(bool)
func (_Floico *FloicoCaller) IsActive(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Floico.contract.Call(opts, out, "isActive")
	return *ret0, err
}

// IsActive is a free data retrieval call binding the contract method 0x22f3e2d4.
//
// Solidity: function isActive() constant returns(bool)
func (_Floico *FloicoSession) IsActive() (bool, error) {
	return _Floico.Contract.IsActive(&_Floico.CallOpts)
}

// IsActive is a free data retrieval call binding the contract method 0x22f3e2d4.
//
// Solidity: function isActive() constant returns(bool)
func (_Floico *FloicoCallerSession) IsActive() (bool, error) {
	return _Floico.Contract.IsActive(&_Floico.CallOpts)
}

// LeftInActualPhase is a free data retrieval call binding the contract method 0x24434169.
//
// Solidity: function leftInActualPhase() constant returns(uint256 res)
func (_Floico *FloicoCaller) LeftInActualPhase(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Floico.contract.Call(opts, out, "leftInActualPhase")
	return *ret0, err
}

// LeftInActualPhase is a free data retrieval call binding the contract method 0x24434169.
//
// Solidity: function leftInActualPhase() constant returns(uint256 res)
func (_Floico *FloicoSession) LeftInActualPhase() (*big.Int, error) {
	return _Floico.Contract.LeftInActualPhase(&_Floico.CallOpts)
}

// LeftInActualPhase is a free data retrieval call binding the contract method 0x24434169.
//
// Solidity: function leftInActualPhase() constant returns(uint256 res)
func (_Floico *FloicoCallerSession) LeftInActualPhase() (*big.Int, error) {
	return _Floico.Contract.LeftInActualPhase(&_Floico.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Floico *FloicoCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Floico.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Floico *FloicoSession) Owner() (common.Address, error) {
	return _Floico.Contract.Owner(&_Floico.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Floico *FloicoCallerSession) Owner() (common.Address, error) {
	return _Floico.Contract.Owner(&_Floico.CallOpts)
}

// Participants is a free data retrieval call binding the contract method 0x35c1d349.
//
// Solidity: function participants(uint256 ) constant returns(bytes32 name, address addr)
func (_Floico *FloicoCaller) Participants(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Name [32]byte
	Addr common.Address
}, error) {
	ret := new(struct {
		Name [32]byte
		Addr common.Address
	})
	out := ret
	err := _Floico.contract.Call(opts, out, "participants", arg0)
	return *ret, err
}

// Participants is a free data retrieval call binding the contract method 0x35c1d349.
//
// Solidity: function participants(uint256 ) constant returns(bytes32 name, address addr)
func (_Floico *FloicoSession) Participants(arg0 *big.Int) (struct {
	Name [32]byte
	Addr common.Address
}, error) {
	return _Floico.Contract.Participants(&_Floico.CallOpts, arg0)
}

// Participants is a free data retrieval call binding the contract method 0x35c1d349.
//
// Solidity: function participants(uint256 ) constant returns(bytes32 name, address addr)
func (_Floico *FloicoCallerSession) Participants(arg0 *big.Int) (struct {
	Name [32]byte
	Addr common.Address
}, error) {
	return _Floico.Contract.Participants(&_Floico.CallOpts, arg0)
}

// ParticipantsLength is a free data retrieval call binding the contract method 0x4158506a.
//
// Solidity: function participantsLength() constant returns(uint256)
func (_Floico *FloicoCaller) ParticipantsLength(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Floico.contract.Call(opts, out, "participantsLength")
	return *ret0, err
}

// ParticipantsLength is a free data retrieval call binding the contract method 0x4158506a.
//
// Solidity: function participantsLength() constant returns(uint256)
func (_Floico *FloicoSession) ParticipantsLength() (*big.Int, error) {
	return _Floico.Contract.ParticipantsLength(&_Floico.CallOpts)
}

// ParticipantsLength is a free data retrieval call binding the contract method 0x4158506a.
//
// Solidity: function participantsLength() constant returns(uint256)
func (_Floico *FloicoCallerSession) ParticipantsLength() (*big.Int, error) {
	return _Floico.Contract.ParticipantsLength(&_Floico.CallOpts)
}

// Phase is a free data retrieval call binding the contract method 0xb1c9fe6e.
//
// Solidity: function phase() constant returns(uint256)
func (_Floico *FloicoCaller) Phase(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Floico.contract.Call(opts, out, "phase")
	return *ret0, err
}

// Phase is a free data retrieval call binding the contract method 0xb1c9fe6e.
//
// Solidity: function phase() constant returns(uint256)
func (_Floico *FloicoSession) Phase() (*big.Int, error) {
	return _Floico.Contract.Phase(&_Floico.CallOpts)
}

// Phase is a free data retrieval call binding the contract method 0xb1c9fe6e.
//
// Solidity: function phase() constant returns(uint256)
func (_Floico *FloicoCallerSession) Phase() (*big.Int, error) {
	return _Floico.Contract.Phase(&_Floico.CallOpts)
}

// PhaseInfo is a free data retrieval call binding the contract method 0x7be6a034.
//
// Solidity: function phaseInfo(uint256 , uint256 ) constant returns(uint256)
func (_Floico *FloicoCaller) PhaseInfo(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Floico.contract.Call(opts, out, "phaseInfo", arg0, arg1)
	return *ret0, err
}

// PhaseInfo is a free data retrieval call binding the contract method 0x7be6a034.
//
// Solidity: function phaseInfo(uint256 , uint256 ) constant returns(uint256)
func (_Floico *FloicoSession) PhaseInfo(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _Floico.Contract.PhaseInfo(&_Floico.CallOpts, arg0, arg1)
}

// PhaseInfo is a free data retrieval call binding the contract method 0x7be6a034.
//
// Solidity: function phaseInfo(uint256 , uint256 ) constant returns(uint256)
func (_Floico *FloicoCallerSession) PhaseInfo(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _Floico.Contract.PhaseInfo(&_Floico.CallOpts, arg0, arg1)
}

// PriceDiscount is a free data retrieval call binding the contract method 0xaa8930d2.
//
// Solidity: function priceDiscount() constant returns(uint256)
func (_Floico *FloicoCaller) PriceDiscount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Floico.contract.Call(opts, out, "priceDiscount")
	return *ret0, err
}

// PriceDiscount is a free data retrieval call binding the contract method 0xaa8930d2.
//
// Solidity: function priceDiscount() constant returns(uint256)
func (_Floico *FloicoSession) PriceDiscount() (*big.Int, error) {
	return _Floico.Contract.PriceDiscount(&_Floico.CallOpts)
}

// PriceDiscount is a free data retrieval call binding the contract method 0xaa8930d2.
//
// Solidity: function priceDiscount() constant returns(uint256)
func (_Floico *FloicoCallerSession) PriceDiscount() (*big.Int, error) {
	return _Floico.Contract.PriceDiscount(&_Floico.CallOpts)
}

// Proposals is a free data retrieval call binding the contract method 0x013cf08b.
//
// Solidity: function proposals(uint256 ) constant returns(bytes32 description, uint256 endTime, uint8 result, uint8 actionType, address actionAddress, bytes32 name, uint256 amount)
func (_Floico *FloicoCaller) Proposals(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Description   [32]byte
	EndTime       *big.Int
	Result        uint8
	ActionType    uint8
	ActionAddress common.Address
	Name          [32]byte
	Amount        *big.Int
}, error) {
	ret := new(struct {
		Description   [32]byte
		EndTime       *big.Int
		Result        uint8
		ActionType    uint8
		ActionAddress common.Address
		Name          [32]byte
		Amount        *big.Int
	})
	out := ret
	err := _Floico.contract.Call(opts, out, "proposals", arg0)
	return *ret, err
}

// Proposals is a free data retrieval call binding the contract method 0x013cf08b.
//
// Solidity: function proposals(uint256 ) constant returns(bytes32 description, uint256 endTime, uint8 result, uint8 actionType, address actionAddress, bytes32 name, uint256 amount)
func (_Floico *FloicoSession) Proposals(arg0 *big.Int) (struct {
	Description   [32]byte
	EndTime       *big.Int
	Result        uint8
	ActionType    uint8
	ActionAddress common.Address
	Name          [32]byte
	Amount        *big.Int
}, error) {
	return _Floico.Contract.Proposals(&_Floico.CallOpts, arg0)
}

// Proposals is a free data retrieval call binding the contract method 0x013cf08b.
//
// Solidity: function proposals(uint256 ) constant returns(bytes32 description, uint256 endTime, uint8 result, uint8 actionType, address actionAddress, bytes32 name, uint256 amount)
func (_Floico *FloicoCallerSession) Proposals(arg0 *big.Int) (struct {
	Description   [32]byte
	EndTime       *big.Int
	Result        uint8
	ActionType    uint8
	ActionAddress common.Address
	Name          [32]byte
	Amount        *big.Int
}, error) {
	return _Floico.Contract.Proposals(&_Floico.CallOpts, arg0)
}

// ProposalsLength is a free data retrieval call binding the contract method 0x44c7c867.
//
// Solidity: function proposalsLength() constant returns(uint256)
func (_Floico *FloicoCaller) ProposalsLength(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Floico.contract.Call(opts, out, "proposalsLength")
	return *ret0, err
}

// ProposalsLength is a free data retrieval call binding the contract method 0x44c7c867.
//
// Solidity: function proposalsLength() constant returns(uint256)
func (_Floico *FloicoSession) ProposalsLength() (*big.Int, error) {
	return _Floico.Contract.ProposalsLength(&_Floico.CallOpts)
}

// ProposalsLength is a free data retrieval call binding the contract method 0x44c7c867.
//
// Solidity: function proposalsLength() constant returns(uint256)
func (_Floico *FloicoCallerSession) ProposalsLength() (*big.Int, error) {
	return _Floico.Contract.ProposalsLength(&_Floico.CallOpts)
}

// RaisedAmount is a free data retrieval call binding the contract method 0xc59ee1dc.
//
// Solidity: function raisedAmount() constant returns(uint256)
func (_Floico *FloicoCaller) RaisedAmount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Floico.contract.Call(opts, out, "raisedAmount")
	return *ret0, err
}

// RaisedAmount is a free data retrieval call binding the contract method 0xc59ee1dc.
//
// Solidity: function raisedAmount() constant returns(uint256)
func (_Floico *FloicoSession) RaisedAmount() (*big.Int, error) {
	return _Floico.Contract.RaisedAmount(&_Floico.CallOpts)
}

// RaisedAmount is a free data retrieval call binding the contract method 0xc59ee1dc.
//
// Solidity: function raisedAmount() constant returns(uint256)
func (_Floico *FloicoCallerSession) RaisedAmount() (*big.Int, error) {
	return _Floico.Contract.RaisedAmount(&_Floico.CallOpts)
}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() constant returns(uint256)
func (_Floico *FloicoCaller) StartTime(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Floico.contract.Call(opts, out, "startTime")
	return *ret0, err
}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() constant returns(uint256)
func (_Floico *FloicoSession) StartTime() (*big.Int, error) {
	return _Floico.Contract.StartTime(&_Floico.CallOpts)
}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() constant returns(uint256)
func (_Floico *FloicoCallerSession) StartTime() (*big.Int, error) {
	return _Floico.Contract.StartTime(&_Floico.CallOpts)
}

// StatusOfProposal is a free data retrieval call binding the contract method 0xbfd6e79f.
//
// Solidity: function statusOfProposal(uint256 index) constant returns(address[], bytes32[], uint8[])
func (_Floico *FloicoCaller) StatusOfProposal(opts *bind.CallOpts, index *big.Int) ([]common.Address, [][32]byte, []uint8, error) {
	var (
		ret0 = new([]common.Address)
		ret1 = new([][32]byte)
		ret2 = new([]uint8)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _Floico.contract.Call(opts, out, "statusOfProposal", index)
	return *ret0, *ret1, *ret2, err
}

// StatusOfProposal is a free data retrieval call binding the contract method 0xbfd6e79f.
//
// Solidity: function statusOfProposal(uint256 index) constant returns(address[], bytes32[], uint8[])
func (_Floico *FloicoSession) StatusOfProposal(index *big.Int) ([]common.Address, [][32]byte, []uint8, error) {
	return _Floico.Contract.StatusOfProposal(&_Floico.CallOpts, index)
}

// StatusOfProposal is a free data retrieval call binding the contract method 0xbfd6e79f.
//
// Solidity: function statusOfProposal(uint256 index) constant returns(address[], bytes32[], uint8[])
func (_Floico *FloicoCallerSession) StatusOfProposal(index *big.Int) ([]common.Address, [][32]byte, []uint8, error) {
	return _Floico.Contract.StatusOfProposal(&_Floico.CallOpts, index)
}

// TokensAvailable is a free data retrieval call binding the contract method 0x60659a92.
//
// Solidity: function tokensAvailable() constant returns(uint256)
func (_Floico *FloicoCaller) TokensAvailable(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Floico.contract.Call(opts, out, "tokensAvailable")
	return *ret0, err
}

// TokensAvailable is a free data retrieval call binding the contract method 0x60659a92.
//
// Solidity: function tokensAvailable() constant returns(uint256)
func (_Floico *FloicoSession) TokensAvailable() (*big.Int, error) {
	return _Floico.Contract.TokensAvailable(&_Floico.CallOpts)
}

// TokensAvailable is a free data retrieval call binding the contract method 0x60659a92.
//
// Solidity: function tokensAvailable() constant returns(uint256)
func (_Floico *FloicoCallerSession) TokensAvailable() (*big.Int, error) {
	return _Floico.Contract.TokensAvailable(&_Floico.CallOpts)
}

// WeiAailable is a free data retrieval call binding the contract method 0x0e222001.
//
// Solidity: function weiAailable() constant returns(uint256)
func (_Floico *FloicoCaller) WeiAailable(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Floico.contract.Call(opts, out, "weiAailable")
	return *ret0, err
}

// WeiAailable is a free data retrieval call binding the contract method 0x0e222001.
//
// Solidity: function weiAailable() constant returns(uint256)
func (_Floico *FloicoSession) WeiAailable() (*big.Int, error) {
	return _Floico.Contract.WeiAailable(&_Floico.CallOpts)
}

// WeiAailable is a free data retrieval call binding the contract method 0x0e222001.
//
// Solidity: function weiAailable() constant returns(uint256)
func (_Floico *FloicoCallerSession) WeiAailable() (*big.Int, error) {
	return _Floico.Contract.WeiAailable(&_Floico.CallOpts)
}

// BuyOverBackend is a paid mutator transaction binding the contract method 0xcb118a72.
//
// Solidity: function buyOverBackend(address buyer, uint256 tokens) returns(bool res)
func (_Floico *FloicoTransactor) BuyOverBackend(opts *bind.TransactOpts, buyer common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Floico.contract.Transact(opts, "buyOverBackend", buyer, tokens)
}

// BuyOverBackend is a paid mutator transaction binding the contract method 0xcb118a72.
//
// Solidity: function buyOverBackend(address buyer, uint256 tokens) returns(bool res)
func (_Floico *FloicoSession) BuyOverBackend(buyer common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Floico.Contract.BuyOverBackend(&_Floico.TransactOpts, buyer, tokens)
}

// BuyOverBackend is a paid mutator transaction binding the contract method 0xcb118a72.
//
// Solidity: function buyOverBackend(address buyer, uint256 tokens) returns(bool res)
func (_Floico *FloicoTransactorSession) BuyOverBackend(buyer common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Floico.Contract.BuyOverBackend(&_Floico.TransactOpts, buyer, tokens)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xd0febe4c.
//
// Solidity: function buyTokens() returns()
func (_Floico *FloicoTransactor) BuyTokens(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Floico.contract.Transact(opts, "buyTokens")
}

// BuyTokens is a paid mutator transaction binding the contract method 0xd0febe4c.
//
// Solidity: function buyTokens() returns()
func (_Floico *FloicoSession) BuyTokens() (*types.Transaction, error) {
	return _Floico.Contract.BuyTokens(&_Floico.TransactOpts)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xd0febe4c.
//
// Solidity: function buyTokens() returns()
func (_Floico *FloicoTransactorSession) BuyTokens() (*types.Transaction, error) {
	return _Floico.Contract.BuyTokens(&_Floico.TransactOpts)
}

// CreateProposal is a paid mutator transaction binding the contract method 0x048bc9ec.
//
// Solidity: function createProposal(bytes32 _description, uint256 _durationHours, uint8 _actionType, address _actionAddress, bytes32 _name, uint256 _amount) returns()
func (_Floico *FloicoTransactor) CreateProposal(opts *bind.TransactOpts, _description [32]byte, _durationHours *big.Int, _actionType uint8, _actionAddress common.Address, _name [32]byte, _amount *big.Int) (*types.Transaction, error) {
	return _Floico.contract.Transact(opts, "createProposal", _description, _durationHours, _actionType, _actionAddress, _name, _amount)
}

// CreateProposal is a paid mutator transaction binding the contract method 0x048bc9ec.
//
// Solidity: function createProposal(bytes32 _description, uint256 _durationHours, uint8 _actionType, address _actionAddress, bytes32 _name, uint256 _amount) returns()
func (_Floico *FloicoSession) CreateProposal(_description [32]byte, _durationHours *big.Int, _actionType uint8, _actionAddress common.Address, _name [32]byte, _amount *big.Int) (*types.Transaction, error) {
	return _Floico.Contract.CreateProposal(&_Floico.TransactOpts, _description, _durationHours, _actionType, _actionAddress, _name, _amount)
}

// CreateProposal is a paid mutator transaction binding the contract method 0x048bc9ec.
//
// Solidity: function createProposal(bytes32 _description, uint256 _durationHours, uint8 _actionType, address _actionAddress, bytes32 _name, uint256 _amount) returns()
func (_Floico *FloicoTransactorSession) CreateProposal(_description [32]byte, _durationHours *big.Int, _actionType uint8, _actionAddress common.Address, _name [32]byte, _amount *big.Int) (*types.Transaction, error) {
	return _Floico.Contract.CreateProposal(&_Floico.TransactOpts, _description, _durationHours, _actionType, _actionAddress, _name, _amount)
}

// FinishProposal is a paid mutator transaction binding the contract method 0x27393278.
//
// Solidity: function finishProposal(uint256 proposalIndex) returns()
func (_Floico *FloicoTransactor) FinishProposal(opts *bind.TransactOpts, proposalIndex *big.Int) (*types.Transaction, error) {
	return _Floico.contract.Transact(opts, "finishProposal", proposalIndex)
}

// FinishProposal is a paid mutator transaction binding the contract method 0x27393278.
//
// Solidity: function finishProposal(uint256 proposalIndex) returns()
func (_Floico *FloicoSession) FinishProposal(proposalIndex *big.Int) (*types.Transaction, error) {
	return _Floico.Contract.FinishProposal(&_Floico.TransactOpts, proposalIndex)
}

// FinishProposal is a paid mutator transaction binding the contract method 0x27393278.
//
// Solidity: function finishProposal(uint256 proposalIndex) returns()
func (_Floico *FloicoTransactorSession) FinishProposal(proposalIndex *big.Int) (*types.Transaction, error) {
	return _Floico.Contract.FinishProposal(&_Floico.TransactOpts, proposalIndex)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address transport) returns()
func (_Floico *FloicoTransactor) Initialize(opts *bind.TransactOpts, transport common.Address) (*types.Transaction, error) {
	return _Floico.contract.Transact(opts, "initialize", transport)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address transport) returns()
func (_Floico *FloicoSession) Initialize(transport common.Address) (*types.Transaction, error) {
	return _Floico.Contract.Initialize(&_Floico.TransactOpts, transport)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address transport) returns()
func (_Floico *FloicoTransactorSession) Initialize(transport common.Address) (*types.Transaction, error) {
	return _Floico.Contract.Initialize(&_Floico.TransactOpts, transport)
}

// SetDiscount is a paid mutator transaction binding the contract method 0xdabd2719.
//
// Solidity: function setDiscount(uint256 discount) returns()
func (_Floico *FloicoTransactor) SetDiscount(opts *bind.TransactOpts, discount *big.Int) (*types.Transaction, error) {
	return _Floico.contract.Transact(opts, "setDiscount", discount)
}

// SetDiscount is a paid mutator transaction binding the contract method 0xdabd2719.
//
// Solidity: function setDiscount(uint256 discount) returns()
func (_Floico *FloicoSession) SetDiscount(discount *big.Int) (*types.Transaction, error) {
	return _Floico.Contract.SetDiscount(&_Floico.TransactOpts, discount)
}

// SetDiscount is a paid mutator transaction binding the contract method 0xdabd2719.
//
// Solidity: function setDiscount(uint256 discount) returns()
func (_Floico *FloicoTransactorSession) SetDiscount(discount *big.Int) (*types.Transaction, error) {
	return _Floico.Contract.SetDiscount(&_Floico.TransactOpts, discount)
}

// SetEnabled is a paid mutator transaction binding the contract method 0x328d8f72.
//
// Solidity: function setEnabled(bool _enabled) returns(bool)
func (_Floico *FloicoTransactor) SetEnabled(opts *bind.TransactOpts, _enabled bool) (*types.Transaction, error) {
	return _Floico.contract.Transact(opts, "setEnabled", _enabled)
}

// SetEnabled is a paid mutator transaction binding the contract method 0x328d8f72.
//
// Solidity: function setEnabled(bool _enabled) returns(bool)
func (_Floico *FloicoSession) SetEnabled(_enabled bool) (*types.Transaction, error) {
	return _Floico.Contract.SetEnabled(&_Floico.TransactOpts, _enabled)
}

// SetEnabled is a paid mutator transaction binding the contract method 0x328d8f72.
//
// Solidity: function setEnabled(bool _enabled) returns(bool)
func (_Floico *FloicoTransactorSession) SetEnabled(_enabled bool) (*types.Transaction, error) {
	return _Floico.Contract.SetEnabled(&_Floico.TransactOpts, _enabled)
}

// TokenFallback is a paid mutator transaction binding the contract method 0xc0ee0b8a.
//
// Solidity: function tokenFallback(address _from, uint256 _value, bytes _data) returns()
func (_Floico *FloicoTransactor) TokenFallback(opts *bind.TransactOpts, _from common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _Floico.contract.Transact(opts, "tokenFallback", _from, _value, _data)
}

// TokenFallback is a paid mutator transaction binding the contract method 0xc0ee0b8a.
//
// Solidity: function tokenFallback(address _from, uint256 _value, bytes _data) returns()
func (_Floico *FloicoSession) TokenFallback(_from common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _Floico.Contract.TokenFallback(&_Floico.TransactOpts, _from, _value, _data)
}

// TokenFallback is a paid mutator transaction binding the contract method 0xc0ee0b8a.
//
// Solidity: function tokenFallback(address _from, uint256 _value, bytes _data) returns()
func (_Floico *FloicoTransactorSession) TokenFallback(_from common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _Floico.Contract.TokenFallback(&_Floico.TransactOpts, _from, _value, _data)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Floico *FloicoTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Floico.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Floico *FloicoSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Floico.Contract.TransferOwnership(&_Floico.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Floico *FloicoTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Floico.Contract.TransferOwnership(&_Floico.TransactOpts, newOwner)
}

// VoteInProposal is a paid mutator transaction binding the contract method 0xd035e72f.
//
// Solidity: function voteInProposal(uint256 proposalIndex, uint8 vote) returns()
func (_Floico *FloicoTransactor) VoteInProposal(opts *bind.TransactOpts, proposalIndex *big.Int, vote uint8) (*types.Transaction, error) {
	return _Floico.contract.Transact(opts, "voteInProposal", proposalIndex, vote)
}

// VoteInProposal is a paid mutator transaction binding the contract method 0xd035e72f.
//
// Solidity: function voteInProposal(uint256 proposalIndex, uint8 vote) returns()
func (_Floico *FloicoSession) VoteInProposal(proposalIndex *big.Int, vote uint8) (*types.Transaction, error) {
	return _Floico.Contract.VoteInProposal(&_Floico.TransactOpts, proposalIndex, vote)
}

// VoteInProposal is a paid mutator transaction binding the contract method 0xd035e72f.
//
// Solidity: function voteInProposal(uint256 proposalIndex, uint8 vote) returns()
func (_Floico *FloicoTransactorSession) VoteInProposal(proposalIndex *big.Int, vote uint8) (*types.Transaction, error) {
	return _Floico.Contract.VoteInProposal(&_Floico.TransactOpts, proposalIndex, vote)
}

// FloicoBoughtTokensIterator is returned from FilterBoughtTokens and is used to iterate over the raw logs and unpacked data for BoughtTokens events raised by the Floico contract.
type FloicoBoughtTokensIterator struct {
	Event *FloicoBoughtTokens // Event containing the contract specifics and raw log

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
func (it *FloicoBoughtTokensIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FloicoBoughtTokens)
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
		it.Event = new(FloicoBoughtTokens)
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
func (it *FloicoBoughtTokensIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FloicoBoughtTokensIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FloicoBoughtTokens represents a BoughtTokens event raised by the Floico contract.
type FloicoBoughtTokens struct {
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterBoughtTokens is a free log retrieval operation binding the contract event 0x61b2357f75eed32a19939598e4c7563879bacd78e9d957c31f0f9f70a3fd14a0.
//
// Solidity: event BoughtTokens(address indexed to, uint256 value)
func (_Floico *FloicoFilterer) FilterBoughtTokens(opts *bind.FilterOpts, to []common.Address) (*FloicoBoughtTokensIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Floico.contract.FilterLogs(opts, "BoughtTokens", toRule)
	if err != nil {
		return nil, err
	}
	return &FloicoBoughtTokensIterator{contract: _Floico.contract, event: "BoughtTokens", logs: logs, sub: sub}, nil
}

// WatchBoughtTokens is a free log subscription operation binding the contract event 0x61b2357f75eed32a19939598e4c7563879bacd78e9d957c31f0f9f70a3fd14a0.
//
// Solidity: event BoughtTokens(address indexed to, uint256 value)
func (_Floico *FloicoFilterer) WatchBoughtTokens(opts *bind.WatchOpts, sink chan<- *FloicoBoughtTokens, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Floico.contract.WatchLogs(opts, "BoughtTokens", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FloicoBoughtTokens)
				if err := _Floico.contract.UnpackLog(event, "BoughtTokens", log); err != nil {
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

// ParseBoughtTokens is a log parse operation binding the contract event 0x61b2357f75eed32a19939598e4c7563879bacd78e9d957c31f0f9f70a3fd14a0.
//
// Solidity: event BoughtTokens(address indexed to, uint256 value)
func (_Floico *FloicoFilterer) ParseBoughtTokens(log types.Log) (*FloicoBoughtTokens, error) {
	event := new(FloicoBoughtTokens)
	if err := _Floico.contract.UnpackLog(event, "BoughtTokens", log); err != nil {
		return nil, err
	}
	return event, nil
}

// FloicoFinishVotingIterator is returned from FilterFinishVoting and is used to iterate over the raw logs and unpacked data for FinishVoting events raised by the Floico contract.
type FloicoFinishVotingIterator struct {
	Event *FloicoFinishVoting // Event containing the contract specifics and raw log

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
func (it *FloicoFinishVotingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FloicoFinishVoting)
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
		it.Event = new(FloicoFinishVoting)
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
func (it *FloicoFinishVotingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FloicoFinishVotingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FloicoFinishVoting represents a FinishVoting event raised by the Floico contract.
type FloicoFinishVoting struct {
	Description   [32]byte
	Result        bool
	ProposalIndex *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterFinishVoting is a free log retrieval operation binding the contract event 0xb89b889fb44aa067abcdcd537d53a8d8e84eb6ffa00368dd937318c80d77e75e.
//
// Solidity: event FinishVoting(bytes32 description, bool result, uint256 proposalIndex)
func (_Floico *FloicoFilterer) FilterFinishVoting(opts *bind.FilterOpts) (*FloicoFinishVotingIterator, error) {

	logs, sub, err := _Floico.contract.FilterLogs(opts, "FinishVoting")
	if err != nil {
		return nil, err
	}
	return &FloicoFinishVotingIterator{contract: _Floico.contract, event: "FinishVoting", logs: logs, sub: sub}, nil
}

// WatchFinishVoting is a free log subscription operation binding the contract event 0xb89b889fb44aa067abcdcd537d53a8d8e84eb6ffa00368dd937318c80d77e75e.
//
// Solidity: event FinishVoting(bytes32 description, bool result, uint256 proposalIndex)
func (_Floico *FloicoFilterer) WatchFinishVoting(opts *bind.WatchOpts, sink chan<- *FloicoFinishVoting) (event.Subscription, error) {

	logs, sub, err := _Floico.contract.WatchLogs(opts, "FinishVoting")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FloicoFinishVoting)
				if err := _Floico.contract.UnpackLog(event, "FinishVoting", log); err != nil {
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

// ParseFinishVoting is a log parse operation binding the contract event 0xb89b889fb44aa067abcdcd537d53a8d8e84eb6ffa00368dd937318c80d77e75e.
//
// Solidity: event FinishVoting(bytes32 description, bool result, uint256 proposalIndex)
func (_Floico *FloicoFilterer) ParseFinishVoting(log types.Log) (*FloicoFinishVoting, error) {
	event := new(FloicoFinishVoting)
	if err := _Floico.contract.UnpackLog(event, "FinishVoting", log); err != nil {
		return nil, err
	}
	return event, nil
}

// FloicoProposalCreatedIterator is returned from FilterProposalCreated and is used to iterate over the raw logs and unpacked data for ProposalCreated events raised by the Floico contract.
type FloicoProposalCreatedIterator struct {
	Event *FloicoProposalCreated // Event containing the contract specifics and raw log

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
func (it *FloicoProposalCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FloicoProposalCreated)
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
		it.Event = new(FloicoProposalCreated)
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
func (it *FloicoProposalCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FloicoProposalCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FloicoProposalCreated represents a ProposalCreated event raised by the Floico contract.
type FloicoProposalCreated struct {
	Description   [32]byte
	EndTime       *big.Int
	ActionType    uint8
	ActionAddress common.Address
	Name          [32]byte
	Amount        *big.Int
	ProposalIndex *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterProposalCreated is a free log retrieval operation binding the contract event 0x5a22095404761c56fdc0a2e7cf9111a91ff494b817685d5b0c34ddc0f134c50e.
//
// Solidity: event ProposalCreated(bytes32 description, uint256 endTime, uint8 actionType, address actionAddress, bytes32 name, uint256 amount, uint256 proposalIndex)
func (_Floico *FloicoFilterer) FilterProposalCreated(opts *bind.FilterOpts) (*FloicoProposalCreatedIterator, error) {

	logs, sub, err := _Floico.contract.FilterLogs(opts, "ProposalCreated")
	if err != nil {
		return nil, err
	}
	return &FloicoProposalCreatedIterator{contract: _Floico.contract, event: "ProposalCreated", logs: logs, sub: sub}, nil
}

// WatchProposalCreated is a free log subscription operation binding the contract event 0x5a22095404761c56fdc0a2e7cf9111a91ff494b817685d5b0c34ddc0f134c50e.
//
// Solidity: event ProposalCreated(bytes32 description, uint256 endTime, uint8 actionType, address actionAddress, bytes32 name, uint256 amount, uint256 proposalIndex)
func (_Floico *FloicoFilterer) WatchProposalCreated(opts *bind.WatchOpts, sink chan<- *FloicoProposalCreated) (event.Subscription, error) {

	logs, sub, err := _Floico.contract.WatchLogs(opts, "ProposalCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FloicoProposalCreated)
				if err := _Floico.contract.UnpackLog(event, "ProposalCreated", log); err != nil {
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

// ParseProposalCreated is a log parse operation binding the contract event 0x5a22095404761c56fdc0a2e7cf9111a91ff494b817685d5b0c34ddc0f134c50e.
//
// Solidity: event ProposalCreated(bytes32 description, uint256 endTime, uint8 actionType, address actionAddress, bytes32 name, uint256 amount, uint256 proposalIndex)
func (_Floico *FloicoFilterer) ParseProposalCreated(log types.Log) (*FloicoProposalCreated, error) {
	event := new(FloicoProposalCreated)
	if err := _Floico.contract.UnpackLog(event, "ProposalCreated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// FloicoVoteIterator is returned from FilterVote and is used to iterate over the raw logs and unpacked data for Vote events raised by the Floico contract.
type FloicoVoteIterator struct {
	Event *FloicoVote // Event containing the contract specifics and raw log

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
func (it *FloicoVoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FloicoVote)
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
		it.Event = new(FloicoVote)
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
func (it *FloicoVoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FloicoVoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FloicoVote represents a Vote event raised by the Floico contract.
type FloicoVote struct {
	Description   [32]byte
	ProposalIndex *big.Int
	Addr          common.Address
	Vote          uint8
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterVote is a free log retrieval operation binding the contract event 0x09859833479d3666edeeab8dc08198cf10941655b46111c955f390c6610a9833.
//
// Solidity: event Vote(bytes32 description, uint256 proposalIndex, address addr, uint8 vote)
func (_Floico *FloicoFilterer) FilterVote(opts *bind.FilterOpts) (*FloicoVoteIterator, error) {

	logs, sub, err := _Floico.contract.FilterLogs(opts, "Vote")
	if err != nil {
		return nil, err
	}
	return &FloicoVoteIterator{contract: _Floico.contract, event: "Vote", logs: logs, sub: sub}, nil
}

// WatchVote is a free log subscription operation binding the contract event 0x09859833479d3666edeeab8dc08198cf10941655b46111c955f390c6610a9833.
//
// Solidity: event Vote(bytes32 description, uint256 proposalIndex, address addr, uint8 vote)
func (_Floico *FloicoFilterer) WatchVote(opts *bind.WatchOpts, sink chan<- *FloicoVote) (event.Subscription, error) {

	logs, sub, err := _Floico.contract.WatchLogs(opts, "Vote")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FloicoVote)
				if err := _Floico.contract.UnpackLog(event, "Vote", log); err != nil {
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

// ParseVote is a log parse operation binding the contract event 0x09859833479d3666edeeab8dc08198cf10941655b46111c955f390c6610a9833.
//
// Solidity: event Vote(bytes32 description, uint256 proposalIndex, address addr, uint8 vote)
func (_Floico *FloicoFilterer) ParseVote(log types.Log) (*FloicoVote, error) {
	event := new(FloicoVote)
	if err := _Floico.contract.UnpackLog(event, "Vote", log); err != nil {
		return nil, err
	}
	return event, nil
}
