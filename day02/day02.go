/*
https://adventofcode.com/2019/day/2
*/

package day02

import (
	"fmt"
	"log"

	"github.com/nicktimko/aoc-2019-golang/intcode"
)

func part2(tape []int) int {
	target := 19690720

	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			workingTape := append([]int(nil), tape...)
			workingTape[1] = noun
			workingTape[2] = verb

			intcode.Processor(workingTape)

			if workingTape[0] == target {
				return noun*100 + verb
			}
		}
	}
	log.Fatal("couldn't find inputs to give desired output")
	return -1
}

// Solve Day 2
func Solve() {
	originalTape, err := intcode.LoadTape("inputs/day02.txt")
	if err != nil {
		log.Fatal(err)
	}

	// make a copy to work on
	tape := append([]int(nil), originalTape...)

	tape[1] = 12
	tape[2] = 2

	intcode.Processor(tape)

	fmt.Println("Part 1:", tape[0])

	nounverb := part2(originalTape)
	fmt.Println("Part 2:", nounverb)
}
