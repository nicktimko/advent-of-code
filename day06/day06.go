package day06

import (
	"fmt"
	"log"

	"github.com/nicktimko/aoc-2019-golang/aocio"
)

func triangleNum(n int) int {
	return (n * (n + 1)) / 2
}

func trapezoidNum(a int, b int) int {
	return triangleNum(b) - triangleNum(a)
}

func countOrbits(parentOrbits map[string]string) (total int) {

	var ok bool
	tallied := make(map[string]bool)
	// fmt.Printf("counting orbits, map size: %d\n", len(parentOrbits))

	for bodyID, parent := range parentOrbits {
		b := 0
		a := 0
		if _, ok := tallied[bodyID]; ok {
			// already saw the final child; can skip scanning entirely
			continue
		}

		for {
			// fmt.Printf(" - bodyID=%s parent=%s, b=%d, ba=%d\n", bodyID, parent, b, a)

			// already tallied?
			if _, ok := tallied[bodyID]; ok {
				// start incrementing a counter to omit it from calculation.
				a++
			}
			tallied[bodyID] = true

			b++

			bodyID = parent
			parent, ok = parentOrbits[parent]
			if !ok {
				break
			}
		}

		// fmt.Printf(" += trapz(%d, %d) = %d\n", a, b, trapezoidNum(a, b))
		total += trapezoidNum(a, b)
		// fmt.Printf(" tot = %d\n", total)
	}
	return
}

// Solve day 6 problem
func Solve() {
	lines, err := aocio.StringLines("inputs/day06.txt")
	if err != nil {
		log.Fatalf("error loading lines, %s", err)
	}

	parentOrbits := make(map[string]string)
	for _, v := range lines {
		// parentOrbits[v[:3]] = v[4:]  // child direction
		parentOrbits[v[4:]] = v[:3]
	}
	// fmt.Printf("%d %#v", len(parentOrbits), parentOrbits)
	nOrbits := countOrbits(parentOrbits)

	fmt.Printf("Part 1, total nested orbits: %d\n", nOrbits)
}
