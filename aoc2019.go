package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/nicktimko/aoc-2019-golang/day01"
	"github.com/nicktimko/aoc-2019-golang/day02"
	"github.com/nicktimko/aoc-2019-golang/day03"
	"github.com/nicktimko/aoc-2019-golang/day04"
	"github.com/nicktimko/aoc-2019-golang/day05"
	"github.com/nicktimko/aoc-2019-golang/day06"
	"github.com/nicktimko/aoc-2019-golang/day07"
	"github.com/nicktimko/aoc-2019-golang/day08"
	"github.com/nicktimko/aoc-2019-golang/day09"
	"github.com/nicktimko/aoc-2019-golang/day10"
)

func main() {
	var ip = flag.Int("day", 1, "which day to run")
	flag.Parse()

	solutions := map[int](func()){
		1:  day01.Solve,
		2:  day02.Solve,
		3:  day03.Solve,
		4:  day04.Solve,
		5:  day05.Solve,
		6:  day06.Solve,
		7:  day07.Solve,
		8:  day08.Solve,
		9:  day09.Solve,
		10: day10.Solve,
	}
	solver, ok := solutions[*ip]
	if !ok {
		log.Fatalf("No solver for day %d", *ip)
	}

	fmt.Printf("Day %d Output\n==============\n", *ip)
	solver()
}
