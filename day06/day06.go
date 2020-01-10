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

// AllParents returns the parents of the origin node in the graph, most distant first
func AllParents(parentGraph map[string]string, origin string) []string {
	var ok bool
	parents := []string{origin}

	for {
		origin, ok = parentGraph[origin]
		if !ok {
			break
		}
		parents = append([]string{origin}, parents...)
	}

	return parents
}

// LatestCommonParent returns the node of the most-closely related parent
func LatestCommonParent(parentGraph map[string]string, n1 string, n2 string) (latest string) {
	p1 := AllParents(parentGraph, n1)
	p2 := AllParents(parentGraph, n2)

	for i := range p1 {
		if p1[i] != p2[i] {
			break
		}
		latest = p1[i]
	}
	return
}

// Distance is how far apart nodes are on the DAG. -1 if infinitely far.
func Distance(parentGraph map[string]string, n1 string, n2 string) int {
	p1 := AllParents(parentGraph, n1)
	p2 := AllParents(parentGraph, n2)

	if p1[0] != p2[0] {
		return -1 // No common parent
	}

	common := 0
	for i := range p1 {
		if p1[i] != p2[i] {
			break
		}
		common++
	}
	return len(p1) + len(p2) - 2*common
}

// CountOrbits and the orbits of orbits, and the orbits of orbits of orbits...
func CountOrbits(parentOrbits map[string]string) (total int) {

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
	nOrbits := CountOrbits(parentOrbits)

	fmt.Printf("Part 1, total nested orbits: %d\n", nOrbits)

	fmt.Printf("Part 2, first common body  : %s\n",
		LatestCommonParent(parentOrbits, "YOU", "SAN"))
	fmt.Printf("            num transfers  : %d\n",
		Distance(parentOrbits, parentOrbits["YOU"], parentOrbits["SAN"]))
}
