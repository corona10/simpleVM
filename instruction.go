package simpleVM

type InstKind int

type Instruction struct {
	kind InstKind
	name string
}

const (
	_ InstKind = iota
	IADD
	ISUB
	IMUL
	ILT
	IEQ
	BR
	BRT
	BRF
	ICONST
	LOAD
	GLOAD
	STORE
	GSTORE
	PRINT
	POP
	CALL
	RET
	HALT
)

// NewInstruction is a function for creating the Instruction
func NewInstruction(raw_inst int) Instruction {
	k := InstKind(raw_inst)
	inst := Instruction{kind: k}
	return inst
}

// GetKind is a method for getting a kind of instruction.
func (inst *Instruction) GetKind() InstKind {
	return inst.kind
}
