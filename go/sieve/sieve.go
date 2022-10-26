package sieve

func Sieve(limit int) []int {
	s := limit + 1
	sieve := make([]int, s)
	for i := range sieve {
		sieve[i] = 1
	}

	for i := 2; i < len(sieve); i++ {
		for j := i; j < len(sieve); j++ {
			if i*j < s {
				sieve[i*j] = 0
			} else {
				break
			}
		}
	}

	var out []int
	for i := 2; i < len(sieve); i++ {
		if sieve[i] == 1 {
			out = append(out, i)
		}
	}
	return out
}
