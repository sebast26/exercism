package pangram

import (
	"unicode"
)

func IsPangram(input string) bool {
	var histogram = map[rune]int{}
	for r := 'a'; r <= 'z'; r++ {
		histogram[r] = 0
	}

	for _, r := range input {
		l := unicode.ToLower(r)
		if !unicode.IsLetter(l) {
			continue
		}
		histogram[l]++
	}

	for _, v := range histogram {
		if v == 0 {
			return false
		}
	}

	return true
}
