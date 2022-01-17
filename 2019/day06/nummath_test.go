package day06

import "testing"

func TestTriangleNum(t *testing.T) {
	pairs := [][2]int{
		{0, 0},
		{1, 1},
		{2, 3},
		{4, 4 + 3 + 2 + 1},
		{10, 10 + 9 + 8 + 7 + 6 + 5 + 4 + 3 + 2 + 1},
		{100, 5050},
	}
	var n, expected, result int
	for _, pair := range pairs {
		n = pair[0]
		expected = pair[1]
		result = triangleNum(n)
		if result != expected {
			t.Errorf("triangle(%d) -> %d != %d", n, result, expected)
		}
	}
}

func TestTrapezoidNum(t *testing.T) {
	pairs := [][3]int{
		{0, 0, 0},
		{0, 1, 1},
		{5, 8, 6 + 7 + 8},
		{0, 100, 5050},
		{2, 100, 5050 - 2 - 1},
		{98, 100, 99 + 100},
	}
	var a, b, expected, result int
	for _, io := range pairs {
		a = io[0]
		b = io[1]
		expected = io[2]
		result = trapezoidNum(a, b)
		if result != expected {
			t.Errorf("trapezoid(%d, %d) -> %d != %d", a, b, result, expected)
		}
	}
}
