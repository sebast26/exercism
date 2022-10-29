package summultiples

func SumMultiples(limit int, divisors ...int) int {
	t := make([]int, limit+1)
	for _, div := range divisors {
		if div == 0 {
			continue
		}
		for i := 1; i*div < limit; i++ {
			t[i*div] = 1
		}
	}

	s := 0
	for i, v := range t {
		if v == 1 {
			s += i
		}
	}
	return s
}
