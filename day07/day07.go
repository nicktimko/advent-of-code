package day07

import (
	"fmt"
	"log"

	"github.com/nicktimko/aoc-2019-golang/intcode"
	"github.com/nicktimko/aoc-2019-golang/toyshop"
)

func amplifier(program []int, phase int, input int) int {
	inputs := []int{phase, input}
	outputs := intcode.IOProcessor(program, inputs)
	return outputs[0]
}

// Solve day 7
func Solve() {
	tape, err := intcode.LoadTape("inputs/day07.txt")
	if err != nil {
		log.Fatal(err)
	}

	permutes := []int{0, 1, 2, 3, 4}
	c := make(chan []int)
	go toyshop.PermutationsInt(c, permutes, 5)

	var maxPermutation []int
	maxOutput := -1_234_567_890

	for permutation := range c {
		signal := 0
		for _, phase := range permutation {
			// outputs := intcode.IOProcessor(tape1, inputs)
			ampTape := append([]int(nil), tape...)
			signal = amplifier(ampTape, phase, signal)
		}
		if signal > maxOutput {
			maxOutput = signal
			maxPermutation = permutation
		}
	}

	fmt.Printf("Part 1, max output:   %d\n", maxOutput)
	fmt.Printf("        phase signal: %v\n", maxPermutation)

	// p2Inputs := []int{5}
	// tape2 := append([]int(nil), tape...)
	// p2Outputs := intcode.IOProcessor(tape2, p2Inputs)

	// fmt.Println("Part 2:", p2Outputs[0])
}
