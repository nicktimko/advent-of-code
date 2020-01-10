package aocio_test

import (
	"testing"

	"github.com/nicktimko/aoc-2019-golang/aocio"
)

func TestStringLines(t *testing.T) {
	sl, err := aocio.StringLines("../test_inputs/strings1.txt")
	if err != nil {
		t.Fatalf("error loading file: %s", err)
	}
	want := []string{"a", "b", "c"}

	if !aocio.EqStringSlice(sl, want) {
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

	if !aocio.EqIntSlice(il, want) {
		t.Errorf("slices differ: %#v %#v", want, il)
	}
}
