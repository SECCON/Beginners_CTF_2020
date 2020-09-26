// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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

// C4BABI is the input ABI used to generate the binding from.
const C4BABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_player\",\"type\":\"address\"},{\"internalType\":\"bytes8\",\"name\":\"_password\",\"type\":\"bytes8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"player\",\"type\":\"address\"}],\"name\":\"CheckPassed\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes8\",\"name\":\"_password\",\"type\":\"bytes8\"},{\"internalType\":\"uint256\",\"name\":\"pin\",\"type\":\"uint256\"}],\"name\":\"check\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"player\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"success\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// C4BFuncSigs maps the 4-byte function signature to its string representation.
var C4BFuncSigs = map[string]string{
	"c577930a": "check(bytes8,uint256)",
	"48db5f89": "player()",
	"0b93381b": "success()",
}

// C4BBin is the compiled bytecode used for deploying new contracts.
var C4BBin = "0x608060405234801561001057600080fd5b506040516102703803806102708339818101604052604081101561003357600080fd5b508051602090910151600080546001600160a01b0319166001600160a01b0390931692909217600160a01b600160e01b031916600160a01b60c09290921c919091021760ff60e01b191681556101e190819061008f90396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80630b93381b1461004657806348db5f8914610062578063c577930a14610086575b600080fd5b61004e6100b3565b604080519115158252519081900360200190f35b61006a6100c3565b604080516001600160a01b039092168252519081900360200190f35b61004e6004803603604081101561009c57600080fd5b506001600160c01b031981351690602001356100d2565b600054600160e01b900460ff1681565b6000546001600160a01b031681565b600060001943014082620f42408206141561016057604080516001600160c01b03198087166020808401919091528351808403820181528385018552805190820120600054600160a01b900460c01b909216606080850191909152845180850390910181526080909301909352815191909201201415610160576000805460ff60e01b1916600160e01b1790555b600080546040516001600160a01b03909116917f320f420edbd58b0816e3c933b93878e7c130093cdc9237d2e3a855b724f69c9491a25050600054600160e01b900460ff169291505056fea2646970667358221220aeab6d521c7fa54a91ba511baec1e6bf53f8d23d43aaf2b13d06602f4302efa664736f6c63430006060033"

// DeployC4B deploys a new Ethereum contract, binding an instance of C4B to it.
func DeployC4B(auth *bind.TransactOpts, backend bind.ContractBackend, _player common.Address, _password [8]byte) (common.Address, *types.Transaction, *C4B, error) {
	parsed, err := abi.JSON(strings.NewReader(C4BABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(C4BBin), backend, _player, _password)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &C4B{C4BCaller: C4BCaller{contract: contract}, C4BTransactor: C4BTransactor{contract: contract}, C4BFilterer: C4BFilterer{contract: contract}}, nil
}

// C4B is an auto generated Go binding around an Ethereum contract.
type C4B struct {
	C4BCaller     // Read-only binding to the contract
	C4BTransactor // Write-only binding to the contract
	C4BFilterer   // Log filterer for contract events
}

// C4BCaller is an auto generated read-only Go binding around an Ethereum contract.
type C4BCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// C4BTransactor is an auto generated write-only Go binding around an Ethereum contract.
type C4BTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// C4BFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type C4BFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// C4BSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type C4BSession struct {
	Contract     *C4B              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// C4BCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type C4BCallerSession struct {
	Contract *C4BCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// C4BTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type C4BTransactorSession struct {
	Contract     *C4BTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// C4BRaw is an auto generated low-level Go binding around an Ethereum contract.
type C4BRaw struct {
	Contract *C4B // Generic contract binding to access the raw methods on
}

// C4BCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type C4BCallerRaw struct {
	Contract *C4BCaller // Generic read-only contract binding to access the raw methods on
}

// C4BTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type C4BTransactorRaw struct {
	Contract *C4BTransactor // Generic write-only contract binding to access the raw methods on
}

// NewC4B creates a new instance of C4B, bound to a specific deployed contract.
func NewC4B(address common.Address, backend bind.ContractBackend) (*C4B, error) {
	contract, err := bindC4B(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &C4B{C4BCaller: C4BCaller{contract: contract}, C4BTransactor: C4BTransactor{contract: contract}, C4BFilterer: C4BFilterer{contract: contract}}, nil
}

// NewC4BCaller creates a new read-only instance of C4B, bound to a specific deployed contract.
func NewC4BCaller(address common.Address, caller bind.ContractCaller) (*C4BCaller, error) {
	contract, err := bindC4B(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &C4BCaller{contract: contract}, nil
}

// NewC4BTransactor creates a new write-only instance of C4B, bound to a specific deployed contract.
func NewC4BTransactor(address common.Address, transactor bind.ContractTransactor) (*C4BTransactor, error) {
	contract, err := bindC4B(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &C4BTransactor{contract: contract}, nil
}

// NewC4BFilterer creates a new log filterer instance of C4B, bound to a specific deployed contract.
func NewC4BFilterer(address common.Address, filterer bind.ContractFilterer) (*C4BFilterer, error) {
	contract, err := bindC4B(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &C4BFilterer{contract: contract}, nil
}

// bindC4B binds a generic wrapper to an already deployed contract.
func bindC4B(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(C4BABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_C4B *C4BRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _C4B.Contract.C4BCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_C4B *C4BRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _C4B.Contract.C4BTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_C4B *C4BRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _C4B.Contract.C4BTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_C4B *C4BCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _C4B.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_C4B *C4BTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _C4B.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_C4B *C4BTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _C4B.Contract.contract.Transact(opts, method, params...)
}

// Player is a free data retrieval call binding the contract method 0x48db5f89.
//
// Solidity: function player() view returns(address)
func (_C4B *C4BCaller) Player(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _C4B.contract.Call(opts, out, "player")
	return *ret0, err
}

// Player is a free data retrieval call binding the contract method 0x48db5f89.
//
// Solidity: function player() view returns(address)
func (_C4B *C4BSession) Player() (common.Address, error) {
	return _C4B.Contract.Player(&_C4B.CallOpts)
}

// Player is a free data retrieval call binding the contract method 0x48db5f89.
//
// Solidity: function player() view returns(address)
func (_C4B *C4BCallerSession) Player() (common.Address, error) {
	return _C4B.Contract.Player(&_C4B.CallOpts)
}

// Success is a free data retrieval call binding the contract method 0x0b93381b.
//
// Solidity: function success() view returns(bool)
func (_C4B *C4BCaller) Success(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _C4B.contract.Call(opts, out, "success")
	return *ret0, err
}

// Success is a free data retrieval call binding the contract method 0x0b93381b.
//
// Solidity: function success() view returns(bool)
func (_C4B *C4BSession) Success() (bool, error) {
	return _C4B.Contract.Success(&_C4B.CallOpts)
}

// Success is a free data retrieval call binding the contract method 0x0b93381b.
//
// Solidity: function success() view returns(bool)
func (_C4B *C4BCallerSession) Success() (bool, error) {
	return _C4B.Contract.Success(&_C4B.CallOpts)
}

// Check is a paid mutator transaction binding the contract method 0xc577930a.
//
// Solidity: function check(bytes8 _password, uint256 pin) returns(bool)
func (_C4B *C4BTransactor) Check(opts *bind.TransactOpts, _password [8]byte, pin *big.Int) (*types.Transaction, error) {
	return _C4B.contract.Transact(opts, "check", _password, pin)
}

// Check is a paid mutator transaction binding the contract method 0xc577930a.
//
// Solidity: function check(bytes8 _password, uint256 pin) returns(bool)
func (_C4B *C4BSession) Check(_password [8]byte, pin *big.Int) (*types.Transaction, error) {
	return _C4B.Contract.Check(&_C4B.TransactOpts, _password, pin)
}

// Check is a paid mutator transaction binding the contract method 0xc577930a.
//
// Solidity: function check(bytes8 _password, uint256 pin) returns(bool)
func (_C4B *C4BTransactorSession) Check(_password [8]byte, pin *big.Int) (*types.Transaction, error) {
	return _C4B.Contract.Check(&_C4B.TransactOpts, _password, pin)
}

// C4BCheckPassedIterator is returned from FilterCheckPassed and is used to iterate over the raw logs and unpacked data for CheckPassed events raised by the C4B contract.
type C4BCheckPassedIterator struct {
	Event *C4BCheckPassed // Event containing the contract specifics and raw log

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
func (it *C4BCheckPassedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(C4BCheckPassed)
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
		it.Event = new(C4BCheckPassed)
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
func (it *C4BCheckPassedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *C4BCheckPassedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// C4BCheckPassed represents a CheckPassed event raised by the C4B contract.
type C4BCheckPassed struct {
	Player common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCheckPassed is a free log retrieval operation binding the contract event 0x320f420edbd58b0816e3c933b93878e7c130093cdc9237d2e3a855b724f69c94.
//
// Solidity: event CheckPassed(address indexed player)
func (_C4B *C4BFilterer) FilterCheckPassed(opts *bind.FilterOpts, player []common.Address) (*C4BCheckPassedIterator, error) {

	var playerRule []interface{}
	for _, playerItem := range player {
		playerRule = append(playerRule, playerItem)
	}

	logs, sub, err := _C4B.contract.FilterLogs(opts, "CheckPassed", playerRule)
	if err != nil {
		return nil, err
	}
	return &C4BCheckPassedIterator{contract: _C4B.contract, event: "CheckPassed", logs: logs, sub: sub}, nil
}

// WatchCheckPassed is a free log subscription operation binding the contract event 0x320f420edbd58b0816e3c933b93878e7c130093cdc9237d2e3a855b724f69c94.
//
// Solidity: event CheckPassed(address indexed player)
func (_C4B *C4BFilterer) WatchCheckPassed(opts *bind.WatchOpts, sink chan<- *C4BCheckPassed, player []common.Address) (event.Subscription, error) {

	var playerRule []interface{}
	for _, playerItem := range player {
		playerRule = append(playerRule, playerItem)
	}

	logs, sub, err := _C4B.contract.WatchLogs(opts, "CheckPassed", playerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(C4BCheckPassed)
				if err := _C4B.contract.UnpackLog(event, "CheckPassed", log); err != nil {
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

// ParseCheckPassed is a log parse operation binding the contract event 0x320f420edbd58b0816e3c933b93878e7c130093cdc9237d2e3a855b724f69c94.
//
// Solidity: event CheckPassed(address indexed player)
func (_C4B *C4BFilterer) ParseCheckPassed(log types.Log) (*C4BCheckPassed, error) {
	event := new(C4BCheckPassed)
	if err := _C4B.contract.UnpackLog(event, "CheckPassed", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ChallengeManagerABI is the input ABI used to generate the binding from.
const ChallengeManagerABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"player\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"challenge\",\"type\":\"address\"}],\"name\":\"ChallengeDeployed\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"player\",\"type\":\"address\"}],\"name\":\"deploy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ChallengeManagerFuncSigs maps the 4-byte function signature to its string representation.
var ChallengeManagerFuncSigs = map[string]string{
	"4c96a389": "deploy(address)",
}

// ChallengeManagerBin is the compiled bytecode used for deploying new contracts.
var ChallengeManagerBin = "0x608060405234801561001057600080fd5b506103cc806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c80634c96a38914610030575b600080fd5b6100566004803603602081101561004657600080fd5b50356001600160a01b0316610058565b005b604080516000194301406020808301919091528251808303820181529183019283905281519082012090916000908490839061009390610119565b6001600160a01b0390921682526001600160c01b0319166020820152604080519182900301906000f0801580156100ce573d6000803e3d6000fd5b50604080516001600160a01b0380841682529151929350908616917fbe099bcc1a1457aefe830cc3873c27e2ace425de5409368e242714a24722af159181900360200190a250505050565b610270806101278339019056fe608060405234801561001057600080fd5b506040516102703803806102708339818101604052604081101561003357600080fd5b508051602090910151600080546001600160a01b0319166001600160a01b0390931692909217600160a01b600160e01b031916600160a01b60c09290921c919091021760ff60e01b191681556101e190819061008f90396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80630b93381b1461004657806348db5f8914610062578063c577930a14610086575b600080fd5b61004e6100b3565b604080519115158252519081900360200190f35b61006a6100c3565b604080516001600160a01b039092168252519081900360200190f35b61004e6004803603604081101561009c57600080fd5b506001600160c01b031981351690602001356100d2565b600054600160e01b900460ff1681565b6000546001600160a01b031681565b600060001943014082620f42408206141561016057604080516001600160c01b03198087166020808401919091528351808403820181528385018552805190820120600054600160a01b900460c01b909216606080850191909152845180850390910181526080909301909352815191909201201415610160576000805460ff60e01b1916600160e01b1790555b600080546040516001600160a01b03909116917f320f420edbd58b0816e3c933b93878e7c130093cdc9237d2e3a855b724f69c9491a25050600054600160e01b900460ff169291505056fea2646970667358221220aeab6d521c7fa54a91ba511baec1e6bf53f8d23d43aaf2b13d06602f4302efa664736f6c63430006060033a2646970667358221220f6242f236924ceee69f0deef808d3980075a97b17f67394c16ad03de97dae0bb64736f6c63430006060033"

// DeployChallengeManager deploys a new Ethereum contract, binding an instance of ChallengeManager to it.
func DeployChallengeManager(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ChallengeManager, error) {
	parsed, err := abi.JSON(strings.NewReader(ChallengeManagerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ChallengeManagerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ChallengeManager{ChallengeManagerCaller: ChallengeManagerCaller{contract: contract}, ChallengeManagerTransactor: ChallengeManagerTransactor{contract: contract}, ChallengeManagerFilterer: ChallengeManagerFilterer{contract: contract}}, nil
}

// ChallengeManager is an auto generated Go binding around an Ethereum contract.
type ChallengeManager struct {
	ChallengeManagerCaller     // Read-only binding to the contract
	ChallengeManagerTransactor // Write-only binding to the contract
	ChallengeManagerFilterer   // Log filterer for contract events
}

// ChallengeManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ChallengeManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChallengeManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ChallengeManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChallengeManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ChallengeManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChallengeManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ChallengeManagerSession struct {
	Contract     *ChallengeManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ChallengeManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ChallengeManagerCallerSession struct {
	Contract *ChallengeManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ChallengeManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ChallengeManagerTransactorSession struct {
	Contract     *ChallengeManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ChallengeManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ChallengeManagerRaw struct {
	Contract *ChallengeManager // Generic contract binding to access the raw methods on
}

// ChallengeManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ChallengeManagerCallerRaw struct {
	Contract *ChallengeManagerCaller // Generic read-only contract binding to access the raw methods on
}

// ChallengeManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ChallengeManagerTransactorRaw struct {
	Contract *ChallengeManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewChallengeManager creates a new instance of ChallengeManager, bound to a specific deployed contract.
func NewChallengeManager(address common.Address, backend bind.ContractBackend) (*ChallengeManager, error) {
	contract, err := bindChallengeManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ChallengeManager{ChallengeManagerCaller: ChallengeManagerCaller{contract: contract}, ChallengeManagerTransactor: ChallengeManagerTransactor{contract: contract}, ChallengeManagerFilterer: ChallengeManagerFilterer{contract: contract}}, nil
}

// NewChallengeManagerCaller creates a new read-only instance of ChallengeManager, bound to a specific deployed contract.
func NewChallengeManagerCaller(address common.Address, caller bind.ContractCaller) (*ChallengeManagerCaller, error) {
	contract, err := bindChallengeManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ChallengeManagerCaller{contract: contract}, nil
}

// NewChallengeManagerTransactor creates a new write-only instance of ChallengeManager, bound to a specific deployed contract.
func NewChallengeManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ChallengeManagerTransactor, error) {
	contract, err := bindChallengeManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ChallengeManagerTransactor{contract: contract}, nil
}

// NewChallengeManagerFilterer creates a new log filterer instance of ChallengeManager, bound to a specific deployed contract.
func NewChallengeManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ChallengeManagerFilterer, error) {
	contract, err := bindChallengeManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ChallengeManagerFilterer{contract: contract}, nil
}

// bindChallengeManager binds a generic wrapper to an already deployed contract.
func bindChallengeManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ChallengeManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChallengeManager *ChallengeManagerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ChallengeManager.Contract.ChallengeManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChallengeManager *ChallengeManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChallengeManager.Contract.ChallengeManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChallengeManager *ChallengeManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChallengeManager.Contract.ChallengeManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChallengeManager *ChallengeManagerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ChallengeManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChallengeManager *ChallengeManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChallengeManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChallengeManager *ChallengeManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChallengeManager.Contract.contract.Transact(opts, method, params...)
}

// Deploy is a paid mutator transaction binding the contract method 0x4c96a389.
//
// Solidity: function deploy(address player) returns()
func (_ChallengeManager *ChallengeManagerTransactor) Deploy(opts *bind.TransactOpts, player common.Address) (*types.Transaction, error) {
	return _ChallengeManager.contract.Transact(opts, "deploy", player)
}

// Deploy is a paid mutator transaction binding the contract method 0x4c96a389.
//
// Solidity: function deploy(address player) returns()
func (_ChallengeManager *ChallengeManagerSession) Deploy(player common.Address) (*types.Transaction, error) {
	return _ChallengeManager.Contract.Deploy(&_ChallengeManager.TransactOpts, player)
}

// Deploy is a paid mutator transaction binding the contract method 0x4c96a389.
//
// Solidity: function deploy(address player) returns()
func (_ChallengeManager *ChallengeManagerTransactorSession) Deploy(player common.Address) (*types.Transaction, error) {
	return _ChallengeManager.Contract.Deploy(&_ChallengeManager.TransactOpts, player)
}

// ChallengeManagerChallengeDeployedIterator is returned from FilterChallengeDeployed and is used to iterate over the raw logs and unpacked data for ChallengeDeployed events raised by the ChallengeManager contract.
type ChallengeManagerChallengeDeployedIterator struct {
	Event *ChallengeManagerChallengeDeployed // Event containing the contract specifics and raw log

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
func (it *ChallengeManagerChallengeDeployedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChallengeManagerChallengeDeployed)
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
		it.Event = new(ChallengeManagerChallengeDeployed)
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
func (it *ChallengeManagerChallengeDeployedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChallengeManagerChallengeDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChallengeManagerChallengeDeployed represents a ChallengeDeployed event raised by the ChallengeManager contract.
type ChallengeManagerChallengeDeployed struct {
	Player    common.Address
	Challenge common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterChallengeDeployed is a free log retrieval operation binding the contract event 0xbe099bcc1a1457aefe830cc3873c27e2ace425de5409368e242714a24722af15.
//
// Solidity: event ChallengeDeployed(address indexed player, address challenge)
func (_ChallengeManager *ChallengeManagerFilterer) FilterChallengeDeployed(opts *bind.FilterOpts, player []common.Address) (*ChallengeManagerChallengeDeployedIterator, error) {

	var playerRule []interface{}
	for _, playerItem := range player {
		playerRule = append(playerRule, playerItem)
	}

	logs, sub, err := _ChallengeManager.contract.FilterLogs(opts, "ChallengeDeployed", playerRule)
	if err != nil {
		return nil, err
	}
	return &ChallengeManagerChallengeDeployedIterator{contract: _ChallengeManager.contract, event: "ChallengeDeployed", logs: logs, sub: sub}, nil
}

// WatchChallengeDeployed is a free log subscription operation binding the contract event 0xbe099bcc1a1457aefe830cc3873c27e2ace425de5409368e242714a24722af15.
//
// Solidity: event ChallengeDeployed(address indexed player, address challenge)
func (_ChallengeManager *ChallengeManagerFilterer) WatchChallengeDeployed(opts *bind.WatchOpts, sink chan<- *ChallengeManagerChallengeDeployed, player []common.Address) (event.Subscription, error) {

	var playerRule []interface{}
	for _, playerItem := range player {
		playerRule = append(playerRule, playerItem)
	}

	logs, sub, err := _ChallengeManager.contract.WatchLogs(opts, "ChallengeDeployed", playerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChallengeManagerChallengeDeployed)
				if err := _ChallengeManager.contract.UnpackLog(event, "ChallengeDeployed", log); err != nil {
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

// ParseChallengeDeployed is a log parse operation binding the contract event 0xbe099bcc1a1457aefe830cc3873c27e2ace425de5409368e242714a24722af15.
//
// Solidity: event ChallengeDeployed(address indexed player, address challenge)
func (_ChallengeManager *ChallengeManagerFilterer) ParseChallengeDeployed(log types.Log) (*ChallengeManagerChallengeDeployed, error) {
	event := new(ChallengeManagerChallengeDeployed)
	if err := _ChallengeManager.contract.UnpackLog(event, "ChallengeDeployed", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SolveC4BABI is the input ABI used to generate the binding from.
const SolveC4BABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"bytes8\",\"name\":\"password\",\"type\":\"bytes8\"}],\"name\":\"solve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// SolveC4BFuncSigs maps the 4-byte function signature to its string representation.
var SolveC4BFuncSigs = map[string]string{
	"231644ee": "solve(bytes8)",
}

// SolveC4BBin is the compiled bytecode used for deploying new contracts.
var SolveC4BBin = "0x608060405234801561001057600080fd5b506040516101a63803806101a68339818101604052602081101561003357600080fd5b5051600080546001600160a01b039092166001600160a01b0319909216919091179055610141806100656000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c8063231644ee14610030575b600080fd5b6100576004803603602081101561004657600080fd5b50356001600160c01b03191661006b565b604080519115158252519081900360200190f35b60008054604080516362bbc98560e11b81526001600160c01b031985166004820152620f424043600019014090810660248301529151919284926001600160a01b039091169163c577930a91604480830192602092919082900301818787803b1580156100d757600080fd5b505af11580156100eb573d6000803e3d6000fd5b505050506040513d602081101561010157600080fd5b505194935050505056fea26469706673582212205f405d65a2456913930c045c143591790eaa5040e019e9b07dba2b400de9b05b64736f6c63430006060033"

// DeploySolveC4B deploys a new Ethereum contract, binding an instance of SolveC4B to it.
func DeploySolveC4B(auth *bind.TransactOpts, backend bind.ContractBackend, addr common.Address) (common.Address, *types.Transaction, *SolveC4B, error) {
	parsed, err := abi.JSON(strings.NewReader(SolveC4BABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SolveC4BBin), backend, addr)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SolveC4B{SolveC4BCaller: SolveC4BCaller{contract: contract}, SolveC4BTransactor: SolveC4BTransactor{contract: contract}, SolveC4BFilterer: SolveC4BFilterer{contract: contract}}, nil
}

// SolveC4B is an auto generated Go binding around an Ethereum contract.
type SolveC4B struct {
	SolveC4BCaller     // Read-only binding to the contract
	SolveC4BTransactor // Write-only binding to the contract
	SolveC4BFilterer   // Log filterer for contract events
}

// SolveC4BCaller is an auto generated read-only Go binding around an Ethereum contract.
type SolveC4BCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SolveC4BTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SolveC4BTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SolveC4BFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SolveC4BFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SolveC4BSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SolveC4BSession struct {
	Contract     *SolveC4B         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SolveC4BCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SolveC4BCallerSession struct {
	Contract *SolveC4BCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SolveC4BTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SolveC4BTransactorSession struct {
	Contract     *SolveC4BTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SolveC4BRaw is an auto generated low-level Go binding around an Ethereum contract.
type SolveC4BRaw struct {
	Contract *SolveC4B // Generic contract binding to access the raw methods on
}

// SolveC4BCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SolveC4BCallerRaw struct {
	Contract *SolveC4BCaller // Generic read-only contract binding to access the raw methods on
}

// SolveC4BTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SolveC4BTransactorRaw struct {
	Contract *SolveC4BTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSolveC4B creates a new instance of SolveC4B, bound to a specific deployed contract.
func NewSolveC4B(address common.Address, backend bind.ContractBackend) (*SolveC4B, error) {
	contract, err := bindSolveC4B(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SolveC4B{SolveC4BCaller: SolveC4BCaller{contract: contract}, SolveC4BTransactor: SolveC4BTransactor{contract: contract}, SolveC4BFilterer: SolveC4BFilterer{contract: contract}}, nil
}

// NewSolveC4BCaller creates a new read-only instance of SolveC4B, bound to a specific deployed contract.
func NewSolveC4BCaller(address common.Address, caller bind.ContractCaller) (*SolveC4BCaller, error) {
	contract, err := bindSolveC4B(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SolveC4BCaller{contract: contract}, nil
}

// NewSolveC4BTransactor creates a new write-only instance of SolveC4B, bound to a specific deployed contract.
func NewSolveC4BTransactor(address common.Address, transactor bind.ContractTransactor) (*SolveC4BTransactor, error) {
	contract, err := bindSolveC4B(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SolveC4BTransactor{contract: contract}, nil
}

// NewSolveC4BFilterer creates a new log filterer instance of SolveC4B, bound to a specific deployed contract.
func NewSolveC4BFilterer(address common.Address, filterer bind.ContractFilterer) (*SolveC4BFilterer, error) {
	contract, err := bindSolveC4B(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SolveC4BFilterer{contract: contract}, nil
}

// bindSolveC4B binds a generic wrapper to an already deployed contract.
func bindSolveC4B(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SolveC4BABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SolveC4B *SolveC4BRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SolveC4B.Contract.SolveC4BCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SolveC4B *SolveC4BRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SolveC4B.Contract.SolveC4BTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SolveC4B *SolveC4BRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SolveC4B.Contract.SolveC4BTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SolveC4B *SolveC4BCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SolveC4B.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SolveC4B *SolveC4BTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SolveC4B.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SolveC4B *SolveC4BTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SolveC4B.Contract.contract.Transact(opts, method, params...)
}

// Solve is a paid mutator transaction binding the contract method 0x231644ee.
//
// Solidity: function solve(bytes8 password) returns(bool)
func (_SolveC4B *SolveC4BTransactor) Solve(opts *bind.TransactOpts, password [8]byte) (*types.Transaction, error) {
	return _SolveC4B.contract.Transact(opts, "solve", password)
}

// Solve is a paid mutator transaction binding the contract method 0x231644ee.
//
// Solidity: function solve(bytes8 password) returns(bool)
func (_SolveC4B *SolveC4BSession) Solve(password [8]byte) (*types.Transaction, error) {
	return _SolveC4B.Contract.Solve(&_SolveC4B.TransactOpts, password)
}

// Solve is a paid mutator transaction binding the contract method 0x231644ee.
//
// Solidity: function solve(bytes8 password) returns(bool)
func (_SolveC4B *SolveC4BTransactorSession) Solve(password [8]byte) (*types.Transaction, error) {
	return _SolveC4B.Contract.Solve(&_SolveC4B.TransactOpts, password)
}
