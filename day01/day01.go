/*
https://adventofcode.com/2019/day/1
*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
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
	/*
		Output
		---
		part 1: 3297909
		part 2: 4943994
	*/
}
