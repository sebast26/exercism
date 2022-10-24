package lsproduct

import (
	"errors"
	"strconv"
	"strings"
)

func LargestSeriesProduct(digits string, span int) (int64, error) {
	if len(digits) < span {
		return 0, errors.New("illegal")
	}
	if span < 0 {
		return 0, errors.New("illegal")
	}
	if strings.IndexFunc(digits, isNotDigit) != -1 {
		return 0, errors.New("illegal")
	}
	if span == 0 {
		return 1, nil
	}

	m := map[string]int64{}
	for i := 0; i < len(digits)-span+1; i++ {
		key := digits[i : i+span]
		m[key] = computeProduct(key)
	}

	max := int64(0)
	for _, v := range m {
		if v > max {
			max = v
		}
	}

	return max, nil
}

func computeProduct(digits string) int64 {
	p := int64(1)
	for _, d := range digits {
		n, _ := strconv.Atoi(string(d))
		p *= int64(n)
	}
	return p
}

var isNotDigit = func(c rune) bool { return c < '0' || c > '9' }
