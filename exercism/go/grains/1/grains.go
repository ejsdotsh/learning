package grains

import "errors"

func Square(number int) (uint64, error) {
	if number < 1 || number > 64 {
		return 0, errors.New("Number must be from 1 to 64.")
	}
	return 1 << (number - 1), nil
}

func Total() uint64 {
	var total uint64 = 0
	for i := 1; i < 65; i++ {
		num, err := Square(i)
		if err != nil {
			return 0
		}
		total += num
	}
	return total
}
