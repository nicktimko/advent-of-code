/*
https://adventofcode.com/2019/day/3
*/

package day03

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

func pointsCrossed(wire []wireSegment) map[[2]int64]int64 {
	pts := make(map[[2]int64]int64)

	pt := [2]int64{0, 0}
	var step int64 = 0
	for _, seg := range wire {
		for i := 0; i < seg.length; i++ {
			step += 1
			pt[0] += seg.dir[0]
			pt[1] += seg.dir[1]
			pts[pt] = step
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

// Solve Day 3
func Solve() {
	file, err := os.Open("inputs/day03.txt")
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

	intersections := make(map[[2]int64]int64)
	var min_intersection int64 = 0xFFF_FFFF_FFFF_FFFF
	var shortest_sum int64 = 0xFFF_FFFF_FFFF_FFFF

	for p1, p1d := range w1pts {
		if p2d, ok := w2pts[p1]; ok {
			intersections[p1] = 1
			dist := AbsInt(p1[0]) + AbsInt(p1[1])
			if dist < min_intersection {
				min_intersection = dist
			}
			total_sum := p1d + p2d
			if total_sum < shortest_sum {
				shortest_sum = total_sum
			}
		}
	}

	fmt.Printf("intersections:   %6d\n", len(intersections))
	fmt.Printf("closest dist:    %6d\n", min_intersection)
	fmt.Printf("short from src:  %6d\n", shortest_sum)
}
