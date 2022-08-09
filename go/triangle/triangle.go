// Package triangle should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package triangle

// Kind of triangle
type Kind int

const (
	NaT Kind = iota // not a triangle
	Equ             // equilateral
	Iso             // isosceles
	Sca             // scalene
)

// KindFromSides detects kind of triangle.
func KindFromSides(a, b, c float64) Kind {
	if a < 0 || b < 0 || c < 0 || a >= b+c || b >= a+c || c >= a+b {
		return NaT
	}

	if (a == b) && (a == c) {
		return Equ
	}
	if a == b || b == c || a == c {
		return Iso
	}
	return Sca
}
