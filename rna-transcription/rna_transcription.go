// Package strand implements functions for transcribing DNA strands
package strand

import "strings"

var RNAReplacer = strings.NewReplacer(
	"G", "C",
	"C", "G",
	"T", "A",
	"A", "U",
)

// ToRNA converts a DNS strand into the RNA compliment
func ToRNA(dna string) string {
	return RNAReplacer.Replace(dna)
}
