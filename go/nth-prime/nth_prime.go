package prime

import "errors"

// Nth returns the nth prime number. An error must be returned if the nth prime number can't be calculated ('n' is equal or less than zero)
func Nth(n int) (int, error) {
	if n <= 0 {
		return 0, errors.New("nth prime number can't be calculated, n is equal or less than zero")
	}

	totalPrimes := 0
	sizeFactor := 2
	s := n * sizeFactor

	primes := initPrimes(s)
	for totalPrimes < n {
		primes = initPrimes(s)
		primes = getPrimes(primes, s)
		totalPrimes = sumPrimes(primes)
		sizeFactor++
		s = n * sizeFactor
	}
	return countPrimes(primes, n), nil
}

// getPrimes generates primes using the Sieve of Eratosthenes.
//
//	Includes the optimization where for every prime p, only factors p >= p^2
//	are verified.
//	The list of primes is represented with a int slice. Each index corresponds
//	to an integer in the list. A value of "1" at the index location indicates
//	the integer is a prime.
func getPrimes(primes []int, s int) []int {
	for i := 2; i < s; i++ {
		if primes[i] == 1 {
			for j := i; j < s; j++ {
				if i*j <= s {
					primes[i*j] = 0
				} else {
					break
				}
			}
		}
	}
	return primes
}

func initPrimes(s int) []int {
	p := make([]int, s+1)
	for i := range p {
		p[i] = 1
	}
	return p
}

// countPrimes returns the n-th prime represented by the index of the n-th "1" in the slice.
func countPrimes(primes []int, n int) int {
	count := 0
	for i := 2; i < len(primes); i++ {
		count += primes[i]
		if count == n {
			return i
		}
	}
	return 0
}

func sumPrimes(primes []int) int {
	sum := 0
	for _, p := range primes[2:] {
		sum += p
	}
	return sum
}
