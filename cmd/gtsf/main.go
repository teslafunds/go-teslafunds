<<<<<<< HEAD:cmd/gtsf/main.go
﻿// Copyright 2014 The go-ethereum Authors && Copyright 2015 go-teslafunds Authors
// This file is part of go-teslafunds.
//
// go-teslafunds is free software: you can redistribute it and/or modify
=======
// Copyright 2014 The go-ethereum Authors
// This file is part of go-ethereum.
//
// go-ethereum is free software: you can redistribute it and/or modify
>>>>>>> 7fdd714... gdbix-update v1.5.0:cmd/gdbix/main.go
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
<<<<<<< HEAD:cmd/gtsf/main.go
// go-teslafunds is distributed in the hope that it will be useful,
=======
// go-ethereum is distributed in the hope that it will be useful,
>>>>>>> 7fdd714... gdbix-update v1.5.0:cmd/gdbix/main.go
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
<<<<<<< HEAD:cmd/gtsf/main.go
// along with go-teslafunds. If not, see <http://www.gnu.org/licenses/>.
=======
// along with go-ethereum. If not, see <http://www.gnu.org/licenses/>.
>>>>>>> 7fdd714... gdbix-update v1.5.0:cmd/gdbix/main.go

// gtsf is the official command-line client for Teslafunds.
package main

import (
	"encoding/hex"
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"

<<<<<<< HEAD:cmd/gtsf/main.go
	"github.com/teslafunds/ethash"
	"github.com/teslafunds/go-teslafunds/cmd/utils"
	"github.com/teslafunds/go-teslafunds/common"
	"github.com/teslafunds/go-teslafunds/console"
	"github.com/teslafunds/go-teslafunds/core"
	"github.com/teslafunds/go-teslafunds/tsf"
	"github.com/teslafunds/go-teslafunds/ethdb"
	"github.com/teslafunds/go-teslafunds/internal/debug"
	"github.com/teslafunds/go-teslafunds/logger"
	"github.com/teslafunds/go-teslafunds/logger/glog"
	"github.com/teslafunds/go-teslafunds/metrics"
	"github.com/teslafunds/go-teslafunds/node"
	"github.com/teslafunds/go-teslafunds/params"
	"github.com/teslafunds/go-teslafunds/release"
	"github.com/teslafunds/go-teslafunds/rlp"
=======
	"github.com/dubaicoin-dbix/go-dubaicoin/cmd/utils"
	"github.com/dubaicoin-dbix/go-dubaicoin/common"
	"github.com/dubaicoin-dbix/go-dubaicoin/console"
	"github.com/dubaicoin-dbix/go-dubaicoin/contracts/release"
	"github.com/dubaicoin-dbix/go-dubaicoin/dbix"
	"github.com/dubaicoin-dbix/go-dubaicoin/internal/debug"
	"github.com/dubaicoin-dbix/go-dubaicoin/logger"
	"github.com/dubaicoin-dbix/go-dubaicoin/logger/glog"
	"github.com/dubaicoin-dbix/go-dubaicoin/metrics"
	"github.com/dubaicoin-dbix/go-dubaicoin/node"
	"github.com/dubaicoin-dbix/go-dubaicoin/params"
	"github.com/dubaicoin-dbix/go-dubaicoin/rlp"
>>>>>>> 7fdd714... gdbix-update v1.5.0:cmd/gdbix/main.go
	"gopkg.in/urfave/cli.v1"
)

const (
<<<<<<< HEAD:cmd/gtsf/main.go
	clientIdentifier = "Gtsf"   // Client identifier to advertise over the network
	versionMajor     = 1        // Major version component of the current release
	versionMinor     = 0       // Minor version component of the current release
	versionPatch     = 1       // Patch version component of the current release
	versionMeta      = "initial" // Version metadata to append to the version string
	versionOracle = "0x926d69cc3bbf81d52cba6886d788df007a15a3cd" // Teslafunds address of the Gtsf release oracle
=======
	clientIdentifier = "gdbix" // Client identifier to advertise over the network
>>>>>>> 7fdd714... gdbix-update v1.5.0:cmd/gdbix/main.go
)

var (
	// Git SHA1 commit hash of the release (set via linker flags)
	gitCommit = ""
	// Dubaicoin address of the Gdbix release oracle.
	relOracle = common.HexToAddress("0x0")
	// The app that holds all commands and flags.
	app = utils.NewApp(gitCommit, "the go-ethereum command line interface")
)

func init() {
<<<<<<< HEAD:cmd/gtsf/main.go
	// Construct the textual version string from the individual components
	verString = fmt.Sprintf("%d.%d.%d", versionMajor, versionMinor, versionPatch)
	if versionMeta != "" {
		verString += "-" + versionMeta
	}
	if gitCommit != "" {
		verString += "-" + gitCommit[:8]
	}
	// Construct the version release oracle configuration
	relConfig.Oracle = common.HexToAddress(versionOracle)

	relConfig.Major = uint32(versionMajor)
	relConfig.Minor = uint32(versionMinor)
	relConfig.Patch = uint32(versionPatch)

	commit, _ := hex.DecodeString(gitCommit)
	copy(relConfig.Commit[:], commit)

	// Initialize the CLI app and start Gtsf
	app = utils.NewApp(verString, "the go-teslafunds command line interface")
	app.Action = gtsf
=======
	// Initialize the CLI app and start Gdbix
	app.Action = gdbix
>>>>>>> 7fdd714... gdbix-update v1.5.0:cmd/gdbix/main.go
	app.HideVersion = true // we have a command to print the version
	app.Copyright = "Copyright 2013-2016 The go-ethereum Authors"
	app.Commands = []cli.Command{
		// See chaincmd.go:
		initCommand,
		importCommand,
		exportCommand,
		upgradedbCommand,
		removedbCommand,
		dumpCommand,
		// See monitorcmd.go:
		monitorCommand,
		// See accountcmd.go:
		accountCommand,
		walletCommand,
		// See consolecmd.go:
		consoleCommand,
		attachCommand,
		javascriptCommand,
<<<<<<< HEAD:cmd/gtsf/main.go
		{
			Action: makedag,
			Name:   "makedag",
			Usage:  "generate ethash dag (for testing)",
			Description: `
The makedag command generates an ethash DAG in /tmp/dag.

This command exists to support the system testing project.
Regular users do not need to execute it.
`,
		},
		{
			Action: gpuinfo,
			Name:   "gpuinfo",
			Usage:  "gpuinfo",
			Description: `
Prints OpenCL device info for all found GPUs.
`,
		},
		{
			Action: gpubench,
			Name:   "gpubench",
			Usage:  "benchmark GPU",
			Description: `
Runs quick benchmark on first GPU found.
`,
		},
		{
			Action: version,
			Name:   "version",
			Usage:  "print teslafunds version numbers",
			Description: `
The output of this command is supposed to be machine-readable.
`,
		},
		{
			Action: initGenesis,
			Name:   "init",
			Usage:  "bootstraps and initialises a new genesis block (JSON)",
			Description: `
The init command initialises a new genesis block and definition for the network.
This is a destructive action and changes the network in which you will be
participating.
`,
		},
=======
		// See misccmd.go:
		makedagCommand,
		versionCommand,
		licenseCommand,
>>>>>>> 7fdd714... gdbix-update v1.5.0:cmd/gdbix/main.go
	}

	app.Flags = []cli.Flag{
		utils.IdentityFlag,
		utils.UnlockedAccountFlag,
		utils.PasswordFileFlag,
		utils.BootnodesFlag,
		utils.DataDirFlag,
		utils.KeyStoreDirFlag,
		utils.FastSyncFlag,
		utils.LightModeFlag,
		utils.LightServFlag,
		utils.LightPeersFlag,
		utils.LightKDFFlag,
		utils.CacheFlag,
		utils.TrieCacheGenFlag,
		utils.JSpathFlag,
		utils.ListenPortFlag,
		utils.MaxPeersFlag,
		utils.MaxPendingPeersFlag,
		utils.EtherbaseFlag,
		utils.GasPriceFlag,
		utils.MinerThreadsFlag,
		utils.MiningEnabledFlag,
		utils.AutoDAGFlag,
		utils.TargetGasLimitFlag,
		utils.NATFlag,
		utils.NoDiscoverFlag,
		utils.DiscoveryV5Flag,
		utils.NetrestrictFlag,
		utils.NodeKeyFileFlag,
		utils.NodeKeyHexFlag,
		utils.RPCEnabledFlag,
		utils.RPCListenAddrFlag,
		utils.RPCPortFlag,
		utils.RPCApiFlag,
		utils.WSEnabledFlag,
		utils.WSListenAddrFlag,
		utils.WSPortFlag,
		utils.WSApiFlag,
		utils.WSAllowedOriginsFlag,
		utils.IPCDisabledFlag,
		utils.IPCApiFlag,
		utils.IPCPathFlag,
		utils.ExecFlag,
		utils.PreloadJSFlag,
		utils.WhisperEnabledFlag,
		utils.DevModeFlag,
		utils.TestNetFlag,
		utils.VMForceJitFlag,
		utils.VMJitCacheFlag,
		utils.VMEnableJitFlag,
		utils.VMEnableDebugFlag,
		utils.NetworkIdFlag,
		utils.RPCCORSDomainFlag,
		utils.EthStatsURLFlag,
		utils.MetricsEnabledFlag,
		utils.FakePoWFlag,
		utils.SolcPathFlag,
		utils.GpoMinGasPriceFlag,
		utils.GpoMaxGasPriceFlag,
		utils.GpoFullBlockRatioFlag,
		utils.GpobaseStepDownFlag,
		utils.GpobaseStepUpFlag,
		utils.GpobaseCorrectionFactorFlag,
		utils.ExtraDataFlag,
	}
	app.Flags = append(app.Flags, debug.Flags...)

	app.Before = func(ctx *cli.Context) error {
		runtime.GOMAXPROCS(runtime.NumCPU())
		if err := debug.Setup(ctx); err != nil {
			return err
		}
		// Start system runtime metrics collection
		go metrics.CollectProcessMetrics(3 * time.Second)

		// This should be the only place where reporting is enabled
		// because it is not intended to run while testing.
		// In addition to this check, bad block reports are sent only
<<<<<<< HEAD:cmd/gtsf/main.go
		// for chains with the main network genesis block and network id 7995.
		tsf.EnableBadBlockReporting = true
=======
		// for chains with the main network genesis block and network id 1.
		eth.EnableBadBlockReporting = true
>>>>>>> 7fdd714... gdbix-update v1.5.0:cmd/gdbix/main.go

		utils.SetupNetwork(ctx)
		return nil
	}

	app.After = func(ctx *cli.Context) error {
		debug.Exit()
		console.Stdin.Close() // Resets terminal mode.
		return nil
	}
}

func main() {
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// gdbix is the main entry point into the system if no special subcommand is ran.
// It creates a default node based on the command line arguments and runs it in
// blocking mode, waiting for it to be shut down.
func gdbix(ctx *cli.Context) error {
	node := makeFullNode(ctx)
	startNode(ctx, node)
	node.Wait()
	return nil
}

func makeFullNode(ctx *cli.Context) *node.Node {
	// Create the default extradata and construct the base node
	var clientInfo = struct {
		Version   uint
		Name      string
		GoVersion string
		Os        string
	}{uint(params.VersionMajor<<16 | params.VersionMinor<<8 | params.VersionPatch), clientIdentifier, runtime.Version(), runtime.GOOS}
	extra, err := rlp.EncodeToBytes(clientInfo)
	if err != nil {
		glog.V(logger.Warn).Infoln("error setting canonical miner information:", err)
	}
	if uint64(len(extra)) > params.MaximumExtraDataSize.Uint64() {
		glog.V(logger.Warn).Infoln("error setting canonical miner information: extra exceeds", params.MaximumExtraDataSize)
		glog.V(logger.Debug).Infof("extra: %x\n", extra)
<<<<<<< HEAD:cmd/gtsf/main.go
		return nil
	}
	return extra
}

// gtsf is the main entry point into the system if no special subcommand is ran.
// It creates a default node based on the command line arguments and runs it in
// blocking mode, waiting for it to be shut down.
func gtsf(ctx *cli.Context) error {
	node := utils.MakeSystemNode(clientIdentifier, verString, relConfig, makeDefaultExtra(), ctx)
	startNode(ctx, node)
	node.Wait()

	return nil
}

// initGenesis will initialise the given JSON format genesis file and writes it as
// the zero'd block (i.e. genesis) or will fail hard if it can't succeed.
func initGenesis(ctx *cli.Context) error {
	genesisPath := ctx.Args().First()
	if len(genesisPath) == 0 {
		utils.Fatalf("must supply path to genesis JSON file")
=======
		extra = nil
>>>>>>> 7fdd714... gdbix-update v1.5.0:cmd/gdbix/main.go
	}
	stack := utils.MakeNode(ctx, clientIdentifier, gitCommit)
	utils.RegisterEthService(ctx, stack, extra)

	// Whisper must be explicitly enabled, but is auto-enabled in --dev mode.
	shhEnabled := ctx.GlobalBool(utils.WhisperEnabledFlag.Name)
	shhAutoEnabled := !ctx.GlobalIsSet(utils.WhisperEnabledFlag.Name) && ctx.GlobalIsSet(utils.DevModeFlag.Name)
	if shhEnabled || shhAutoEnabled {
		utils.RegisterShhService(stack)
	}
	// Add the Dubaicoin Stats daemon if requested
	if url := ctx.GlobalString(utils.EthStatsURLFlag.Name); url != "" {
		utils.RegisterEthStatsService(stack, url)
	}
	// Add the release oracle service so it boots along with node.
	if err := stack.Register(func(ctx *node.ServiceContext) (node.Service, error) {
		config := release.Config{
			Oracle: relOracle,
			Major:  uint32(params.VersionMajor),
			Minor:  uint32(params.VersionMinor),
			Patch:  uint32(params.VersionPatch),
		}
		commit, _ := hex.DecodeString(gitCommit)
		copy(config.Commit[:], commit)
		return release.NewReleaseService(ctx, config)
	}); err != nil {
		utils.Fatalf("Failed to register the Gdbix release oracle service: %v", err)
	}
	return stack
}

// startNode boots up the system node and all registered protocols, after which
// it unlocks any requested accounts, and starts the RPC/IPC interfaces and the
// miner.
func startNode(ctx *cli.Context, stack *node.Node) {
	// Start up the node itself
	utils.StartNode(stack)

	// Unlock any account specifically requested
<<<<<<< HEAD:cmd/gtsf/main.go
	var teslafunds *tsf.Teslafunds
	if err := stack.Service(&teslafunds); err != nil {
		utils.Fatalf("ethereum service not running: %v", err)
	}
	accman := teslafunds.AccountManager()
=======
	accman := stack.AccountManager()
>>>>>>> 7fdd714... gdbix-update v1.5.0:cmd/gdbix/main.go
	passwords := utils.MakePasswordList(ctx)
	accounts := strings.Split(ctx.GlobalString(utils.UnlockedAccountFlag.Name), ",")
	for i, account := range accounts {
		if trimmed := strings.TrimSpace(account); trimmed != "" {
			unlockAccount(ctx, accman, trimmed, i, passwords)
		}
	}
	// Start auxiliary services if enabled
	if ctx.GlobalBool(utils.MiningEnabledFlag.Name) {
<<<<<<< HEAD:cmd/gtsf/main.go
		if err := teslafunds.StartMining(ctx.GlobalInt(utils.MinerThreadsFlag.Name), ctx.GlobalString(utils.MiningGPUFlag.Name)); err != nil {
			utils.Fatalf("Failed to start mining: %v", err)
		}
	}
}

func makedag(ctx *cli.Context) error {
	args := ctx.Args()
	wrongArgs := func() {
		utils.Fatalf(`Usage: gtsf makedag <block number> <outputdir>`)
	}
	switch {
	case len(args) == 2:
		blockNum, err := strconv.ParseUint(args[0], 0, 64)
		dir := args[1]
		if err != nil {
			wrongArgs()
		} else {
			dir = filepath.Clean(dir)
			// seems to require a trailing slash
			if !strings.HasSuffix(dir, "/") {
				dir = dir + "/"
			}
			_, err = ioutil.ReadDir(dir)
			if err != nil {
				utils.Fatalf("Can't find dir")
			}
			fmt.Println("making DAG, this could take awhile...")
			ethash.MakeDAG(blockNum, dir)
		}
	default:
		wrongArgs()
	}
	return nil
}

func gpuinfo(ctx *cli.Context) error {
	tsf.PrintOpenCLDevices()
	return nil
}

func gpubench(ctx *cli.Context) error {
	args := ctx.Args()
	wrongArgs := func() {
		utils.Fatalf(`Usage: gtsf gpubench <gpu number>`)
	}
	switch {
	case len(args) == 1:
		n, err := strconv.ParseUint(args[0], 0, 64)
		if err != nil {
			wrongArgs()
		}
		tsf.GPUBench(n)
	case len(args) == 0:
		tsf.GPUBench(0)
	default:
		wrongArgs()
	}
	return nil
}

func version(c *cli.Context) error {
	fmt.Println(clientIdentifier)
	fmt.Println("Version:", verString)
	fmt.Println("Protocol Versions:", tsf.ProtocolVersions)
	fmt.Println("Network Id:", c.GlobalInt(utils.NetworkIdFlag.Name))
	fmt.Println("Go Version:", runtime.Version())
	fmt.Println("OS:", runtime.GOOS)
	fmt.Printf("GOPATH=%s\n", os.Getenv("GOPATH"))
	fmt.Printf("GOROOT=%s\n", runtime.GOROOT())

	return nil
=======
		var ethereum *eth.Ethereum
		if err := stack.Service(&ethereum); err != nil {
			utils.Fatalf("ethereum service not running: %v", err)
		}
		if err := ethereum.StartMining(ctx.GlobalInt(utils.MinerThreadsFlag.Name)); err != nil {
			utils.Fatalf("Failed to start mining: %v", err)
		}
	}
>>>>>>> 7fdd714... gdbix-update v1.5.0:cmd/gdbix/main.go
}
