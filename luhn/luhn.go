// Package luhn implements functions for working with the Luhn checksum algorithm
package luhn

import "strings"

// Valid checks if the given input passes the Luhn algorithm
func Valid(input string) bool {
	var sum int
	var sequence = strings.Replace(input, " ", "", -1)
	var length = len(sequence)

	if length == 1 {
		return false
	}

	for i := length; i > 0; i-- {
		// Substracting the '0' rune gets the character code offset of subsequent numbers
		digit := int(sequence[i-1] - '0')

		// Anything outside of 0-9 is some other character
		if digit > 9 || digit < 0 {
			return false
		}

		// Double every second digit, subtracting 9 for values greater than 9
		sum += ((digit*(((length-i)%2)+1))-1)%9 + 1
	}

	return sum%10 == 0
}
