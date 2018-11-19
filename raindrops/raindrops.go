// Package raindrops implements functions for translating to raindropese
package raindrops

import "strconv"

// Convert changes an integer into the raindrop-speak equivalent
func Convert(number int) string {
	var translation string

	if number%3 == 0 {
		translation += "Pling"
	}

	if number%5 == 0 {
		translation += "Plang"
	}

	if number%7 == 0 {
		translation += "Plong"
	}

	if len(translation) > 0 {
		return translation
	} else {
		return strconv.Itoa(number)
	}
}
