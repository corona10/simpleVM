package main

import "github.com/corona10/simpleVM"

func main() {
	vm := simpleVM.CreateNewVM()
	program := []int{
		int(simpleVM.ICONST),
		3,
		int(simpleVM.ICONST),
		5,
		int(simpleVM.IADD),
		int(simpleVM.ICONST),
		4,
		int(simpleVM.ISUB),
		int(simpleVM.ICONST),
		9,
		int(simpleVM.IMUL),
		int(simpleVM.PRINT),
	}
	vm.LoadProgram(program, 0)
	vm.Run()
}
