package day08

import (
	"fmt"
	"log"
	"os"
)

func part1() {

	height := 6
	width := 25
	layerSize := height * width

	file, err := os.Open("inputs/day08.txt")
	if err != nil {
		log.Fatal(err)
	}

	var data []byte
	stackUp := make([]rune, layerSize)

	fewestZeroes := layerSize + 1
	var layerSignature int

	for layer := 0; ; layer++ {
		data = make([]byte, layerSize)
		readCount, err := file.Read(data)
		if err != nil || readCount != layerSize {
			break
		}

		// part 1
		//////////////////////////
		// only care about 0/1/2 anyway. preallocate to avoid excessive
		// gymnastics later?
		// counts := map[byte]int{
		// 	'0': 0,
		// 	'1': 0,
		// 	'2': 0,
		// }
		counts := make(map[byte]int) // seems like all layers have them in it anyway.
		for _, b := range data {
			if _, ok := counts[b]; !ok {
				counts[b] = 1
			} else {
				counts[b]++
			}
		}
		if counts['0'] < fewestZeroes {
			fewestZeroes = counts['0']
			layerSignature = counts['1'] * counts['2']
		}

		// part 2
		for n, b := range stackUp {
			if b == '2' || b == 0 {
				if data[n] == '0' {
					stackUp[n] = ' '
				}
				if data[n] == '1' {
					stackUp[n] = '\u2588'
				}
			}
		}
	}

	fmt.Printf("part 1: %d\n", layerSignature)
	fmt.Print("part 2:\n")
	for i := 0; i < height; i++ {
		fmt.Printf("   %s\n", string(stackUp[width*i:width*(i+1)]))
	}
}

// Solve day 8
func Solve() {
	part1()
	// part2()
}
