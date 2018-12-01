// Package reverse implements functions for changing the direction of things
package reverse

import "strings"

// String reverses the order of the runes in a string
func String(input string) string {
	var output string
	var sequence = strings.Split(input, "")

	for i := len(sequence) - 1; i >= 0; i-- {
		output += sequence[i]
	}

	return output
}
