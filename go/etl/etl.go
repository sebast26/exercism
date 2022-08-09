package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
	ret := map[string]int{}
	for score, letters := range in {
		for _, letter := range letters {
			l := strings.ToLower(letter)
			ret[l] = score
		}
	}
	return ret
}
