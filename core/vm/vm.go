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
	"fmt"
	"math/big"
	"sync/atomic"
	"time"

<<<<<<< HEAD
	"github.com/teslafunds/go-teslafunds/common"
	"github.com/teslafunds/go-teslafunds/crypto"
	"github.com/teslafunds/go-teslafunds/logger"
	"github.com/teslafunds/go-teslafunds/logger/glog"
	"github.com/teslafunds/go-teslafunds/params"
=======
	"github.com/dubaicoin-dbix/go-dubaicoin/common"
	"github.com/dubaicoin-dbix/go-dubaicoin/crypto"
	"github.com/dubaicoin-dbix/go-dubaicoin/logger"
	"github.com/dubaicoin-dbix/go-dubaicoin/logger/glog"
	"github.com/dubaicoin-dbix/go-dubaicoin/params"
>>>>>>> 7fdd714... gdbix-update v1.5.0
)

// Config are the configuration options for the Interpreter
type Config struct {
	// Debug enabled debugging Interpreter options
	Debug bool
	// EnableJit enabled the JIT VM
	EnableJit bool
<<<<<<< HEAD
	ForceJit  bool
	Tracer    Tracer
}

// EVM is used to run Teslafunds based contracts and will utilise the
=======
	// ForceJit forces the JIT VM
	ForceJit bool
	// Tracer is the op code logger
	Tracer Tracer
	// NoRecursion disabled Interpreter call, callcode,
	// delegate call and create.
	NoRecursion bool
	// Disable gas metering
	DisableGasMetering bool
	// Enable recording of SHA3/keccak preimages
	EnablePreimageRecording bool
	// JumpTable contains the EVM instruction table. This
	// may me left uninitialised and will be set the default
	// table.
	JumpTable [256]operation
}

// Interpreter is used to run Ethereum based contracts and will utilise the
>>>>>>> 7fdd714... gdbix-update v1.5.0
// passed environment to query external sources for state information.
// The Interpreter will run the byte code VM or JIT VM based on the passed
// configuration.
<<<<<<< HEAD
type EVM struct {
	env       Environment
	jumpTable vmJumpTable
	cfg       Config


}

// New returns a new instance of the EVM.
func New(env Environment, cfg Config) *EVM {

	return &EVM{
		env:       env,
		jumpTable: newJumpTable(env.RuleSet(), env.BlockNumber()),
		cfg:       cfg,

=======
type Interpreter struct {
	env      *EVM
	cfg      Config
	gasTable params.GasTable
}

// NewInterpreter returns a new instance of the Interpreter.
func NewInterpreter(env *EVM, cfg Config) *Interpreter {
	// We use the STOP instruction whether to see
	// the jump table was initialised. If it was not
	// we'll set the default jump table.
	if !cfg.JumpTable[STOP].valid {
		cfg.JumpTable = defaultJumpTable
	}

	return &Interpreter{
		env:      env,
		cfg:      cfg,
		gasTable: env.ChainConfig().GasTable(env.BlockNumber),
>>>>>>> 7fdd714... gdbix-update v1.5.0
	}
}

// Run loops and evaluates the contract's code with the given input data
func (evm *Interpreter) Run(contract *Contract, input []byte) (ret []byte, err error) {
	evm.env.depth++
	defer func() { evm.env.depth-- }()

	if contract.CodeAddr != nil {
		if p := PrecompiledContracts[*contract.CodeAddr]; p != nil {
			return RunPrecompiledContract(p, input, contract)
		}
	}

	// Don't bother with the execution if there's no code.
	if len(contract.Code) == 0 {
		return nil, nil
	}

	codehash := contract.CodeHash // codehash is used when doing jump dest caching
	if codehash == (common.Hash{}) {
		codehash = crypto.Keccak256Hash(contract.Code)
	}

	var (
		op    OpCode        // current opcode
		mem   = NewMemory() // bound memory
		stack = newstack()  // local stack
		// For optimisation reason we're using uint64 as the program counter.
		// It's theoretically possible to go above 2^64. The YP defines the PC to be uint256. Practically much less so feasible.
		pc   = uint64(0) // program counter
		cost *big.Int
	)
	contract.Input = input

	// User defer pattern to check for an error and, based on the error being nil or not, use all gas and return.
	defer func() {
		if err != nil && evm.cfg.Debug {
<<<<<<< HEAD
			evm.cfg.Tracer.CaptureState(evm.env, pc, op, contract.Gas, cost, mem, stack, contract, evm.env.Depth(), err)
=======
			evm.cfg.Tracer.CaptureState(evm.env, pc, op, contract.Gas, cost, mem, stack, contract, evm.env.depth, err)
>>>>>>> 7fdd714... gdbix-update v1.5.0
		}
	}()

	if glog.V(logger.Debug) {
		glog.Infof("evm running: %x\n", codehash[:4])
		tstart := time.Now()
		defer func() {
			glog.Infof("evm done: %x. time: %v\n", codehash[:4], time.Since(tstart))
		}()
	}

	// The Interpreter main run loop (contextual). This loop runs until either an
	// explicit STOP, RETURN or SUICIDE is executed, an error accured during
	// the execution of one of the operations or until the evm.done is set by
	// the parent context.Context.
	for atomic.LoadInt32(&evm.env.abort) == 0 {
		// Get the memory location of pc
		op = contract.GetOp(pc)

		// get the operation from the jump table matching the opcode
		operation := evm.cfg.JumpTable[op]

<<<<<<< HEAD
		// Resize the memory calculated previously
		mem.Resize(newMemSize.Uint64())
		// Add a log message
		if evm.cfg.Debug {
			evm.cfg.Tracer.CaptureState(evm.env, pc, op, contract.Gas, cost, mem, stack, contract, evm.env.Depth(), nil)
		}

		if opPtr := evm.jumpTable[op]; opPtr.valid {
			if opPtr.fn != nil {
				opPtr.fn(instruction{}, &pc, evm.env, contract, mem, stack)
			} else {
				switch op {
				case PC:
					opPc(instruction{data: new(big.Int).SetUint64(pc)}, &pc, evm.env, contract, mem, stack)
				case JUMP:
					if err := jump(pc, stack.pop()); err != nil {
						return nil, err
					}

					continue
				case JUMPI:
					pos, cond := stack.pop(), stack.pop()

					if cond.Cmp(common.BigTrue) >= 0 {
						if err := jump(pc, pos); err != nil {
							return nil, err
						}

						continue
					}
				case RETURN:
					offset, size := stack.pop(), stack.pop()
					ret := mem.GetPtr(offset.Int64(), size.Int64())

					return ret, nil
				case SUICIDE:
					opSuicide(instruction{}, nil, evm.env, contract, mem, stack)

					fallthrough
				case STOP: // Stop the contract
					return nil, nil
				}
			}
		} else {
			return nil, fmt.Errorf("Invalid opcode %x", op)
		}

		pc++

	}
}

// calculateGasAndSize calculates the required given the opcode and stack items calculates the new memorysize for
// the operation. This does not reduce gas or resizes the memory.
func calculateGasAndSize(env Environment, contract *Contract, caller ContractRef, op OpCode, statedb Database, mem *Memory, stack *Stack) (*big.Int, *big.Int, error) {
	var (
		gas                 = new(big.Int)
		newMemSize *big.Int = new(big.Int)
	)
	err := baseCheck(op, stack, gas)
	if err != nil {
		return nil, nil, err
	}

	// stack Check, memory resize & gas phase
	switch op {
	case SWAP1, SWAP2, SWAP3, SWAP4, SWAP5, SWAP6, SWAP7, SWAP8, SWAP9, SWAP10, SWAP11, SWAP12, SWAP13, SWAP14, SWAP15, SWAP16:
		n := int(op - SWAP1 + 2)
		err := stack.require(n)
		if err != nil {
			return nil, nil, err
		}
		gas.Set(GasFastestStep)
	case DUP1, DUP2, DUP3, DUP4, DUP5, DUP6, DUP7, DUP8, DUP9, DUP10, DUP11, DUP12, DUP13, DUP14, DUP15, DUP16:
		n := int(op - DUP1 + 1)
		err := stack.require(n)
		if err != nil {
			return nil, nil, err
		}
		gas.Set(GasFastestStep)
	case LOG0, LOG1, LOG2, LOG3, LOG4:
		n := int(op - LOG0)
		err := stack.require(n + 2)
		if err != nil {
			return nil, nil, err
=======
		// if the op is invalid abort the process and return an error
		if !operation.valid {
			return nil, fmt.Errorf("invalid opcode %x", op)
		}

		// validate the stack and make sure there enough stack items available
		// to perform the operation
		if err := operation.validateStack(stack); err != nil {
			return nil, err
>>>>>>> 7fdd714... gdbix-update v1.5.0
		}

		var memorySize *big.Int
		// calculate the new memory size and expand the memory to fit
		// the operation
		if operation.memorySize != nil {
			memorySize = operation.memorySize(stack)
			// memory is expanded in words of 32 bytes. Gas
			// is also calculated in words.
			memorySize.Mul(toWordSize(memorySize), big.NewInt(32))
		}

		if !evm.cfg.DisableGasMetering {
			// consume the gas and return an error if not enough gas is available.
			// cost is explicitly set so that the capture state defer method cas get the proper cost
			cost = operation.gasCost(evm.gasTable, evm.env, contract, stack, mem, memorySize)
			if !contract.UseGas(cost) {
				return nil, ErrOutOfGas
			}
		}
		if memorySize != nil {
			mem.Resize(memorySize.Uint64())
		}

		if evm.cfg.Debug {
			evm.cfg.Tracer.CaptureState(evm.env, pc, op, contract.Gas, cost, mem, stack, contract, evm.env.depth, err)
		}

		// execute the operation
		res, err := operation.execute(&pc, evm.env, contract, mem, stack)
		switch {
		case err != nil:
			return nil, err
		case operation.halts:
			return res, nil
		case !operation.jumps:
			pc++
		}
	}
	return nil, nil
}
