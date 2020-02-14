// https://adventofcode.com/2019/day/9

package day09

import (
	"fmt"
	"log"

	"github.com/nicktimko/aoc-2019-golang/intcode"
)

func printPart(n int, output []int) {
	if len(output) == 1 {
		fmt.Printf("Part %d: %d\n", n, output[0])
	} else {
		fmt.Printf("Part %d (error?): %#v\n", n, output)
	}
}

func part1(tape []int) {
	testModeInputs := []int{1}
	output := intcode.IOProcessor(tape, testModeInputs)
	printPart(1, output)
}

func part2(tape []int) {
	boostModeInputs := []int{2}
	output := intcode.IOProcessor(tape, boostModeInputs)
	printPart(2, output)
}

// Solve day 9
func Solve() {
	tape, err := intcode.LoadTape("inputs/day09.txt")
	if err != nil {
		log.Fatal(err)
	}
	part1(append([]int(nil), tape...))
	part2(append([]int(nil), tape...))
}
