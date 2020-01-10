package aocio_test

import (
	"testing"

	"github.com/nicktimko/aoc-2019-golang/aocio"
)

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

func TestStringLines(t *testing.T) {
	sl, err := aocio.StringLines("../test_inputs/strings1.txt")
	if err != nil {
		t.Fatalf("error loading file: %s", err)
	}
	want := []string{"a", "b", "c"}

	if !EqStringSlice(sl, want) {
		t.Errorf("slices differ: %#v %#v", want, sl)
	}
}

func TestIntLines(t *testing.T) {
	il, err := aocio.IntLines("../test_inputs/ints1.txt")
	if err != nil {
		t.Fatalf("error loading file: %s", err)
	}
	want := []int64{
		1,
		12345,
		1234567890987654321, // ~60 bits
	}

	if !EqIntSlice(il, want) {
		t.Errorf("slices differ: %#v %#v", want, il)
	}
}
