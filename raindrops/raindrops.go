// Package raindrops implements functions for traslating to raindropese
package raindrops

import (
	"strconv"
	"strings"
)

// Convert changes an integer into the raindrop-speak equivalent
func Convert(number int) string {
	var parts []string

	if number%3 == 0 {
		parts = append(parts, "Pling")
	}

	if number%5 == 0 {
		parts = append(parts, "Plang")
	}

	if number%7 == 0 {
		parts = append(parts, "Plong")
	}

	if len(parts) > 0 {
		return strings.Join(parts, "")
	} else {
		return strconv.Itoa(number)
	}
}
