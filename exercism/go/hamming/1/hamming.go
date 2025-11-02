package hamming

import (
	"errors"
	"strings"
)

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("The strings are not the same length.")
	}

	var hammingDist int
	a1 := strings.Split(a, "")
	b1 := strings.Split(b, "")

	for i := range len(a1) {
		if a1[i] != b1[i] {
			hammingDist++
		}
	}

	return hammingDist, nil
}
