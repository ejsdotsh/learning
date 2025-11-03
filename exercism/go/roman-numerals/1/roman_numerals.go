package romannumerals

import (
	"errors"
	"strings"
)

// ToRomanNumeral converts an integer between 1 and 3999 to its equivalent in
// Roman numerals.
func ToRomanNumeral(input int) (string, error) {
	if input < 1 || 3999 < input {
		return "", errors.New("Input must be a number between 1 and 3999, inclusive.")
	}

	// Arabic numbers
	an := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	// Roman numbers
	rn := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	// a slice of string to collect the Roman numerals
	numerals := []string{}

	// iterate over the Arabic number slice
	for i := range len(an) {
		// if the value of 'input' is greater than the Arabic number at that
		// index, subtract the number from 'input' and append the corresponding
		// Roman number to the return slice
		for input >= an[i] {
			input -= an[i]
			numerals = append(numerals, rn[i])
		}
	}

	return strings.Join(numerals, ""), nil
}
