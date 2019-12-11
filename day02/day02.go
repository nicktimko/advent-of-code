/*
https://adventofcode.com/2019/day/2
*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func icAdd(tape []int, ptr *int) bool {
	sum := tape[tape[*ptr+1]] + tape[tape[*ptr+2]]
	targetIndex := tape[*ptr+3]
	tape[targetIndex] = sum
	// fmt.Println("ADD wrote", sum, "to index", targetIndex)

	*ptr += 4
	return false
}

func icMul(tape []int, ptr *int) bool {
	product := tape[tape[*ptr+1]] * tape[tape[*ptr+2]]
	targetIndex := tape[*ptr+3]
	tape[targetIndex] = product
	// fmt.Println("MUL wrote", product, "to index", targetIndex)

	*ptr += 4
	return false
}

func icHalt(tape []int, ptr *int) bool {
	*ptr++
	return true
}

func intcodeProcessor(tape []int) {
	icOps := map[int](func([]int, *int) bool){
		1:  icAdd,
		2:  icMul,
		99: icHalt,
	}
	ptr := 0
	halt := false
	for !halt {
		halt = icOps[tape[ptr]](tape, &ptr)
	}
}

func part2(tape []int) int {
	target := 19690720

	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			workingTape := append([]int(nil), tape...)
			workingTape[1] = noun
			workingTape[2] = verb

			intcodeProcessor(workingTape)

			if workingTape[0] == target {
				return noun*100 + verb
			}
		}
	}
	log.Fatal("couldn't find inputs to give desired output")
	return -1
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	contents, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	var originalTape = []int{}
	instructions := strings.Split(strings.TrimSpace(contents), ",")
	for _, i := range instructions {
		j, err := strconv.Atoi(i)
		if err != nil {
			log.Fatal(err)
		}
		originalTape = append(originalTape, j)
	}

	// make a copy to work on
	tape := append([]int(nil), originalTape...)

	tape[1] = 12
	tape[2] = 2

	intcodeProcessor(tape)

	fmt.Println("Part 1:", tape[0])

	nounverb := part2(originalTape)
	fmt.Println("Part 2:", nounverb)
	/*
		Output
		---
		Part 1: 5290681
		Part 2: 5741
	*/
}
