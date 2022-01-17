package intcode

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// LoadTape from a file
func LoadTape(fn string) ([]int, error) {
	file, err := os.Open(fn)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	contents, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}

	var tape = []int{}
	instructions := strings.Split(strings.TrimSpace(contents), ",")
	for _, i := range instructions {
		j, err := strconv.Atoi(i)
		if err != nil {
			return nil, err
		}
		tape = append(tape, j)
	}
	return tape, nil
}
