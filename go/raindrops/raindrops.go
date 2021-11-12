package raindrops

import (
	"strconv"
)

// Convert a number into Pling/Plang/Plong string or number string
func Convert(number int) string {
	var ret string
	if number%3 == 0 {
		ret += "Pling"
	}
	if number%5 == 0 {
		ret += "Plang"
	}
	if number%7 == 0 {
		ret += "Plong"
	}
	if ret != "" {
		return ret
	}

	return strconv.Itoa(number)
}
