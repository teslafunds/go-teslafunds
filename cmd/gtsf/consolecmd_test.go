// Copyright 2016 The go-ethereum Authors
// This file is part of go-ethereum.
//
// go-ethereum is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// go-ethereum is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with go-ethereum. If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"crypto/rand"
	"math/big"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"testing"
	"time"

<<<<<<< HEAD:cmd/gtsf/consolecmd_test.go
	"github.com/teslafunds/go-teslafunds/rpc"
=======
	"github.com/dubaicoin-dbix/go-dubaicoin/params"
	"github.com/dubaicoin-dbix/go-dubaicoin/rpc"
>>>>>>> 7fdd714... gdbix-update v1.5.0:cmd/gdbix/consolecmd_test.go
)

// Tests that a node embedded within a console can be started up properly and
// then terminated by closing the input stream.
func TestConsoleWelcome(t *testing.T) {
	coinbase := "0x8605cdbbdb6d264aa742e77020dcbc58fcdce182"

<<<<<<< HEAD:cmd/gtsf/consolecmd_test.go
	// Start a gtsf console, make sure it's cleaned up and terminate the console
	gtsf := runGtsf(t,
=======
	// Start a gdbix console, make sure it's cleaned up and terminate the console
	gdbix := runGeth(t,
>>>>>>> 7fdd714... gdbix-update v1.5.0:cmd/gdbix/consolecmd_test.go
		"--port", "0", "--maxpeers", "0", "--nodiscover", "--nat", "none",
		"--etherbase", coinbase, "--shh",
		"console")

	// Gather all the infos the welcome message needs to contain
<<<<<<< HEAD:cmd/gtsf/consolecmd_test.go
	gtsf.setTemplateFunc("goos", func() string { return runtime.GOOS })
	gtsf.setTemplateFunc("gover", runtime.Version)
	gtsf.setTemplateFunc("gtsfver", func() string { return verString })
	gtsf.setTemplateFunc("niltime", func() string { return time.Unix(0, 0).Format(time.RFC1123) })
	gtsf.setTemplateFunc("apis", func() []string {
=======
	gdbix.setTemplateFunc("goos", func() string { return runtime.GOOS })
	gdbix.setTemplateFunc("gover", runtime.Version)
	gdbix.setTemplateFunc("gethver", func() string { return params.Version })
	gdbix.setTemplateFunc("niltime", func() string { return time.Unix(0, 0).Format(time.RFC1123) })
	gdbix.setTemplateFunc("apis", func() []string {
>>>>>>> 7fdd714... gdbix-update v1.5.0:cmd/gdbix/consolecmd_test.go
		apis := append(strings.Split(rpc.DefaultIPCApis, ","), rpc.MetadataApi)
		sort.Strings(apis)
		return apis
	})

	// Verify the actual welcome message to the required template
	gtsf.expect(`
Welcome to the Gtsf JavaScript console!

<<<<<<< HEAD:cmd/gtsf/consolecmd_test.go
instance: Gtsf/v{{gtsfver}}/{{goos}}/{{gover}}
=======
instance: Gdbix/v{{gethver}}/{{goos}}/{{gover}}
>>>>>>> 7fdd714... gdbix-update v1.5.0:cmd/gdbix/consolecmd_test.go
coinbase: {{.Etherbase}}
at block: 0 ({{niltime}})
 datadir: {{.Datadir}}
 modules:{{range apis}} {{.}}:1.0{{end}}

> {{.InputLine "exit"}}
`)
	gtsf.expectExit()
}

// Tests that a console can be attached to a running node via various means.
func TestIPCAttachWelcome(t *testing.T) {
	// Configure the instance for IPC attachement
	coinbase := "0x8605cdbbdb6d264aa742e77020dcbc58fcdce182"
	var ipc string
	if runtime.GOOS == "windows" {
		ipc = `\\.\pipe\gtsf` + strconv.Itoa(trulyRandInt(100000, 999999))
	} else {
		ws := tmpdir(t)
		defer os.RemoveAll(ws)
		ipc = filepath.Join(ws, "gtsf.ipc")
	}
	// Note: we need --shh because testAttachWelcome checks for default
	// list of ipc modules and shh is included there.
<<<<<<< HEAD:cmd/gtsf/consolecmd_test.go
	gtsf := runGtsf(t,
=======
	gdbix := runGeth(t,
>>>>>>> 7fdd714... gdbix-update v1.5.0:cmd/gdbix/consolecmd_test.go
		"--port", "0", "--maxpeers", "0", "--nodiscover", "--nat", "none",
		"--etherbase", coinbase, "--shh", "--ipcpath", ipc)

	time.Sleep(2 * time.Second) // Simple way to wait for the RPC endpoint to open
	testAttachWelcome(t, gtsf, "ipc:"+ipc)

	gtsf.interrupt()
	gtsf.expectExit()
}

func TestHTTPAttachWelcome(t *testing.T) {
	coinbase := "0x8605cdbbdb6d264aa742e77020dcbc58fcdce182"
	port := strconv.Itoa(trulyRandInt(1024, 65536)) // Yeah, sometimes this will fail, sorry :P
<<<<<<< HEAD:cmd/gtsf/consolecmd_test.go
	gtsf := runGtsf(t,
=======
	gdbix := runGeth(t,
>>>>>>> 7fdd714... gdbix-update v1.5.0:cmd/gdbix/consolecmd_test.go
		"--port", "0", "--maxpeers", "0", "--nodiscover", "--nat", "none",
		"--etherbase", coinbase, "--rpc", "--rpcport", port)

	time.Sleep(2 * time.Second) // Simple way to wait for the RPC endpoint to open
	testAttachWelcome(t, gtsf, "http://localhost:"+port)

	gtsf.interrupt()
	gtsf.expectExit()
}

func TestWSAttachWelcome(t *testing.T) {
	coinbase := "0x8605cdbbdb6d264aa742e77020dcbc58fcdce182"
	port := strconv.Itoa(trulyRandInt(1024, 65536)) // Yeah, sometimes this will fail, sorry :P

<<<<<<< HEAD:cmd/gtsf/consolecmd_test.go
	gtsf := runGtsf(t,
=======
	gdbix := runGeth(t,
>>>>>>> 7fdd714... gdbix-update v1.5.0:cmd/gdbix/consolecmd_test.go
		"--port", "0", "--maxpeers", "0", "--nodiscover", "--nat", "none",
		"--etherbase", coinbase, "--ws", "--wsport", port)

	time.Sleep(2 * time.Second) // Simple way to wait for the RPC endpoint to open
	testAttachWelcome(t, gtsf, "ws://localhost:"+port)

	gtsf.interrupt()
	gtsf.expectExit()
}

<<<<<<< HEAD:cmd/gtsf/consolecmd_test.go
func testAttachWelcome(t *testing.T, gtsf *testgtsf, endpoint string) {
	// Attach to a running gtsf note and terminate immediately
	attach := runGtsf(t, "attach", endpoint)
=======
func testAttachWelcome(t *testing.T, gdbix *testgeth, endpoint string) {
	// Attach to a running gdbix note and terminate immediately
	attach := runGeth(t, "attach", endpoint)
>>>>>>> 7fdd714... gdbix-update v1.5.0:cmd/gdbix/consolecmd_test.go
	defer attach.expectExit()
	attach.stdin.Close()

	// Gather all the infos the welcome message needs to contain
	attach.setTemplateFunc("goos", func() string { return runtime.GOOS })
	attach.setTemplateFunc("gover", runtime.Version)
<<<<<<< HEAD:cmd/gtsf/consolecmd_test.go
	attach.setTemplateFunc("gtsfver", func() string { return verString })
	attach.setTemplateFunc("etherbase", func() string { return gtsf.Etherbase })
=======
	attach.setTemplateFunc("gethver", func() string { return params.Version })
	attach.setTemplateFunc("etherbase", func() string { return gdbix.Etherbase })
>>>>>>> 7fdd714... gdbix-update v1.5.0:cmd/gdbix/consolecmd_test.go
	attach.setTemplateFunc("niltime", func() string { return time.Unix(0, 0).Format(time.RFC1123) })
	attach.setTemplateFunc("ipc", func() bool { return strings.HasPrefix(endpoint, "ipc") })
	attach.setTemplateFunc("datadir", func() string { return gtsf.Datadir })
	attach.setTemplateFunc("apis", func() []string {
		var apis []string
		if strings.HasPrefix(endpoint, "ipc") {
			apis = append(strings.Split(rpc.DefaultIPCApis, ","), rpc.MetadataApi)
		} else {
			apis = append(strings.Split(rpc.DefaultHTTPApis, ","), rpc.MetadataApi)
		}
		sort.Strings(apis)
		return apis
	})

	// Verify the actual welcome message to the required template
	attach.expect(`
Welcome to the Gtsf JavaScript console!

<<<<<<< HEAD:cmd/gtsf/consolecmd_test.go
instance: Gtsf/v{{gtsfver}}/{{goos}}/{{gover}}
=======
instance: Gdbix/v{{gethver}}/{{goos}}/{{gover}}
>>>>>>> 7fdd714... gdbix-update v1.5.0:cmd/gdbix/consolecmd_test.go
coinbase: {{etherbase}}
at block: 0 ({{niltime}}){{if ipc}}
 datadir: {{datadir}}{{end}}
 modules:{{range apis}} {{.}}:1.0{{end}}

> {{.InputLine "exit" }}
`)
	attach.expectExit()
}

// trulyRandInt generates a crypto random integer used by the console tests to
// not clash network ports with other tests running cocurrently.
func trulyRandInt(lo, hi int) int {
	num, _ := rand.Int(rand.Reader, big.NewInt(int64(hi-lo)))
	return int(num.Int64()) + lo
}
