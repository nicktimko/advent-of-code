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
	fmt.Println("ADD wrote", sum, "to index", targetIndex)

	*ptr += 4
	return false
}

func icMul(tape []int, ptr *int) bool {
	product := tape[tape[*ptr+1]] * tape[tape[*ptr+2]]
	targetIndex := tape[*ptr+3]
	tape[targetIndex] = product
	fmt.Println("MUL wrote", product, "to index", targetIndex)

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

	fmt.Println(tape[0])
	/*
		Output
		---
		5290681
	*/
}
