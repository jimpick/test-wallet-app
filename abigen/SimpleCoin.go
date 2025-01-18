// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abigen

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
	_ = abi.ConvertType
)

// SimpleCoinMetaData contains all meta data concerning the SimpleCoin contract.
var SimpleCoinMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getBalance\",\"inputs\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBalanceInEth\",\"inputs\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"sendCoin\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"sufficient\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"_from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
}

// SimpleCoinABI is the input ABI used to generate the binding from.
// Deprecated: Use SimpleCoinMetaData.ABI instead.
var SimpleCoinABI = SimpleCoinMetaData.ABI

// SimpleCoin is an auto generated Go binding around an Ethereum contract.
type SimpleCoin struct {
	SimpleCoinCaller     // Read-only binding to the contract
	SimpleCoinTransactor // Write-only binding to the contract
	SimpleCoinFilterer   // Log filterer for contract events
}

// SimpleCoinCaller is an auto generated read-only Go binding around an Ethereum contract.
type SimpleCoinCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleCoinTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SimpleCoinTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleCoinFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SimpleCoinFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleCoinSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SimpleCoinSession struct {
	Contract     *SimpleCoin       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SimpleCoinCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SimpleCoinCallerSession struct {
	Contract *SimpleCoinCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// SimpleCoinTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SimpleCoinTransactorSession struct {
	Contract     *SimpleCoinTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// SimpleCoinRaw is an auto generated low-level Go binding around an Ethereum contract.
type SimpleCoinRaw struct {
	Contract *SimpleCoin // Generic contract binding to access the raw methods on
}

// SimpleCoinCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SimpleCoinCallerRaw struct {
	Contract *SimpleCoinCaller // Generic read-only contract binding to access the raw methods on
}

// SimpleCoinTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SimpleCoinTransactorRaw struct {
	Contract *SimpleCoinTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSimpleCoin creates a new instance of SimpleCoin, bound to a specific deployed contract.
func NewSimpleCoin(address common.Address, backend bind.ContractBackend) (*SimpleCoin, error) {
	contract, err := bindSimpleCoin(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SimpleCoin{SimpleCoinCaller: SimpleCoinCaller{contract: contract}, SimpleCoinTransactor: SimpleCoinTransactor{contract: contract}, SimpleCoinFilterer: SimpleCoinFilterer{contract: contract}}, nil
}

// NewSimpleCoinCaller creates a new read-only instance of SimpleCoin, bound to a specific deployed contract.
func NewSimpleCoinCaller(address common.Address, caller bind.ContractCaller) (*SimpleCoinCaller, error) {
	contract, err := bindSimpleCoin(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SimpleCoinCaller{contract: contract}, nil
}

// NewSimpleCoinTransactor creates a new write-only instance of SimpleCoin, bound to a specific deployed contract.
func NewSimpleCoinTransactor(address common.Address, transactor bind.ContractTransactor) (*SimpleCoinTransactor, error) {
	contract, err := bindSimpleCoin(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SimpleCoinTransactor{contract: contract}, nil
}

// NewSimpleCoinFilterer creates a new log filterer instance of SimpleCoin, bound to a specific deployed contract.
func NewSimpleCoinFilterer(address common.Address, filterer bind.ContractFilterer) (*SimpleCoinFilterer, error) {
	contract, err := bindSimpleCoin(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SimpleCoinFilterer{contract: contract}, nil
}

// bindSimpleCoin binds a generic wrapper to an already deployed contract.
func bindSimpleCoin(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SimpleCoinMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SimpleCoin *SimpleCoinRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SimpleCoin.Contract.SimpleCoinCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SimpleCoin *SimpleCoinRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleCoin.Contract.SimpleCoinTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SimpleCoin *SimpleCoinRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SimpleCoin.Contract.SimpleCoinTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SimpleCoin *SimpleCoinCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SimpleCoin.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SimpleCoin *SimpleCoinTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleCoin.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SimpleCoin *SimpleCoinTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SimpleCoin.Contract.contract.Transact(opts, method, params...)
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address addr) view returns(uint256)
func (_SimpleCoin *SimpleCoinCaller) GetBalance(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SimpleCoin.contract.Call(opts, &out, "getBalance", addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address addr) view returns(uint256)
func (_SimpleCoin *SimpleCoinSession) GetBalance(addr common.Address) (*big.Int, error) {
	return _SimpleCoin.Contract.GetBalance(&_SimpleCoin.CallOpts, addr)
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address addr) view returns(uint256)
func (_SimpleCoin *SimpleCoinCallerSession) GetBalance(addr common.Address) (*big.Int, error) {
	return _SimpleCoin.Contract.GetBalance(&_SimpleCoin.CallOpts, addr)
}

// GetBalanceInEth is a free data retrieval call binding the contract method 0x7bd703e8.
//
// Solidity: function getBalanceInEth(address addr) view returns(uint256)
func (_SimpleCoin *SimpleCoinCaller) GetBalanceInEth(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SimpleCoin.contract.Call(opts, &out, "getBalanceInEth", addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalanceInEth is a free data retrieval call binding the contract method 0x7bd703e8.
//
// Solidity: function getBalanceInEth(address addr) view returns(uint256)
func (_SimpleCoin *SimpleCoinSession) GetBalanceInEth(addr common.Address) (*big.Int, error) {
	return _SimpleCoin.Contract.GetBalanceInEth(&_SimpleCoin.CallOpts, addr)
}

// GetBalanceInEth is a free data retrieval call binding the contract method 0x7bd703e8.
//
// Solidity: function getBalanceInEth(address addr) view returns(uint256)
func (_SimpleCoin *SimpleCoinCallerSession) GetBalanceInEth(addr common.Address) (*big.Int, error) {
	return _SimpleCoin.Contract.GetBalanceInEth(&_SimpleCoin.CallOpts, addr)
}

// SendCoin is a paid mutator transaction binding the contract method 0x90b98a11.
//
// Solidity: function sendCoin(address receiver, uint256 amount) returns(bool sufficient)
func (_SimpleCoin *SimpleCoinTransactor) SendCoin(opts *bind.TransactOpts, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SimpleCoin.contract.Transact(opts, "sendCoin", receiver, amount)
}

// SendCoin is a paid mutator transaction binding the contract method 0x90b98a11.
//
// Solidity: function sendCoin(address receiver, uint256 amount) returns(bool sufficient)
func (_SimpleCoin *SimpleCoinSession) SendCoin(receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SimpleCoin.Contract.SendCoin(&_SimpleCoin.TransactOpts, receiver, amount)
}

// SendCoin is a paid mutator transaction binding the contract method 0x90b98a11.
//
// Solidity: function sendCoin(address receiver, uint256 amount) returns(bool sufficient)
func (_SimpleCoin *SimpleCoinTransactorSession) SendCoin(receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SimpleCoin.Contract.SendCoin(&_SimpleCoin.TransactOpts, receiver, amount)
}

// SimpleCoinTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the SimpleCoin contract.
type SimpleCoinTransferIterator struct {
	Event *SimpleCoinTransfer // Event containing the contract specifics and raw log

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
func (it *SimpleCoinTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleCoinTransfer)
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
		it.Event = new(SimpleCoinTransfer)
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
func (it *SimpleCoinTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleCoinTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleCoinTransfer represents a Transfer event raised by the SimpleCoin contract.
type SimpleCoinTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_SimpleCoin *SimpleCoinFilterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*SimpleCoinTransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _SimpleCoin.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &SimpleCoinTransferIterator{contract: _SimpleCoin.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_SimpleCoin *SimpleCoinFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *SimpleCoinTransfer, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _SimpleCoin.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleCoinTransfer)
				if err := _SimpleCoin.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_SimpleCoin *SimpleCoinFilterer) ParseTransfer(log types.Log) (*SimpleCoinTransfer, error) {
	event := new(SimpleCoinTransfer)
	if err := _SimpleCoin.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
