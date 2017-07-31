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

<<<<<<< HEAD:tsf/api.go
package tsf
=======
package eth
>>>>>>> 7fdd714... gdbix-update v1.5.0:dbix/api.go

import (
	"bytes"
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"math/big"
	"os"
	"runtime"
	"strings"
	"time"

<<<<<<< HEAD:tsf/api.go
	"github.com/teslafunds/ethash"
	"github.com/teslafunds/go-teslafunds/accounts"
	"github.com/teslafunds/go-teslafunds/common"
	"github.com/teslafunds/go-teslafunds/common/compiler"
	"github.com/teslafunds/go-teslafunds/core"
	"github.com/teslafunds/go-teslafunds/core/state"
	"github.com/teslafunds/go-teslafunds/core/types"
	"github.com/teslafunds/go-teslafunds/core/vm"
	"github.com/teslafunds/go-teslafunds/crypto"
	"github.com/teslafunds/go-teslafunds/ethdb"
	"github.com/teslafunds/go-teslafunds/event"
	"github.com/teslafunds/go-teslafunds/logger"
	"github.com/teslafunds/go-teslafunds/logger/glog"
	"github.com/teslafunds/go-teslafunds/miner"
	"github.com/teslafunds/go-teslafunds/p2p"
	"github.com/teslafunds/go-teslafunds/rlp"
	"github.com/teslafunds/go-teslafunds/rpc"
	"github.com/syndtr/goleveldb/leveldb"
	"golang.org/x/net/context"
)

const defaultGas = uint64(90000)

// blockByNumber is a commonly used helper function which retrieves and returns
// the block for the given block number, capable of handling two special blocks:
// rpc.LatestBlockNumber and rpc.PendingBlockNumber. It returns nil when no block
// could be found.
func blockByNumber(m *miner.Miner, bc *core.BlockChain, blockNr rpc.BlockNumber) *types.Block {
	// Pending block is only known by the miner
	if blockNr == rpc.PendingBlockNumber {
		block, _ := m.Pending()
		return block
	}
	// Otherwise resolve and return the block
	if blockNr == rpc.LatestBlockNumber {
		return bc.CurrentBlock()
	}
	return bc.GetBlockByNumber(uint64(blockNr))
}

// stateAndBlockByNumber is a commonly used helper function which retrieves and
// returns the state and containing block for the given block number, capable of
// handling two special states: rpc.LatestBlockNumber and rpc.PendingBlockNumber.
// It returns nil when no block or state could be found.
func stateAndBlockByNumber(m *miner.Miner, bc *core.BlockChain, blockNr rpc.BlockNumber, chainDb ethdb.Database) (*state.StateDB, *types.Block, error) {
	// Pending state is only known by the miner
	if blockNr == rpc.PendingBlockNumber {
		block, state := m.Pending()
		return state, block, nil
	}
	// Otherwise resolve the block number and return its state
	block := blockByNumber(m, bc, blockNr)
	if block == nil {
		return nil, nil, nil
	}
	stateDb, err := state.New(block.Root(), chainDb)
	return stateDb, block, err
}

// PublicEthereumAPI provides an API to access Teslafunds related information.
// It offers only methods that operate on public data that is freely available to anyone.
type PublicEthereumAPI struct {
	e   *Teslafunds
	gpo *GasPriceOracle
}

// NewPublicEthereumAPI creates a new Teslafunds protocol API.
func NewPublicEthereumAPI(e *Teslafunds) *PublicEthereumAPI {
	return &PublicEthereumAPI{
		e:   e,
		gpo: e.gpo,
	}
}

// GasPrice returns a suggestion for a gas price.
func (s *PublicEthereumAPI) GasPrice() *big.Int {
	return s.gpo.SuggestPrice()
}

// GetCompilers returns the collection of available smart contract compilers
func (s *PublicEthereumAPI) GetCompilers() ([]string, error) {
	solc, err := s.e.Solc()
	if err == nil && solc != nil {
		return []string{"Solidity"}, nil
	}

	return []string{}, nil
}

// CompileSolidity compiles the given solidity source
func (s *PublicEthereumAPI) CompileSolidity(source string) (map[string]*compiler.Contract, error) {
	solc, err := s.e.Solc()
	if err != nil {
		return nil, err
	}

	if solc == nil {
		return nil, errors.New("solc (solidity compiler) not found")
	}

	return solc.Compile(source)
}

// Etherbase is the address that mining rewards will be send to
func (s *PublicEthereumAPI) Etherbase() (common.Address, error) {
	return s.e.Etherbase()
}

// Coinbase is the address that mining rewards will be send to (alias for Etherbase)
func (s *PublicEthereumAPI) Coinbase() (common.Address, error) {
	return s.Etherbase()
}

// ProtocolVersion returns the current Teslafunds protocol version this node supports
func (s *PublicEthereumAPI) ProtocolVersion() *rpc.HexNumber {
	return rpc.NewHexNumber(s.e.EthVersion())
}

// Hashrate returns the POW hashrate
func (s *PublicEthereumAPI) Hashrate() *rpc.HexNumber {
	return rpc.NewHexNumber(s.e.Miner().HashRate())
}

// Syncing returns false in case the node is currently not syncing with the network. It can be up to date or has not
// yet received the latest block headers from its pears. In case it is synchronizing:
// - startingBlock: block number this node started to synchronise from
// - currentBlock:  block number this node is currently importing
// - highestBlock:  block number of the highest block header this node has received from peers
// - pulledStates:  number of state entries processed until now
// - knownStates:   number of known state entries that still need to be pulled
func (s *PublicEthereumAPI) Syncing() (interface{}, error) {
	origin, current, height, pulled, known := s.e.Downloader().Progress()

	// Return not syncing if the synchronisation already completed
	if current >= height {
		return false, nil
	}
	// Otherwise gather the block sync stats
	return map[string]interface{}{
		"startingBlock": rpc.NewHexNumber(origin),
		"currentBlock":  rpc.NewHexNumber(current),
		"highestBlock":  rpc.NewHexNumber(height),
		"pulledStates":  rpc.NewHexNumber(pulled),
		"knownStates":   rpc.NewHexNumber(known),
	}, nil
}

// PublicMinerAPI provides an API to control the miner.
// It offers only methods that operate on data that pose no security risk when it is publicly accessible.
type PublicMinerAPI struct {
	e     *Teslafunds
	agent *miner.RemoteAgent
}

// NewPublicMinerAPI create a new PublicMinerAPI instance.
func NewPublicMinerAPI(e *Teslafunds) *PublicMinerAPI {
	agent := miner.NewRemoteAgent()
	e.Miner().Register(agent)

	return &PublicMinerAPI{e, agent}
}

// Mining returns an indication if this node is currently mining.
func (s *PublicMinerAPI) Mining() bool {
	return s.e.IsMining()
}

// SubmitWork can be used by external miner to submit their POW solution. It returns an indication if the work was
// accepted. Note, this is not an indication if the provided work was valid!
func (s *PublicMinerAPI) SubmitWork(nonce rpc.HexNumber, solution, digest common.Hash) bool {
	return s.agent.SubmitWork(nonce.Uint64(), digest, solution)
}

// GetWork returns a work package for external miner. The work package consists of 3 strings
// result[0], 32 bytes hex encoded current block header pow-hash
// result[1], 32 bytes hex encoded seed hash used for DAG
// result[2], 32 bytes hex encoded boundary condition ("target"), 2^256/difficulty
func (s *PublicMinerAPI) GetWork() (work [3]string, err error) {
	if !s.e.IsMining() {
		if err := s.e.StartMining(0, ""); err != nil {
			return work, err
		}
	}
	if work, err = s.agent.GetWork(); err == nil {
		return
	}
	glog.V(logger.Debug).Infof("%v", err)
	return work, fmt.Errorf("mining not ready")
}

// SubmitHashrate can be used for remote miners to submit their hash rate. This enables the node to report the combined
// hash rate of all miners which submit work through this node. It accepts the miner hash rate and an identifier which
// must be unique between nodes.
func (s *PublicMinerAPI) SubmitHashrate(hashrate rpc.HexNumber, id common.Hash) bool {
	s.agent.SubmitHashrate(id, hashrate.Uint64())
	return true
}

// PrivateMinerAPI provides private RPC methods to control the miner.
// These methods can be abused by external users and must be considered insecure for use by untrusted users.
type PrivateMinerAPI struct {
	e *Teslafunds
}

// NewPrivateMinerAPI create a new RPC service which controls the miner of this node.
func NewPrivateMinerAPI(e *Teslafunds) *PrivateMinerAPI {
	return &PrivateMinerAPI{e: e}
}

// Start the miner with the given number of threads. If threads is nil the number of
// workers started is equal to the number of logical CPU's that are usable by this process.
func (s *PrivateMinerAPI) Start(threads *rpc.HexNumber) (bool, error) {
	s.e.StartAutoDAG()

	if threads == nil {
		threads = rpc.NewHexNumber(runtime.NumCPU())
	}

	err := s.e.StartMining(threads.Int(), "")
	if err == nil {
		return true, nil
	}
	return false, err
}

// Stop the miner
func (s *PrivateMinerAPI) Stop() bool {
	s.e.StopMining()
	return true
}

// SetExtra sets the extra data string that is included when this miner mines a block.
func (s *PrivateMinerAPI) SetExtra(extra string) (bool, error) {
	if err := s.e.Miner().SetExtra([]byte(extra)); err != nil {
		return false, err
	}
	return true, nil
}

// SetGasPrice sets the minimum accepted gas price for the miner.
func (s *PrivateMinerAPI) SetGasPrice(gasPrice rpc.HexNumber) bool {
	s.e.Miner().SetGasPrice(gasPrice.BigInt())
	return true
}

// SetEtherbase sets the etherbase of the miner
func (s *PrivateMinerAPI) SetEtherbase(etherbase common.Address) bool {
	s.e.SetEtherbase(etherbase)
	return true
}

// StartAutoDAG starts auto DAG generation. This will prevent the DAG generating on epoch change
// which will cause the node to stop mining during the generation process.
func (s *PrivateMinerAPI) StartAutoDAG() bool {
	s.e.StartAutoDAG()
	return true
}

// StopAutoDAG stops auto DAG generation
func (s *PrivateMinerAPI) StopAutoDAG() bool {
	s.e.StopAutoDAG()
	return true
}

// MakeDAG creates the new DAG for the given block number
func (s *PrivateMinerAPI) MakeDAG(blockNr rpc.BlockNumber) (bool, error) {
	if err := ethash.MakeDAG(uint64(blockNr.Int64()), ""); err != nil {
		return false, err
	}
	return true, nil
}

// PublicTxPoolAPI offers and API for the transaction pool. It only operates on data that is non confidential.
type PublicTxPoolAPI struct {
	e *Teslafunds
}

// NewPublicTxPoolAPI creates a new tx pool service that gives information about the transaction pool.
func NewPublicTxPoolAPI(e *Teslafunds) *PublicTxPoolAPI {
	return &PublicTxPoolAPI{e}
}

// Content returns the transactions contained within the transaction pool.
func (s *PublicTxPoolAPI) Content() map[string]map[string]map[string][]*RPCTransaction {
	content := map[string]map[string]map[string][]*RPCTransaction{
		"pending": make(map[string]map[string][]*RPCTransaction),
		"queued":  make(map[string]map[string][]*RPCTransaction),
	}
	pending, queue := s.e.TxPool().Content()

	// Flatten the pending transactions
	for account, batches := range pending {
		dump := make(map[string][]*RPCTransaction)
		for _, tx := range batches {
			nonce := fmt.Sprintf("%d", tx.Nonce())
			dump[nonce] = []*RPCTransaction{newRPCPendingTransaction(tx)}
		}
		content["pending"][account.Hex()] = dump
	}
	// Flatten the queued transactions
	for account, batches := range queue {
		dump := make(map[string][]*RPCTransaction)
		for _, tx := range batches {
			nonce := fmt.Sprintf("%d", tx.Nonce())
			dump[nonce] = []*RPCTransaction{newRPCPendingTransaction(tx)}
		}
		content["queued"][account.Hex()] = dump
	}
	return content
}

// Status returns the number of pending and queued transaction in the pool.
func (s *PublicTxPoolAPI) Status() map[string]*rpc.HexNumber {
	pending, queue := s.e.TxPool().Stats()
	return map[string]*rpc.HexNumber{
		"pending": rpc.NewHexNumber(pending),
		"queued":  rpc.NewHexNumber(queue),
	}
}

// Inspect retrieves the content of the transaction pool and flattens it into an
// easily inspectable list.
func (s *PublicTxPoolAPI) Inspect() map[string]map[string]map[string][]string {
	content := map[string]map[string]map[string][]string{
		"pending": make(map[string]map[string][]string),
		"queued":  make(map[string]map[string][]string),
	}
	pending, queue := s.e.TxPool().Content()

	// Define a formatter to flatten a transaction into a string
	var format = func(tx *types.Transaction) string {
		if to := tx.To(); to != nil {
			return fmt.Sprintf("%s: %v wei + %v ?? %v gas", tx.To().Hex(), tx.Value(), tx.Gas(), tx.GasPrice())
		}
		return fmt.Sprintf("contract creation: %v wei + %v ?? %v gas", tx.Value(), tx.Gas(), tx.GasPrice())
	}
	// Flatten the pending transactions
	for account, batches := range pending {
		dump := make(map[string][]string)
		for _, tx := range batches {
			nonce := fmt.Sprintf("%d", tx.Nonce())
			dump[nonce] = []string{format(tx)}
		}
		content["pending"][account.Hex()] = dump
	}
	// Flatten the queued transactions
	for account, batches := range queue {
		dump := make(map[string][]string)
		for _, tx := range batches {
			nonce := fmt.Sprintf("%d", tx.Nonce())
			dump[nonce] = []string{format(tx)}
		}
		content["queued"][account.Hex()] = dump
	}
	return content
}

// PublicAccountAPI provides an API to access accounts managed by this node.
// It offers only methods that can retrieve accounts.
type PublicAccountAPI struct {
	am *accounts.Manager
}

// NewPublicAccountAPI creates a new PublicAccountAPI.
func NewPublicAccountAPI(am *accounts.Manager) *PublicAccountAPI {
	return &PublicAccountAPI{am: am}
}

// Accounts returns the collection of accounts this node manages
func (s *PublicAccountAPI) Accounts() []accounts.Account {
	return s.am.Accounts()
}

// PrivateAccountAPI provides an API to access accounts managed by this node.
// It offers methods to create, (un)lock en list accounts. Some methods accept
// passwords and are therefore considered private by default.
type PrivateAccountAPI struct {
	am     *accounts.Manager
	txPool *core.TxPool
	txMu   *sync.Mutex
	gpo    *GasPriceOracle
}

// NewPrivateAccountAPI create a new PrivateAccountAPI.
func NewPrivateAccountAPI(e *Teslafunds) *PrivateAccountAPI {
	return &PrivateAccountAPI{
		am:     e.accountManager,
		txPool: e.txPool,
		txMu:   &e.txMu,
		gpo:    e.gpo,
	}
}

// ListAccounts will return a list of addresses for accounts this node manages.
func (s *PrivateAccountAPI) ListAccounts() []common.Address {
	accounts := s.am.Accounts()
	addresses := make([]common.Address, len(accounts))
	for i, acc := range accounts {
		addresses[i] = acc.Address
	}
	return addresses
}

// NewAccount will create a new account and returns the address for the new account.
func (s *PrivateAccountAPI) NewAccount(password string) (common.Address, error) {
	acc, err := s.am.NewAccount(password)
	if err == nil {
		return acc.Address, nil
	}
	return common.Address{}, err
}

// ImportRawKey stores the given hex encoded ECDSA key into the key directory,
// encrypting it with the passphrase.
func (s *PrivateAccountAPI) ImportRawKey(privkey string, password string) (common.Address, error) {
	hexkey, err := hex.DecodeString(privkey)
	if err != nil {
		return common.Address{}, err
	}

	acc, err := s.am.ImportECDSA(crypto.ToECDSA(hexkey), password)
	return acc.Address, err
}

// UnlockAccount will unlock the account associated with the given address with
// the given password for duration seconds. If duration is nil it will use a
// default of 300 seconds. It returns an indication if the account was unlocked.
func (s *PrivateAccountAPI) UnlockAccount(addr common.Address, password string, duration *rpc.HexNumber) (bool, error) {
	if duration == nil {
		duration = rpc.NewHexNumber(300)
	}
	a := accounts.Account{Address: addr}
	d := time.Duration(duration.Int64()) * time.Second
	if err := s.am.TimedUnlock(a, password, d); err != nil {
		return false, err
	}
	return true, nil
}

// LockAccount will lock the account associated with the given address when it's unlocked.
func (s *PrivateAccountAPI) LockAccount(addr common.Address) bool {
	return s.am.Lock(addr) == nil
}

// SignAndSendTransaction will create a transaction from the given arguments and
// tries to sign it with the key associated with args.To. If the given passwd isn't
// able to decrypt the key it fails.
func (s *PrivateAccountAPI) SignAndSendTransaction(args SendTxArgs, passwd string) (common.Hash, error) {
	args = prepareSendTxArgs(args, s.gpo)

	s.txMu.Lock()
	defer s.txMu.Unlock()

	if args.Nonce == nil {
		args.Nonce = rpc.NewHexNumber(s.txPool.State().GetNonce(args.From))
	}

	var tx *types.Transaction
	if args.To == nil {
		tx = types.NewContractCreation(args.Nonce.Uint64(), args.Value.BigInt(), args.Gas.BigInt(), args.GasPrice.BigInt(), common.FromHex(args.Data))
	} else {
		tx = types.NewTransaction(args.Nonce.Uint64(), *args.To, args.Value.BigInt(), args.Gas.BigInt(), args.GasPrice.BigInt(), common.FromHex(args.Data))
	}

	signature, err := s.am.SignWithPassphrase(args.From, passwd, tx.SigHash().Bytes())
	if err != nil {
		return common.Hash{}, err
	}

	return submitTransaction(s.txPool, tx, signature)
}

// PublicBlockChainAPI provides an API to access the Ethereum blockchain.
// It offers only methods that operate on public data that is freely available to anyone.
type PublicBlockChainAPI struct {
	config                  *core.ChainConfig
	bc                      *core.BlockChain
	chainDb                 ethdb.Database
	eventMux                *event.TypeMux
	muNewBlockSubscriptions sync.Mutex                             // protects newBlocksSubscriptions
	newBlockSubscriptions   map[string]func(core.ChainEvent) error // callbacks for new block subscriptions
	am                      *accounts.Manager
	miner                   *miner.Miner
	gpo                     *GasPriceOracle
}

// NewPublicBlockChainAPI creates a new Etheruem blockchain API.
func NewPublicBlockChainAPI(config *core.ChainConfig, bc *core.BlockChain, m *miner.Miner, chainDb ethdb.Database, gpo *GasPriceOracle, eventMux *event.TypeMux, am *accounts.Manager) *PublicBlockChainAPI {
	api := &PublicBlockChainAPI{
		config:   config,
		bc:       bc,
		miner:    m,
		chainDb:  chainDb,
		eventMux: eventMux,
		am:       am,
		newBlockSubscriptions: make(map[string]func(core.ChainEvent) error),
		gpo: gpo,
	}

	go api.subscriptionLoop()

	return api
}

// subscriptionLoop reads events from the global event mux and creates notifications for the matched subscriptions.
func (s *PublicBlockChainAPI) subscriptionLoop() {
	sub := s.eventMux.Subscribe(core.ChainEvent{})
	for event := range sub.Chan() {
		if chainEvent, ok := event.Data.(core.ChainEvent); ok {
			s.muNewBlockSubscriptions.Lock()
			for id, notifyOf := range s.newBlockSubscriptions {
				if notifyOf(chainEvent) == rpc.ErrNotificationNotFound {
					delete(s.newBlockSubscriptions, id)
				}
			}
			s.muNewBlockSubscriptions.Unlock()
		}
	}
}

// BlockNumber returns the block number of the chain head.
func (s *PublicBlockChainAPI) BlockNumber() *big.Int {
	return s.bc.CurrentHeader().Number
}

// GetBalance returns the amount of wei for the given address in the state of the
// given block number. The rpc.LatestBlockNumber and rpc.PendingBlockNumber meta
// block numbers are also allowed.
func (s *PublicBlockChainAPI) GetBalance(address common.Address, blockNr rpc.BlockNumber) (*big.Int, error) {
	state, _, err := stateAndBlockByNumber(s.miner, s.bc, blockNr, s.chainDb)
	if state == nil || err != nil {
		return nil, err
	}
	return state.GetBalance(address), nil
}

// GetBlockByNumber returns the requested block. When blockNr is -1 the chain head is returned. When fullTx is true all
// transactions in the block are returned in full detail, otherwise only the transaction hash is returned.
func (s *PublicBlockChainAPI) GetBlockByNumber(blockNr rpc.BlockNumber, fullTx bool) (map[string]interface{}, error) {
	if block := blockByNumber(s.miner, s.bc, blockNr); block != nil {
		response, err := s.rpcOutputBlock(block, true, fullTx)
		if err == nil && blockNr == rpc.PendingBlockNumber {
			// Pending blocks need to nil out a few fields
			for _, field := range []string{"hash", "nonce", "logsBloom", "miner"} {
				response[field] = nil
			}
		}
		return response, err
	}
	return nil, nil
}

// GetBlockByHash returns the requested block. When fullTx is true all transactions in the block are returned in full
// detail, otherwise only the transaction hash is returned.
func (s *PublicBlockChainAPI) GetBlockByHash(blockHash common.Hash, fullTx bool) (map[string]interface{}, error) {
	if block := s.bc.GetBlock(blockHash); block != nil {
		return s.rpcOutputBlock(block, true, fullTx)
	}
	return nil, nil
}

// GetUncleByBlockNumberAndIndex returns the uncle block for the given block hash and index. When fullTx is true
// all transactions in the block are returned in full detail, otherwise only the transaction hash is returned.
func (s *PublicBlockChainAPI) GetUncleByBlockNumberAndIndex(blockNr rpc.BlockNumber, index rpc.HexNumber) (map[string]interface{}, error) {
	if block := blockByNumber(s.miner, s.bc, blockNr); block != nil {
		uncles := block.Uncles()
		if index.Int() < 0 || index.Int() >= len(uncles) {
			glog.V(logger.Debug).Infof("uncle block on index %d not found for block #%d", index.Int(), blockNr)
			return nil, nil
		}
		block = types.NewBlockWithHeader(uncles[index.Int()])
		return s.rpcOutputBlock(block, false, false)
	}
	return nil, nil
}

// GetUncleByBlockHashAndIndex returns the uncle block for the given block hash and index. When fullTx is true
// all transactions in the block are returned in full detail, otherwise only the transaction hash is returned.
func (s *PublicBlockChainAPI) GetUncleByBlockHashAndIndex(blockHash common.Hash, index rpc.HexNumber) (map[string]interface{}, error) {
	if block := s.bc.GetBlock(blockHash); block != nil {
		uncles := block.Uncles()
		if index.Int() < 0 || index.Int() >= len(uncles) {
			glog.V(logger.Debug).Infof("uncle block on index %d not found for block %s", index.Int(), blockHash.Hex())
			return nil, nil
		}
		block = types.NewBlockWithHeader(uncles[index.Int()])
		return s.rpcOutputBlock(block, false, false)
	}
	return nil, nil
}

// GetUncleCountByBlockNumber returns number of uncles in the block for the given block number
func (s *PublicBlockChainAPI) GetUncleCountByBlockNumber(blockNr rpc.BlockNumber) *rpc.HexNumber {
	if block := blockByNumber(s.miner, s.bc, blockNr); block != nil {
		return rpc.NewHexNumber(len(block.Uncles()))
	}
	return nil
}

// GetUncleCountByBlockHash returns number of uncles in the block for the given block hash
func (s *PublicBlockChainAPI) GetUncleCountByBlockHash(blockHash common.Hash) *rpc.HexNumber {
	if block := s.bc.GetBlock(blockHash); block != nil {
		return rpc.NewHexNumber(len(block.Uncles()))
	}
	return nil
}

// NewBlocksArgs allows the user to specify if the returned block should include transactions and in which format.
type NewBlocksArgs struct {
	IncludeTransactions bool `json:"includeTransactions"`
	TransactionDetails  bool `json:"transactionDetails"`
}

// NewBlocks triggers a new block event each time a block is appended to the chain. It accepts an argument which allows
// the caller to specify whether the output should contain transactions and in what format.
func (s *PublicBlockChainAPI) NewBlocks(ctx context.Context, args NewBlocksArgs) (rpc.Subscription, error) {
	notifier, supported := rpc.NotifierFromContext(ctx)
	if !supported {
		return nil, rpc.ErrNotificationsUnsupported
	}

	// create a subscription that will remove itself when unsubscribed/cancelled
	subscription, err := notifier.NewSubscription(func(subId string) {
		s.muNewBlockSubscriptions.Lock()
		delete(s.newBlockSubscriptions, subId)
		s.muNewBlockSubscriptions.Unlock()
	})

	if err != nil {
		return nil, err
	}

	// add a callback that is called on chain events which will format the block and notify the client
	s.muNewBlockSubscriptions.Lock()
	s.newBlockSubscriptions[subscription.ID()] = func(e core.ChainEvent) error {
		notification, err := s.rpcOutputBlock(e.Block, args.IncludeTransactions, args.TransactionDetails)
		if err == nil {
			return subscription.Notify(notification)
		}
		glog.V(logger.Warn).Info("unable to format block %v\n", err)
		return nil
	}
	s.muNewBlockSubscriptions.Unlock()
	return subscription, nil
}

// GetCode returns the code stored at the given address in the state for the given block number.
func (s *PublicBlockChainAPI) GetCode(address common.Address, blockNr rpc.BlockNumber) (string, error) {
	state, _, err := stateAndBlockByNumber(s.miner, s.bc, blockNr, s.chainDb)
	if state == nil || err != nil {
		return "", err
	}
	res := state.GetCode(address)
	if len(res) == 0 { // backwards compatibility
		return "0x", nil
	}
	return common.ToHex(res), nil
}

// GetStorageAt returns the storage from the state at the given address, key and
// block number. The rpc.LatestBlockNumber and rpc.PendingBlockNumber meta block
// numbers are also allowed.
func (s *PublicBlockChainAPI) GetStorageAt(address common.Address, key string, blockNr rpc.BlockNumber) (string, error) {
	state, _, err := stateAndBlockByNumber(s.miner, s.bc, blockNr, s.chainDb)
	if state == nil || err != nil {
		return "0x", err
	}
	return state.GetState(address, common.HexToHash(key)).Hex(), nil
}

// callmsg is the message type used for call transactions.
type callmsg struct {
	from          *state.StateObject
	to            *common.Address
	gas, gasPrice *big.Int
	value         *big.Int
	data          []byte
}

// accessor boilerplate to implement core.Message
func (m callmsg) From() (common.Address, error)         { return m.from.Address(), nil }
func (m callmsg) FromFrontier() (common.Address, error) { return m.from.Address(), nil }
func (m callmsg) Nonce() uint64                         { return m.from.Nonce() }
func (m callmsg) To() *common.Address                   { return m.to }
func (m callmsg) GasPrice() *big.Int                    { return m.gasPrice }
func (m callmsg) Gas() *big.Int                         { return m.gas }
func (m callmsg) Value() *big.Int                       { return m.value }
func (m callmsg) Data() []byte                          { return m.data }

// CallArgs represents the arguments for a call.
type CallArgs struct {
	From     common.Address  `json:"from"`
	To       *common.Address `json:"to"`
	Gas      *rpc.HexNumber  `json:"gas"`
	GasPrice *rpc.HexNumber  `json:"gasPrice"`
	Value    rpc.HexNumber   `json:"value"`
	Data     string          `json:"data"`
}

func (s *PublicBlockChainAPI) doCall(args CallArgs, blockNr rpc.BlockNumber) (string, *big.Int, error) {
	// Fetch the state associated with the block number
	stateDb, block, err := stateAndBlockByNumber(s.miner, s.bc, blockNr, s.chainDb)
	if stateDb == nil || err != nil {
		return "0x", nil, err
	}
	stateDb = stateDb.Copy()

	// Retrieve the account state object to interact with
	var from *state.StateObject
	if args.From == (common.Address{}) {
		accounts := s.am.Accounts()
		if len(accounts) == 0 {
			from = stateDb.GetOrNewStateObject(common.Address{})
		} else {
			from = stateDb.GetOrNewStateObject(accounts[0].Address)
		}
	} else {
		from = stateDb.GetOrNewStateObject(args.From)
	}
	from.SetBalance(common.MaxBig)

	// Assemble the CALL invocation
	msg := callmsg{
		from:     from,
		to:       args.To,
		gas:      args.Gas.BigInt(),
		gasPrice: args.GasPrice.BigInt(),
		value:    args.Value.BigInt(),
		data:     common.FromHex(args.Data),
	}
	if msg.gas == nil {
		msg.gas = big.NewInt(50000000)
	}
	if msg.gasPrice == nil {
		msg.gasPrice = s.gpo.SuggestPrice()
	}

	// Execute the call and return
	vmenv := core.NewEnv(stateDb, s.config, s.bc, msg, block.Header(), s.config.VmConfig)
	gp := new(core.GasPool).AddGas(common.MaxBig)

	res, requiredGas, _, err := core.NewStateTransition(vmenv, msg, gp).TransitionDb()
	if len(res) == 0 { // backwards compatibility
		return "0x", requiredGas, err
	}
	return common.ToHex(res), requiredGas, err
}

// Call executes the given transaction on the state for the given block number.
// It doesn't make and changes in the state/blockchain and is useful to execute and retrieve values.
func (s *PublicBlockChainAPI) Call(args CallArgs, blockNr rpc.BlockNumber) (string, error) {
	result, _, err := s.doCall(args, blockNr)
	return result, err
}

// EstimateGas returns an estimate of the amount of gas needed to execute the given transaction.
func (s *PublicBlockChainAPI) EstimateGas(args CallArgs) (*rpc.HexNumber, error) {
	_, gas, err := s.doCall(args, rpc.PendingBlockNumber)
	return rpc.NewHexNumber(gas), err
}

// rpcOutputBlock converts the given block to the RPC output which depends on fullTx. If inclTx is true transactions are
// returned. When fullTx is true the returned block contains full transaction details, otherwise it will only contain
// transaction hashes.
func (s *PublicBlockChainAPI) rpcOutputBlock(b *types.Block, inclTx bool, fullTx bool) (map[string]interface{}, error) {
	fields := map[string]interface{}{
		"number":           rpc.NewHexNumber(b.Number()),
		"hash":             b.Hash(),
		"parentHash":       b.ParentHash(),
		"nonce":            b.Header().Nonce,
		"sha3Uncles":       b.UncleHash(),
		"logsBloom":        b.Bloom(),
		"stateRoot":        b.Root(),
		"miner":            b.Coinbase(),
		"difficulty":       rpc.NewHexNumber(b.Difficulty()),
		"totalDifficulty":  rpc.NewHexNumber(s.bc.GetTd(b.Hash())),
		"extraData":        fmt.Sprintf("0x%x", b.Extra()),
		"size":             rpc.NewHexNumber(b.Size().Int64()),
		"gasLimit":         rpc.NewHexNumber(b.GasLimit()),
		"gasUsed":          rpc.NewHexNumber(b.GasUsed()),
		"timestamp":        rpc.NewHexNumber(b.Time()),
		"transactionsRoot": b.TxHash(),
		"receiptRoot":      b.ReceiptHash(),
	}

	if inclTx {
		formatTx := func(tx *types.Transaction) (interface{}, error) {
			return tx.Hash(), nil
		}

		if fullTx {
			formatTx = func(tx *types.Transaction) (interface{}, error) {
				return newRPCTransaction(b, tx.Hash())
			}
		}

		txs := b.Transactions()
		transactions := make([]interface{}, len(txs))
		var err error
		for i, tx := range b.Transactions() {
			if transactions[i], err = formatTx(tx); err != nil {
				return nil, err
			}
		}
		fields["transactions"] = transactions
	}

	uncles := b.Uncles()
	uncleHashes := make([]common.Hash, len(uncles))
	for i, uncle := range uncles {
		uncleHashes[i] = uncle.Hash()
	}
	fields["uncles"] = uncleHashes

	return fields, nil
}

// RPCTransaction represents a transaction that will serialize to the RPC representation of a transaction
type RPCTransaction struct {
	BlockHash        common.Hash     `json:"blockHash"`
	BlockNumber      *rpc.HexNumber  `json:"blockNumber"`
	From             common.Address  `json:"from"`
	Gas              *rpc.HexNumber  `json:"gas"`
	GasPrice         *rpc.HexNumber  `json:"gasPrice"`
	Hash             common.Hash     `json:"hash"`
	Input            string          `json:"input"`
	Nonce            *rpc.HexNumber  `json:"nonce"`
	To               *common.Address `json:"to"`
	TransactionIndex *rpc.HexNumber  `json:"transactionIndex"`
	Value            *rpc.HexNumber  `json:"value"`
}

// newRPCPendingTransaction returns a pending transaction that will serialize to the RPC representation
func newRPCPendingTransaction(tx *types.Transaction) *RPCTransaction {
	from, _ := tx.FromFrontier()

	return &RPCTransaction{
		From:     from,
		Gas:      rpc.NewHexNumber(tx.Gas()),
		GasPrice: rpc.NewHexNumber(tx.GasPrice()),
		Hash:     tx.Hash(),
		Input:    fmt.Sprintf("0x%x", tx.Data()),
		Nonce:    rpc.NewHexNumber(tx.Nonce()),
		To:       tx.To(),
		Value:    rpc.NewHexNumber(tx.Value()),
	}
}

// newRPCTransaction returns a transaction that will serialize to the RPC representation.
func newRPCTransactionFromBlockIndex(b *types.Block, txIndex int) (*RPCTransaction, error) {
	if txIndex >= 0 && txIndex < len(b.Transactions()) {
		tx := b.Transactions()[txIndex]
		from, err := tx.FromFrontier()
		if err != nil {
			return nil, err
		}

		return &RPCTransaction{
			BlockHash:        b.Hash(),
			BlockNumber:      rpc.NewHexNumber(b.Number()),
			From:             from,
			Gas:              rpc.NewHexNumber(tx.Gas()),
			GasPrice:         rpc.NewHexNumber(tx.GasPrice()),
			Hash:             tx.Hash(),
			Input:            fmt.Sprintf("0x%x", tx.Data()),
			Nonce:            rpc.NewHexNumber(tx.Nonce()),
			To:               tx.To(),
			TransactionIndex: rpc.NewHexNumber(txIndex),
			Value:            rpc.NewHexNumber(tx.Value()),
		}, nil
	}

	return nil, nil
}

// newRPCTransaction returns a transaction that will serialize to the RPC representation.
func newRPCTransaction(b *types.Block, txHash common.Hash) (*RPCTransaction, error) {
	for idx, tx := range b.Transactions() {
		if tx.Hash() == txHash {
			return newRPCTransactionFromBlockIndex(b, idx)
		}
	}

	return nil, nil
}

// PublicTransactionPoolAPI exposes methods for the RPC interface
type PublicTransactionPoolAPI struct {
	eventMux        *event.TypeMux
	chainDb         ethdb.Database
	gpo             *GasPriceOracle
	bc              *core.BlockChain
	miner           *miner.Miner
	am              *accounts.Manager
	txPool          *core.TxPool
	txMu            *sync.Mutex
	muPendingTxSubs sync.Mutex
	pendingTxSubs   map[string]rpc.Subscription
}

// NewPublicTransactionPoolAPI creates a new RPC service with methods specific for the transaction pool.
func NewPublicTransactionPoolAPI(e *Teslafunds) *PublicTransactionPoolAPI {
	api := &PublicTransactionPoolAPI{
		eventMux:      e.eventMux,
		gpo:           e.gpo,
		chainDb:       e.chainDb,
		bc:            e.blockchain,
		am:            e.accountManager,
		txPool:        e.txPool,
		txMu:          &e.txMu,
		miner:         e.miner,
		pendingTxSubs: make(map[string]rpc.Subscription),
	}
	go api.subscriptionLoop()

	return api
}

// subscriptionLoop listens for events on the global event mux and creates notifications for subscriptions.
func (s *PublicTransactionPoolAPI) subscriptionLoop() {
	sub := s.eventMux.Subscribe(core.TxPreEvent{})
	for event := range sub.Chan() {
		tx := event.Data.(core.TxPreEvent)
		if from, err := tx.Tx.FromFrontier(); err == nil {
			if s.am.HasAddress(from) {
				s.muPendingTxSubs.Lock()
				for id, sub := range s.pendingTxSubs {
					if sub.Notify(tx.Tx.Hash()) == rpc.ErrNotificationNotFound {
						delete(s.pendingTxSubs, id)
					}
				}
				s.muPendingTxSubs.Unlock()
			}
		}
	}
}

func getTransaction(chainDb ethdb.Database, txPool *core.TxPool, txHash common.Hash) (*types.Transaction, bool, error) {
	txData, err := chainDb.Get(txHash.Bytes())
	isPending := false
	tx := new(types.Transaction)

	if err == nil && len(txData) > 0 {
		if err := rlp.DecodeBytes(txData, tx); err != nil {
			return nil, isPending, err
		}
	} else {
		// pending transaction?
		tx = txPool.Get(txHash)
		isPending = true
	}

	return tx, isPending, nil
}

// GetBlockTransactionCountByNumber returns the number of transactions in the block with the given block number.
func (s *PublicTransactionPoolAPI) GetBlockTransactionCountByNumber(blockNr rpc.BlockNumber) *rpc.HexNumber {
	if block := blockByNumber(s.miner, s.bc, blockNr); block != nil {
		return rpc.NewHexNumber(len(block.Transactions()))
	}
	return nil
}

// GetBlockTransactionCountByHash returns the number of transactions in the block with the given hash.
func (s *PublicTransactionPoolAPI) GetBlockTransactionCountByHash(blockHash common.Hash) *rpc.HexNumber {
	if block := s.bc.GetBlock(blockHash); block != nil {
		return rpc.NewHexNumber(len(block.Transactions()))
	}
	return nil
}

// GetTransactionByBlockNumberAndIndex returns the transaction for the given block number and index.
func (s *PublicTransactionPoolAPI) GetTransactionByBlockNumberAndIndex(blockNr rpc.BlockNumber, index rpc.HexNumber) (*RPCTransaction, error) {
	if block := blockByNumber(s.miner, s.bc, blockNr); block != nil {
		return newRPCTransactionFromBlockIndex(block, index.Int())
	}
	return nil, nil
}

// GetTransactionByBlockHashAndIndex returns the transaction for the given block hash and index.
func (s *PublicTransactionPoolAPI) GetTransactionByBlockHashAndIndex(blockHash common.Hash, index rpc.HexNumber) (*RPCTransaction, error) {
	if block := s.bc.GetBlock(blockHash); block != nil {
		return newRPCTransactionFromBlockIndex(block, index.Int())
	}
	return nil, nil
}

// GetTransactionCount returns the number of transactions the given address has sent for the given block number
func (s *PublicTransactionPoolAPI) GetTransactionCount(address common.Address, blockNr rpc.BlockNumber) (*rpc.HexNumber, error) {
	state, _, err := stateAndBlockByNumber(s.miner, s.bc, blockNr, s.chainDb)
	if state == nil || err != nil {
		return nil, err
	}
	return rpc.NewHexNumber(state.GetNonce(address)), nil
}

// getTransactionBlockData fetches the meta data for the given transaction from the chain database. This is useful to
// retrieve block information for a hash. It returns the block hash, block index and transaction index.
func getTransactionBlockData(chainDb ethdb.Database, txHash common.Hash) (common.Hash, uint64, uint64, error) {
	var txBlock struct {
		BlockHash  common.Hash
		BlockIndex uint64
		Index      uint64
	}

	blockData, err := chainDb.Get(append(txHash.Bytes(), 0x0001))
	if err != nil {
		return common.Hash{}, uint64(0), uint64(0), err
	}

	reader := bytes.NewReader(blockData)
	if err = rlp.Decode(reader, &txBlock); err != nil {
		return common.Hash{}, uint64(0), uint64(0), err
	}

	return txBlock.BlockHash, txBlock.BlockIndex, txBlock.Index, nil
}

// GetTransactionByHash returns the transaction for the given hash
func (s *PublicTransactionPoolAPI) GetTransactionByHash(txHash common.Hash) (*RPCTransaction, error) {
	var tx *types.Transaction
	var isPending bool
	var err error

	if tx, isPending, err = getTransaction(s.chainDb, s.txPool, txHash); err != nil {
		glog.V(logger.Debug).Infof("%v\n", err)
		return nil, nil
	} else if tx == nil {
		return nil, nil
	}

	if isPending {
		return newRPCPendingTransaction(tx), nil
	}

	blockHash, _, _, err := getTransactionBlockData(s.chainDb, txHash)
	if err != nil {
		glog.V(logger.Debug).Infof("%v\n", err)
		return nil, nil
	}

	if block := s.bc.GetBlock(blockHash); block != nil {
		return newRPCTransaction(block, txHash)
	}

	return nil, nil
}

// GetTransactionReceipt returns the transaction receipt for the given transaction hash.
func (s *PublicTransactionPoolAPI) GetTransactionReceipt(txHash common.Hash) (map[string]interface{}, error) {
	receipt := core.GetReceipt(s.chainDb, txHash)
	if receipt == nil {
		glog.V(logger.Debug).Infof("receipt not found for transaction %s", txHash.Hex())
		return nil, nil
	}

	tx, _, err := getTransaction(s.chainDb, s.txPool, txHash)
	if err != nil {
		glog.V(logger.Debug).Infof("%v\n", err)
		return nil, nil
	}

	txBlock, blockIndex, index, err := getTransactionBlockData(s.chainDb, txHash)
	if err != nil {
		glog.V(logger.Debug).Infof("%v\n", err)
		return nil, nil
	}

	from, err := tx.FromFrontier()
	if err != nil {
		glog.V(logger.Debug).Infof("%v\n", err)
		return nil, nil
	}

	fields := map[string]interface{}{
		"root":              common.Bytes2Hex(receipt.PostState),
		"blockHash":         txBlock,
		"blockNumber":       rpc.NewHexNumber(blockIndex),
		"transactionHash":   txHash,
		"transactionIndex":  rpc.NewHexNumber(index),
		"from":              from,
		"to":                tx.To(),
		"gasUsed":           rpc.NewHexNumber(receipt.GasUsed),
		"cumulativeGasUsed": rpc.NewHexNumber(receipt.CumulativeGasUsed),
		"contractAddress":   nil,
		"logs":              receipt.Logs,
	}

	if receipt.Logs == nil {
		fields["logs"] = []vm.Logs{}
	}
=======
	"github.com/ethereum/ethash"
	"github.com/dubaicoin-dbix/go-dubaicoin/common"
	"github.com/dubaicoin-dbix/go-dubaicoin/common/hexutil"
	"github.com/dubaicoin-dbix/go-dubaicoin/core"
	"github.com/dubaicoin-dbix/go-dubaicoin/core/state"
	"github.com/dubaicoin-dbix/go-dubaicoin/core/types"
	"github.com/dubaicoin-dbix/go-dubaicoin/core/vm"
	"github.com/dubaicoin-dbix/go-dubaicoin/internal/ethapi"
	"github.com/dubaicoin-dbix/go-dubaicoin/logger"
	"github.com/dubaicoin-dbix/go-dubaicoin/logger/glog"
	"github.com/dubaicoin-dbix/go-dubaicoin/miner"
	"github.com/dubaicoin-dbix/go-dubaicoin/params"
	"github.com/dubaicoin-dbix/go-dubaicoin/rlp"
	"github.com/dubaicoin-dbix/go-dubaicoin/rpc"
	"golang.org/x/net/context"
)

const defaultTraceTimeout = 5 * time.Second
>>>>>>> 7fdd714... gdbix-update v1.5.0:dbix/api.go

// PublicEthereumAPI provides an API to access Ethereum full node-related
// information.
type PublicEthereumAPI struct {
	e *Ethereum
}

// NewPublicEthereumAPI creates a new Etheruem protocol API for full nodes.
func NewPublicEthereumAPI(e *Ethereum) *PublicEthereumAPI {
	return &PublicEthereumAPI{e}
}

// Etherbase is the address that mining rewards will be send to
func (s *PublicEthereumAPI) Etherbase() (common.Address, error) {
	return s.e.Etherbase()
}

// Coinbase is the address that mining rewards will be send to (alias for Etherbase)
func (s *PublicEthereumAPI) Coinbase() (common.Address, error) {
	return s.Etherbase()
}

// Hashrate returns the POW hashrate
func (s *PublicEthereumAPI) Hashrate() hexutil.Uint64 {
	return hexutil.Uint64(s.e.Miner().HashRate())
}

// PublicMinerAPI provides an API to control the miner.
// It offers only methods that operate on data that pose no security risk when it is publicly accessible.
type PublicMinerAPI struct {
	e     *Ethereum
	agent *miner.RemoteAgent
}

// NewPublicMinerAPI create a new PublicMinerAPI instance.
func NewPublicMinerAPI(e *Ethereum) *PublicMinerAPI {
	agent := miner.NewRemoteAgent(e.Pow())
	e.Miner().Register(agent)

	return &PublicMinerAPI{e, agent}
}

// Mining returns an indication if this node is currently mining.
func (s *PublicMinerAPI) Mining() bool {
	return s.e.IsMining()
}

// SubmitWork can be used by external miner to submit their POW solution. It returns an indication if the work was
// accepted. Note, this is not an indication if the provided work was valid!
func (s *PublicMinerAPI) SubmitWork(nonce types.BlockNonce, solution, digest common.Hash) bool {
	return s.agent.SubmitWork(nonce, digest, solution)
}

// GetWork returns a work package for external miner. The work package consists of 3 strings
// result[0], 32 bytes hex encoded current block header pow-hash
// result[1], 32 bytes hex encoded seed hash used for DAG
// result[2], 32 bytes hex encoded boundary condition ("target"), 2^256/difficulty
func (s *PublicMinerAPI) GetWork() (work [3]string, err error) {
	if !s.e.IsMining() {
		if err := s.e.StartMining(0); err != nil {
			return work, err
		}
	}
	if work, err = s.agent.GetWork(); err == nil {
		return
	}
	glog.V(logger.Debug).Infof("%v", err)
	return work, fmt.Errorf("mining not ready")
}

// SubmitHashrate can be used for remote miners to submit their hash rate. This enables the node to report the combined
// hash rate of all miners which submit work through this node. It accepts the miner hash rate and an identifier which
// must be unique between nodes.
func (s *PublicMinerAPI) SubmitHashrate(hashrate hexutil.Uint64, id common.Hash) bool {
	s.agent.SubmitHashrate(id, uint64(hashrate))
	return true
}

// PrivateMinerAPI provides private RPC methods to control the miner.
// These methods can be abused by external users and must be considered insecure for use by untrusted users.
type PrivateMinerAPI struct {
	e *Ethereum
}

// NewPrivateMinerAPI create a new RPC service which controls the miner of this node.
func NewPrivateMinerAPI(e *Ethereum) *PrivateMinerAPI {
	return &PrivateMinerAPI{e: e}
}

// Start the miner with the given number of threads. If threads is nil the number of
// workers started is equal to the number of logical CPU's that are usable by this process.
func (s *PrivateMinerAPI) Start(threads *int) (bool, error) {
	s.e.StartAutoDAG()
	var err error
	if threads == nil {
		err = s.e.StartMining(runtime.NumCPU())
	} else {
		err = s.e.StartMining(*threads)
	}
	return err == nil, err
}

// Stop the miner
func (s *PrivateMinerAPI) Stop() bool {
	s.e.StopMining()
	return true
}

// SetExtra sets the extra data string that is included when this miner mines a block.
func (s *PrivateMinerAPI) SetExtra(extra string) (bool, error) {
	if err := s.e.Miner().SetExtra([]byte(extra)); err != nil {
		return false, err
	}
	return true, nil
}

// SetGasPrice sets the minimum accepted gas price for the miner.
func (s *PrivateMinerAPI) SetGasPrice(gasPrice hexutil.Big) bool {
	s.e.Miner().SetGasPrice((*big.Int)(&gasPrice))
	return true
}

// SetEtherbase sets the etherbase of the miner
func (s *PrivateMinerAPI) SetEtherbase(etherbase common.Address) bool {
	s.e.SetEtherbase(etherbase)
	return true
}

// StartAutoDAG starts auto DAG generation. This will prevent the DAG generating on epoch change
// which will cause the node to stop mining during the generation process.
func (s *PrivateMinerAPI) StartAutoDAG() bool {
	s.e.StartAutoDAG()
	return true
}

// StopAutoDAG stops auto DAG generation
func (s *PrivateMinerAPI) StopAutoDAG() bool {
	s.e.StopAutoDAG()
	return true
}

// MakeDAG creates the new DAG for the given block number
func (s *PrivateMinerAPI) MakeDAG(blockNr rpc.BlockNumber) (bool, error) {
	if err := ethash.MakeDAG(uint64(blockNr.Int64()), ""); err != nil {
		return false, err
	}
	return true, nil
}

// PrivateAdminAPI is the collection of Etheruem full node-related APIs
// exposed over the private admin endpoint.
type PrivateAdminAPI struct {
<<<<<<< HEAD:tsf/api.go
	tsf *Teslafunds
}

// NewPrivateAdminAPI creates a new API definition for the private admin methods
// of the Teslafunds service.
func NewPrivateAdminAPI(tsf *Teslafunds) *PrivateAdminAPI {
	return &PrivateAdminAPI{tsf: tsf}
}

// SetSolc sets the Solidity compiler path to be used by the node.
func (api *PrivateAdminAPI) SetSolc(path string) (string, error) {
	solc, err := api.tsf.SetSolc(path)
	if err != nil {
		return "", err
	}
	return solc.Info(), nil
=======
	eth *Ethereum
}

// NewPrivateAdminAPI creates a new API definition for the full node private
// admin methods of the Dubaicoin service.
func NewPrivateAdminAPI(eth *Ethereum) *PrivateAdminAPI {
	return &PrivateAdminAPI{eth: eth}
>>>>>>> 7fdd714... gdbix-update v1.5.0:dbix/api.go
}

// ExportChain exports the current blockchain into a local file.
func (api *PrivateAdminAPI) ExportChain(file string) (bool, error) {
	// Make sure we can create the file to export into
	out, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return false, err
	}
	defer out.Close()

	var writer io.Writer = out
	if strings.HasSuffix(file, ".gz") {
		writer = gzip.NewWriter(writer)
		defer writer.(*gzip.Writer).Close()
	}

	// Export the blockchain
<<<<<<< HEAD:tsf/api.go
	if err := api.tsf.BlockChain().Export(out); err != nil {
=======
	if err := api.eth.BlockChain().Export(writer); err != nil {
>>>>>>> 7fdd714... gdbix-update v1.5.0:dbix/api.go
		return false, err
	}
	return true, nil
}

func hasAllBlocks(chain *core.BlockChain, bs []*types.Block) bool {
	for _, b := range bs {
		if !chain.HasBlock(b.Hash()) {
			return false
		}
	}

	return true
}

// ImportChain imports a blockchain from a local file.
func (api *PrivateAdminAPI) ImportChain(file string) (bool, error) {
	// Make sure the can access the file to import
	in, err := os.Open(file)
	if err != nil {
		return false, err
	}
	defer in.Close()

	var reader io.Reader = in
	if strings.HasSuffix(file, ".gz") {
		if reader, err = gzip.NewReader(reader); err != nil {
			return false, err
		}
	}

	// Run actual the import in pre-configured batches
	stream := rlp.NewStream(reader, 0)

	blocks, index := make([]*types.Block, 0, 2500), 0
	for batch := 0; ; batch++ {
		// Load a batch of blocks from the input file
		for len(blocks) < cap(blocks) {
			block := new(types.Block)
			if err := stream.Decode(block); err == io.EOF {
				break
			} else if err != nil {
				return false, fmt.Errorf("block %d: failed to parse: %v", index, err)
			}
			blocks = append(blocks, block)
			index++
		}
		if len(blocks) == 0 {
			break
		}

<<<<<<< HEAD:tsf/api.go
		if hasAllBlocks(api.tsf.BlockChain(), blocks) {
=======
		if hasAllBlocks(api.eth.BlockChain(), blocks) {
>>>>>>> 7fdd714... gdbix-update v1.5.0:dbix/api.go
			blocks = blocks[:0]
			continue
		}
		// Import the batch and reset the buffer
<<<<<<< HEAD:tsf/api.go
		if _, err := api.tsf.BlockChain().InsertChain(blocks); err != nil {
=======
		if _, err := api.eth.BlockChain().InsertChain(blocks); err != nil {
>>>>>>> 7fdd714... gdbix-update v1.5.0:dbix/api.go
			return false, fmt.Errorf("batch %d: failed to insert: %v", batch, err)
		}
		blocks = blocks[:0]
	}
	return true, nil
}

// PublicDebugAPI is the collection of Etheruem full node APIs exposed
// over the public debugging endpoint.
type PublicDebugAPI struct {
<<<<<<< HEAD:tsf/api.go
	tsf *Teslafunds
}

// NewPublicDebugAPI creates a new API definition for the public debug methods
// of the Teslafunds service.
func NewPublicDebugAPI(tsf *Teslafunds) *PublicDebugAPI {
	return &PublicDebugAPI{tsf: tsf}
=======
	eth *Ethereum
}

// NewPublicDebugAPI creates a new API definition for the full node-
// related public debug methods of the Dubaicoin service.
func NewPublicDebugAPI(eth *Ethereum) *PublicDebugAPI {
	return &PublicDebugAPI{eth: eth}
>>>>>>> 7fdd714... gdbix-update v1.5.0:dbix/api.go
}

// DumpBlock retrieves the entire state of the database at a given block.
func (api *PublicDebugAPI) DumpBlock(number uint64) (state.Dump, error) {
<<<<<<< HEAD:tsf/api.go
	block := api.tsf.BlockChain().GetBlockByNumber(number)
	if block == nil {
		return state.Dump{}, fmt.Errorf("block #%d not found", number)
	}
	stateDb, err := api.tsf.BlockChain().StateAt(block.Root())
=======
	block := api.eth.BlockChain().GetBlockByNumber(number)
	if block == nil {
		return state.Dump{}, fmt.Errorf("block #%d not found", number)
	}
	stateDb, err := api.eth.BlockChain().StateAt(block.Root())
>>>>>>> 7fdd714... gdbix-update v1.5.0:dbix/api.go
	if err != nil {
		return state.Dump{}, err
	}
	return stateDb.RawDump(), nil
}

<<<<<<< HEAD:tsf/api.go
// GetBlockRlp retrieves the RLP encoded for of a single block.
func (api *PublicDebugAPI) GetBlockRlp(number uint64) (string, error) {
	block := api.tsf.BlockChain().GetBlockByNumber(number)
	if block == nil {
		return "", fmt.Errorf("block #%d not found", number)
	}
	encoded, err := rlp.EncodeToBytes(block)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", encoded), nil
}

// PrintBlock retrieves a block and returns its pretty printed form.
func (api *PublicDebugAPI) PrintBlock(number uint64) (string, error) {
	block := api.tsf.BlockChain().GetBlockByNumber(number)
	if block == nil {
		return "", fmt.Errorf("block #%d not found", number)
	}
	return fmt.Sprintf("%s", block), nil
}

// SeedHash retrieves the seed hash of a block.
func (api *PublicDebugAPI) SeedHash(number uint64) (string, error) {
	block := api.tsf.BlockChain().GetBlockByNumber(number)
	if block == nil {
		return "", fmt.Errorf("block #%d not found", number)
	}
	hash, err := ethash.GetSeedHash(number)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("0x%x", hash), nil
}

// PrivateDebugAPI is the collection of Etheruem APIs exposed over the private
// debugging endpoint.
type PrivateDebugAPI struct {
	config *core.ChainConfig
	tsf    *Teslafunds
}

// NewPrivateDebugAPI creates a new API definition for the private debug methods
// of the Teslafunds service.
func NewPrivateDebugAPI(config *core.ChainConfig, tsf *Teslafunds) *PrivateDebugAPI {
	return &PrivateDebugAPI{config: config, tsf: tsf}
}

// ChaindbProperty returns leveldb properties of the chain database.
func (api *PrivateDebugAPI) ChaindbProperty(property string) (string, error) {
	ldb, ok := api.tsf.chainDb.(interface {
		LDB() *leveldb.DB
	})
	if !ok {
		return "", fmt.Errorf("chaindbProperty does not work for memory databases")
	}
	if property == "" {
		property = "leveldb.stats"
	} else if !strings.HasPrefix(property, "leveldb.") {
		property = "leveldb." + property
	}
	return ldb.LDB().GetProperty(property)
=======
// PrivateDebugAPI is the collection of Etheruem full node APIs exposed over
// the private debugging endpoint.
type PrivateDebugAPI struct {
	config *params.ChainConfig
	eth    *Ethereum
}

// NewPrivateDebugAPI creates a new API definition for the full node-related
// private debug methods of the Dubaicoin service.
func NewPrivateDebugAPI(config *params.ChainConfig, eth *Ethereum) *PrivateDebugAPI {
	return &PrivateDebugAPI{config: config, eth: eth}
>>>>>>> 7fdd714... gdbix-update v1.5.0:dbix/api.go
}

// BlockTraceResult is the returned value when replaying a block to check for
// consensus results and full VM trace logs for all included transactions.
type BlockTraceResult struct {
	Validated  bool                  `json:"validated"`
	StructLogs []ethapi.StructLogRes `json:"structLogs"`
	Error      string                `json:"error"`
}

// TraceArgs holds extra parameters to trace functions
type TraceArgs struct {
	*vm.LogConfig
	Tracer  *string
	Timeout *string
}

// TraceBlock processes the given block's RLP but does not import the block in to
// the chain.
func (api *PrivateDebugAPI) TraceBlock(blockRlp []byte, config *vm.LogConfig) BlockTraceResult {
	var block types.Block
	err := rlp.Decode(bytes.NewReader(blockRlp), &block)
	if err != nil {
		return BlockTraceResult{Error: fmt.Sprintf("could not decode block: %v", err)}
	}

	validated, logs, err := api.traceBlock(&block, config)
	return BlockTraceResult{
		Validated:  validated,
		StructLogs: ethapi.FormatLogs(logs),
		Error:      formatError(err),
	}
}

// TraceBlockFromFile loads the block's RLP from the given file name and attempts to
// process it but does not import the block in to the chain.
func (api *PrivateDebugAPI) TraceBlockFromFile(file string, config *vm.LogConfig) BlockTraceResult {
	blockRlp, err := ioutil.ReadFile(file)
	if err != nil {
		return BlockTraceResult{Error: fmt.Sprintf("could not read file: %v", err)}
	}
	return api.TraceBlock(blockRlp, config)
}

// TraceBlockByNumber processes the block by canonical block number.
func (api *PrivateDebugAPI) TraceBlockByNumber(number uint64, config *vm.LogConfig) BlockTraceResult {
	// Fetch the block that we aim to reprocess
<<<<<<< HEAD:tsf/api.go
	block := api.tsf.BlockChain().GetBlockByNumber(number)
=======
	block := api.eth.BlockChain().GetBlockByNumber(number)
>>>>>>> 7fdd714... gdbix-update v1.5.0:dbix/api.go
	if block == nil {
		return BlockTraceResult{Error: fmt.Sprintf("block #%d not found", number)}
	}

	validated, logs, err := api.traceBlock(block, config)
	return BlockTraceResult{
		Validated:  validated,
		StructLogs: ethapi.FormatLogs(logs),
		Error:      formatError(err),
	}
}

// TraceBlockByHash processes the block by hash.
func (api *PrivateDebugAPI) TraceBlockByHash(hash common.Hash, config *vm.LogConfig) BlockTraceResult {
	// Fetch the block that we aim to reprocess
<<<<<<< HEAD:tsf/api.go
	block := api.tsf.BlockChain().GetBlock(hash)
=======
	block := api.eth.BlockChain().GetBlockByHash(hash)
>>>>>>> 7fdd714... gdbix-update v1.5.0:dbix/api.go
	if block == nil {
		return BlockTraceResult{Error: fmt.Sprintf("block #%x not found", hash)}
	}

	validated, logs, err := api.traceBlock(block, config)
	return BlockTraceResult{
		Validated:  validated,
		StructLogs: ethapi.FormatLogs(logs),
		Error:      formatError(err),
	}
}

<<<<<<< HEAD:tsf/api.go

=======
>>>>>>> 7fdd714... gdbix-update v1.5.0:dbix/api.go
// traceBlock processes the given block but does not save the state.
func (api *PrivateDebugAPI) traceBlock(block *types.Block, logConfig *vm.LogConfig) (bool, []vm.StructLog, error) {
	// Validate and reprocess the block
	var (
<<<<<<< HEAD:tsf/api.go
		blockchain = api.tsf.BlockChain()
		validator  = blockchain.Validator()
		processor  = blockchain.Processor()
	)

	structLogger := vm.NewStructLogger(logConfig)

	config := vm.Config{
		Debug:  true,
		Tracer: structLogger,
	}


	if err := core.ValidateHeader(api.config, blockchain.AuxValidator(), block.Header(), blockchain.GetHeader(block.ParentHash()), true, false); err != nil {
		return false, structLogger.StructLogs(), err
=======
		blockchain = api.eth.BlockChain()
		validator  = blockchain.Validator()
		processor  = blockchain.Processor()
	)

	structLogger := vm.NewStructLogger(logConfig)

	config := vm.Config{
		Debug:  true,
		Tracer: structLogger,
>>>>>>> 7fdd714... gdbix-update v1.5.0:dbix/api.go
	}

	if err := core.ValidateHeader(api.config, blockchain.AuxValidator(), block.Header(), blockchain.GetHeader(block.ParentHash(), block.NumberU64()-1), true, false, blockchain); err != nil {
		return false, structLogger.StructLogs(), err
	} //FluxDbix
	statedb, err := blockchain.StateAt(blockchain.GetBlock(block.ParentHash(), block.NumberU64()-1).Root())
	if err != nil {
		return false, structLogger.StructLogs(), err
	}

	receipts, _, usedGas, err := processor.Process(block, statedb, config)
	if err != nil {
		return false, structLogger.StructLogs(), err
	}
<<<<<<< HEAD:tsf/api.go
	if err := validator.ValidateState(block, blockchain.GetBlock(block.ParentHash()), statedb, receipts, usedGas); err != nil {
		return false, structLogger.StructLogs(), err
	}
	return true, structLogger.StructLogs(), nil
}

// SetHead rewinds the head of the blockchain to a previous block.
func (api *PrivateDebugAPI) SetHead(number uint64) {
	api.tsf.BlockChain().SetHead(number)
}

// ExecutionResult groups all structured logs emitted by the EVM
// while replaying a transaction in debug mode as well as the amount of
// gas used and the return value
type ExecutionResult struct {
	Gas         *big.Int       `json:"gas"`
	ReturnValue string         `json:"returnValue"`
	StructLogs  []structLogRes `json:"structLogs"`
=======
	if err := validator.ValidateState(block, blockchain.GetBlock(block.ParentHash(), block.NumberU64()-1), statedb, receipts, usedGas); err != nil {
		return false, structLogger.StructLogs(), err
	}
	return true, structLogger.StructLogs(), nil
>>>>>>> 7fdd714... gdbix-update v1.5.0:dbix/api.go
}

// callmsg is the message type used for call transitions.
type callmsg struct {
	addr          common.Address
	to            *common.Address
	gas, gasPrice *big.Int
	value         *big.Int
	data          []byte
}

// accessor boilerplate to implement core.Message
func (m callmsg) From() (common.Address, error)         { return m.addr, nil }
func (m callmsg) FromFrontier() (common.Address, error) { return m.addr, nil }
func (m callmsg) Nonce() uint64                         { return 0 }
func (m callmsg) CheckNonce() bool                      { return false }
func (m callmsg) To() *common.Address                   { return m.to }
func (m callmsg) GasPrice() *big.Int                    { return m.gasPrice }
func (m callmsg) Gas() *big.Int                         { return m.gas }
func (m callmsg) Value() *big.Int                       { return m.value }
func (m callmsg) Data() []byte                          { return m.data }

// formatError formats a Go error into either an empty string or the data content
// of the error itself.
func formatError(err error) string {
	if err == nil {
		return ""
	}
	return err.Error()
}

type timeoutError struct{}

func (t *timeoutError) Error() string {
	return "Execution time exceeded"
}

// TraceTransaction returns the structured logs created during the execution of EVM
// and returns them as a JSON object.
<<<<<<< HEAD:tsf/api.go
func (api *PrivateDebugAPI) TraceTransaction(txHash common.Hash, logConfig *vm.LogConfig) (*ExecutionResult, error) {
	logger := vm.NewStructLogger(logConfig)


	// Retrieve the tx from the chain and the containing block
	tx, blockHash, _, txIndex := core.GetTransaction(api.tsf.ChainDb(), txHash)
	if tx == nil {
		return nil, fmt.Errorf("transaction %x not found", txHash)
	}
	block := api.tsf.BlockChain().GetBlock(blockHash)
=======
func (api *PrivateDebugAPI) TraceTransaction(ctx context.Context, txHash common.Hash, config *TraceArgs) (interface{}, error) {
	var tracer vm.Tracer
	if config != nil && config.Tracer != nil {
		timeout := defaultTraceTimeout
		if config.Timeout != nil {
			var err error
			if timeout, err = time.ParseDuration(*config.Timeout); err != nil {
				return nil, err
			}
		}

		var err error
		if tracer, err = ethapi.NewJavascriptTracer(*config.Tracer); err != nil {
			return nil, err
		}

		// Handle timeouts and RPC cancellations
		deadlineCtx, cancel := context.WithTimeout(ctx, timeout)
		go func() {
			<-deadlineCtx.Done()
			tracer.(*ethapi.JavascriptTracer).Stop(&timeoutError{})
		}()
		defer cancel()
	} else if config == nil {
		tracer = vm.NewStructLogger(nil)
	} else {
		tracer = vm.NewStructLogger(config.LogConfig)
	}

	// Retrieve the tx from the chain and the containing block
	tx, blockHash, _, txIndex := core.GetTransaction(api.eth.ChainDb(), txHash)
	if tx == nil {
		return nil, fmt.Errorf("transaction %x not found", txHash)
	}
	block := api.eth.BlockChain().GetBlockByHash(blockHash)
>>>>>>> 7fdd714... gdbix-update v1.5.0:dbix/api.go
	if block == nil {
		return nil, fmt.Errorf("block %x not found", blockHash)
	}
	// Create the state database to mutate and eventually trace
<<<<<<< HEAD:tsf/api.go
	parent := api.tsf.BlockChain().GetBlock(block.ParentHash())
	if parent == nil {
		return nil, fmt.Errorf("block parent %x not found", block.ParentHash())
	}
	stateDb, err := api.tsf.BlockChain().StateAt(parent.Root())
=======
	parent := api.eth.BlockChain().GetBlock(block.ParentHash(), block.NumberU64()-1)
	if parent == nil {
		return nil, fmt.Errorf("block parent %x not found", block.ParentHash())
	}
	stateDb, err := api.eth.BlockChain().StateAt(parent.Root())
>>>>>>> 7fdd714... gdbix-update v1.5.0:dbix/api.go
	if err != nil {
		return nil, err
	}

	signer := types.MakeSigner(api.config, block.Number())
	// Mutate the state and trace the selected transaction
	for idx, tx := range block.Transactions() {
		// Assemble the transaction call message
		msg, err := tx.AsMessage(signer)
		if err != nil {
			return nil, fmt.Errorf("sender retrieval failed: %v", err)
		}
		context := core.NewEVMContext(msg, block.Header(), api.eth.BlockChain())

		// Mutate the state if we haven't reached the tracing transaction yet
		if uint64(idx) < txIndex {
<<<<<<< HEAD:tsf/api.go
			vmenv := core.NewEnv(stateDb, api.config, api.tsf.BlockChain(), msg, block.Header(), vm.Config{})
=======
			vmenv := vm.NewEVM(context, stateDb, api.config, vm.Config{})
>>>>>>> 7fdd714... gdbix-update v1.5.0:dbix/api.go
			_, _, err := core.ApplyMessage(vmenv, msg, new(core.GasPool).AddGas(tx.Gas()))
			if err != nil {
				return nil, fmt.Errorf("mutation failed: %v", err)
			}
			stateDb.DeleteSuicides()
			continue
		}
<<<<<<< HEAD:tsf/api.go
		// Otherwise trace the transaction and return
		vmenv := core.NewEnv(stateDb, api.config, api.tsf.BlockChain(), msg, block.Header(), vm.Config{Debug: true, Tracer: logger})
=======

		vmenv := vm.NewEVM(context, stateDb, api.config, vm.Config{Debug: true, Tracer: tracer})
>>>>>>> 7fdd714... gdbix-update v1.5.0:dbix/api.go
		ret, gas, err := core.ApplyMessage(vmenv, msg, new(core.GasPool).AddGas(tx.Gas()))
		if err != nil {
			return nil, fmt.Errorf("tracing failed: %v", err)
		}
<<<<<<< HEAD:tsf/api.go
		return &ExecutionResult{
			Gas:         gas,
			ReturnValue: fmt.Sprintf("%x", ret),
			StructLogs:  formatLogs(logger.StructLogs()),
		}, nil
	}
	return nil, errors.New("database inconsistency")
}
=======
>>>>>>> 7fdd714... gdbix-update v1.5.0:dbix/api.go

		switch tracer := tracer.(type) {
		case *vm.StructLogger:
			return &ethapi.ExecutionResult{
				Gas:         gas,
				ReturnValue: fmt.Sprintf("%x", ret),
				StructLogs:  ethapi.FormatLogs(tracer.StructLogs()),
			}, nil
		case *ethapi.JavascriptTracer:
			return tracer.GetResult()
		}
	}
<<<<<<< HEAD:tsf/api.go
	if msg.gas.Cmp(common.Big0) == 0 {
		msg.gas = big.NewInt(50000000)
	}
	if msg.gasPrice.Cmp(common.Big0) == 0 {
		msg.gasPrice = new(big.Int).Mul(big.NewInt(50), common.Shannon)
	}

	// Execute the call and return
	vmenv := core.NewEnv(stateDb, s.config, s.bc, msg, block.Header(), vm.Config{
		Debug: true,
	})
	gp := new(core.GasPool).AddGas(common.MaxBig)

	ret, gas, err := core.ApplyMessage(vmenv, msg, gp)
	return &ExecutionResult{
		Gas:         gas,
		ReturnValue: fmt.Sprintf("%x", ret),
		StructLogs:  formatLogs(logger.StructLogs()),
	}, nil
}

// PublicNetAPI offers network related RPC methods
type PublicNetAPI struct {
	net            *p2p.Server
	networkVersion int
}

// NewPublicNetAPI creates a new net API instance.
func NewPublicNetAPI(net *p2p.Server, networkVersion int) *PublicNetAPI {
	return &PublicNetAPI{net, networkVersion}
}

// Listening returns an indication if the node is listening for network connections.
func (s *PublicNetAPI) Listening() bool {
	return true // always listening
}

// PeerCount returns the number of connected peers
func (s *PublicNetAPI) PeerCount() *rpc.HexNumber {
	return rpc.NewHexNumber(s.net.PeerCount())
=======
	return nil, errors.New("database inconsistency")
>>>>>>> 7fdd714... gdbix-update v1.5.0:dbix/api.go
}

// Preimage is a debug API function that returns the preimage for a sha3 hash, if known.
func (api *PrivateDebugAPI) Preimage(ctx context.Context, hash common.Hash) (hexutil.Bytes, error) {
	db := core.PreimageTable(api.eth.ChainDb())
	return db.Get(hash.Bytes())
}
