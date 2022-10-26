package armstrong

import (
	"fmt"
	"math"
)

func IsNumber(n int) bool {
	pow, tmp, s := len(fmt.Sprintf("%d", n)), n, 0
	for i := pow - 1; i >= 0; i-- {
		s += int(math.Pow(float64(tmp%10), float64(pow)))
		tmp /= 10
	}
	return s == n
}
