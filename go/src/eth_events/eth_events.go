// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package eth_events

import (
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// EthEventsABI is the input ABI used to generate the binding from.
const EthEventsABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"emitEvent\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_a\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_b\",\"type\":\"string\"}],\"name\":\"TestEvent\",\"type\":\"event\"}]"

// EthEventsBin is the compiled bytecode used for deploying new contracts.
const EthEventsBin = `60606040523415600e57600080fd5b6101188061001d6000396000f300606060405260043610603e5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416637b0cb83981146043575b600080fd5b3415604d57600080fd5b60536055565b005b7fee93dc634bc91ebdab8284ef96a0cfe42aa5c96aed9a8409f8d013f14fa7889a604051604080825260018183018190527f610000000000000000000000000000000000000000000000000000000000000060608401526080602084018190528301527f620000000000000000000000000000000000000000000000000000000000000060a083015260c0909101905180910390a15600a165627a7a72305820f91bd52241bed6b150c3daf8818beefa57a582847af4abc6e0791d285a67939f0029`

// DeployEthEvents deploys a new Ethereum contract, binding an instance of EthEvents to it.
func DeployEthEvents(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *EthEvents, error) {
	parsed, err := abi.JSON(strings.NewReader(EthEventsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(EthEventsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EthEvents{EthEventsCaller: EthEventsCaller{contract: contract}, EthEventsTransactor: EthEventsTransactor{contract: contract}, EthEventsFilterer: EthEventsFilterer{contract: contract}}, nil
}

// EthEvents is an auto generated Go binding around an Ethereum contract.
type EthEvents struct {
	EthEventsCaller     // Read-only binding to the contract
	EthEventsTransactor // Write-only binding to the contract
	EthEventsFilterer   // Log filterer for contract events
}

// EthEventsCaller is an auto generated read-only Go binding around an Ethereum contract.
type EthEventsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthEventsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EthEventsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthEventsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EthEventsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthEventsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EthEventsSession struct {
	Contract     *EthEvents        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EthEventsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EthEventsCallerSession struct {
	Contract *EthEventsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// EthEventsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EthEventsTransactorSession struct {
	Contract     *EthEventsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// EthEventsRaw is an auto generated low-level Go binding around an Ethereum contract.
type EthEventsRaw struct {
	Contract *EthEvents // Generic contract binding to access the raw methods on
}

// EthEventsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EthEventsCallerRaw struct {
	Contract *EthEventsCaller // Generic read-only contract binding to access the raw methods on
}

// EthEventsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EthEventsTransactorRaw struct {
	Contract *EthEventsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEthEvents creates a new instance of EthEvents, bound to a specific deployed contract.
func NewEthEvents(address common.Address, backend bind.ContractBackend) (*EthEvents, error) {
	contract, err := bindEthEvents(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EthEvents{EthEventsCaller: EthEventsCaller{contract: contract}, EthEventsTransactor: EthEventsTransactor{contract: contract}, EthEventsFilterer: EthEventsFilterer{contract: contract}}, nil
}

// NewEthEventsCaller creates a new read-only instance of EthEvents, bound to a specific deployed contract.
func NewEthEventsCaller(address common.Address, caller bind.ContractCaller) (*EthEventsCaller, error) {
	contract, err := bindEthEvents(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EthEventsCaller{contract: contract}, nil
}

// NewEthEventsTransactor creates a new write-only instance of EthEvents, bound to a specific deployed contract.
func NewEthEventsTransactor(address common.Address, transactor bind.ContractTransactor) (*EthEventsTransactor, error) {
	contract, err := bindEthEvents(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EthEventsTransactor{contract: contract}, nil
}

// NewEthEventsFilterer creates a new log filterer instance of EthEvents, bound to a specific deployed contract.
func NewEthEventsFilterer(address common.Address, filterer bind.ContractFilterer) (*EthEventsFilterer, error) {
	contract, err := bindEthEvents(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EthEventsFilterer{contract: contract}, nil
}

// bindEthEvents binds a generic wrapper to an already deployed contract.
func bindEthEvents(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EthEventsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EthEvents *EthEventsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _EthEvents.Contract.EthEventsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EthEvents *EthEventsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthEvents.Contract.EthEventsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EthEvents *EthEventsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EthEvents.Contract.EthEventsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EthEvents *EthEventsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _EthEvents.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EthEvents *EthEventsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthEvents.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EthEvents *EthEventsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EthEvents.Contract.contract.Transact(opts, method, params...)
}

// EmitEvent is a paid mutator transaction binding the contract method 0x7b0cb839.
//
// Solidity: function emitEvent() returns()
func (_EthEvents *EthEventsTransactor) EmitEvent(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthEvents.contract.Transact(opts, "emitEvent")
}

// EmitEvent is a paid mutator transaction binding the contract method 0x7b0cb839.
//
// Solidity: function emitEvent() returns()
func (_EthEvents *EthEventsSession) EmitEvent() (*types.Transaction, error) {
	return _EthEvents.Contract.EmitEvent(&_EthEvents.TransactOpts)
}

// EmitEvent is a paid mutator transaction binding the contract method 0x7b0cb839.
//
// Solidity: function emitEvent() returns()
func (_EthEvents *EthEventsTransactorSession) EmitEvent() (*types.Transaction, error) {
	return _EthEvents.Contract.EmitEvent(&_EthEvents.TransactOpts)
}

// EthEventsTestEventIterator is returned from FilterTestEvent and is used to iterate over the raw logs and unpacked data for TestEvent events raised by the EthEvents contract.
type EthEventsTestEventIterator struct {
	Event *EthEventsTestEvent // Event containing the contract specifics and raw log

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
func (it *EthEventsTestEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthEventsTestEvent)
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
		it.Event = new(EthEventsTestEvent)
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
func (it *EthEventsTestEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthEventsTestEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthEventsTestEvent represents a TestEvent event raised by the EthEvents contract.
type EthEventsTestEvent struct {
	A   string
	B   string
	Raw types.Log // Blockchain specific contextual infos
}

// FilterTestEvent is a free log retrieval operation binding the contract event 0xee93dc634bc91ebdab8284ef96a0cfe42aa5c96aed9a8409f8d013f14fa7889a.
//
// Solidity: event TestEvent(_a string, _b string)
func (_EthEvents *EthEventsFilterer) FilterTestEvent(opts *bind.FilterOpts) (*EthEventsTestEventIterator, error) {

	logs, sub, err := _EthEvents.contract.FilterLogs(opts, "TestEvent")
	if err != nil {
		return nil, err
	}
	return &EthEventsTestEventIterator{contract: _EthEvents.contract, event: "TestEvent", logs: logs, sub: sub}, nil
}

// WatchTestEvent is a free log subscription operation binding the contract event 0xee93dc634bc91ebdab8284ef96a0cfe42aa5c96aed9a8409f8d013f14fa7889a.
//
// Solidity: event TestEvent(_a string, _b string)
func (_EthEvents *EthEventsFilterer) WatchTestEvent(opts *bind.WatchOpts, sink chan<- *EthEventsTestEvent) (event.Subscription, error) {

	logs, sub, err := _EthEvents.contract.WatchLogs(opts, "TestEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthEventsTestEvent)
				if err := _EthEvents.contract.UnpackLog(event, "TestEvent", log); err != nil {
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
