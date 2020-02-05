package intcode_test

import (
	"testing"

	"github.com/nicktimko/aoc-2019-golang/intcode"
)

func TestQuine(t *testing.T) {
	quine := []int{
		// takes no input and produces a copy of itself as output.
		109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99,
	}

	// copy input to avoid getting trashed for comparison
	inputTape := append([]int(nil), quine...)

	proc := intcode.New(inputTape, []int{})

	for proc.Running() {
		proc.ProcessInstruction()
	}

	// if proc.Crashed() {
	// 	t.Errorf("unexpected processor halt!")
	// }

	// output := proc.Output()
	// if !toyshop.EqIntSlice(quine, output) {
	// 	t.Errorf("slices differ: %#v %#v", quine, output)
	// }
}
