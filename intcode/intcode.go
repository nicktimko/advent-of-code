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

// State is the set of values that compose an IntCode computer's internal state.
type State struct {
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

func (c *State) getParam(n int, modes [3]ParameterMode) (int, bool) {
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
func (c *State) ProcessInstruction() {
	if c.ptr >= len(c.tape) {
		c.status = Crashed
		c.faultReason = fmt.Sprintf("pointer left tape (address %d)", c.ptr)
		return
	}
	op := decodeOp(c.tape[c.ptr])

	opHandler, ok := icOps[op.op]
	if !ok {
		c.status = Crashed
		c.faultReason = fmt.Sprintf("unknown op %d at address %d", op.op, c.ptr)
		return
	}
	opHandler(c, op.pm)
}

// Processor for simple Intcode tapes with no I/O
func Processor(tape []int) {
	IOProcessor(tape, []int{})
}

// IOProcessor supports Intcode tapes with input/output
func IOProcessor(tape []int, inputs []int) []int {
	var c State

	c.tape = tape
	c.ptr = 0
	c.status = Running
	c.inputs = inputs

	for c.status == Running {
		c.ProcessInstruction()
	}

	return c.outputs
}

// New Intcode processor
func New(tape []int, inputs []int) State {
	var c State

	c.tape = tape
	c.ptr = 0
	c.status = Running
	c.inputs = inputs

	return c
}

// CommunicatingProcessor uses input and output channels for I/O.
func CommunicatingProcessor(tape []int, input chan int, output chan int) {
	var c State

	c.tape = tape
	c.ptr = 0
	c.status = Running

	for c.status == Running {
		op := decodeOp(c.tape[c.ptr])

		if op.op == opcInput {
			c.inputs = []int{<-input}
			// fmt.Printf("got input %d\n", c.inputs[0])
		}

		icOps[op.op](&c, op.pm)

		if op.op == opcOutput {
			// fmt.Printf("sent output %d\n", c.outputs[0])
			output <- c.outputs[0]
			c.outputs = nil
		}
	}
	close(output)
}

// Running checks if the computer is running and has not halted or crashed
func (c *State) Running() bool {
	return c.status == Running
}

// Crashed checks if the computer halted in an abnormal manner and why
func (c *State) Crashed() (bool, string) {
	return (c.status == Crashed), c.faultReason
}

// Output gets the processor's output stream
func (c *State) Output() []int {
	return c.outputs
}
