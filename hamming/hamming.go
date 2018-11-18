package hamming

import (
	"errors"
	"strings"
)

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("DNA strands are of unequal length")
	}

	aNucleotides := strings.Split(a, "")
	bNucleotides := strings.Split(b, "")
	distance := 0

	for i := 0; i < len(aNucleotides); i++ {
		if aNucleotides[i] != bNucleotides[i] {
			distance++
		}
	}

	return distance, nil
}
