package romannumerals

import "fmt"

var numerals = map[string][]string{
	"ones":      {"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"},
	"tens":      {"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"},
	"hundreds":  {"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"},
	"thousands": {"", "M", "MM", "MMM"},
}

func ToRomanNumeral(input int) (string, error) {
	if input <= 0 || input > 3000 {
		return "", fmt.Errorf("%d out of range", input)
	}
	
	roman := make([]byte, 0)
	for _, i := range []string{"ones", "tens", "hundreds", "thousands"} {
		quot := input / 10
		rem := input % 10
		roman = append([]byte(numerals[i][rem]), roman...)
		input = quot
	}
	return string(roman), nil
}
