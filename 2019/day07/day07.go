// https://adventofcode.com/2019/day/7

package day07

import (
	"fmt"
	"log"

	"github.com/nicktimko/aoc-2019-golang/intcode"
	"github.com/nicktimko/aoc-2019-golang/toyshop"
)

func amplifier(tape []int, phase int, input int) int {
	inputs := []int{phase, input}
	outputs := intcode.IOProcessor(append([]int(nil), tape...), inputs)
	return outputs[0]
}

func part1(tape []int) {
	permutes := []int{0, 1, 2, 3, 4}
	c := make(chan []int)
	go toyshop.PermutationsInt(c, permutes, 5)

	var maxPermutation []int
	maxOutput := -1_234_567_890

	for permutation := range c {
		signal := 0
		for _, phase := range permutation {
			signal = amplifier(tape, phase, signal)
		}
		if signal > maxOutput {
			maxOutput = signal
			maxPermutation = permutation
		}
	}

	fmt.Printf("Part 1, max output:   %d\n", maxOutput)
	fmt.Printf("        phase signal: %v\n", maxPermutation)
}

func part2(tape []int) {
	permutes := []int{5, 6, 7, 8, 9}
	c := make(chan []int)
	go toyshop.PermutationsInt(c, permutes, 5)

	var maxPermutationFb []int
	maxOutputFb := -1_234_567_890

	for permutation := range c {
		var links [6]chan int
		for i := range links {
			bsize := 1
			if i == 0 {
				bsize = 2 // allows setting initial input/phase without problems
			}
			links[i] = make(chan int, bsize)
		}
		output := make(chan int, 1)

		for i := 0; i < len(links)-1; i++ {
			// set phase
			links[i] <- permutation[i]

			go intcode.CommunicatingProcessor(
				append([]int(nil), tape...),
				links[i],
				links[i+1],
			)

			// initial input
			if i == 0 {
				links[i] <- 0
			}
		}
		go toyshop.TeeInt(links[len(links)-1], links[0], output)

		var signal int
		for signal = range output {
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

// Solve day 7
func Solve() {
	tape, err := intcode.LoadTape("inputs/day07.txt")
	if err != nil {
		log.Fatal(err)
	}
	part1(tape)
	part2(tape)
}
