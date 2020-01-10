package aocio

import (
	"bufio"
	"os"
	"strconv"
)

// EqStringSlice returns true if the slices of strings are identical
func EqStringSlice(a []string, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// EqIntSlice returns true if the slices of int64s are identical
func EqIntSlice(a []int64, b []int64) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

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
