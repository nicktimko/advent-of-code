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
	tape1 := append([]int(nil), tape...)
	outputs := intcode.IOProcessor(tape1, inputs)

	fmt.Println("Part 1:", outputs[len(outputs)-1])

	p2Inputs := []int{5}
	tape2 := append([]int(nil), tape...)
	p2Outputs := intcode.IOProcessor(tape2, p2Inputs)

	fmt.Println("Part 2:", p2Outputs[0])
}
