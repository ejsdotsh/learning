package darts

import "math"

func Score(x, y float64) int {
	r := math.Sqrt(x*x + y*y)
	switch {
	case r <= float64(1):
		return 10
	case r <= float64(5):
		return 5
	case r <= float64(10):
		return 1
	default:
		return 0
	}
}
