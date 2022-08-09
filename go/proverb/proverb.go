// Package proverb should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package proverb

import "fmt"

const (
	verse = "For want of a %s the %s was lost."
	last  = "And all for the want of a %s."
)

// Proverb should have a comment documenting it.
func Proverb(rhyme []string) []string {
	if len(rhyme) == 0 {
		return []string{}
	}

	ret := make([]string, 0, len(rhyme))
	for i := 0; i < len(rhyme)-1; i++ {
		ret = append(ret, fmt.Sprintf(verse, rhyme[i], rhyme[i+1]))
	}
	ret = append(ret, fmt.Sprintf(last, rhyme[0]))

	return ret
}
