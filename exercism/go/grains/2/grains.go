package grains

import "errors"

// Square uses bitshifting to efficiently return the number of grains on the
// given square. If the number is less-then 1 or greater-than 64, it returns
// an error.
func Square(number int) (uint64, error) {
	if number < 1 || number > 64 {
		return uint64(0), errors.New("Number must be from 1 to 64, inclusive.")
	}
	return 1 << uint64(number-1), nil
}

// maxValue uses bitshifting to determine the total number of grains once, at
// compile time.
const maxValue = (1 << 64) - 1

func Total() uint64 {
	return uint64(maxValue)
}
