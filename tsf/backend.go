<<<<<<< HEAD:tsf/backend.go
// Copyright 2014 The go-ethereum Authors && Copyright 2015 go-teslafunds Authors
// This file is part of the go-teslafunds library.
//
// The go-teslafunds library is free software: you can redistribute it and/or modify
=======
// Copyright 2014 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
>>>>>>> 7fdd714... gdbix-update v1.5.0:dbix/backend.go
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
<<<<<<< HEAD:tsf/backend.go
// The go-teslafunds library is distributed in the hope that it will be useful,
=======
// The go-ethereum library is distributed in the hope that it will be useful,
>>>>>>> 7fdd714... gdbix-update v1.5.0:dbix/backend.go
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
<<<<<<< HEAD:tsf/backend.go
// along with the go-teslafunds library. If not, see <http://www.gnu.org/licenses/>.

// Package tsf implements the Teslafunds protocol.
package tsf
=======
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

// Package eth implements the Dubaicoin protocol.
package eth
>>>>>>> 7fdd714... gdbix-update v1.5.0:dbix/backend.go

import (
	"errors"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
	"time"

<<<<<<< HEAD:tsf/backend.go
	"github.com/teslafunds/ethash"
	"github.com/teslafunds/go-teslafunds/accounts"
	"github.com/teslafunds/go-teslafunds/common"
	"github.com/teslafunds/go-teslafunds/common/compiler"
	"github.com/teslafunds/go-teslafunds/common/httpclient"
	"github.com/teslafunds/go-teslafunds/common/registrar/ethreg"
	"github.com/teslafunds/go-teslafunds/core"
	"github.com/teslafunds/go-teslafunds/core/types"
	"github.com/teslafunds/go-teslafunds/core/vm"
	"github.com/teslafunds/go-teslafunds/tsf/downloader"
	"github.com/teslafunds/go-teslafunds/tsf/filters"
	"github.com/teslafunds/go-teslafunds/ethdb"
	"github.com/teslafunds/go-teslafunds/event"
	"github.com/teslafunds/go-teslafunds/logger"
	"github.com/teslafunds/go-teslafunds/logger/glog"
	"github.com/teslafunds/go-teslafunds/miner"
	"github.com/teslafunds/go-teslafunds/node"
	"github.com/teslafunds/go-teslafunds/p2p"
	"github.com/teslafunds/go-teslafunds/rlp"
	"github.com/teslafunds/go-teslafunds/rpc"
=======
	"github.com/ethereum/ethash"
	"github.com/dubaicoin-dbix/go-dubaicoin/accounts"
	"github.com/dubaicoin-dbix/go-dubaicoin/common"
	"github.com/dubaicoin-dbix/go-dubaicoin/core"
	"github.com/dubaicoin-dbix/go-dubaicoin/core/types"
	"github.com/dubaicoin-dbix/go-dubaicoin/core/vm"
	"github.com/dubaicoin-dbix/go-dubaicoin/dbix/downloader"
	"github.com/dubaicoin-dbix/go-dubaicoin/dbix/filters"
	"github.com/dubaicoin-dbix/go-dubaicoin/dbix/gasprice"
	"github.com/dubaicoin-dbix/go-dubaicoin/dbixdb"
	"github.com/dubaicoin-dbix/go-dubaicoin/event"
	"github.com/dubaicoin-dbix/go-dubaicoin/internal/ethapi"
	"github.com/dubaicoin-dbix/go-dubaicoin/logger"
	"github.com/dubaicoin-dbix/go-dubaicoin/logger/glog"
	"github.com/dubaicoin-dbix/go-dubaicoin/miner"
	"github.com/dubaicoin-dbix/go-dubaicoin/node"
	"github.com/dubaicoin-dbix/go-dubaicoin/p2p"
	"github.com/dubaicoin-dbix/go-dubaicoin/params"
	"github.com/dubaicoin-dbix/go-dubaicoin/pow"
	"github.com/dubaicoin-dbix/go-dubaicoin/rpc"
>>>>>>> 7fdd714... gdbix-update v1.5.0:dbix/backend.go
)

const (
	epochLength    = 30000
	ethashRevision = 23

	autoDAGcheckInterval = 10 * time.Hour
	autoDAGepochHeight   = epochLength / 2
)

var (
	datadirInUseErrnos = map[uint]bool{11: true, 32: true, 35: true}
	portInUseErrRE     = regexp.MustCompile("address already in use")
)

type Config struct {
	ChainConfig *params.ChainConfig // chain configuration

	NetworkId  int    // Network ID to use for selecting peers to connect to
	Genesis    string // Genesis JSON to seed the chain database with
	FastSync   bool   // Enables the state download based fast synchronisation algorithm
	LightMode  bool   // Running in light client mode
	LightServ  int    // Maximum percentage of time allowed for serving LES requests
	LightPeers int    // Maximum number of LES client peers
	MaxPeers   int    // Maximum number of global peers

	SkipBcVersionCheck bool // e.g. blockchain export
	DatabaseCache      int
	DatabaseHandles    int

	DocRoot   string
	AutoDAG   bool
	PowFake   bool
	PowTest   bool
	PowShared bool
	ExtraData []byte

	Etherbase    common.Address
	GasPrice     *big.Int
	MinerThreads int
	SolcPath     string

	GpoMinGasPrice          *big.Int
	GpoMaxGasPrice          *big.Int
	GpoFullBlockRatio       int
	GpobaseStepDown         int
	GpobaseStepUp           int
	GpobaseCorrectionFactor int

	EnablePreimageRecording bool

	TestGenesisBlock *types.Block   // Genesis block to seed the chain database with (testing only!)
	TestGenesisState ethdb.Database // Genesis state to seed the database with (testing only!)
}

<<<<<<< HEAD:tsf/backend.go
type Teslafunds struct {
	chainConfig *core.ChainConfig
	// Channel for shutting down the teslafunds
	shutdownChan chan bool

	// DB interfaces
	chainDb ethdb.Database // Block chain database
	dappDb  ethdb.Database // Dapp database
=======
type LesServer interface {
	Start(srvr *p2p.Server)
	Stop()
	Protocols() []p2p.Protocol
}
>>>>>>> 7fdd714... gdbix-update v1.5.0:dbix/backend.go

// Ethereum implements the Dubaicoin full node service.
type Ethereum struct {
	chainConfig *params.ChainConfig
	// Channel for shutting down the service
	shutdownChan  chan bool // Channel for shutting down the dubaicoin
	stopDbUpgrade func()    // stop chain db sequential key upgrade
	// Handlers
	txPool          *core.TxPool
	txMu            sync.Mutex
	blockchain      *core.BlockChain
	protocolManager *ProtocolManager
	lesServer       LesServer
	// DB interfaces
	chainDb ethdb.Database // Block chain database

	eventMux       *event.TypeMux
	pow            pow.PoW
	accountManager *accounts.Manager

	ApiBackend *EthApiBackend

	miner        *miner.Miner
	Mining       bool
	MinerThreads int
	AutoDAG      bool
	autodagquit  chan bool
	etherbase    common.Address
	solcPath     string

	netVersionId  int
	netRPCService *ethapi.PublicNetAPI
}

func (s *Ethereum) AddLesServer(ls LesServer) {
	s.lesServer = ls
	s.protocolManager.lesServer = ls
}

<<<<<<< HEAD:tsf/backend.go
func New(ctx *node.ServiceContext, config *Config) (*Teslafunds, error) {
	// Open the chain database and perform any upgrades needed
	chainDb, err := ctx.OpenDatabase("chaindata", config.DatabaseCache, config.DatabaseHandles)
	if err != nil {
		return nil, err
	}
	if db, ok := chainDb.(*ethdb.LDBDatabase); ok {
		db.Meter("tsf/db/chaindata/")
	}
	if err := upgradeChainDatabase(chainDb); err != nil {
=======
// New creates a new Ethereum object (including the
// initialisation of the common Ethereum object)
func New(ctx *node.ServiceContext, config *Config) (*Ethereum, error) {
	chainDb, err := CreateDB(ctx, config, "chaindata")
	if err != nil {
		return nil, err
	}
	stopDbUpgrade := upgradeSequentialKeys(chainDb)
	if err := SetupGenesisBlock(&chainDb, config); err != nil {
>>>>>>> 7fdd714... gdbix-update v1.5.0:dbix/backend.go
		return nil, err
	}
	pow, err := CreatePoW(config)
	if err != nil {
		return nil, err
	}
<<<<<<< HEAD:tsf/backend.go
	if db, ok := dappDb.(*ethdb.LDBDatabase); ok {
		db.Meter("tsf/db/dapp/")
	}
	glog.V(logger.Info).Infof("Protocol Versions: %v, Network Id: %v", ProtocolVersions, config.NetworkId)
=======
>>>>>>> 7fdd714... gdbix-update v1.5.0:dbix/backend.go

	eth := &Ethereum{
		chainDb:        chainDb,
		eventMux:       ctx.EventMux,
		accountManager: ctx.AccountManager,
		pow:            pow,
		shutdownChan:   make(chan bool),
		stopDbUpgrade:  stopDbUpgrade,
		netVersionId:   config.NetworkId,
		etherbase:      config.Etherbase,
		MinerThreads:   config.MinerThreads,
		AutoDAG:        config.AutoDAG,
		solcPath:       config.SolcPath,
	}

	if err := upgradeChainDatabase(chainDb); err != nil {
		return nil, err
	}
	if err := addMipmapBloomBins(chainDb); err != nil {
		return nil, err
	}

	glog.V(logger.Info).Infof("Protocol Versions: %v, Network Id: %v", ProtocolVersions, config.NetworkId)

	if !config.SkipBcVersionCheck {
		bcVersion := core.GetBlockChainVersion(chainDb)
<<<<<<< HEAD:tsf/backend.go
		if bcVersion != config.BlockChainVersion && bcVersion != 0 {
			return nil, fmt.Errorf("Blockchain DB version mismatch (%d / %d). Run gtsf upgradedb.\n", bcVersion, config.BlockChainVersion)
		}
		core.WriteBlockChainVersion(chainDb, config.BlockChainVersion)
	}
	glog.V(logger.Info).Infof("Blockchain DB Version: %d", config.BlockChainVersion)

	tsf := &Teslafunds{
		shutdownChan:            make(chan bool),
		chainDb:                 chainDb,
		dappDb:                  dappDb,
		eventMux:                ctx.EventMux,
		accountManager:          config.AccountManager,
		etherbase:               config.Etherbase,
		netVersionId:            config.NetworkId,
		NatSpec:                 config.NatSpec,
		MinerThreads:            config.MinerThreads,
		SolcPath:                config.SolcPath,
		AutoDAG:                 config.AutoDAG,
		PowTest:                 config.PowTest,
		GpoMinGasPrice:          config.GpoMinGasPrice,
		GpoMaxGasPrice:          config.GpoMaxGasPrice,
		GpoFullBlockRatio:       config.GpoFullBlockRatio,
		GpobaseStepDown:         config.GpobaseStepDown,
		GpobaseStepUp:           config.GpobaseStepUp,
		GpobaseCorrectionFactor: config.GpobaseCorrectionFactor,
		httpclient:              httpclient.New(config.DocRoot),
	}
	switch {
	case config.PowTest:
		glog.V(logger.Info).Infof("ethash used in test mode")
		tsf.pow, err = ethash.NewForTesting()
		if err != nil {
			return nil, err
		}
	case config.PowShared:
		glog.V(logger.Info).Infof("ethash used in shared mode")
		tsf.pow = ethash.NewShared()

	default:
		tsf.pow = ethash.New()
=======
		if bcVersion != core.BlockChainVersion && bcVersion != 0 {
			return nil, fmt.Errorf("Blockchain DB version mismatch (%d / %d). Run gdbix upgradedb.\n", bcVersion, core.BlockChainVersion)
		}
		core.WriteBlockChainVersion(chainDb, core.BlockChainVersion)
>>>>>>> 7fdd714... gdbix-update v1.5.0:dbix/backend.go
	}

	// load the genesis block or write a new one if no genesis
	// block is prenent in the database.
	genesis := core.GetBlock(chainDb, core.GetCanonicalHash(chainDb, 0), 0)
	if genesis == nil {
		genesis, err = core.WriteDefaultGenesisBlock(chainDb)
		if err != nil {
			return nil, err
		}
		glog.V(logger.Info).Infoln("WARNING: Wrote default teslafunds genesis block")
	}

	if config.ChainConfig == nil {
		return nil, errors.New("missing chain config")
	}
	core.WriteChainConfig(chainDb, genesis.Hash(), config.ChainConfig)

<<<<<<< HEAD:tsf/backend.go
	tsf.chainConfig = config.ChainConfig
	tsf.chainConfig.VmConfig = vm.Config{
		EnableJit: config.EnableJit,
		ForceJit:  config.ForceJit,
	}

	tsf.blockchain, err = core.NewBlockChain(chainDb, tsf.chainConfig, tsf.pow, tsf.EventMux())
=======
	eth.chainConfig = config.ChainConfig

	glog.V(logger.Info).Infoln("Chain config:", eth.chainConfig)

	eth.blockchain, err = core.NewBlockChain(chainDb, eth.chainConfig, eth.pow, eth.EventMux(), vm.Config{EnablePreimageRecording: config.EnablePreimageRecording})
>>>>>>> 7fdd714... gdbix-update v1.5.0:dbix/backend.go
	if err != nil {
		if err == core.ErrNoGenesis {
			return nil, fmt.Errorf(`No chain found. Please initialise a new chain using the "init" subcommand.`)
		}
		return nil, err
	}
<<<<<<< HEAD:tsf/backend.go
	tsf.gpo = NewGasPriceOracle(tsf)

	newPool := core.NewTxPool(tsf.chainConfig, tsf.EventMux(), tsf.blockchain.State, tsf.blockchain.GasLimit)
	tsf.txPool = newPool

	if tsf.protocolManager, err = NewProtocolManager(tsf.chainConfig, config.FastSync, config.NetworkId, tsf.eventMux, tsf.txPool, tsf.pow, tsf.blockchain, chainDb); err != nil {
		return nil, err
	}
	tsf.miner = miner.New(tsf, tsf.chainConfig, tsf.EventMux(), tsf.pow)
	tsf.miner.SetGasPrice(config.GasPrice)
	tsf.miner.SetExtra(config.ExtraData)

	return tsf, nil
=======
	newPool := core.NewTxPool(eth.chainConfig, eth.EventMux(), eth.blockchain.State, eth.blockchain.GasLimit)
	eth.txPool = newPool

	maxPeers := config.MaxPeers
	if config.LightServ > 0 {
		// if we are running a light server, limit the number of ETH peers so that we reserve some space for incoming LES connections
		// temporary solution until the new peer connectivity API is finished
		halfPeers := maxPeers / 2
		maxPeers -= config.LightPeers
		if maxPeers < halfPeers {
			maxPeers = halfPeers
		}
	}

	if eth.protocolManager, err = NewProtocolManager(eth.chainConfig, config.FastSync, config.NetworkId, maxPeers, eth.eventMux, eth.txPool, eth.pow, eth.blockchain, chainDb); err != nil {
		return nil, err
	}
	eth.miner = miner.New(eth, eth.chainConfig, eth.EventMux(), eth.pow)
	eth.miner.SetGasPrice(config.GasPrice)
	eth.miner.SetExtra(config.ExtraData)

	gpoParams := &gasprice.GpoParams{
		GpoMinGasPrice:          config.GpoMinGasPrice,
		GpoMaxGasPrice:          config.GpoMaxGasPrice,
		GpoFullBlockRatio:       config.GpoFullBlockRatio,
		GpobaseStepDown:         config.GpobaseStepDown,
		GpobaseStepUp:           config.GpobaseStepUp,
		GpobaseCorrectionFactor: config.GpobaseCorrectionFactor,
	}
	gpo := gasprice.NewGasPriceOracle(eth.blockchain, chainDb, eth.eventMux, gpoParams)
	eth.ApiBackend = &EthApiBackend{eth, gpo}

	return eth, nil
}

// CreateDB creates the chain database.
func CreateDB(ctx *node.ServiceContext, config *Config, name string) (ethdb.Database, error) {
	db, err := ctx.OpenDatabase(name, config.DatabaseCache, config.DatabaseHandles)
	if db, ok := db.(*ethdb.LDBDatabase); ok {
		db.Meter("eth/db/chaindata/")
	}
	return db, err
}

// SetupGenesisBlock initializes the genesis block for a Dubaicoin service
func SetupGenesisBlock(chainDb *ethdb.Database, config *Config) error {
	// Load up any custom genesis block if requested
	if len(config.Genesis) > 0 {
		block, err := core.WriteGenesisBlock(*chainDb, strings.NewReader(config.Genesis))
		if err != nil {
			return err
		}
		glog.V(logger.Info).Infof("Successfully wrote custom genesis block: %x", block.Hash())
	}
	// Load up a test setup if directly injected
	if config.TestGenesisState != nil {
		*chainDb = config.TestGenesisState
	}
	if config.TestGenesisBlock != nil {
		core.WriteTd(*chainDb, config.TestGenesisBlock.Hash(), config.TestGenesisBlock.NumberU64(), config.TestGenesisBlock.Difficulty())
		core.WriteBlock(*chainDb, config.TestGenesisBlock)
		core.WriteCanonicalHash(*chainDb, config.TestGenesisBlock.Hash(), config.TestGenesisBlock.NumberU64())
		core.WriteHeadBlockHash(*chainDb, config.TestGenesisBlock.Hash())
	}
	return nil
}

// CreatePoW creates the required type of PoW instance for a Dubaicoin service
func CreatePoW(config *Config) (pow.PoW, error) {
	switch {
	case config.PowFake:
		glog.V(logger.Info).Infof("ethash used in fake mode")
		return pow.PoW(core.FakePow{}), nil
	case config.PowTest:
		glog.V(logger.Info).Infof("ethash used in test mode")
		return ethash.NewForTesting()
	case config.PowShared:
		glog.V(logger.Info).Infof("ethash used in shared mode")
		return ethash.NewShared(), nil
	default:
		return ethash.New(), nil
	}
>>>>>>> 7fdd714... gdbix-update v1.5.0:dbix/backend.go
}

// APIs returns the collection of RPC services the teslafunds package offers.
// NOTE, some of these services probably need to be moved to somewhere else.
<<<<<<< HEAD:tsf/backend.go
func (s *Teslafunds) APIs() []rpc.API {
	return []rpc.API{
		{
			Namespace: "tsf",
=======
func (s *Ethereum) APIs() []rpc.API {
	return append(ethapi.GetAPIs(s.ApiBackend, s.solcPath), []rpc.API{
		{
			Namespace: "eth",
>>>>>>> 7fdd714... gdbix-update v1.5.0:dbix/backend.go
			Version:   "1.0",
			Service:   NewPublicEthereumAPI(s),
			Public:    true,
		}, {
<<<<<<< HEAD:tsf/backend.go
			Namespace: "tsf",
=======
			Namespace: "eth",
			Version:   "1.0",
			Service:   NewPublicMinerAPI(s),
			Public:    true,
		}, {
			Namespace: "eth",
>>>>>>> 7fdd714... gdbix-update v1.5.0:dbix/backend.go
			Version:   "1.0",
			Service:   downloader.NewPublicDownloaderAPI(s.protocolManager.downloader, s.eventMux),
			Public:    true,
		}, {
			Namespace: "miner",
			Version:   "1.0",
			Service:   NewPrivateMinerAPI(s),
			Public:    false,
		}, {
<<<<<<< HEAD:tsf/backend.go
			Namespace: "tsf",
=======
			Namespace: "eth",
>>>>>>> 7fdd714... gdbix-update v1.5.0:dbix/backend.go
			Version:   "1.0",
			Service:   filters.NewPublicFilterAPI(s.ApiBackend, false),
			Public:    true,
		}, {
			Namespace: "tsf",
			Version:   "1.0",
			Service:   NewPublicEthereumAPI(s),
			Public:    true,
		}, {
			Namespace: "tsf",
			Version:   "1.0",
			Service:   NewPublicMinerAPI(s),
			Public:    true,
		}, {
			Namespace: "tsf",
			Version:   "1.0",
			Service:   downloader.NewPublicDownloaderAPI(s.protocolManager.downloader, s.eventMux),
			Public:    true,
		}, {
<<<<<<< HEAD:tsf/backend.go
			Namespace: "miner",
			Version:   "1.0",
			Service:   NewPrivateMinerAPI(s),
			Public:    false,
		}, {
			Namespace: "txpool",
			Version:   "1.0",
			Service:   NewPublicTxPoolAPI(s),
			Public:    true,
		}, {
			Namespace: "tsf",
=======
			Namespace: "dbix",
>>>>>>> 7fdd714... gdbix-update v1.5.0:dbix/backend.go
			Version:   "1.0",
			Service:   filters.NewPublicFilterAPI(s.ApiBackend, false),
			Public:    true,
		}, {
			Namespace: "admin",
			Version:   "1.0",
			Service:   NewPrivateAdminAPI(s),
		}, {
			Namespace: "debug",
			Version:   "1.0",
			Service:   NewPublicDebugAPI(s),
			Public:    true,
		}, {
			Namespace: "debug",
			Version:   "1.0",
			Service:   NewPrivateDebugAPI(s.chainConfig, s),
		}, {
			Namespace: "net",
			Version:   "1.0",
			Service:   s.netRPCService,
			Public:    true,
		},
	}...)
}

<<<<<<< HEAD:tsf/backend.go
func (s *Teslafunds) ResetWithGenesisBlock(gb *types.Block) {
	s.blockchain.ResetWithGenesisBlock(gb)
}

func (s *Teslafunds) Etherbase() (eb common.Address, err error) {
=======
func (s *Ethereum) ResetWithGenesisBlock(gb *types.Block) {
	s.blockchain.ResetWithGenesisBlock(gb)
}

func (s *Ethereum) Etherbase() (eb common.Address, err error) {
>>>>>>> 7fdd714... gdbix-update v1.5.0:dbix/backend.go
	eb = s.etherbase
	if (eb == common.Address{}) {
		firstAccount, err := s.AccountManager().AccountByIndex(0)
		eb = firstAccount.Address
		if err != nil {
			return eb, fmt.Errorf("etherbase address must be explicitly specified")
		}
	}
	return eb, nil
}

// set in js console via admin interface or wrapper from cli flags
<<<<<<< HEAD:tsf/backend.go
func (self *Teslafunds) SetEtherbase(etherbase common.Address) {
=======
func (self *Ethereum) SetEtherbase(etherbase common.Address) {
>>>>>>> 7fdd714... gdbix-update v1.5.0:dbix/backend.go
	self.etherbase = etherbase
	self.miner.SetEtherbase(etherbase)
}

<<<<<<< HEAD:tsf/backend.go
func (s *Teslafunds) StopMining()         { s.miner.Stop() }
func (s *Teslafunds) IsMining() bool      { return s.miner.Mining() }
func (s *Teslafunds) Miner() *miner.Miner { return s.miner }

func (s *Teslafunds) AccountManager() *accounts.Manager  { return s.accountManager }
func (s *Teslafunds) BlockChain() *core.BlockChain       { return s.blockchain }
func (s *Teslafunds) TxPool() *core.TxPool               { return s.txPool }
func (s *Teslafunds) EventMux() *event.TypeMux           { return s.eventMux }
func (s *Teslafunds) ChainDb() ethdb.Database            { return s.chainDb }
func (s *Teslafunds) DappDb() ethdb.Database             { return s.dappDb }
func (s *Teslafunds) IsListening() bool                  { return true } // Always listening
func (s *Teslafunds) EthVersion() int                    { return int(s.protocolManager.SubProtocols[0].Version) }
func (s *Teslafunds) NetVersion() int                    { return s.netVersionId }
func (s *Teslafunds) Downloader() *downloader.Downloader { return s.protocolManager.downloader }

// Protocols implements node.Service, returning all the currently configured
// network protocols to start.
func (s *Teslafunds) Protocols() []p2p.Protocol {
	return s.protocolManager.SubProtocols
}

// Start implements node.Service, starting all internal goroutines needed by the
// Teslafunds protocol implementation.
func (s *Teslafunds) Start(srvr *p2p.Server) error {
=======
func (s *Ethereum) StartMining(threads int) error {
	eb, err := s.Etherbase()
	if err != nil {
		err = fmt.Errorf("Cannot start mining without etherbase address: %v", err)
		glog.V(logger.Error).Infoln(err)
		return err
	}
	go s.miner.Start(eb, threads)
	return nil
}

func (s *Ethereum) StopMining()         { s.miner.Stop() }
func (s *Ethereum) IsMining() bool      { return s.miner.Mining() }
func (s *Ethereum) Miner() *miner.Miner { return s.miner }

func (s *Ethereum) AccountManager() *accounts.Manager  { return s.accountManager }
func (s *Ethereum) BlockChain() *core.BlockChain       { return s.blockchain }
func (s *Ethereum) TxPool() *core.TxPool               { return s.txPool }
func (s *Ethereum) EventMux() *event.TypeMux           { return s.eventMux }
func (s *Ethereum) Pow() pow.PoW                       { return s.pow }
func (s *Ethereum) ChainDb() ethdb.Database            { return s.chainDb }
func (s *Ethereum) IsListening() bool                  { return true } // Always listening
func (s *Ethereum) EthVersion() int                    { return int(s.protocolManager.SubProtocols[0].Version) }
func (s *Ethereum) NetVersion() int                    { return s.netVersionId }
func (s *Ethereum) Downloader() *downloader.Downloader { return s.protocolManager.downloader }

// Protocols implements node.Service, returning all the currently configured
// network protocols to start.
func (s *Ethereum) Protocols() []p2p.Protocol {
	if s.lesServer == nil {
		return s.protocolManager.SubProtocols
	} else {
		return append(s.protocolManager.SubProtocols, s.lesServer.Protocols()...)
	}
}

// Start implements node.Service, starting all internal goroutines needed by the
// Dubaicoin protocol implementation.
func (s *Ethereum) Start(srvr *p2p.Server) error {
	s.netRPCService = ethapi.NewPublicNetAPI(srvr, s.NetVersion())
>>>>>>> 7fdd714... gdbix-update v1.5.0:dbix/backend.go
	if s.AutoDAG {
		s.StartAutoDAG()
	}
	s.protocolManager.Start()
	if s.lesServer != nil {
		s.lesServer.Start(srvr)
	}
	return nil
}

// Stop implements node.Service, terminating all internal goroutines used by the
<<<<<<< HEAD:tsf/backend.go
// Teslafunds protocol.
func (s *Teslafunds) Stop() error {
=======
// Dubaicoin protocol.
func (s *Ethereum) Stop() error {
	if s.stopDbUpgrade != nil {
		s.stopDbUpgrade()
	}
>>>>>>> 7fdd714... gdbix-update v1.5.0:dbix/backend.go
	s.blockchain.Stop()
	s.protocolManager.Stop()
	if s.lesServer != nil {
		s.lesServer.Stop()
	}
	s.txPool.Stop()
	s.miner.Stop()
	s.eventMux.Stop()

	s.StopAutoDAG()

	s.chainDb.Close()
	close(s.shutdownChan)

	return nil
}

// This function will wait for a shutdown and resumes main thread execution
<<<<<<< HEAD:tsf/backend.go
func (s *Teslafunds) WaitForShutdown() {
=======
func (s *Ethereum) WaitForShutdown() {
>>>>>>> 7fdd714... gdbix-update v1.5.0:dbix/backend.go
	<-s.shutdownChan
}

// StartAutoDAG() spawns a go routine that checks the DAG every autoDAGcheckInterval
// by default that is 10 times per epoch
// in epoch n, if we past autoDAGepochHeight within-epoch blocks,
// it calls ethash.MakeDAG  to pregenerate the DAG for the next epoch n+1
// if it does not exist yet as well as remove the DAG for epoch n-1
// the loop quits if autodagquit channel is closed, it can safely restart and
// stop any number of times.
// For any more sophisticated pattern of DAG generation, use CLI subcommand
// makedag
<<<<<<< HEAD:tsf/backend.go
func (self *Teslafunds) StartAutoDAG() {
=======
func (self *Ethereum) StartAutoDAG() {
>>>>>>> 7fdd714... gdbix-update v1.5.0:dbix/backend.go
	if self.autodagquit != nil {
		return // already started
	}
	go func() {
		glog.V(logger.Info).Infof("Automatic pregeneration of ethash DAG ON (ethash dir: %s)", ethash.DefaultDir)
		var nextEpoch uint64
		timer := time.After(0)
		self.autodagquit = make(chan bool)
		for {
			select {
			case <-timer:
				glog.V(logger.Info).Infof("checking DAG (ethash dir: %s)", ethash.DefaultDir)
				currentBlock := self.BlockChain().CurrentBlock().NumberU64()
				thisEpoch := currentBlock / epochLength
				if nextEpoch <= thisEpoch {
					if currentBlock%epochLength > autoDAGepochHeight {
						if thisEpoch > 0 {
							previousDag, previousDagFull := dagFiles(thisEpoch - 1)
							os.Remove(filepath.Join(ethash.DefaultDir, previousDag))
							os.Remove(filepath.Join(ethash.DefaultDir, previousDagFull))
							glog.V(logger.Info).Infof("removed DAG for epoch %d (%s)", thisEpoch-1, previousDag)
						}
						nextEpoch = thisEpoch + 1
						dag, _ := dagFiles(nextEpoch)
						if _, err := os.Stat(dag); os.IsNotExist(err) {
							glog.V(logger.Info).Infof("Pregenerating DAG for epoch %d (%s)", nextEpoch, dag)
							err := ethash.MakeDAG(nextEpoch*epochLength, "") // "" -> ethash.DefaultDir
							if err != nil {
								glog.V(logger.Error).Infof("Error generating DAG for epoch %d (%s)", nextEpoch, dag)
								return
							}
						} else {
							glog.V(logger.Error).Infof("DAG for epoch %d (%s)", nextEpoch, dag)
						}
					}
				}
				timer = time.After(autoDAGcheckInterval)
			case <-self.autodagquit:
				return
			}
		}
	}()
}

// stopAutoDAG stops automatic DAG pregeneration by quitting the loop
<<<<<<< HEAD:tsf/backend.go
func (self *Teslafunds) StopAutoDAG() {
=======
func (self *Ethereum) StopAutoDAG() {
>>>>>>> 7fdd714... gdbix-update v1.5.0:dbix/backend.go
	if self.autodagquit != nil {
		close(self.autodagquit)
		self.autodagquit = nil
	}
	glog.V(logger.Info).Infof("Automatic pregeneration of ethash DAG OFF (ethash dir: %s)", ethash.DefaultDir)
}

<<<<<<< HEAD:tsf/backend.go

// HTTPClient returns the light http client used for fetching offchain docs
// (natspec, source for verification)
func (self *Teslafunds) HTTPClient() *httpclient.HTTPClient {
	return self.httpclient
}

func (self *Teslafunds) Solc() (*compiler.Solidity, error) {
	var err error
	if self.solc == nil {
		self.solc, err = compiler.New(self.SolcPath)
	}
	return self.solc, err
}

// set in js console via admin interface or wrapper from cli flags
func (self *Teslafunds) SetSolc(solcPath string) (*compiler.Solidity, error) {
	self.SolcPath = solcPath
	self.solc = nil
	return self.Solc()
}

=======
>>>>>>> 7fdd714... gdbix-update v1.5.0:dbix/backend.go
// dagFiles(epoch) returns the two alternative DAG filenames (not a path)
// 1) <revision>-<hex(seedhash[8])> 2) full-R<revision>-<hex(seedhash[8])>
func dagFiles(epoch uint64) (string, string) {
	seedHash, _ := ethash.GetSeedHash(epoch * epochLength)
	dag := fmt.Sprintf("full-R%d-%x", ethashRevision, seedHash[:8])
	return dag, "full-R" + dag
}
