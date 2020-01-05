package intcode

import (
	"testing"
	// "github.com/nicktimko/aoc-2019-golang/intcode"
)

func expectTapeIndexEq(t *testing.T, s *intcodeComputerState, index int, want int) {
	if s.tape[index] != want {
		t.Errorf("Unexpected result in slot (was %d, wanted %d)", s.tape[index], want)
	}
	if s.status != Running {
		t.Error("Halted unexpectedly")
	}
}

func TestAddIndirect(t *testing.T) {
	s := intcodeComputerState{tape: []int{1, 4, 4, 4, 10}}
	want := s.tape[s.tape[1]] + s.tape[s.tape[2]]

	ProcessInstruction(&s)

	expectTapeIndexEq(t, &s, 4, want)
}

func TestAddImmediate(t *testing.T) {
	s := intcodeComputerState{tape: []int{1101, 1030, 204, 4, 0}}
	want := s.tape[1] + s.tape[2]

	ProcessInstruction(&s)

	expectTapeIndexEq(t, &s, 4, want)
}

func TestAddMixed(t *testing.T) {
	s := intcodeComputerState{tape: []int{101, 1030, 3, 4, 0}}
	want := s.tape[1] + s.tape[s.tape[2]]

	ProcessInstruction(&s)

	expectTapeIndexEq(t, &s, 4, want)
}

func TestMulIndirect(t *testing.T) {
	s := intcodeComputerState{tape: []int{2, 4, 4, 4, 10}}
	want := s.tape[s.tape[1]] * s.tape[s.tape[2]]

	ProcessInstruction(&s)

	expectTapeIndexEq(t, &s, 4, want)
}

func TestMulImmediate(t *testing.T) {
	s := intcodeComputerState{tape: []int{1102, 1030, 204, 4, 0}}
	want := s.tape[1] * s.tape[2]

	ProcessInstruction(&s)

	expectTapeIndexEq(t, &s, 4, want)
}

func TestInput(t *testing.T) {
	want := 123456
	s := intcodeComputerState{tape: []int{3, 2, 0}, inputs: []int{want}}

	ProcessInstruction(&s)

	expectTapeIndexEq(t, &s, 2, want)
}

func TestOutput(t *testing.T) {
	want := 6543
	s := intcodeComputerState{tape: []int{4, 2, want}}

	ProcessInstruction(&s)

	if s.outputs[0] != want {
		t.Errorf("bad output, wanted %d but got %d", want, s.outputs[0])
	}
}
