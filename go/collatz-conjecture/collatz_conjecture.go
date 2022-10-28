package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
		return 0, errors.New("invalid n")
	}
	return step(n), nil
}

func step(n int) int {
	if n == 1 {
		return 0
	} else if n%2 == 0 {
		return step(n/2) + 1
	} else {
		return step(3*n+1) + 1
	}
}
