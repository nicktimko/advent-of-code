package intcode

import (
	"testing"
	// "github.com/nicktimko/aoc-2019-golang/intcode"
)

func expectTapeIndexEq(t *testing.T, s *State, index int, want int) {
	if s.tape[index] != want {
		t.Errorf("Unexpected result in slot (was %d, wanted %d)", s.tape[index], want)
	}
}

func expectRunning(t *testing.T, s *State) {
	if !s.Running() {
		t.Error("Halted unexpectedly")
	}
}

func TestAddIndirect(t *testing.T) {
	s := State{tape: []int{1, 4, 4, 4, 10}}
	want := s.tape[s.tape[1]] + s.tape[s.tape[2]]

	s.ProcessInstruction()

	expectTapeIndexEq(t, &s, 4, want)
	expectRunning(t, &s)
}

func TestAddImmediate(t *testing.T) {
	s := State{tape: []int{1101, 1030, 204, 4, 0}}
	want := s.tape[1] + s.tape[2]

	s.ProcessInstruction()

	expectTapeIndexEq(t, &s, 4, want)
	expectRunning(t, &s)
}

func TestAddMixed(t *testing.T) {
	s := State{tape: []int{101, 1030, 3, 4, 0}}
	want := s.tape[1] + s.tape[s.tape[2]]

	s.ProcessInstruction()

	expectTapeIndexEq(t, &s, 4, want)
	expectRunning(t, &s)
}

func TestMulIndirect(t *testing.T) {
	s := State{tape: []int{2, 4, 4, 4, 10}}
	want := s.tape[s.tape[1]] * s.tape[s.tape[2]]

	s.ProcessInstruction()

	expectTapeIndexEq(t, &s, 4, want)
	expectRunning(t, &s)
}

func TestMulImmediate(t *testing.T) {
	s := State{tape: []int{1102, 1030, 204, 4, 0}}
	want := s.tape[1] * s.tape[2]

	s.ProcessInstruction()

	expectTapeIndexEq(t, &s, 4, want)
	expectRunning(t, &s)
}

func TestInput(t *testing.T) {
	want := 123456
	s := State{tape: []int{3, 2, 0}, inputs: []int{want}}

	s.ProcessInstruction()

	expectTapeIndexEq(t, &s, 2, want)
	expectRunning(t, &s)
}

func TestInputRelbase(t *testing.T) {
	want := 123456
	s := State{
		tape:         []int{0, 203, 2 - 44, 0},
		inputs:       []int{want},
		relativeBase: 44,
		ptr:          1,
	}

	s.ProcessInstruction()

	if s.faultReason != "" {
		t.Errorf("crashed processor: %s\n", s.faultReason)
	}
	expectTapeIndexEq(t, &s, 2, want)
	expectRunning(t, &s)
	t.Logf("tape: %#v", s.tape)
}

func TestOutput(t *testing.T) {
	want := 6543
	s := State{tape: []int{4, 2, want}}

	s.ProcessInstruction()

	if s.outputs[0] != want {
		t.Errorf("bad output, wanted %d but got %d", want, s.outputs[0])
	}
	expectRunning(t, &s)
}

func TestLTTrue(t *testing.T) {
	s := State{tape: []int{1107, 1, 2, 4, -1}}
	want := 1

	s.ProcessInstruction()

	expectTapeIndexEq(t, &s, 4, want)
	expectRunning(t, &s)
}

func TestLTFalse(t *testing.T) {
	s := State{tape: []int{1107, 2, 2, 4, -1}}
	want := 0

	s.ProcessInstruction()

	expectTapeIndexEq(t, &s, 4, want)
	expectRunning(t, &s)
}

func TestEQTrue(t *testing.T) {
	s := State{tape: []int{1108, 2, 2, 4, -1}}
	want := 1

	s.ProcessInstruction()

	expectTapeIndexEq(t, &s, 4, want)
	expectRunning(t, &s)
}

func TestEQFalse(t *testing.T) {
	s := State{tape: []int{1108, 1, 2, 4, -1}}
	want := 0

	s.ProcessInstruction()

	expectTapeIndexEq(t, &s, 4, want)
	expectRunning(t, &s)
}

func TestJumpIfTrueJumping(t *testing.T) {
	wantPtr := 100
	s := State{tape: []int{105, 1, 3, wantPtr}}

	s.ProcessInstruction()

	if s.ptr != 100 {
		t.Errorf("Unexpected pointer (was %d, wanted %d)", s.ptr, wantPtr)
	}
	expectRunning(t, &s)
}

func TestJumpIfTrueNoJumping(t *testing.T) {
	wantPtr := 3
	s := State{tape: []int{105, 0, 3, 9999}}

	s.ProcessInstruction()

	if s.ptr != wantPtr {
		t.Errorf("Unexpected pointer (was %d, wanted %d)", s.ptr, wantPtr)
	}
	expectRunning(t, &s)
}

func TestJumpIfFalseJumping(t *testing.T) {
	wantPtr := 100
	s := State{tape: []int{106, 0, 3, wantPtr}}

	s.ProcessInstruction()

	if s.ptr != 100 {
		t.Errorf("Unexpected pointer (was %d, wanted %d)", s.ptr, wantPtr)
	}
	expectRunning(t, &s)
}

func TestJumpIfFalseNoJumping(t *testing.T) {
	wantPtr := 3
	s := State{tape: []int{106, 1, 3, 9999}}

	s.ProcessInstruction()

	if s.ptr != wantPtr {
		t.Errorf("Unexpected pointer (was %d, wanted %d)", s.ptr, wantPtr)
	}
	expectRunning(t, &s)
}
