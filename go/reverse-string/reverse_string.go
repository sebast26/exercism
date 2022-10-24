package reverse

import (
	"unicode/utf8"
)

func Reverse(input string) string {
	l, i := utf8.RuneCountInString(input), 0
	rev := make([]rune, l)
	for _, r := range input {
		rev[l-1-i] = r
		i++
	}
	return string(rev)
}
