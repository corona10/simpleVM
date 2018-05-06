package simpleVM

import (
	"fmt"
	"log"
	"os"
)

type VM struct {
	is_running bool
	program    []int
	stack      []int
	globals    []int
	pc         int
	sp         int
	fp         int
	logger     *log.Logger
}

const (
	DEFAULT_STACK_SIZE = 1000
)

func CreateNewVM() *VM {
	return &VM{is_running: false, program: nil,
		pc: 0, sp: -1, fp: -1, logger: log.New(os.Stdout, "vm: ", log.LstdFlags)}
}

// LoadProgram is a method to load a program
// into the virtual machine.
func (vm *VM) LoadProgram(program []int, start_pc int) {
	vm.program = program
	vm.pc = start_pc
	vm.stack = make([]int, DEFAULT_STACK_SIZE)
}

// Run is a method to run a program on
// the virtual machine.
func (vm *VM) Run() {
	vm.logger.Println("Now VM runs a code.")
	vm.is_running = true
	for vm.is_running && vm.pc < len(vm.program) {
		raw_inst := vm.fetch()
		inst := vm.decode(raw_inst)
		vm.execute(inst)
	}
	vm.logger.Println("End of the program for the VM.")
}

func (vm *VM) decode(raw_inst int) Instruction {
	inst := NewInstruction(raw_inst)
	return inst
}

func (vm *VM) execute(inst Instruction) {
	switch in := inst.GetKind(); in {
	case IADD:
		b := vm.stack[vm.sp]
		vm.sp--
		a := vm.stack[vm.sp]
		vm.sp++
		vm.stack[vm.sp] = a + b
	case ICONST:
		vm.sp++
		vm.stack[vm.sp] = vm.program[vm.pc]
		vm.pc++
	case PRINT:
		fmt.Println(vm.stack[vm.sp])
		vm.sp--
	case HALT:
		vm.is_running = false
	}
}

func (vm *VM) fetch() int {
	ret := vm.program[vm.pc]
	vm.pc++
	return ret
}
