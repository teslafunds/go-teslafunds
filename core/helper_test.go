<<<<<<< HEAD
// Copyright 2014 The go-ethereum Authors && Copyright 2015 go-teslafunds Authors
// This file is part of the go-teslafunds library.
//
// The go-teslafunds library is free software: you can redistribute it and/or modify
=======
// Copyright 2014 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
>>>>>>> 7fdd714... gdbix-update v1.5.0
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
<<<<<<< HEAD
// The go-teslafunds library is distributed in the hope that it will be useful,
=======
// The go-ethereum library is distributed in the hope that it will be useful,
>>>>>>> 7fdd714... gdbix-update v1.5.0
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
<<<<<<< HEAD
// along with the go-teslafunds library. If not, see <http://www.gnu.org/licenses/>.
=======
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.
>>>>>>> 7fdd714... gdbix-update v1.5.0

package core

import (
	"container/list"
	"fmt"

<<<<<<< HEAD
	"github.com/teslafunds/go-teslafunds/core/types"
	// "github.com/teslafunds/go-teslafunds/crypto"

	"github.com/teslafunds/go-teslafunds/ethdb"
	"github.com/teslafunds/go-teslafunds/event"
=======
	"github.com/dubaicoin-dbix/go-dubaicoin/core/types"
	// "github.com/dubaicoin-dbix/go-dubaicoin/crypto"

	"github.com/dubaicoin-dbix/go-dubaicoin/dbixdb"
	"github.com/dubaicoin-dbix/go-dubaicoin/event"
>>>>>>> 7fdd714... gdbix-update v1.5.0
)

// Implement our EthTest Manager
type TestManager struct {
	// stateManager *StateManager
	eventMux *event.TypeMux

	db         ethdb.Database
	txPool     *TxPool
	blockChain *BlockChain
	Blocks     []*types.Block
}

func (s *TestManager) IsListening() bool {
	return false
}

func (s *TestManager) IsMining() bool {
	return false
}

func (s *TestManager) PeerCount() int {
	return 0
}

func (s *TestManager) Peers() *list.List {
	return list.New()
}

func (s *TestManager) BlockChain() *BlockChain {
	return s.blockChain
}

func (tm *TestManager) TxPool() *TxPool {
	return tm.txPool
}

// func (tm *TestManager) StateManager() *StateManager {
// 	return tm.stateManager
// }

func (tm *TestManager) EventMux() *event.TypeMux {
	return tm.eventMux
}

// func (tm *TestManager) KeyManager() *crypto.KeyManager {
// 	return nil
// }

func (tm *TestManager) Db() ethdb.Database {
	return tm.db
}

func NewTestManager() *TestManager {
	db, err := ethdb.NewMemDatabase()
	if err != nil {
		fmt.Println("Could not create mem-db, failing")
		return nil
	}

	testManager := &TestManager{}
	testManager.eventMux = new(event.TypeMux)
	testManager.db = db
	// testManager.txPool = NewTxPool(testManager)
	// testManager.blockChain = NewBlockChain(testManager)
	// testManager.stateManager = NewStateManager(testManager)

	return testManager
}
