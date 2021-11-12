package diffsquares

import "math"

// SquareOfSum
// https://brilliant.org/wiki/sum-of-n-n2-or-n3/
func SquareOfSum(n int) int {
	s := n * (n + 1) / 2
	return int(math.Pow(float64(s), 2))
}

// SumOfSquares
// https://brilliant.org/wiki/sum-of-n-n2-or-n3/
func SumOfSquares(n int) int {
	return n * (n + 1) * (2*n + 1) / 6
}

func Difference(n int) int {
	return int(math.Abs(float64(SumOfSquares(n) - SquareOfSum(n))))
}
