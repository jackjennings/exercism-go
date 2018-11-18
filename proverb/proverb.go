package proverb

import "fmt"

const (
	stanza = "For want of a %s the %s was lost."
	last   = "And all for the want of a %s."
)

func Proverb(rhyme []string) []string {
	var length = len(rhyme)
	var lines []string

	for i := 0; i < length; i++ {
		var line string

		if i < length-1 {
			line = fmt.Sprintf(stanza, rhyme[i], rhyme[i+1])
		} else {
			line = fmt.Sprintf(last, rhyme[0])
		}

		lines = append(lines, line)
	}

	return lines
}
