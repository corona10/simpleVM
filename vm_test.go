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
		0,
		int(ICONST),
		2,
		int(ILT),
		int(BRF),
		10,
		int(ICONST),
		1,
		int(RET),
		int(LOAD),
		0,
		int(LOAD),
		0,
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

func TestVMCase4(t *testing.T) {
	vm := CreateNewVM()
	program := []int{
		int(ICONST),
		10,
		int(GSTORE),
		0,
		int(ICONST),
		0,
		int(GSTORE),
		1,
		int(GLOAD),
		1,
		int(GLOAD),
		0,
		int(ILT),
		int(BRF),
		24,
		int(GLOAD),
		1,
		int(ICONST),
		1,
		int(IADD),
		int(GSTORE),
		1,
		int(BR),
		8,
		int(HALT),
	}
	vm.LoadProgram(program, 0, 10)
	vm.Run()
}
