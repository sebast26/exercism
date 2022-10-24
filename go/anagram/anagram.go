package anagram

import (
	"sort"
	"strings"
)

type byteSlice []byte

func Detect(subject string, candidates []string) []string {
	key := byteSlice(strings.ToLower(subject))
	sort.Slice(key, func(i, j int) bool {
		return key[i] < key[j]
	})

	var anagrams []string
	for _, candidate := range candidates {
		cand := byteSlice(strings.ToLower(candidate))
		sort.Slice(cand, func(i, j int) bool {
			return cand[i] < cand[j]
		})
		if string(key) == string(cand) && strings.ToLower(candidate) != strings.ToLower(subject) {
			anagrams = append(anagrams, candidate)
		}
	}
	return anagrams
}
