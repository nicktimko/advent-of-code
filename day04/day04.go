/*
https://adventofcode.com/2019/day/4
*/

package main

import "fmt"

var rangeStart int32 = 152085
var rangeEnd int32 = 670283

func decomposePassword(password int32) [6]int8 {
	var digits [6]int8
	for n := range digits {
		digits[5-n] = int8(password % 10)
		password = password / 10
	}
	return digits
}

func monotonicDigits(digits []int8) bool {
	for n := range digits {
		if n == 0 {
			continue
		}
		// Going from left to right, the digits never decrease;
		// they only ever increase or stay the same (like 111123
		// or 135679).
		if digits[n-1] > digits[n] {
			return false
		}
	}
	return true
}

func validPasswordStep1(password int32) bool {
	digits := decomposePassword(password)
	if !monotonicDigits(digits[:]) {
		return false
	}
	// Two adjacent digits are the same (like 22 in 122345).
	sameFound := false
	for n := range digits {
		if n == 0 {
			continue
		}
		if digits[n-1] == digits[n] {
			sameFound = true
		}
	}
	return sameFound
}

func main() {

	nCheckedPasswords := 0
	nValidPasswords := 0

	for password := rangeStart; password <= rangeEnd; password++ {
		nCheckedPasswords++
		if validPasswordStep1(password) {
			nValidPasswords++
		}
	}

	fmt.Printf("passwords in range: %6d\n", nCheckedPasswords)
	fmt.Printf("valid             : %6d\n", nValidPasswords)
	/*
		Output
		========
		passwords in range: 518199
		valid             :   1764
	*/
}
