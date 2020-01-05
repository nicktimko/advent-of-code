package intcode

import (
	"fmt"
)

// ParameterMode ...
type ParameterMode int

const (
	modeIndirect  ParameterMode = 0
	modeImmediate ParameterMode = 1
)

// Status indicates the status of the intcode processor.
type Status int

const (
	// Running means the machine has not halted
	Running Status = iota
	// Halted means the machine ended the program normally with a HALT instruction
	Halted
	// Crashed means the machine encountered an error, and has halted abnormally
	Crashed
)

// OpCode ...
type OpCode struct {
	op int
	pm [3]ParameterMode
}

type intcodeComputerState struct {
	ptr         int
	tape        []int
	inputs      []int
	outputs     []int
	status      Status
	faultReason string
}

func decodeOp(opcode int) OpCode {
	var oc OpCode

	oc.op = opcode % 100
	opcode = opcode / 100

	for i := 0; i < 3; i++ {
		oc.pm[i] = ParameterMode(opcode % 10)
		opcode = opcode / 10
	}

	return oc
}

func getParam(c *intcodeComputerState, n int, modes [3]ParameterMode) (int, bool) {
	switch modes[n] {
	case modeImmediate:
		return c.tape[c.ptr+n+1], false
	case modeIndirect:
		return c.tape[c.tape[c.ptr+n+1]], false
	}
	c.faultReason = fmt.Sprintf("error in parameter %d at instruction %d", n, c.ptr)
	c.status = Crashed
	return 0, true
}

// ProcessInstruction (single) for Intcode tapes
func ProcessInstruction(c *intcodeComputerState) {
	op := decodeOp(c.tape[c.ptr])
	icOps[op.op](c, op.pm)
}

// Processor for simple Intcode tapes with no I/O
func Processor(tape []int) {
	IOProcessor(tape, []int{})
}

// IOProcessor supports Intcode tapes with input/output
func IOProcessor(tape []int, inputs []int) []int {
	var c intcodeComputerState

	c.tape = tape
	c.ptr = 0
	c.status = Running
	c.inputs = inputs

	for c.status == Running {
		ProcessInstruction(&c)
	}

	return c.outputs
}
