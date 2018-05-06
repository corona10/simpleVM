package simplevm

import (
	"testing"
)

func TestVMCase1(t *testing.T) {
	vm := CreateNewVM()
	program := []int{
		int(ICONST),
		3,
		int(ICONST),
		5,
		int(IADD),
		int(ICONST),
		4,
		int(ISUB),
		int(ICONST),
		9,
		int(IMUL),
		int(PRINT),
	}
	vm.LoadProgram(program, 0, 0)
	vm.Run()
}

func TestVMCase2(t *testing.T) {
	vm := CreateNewVM()
	program := []int{
		int(ICONST),
		3,
		int(ICONST),
		5,
		int(HALT),
		int(IADD),
		int(ICONST),
		4,
		int(ISUB),
		int(ICONST),
		9,
		int(IMUL),
		int(PRINT),
	}
	vm.LoadProgram(program, 0, 0)
	vm.Run()
}

func TestVMCase3(t *testing.T) {
	vm := CreateNewVM()
	factorial := []int{
		int(LOAD),
		-3,
		int(ICONST),
		2,
		int(ILT),
		int(BRF),
		10,
		int(ICONST),
		1,
		int(RET),
		int(LOAD),
		-3,
		int(LOAD),
		-3,
		int(ICONST),
		1,
		int(ISUB),
		int(CALL),
		0,
		1,
		int(IMUL),
		int(RET),
		int(ICONST),
		5,
		int(CALL),
		0,
		1,
		int(PRINT),
		int(HALT),
	}
	vm.LoadProgram(factorial, 22, 0)
	vm.Run()
}
