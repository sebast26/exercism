package isogram

import (
	"sort"
	"strings"
	"unicode"
)

type sortedRune []rune

func (s sortedRune) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortedRune) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortedRune) Len() int {
	return len(s)
}

// IsIsogram determines if a word or phrase is an isogram.
func IsIsogram(word string) bool {
	r := []rune(strings.ToUpper(word))
	sort.Sort(sortedRune(r))

	for i := 0; i < len(r)-1; i++ {
		if unicode.IsLetter(r[i]) && r[i] == r[i+1] {
			return false
		}
	}
	return true
}
