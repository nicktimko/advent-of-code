/*
https://adventofcode.com/2019/day/4
*/

package day04

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

func validPasswordStep1(digits [6]int8) bool {
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

func validPasswordStep2(digits [6]int8) bool {
	currentRunLen := 0
	var currentRunDigit int8 = -1
	for _, d := range digits {
		if d != currentRunDigit {
			// new run
			if currentRunLen == 2 {
				return true
			}
			currentRunLen = 1
			currentRunDigit = d
		} else {
			// continuing run
			currentRunLen++
		}
	}
	if currentRunLen == 2 {
		// check at end of digits
		return true
	}
	return false
}

// Solve Day 4
func Solve() {

	nCheckedPasswords := 0
	nValidPasswordsStep1 := 0
	nValidPasswordsStep2 := 0

	var digits [6]int8

	for password := rangeStart; password <= rangeEnd; password++ {
		nCheckedPasswords++

		digits = decomposePassword(password)

		if validPasswordStep1(digits) {
			nValidPasswordsStep1++
			if validPasswordStep2(digits) {
				nValidPasswordsStep2++
			}
		}
	}

	fmt.Printf("passwords in range   : %6d\n", nCheckedPasswords)
	fmt.Printf("valid for first step : %6d\n", nValidPasswordsStep1)
	fmt.Printf("valid for second step: %6d\n", nValidPasswordsStep2)
	/*
		Output
		========
		passwords in range   : 518199
		valid for first step :   1764
		valid for second step:   1196
	*/
}
