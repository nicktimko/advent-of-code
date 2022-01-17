package aocio

import (
	"bufio"
	"os"
	"strconv"
)

// StringLines loads a file at fn and slices it up by line
func StringLines(fn string) ([]string, error) {
	file, err := os.Open(fn)
	if err != nil {
		return []string(nil), err
	}
	defer file.Close()

	lines := []string(nil)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	return lines, nil
}

// IntLines loads a file at fn and slices it up by line, and
// converts each line to an integer.
func IntLines(fn string) ([]int64, error) {
	var il []int64

	sl, err := StringLines(fn)
	if err != nil {
		return il, err
	}

	for _, s := range sl {
		i, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return il, err
		}
		il = append(il, i)
	}
	return il, nil
}
