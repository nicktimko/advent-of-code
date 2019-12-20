/*
https://adventofcode.com/2019/day/3
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

type wireSegment struct {
	length int
	dir    *[2]int64
}

var dirs = map[string]([2]int64){
	"U": [2]int64{0, +1},
	"D": [2]int64{0, -1},
	"R": [2]int64{+1, 0},
	"L": [2]int64{-1, 0},
}

func readWire(reader *bufio.Reader) []wireSegment {
	wireCoords, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	rawTurns := strings.Split(strings.TrimSpace(wireCoords), ",")

	var turns = []wireSegment{}
	for _, i := range rawTurns {
		var seg wireSegment

		dir := dirs[i[:1]]

		seg.dir = &dir
		seg.length, err = strconv.Atoi(i[1:])
		if err != nil {
			log.Fatal(err)
		}
		turns = append(turns, seg)
	}

	return turns
}

func wireLen(wire []wireSegment) (sum int) {
	for _, seg := range wire {
		sum += seg.length
	}
	return
}

func pointsCrossed(wire []wireSegment) map[[2]int64]int8 {
	pts := make(map[[2]int64]int8)

	pt := [2]int64{0, 0}
	for _, seg := range wire {
		for i := 0; i < seg.length; i++ {
			pt[0] += seg.dir[0]
			pt[1] += seg.dir[1]
			pts[pt] = 1 // dummy value, doesn't matter.
		}
	}

	return pts
}

func AbsInt(n int64) int64 {
	if n < 0 {
		return -n
	}
	return n
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	wire1 := readWire(reader)
	wire2 := readWire(reader)

	w1pts := pointsCrossed(wire1)
	w2pts := pointsCrossed(wire2)

	fmt.Printf("wire 1 segments: %6d\n", len(wire1))
	fmt.Printf("       length:   %6d\n", wireLen(wire1))
	fmt.Printf("       points:   %6d\n", len(w1pts))
	fmt.Printf("wire 2 segments: %6d\n", len(wire2))
	fmt.Printf("       length:   %6d\n", wireLen(wire2))
	fmt.Printf("       points:   %6d\n", len(w2pts))

	intersections := make(map[[2]int64]int8)
	var min_intersection int64 = 0xFFF_FFFF_FFFF_FFFF

	for p1, _ := range w1pts {
		if _, ok := w2pts[p1]; ok {
			intersections[p1] = 1
			dist := AbsInt(p1[0]) + AbsInt(p1[1])
			if dist < min_intersection {
				min_intersection = dist
			}
		}
	}

	fmt.Printf("intersections:   %6d\n", len(intersections))
	fmt.Printf("closest dist :   %6d\n", min_intersection)
	// sum = 0
	// for _, seg := range wire2 {
	// 	sum += seg.length
	// }
	// fmt.Println("wire 2 pts: " + strconv.Itoa(sum))

	// fmt.Println("part 2: " + strconv.Itoa(fuelReq))
	/*
		Output
		---
		part 1: 3297909
		part 2: 4943994
	*/
}
