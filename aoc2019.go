package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/nicktimko/aoc-2019-golang/day01"
	"github.com/nicktimko/aoc-2019-golang/day02"
	"github.com/nicktimko/aoc-2019-golang/day03"
	"github.com/nicktimko/aoc-2019-golang/day04"
)

func main() {
	var ip = flag.Int("day", 1, "which day to run")
	flag.Parse()

	solutions := map[int](func()){
		1: day01.Solve,
		2: day02.Solve,
		3: day03.Solve,
		4: day04.Solve,
	}
	solver, ok := solutions[*ip]
	if !ok {
		log.Fatalf("No solver for day %d", *ip)
	}

	fmt.Printf("Day %d Output\n==============\n", *ip)
	solver()
}