// Package isbn checks the validity of a provided ISBN.
package isbn

import (
	"strings"
)

// IsValidISBN checks the validity of a provided ISBN and returns a boolean.
func IsValidISBN(isbn string) bool {
	s := strings.ReplaceAll(isbn, "-", "")
	if len(s) != 10 {
		return false
	}
	d := make([]int, 10)
	for i, c := range s {
		switch {
		case (0 <= int(c-'0') && int(c-'0') <= 9):
			d[i] = int(c - '0')
		case (i == 9 && (c == 'x' || c == 'X')):
			d[i] = 10
		default:
			return false
		}
	}

	return (d[0]*10+d[1]*9+d[2]*8+d[3]*7+d[4]*6+d[5]*5+d[6]*4+d[7]*3+d[8]*2+d[9]*1)%11 == 0
}
