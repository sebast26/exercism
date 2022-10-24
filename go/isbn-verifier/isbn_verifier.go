package isbn

import (
	"fmt"
	"strconv"
	"strings"
)

// IsValidISBN checks if isbn string is valid ISBN number.
func IsValidISBN(isbn string) bool {
	nums, err := isbnNumbers(isbn)
	if err != nil {
		return false
	}

	// Using modulo distributive property
	// (a + b) mod n = [(a mod n) + (b mod n)] mod n
	sum := 0
	for i, num := range nums {
		x := (10 - i) * num
		sum += x % 11
	}

	return sum%11 == 0
}

func isbnNumbers(s string) ([]int, error) {
	isbn := strings.ReplaceAll(s, "-", "")

	if len(isbn) != isbnLength {
		return nil, fmt.Errorf("invalid isbn, wanted %d got %d characters", isbnLength, len(isbn))
	}

	var nums []int
	for i := 0; i < len(isbn)-1; i++ {
		c := isbn[i]
		n, err := strconv.Atoi(fmt.Sprintf("%c", c))
		if err != nil {
			return nil, err
		}
		nums = append(nums, n)
	}

	// handle check digit
	n, err := checkDigit(rune(isbn[checkDigitPos]))
	if err != nil {
		return nil, err
	}
	nums = append(nums, n)

	return nums, nil
}

func checkDigit(c rune) (n int, err error) {
	if c == specialCheckChar {
		return specialCheckValue, nil
	}
	n, err = strconv.Atoi(fmt.Sprintf("%c", c))
	return
}

const (
	isbnLength        = 10
	checkDigitPos     = 9
	specialCheckChar  = 'X'
	specialCheckValue = 10
)
