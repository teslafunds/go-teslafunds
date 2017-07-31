<<<<<<< HEAD
// Copyright 2014 The go-ethereum Authors
// Copyright 2015 go-teslafunds Authors
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

package vm

import (
	"fmt"
	"math/big"

<<<<<<< HEAD

	"github.com/teslafunds/go-teslafunds/common"
	"github.com/teslafunds/go-teslafunds/crypto"
	"github.com/teslafunds/go-teslafunds/params"
)

type programInstruction interface {
	// executes the program instruction and allows the instruction to modify the state of the program
	do(program *Program, pc *uint64, env Environment, contract *Contract, memory *Memory, Stack *Stack) ([]byte, error)
	// returns whether the program instruction halts the execution of the JIT
	halts() bool
	// Returns the current op code (debugging purposes)
	Op() OpCode
}

type instrFn func(instr instruction, pc *uint64, env Environment, contract *Contract, memory *Memory, Stack *Stack)

type instruction struct {
	op   OpCode
	pc   uint64
	fn   instrFn
	data *big.Int

	gas   *big.Int
	spop  int
	spush int

	returns bool
}

func jump(mapping map[uint64]uint64, destinations map[uint64]struct{}, contract *Contract, to *big.Int) (uint64, error) {
	if !validDest(destinations, to) {
		nop := contract.GetOp(to.Uint64())
		return 0, fmt.Errorf("invalid jump destination (%v) %v", nop, to)
	}

	return mapping[to.Uint64()], nil
}

func (instr instruction) do(program *Program, pc *uint64, env Environment, contract *Contract, memory *Memory, Stack *Stack) ([]byte, error) {
	// calculate the new memory size and gas price for the current executing opcode
	newMemSize, cost, err := jitCalculateGasAndSize(env, contract, instr, env.Db(), memory, Stack)
	if err != nil {
		return nil, err
	}

	// Use the calculated gas. When insufficient gas is present, use all gas and return an
	// Out Of Gas error
	if !contract.UseGas(cost) {
		return nil, OutOfGasError
	}
	// Resize the memory calculated previously
	memory.Resize(newMemSize.Uint64())

	// These opcodes return an argument and are therefor handled
	// differently from the rest of the opcodes
	switch instr.op {
	case JUMP:
		if pos, err := jump(program.mapping, program.destinations, contract, Stack.pop()); err != nil {
			return nil, err
		} else {
			*pc = pos
			return nil, nil
		}
	case JUMPI:
		pos, cond := Stack.pop(), Stack.pop()
		if cond.Cmp(common.BigTrue) >= 0 {
			if pos, err := jump(program.mapping, program.destinations, contract, pos); err != nil {
				return nil, err
			} else {
				*pc = pos
				return nil, nil
			}
		}
	case RETURN:
		offset, size := Stack.pop(), Stack.pop()
		return memory.GetPtr(offset.Int64(), size.Int64()), nil
	default:
		if instr.fn == nil {
			return nil, fmt.Errorf("Invalid opcode 0x%x", instr.op)
		}
		instr.fn(instr, pc, env, contract, memory, Stack)
	}
	*pc++
	return nil, nil
}

func (instr instruction) halts() bool {
	return instr.returns
}

func (instr instruction) Op() OpCode {
	return instr.op
}

func opStaticJump(instr instruction, pc *uint64, ret *big.Int, env Environment, contract *Contract, memory *Memory, Stack *Stack) {
	ret.Set(instr.data)
}

func opAdd(instr instruction, pc *uint64, env Environment, contract *Contract, memory *Memory, Stack *Stack) {
	x, y := Stack.pop(), Stack.pop()
	Stack.push(U256(x.Add(x, y)))
}

func opSub(instr instruction, pc *uint64, env Environment, contract *Contract, memory *Memory, Stack *Stack) {
	x, y := Stack.pop(), Stack.pop()
	Stack.push(U256(x.Sub(x, y)))
}

func opMul(instr instruction, pc *uint64, env Environment, contract *Contract, memory *Memory, Stack *Stack) {
	x, y := Stack.pop(), Stack.pop()
	Stack.push(U256(x.Mul(x, y)))
}

func opDiv(instr instruction, pc *uint64, env Environment, contract *Contract, memory *Memory, Stack *Stack) {
	x, y := Stack.pop(), Stack.pop()
=======
	"github.com/dubaicoin-dbix/go-dubaicoin/common"
	"github.com/dubaicoin-dbix/go-dubaicoin/common/math"
	"github.com/dubaicoin-dbix/go-dubaicoin/core/types"
	"github.com/dubaicoin-dbix/go-dubaicoin/crypto"
	"github.com/dubaicoin-dbix/go-dubaicoin/params"
)

func opAdd(pc *uint64, env *EVM, contract *Contract, memory *Memory, stack *Stack) ([]byte, error) {
	x, y := stack.pop(), stack.pop()
	stack.push(U256(x.Add(x, y)))
	return nil, nil
}

func opSub(pc *uint64, env *EVM, contract *Contract, memory *Memory, stack *Stack) ([]byte, error) {
	x, y := stack.pop(), stack.pop()
	stack.push(U256(x.Sub(x, y)))
	return nil, nil
}

func opMul(pc *uint64, env *EVM, contract *Contract, memory *Memory, stack *Stack) ([]byte, error) {
	x, y := stack.pop(), stack.pop()
	stack.push(U256(x.Mul(x, y)))
	return nil, nil
}

func opDiv(pc *uint64, env *EVM, contract *Contract, memory *Memory, stack *Stack) ([]byte, error) {
	x, y := stack.pop(), stack.pop()
>>>>>>> 7fdd714... gdbix-update v1.5.0
	if y.Cmp(common.Big0) != 0 {
		Stack.push(U256(x.Div(x, y)))
	} else {
		Stack.push(new(big.Int))
	}
	return nil, nil
}

<<<<<<< HEAD
func opSdiv(instr instruction, pc *uint64, env Environment, contract *Contract, memory *Memory, Stack *Stack) {
	x, y := S256(Stack.pop()), S256(Stack.pop())
	if y.Cmp(common.Big0) == 0 {
		Stack.push(new(big.Int))
		return
=======
func opSdiv(pc *uint64, env *EVM, contract *Contract, memory *Memory, stack *Stack) ([]byte, error) {
	x, y := S256(stack.pop()), S256(stack.pop())
	if y.Cmp(common.Big0) == 0 {
		stack.push(new(big.Int))
		return nil, nil
>>>>>>> 7fdd714... gdbix-update v1.5.0
	} else {
		n := new(big.Int)
		if new(big.Int).Mul(x, y).Cmp(common.Big0) < 0 {
			n.SetInt64(-1)
		} else {
			n.SetInt64(1)
		}

		res := x.Div(x.Abs(x), y.Abs(y))
		res.Mul(res, n)

		Stack.push(U256(res))
	}
	return nil, nil
}

<<<<<<< HEAD
func opMod(instr instruction, pc *uint64, env Environment, contract *Contract, memory *Memory, Stack *Stack) {
	x, y := Stack.pop(), Stack.pop()
=======
func opMod(pc *uint64, env *EVM, contract *Contract, memory *Memory, stack *Stack) ([]byte, error) {
	x, y := stack.pop(), stack.pop()
>>>>>>> 7fdd714... gdbix-update v1.5.0
	if y.Cmp(common.Big0) == 0 {
		Stack.push(new(big.Int))
	} else {
		Stack.push(U256(x.Mod(x, y)))
	}
	return nil, nil
}

<<<<<<< HEAD
func opSmod(instr instruction, pc *uint64, env Environment, contract *Contract, memory *Memory, Stack *Stack) {
	x, y := S256(Stack.pop()), S256(Stack.pop())
=======
func opSmod(pc *uint64, env *EVM, contract *Contract, memory *Memory, stack *Stack) ([]byte, error) {
	x, y := S256(stack.pop()), S256(stack.pop())
>>>>>>> 7fdd714... gdbix-update v1.5.0

	if y.Cmp(common.Big0) == 0 {
		Stack.push(new(big.Int))
	} else {
		n := new(big.Int)
		if x.Cmp(common.Big0) < 0 {
			n.SetInt64(-1)
		} else {
			n.SetInt64(1)
		}

		res := x.Mod(x.Abs(x), y.Abs(y))
		res.Mul(res, n)

		Stack.push(U256(res))
	}
	return nil, nil
}

<<<<<<< HEAD
func opExp(instr instruction, pc *uint64, env Environment, contract *Contract, memory *Memory, Stack *Stack) {
	x, y := Stack.pop(), Stack.pop()
	Stack.push(U256(x.Exp(x, y, Pow256)))
}

func opSignExtend(instr instruction, pc *uint64, env Environment, contract *Contract, memory *Memory, Stack *Stack) {
	back := Stack.pop()
=======
func opExp(pc *uint64, env *EVM, contract *Contract, memory *Memory, stack *Stack) ([]byte, error) {
	base, exponent := stack.pop(), stack.pop()
	stack.push(math.Exp(base, exponent))
	return nil, nil
}

func opSignExtend(pc *uint64, env *EVM, contract *Contract, memory *Memory, stack *Stack) ([]byte, error) {
	back := stack.pop()
>>>>>>> 7fdd714... gdbix-update v1.5.0
	if back.Cmp(big.NewInt(31)) < 0 {
		bit := uint(back.Uint64()*8 + 7)
		num := Stack.pop()
		mask := back.Lsh(common.Big1, bit)
		mask.Sub(mask, common.Big1)
		if common.BitTest(num, int(bit)) {
			num.Or(num, mask.Not(mask))
		} else {
			num.And(num, mask)
		}

		Stack.push(U256(num))
	}
	return nil, nil
}

<<<<<<< HEAD
func opNot(instr instruction, pc *uint64, env Environment, contract *Contract, memory *Memory, Stack *Stack) {
	x := Stack.pop()
	Stack.push(U256(x.Not(x)))
}

func opLt(instr instruction, pc *uint64, env Environment, contract *Contract, memory *Memory, Stack *Stack) {
	x, y := Stack.pop(), Stack.pop()
=======
func opNot(pc *uint64, env *EVM, contract *Contract, memory *Memory, stack *Stack) ([]byte, error) {
	x := stack.pop()
	stack.push(U256(x.Not(x)))
	return nil, nil
}

func opLt(pc *uint64, env *EVM, contract *Contract, memory *Memory, stack *Stack) ([]byte, error) {
	x, y := stack.pop(), stack.pop()
>>>>>>> 7fdd714... gdbix-update v1.5.0
	if x.Cmp(y) < 0 {
		Stack.push(big.NewInt(1))
	} else {
		Stack.push(new(big.Int))
	}
	return nil, nil
}

<<<<<<< HEAD
func opGt(instr instruction, pc *uint64, env Environment, contract *Contract, memory *Memory, Stack *Stack) {
	x, y := Stack.pop(), Stack.pop()
=======
func opGt(pc *uint64, env *EVM, contract *Contract, memory *Memory, stack *Stack) ([]byte, error) {
	x, y := stack.pop(), stack.pop()
>>>>>>> 7fdd714... gdbix-update v1.5.0
	if x.Cmp(y) > 0 {
		Stack.push(big.NewInt(1))
	} else {
		Stack.push(new(big.Int))
	}
	return nil, nil
}

<<<<<<< HEAD
func opSlt(instr instruction, pc *uint64, env Environment, contract *Contract, memory *Memory, Stack *Stack) {
	x, y := S256(Stack.pop()), S256(Stack.pop())
=======
func opSlt(pc *uint64, env *EVM, contract *Contract, memory *Memory, stack *Stack) ([]byte, error) {
	x, y := S256(stack.pop()), S256(stack.pop())
>>>>>>> 7fdd714... gdbix-update v1.5.0
	if x.Cmp(S256(y)) < 0 {
		Stack.push(big.NewInt(1))
	} else {
		Stack.push(new(big.Int))
	}
	return nil, nil
}

<<<<<<< HEAD
func opSgt(instr instruction, pc *uint64, env Environment, contract *Contract, memory *Memory, Stack *Stack) {
	x, y := S256(Stack.pop()), S256(Stack.pop())
=======
func opSgt(pc *uint64, env *EVM, contract *Contract, memory *Memory, stack *Stack) ([]byte, error) {
	x, y := S256(stack.pop()), S256(stack.pop())
>>>>>>> 7fdd714... gdbix-update v1.5.0
	if x.Cmp(y) > 0 {
		Stack.push(big.NewInt(1))
	} else {
		Stack.push(new(big.Int))
	}
	return nil, nil
}

<<<<<<< HEAD
func opEq(instr instruction, pc *uint64, env Environment, contract *Contract, memory *Memory, Stack *Stack) {
	x, y := Stack.pop(), Stack.pop()
=======
func opEq(pc *uint64, env *EVM, contract *Contract, memory *Memory, stack *Stack) ([]byte, error) {
	x, y := stack.pop(), stack.pop()
>>>>>>> 7fdd714... gdbix-update v1.5.0
	if x.Cmp(y) == 0 {
		Stack.push(big.NewInt(1))
	} else {
		Stack.push(new(big.Int))
	}
	return nil, nil
}

<<<<<<< HEAD
func opIszero(instr instruction, pc *uint64, env Environment, contract *Contract, memory *Memory, Stack *Stack) {
	x := Stack.pop()
=======
func opIszero(pc *uint64, env *EVM, contract *Contract, memory *Memory, stack *Stack) ([]byte, error) {
	x := stack.pop()
>>>>>>> 7fdd714... gdbix-update v1.5.0
	if x.Cmp(common.Big0) > 0 {
		Stack.push(new(big.Int))
	} else {
		Stack.push(big.NewInt(1))
	}
	return nil, nil
}

<<<<<<< HEAD
func opAnd(instr instruction, pc *uint64, env Environment, contract *Contract, memory *Memory, Stack *Stack) {
	x, y := Stack.pop(), Stack.pop()
	Stack.push(x.And(x, y))
}
func opOr(instr instruction, pc *uint64, env Environment, contract *Contract, memory *Memory, Stack *Stack) {
	x, y := Stack.pop(), Stack.pop()
	Stack.push(x.Or(x, y))
}
func opXor(instr instruction, pc *uint64, env Environment, contract *Contract, memory *Memory, Stack *Stack) {
	x, y := Stack.pop(), Stack.pop()
	Stack.push(x.Xor(x, y))
}
func opByte(instr instruction, pc *uint64, env Environment, contract *Contract, memory *Memory, Stack *Stack) {
	th, val := Stack.pop(), Stack.pop()
=======
func opAnd(pc *uint64, env *EVM, contract *Contract, memory *Memory, stack *Stack) ([]byte, error) {
	x, y := stack.pop(), stack.pop()
	stack.push(x.And(x, y))
	return nil, nil
}
func opOr(pc *uint64, env *EVM, contract *Contract, memory *Memory, stack *Stack) ([]byte, error) {
	x, y := stack.pop(), stack.pop()
	stack.push(x.Or(x, y))
	return nil, nil
}
func opXor(pc *uint64, env *EVM, contract *Contract, memory *Memory, stack *Stack) ([]byte, error) {
	x, y := stack.pop(), stack.pop()
	stack.push(x.Xor(x, y))
	return nil, nil
}
func opByte(pc *uint64, env *EVM, contract *Contract, memory *Memory, stack *Stack) ([]byte, error) {
	th, val := stack.pop(), stack.pop()
>>>>>>> 7fdd714... gdbix-update v1.5.0
	if th.Cmp(big.NewInt(32)) < 0 {
		byte := big.NewInt(int64(common.LeftPadBytes(val.Bytes(), 32)[th.Int64()]))
		Stack.push(byte)
	} else {
		Stack.push(new(big.Int))
	}
	return nil, nil
}
<<<<<<< HEAD
func opAddmod(instr instruction, pc *uint64, env Environment, contract *Contract, memory *Memory, Stack *Stack) {
	x, y, z := Stack.pop(), Stack.pop(), Stack.pop()
=======
func opAddmod(pc *uint64, env *EVM, contract *Contract, memory *Memory, stack *Stack) ([]byte, error) {
	x, y, z := stack.pop(), stack.pop(), stack.pop()
>>>>>>> 7fdd714... gdbix-update v1.5.0
	if z.Cmp(Zero) > 0 {
		add := x.Add(x, y)
		add.Mod(add, z)
		Stack.push(U256(add))
	} else {
		Stack.push(new(big.Int))
	}
	return nil, nil
}
<<<<<<< HEAD
func opMulmod(instr instruction, pc *uint64, env Environment, contract *Contract, memory *Memory, Stack *Stack) {
	x, y, z := Stack.pop(), Stack.pop(), Stack.pop()
=======
func opMulmod(pc *uint64, env *EVM, contract *Contract, memory *Memory, stack *Stack) ([]byte, error) {
	x, y, z := stack.pop(), stack.pop(), stack.pop()
>>>>>>> 7fdd714... gdbix-update v1.5.0
	if z.Cmp(Zero) > 0 {
		mul := x.Mul(x, y)
		mul.Mod(mul, z)
		Stack.push(U256(mul))
	} else {
		Stack.push(new(big.Int))
	}
	return nil, nil
}

<<<<<<< HEAD
func opSha3(instr instruction, pc *uint64, env Environment, contract *Contract, memory *Memory, Stack *Stack) {
	offset, size := Stack.pop(), Stack.pop()
	hash := crypto.Keccak256(memory.Get(offset.Int64(), size.Int64()))

	Stack.push(common.BytesToBig(hash))
}

func opAddress(instr instruction, pc *uint64, env Environment, contract *Contract, memory *Memory, Stack *Stack) {
	Stack.push(common.Bytes2Big(contract.Address().Bytes()))
}

func opBalance(instr instruction, pc *uint64, env Environment, contract *Contract, memory *Memory, Stack *Stack) {
	addr := common.BigToAddress(Stack.pop())
	balance := env.Db().GetBalance(addr)

	Stack.push(new(big.Int).Set(balance))
}

func opOrigin(instr instruction, pc *uint64, env Environment, contract *Contract, memory *Memory, Stack *Stack) {
	Stack.push(env.Origin().Big())
}

func opCaller(instr instruction, pc *uint64, env Environment, contract *Contract, memory *Memory, Stack *Stack) {
	Stack.push(contract.Caller().Big())
}

func opCallValue(instr instruction, pc *uint64, env Environment, contract *Contract, memory *Memory, Stack *Stack) {
	Stack.push(new(big.Int).Set(contract.value))
}

func opCalldataLoad(instr instruction, pc *uint64, env Environment, contract *Contract, memory *Memory, Stack *Stack) {
	Stack.push(common.Bytes2Big(getData(contract.Input, Stack.pop(), common.Big32)))
}

func opCalldataSize(instr instruction, pc *uint64, env Environment, contract *Contract, memory *Memory, Stack *Stack) {
	Stack.push(big.NewInt(int64(len(contract.Input))))
}

func opCalldataCopy(instr instruction, pc *uint64, env Environment, contract *Contract, memory *Memory, Stack *Stack) {
=======
func opSha3(pc *uint64, env *EVM, contract *Contract, memory *Memory, stack *Stack) ([]byte, error) {
	offset, size := stack.pop(), stack.pop()
	data := memory.Get(offset.Int64(), size.Int64())
	hash := crypto.Keccak256(data)

	if env.vmConfig.EnablePreimageRecording {
		env.StateDB.AddPreimage(common.BytesToHash(hash), data)
	}

	stack.push(common.BytesToBig(hash))
	return nil, nil
}

func opAddress(pc *uint64, env *EVM, contract *Contract, memory *Memory, stack *Stack) ([]byte, error) {
	stack.push(common.Bytes2Big(contract.Address().Bytes()))
	return nil, nil
}

func opBalance(pc *uint64, env *EVM, contract *Contract, memory *Memory, stack *Stack) ([]byte, error) {
	addr := common.BigToAddress(stack.pop())
	balance := env.StateDB.GetBalance(addr)

	stack.push(new(big.Int).Set(balance))
	return nil, nil
}

func opOrigin(pc *uint64, env *EVM, contract *Contract, memory *Memory, stack *Stack) ([]byte, error) {
	stack.push(env.Origin.Big())
	return nil, nil
}

func opCaller(pc *uint64, env *EVM, contract *Contract, memory *Memory, stack *Stack) ([]byte, error) {
	stack.push(contract.Caller().Big())
	return nil, nil
}

func opCallValue(pc *uint64, env *EVM, contract *Contract, memory *Memory, stack *Stack) ([]byte, error) {
	stack.push(new(big.Int).Set(contract.value))
	return nil, nil
}

func opCalldataLoad(pc *uint64, env *EVM, contract *Contract, memory *Memory, stack *Stack) ([]byte, error) {
	stack.push(common.Bytes2Big(getData(contract.Input, stack.pop(), common.Big32)))
	return nil, nil
}

func opCalldataSize(pc *uint64, env *EVM, contract *Contract, memory *Memory, stack *Stack) ([]byte, error) {
	stack.push(big.NewInt(int64(len(contract.Input))))
	return nil, nil
}

func opCalldataCopy(pc *uint64, env *EVM, contract *Contract, memory *Memory, stack *Stack) ([]byte, error) {
>>>>>>> 7fdd714... gdbix-update v1.5.0
	var (
		mOff = Stack.pop()
		cOff = Stack.pop()
		l    = Stack.pop()
	)
	memory.Set(mOff.Uint64(), l.Uint64(), getData(contract.Input, cOff, l))
	return nil, nil
}

<<<<<<< HEAD
func opExtCodeSize(instr instruction, pc *uint64, env Environment, contract *Contract, memory *Memory, Stack *Stack) {
	addr := common.BigToAddress(Stack.pop())
	l := big.NewInt(int64(env.Db().GetCodeSize(addr)))
	Stack.push(l)
}

func opCodeSize(instr instruction, pc *uint64, env Environment, contract *Contract, memory *Memory, Stack *Stack) {
	l := big.NewInt(int64(len(contract.Code)))
	Stack.push(l)
}

func opCodeCopy(instr instruction, pc *uint64, env Environment, contract *Contract, memory *Memory, Stack *Stack) {
=======
func opExtCodeSize(pc *uint64, env *EVM, contract *Contract, memory *Memory, stack *Stack) ([]byte, error) {
	addr := common.BigToAddress(stack.pop())
	l := big.NewInt(int64(env.StateDB.GetCodeSize(addr)))
	stack.push(l)
	return nil, nil
}

func opCodeSize(pc *uint64, env *EVM, contract *Contract, memory *Memory, stack *Stack) ([]byte, error) {
	l := big.NewInt(int64(len(contract.Code)))
	stack.push(l)
	return nil, nil
}

func opCodeCopy(pc *uint64, env *EVM, contract *Contract, memory *Memory, stack *Stack) ([]byte, error) {
>>>>>>> 7fdd714... gdbix-update v1.5.0
	var (
		mOff = Stack.pop()
		cOff = Stack.pop()
		l    = Stack.pop()
	)
	codeCopy := getData(contract.Code, cOff, l)

	memory.Set(mOff.Uint64(), l.Uint64(), codeCopy)
	return nil, nil
}

<<<<<<< HEAD
func opExtCodeCopy(instr instruction, pc *uint64, env Environment, contract *Contract, memory *Memory, Stack *Stack) {
=======
func opExtCodeCopy(pc *uint64, env *EVM, contract *Contract, memory *Memory, stack *Stack) ([]byte, error) {
>>>>>>> 7fdd714... gdbix-update v1.5.0
	var (
		addr = common.BigToAddress(Stack.pop())
		mOff = Stack.pop()
		cOff = Stack.pop()
		l    = Stack.pop()
	)
	codeCopy := getData(env.StateDB.GetCode(addr), cOff, l)

	memory.Set(mOff.Uint64(), l.Uint64(), codeCopy)
	return nil, nil
}

<<<<<<< HEAD
func opGasprice(instr instruction, pc *uint64, env Environment, contract *Contract, memory *Memory, Stack *Stack) {
	Stack.push(new(big.Int).Set(contract.Price))
}

func opBlockhash(instr instruction, pc *uint64, env Environment, contract *Contract, memory *Memory, Stack *Stack) {
	num := Stack.pop()

	n := new(big.Int).Sub(env.BlockNumber(), common.Big257)
	if num.Cmp(n) > 0 && num.Cmp(env.BlockNumber()) < 0 {
		Stack.push(env.GetHash(num.Uint64()).Big())
=======
func opGasprice(pc *uint64, env *EVM, contract *Contract, memory *Memory, stack *Stack) ([]byte, error) {
	stack.push(new(big.Int).Set(env.GasPrice))
	return nil, nil
}

func opBlockhash(pc *uint64, env *EVM, contract *Contract, memory *Memory, stack *Stack) ([]byte, error) {
	num := stack.pop()

	n := new(big.Int).Sub(env.BlockNumber, common.Big257)
	if num.Cmp(n) > 0 && num.Cmp(env.BlockNumber) < 0 {
		stack.push(env.GetHash(num.Uint64()).Big())
>>>>>>> 7fdd714... gdbix-update v1.5.0
	} else {
		Stack.push(new(big.Int))
	}
	return nil, nil
}

<<<<<<< HEAD
func opCoinbase(instr instruction, pc *uint64, env Environment, contract *Contract, memory *Memory, Stack *Stack) {
	Stack.push(env.Coinbase().Big())
}

func opTimestamp(instr instruction, pc *uint64, env Environment, contract *Contract, memory *Memory, Stack *Stack) {
	Stack.push(U256(new(big.Int).Set(env.Time())))
}

func opNumber(instr instruction, pc *uint64, env Environment, contract *Contract, memory *Memory, Stack *Stack) {
	Stack.push(U256(new(big.Int).Set(env.BlockNumber())))
}

func opDifficulty(instr instruction, pc *uint64, env Environment, contract *Contract, memory *Memory, Stack *Stack) {
	Stack.push(U256(new(big.Int).Set(env.Difficulty())))
}

func opGasLimit(instr instruction, pc *uint64, env Environment, contract *Contract, memory *Memory, Stack *Stack) {
	Stack.push(U256(new(big.Int).Set(env.GasLimit())))
}

func opPop(instr instruction, pc *uint64, env Environment, contract *Contract, memory *Memory, Stack *Stack) {
	Stack.pop()
}

func opPush(instr instruction, pc *uint64, env Environment, contract *Contract, memory *Memory, Stack *Stack) {
	Stack.push(new(big.Int).Set(instr.data))
}

func opDup(instr instruction, pc *uint64, env Environment, contract *Contract, memory *Memory, Stack *Stack) {
	Stack.dup(int(instr.data.Int64()))
}

func opSwap(instr instruction, pc *uint64, env Environment, contract *Contract, memory *Memory, Stack *Stack) {
	Stack.swap(int(instr.data.Int64()))
}

func opLog(instr instruction, pc *uint64, env Environment, contract *Contract, memory *Memory, Stack *Stack) {
	n := int(instr.data.Int64())
	topics := make([]common.Hash, n)
	mStart, mSize := Stack.pop(), Stack.pop()
	for i := 0; i < n; i++ {
		topics[i] = common.BigToHash(Stack.pop())
	}

	d := memory.Get(mStart.Int64(), mSize.Int64())
	log := NewLog(contract.Address(), topics, d, env.BlockNumber().Uint64())
	env.AddLog(log)
}

func opMload(instr instruction, pc *uint64, env Environment, contract *Contract, memory *Memory, Stack *Stack) {
	offset := Stack.pop()
	val := common.BigD(memory.Get(offset.Int64(), 32))
	Stack.push(val)
}

func opMstore(instr instruction, pc *uint64, env Environment, contract *Contract, memory *Memory, Stack *Stack) {
	// pop value of the Stack
	mStart, val := Stack.pop(), Stack.pop()
=======
func opCoinbase(pc *uint64, env *EVM, contract *Contract, memory *Memory, stack *Stack) ([]byte, error) {
	stack.push(env.Coinbase.Big())
	return nil, nil
}

func opTimestamp(pc *uint64, env *EVM, contract *Contract, memory *Memory, stack *Stack) ([]byte, error) {
	stack.push(U256(new(big.Int).Set(env.Time)))
	return nil, nil
}

func opNumber(pc *uint64, env *EVM, contract *Contract, memory *Memory, stack *Stack) ([]byte, error) {
	stack.push(U256(new(big.Int).Set(env.BlockNumber)))
	return nil, nil
}

func opDifficulty(pc *uint64, env *EVM, contract *Contract, memory *Memory, stack *Stack) ([]byte, error) {
	stack.push(U256(new(big.Int).Set(env.Difficulty)))
	return nil, nil
}

func opGasLimit(pc *uint64, env *EVM, contract *Contract, memory *Memory, stack *Stack) ([]byte, error) {
	stack.push(U256(new(big.Int).Set(env.GasLimit)))
	return nil, nil
}

func opPop(pc *uint64, env *EVM, contract *Contract, memory *Memory, stack *Stack) ([]byte, error) {
	stack.pop()
	return nil, nil
}

func opMload(pc *uint64, env *EVM, contract *Contract, memory *Memory, stack *Stack) ([]byte, error) {
	offset := stack.pop()
	val := common.BigD(memory.Get(offset.Int64(), 32))
	stack.push(val)
	return nil, nil
}

func opMstore(pc *uint64, env *EVM, contract *Contract, memory *Memory, stack *Stack) ([]byte, error) {
	// pop value of the stack
	mStart, val := stack.pop(), stack.pop()
>>>>>>> 7fdd714... gdbix-update v1.5.0
	memory.Set(mStart.Uint64(), 32, common.BigToBytes(val, 256))
	return nil, nil
}

<<<<<<< HEAD
func opMstore8(instr instruction, pc *uint64, env Environment, contract *Contract, memory *Memory, Stack *Stack) {
	off, val := Stack.pop().Int64(), Stack.pop().Int64()
=======
func opMstore8(pc *uint64, env *EVM, contract *Contract, memory *Memory, stack *Stack) ([]byte, error) {
	off, val := stack.pop().Int64(), stack.pop().Int64()
>>>>>>> 7fdd714... gdbix-update v1.5.0
	memory.store[off] = byte(val & 0xff)
	return nil, nil
}

<<<<<<< HEAD
func opSload(instr instruction, pc *uint64, env Environment, contract *Contract, memory *Memory, Stack *Stack) {
	loc := common.BigToHash(Stack.pop())
	val := env.Db().GetState(contract.Address(), loc).Big()
	Stack.push(val)
}

func opSstore(instr instruction, pc *uint64, env Environment, contract *Contract, memory *Memory, Stack *Stack) {
	loc := common.BigToHash(Stack.pop())
	val := Stack.pop()
	env.Db().SetState(contract.Address(), loc, common.BigToHash(val))
}

func opJump(instr instruction, pc *uint64, env Environment, contract *Contract, memory *Memory, Stack *Stack) {
}
func opJumpi(instr instruction, pc *uint64, env Environment, contract *Contract, memory *Memory, Stack *Stack) {
}
func opJumpdest(instr instruction, pc *uint64, env Environment, contract *Contract, memory *Memory, Stack *Stack) {
}

func opPc(instr instruction, pc *uint64, env Environment, contract *Contract, memory *Memory, Stack *Stack) {
	Stack.push(new(big.Int).Set(instr.data))
}

func opMsize(instr instruction, pc *uint64, env Environment, contract *Contract, memory *Memory, Stack *Stack) {
	Stack.push(big.NewInt(int64(memory.Len())))
}

func opGas(instr instruction, pc *uint64, env Environment, contract *Contract, memory *Memory, Stack *Stack) {
	Stack.push(new(big.Int).Set(contract.Gas))
}

func opCreate(instr instruction, pc *uint64, env Environment, contract *Contract, memory *Memory, Stack *Stack) {
=======
func opSload(pc *uint64, env *EVM, contract *Contract, memory *Memory, stack *Stack) ([]byte, error) {
	loc := common.BigToHash(stack.pop())
	val := env.StateDB.GetState(contract.Address(), loc).Big()
	stack.push(val)
	return nil, nil
}

func opSstore(pc *uint64, env *EVM, contract *Contract, memory *Memory, stack *Stack) ([]byte, error) {
	loc := common.BigToHash(stack.pop())
	val := stack.pop()
	env.StateDB.SetState(contract.Address(), loc, common.BigToHash(val))
	return nil, nil
}

func opJump(pc *uint64, env *EVM, contract *Contract, memory *Memory, stack *Stack) ([]byte, error) {
	pos := stack.pop()
	if !contract.jumpdests.has(contract.CodeHash, contract.Code, pos) {
		nop := contract.GetOp(pos.Uint64())
		return nil, fmt.Errorf("invalid jump destination (%v) %v", nop, pos)
	}
	*pc = pos.Uint64()
	return nil, nil
}
func opJumpi(pc *uint64, env *EVM, contract *Contract, memory *Memory, stack *Stack) ([]byte, error) {
	pos, cond := stack.pop(), stack.pop()
	if cond.Cmp(common.BigTrue) >= 0 {
		if !contract.jumpdests.has(contract.CodeHash, contract.Code, pos) {
			nop := contract.GetOp(pos.Uint64())
			return nil, fmt.Errorf("invalid jump destination (%v) %v", nop, pos)
		}
		*pc = pos.Uint64()
	} else {
		*pc++
	}
	return nil, nil
}
func opJumpdest(pc *uint64, env *EVM, contract *Contract, memory *Memory, stack *Stack) ([]byte, error) {
	return nil, nil
}

func opPc(pc *uint64, env *EVM, contract *Contract, memory *Memory, stack *Stack) ([]byte, error) {
	stack.push(new(big.Int).SetUint64(*pc))
	return nil, nil
}

func opMsize(pc *uint64, env *EVM, contract *Contract, memory *Memory, stack *Stack) ([]byte, error) {
	stack.push(big.NewInt(int64(memory.Len())))
	return nil, nil
}

func opGas(pc *uint64, env *EVM, contract *Contract, memory *Memory, stack *Stack) ([]byte, error) {
	stack.push(new(big.Int).Set(contract.Gas))
	return nil, nil
}

func opCreate(pc *uint64, env *EVM, contract *Contract, memory *Memory, stack *Stack) ([]byte, error) {
>>>>>>> 7fdd714... gdbix-update v1.5.0
	var (
		value        = Stack.pop()
		offset, size = Stack.pop(), Stack.pop()
		input        = memory.Get(offset.Int64(), size.Int64())
		gas          = new(big.Int).Set(contract.Gas)
	)
<<<<<<< HEAD
	contract.UseGas(contract.Gas)
	_, addr, suberr := env.Create(contract, input, gas, contract.Price, value)
	// Push item on the Stack based on the returned error. If the ruleset is
	// homestead we must check for CodeStoreOutOfGasError (homestead only
	// rule) and treat as an error, if the ruleset is frontier we must
	// ignore this error and pretend the operation was successful.
	if env.RuleSet().IsHomestead(env.BlockNumber()) && suberr == CodeStoreOutOfGasError {
		Stack.push(new(big.Int))
	} else if suberr != nil && suberr != CodeStoreOutOfGasError {
		Stack.push(new(big.Int))
=======
	if env.ChainConfig().IsEIP150(env.BlockNumber) {
		gas.Div(gas, n64)
		gas = gas.Sub(contract.Gas, gas)
	}

	contract.UseGas(gas)
	_, addr, suberr := env.Create(contract, input, gas, value)
	// Push item on the stack based on the returned error. If the ruleset is
	// homestead we must check for CodeStoreOutOfGasError (homestead only
	// rule) and treat as an error, if the ruleset is frontier we must
	// ignore this error and pretend the operation was successful.
	if env.ChainConfig().IsHomestead(env.BlockNumber) && suberr == ErrCodeStoreOutOfGas {
		stack.push(new(big.Int))
	} else if suberr != nil && suberr != ErrCodeStoreOutOfGas {
		stack.push(new(big.Int))
>>>>>>> 7fdd714... gdbix-update v1.5.0
	} else {
		Stack.push(addr.Big())
	}
	return nil, nil
}

<<<<<<< HEAD
func opCall(instr instruction, pc *uint64, env Environment, contract *Contract, memory *Memory, Stack *Stack) {
	gas := Stack.pop()
	// pop gas and value of the Stack.
	addr, value := Stack.pop(), Stack.pop()
=======
func opCall(pc *uint64, env *EVM, contract *Contract, memory *Memory, stack *Stack) ([]byte, error) {
	gas := stack.pop()
	// pop gas and value of the stack.
	addr, value := stack.pop(), stack.pop()
>>>>>>> 7fdd714... gdbix-update v1.5.0
	value = U256(value)
	// pop input size and offset
	inOffset, inSize := Stack.pop(), Stack.pop()
	// pop return size and offset
	retOffset, retSize := Stack.pop(), Stack.pop()

	address := common.BigToAddress(addr)

	// Get the arguments from the memory
	args := memory.Get(inOffset.Int64(), inSize.Int64())

	if len(value.Bytes()) > 0 {
		gas.Add(gas, params.CallStipend)
	}

	ret, err := env.Call(contract, address, args, gas, value)

	if err != nil {
		Stack.push(new(big.Int))

	} else {
		Stack.push(big.NewInt(1))

		memory.Set(retOffset.Uint64(), retSize.Uint64(), ret)
	}
	return nil, nil
}

<<<<<<< HEAD
func opCallCode(instr instruction, pc *uint64, env Environment, contract *Contract, memory *Memory, Stack *Stack) {
	gas := Stack.pop()
	// pop gas and value of the Stack.
	addr, value := Stack.pop(), Stack.pop()
=======
func opCallCode(pc *uint64, env *EVM, contract *Contract, memory *Memory, stack *Stack) ([]byte, error) {
	gas := stack.pop()
	// pop gas and value of the stack.
	addr, value := stack.pop(), stack.pop()
>>>>>>> 7fdd714... gdbix-update v1.5.0
	value = U256(value)
	// pop input size and offset
	inOffset, inSize := Stack.pop(), Stack.pop()
	// pop return size and offset
	retOffset, retSize := Stack.pop(), Stack.pop()

	address := common.BigToAddress(addr)

	// Get the arguments from the memory
	args := memory.Get(inOffset.Int64(), inSize.Int64())

	if len(value.Bytes()) > 0 {
		gas.Add(gas, params.CallStipend)
	}

	ret, err := env.CallCode(contract, address, args, gas, value)

	if err != nil {
		Stack.push(new(big.Int))

	} else {
		Stack.push(big.NewInt(1))

		memory.Set(retOffset.Uint64(), retSize.Uint64(), ret)
	}
	return nil, nil
}

<<<<<<< HEAD
func opDelegateCall(instr instruction, pc *uint64, env Environment, contract *Contract, memory *Memory, Stack *Stack) {
	gas, to, inOffset, inSize, outOffset, outSize := Stack.pop(), Stack.pop(), Stack.pop(), Stack.pop(), Stack.pop(), Stack.pop()
=======
func opDelegateCall(pc *uint64, env *EVM, contract *Contract, memory *Memory, stack *Stack) ([]byte, error) {
	// if not homestead return an error. DELEGATECALL is not supported
	// during pre-homestead.
	if !env.ChainConfig().IsHomestead(env.BlockNumber) {
		return nil, fmt.Errorf("invalid opcode %x", DELEGATECALL)
	}

	gas, to, inOffset, inSize, outOffset, outSize := stack.pop(), stack.pop(), stack.pop(), stack.pop(), stack.pop(), stack.pop()
>>>>>>> 7fdd714... gdbix-update v1.5.0

	toAddr := common.BigToAddress(to)
	args := memory.Get(inOffset.Int64(), inSize.Int64())
	ret, err := env.DelegateCall(contract, toAddr, args, gas)
	if err != nil {
		Stack.push(new(big.Int))
	} else {
		Stack.push(big.NewInt(1))
		memory.Set(outOffset.Uint64(), outSize.Uint64(), ret)
	}
	return nil, nil
}

<<<<<<< HEAD
func opReturn(instr instruction, pc *uint64, env Environment, contract *Contract, memory *Memory, Stack *Stack) {
}
func opStop(instr instruction, pc *uint64, env Environment, contract *Contract, memory *Memory, Stack *Stack) {
}

func opSuicide(instr instruction, pc *uint64, env Environment, contract *Contract, memory *Memory, Stack *Stack) {
	balance := env.Db().GetBalance(contract.Address())
	env.Db().AddBalance(common.BigToAddress(Stack.pop()), balance)
=======
func opReturn(pc *uint64, env *EVM, contract *Contract, memory *Memory, stack *Stack) ([]byte, error) {
	offset, size := stack.pop(), stack.pop()
	ret := memory.GetPtr(offset.Int64(), size.Int64())

	return ret, nil
}

func opStop(pc *uint64, env *EVM, contract *Contract, memory *Memory, stack *Stack) ([]byte, error) {
	return nil, nil
}

func opSuicide(pc *uint64, env *EVM, contract *Contract, memory *Memory, stack *Stack) ([]byte, error) {
	balance := env.StateDB.GetBalance(contract.Address())
	env.StateDB.AddBalance(common.BigToAddress(stack.pop()), balance)

	env.StateDB.Suicide(contract.Address())
>>>>>>> 7fdd714... gdbix-update v1.5.0

	return nil, nil
}

// following functions are used by the instruction jump  table

// make log instruction function
<<<<<<< HEAD
func makeLog(size int) instrFn {
	return func(instr instruction, pc *uint64, env Environment, contract *Contract, memory *Memory, Stack *Stack) {
=======
func makeLog(size int) executionFunc {
	return func(pc *uint64, env *EVM, contract *Contract, memory *Memory, stack *Stack) ([]byte, error) {
>>>>>>> 7fdd714... gdbix-update v1.5.0
		topics := make([]common.Hash, size)
		mStart, mSize := Stack.pop(), Stack.pop()
		for i := 0; i < size; i++ {
			topics[i] = common.BigToHash(Stack.pop())
		}

		d := memory.Get(mStart.Int64(), mSize.Int64())
		env.StateDB.AddLog(&types.Log{
			Address: contract.Address(),
			Topics:  topics,
			Data:    d,
			// This is a non-consensus field, but assigned here because
			// core/state doesn't know the current block number.
			BlockNumber: env.BlockNumber.Uint64(),
		})
		return nil, nil
	}
}

// make push instruction function
<<<<<<< HEAD
func makePush(size uint64, bsize *big.Int) instrFn {
	return func(instr instruction, pc *uint64, env Environment, contract *Contract, memory *Memory, Stack *Stack) {
=======
func makePush(size uint64, bsize *big.Int) executionFunc {
	return func(pc *uint64, env *EVM, contract *Contract, memory *Memory, stack *Stack) ([]byte, error) {
>>>>>>> 7fdd714... gdbix-update v1.5.0
		byts := getData(contract.Code, new(big.Int).SetUint64(*pc+1), bsize)
		Stack.push(common.Bytes2Big(byts))
		*pc += size
		return nil, nil
	}
}

// make push instruction function
<<<<<<< HEAD
func makeDup(size int64) instrFn {
	return func(instr instruction, pc *uint64, env Environment, contract *Contract, memory *Memory, Stack *Stack) {
		Stack.dup(int(size))
=======
func makeDup(size int64) executionFunc {
	return func(pc *uint64, env *EVM, contract *Contract, memory *Memory, stack *Stack) ([]byte, error) {
		stack.dup(int(size))
		return nil, nil
>>>>>>> 7fdd714... gdbix-update v1.5.0
	}
}

// make swap instruction function
func makeSwap(size int64) executionFunc {
	// switch n + 1 otherwise n would be swapped with n
	size += 1
<<<<<<< HEAD
	return func(instr instruction, pc *uint64, env Environment, contract *Contract, memory *Memory, Stack *Stack) {
		Stack.swap(int(size))
=======
	return func(pc *uint64, env *EVM, contract *Contract, memory *Memory, stack *Stack) ([]byte, error) {
		stack.swap(int(size))
		return nil, nil
>>>>>>> 7fdd714... gdbix-update v1.5.0
	}
}
