package intcode_test

import (
	"testing"

	"github.com/nicktimko/aoc-2019-golang/intcode"
)

func TestAddIndirect(t *testing.T) {
	tape := []int{1, 4, 4, 4, 10}
	want := tape[tape[1]] + tape[tape[2]]

	ptr := 0
	halt := intcode.ProcessInstruction(tape, &ptr)

	if tape[4] != want {
		t.Errorf("Unexpected result in slot 4 (was %d, wanted %d)", tape[4], want)
	}
	if halt {
		t.Error("Halted unexpectedly")
	}
}

// func TestAddImmediate(t *testing.T) {
// 	tape := []int{1101, 1030, 204, 4, 0}
// 	want := tape[1] + tape[2]

// 	ptr := 0
// 	halt := intcode.ProcessInstruction(tape, &ptr)

// 	if tape[4] != want {
// 		t.Errorf("Unexpected result in slot 4 (was %d, wanted %d)", tape[4], want)
// 	}
// 	if halt {
// 		t.Error("Halted unexpectedly")
// 	}
// }
