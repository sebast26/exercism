package dna

import (
	"fmt"
	"strings"
)

// Histogram is a mapping from nucleotide to its count in given DNA.
type Histogram map[rune]int

// DNA is a list of nucleotides.
type DNA string

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
func (d DNA) Counts() (Histogram, error) {
	if strings.Count(string(d), "A")+
		strings.Count(string(d), "G")+
		strings.Count(string(d), "T")+
		strings.Count(string(d), "C") != len(d) {
		return nil, fmt.Errorf("invalid DNA: %s", string(d))
	}
	var h Histogram = map[rune]int{'A': 0, 'G': 0, 'T': 0, 'C': 0}
	for _, r := range d {
		h[r]++
	}
	return h, nil
}
