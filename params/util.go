<<<<<<< HEAD
// Copyright 2015 The go-teslafunds Authors
// This file is part of the go-teslafunds library.
//
// The go-teslafunds library is free software: you can redistribute it and/or modify
=======
// Copyright 2015 The go-ethereum Authors
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

package params

import (
	"math/big"

	"github.com/dubaicoin-dbix/go-dubaicoin/common"
)

var (
	TestNetGenesisHash = common.HexToHash("0x991a2a023969dac3739cea2ba92df6c40cf494f156fbc76ea17469a81e412da5") // Testnet genesis hash to enforce below configs on
	MainNetGenesisHash = common.HexToHash("0x4f09f80efaa0ac22046320f6afa92b96371343f7d6da68d2d7d1b44dcc0bc629") // Mainnet genesis hash to enforce below configs on

	TestNetHomesteadBlock = big.NewInt(0)       // Testnet homestead block
	MainNetHomesteadBlock = big.NewInt(90000) // Mainnet homestead block

	TestNetHomesteadGasRepriceBlock = big.NewInt(0)       // Testnet gas reprice block
	MainNetHomesteadGasRepriceBlock = big.NewInt(150000) // Mainnet gas reprice block

	TestNetHomesteadGasRepriceHash = common.HexToHash("0x991a2a023969dac3739cea2ba92df6c40cf494f156fbc76ea17469a81e412da5") // Testnet gas reprice block hash (used by fast sync) // TestNet--GasRepriceHash updated at block 0
	//
	MainNetHomesteadGasRepriceHash = common.HexToHash("0x0") // Mainnet gas reprice block hash (used by fast sync) // MainNet--GasRepriceHash will be updated at Block 150000

	TestNetSpuriousDragon = big.NewInt(10)
	MainNetSpuriousDragon = big.NewInt(190000)

	TestNetChainID = big.NewInt(6) // Test net default chain ID
	MainNetChainID = big.NewInt(5) // main net default chain ID
)
