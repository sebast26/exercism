package wordcount

import (
	"regexp"
	"strings"
)

type Frequency map[string]int
var (
	punct = regexp.MustCompile(`[,.!$@%^&:]`)
	quot = regexp.MustCompile(`^'|'$|'\s|\s'`)
)

func WordCount(phrase string) Frequency {
	p := punct.ReplaceAllLiteralString(phrase, " ")
	p = quot.ReplaceAllString(p, " ")
	words := strings.Fields(p)

	var freq = Frequency{}
	for _, word := range words {
		w := strings.ToLower(word)
		if _, ok := freq[w]; !ok {
			freq[w] = 0
		}
		freq[w]++
	}

	return freq
}
