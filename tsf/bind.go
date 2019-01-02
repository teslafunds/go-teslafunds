// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

<<<<<<< HEAD:tsf/bind.go
package tsf
=======
package eth
>>>>>>> 7fdd714... gdbix-update v1.5.0:dbix/bind.go

import (
	"math/big"

<<<<<<< HEAD:tsf/bind.go
	"github.com/teslafunds/go-teslafunds/common"
	"github.com/teslafunds/go-teslafunds/core/types"
	"github.com/teslafunds/go-teslafunds/rlp"
	"github.com/teslafunds/go-teslafunds/rpc"
=======
	"github.com/dubaicoin-dbix/go-dubaicoin"
	"github.com/dubaicoin-dbix/go-dubaicoin/common"
	"github.com/dubaicoin-dbix/go-dubaicoin/common/hexutil"
	"github.com/dubaicoin-dbix/go-dubaicoin/core/types"
	"github.com/dubaicoin-dbix/go-dubaicoin/internal/ethapi"
	"github.com/dubaicoin-dbix/go-dubaicoin/rlp"
	"github.com/dubaicoin-dbix/go-dubaicoin/rpc"
	"golang.org/x/net/context"
>>>>>>> 7fdd714... gdbix-update v1.5.0:dbix/bind.go
)

// ContractBackend implements bind.ContractBackend with direct calls to Ethereum
// internals to support operating on contracts within subprotocols like eth and
// swarm.
//
// Internally this backend uses the already exposed API endpoints of the Dubaicoin
// object. These should be rewritten to internal Go method calls when the Go API
// is refactored to support a clean library use.
type ContractBackend struct {
<<<<<<< HEAD:tsf/bind.go
	eapi  *PublicEthereumAPI        // Wrapper around the Teslafunds object to access metadata
	bcapi *PublicBlockChainAPI      // Wrapper around the blockchain to access chain data
	txapi *PublicTransactionPoolAPI // Wrapper around the transaction pool to access transaction data
=======
	eapi  *ethapi.PublicEthereumAPI        // Wrapper around the Dubaicoin object to access metadata
	bcapi *ethapi.PublicBlockChainAPI      // Wrapper around the blockchain to access chain data
	txapi *ethapi.PublicTransactionPoolAPI // Wrapper around the transaction pool to access transaction data
>>>>>>> 7fdd714... gdbix-update v1.5.0:dbix/bind.go
}

// NewContractBackend creates a new native contract backend using an existing
// Etheruem object.
<<<<<<< HEAD:tsf/bind.go
func NewContractBackend(tsf *Teslafunds) *ContractBackend {
	return &ContractBackend{
		eapi:  NewPublicEthereumAPI(tsf),
		bcapi: NewPublicBlockChainAPI(tsf.chainConfig, tsf.blockchain, tsf.miner, tsf.chainDb, tsf.gpo, tsf.eventMux, tsf.accountManager),
		txapi: NewPublicTransactionPoolAPI(tsf),
=======
func NewContractBackend(apiBackend ethapi.Backend) *ContractBackend {
	return &ContractBackend{
		eapi:  ethapi.NewPublicEthereumAPI(apiBackend),
		bcapi: ethapi.NewPublicBlockChainAPI(apiBackend),
		txapi: ethapi.NewPublicTransactionPoolAPI(apiBackend),
>>>>>>> 7fdd714... gdbix-update v1.5.0:dbix/bind.go
	}
}

// CodeAt retrieves any code associated with the contract from the local API.
func (b *ContractBackend) CodeAt(ctx context.Context, contract common.Address, blockNum *big.Int) ([]byte, error) {
	out, err := b.bcapi.GetCode(ctx, contract, toBlockNumber(blockNum))
	return common.FromHex(out), err
}

// CodeAt retrieves any code associated with the contract from the local API.
func (b *ContractBackend) PendingCodeAt(ctx context.Context, contract common.Address) ([]byte, error) {
	out, err := b.bcapi.GetCode(ctx, contract, rpc.PendingBlockNumber)
	return common.FromHex(out), err
}

// ContractCall implements bind.ContractCaller executing a Dubaicoin contract
// call with the specified data as the input. The pending flag requests execution
// against the pending block, not the stable head of the chain.
func (b *ContractBackend) CallContract(ctx context.Context, msg ethereum.CallMsg, blockNum *big.Int) ([]byte, error) {
	out, err := b.bcapi.Call(ctx, toCallArgs(msg), toBlockNumber(blockNum))
	return common.FromHex(out), err
}

// ContractCall implements bind.ContractCaller executing a Dubaicoin contract
// call with the specified data as the input. The pending flag requests execution
// against the pending block, not the stable head of the chain.
func (b *ContractBackend) PendingCallContract(ctx context.Context, msg ethereum.CallMsg) ([]byte, error) {
	out, err := b.bcapi.Call(ctx, toCallArgs(msg), rpc.PendingBlockNumber)
	return common.FromHex(out), err
}

func toCallArgs(msg ethereum.CallMsg) ethapi.CallArgs {
	args := ethapi.CallArgs{
		To:   msg.To,
		From: msg.From,
		Data: msg.Data,
	}
	if msg.Gas != nil {
		args.Gas = hexutil.Big(*msg.Gas)
	}
	if msg.GasPrice != nil {
		args.GasPrice = hexutil.Big(*msg.GasPrice)
	}
	if msg.Value != nil {
		args.Value = hexutil.Big(*msg.Value)
	}
	return args
}

func toBlockNumber(num *big.Int) rpc.BlockNumber {
	if num == nil {
		return rpc.LatestBlockNumber
	}
	return rpc.BlockNumber(num.Int64())
}

// PendingAccountNonce implements bind.ContractTransactor retrieving the current
// pending nonce associated with an account.
func (b *ContractBackend) PendingNonceAt(ctx context.Context, account common.Address) (nonce uint64, err error) {
	out, err := b.txapi.GetTransactionCount(ctx, account, rpc.PendingBlockNumber)
	if out != nil {
		nonce = uint64(*out)
	}
	return nonce, err
}

// SuggestGasPrice implements bind.ContractTransactor retrieving the currently
// suggested gas price to allow a timely execution of a transaction.
func (b *ContractBackend) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	return b.eapi.GasPrice(ctx)
}

// EstimateGasLimit implements bind.ContractTransactor triing to estimate the gas
// needed to execute a specific transaction based on the current pending state of
// the backend blockchain. There is no guarantee that this is the true gas limit
// requirement as other transactions may be added or removed by miners, but it
// should provide a basis for setting a reasonable default.
func (b *ContractBackend) EstimateGas(ctx context.Context, msg ethereum.CallMsg) (*big.Int, error) {
	out, err := b.bcapi.EstimateGas(ctx, toCallArgs(msg))
	return out.ToInt(), err
}

// SendTransaction implements bind.ContractTransactor injects the transaction
// into the pending pool for execution.
func (b *ContractBackend) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	raw, _ := rlp.EncodeToBytes(tx)
	_, err := b.txapi.SendRawTransaction(ctx, raw)
	return err
}
