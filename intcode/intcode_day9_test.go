// Here are some example programs that use these features:
package intcode_test

import (
	"testing"

	"github.com/nicktimko/aoc-2019-golang/intcode"
)

func TestLargeOutput(t *testing.T) {
	largeOutput := []int{
		// should output a 16-digit number.
		1102, 34915192, 34915192, 7, 4, 7, 99, 0,
	}

	expected := 34915192 * 34915192
	output := intcode.IOProcessor(largeOutput, []int{})[0]

	if output != expected {
		t.Errorf("output differs: %#v %#v", output, expected)
	}
	/*
		NOTE: thankfully this works because we're on a 64-bit system, and int means
		int64. I'm too lazy to actually add a new computer to explicitly define that,
		and not sure if there's a way to make it versatile in Go...
	*/
}

func TestLargeOutput2(t *testing.T) {
	largeOutput2 := []int{
		// should output the large number in the middle.
		104, 1125899906842624, 99,
	}

	expected := 1125899906842624
	output := intcode.IOProcessor(largeOutput2, []int{})[0]

	if output != expected {
		t.Errorf("output differs: %#v %#v", output, expected)
	}
}

// func TestQuine(t *testing.T) {
// 	quine := []int{
// 		// takes no input and produces a copy of itself as output.
// 		109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99,
// 	}

// 	inputQuine := append([]int(nil), quine...)

// 	output := intcode.IOProcessor(inputQuine, []int{})

// 	if !toyshop.EqIntSlice(quine, output) {
// 		t.Errorf("slices differ: %#v %#v", quine, output)
// 	}
// }
