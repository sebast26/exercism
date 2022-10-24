package encode

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type byteSlice []byte

func RunLengthEncode(input string) string {
	bs := byteSlice(input)
	var char byte
	var count int
	var sb strings.Builder
	for _, b := range bs {
		if char == b {
			count++
			continue
		}
		if char != 0 {
			encodeOut(&sb, count, char)
		}
		char = b
		count = 1
	}

	if char != 0 {
		encodeOut(&sb, count, char)
	}

	return sb.String()
}

func encodeOut(sb *strings.Builder, count int, char byte) {
	if count > 1 {
		sb.WriteString(fmt.Sprintf("%d%c", count, char))
	} else {
		sb.WriteString(fmt.Sprintf("%c", char))
	}
}

func RunLengthDecode(input string) string {
	splits := split(input)
	fmt.Println(splits)

	var sb strings.Builder
	for _, split := range splits {
		if len(split) == 1 {
			sb.WriteString(split)
			continue
		}
		i, _ := strconv.Atoi(split[0 : len(split)-2])
		decodeTo(&sb, i, split[len(split)-1:])
	}

	return sb.String()
}

func split(input string) []string {
	var splits []string
	for _, char := range input {
		if unicode.IsLetter(char) {
			splits = append(splits, string(char))
		}
	}
	return splits
}

func decodeTo(sb *strings.Builder, count int, s string) {
	for i := 0; i < count; i++ {
		sb.WriteString(s)
	}
}
