package rotationalcipher

import (
	"strings"
	"unicode"
)

func RotationalCipher(plain string, shiftKey int) string {
	var sb strings.Builder
	for _, c := range plain {
		if unicode.IsSpace(c) || unicode.IsPunct(c) || unicode.IsDigit(c) {
			sb.WriteRune(c)
			continue
		}
		base := 'A'
		if c >= 'a' && c <= 'z' {
			base = 'a'
		}

		idx := (c + int32(shiftKey) - base) % 26
		sb.WriteRune(base + idx)
	}
	return sb.String()
}
