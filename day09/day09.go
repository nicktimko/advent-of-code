// https://adventofcode.com/2019/day/9

package day09

import (
	"log"

	"github.com/nicktimko/aoc-2019-golang/intcode"
)

func part1(tape []int) {

}

func part2(tape []int) {

}

// Solve day 9
func Solve() {
	tape, err := intcode.LoadTape("inputs/day09.txt")
	if err != nil {
		log.Fatal(err)
	}
	part1(tape)
	part2(tape)
}
