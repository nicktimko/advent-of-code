/*
https://adventofcode.com/2019/day/1
*/

package day01

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// Solve Day 1
func Solve() {
	// something
	file, err := os.Open("inputs/day01.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var fuelReqNaive, fuelReq int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		i, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}

		moduleFuel := i/3 - 2
		fuelReqNaive += moduleFuel
		for moduleFuel > 0 {
			fuelReq += moduleFuel
			moduleFuel = moduleFuel/3 - 2
		}
	}
	fmt.Println("part 1: " + strconv.Itoa(fuelReqNaive))
	fmt.Println("part 2: " + strconv.Itoa(fuelReq))
}
