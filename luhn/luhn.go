// Package luhn implements functions for working with the Luhn checksum algorithm
package luhn

import (
	"strconv"
	"strings"
)

// Valid checks if the given input passes the Luhn algorithm
func Valid(input string) bool {
	var sum int
	var sequence = strings.Split(strings.Replace(input, " ", "", -1), "")
	var length = len(sequence)

	if length == 1 {
		return false
	}

	for i := length; i > 0; i-- {
		v, err := strconv.Atoi(sequence[i-1])

		if err != nil {
			return false
		}

		sum += ((v*(((length-i)%2)+1))-1)%9 + 1
	}

	return sum%10 == 0
}
