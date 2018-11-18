package proverb

import "fmt"

func Proverb(rhyme []string) []string {
	var length = len(rhyme)
	var lines []string

	for i := 0; i < length; i++ {
		var line string

		if i < length-1 {
			line = fmt.Sprintf(
				"For want of a %s the %s was lost.",
				rhyme[i],
				rhyme[i+1],
			)
		} else {
			line = fmt.Sprintf("And all for the want of a %s.", rhyme[0])
		}

		lines = append(lines, line)
	}
	return lines
}
