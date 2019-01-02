// Copyright 2016 The go-ethereum Authors
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

<<<<<<< HEAD:tsf/bad_block.go
package tsf
=======
package eth
>>>>>>> 7fdd714... gdbix-update v1.5.0:dbix/bad_block.go

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

<<<<<<< HEAD:tsf/bad_block.go
	"github.com/teslafunds/go-teslafunds/common"
	"github.com/teslafunds/go-teslafunds/core/types"
	"github.com/teslafunds/go-teslafunds/logger"
	"github.com/teslafunds/go-teslafunds/logger/glog"
	"github.com/teslafunds/go-teslafunds/rlp"
)

const (
	// The Teslafunds main network genesis block hash.
	defaultGenesisHash = "0x4f09f80efaa0ac22046320f6afa92b96371343f7d6da68d2d7d1b44dcc0bc629"
	badBlocksURL       = "https://badblocks.tsfchain.io"
=======
	"github.com/dubaicoin-dbix/go-dubaicoin/common"
	"github.com/dubaicoin-dbix/go-dubaicoin/core/types"
	"github.com/dubaicoin-dbix/go-dubaicoin/logger"
	"github.com/dubaicoin-dbix/go-dubaicoin/logger/glog"
	"github.com/dubaicoin-dbix/go-dubaicoin/rlp"
)

const (
	// The Dubaicoin main network genesis block.
	defaultGenesisHash = "0x4f09f80efaa0ac22046320f6afa92b96371343f7d6da68d2d7d1b44dcc0bc629"
	badBlocksURL       = "https://badblocks.dbixscan.io"
>>>>>>> 7fdd714... gdbix-update v1.5.0:dbix/bad_block.go
)

var EnableBadBlockReporting = false

func sendBadBlockReport(block *types.Block, err error) {
	if !EnableBadBlockReporting {
		return
	}

	var (
		blockRLP, _ = rlp.EncodeToBytes(block)
		params      = map[string]interface{}{
			"block":     common.Bytes2Hex(blockRLP),
			"blockHash": block.Hash().Hex(),
			"errortype": err.Error(),
			"client":    "go",
		}
	)
	if !block.ReceivedAt.IsZero() {
		params["receivedAt"] = block.ReceivedAt.UTC().String()
	}
	if p, ok := block.ReceivedFrom.(*peer); ok {
		params["receivedFrom"] = map[string]interface{}{
			"enode":           fmt.Sprintf("enode://%x@%v", p.ID(), p.RemoteAddr()),
			"name":            p.Name(),
			"protocolVersion": p.version,
		}
	}
	jsonStr, _ := json.Marshal(map[string]interface{}{"method": "eth_badBlock", "id": "1", "jsonrpc": "2.0", "params": []interface{}{params}})
	client := http.Client{Timeout: 8 * time.Second}
	resp, err := client.Post(badBlocksURL, "application/json", bytes.NewReader(jsonStr))
	if err != nil {
		glog.V(logger.Debug).Infoln(err)
		return
	}
	glog.V(logger.Debug).Infof("Bad Block Report posted (%d)", resp.StatusCode)
	resp.Body.Close()
}
