package darts

func Score(x, y float64) int {
	r := x*x + y*y
	switch {
	case r <= 1:
		return 10
	case r <= 25:
		return 5
	case r <= 100:
		return 1
	default:
		return 0
	}
}
