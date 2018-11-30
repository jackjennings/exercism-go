// Package isogram implements functions relating to "nonpattern words"
package isogram

import "strings"

var sanitizer = strings.NewReplacer(
  " ", "",
  "-", "",
)

// IsIsogram determines if a provided word or phrase is an isogram.
// Spaces and hyphens are not counted as repeating characters for this purpose.
func IsIsogram(word string) bool {
  var sequence = sanitizer.Replace(strings.ToLower(word))
	var letters = make(map[rune]bool)

	for _, letter := range sequence {
		if _, ok := letters[letter]; ok {
			return false
		}

		letters[letter] = true
	}

	return true
}
