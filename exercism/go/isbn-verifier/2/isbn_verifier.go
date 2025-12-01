// Package isbn checks the validity of a provided ISBN.
package isbn

import (
	"strings"
	"unicode"
)

// IsValidISBN checks the validity of a provided ISBN and returns a boolean.
func IsValidISBN(isbn string) bool {
	s := strings.ReplaceAll(isbn, "-", "")
	if len(s) != 10 {
		return false
	}
	sum := 0
	for i, c := range s {
		switch {
		// if it is a digit, add it to the running sum
		case unicode.IsDigit(c):
			sum += int(c-'0') * (10 - i)
		// if the last digit is 'X', add 10
		case (i == 9 && (c == 'x' || c == 'X')):
			sum += (10 - i) * 10
		// anything else is invalid
		default:
			return false
		}
	}

	return sum%11 == 0
}
