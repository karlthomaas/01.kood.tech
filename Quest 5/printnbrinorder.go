package piscine

import (
	"github.com/01-edu/z01"
)

func intToDigits(n int) (digits []int) {
	for n > 0 {
		if n == 0 {
			digits = append(digits, 0)
		} else {
			digits = append(digits, n%10)
		}
		n /= 10
	}
	return
}

func remove(s []int, i int) []int {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	digits := intToDigits(n) // digits in list
	var sorted_digits []int  // empty int list

	for {

		lock_state := false
		lowest_digit := -1
		lowest_index := -1

		for index, digit := range digits {

			// First iteration
			if lock_state == false {
				lowest_digit = digit
				lowest_index = index
				lock_state = true
				continue
			}

			// if the next digit is <= than current lower
			if digit < lowest_digit {
				lowest_digit = digit
				lowest_index = index
			}

		}
		digits = remove(digits, lowest_index)
		sorted_digits = append(sorted_digits, lowest_digit)

		if len(digits) == 0 {
			for _, digit := range sorted_digits {
				value := digit + '0'
				z01.PrintRune(rune(value))
			}
			break
		}
	}
}
