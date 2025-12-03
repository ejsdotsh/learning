// Package triangle determines if a triangle is equilateral, isosceles, or
// scalene, or not a triangle
package triangle

// Kind is an integer to represent the constant value.
type Kind int

const (
	NaT Kind = iota // not a triangle
	Equ             // equilateral
	Iso             // isosceles
	Sca             // scalene
)

// KindFromSides returns what kind of triangle it is from the three values.
func KindFromSides(a, b, c float64) Kind {
	var k Kind
	switch {
	case a <= 0 || b <= 0 || c <= 0:
		k = NaT // this is the default value
	case a == b && b == c:
		k = Equ
	case (a == b || a == c || b == c) && isTriangle(a, b, c):
		k = Iso
	case (a != b && a != c && b != c) && isTriangle(a, b, c):
		k = Sca
	}

	return k
}

func isTriangle(a, b, c float64) bool {
	return (a+b >= c) && (b+c >= a) && (a+c >= b)
}
