// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package triangle should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package triangle

import "math"

// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind int

const (
	NaT = iota // not a triangle
	Equ // equilateral
	Iso // isosceles
	Sca // scalene
	Deg // degenerate
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	var k Kind = NaT

	if math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0) {
		k = NaT
	} else if math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) {
		k = NaT
	} else if a <= 0 || b <= 0 || c <= 0 {
		k = NaT
	} else if (a+b < c) || (a+c < b) || (b+c < a) {
		k = NaT
	} else if (a+b == c) || (a+c == b) || (b+c == a) {
		k = Deg
	} else if a == b && b == c {
		k = Equ
	} else if a == b || b == c || c == a {
		k = Iso
	} else if a != b && b != c && c != a {
		k = Sca
	}

	return k
}
