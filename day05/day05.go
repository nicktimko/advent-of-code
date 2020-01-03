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

	// intcode.Processor(tape)

	fmt.Println("Part 1:", tape[0])
}
