// Package lsproduct uses the largest series product to find patterns in a
// sequence of input digits.
package lsproduct

import (
	"errors"
	"strconv"
	"unicode"
)

var (
	ErrInvalidSpan  = errors.New("span must be greater than 1; invalid span size")
	ErrSpanLength   = errors.New("span must be smaller than string length")
	ErrInvalidInput = errors.New("digits input must only contain digits")
)

// LargestSeriesProduct finds the largest series product given an input sequence
// of digits and a span size.
func LargestSeriesProduct(digits string, span int) (int64, error) {
	if span < 1 {
		return -1, ErrInvalidSpan
	}
	if span > len(digits) {
		return -1, ErrSpanLength
	}

	var largestSum int64 = 0
	digitsLen := len(digits)

	for i, d := range digits {
		if !unicode.IsDigit(d) {
			return -1, ErrInvalidInput
		}
		if i+span > digitsLen {
			break
		}

		runningSum := findProduct(digits[i : i+span])
		if runningSum > largestSum {
			largestSum = runningSum
		}
	}

	return largestSum, nil
}

func findProduct(digits string) int64 {
	var product = 1
	for _, d := range digits {
		num, err := strconv.Atoi(string(d))
		if err != nil {
			return -1
		}
		product *= num
	}
	return int64(product)
}
