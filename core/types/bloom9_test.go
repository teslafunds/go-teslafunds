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

package types

import (
	"math/big"
	"testing"
)

func TestBloom(t *testing.T) {
	positive := []string{
		"testtest",
		"test",
		"hallo",
		"other",
	}
	negative := []string{
		"tes",
		"lo",
	}

	var bloom Bloom
	for _, data := range positive {
		bloom.Add(new(big.Int).SetBytes([]byte(data)))
	}

	for _, data := range positive {
		if !bloom.TestBytes([]byte(data)) {
			t.Error("expected", data, "to test true")
		}
	}
	for _, data := range negative {
		if bloom.TestBytes([]byte(data)) {
			t.Error("did not expect", data, "to test true")
		}
	}
}

/*
import (
	"testing"

<<<<<<< HEAD
	"github.com/teslafunds/go-teslafunds/core/state"
=======
	"github.com/dubaicoin-dbix/go-dubaicoin/core/state"
>>>>>>> 7fdd714... gdbix-update v1.5.0
)

func TestBloom9(t *testing.T) {
	testCase := []byte("testtest")
	bin := LogsBloom([]state.Log{
		{testCase, [][]byte{[]byte("hellohello")}, nil},
	}).Bytes()
	res := BloomLookup(bin, testCase)

	if !res {
		t.Errorf("Bloom lookup failed")
	}
}


func TestAddress(t *testing.T) {
	block := &Block{}
	block.Coinbase = common.Hex2Bytes("22341ae42d6dd7384bc8584e50419ea3ac75b83f")
	fmt.Printf("%x\n", crypto.Keccak256(block.Coinbase))

	bin := CreateBloom(block)
	fmt.Printf("bin = %x\n", common.LeftPadBytes(bin, 64))
}
*/
