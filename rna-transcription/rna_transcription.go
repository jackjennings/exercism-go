package strand

import "strings"

var ComplementTable = map[string]string{
	"C": "G",
	"G": "C",
	"T": "A",
	"A": "U",
}

func ToRNA(dna string) string {
	var nucleotides = strings.Split(dna, "")
	var complement []string

	for i := 0; i < len(nucleotides); i++ {
		complement = append(complement, ComplementTable[nucleotides[i]])
	}

	return strings.Join(complement, "")
}
