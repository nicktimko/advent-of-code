package intcode

import (
	"testing"
)

func TestSomeProblem(t *testing.T) {
	badTapeFragment := []int{
		// should output a 16-digit number.
		0, 203, -100, 0,
	}

	expected := 42
	s := New(badTapeFragment, []int{expected})
	s.ptr = 1
	s.relativeBase = 100

	s.ProcessInstruction()

	produced := badTapeFragment[0]
	if produced != expected {
		t.Errorf("produced differs from expected: %d != %d", produced, expected)
		t.Errorf("tape: %#v\n", badTapeFragment)
	}
}
