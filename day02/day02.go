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

	var tape = []int{}
	instructions := strings.Split(strings.TrimSpace(contents), ",")
	for _, i := range instructions {
		j, err := strconv.Atoi(i)

		if err != nil {
			log.Fatal(err)
		}
		tape = append(tape, j)
	}

	tape[1] = 12
	tape[2] = 2

	var instruction int
	ptr := 0
	for {
		instruction = tape[ptr]
		if instruction == 1 {
			// add
			sum := tape[tape[ptr+1]] + tape[tape[ptr+2]]
			targetIndex := tape[ptr+3]
			tape[targetIndex] = sum
			fmt.Println("ADD wrote", sum, "to index", targetIndex)
			ptr += 4
		} else if instruction == 2 {
			// mul
			product := tape[tape[ptr+1]] * tape[tape[ptr+2]]
			targetIndex := tape[ptr+3]
			tape[targetIndex] = product
			fmt.Println("MUL wrote", product, "to index", targetIndex)
			ptr += 4
		} else if instruction == 99 {
			fmt.Println("HALT at index", ptr)
			break
		} else {
			log.Fatal("Bad instruction!")
		}
	}
	fmt.Println(tape[0])
	/*
		Output
		---
		5290681
	*/
}
