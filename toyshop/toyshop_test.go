package toyshop_test

import (
	"testing"

	"github.com/nicktimko/aoc-2019-golang/toyshop"
)

func TestRangeSimple(t *testing.T) {
	r := toyshop.Range(0, 5, 1)
	want := []int{0, 1, 2, 3, 4}
	if !toyshop.EqIntSlice(r, want) {
		t.Errorf("slices differ: %#v %#v", want, r)
	}
}

func TestRangeReverse(t *testing.T) {
	r := toyshop.Range(5, 0, -1)
	want := []int{5, 4, 3, 2, 1}
	if !toyshop.EqIntSlice(r, want) {
		t.Errorf("slices differ: %#v %#v", want, r)
	}
}

func TestRangeStride(t *testing.T) {
	r := toyshop.Range(0, 5, 2)
	want := []int{0, 2, 4}
	if !toyshop.EqIntSlice(r, want) {
		t.Errorf("slices differ: %#v %#v", want, r)
	}
}

func TestRangeZeroSize(t *testing.T) {
	r := toyshop.Range(0, 0, 1)
	want := []int{}
	if !toyshop.EqIntSlice(r, want) {
		t.Errorf("slices differ: %#v %#v", want, r)
	}
}

func TestRangeLessThanZeroSize(t *testing.T) {
	r := toyshop.Range(0, -2, 1)
	want := []int{}
	if !toyshop.EqIntSlice(r, want) {
		t.Errorf("slices differ: %#v %#v", want, r)
	}
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return factorial(n-1) * n
}

func nPr(n int, r int) int {
	return factorial(n) / factorial(n-r)
}

func TestPermutationsInt(t *testing.T) {
	permutes := []int{1, 2, 3}
	c := make(chan []int)
	go toyshop.PermutationsInt(c, permutes, 2)

	outputs := make(map[[2]int]bool)

	for val := range c {
		aval := [2]int{val[0], val[1]}
		outputs[aval] = true
	}
	if len(outputs) != nPr(len(permutes), 2) {
		t.Errorf("len unexpected, was %d expected %d", len(outputs), nPr(4, 2))
	}
	expected := map[[2]int]bool{
		{1, 2}: true,
		{1, 3}: true,
		{2, 1}: true,
		{2, 3}: true,
		{3, 1}: true,
		{3, 2}: true,
	}
	for k, av := range expected {
		bv, ok := outputs[k]
		if !ok {
			t.Errorf("missing in outputs: %v", k)
		}
		if av != bv {
			t.Errorf("differing in outputs for k=%v: %v and %v", k, av, bv)
		}
	}
}
