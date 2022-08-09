package strand

import "strings"

var complements = map[rune]rune{
	'G': 'C',
	'C': 'G',
	'T': 'A',
	'A': 'U',
}
func ToRNA(dna string) string {
	var sb strings.Builder
	for _, c := range dna {
		t, _ := complements[c]
		sb.WriteRune(t)
	}
	return sb.String()
}
