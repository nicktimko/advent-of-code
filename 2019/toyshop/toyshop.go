package toyshop

import (
	"math"
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

// EqInt64Slice returns true if the slices of int64s are identical
func EqInt64Slice(a []int64, b []int64) bool {
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

// EqIntSlice returns true if the slices of ints are identical
func EqIntSlice(a []int, b []int) bool {
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

// Range replicates Python's range()
func Range(start int, stop int, step int) []int {
	size := int(math.Ceil(float64(stop-start) / float64(step)))
	if size < 0 {
		size = 0
	}
	r := make([]int, size)
	for n := range r {
		r[n] = start + n*step
	}
	return r
}

// PermutationsInt returns all permutations of iterable of length r via the channel
func PermutationsInt(c chan []int, iterable []int, r int) {
	// adapted from Python pseudo-implementation at
	// https://docs.python.org/3/library/itertools.html#itertools.permutations
	n := len(iterable)
	if r > n {
		return
	}
	indices := Range(0, n, 1)
	cycles := Range(n, n-r, -1)

	var outp []int

	outp = make([]int, r)
	for i := range outp {
		outp[i] = iterable[indices[i]]
	}
	c <- outp

	if n == 0 {
		return
	}
	for {
		for i := r - 1; i >= 0; i-- {
			cycles[i]--
			if cycles[i] == 0 {
				// indices[i:] = indices[i+1:] + indices[i:i+1]
				for j := i; j <= len(indices)-2; j++ {
					indices[j], indices[j+1] = indices[j+1], indices[j]
				}
				cycles[i] = n - i
			} else {
				// "negative j" with respect to indices
				nj := len(indices) - cycles[i]
				indices[i], indices[nj] = indices[nj], indices[i]
				outp = make([]int, r)
				for i := range outp {
					outp[i] = iterable[indices[i]]
				}
				c <- outp
				break
			}
			if i == 0 {
				close(c)
				return
			}
		}
	}
}

func TeeInt(input chan int, output1 chan int, output2 chan int) {
	for x := range input {
		output1 <- x
		output2 <- x
	}
	close(output1)
	close(output2)
}
