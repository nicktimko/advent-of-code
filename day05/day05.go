/*
https://adventofcode.com/2019/day/5
*/

package day05

import (
	"fmt"
	"log"

	"github.com/nicktimko/aoc-2019-golang/intcode"
)

// Solve Day 5
func Solve() {
	tape, err := intcode.LoadTape("inputs/day05.txt")
	if err != nil {
		log.Fatal(err)
	}

	inputs := []int{1}

	outputs := intcode.IOProcessor(tape, inputs)

	fmt.Println("Part 1:", outputs[len(outputs)-1])
}
