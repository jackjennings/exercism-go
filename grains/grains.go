// Package grains implements functions relating to chess-based wheat cultivation
package grains

import "fmt"

// Square returns the number of wheat grains on the nth square of a chess board
func Square(index int) (uint64, error) {
	if index <= 0 || index > 64 {
		return 0, fmt.Errorf("Square number %d is invalid", index)
	}

	var total uint64 = 1

	for n := 1; n < index; n++ {
		total = total << 1
	}

	return total, nil
}

// Total returns the number of wheat grains on a (very large) chess board
func Total() uint64 {
	var total uint64

	for n := 1; n <= 64; n++ {
		value, _ := Square(n)
		total += value
	}

	return total
}
