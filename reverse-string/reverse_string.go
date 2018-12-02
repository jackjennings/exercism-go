// Package reverse implements functions for changing the direction of things
package reverse

import "strings"

// String reverses the order of the runes in a string
func String(input string) string {
	var b strings.Builder

	b.Grow(len(input))

	runes := []rune(input)
	for i := len(runes) - 1; i >= 0; i-- {
		b.WriteRune(runes[i])
	}

	return b.String()
}
