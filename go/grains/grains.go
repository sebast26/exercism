package grains

import (
	"errors"
)

// Square shows how many grains were on a given square
func Square(number int) (uint64, error) {
	if number < 1 || number > 64 {
		return 0, errors.New("error")
	}
	return 1 << (number - 1), nil
}

// Total shows the total number of grains on the chessboard
func Total() uint64 {
	return 1<<64 - 1
}
