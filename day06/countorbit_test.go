package day06

import "testing"

type countOrbitTrial struct {
	input    map[string]string
	expected int
}

func TestCounter(t *testing.T) {
	trials := []countOrbitTrial{
		{map[string]string{"B": "COM"}, 1},
		{map[string]string{"B": "COM", "C": "B"}, 3},
		{map[string]string{"B": "COM", "C": "COM"}, 2},
		{map[string]string{
			"B": "COM",
			"C": "B",
			"D": "C",
			"E": "D",
			"F": "E",
			"G": "B",
			"H": "G",
			"I": "D",
			"J": "E",
			"K": "J",
			"L": "K",
		}, 42},
	}
	for n, io := range trials {
		result := countOrbits(io.input)
		if result != io.expected {
			t.Errorf("countOrbits(trials[%d].input) -> %d != %d", n, result, io.expected)
		}
	}
}
