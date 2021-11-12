package luhn

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

var doubling = []byte{'0', '2', '4', '6', '8', '1', '3', '5', '7', '9'}

// Valid given a number determine whether or not it is valid per the Luhn formula.
func Valid(id string) bool {
	id, err := reformat(id)
	if err != nil {
		return false
	}

	t := make([]byte, len(id))
	for i := len(id) - 1; i >= 0; i-- {
		if i%2 != 0 {
			t[i] = id[i] - '0'
		} else {
			t[i] = doubling[id[i]-'0'] - '0'
		}
	}

	s := byte(0)
	for _, n := range t {
		s += n
	}

	return s%10 == 0
}

func reformat(id string) (string, error) {
	id = strings.ReplaceAll(id, " ", "")

	for _, c := range id {
		if !unicode.IsDigit(c) {
			return "", errors.New("not a digit")
		}
	}

	if len(id) <= 1 {
		return "", errors.New("too short")
	}

	if len(id)%2 != 0 {
		id = fmt.Sprintf("0%s", id)
	}

	return id, nil
}
