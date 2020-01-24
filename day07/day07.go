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

	// Part 2
	permutes2 := []int{5, 6, 7, 8, 9}
	c2 := make(chan []int)
	go toyshop.PermutationsInt(c2, permutes2, 5)

	var maxPermutationFb []int
	maxOutputFb := -1_234_567_890

	for permutation := range c2 {
		c0 := make(chan int, 1)
		c1 := make(chan int, 1)
		c2 := make(chan int, 1)
		c3 := make(chan int, 1)
		c4 := make(chan int, 1)
		c5 := make(chan int, 1)
		cOutput := make(chan int, 1)

		// fmt.Printf("%v\n", permutation)

		go intcode.CommunicatingProcessor(append([]int(nil), tape...), c0, c1)
		go intcode.CommunicatingProcessor(append([]int(nil), tape...), c1, c2)
		go intcode.CommunicatingProcessor(append([]int(nil), tape...), c2, c3)
		go intcode.CommunicatingProcessor(append([]int(nil), tape...), c3, c4)
		go intcode.CommunicatingProcessor(append([]int(nil), tape...), c4, c5)
		go toyshop.TeeInt(c5, c0, cOutput)

		c0 <- permutation[0]
		c1 <- permutation[1]
		c2 <- permutation[2]
		c3 <- permutation[3]
		c4 <- permutation[4]
		c0 <- 0

		var signal int
		for signal = range cOutput {
			// just exhaust the channel
		}
		if signal > maxOutputFb {
			maxOutputFb = signal
			maxPermutationFb = permutation
		}
	}

	fmt.Printf("Part 2, max output:   %d\n", maxOutputFb)
	fmt.Printf("        phase signal: %v\n", maxPermutationFb)
}
