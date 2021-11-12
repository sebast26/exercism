package hamming

import "errors"

// Distance calculates the Hamming Distance between two DNA strands.
func Distance(a, b string) (int, error) {
	ra := []rune(a)
	rb := []rune(b)
	if len(ra) != len(rb) {
		return 0, errors.New("works only on strings with equal length")
	}

	d := 0
	for i := range ra {
		if ra[i] != rb[i] {
			d++
		}
	}
	return d, nil
}
