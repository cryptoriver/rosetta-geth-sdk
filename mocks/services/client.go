// Code generated by mockery v1.0.0. DO NOT EDIT.

package services

import (
	big "math/big"

	client "github.com/coinbase/rosetta-geth-sdk/client"
	common "github.com/ethereum/go-ethereum/common"

	configuration "github.com/coinbase/rosetta-geth-sdk/configuration"

	context "context"

	coretypes "github.com/ethereum/go-ethereum/core/types"

	json "encoding/json"

	mock "github.com/stretchr/testify/mock"

	rpc "github.com/ethereum/go-ethereum/rpc"

	types "github.com/coinbase/rosetta-sdk-go/types"
)

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// Balance provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *Client) Balance(_a0 context.Context, _a1 *types.AccountIdentifier, _a2 *types.PartialBlockIdentifier, _a3 []*types.Currency) (*types.AccountBalanceResponse, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	var r0 *types.AccountBalanceResponse
	if rf, ok := ret.Get(0).(func(context.Context, *types.AccountIdentifier, *types.PartialBlockIdentifier, []*types.Currency) *types.AccountBalanceResponse); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.AccountBalanceResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.AccountIdentifier, *types.PartialBlockIdentifier, []*types.Currency) error); ok {
		r1 = rf(_a0, _a1, _a2, _a3)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BatchCallContext provides a mock function with given fields: ctx, b
func (_m *Client) BatchCallContext(ctx context.Context, b []rpc.BatchElem) error {
	ret := _m.Called(ctx, b)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []rpc.BatchElem) error); ok {
		r0 = rf(ctx, b)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// BlockAuthor provides a mock function with given fields: ctx, blockIndex
func (_m *Client) BlockAuthor(ctx context.Context, blockIndex int64) (string, error) {
	ret := _m.Called(ctx, blockIndex)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, int64) string); ok {
		r0 = rf(ctx, blockIndex)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, blockIndex)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BlockRewardTransaction provides a mock function with given fields: blockIdentifier, miner, uncles
func (_m *Client) BlockRewardTransaction(blockIdentifier *types.BlockIdentifier, miner string, uncles []*coretypes.Header) *types.Transaction {
	ret := _m.Called(blockIdentifier, miner, uncles)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(*types.BlockIdentifier, string, []*coretypes.Header) *types.Transaction); ok {
		r0 = rf(blockIdentifier, miner, uncles)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	return r0
}

// CallContext provides a mock function with given fields: ctx, result, method, args
func (_m *Client) CallContext(ctx context.Context, result interface{}, method string, args ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, ctx, result, method)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, string, ...interface{}) error); ok {
		r0 = rf(ctx, result, method, args...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetBlockReceipts provides a mock function with given fields: ctx, blockHash, txs, baseFee
func (_m *Client) GetBlockReceipts(ctx context.Context, blockHash common.Hash, txs []client.RPCTransaction, baseFee *big.Int) ([]*client.RosettaTxReceipt, error) {
	ret := _m.Called(ctx, blockHash, txs, baseFee)

	var r0 []*client.RosettaTxReceipt
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash, []client.RPCTransaction, *big.Int) []*client.RosettaTxReceipt); ok {
		r0 = rf(ctx, blockHash, txs, baseFee)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*client.RosettaTxReceipt)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, common.Hash, []client.RPCTransaction, *big.Int) error); ok {
		r1 = rf(ctx, blockHash, txs, baseFee)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetContractCallGasLimit provides a mock function with given fields: ctx, toAddress, fromAddress, data
func (_m *Client) GetContractCallGasLimit(ctx context.Context, toAddress string, fromAddress string, data []byte) (uint64, error) {
	ret := _m.Called(ctx, toAddress, fromAddress, data)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(context.Context, string, string, []byte) uint64); ok {
		r0 = rf(ctx, toAddress, fromAddress, data)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, []byte) error); ok {
		r1 = rf(ctx, toAddress, fromAddress, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetContractCurrency provides a mock function with given fields: addr, erc20
func (_m *Client) GetContractCurrency(addr common.Address, erc20 bool) (*client.ContractCurrency, error) {
	ret := _m.Called(addr, erc20)

	var r0 *client.ContractCurrency
	if rf, ok := ret.Get(0).(func(common.Address, bool) *client.ContractCurrency); ok {
		r0 = rf(addr, erc20)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.ContractCurrency)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(common.Address, bool) error); ok {
		r1 = rf(addr, erc20)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetErc20TransferGasLimit provides a mock function with given fields: ctx, toAddress, fromAddress, value, currency
func (_m *Client) GetErc20TransferGasLimit(ctx context.Context, toAddress string, fromAddress string, value *big.Int, currency *types.Currency) (uint64, error) {
	ret := _m.Called(ctx, toAddress, fromAddress, value, currency)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(context.Context, string, string, *big.Int, *types.Currency) uint64); ok {
		r0 = rf(ctx, toAddress, fromAddress, value, currency)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, *big.Int, *types.Currency) error); ok {
		r1 = rf(ctx, toAddress, fromAddress, value, currency)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetGasPrice provides a mock function with given fields: _a0, _a1
func (_m *Client) GetGasPrice(_a0 context.Context, _a1 client.Options) (*big.Int, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(context.Context, client.Options) *big.Int); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, client.Options) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLoadedTransaction provides a mock function with given fields: ctx, request
func (_m *Client) GetLoadedTransaction(ctx context.Context, request *types.BlockTransactionRequest) (*client.LoadedTransaction, error) {
	ret := _m.Called(ctx, request)

	var r0 *client.LoadedTransaction
	if rf, ok := ret.Get(0).(func(context.Context, *types.BlockTransactionRequest) *client.LoadedTransaction); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.LoadedTransaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.BlockTransactionRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNativeTransferGasLimit provides a mock function with given fields: ctx, toAddress, fromAddress, value
func (_m *Client) GetNativeTransferGasLimit(ctx context.Context, toAddress string, fromAddress string, value *big.Int) (uint64, error) {
	ret := _m.Called(ctx, toAddress, fromAddress, value)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(context.Context, string, string, *big.Int) uint64); ok {
		r0 = rf(ctx, toAddress, fromAddress, value)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, *big.Int) error); ok {
		r1 = rf(ctx, toAddress, fromAddress, value)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNonce provides a mock function with given fields: _a0, _a1
func (_m *Client) GetNonce(_a0 context.Context, _a1 client.Options) (uint64, error) {
	ret := _m.Called(_a0, _a1)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(context.Context, client.Options) uint64); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, client.Options) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRosettaConfig provides a mock function with given fields:
func (_m *Client) GetRosettaConfig() configuration.RosettaConfig {
	ret := _m.Called()

	var r0 configuration.RosettaConfig
	if rf, ok := ret.Get(0).(func() configuration.RosettaConfig); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(configuration.RosettaConfig)
	}

	return r0
}

// GetTransactionReceipt provides a mock function with given fields: ctx, tx
func (_m *Client) GetTransactionReceipt(ctx context.Context, tx *client.LoadedTransaction) (*client.RosettaTxReceipt, error) {
	ret := _m.Called(ctx, tx)

	var r0 *client.RosettaTxReceipt
	if rf, ok := ret.Get(0).(func(context.Context, *client.LoadedTransaction) *client.RosettaTxReceipt); ok {
		r0 = rf(ctx, tx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.RosettaTxReceipt)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *client.LoadedTransaction) error); ok {
		r1 = rf(ctx, tx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUncles provides a mock function with given fields: ctx, head, body
func (_m *Client) GetUncles(ctx context.Context, head *coretypes.Header, body *client.RPCBlock) ([]*coretypes.Header, error) {
	ret := _m.Called(ctx, head, body)

	var r0 []*coretypes.Header
	if rf, ok := ret.Get(0).(func(context.Context, *coretypes.Header, *client.RPCBlock) []*coretypes.Header); ok {
		r0 = rf(ctx, head, body)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*coretypes.Header)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *coretypes.Header, *client.RPCBlock) error); ok {
		r1 = rf(ctx, head, body)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PopulateCrossChainTransactions provides a mock function with given fields: _a0, _a1
func (_m *Client) PopulateCrossChainTransactions(_a0 *coretypes.Block, _a1 []*client.LoadedTransaction) ([]*types.Transaction, error) {
	ret := _m.Called(_a0, _a1)

	var r0 []*types.Transaction
	if rf, ok := ret.Get(0).(func(*coretypes.Block, []*client.LoadedTransaction) []*types.Transaction); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*types.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*coretypes.Block, []*client.LoadedTransaction) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Status provides a mock function with given fields: _a0
func (_m *Client) Status(_a0 context.Context) (*types.BlockIdentifier, int64, *types.SyncStatus, []*types.Peer, error) {
	ret := _m.Called(_a0)

	var r0 *types.BlockIdentifier
	if rf, ok := ret.Get(0).(func(context.Context) *types.BlockIdentifier); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.BlockIdentifier)
		}
	}

	var r1 int64
	if rf, ok := ret.Get(1).(func(context.Context) int64); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Get(1).(int64)
	}

	var r2 *types.SyncStatus
	if rf, ok := ret.Get(2).(func(context.Context) *types.SyncStatus); ok {
		r2 = rf(_a0)
	} else {
		if ret.Get(2) != nil {
			r2 = ret.Get(2).(*types.SyncStatus)
		}
	}

	var r3 []*types.Peer
	if rf, ok := ret.Get(3).(func(context.Context) []*types.Peer); ok {
		r3 = rf(_a0)
	} else {
		if ret.Get(3) != nil {
			r3 = ret.Get(3).([]*types.Peer)
		}
	}

	var r4 error
	if rf, ok := ret.Get(4).(func(context.Context) error); ok {
		r4 = rf(_a0)
	} else {
		r4 = ret.Error(4)
	}

	return r0, r1, r2, r3, r4
}

// Submit provides a mock function with given fields: _a0, _a1
func (_m *Client) Submit(_a0 context.Context, _a1 *coretypes.Transaction) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *coretypes.Transaction) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TraceBlockByHash provides a mock function with given fields: _a0, _a1, _a2
func (_m *Client) TraceBlockByHash(_a0 context.Context, _a1 common.Hash, _a2 []client.RPCTransaction) (map[string][]*client.FlatCall, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 map[string][]*client.FlatCall
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash, []client.RPCTransaction) map[string][]*client.FlatCall); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string][]*client.FlatCall)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, common.Hash, []client.RPCTransaction) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TraceReplayBlockTransactions provides a mock function with given fields: ctx, hsh
func (_m *Client) TraceReplayBlockTransactions(ctx context.Context, hsh string) (map[string][]*client.FlatCall, error) {
	ret := _m.Called(ctx, hsh)

	var r0 map[string][]*client.FlatCall
	if rf, ok := ret.Get(0).(func(context.Context, string) map[string][]*client.FlatCall); ok {
		r0 = rf(ctx, hsh)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string][]*client.FlatCall)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, hsh)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TraceReplayTransaction provides a mock function with given fields: ctx, hsh
func (_m *Client) TraceReplayTransaction(ctx context.Context, hsh string) (json.RawMessage, []*client.FlatCall, error) {
	ret := _m.Called(ctx, hsh)

	var r0 json.RawMessage
	if rf, ok := ret.Get(0).(func(context.Context, string) json.RawMessage); ok {
		r0 = rf(ctx, hsh)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(json.RawMessage)
		}
	}

	var r1 []*client.FlatCall
	if rf, ok := ret.Get(1).(func(context.Context, string) []*client.FlatCall); ok {
		r1 = rf(ctx, hsh)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]*client.FlatCall)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string) error); ok {
		r2 = rf(ctx, hsh)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// TraceTransaction provides a mock function with given fields: ctx, hash
func (_m *Client) TraceTransaction(ctx context.Context, hash common.Hash) (json.RawMessage, []*client.FlatCall, error) {
	ret := _m.Called(ctx, hash)

	var r0 json.RawMessage
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash) json.RawMessage); ok {
		r0 = rf(ctx, hash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(json.RawMessage)
		}
	}

	var r1 []*client.FlatCall
	if rf, ok := ret.Get(1).(func(context.Context, common.Hash) []*client.FlatCall); ok {
		r1 = rf(ctx, hash)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]*client.FlatCall)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, common.Hash) error); ok {
		r2 = rf(ctx, hash)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}