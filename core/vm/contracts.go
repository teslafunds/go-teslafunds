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

package vm

import (
	"math/big"

<<<<<<< HEAD
	"github.com/teslafunds/go-teslafunds/common"
	"github.com/teslafunds/go-teslafunds/crypto"
	"github.com/teslafunds/go-teslafunds/logger"
	"github.com/teslafunds/go-teslafunds/logger/glog"
	"github.com/teslafunds/go-teslafunds/params"
)

// PrecompiledAccount represents a native teslafunds contract
type PrecompiledAccount struct {
	Gas func(l int) *big.Int
	fn  func(in []byte) []byte
=======
	"github.com/dubaicoin-dbix/go-dubaicoin/common"
	"github.com/dubaicoin-dbix/go-dubaicoin/crypto"
	"github.com/dubaicoin-dbix/go-dubaicoin/logger"
	"github.com/dubaicoin-dbix/go-dubaicoin/logger/glog"
	"github.com/dubaicoin-dbix/go-dubaicoin/params"
)

// Precompiled contract is the basic interface for native Go contracts. The implementation
// requires a deterministic gas count based on the input size of the Run method of the
// contract.
type PrecompiledContract interface {
	RequiredGas(inputSize int) *big.Int // RequiredPrice calculates the contract gas use
	Run(input []byte) []byte            // Run runs the precompiled contract
>>>>>>> 7fdd714... gdbix-update v1.5.0
}

// Precompiled contains the default set of ethereum contracts
var PrecompiledContracts = map[common.Address]PrecompiledContract{
	common.BytesToAddress([]byte{1}): &ecrecover{},
	common.BytesToAddress([]byte{2}): &sha256{},
	common.BytesToAddress([]byte{3}): &ripemd160{},
	common.BytesToAddress([]byte{4}): &dataCopy{},
}

<<<<<<< HEAD
// Precompiled contains the default set of teslafunds contracts
var Precompiled = PrecompiledContracts()

// PrecompiledContracts returns the default set of precompiled ethereum
// contracts defined by the teslafunds yellow paper.
func PrecompiledContracts() map[string]*PrecompiledAccount {
	return map[string]*PrecompiledAccount{
		// ECRECOVER
		string(common.LeftPadBytes([]byte{1}, 20)): &PrecompiledAccount{func(l int) *big.Int {
			return params.EcrecoverGas
		}, ecrecoverFunc},

		// SHA256
		string(common.LeftPadBytes([]byte{2}, 20)): &PrecompiledAccount{func(l int) *big.Int {
			n := big.NewInt(int64(l+31) / 32)
			n.Mul(n, params.Sha256WordGas)
			return n.Add(n, params.Sha256Gas)
		}, sha256Func},

		// RIPEMD160
		string(common.LeftPadBytes([]byte{3}, 20)): &PrecompiledAccount{func(l int) *big.Int {
			n := big.NewInt(int64(l+31) / 32)
			n.Mul(n, params.Ripemd160WordGas)
			return n.Add(n, params.Ripemd160Gas)
		}, ripemd160Func},

		string(common.LeftPadBytes([]byte{4}, 20)): &PrecompiledAccount{func(l int) *big.Int {
			n := big.NewInt(int64(l+31) / 32)
			n.Mul(n, params.IdentityWordGas)

			return n.Add(n, params.IdentityGas)
		}, memCpy},
=======
// RunPrecompile runs and evaluate the output of a precompiled contract defined in contracts.go
func RunPrecompiledContract(p PrecompiledContract, input []byte, contract *Contract) (ret []byte, err error) {
	gas := p.RequiredGas(len(input))
	if contract.UseGas(gas) {
		ret = p.Run(input)

		return ret, nil
	} else {
		return nil, ErrOutOfGas
>>>>>>> 7fdd714... gdbix-update v1.5.0
	}
}

// ECRECOVER implemented as a native contract
type ecrecover struct{}

func (c *ecrecover) RequiredGas(inputSize int) *big.Int {
	return params.EcrecoverGas
}

func (c *ecrecover) Run(in []byte) []byte {
	const ecRecoverInputLength = 128

	in = common.RightPadBytes(in, ecRecoverInputLength)
	// "in" is (hash, v, r, s), each 32 bytes
	// but for ecrecover we want (r, s, v)

	r := common.BytesToBig(in[64:96])
	s := common.BytesToBig(in[96:128])
	v := in[63] - 27

	// tighter sig s values in homestead only apply to tx sigs
	if common.Bytes2Big(in[32:63]).BitLen() > 0 || !crypto.ValidateSignatureValues(v, r, s, false) {
		glog.V(logger.Detail).Infof("ECRECOVER error: v, r or s value invalid")
		return nil
	}
	// v needs to be at the end for libsecp256k1
	pubKey, err := crypto.Ecrecover(in[:32], append(in[64:128], v))
	// make sure the public key is a valid one
	if err != nil {
		glog.V(logger.Detail).Infoln("ECRECOVER error: ", err)
		return nil
	}

	// the first byte of pubkey is bitcoin heritage
	return common.LeftPadBytes(crypto.Keccak256(pubKey[1:])[12:], 32)
}

// SHA256 implemented as a native contract
type sha256 struct{}

func (c *sha256) RequiredGas(inputSize int) *big.Int {
	n := big.NewInt(int64(inputSize+31) / 32)
	n.Mul(n, params.Sha256WordGas)
	return n.Add(n, params.Sha256Gas)
}
func (c *sha256) Run(in []byte) []byte {
	return crypto.Sha256(in)
}

// RIPMED160 implemented as a native contract
type ripemd160 struct{}

func (c *ripemd160) RequiredGas(inputSize int) *big.Int {
	n := big.NewInt(int64(inputSize+31) / 32)
	n.Mul(n, params.Ripemd160WordGas)
	return n.Add(n, params.Ripemd160Gas)
}
func (c *ripemd160) Run(in []byte) []byte {
	return common.LeftPadBytes(crypto.Ripemd160(in), 32)
}

// data copy implemented as a native contract
type dataCopy struct{}

func (c *dataCopy) RequiredGas(inputSize int) *big.Int {
	n := big.NewInt(int64(inputSize+31) / 32)
	n.Mul(n, params.IdentityWordGas)

	return n.Add(n, params.IdentityGas)
}
func (c *dataCopy) Run(in []byte) []byte {
	return in
}
