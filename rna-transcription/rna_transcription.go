// Package strand implements functions for transcribing DNA strands
package strand

import "strings"

// ToRNA converts a DNS strand into the RNA compliment
func ToRNA(dna string) string {
	return strings.NewReplacer(
		"G", "C",
		"C", "G",
		"T", "A",
		"A", "U",
	).Replace(dna)
}
