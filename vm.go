package simplevm

import (
	"fmt"
	"log"
	"os"
)

type VM struct {
	isRunning  bool
	program    []int
	stack      []int
	locals     []int
	globals    []int
	fCallstack []*Context
	pc         int // program counter
	callsp     int
	sp         int // stack pointer
	fp         int // frame pointer
	logger     *log.Logger
}

const (
	DEFAULT_CALL_STACK_SIZE = 1000
	DEFAULT_STACK_SIZE      = 1000
	TRUE                    = 1
	FALSE                   = 0
)

// CreateNewVM is a function for create a new virtual machine.
func CreateNewVM() *VM {
	return &VM{isRunning: false, program: nil,
		pc: 0, sp: -1, fp: -1, callsp: -1, logger: log.New(os.Stdout, "simpleVM: ", log.LstdFlags)}
}

// LoadProgram is a method to load a program
// into the virtual machine.
func (vm *VM) LoadProgram(program []int, startPc int, nGlobals int) {
	vm.program = program
	vm.pc = startPc
	vm.stack = make([]int, DEFAULT_STACK_SIZE)
	vm.fCallstack = make([]*Context, DEFAULT_CALL_STACK_SIZE)
	vm.globals = make([]int, nGlobals)
}

// Run is a method to run a program on
// the virtual machine.
func (vm *VM) Run() {
	vm.logger.Println("Now VM launch a program")
	vm.isRunning = true
	for vm.isRunning && vm.pc < len(vm.program) {
		rawInst := vm.fetch()
		inst := vm.decode(rawInst)
		vm.execute(inst)
	}
	vm.logger.Println("End of the program for the VM.")
}

func (vm *VM) decode(rawInst int) *Instruction {
	inst := NewInstruction(rawInst)
	return inst
}

func (vm *VM) execute(inst *Instruction) {
	switch in := inst.GetKind(); in {
	case BR:
		vm.pc = vm.program[vm.pc]
	case BRT:
		addr := vm.program[vm.pc]
		vm.pc++
		if vm.stack[vm.sp] == TRUE {
			vm.pc = addr
		}
		vm.sp--
	case BRF:
		addr := vm.program[vm.pc]
		vm.pc++
		if vm.stack[vm.sp] == FALSE {
			vm.pc = addr
		}
		vm.sp--
	case CALL:
		address := vm.program[vm.pc] // Get address function
		vm.pc++
		nargs := vm.program[vm.pc] // Get number of arguments
		vm.pc++
		vm.callsp++
		vm.fCallstack[vm.callsp] = NewContext(vm.pc, nargs)
		firstarg := vm.sp - nargs + 1
		for i := 0; i < nargs; i++ {
			vm.fCallstack[vm.callsp].locals[i] = vm.stack[firstarg+i]
		}
		vm.sp -= nargs
		vm.pc = address
	case RET:
		vm.pc = vm.fCallstack[vm.callsp].GetReturnPC()
		vm.callsp--
	case IADD:
		b := vm.stack[vm.sp]
		vm.sp--
		a := vm.stack[vm.sp]
		vm.stack[vm.sp] = a + b
	case ICONST:
		vm.sp++
		vm.stack[vm.sp] = vm.program[vm.pc]
		vm.pc++
	case IMUL:
		b := vm.stack[vm.sp]
		vm.sp--
		a := vm.stack[vm.sp]
		vm.stack[vm.sp] = a * b
	case IDIV:
		b := vm.stack[vm.sp]
		vm.sp--
		a := vm.stack[vm.sp]
		vm.stack[vm.sp] = a / b
	case IMOD:
		b := vm.stack[vm.sp]
		vm.sp--
		a := vm.stack[vm.sp]
		vm.stack[vm.sp] = a % b
	case ILT:
		b := vm.stack[vm.sp]
		vm.sp--
		a := vm.stack[vm.sp]
		if a < b {
			vm.stack[vm.sp] = TRUE
		} else {
			vm.stack[vm.sp] = FALSE
		}
	case PRINT:
		fmt.Println(vm.stack[vm.sp])
		vm.sp--
	case POP:
		vm.sp--
	case LOAD:
		regnum := vm.program[vm.pc]
		vm.pc++
		vm.sp++
		vm.stack[vm.sp] = vm.fCallstack[vm.callsp].locals[regnum]
	case STORE:
		regnum := vm.program[vm.pc]
		vm.pc++
		vm.fCallstack[vm.callsp].locals[regnum] = vm.stack[vm.sp]
		vm.sp--
	case GLOAD:
		address := vm.program[vm.pc]
		vm.pc++
		vm.sp++
		vm.stack[vm.sp] = vm.globals[address]
	case GSTORE:
		address := vm.program[vm.pc]
		vm.pc++
		vm.globals[address] = vm.stack[vm.sp]
		vm.sp--
	case ISUB:
		b := vm.stack[vm.sp]
		vm.sp--
		a := vm.stack[vm.sp]
		vm.stack[vm.sp] = a - b
	case HALT:
		vm.logger.Printf("Got HALT instruction from pc[%v]\n", vm.pc)
		vm.isRunning = false
	default:
		msg := fmt.Sprintf("Invalid bytecode: %v pc[%v]\n", int(in), vm.pc)
		panic(msg)
	}
}

func (vm *VM) fetch() int {
	ret := vm.program[vm.pc]
	vm.pc++
	return ret
}
