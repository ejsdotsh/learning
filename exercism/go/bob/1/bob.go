// bob returns some limited responses that could be attributable to a teenager.
package bob

import (
	"strings"
	"unicode"
)

// isUpper returns a bool if the string is all upper case
func isUpper(s string) bool {
	letters := 0
	upper := 0
	for _, ch := range s {
		if !unicode.IsLetter(ch) {
			continue
		}
		if !unicode.IsUpper(ch) && unicode.IsLetter(ch) {
			letters++
		}
		if unicode.IsUpper(ch) && unicode.IsLetter(ch) {
			upper++
			letters++
		}
	}
	if upper > 0 {
		return upper == letters
	}
	return false
}

// Hey returns Bob's limited responses.
func Hey(remark string) string {
	r := (strings.TrimSpace(remark))
	switch {
	case isUpper(r):
		if strings.HasSuffix(r, "?") {
			return "Calm down, I know what I'm doing!"
		}
		return "Whoa, chill out!"
	case strings.HasSuffix(r, "?"):
		return "Sure."
	case len(r) == 0:
		return "Fine. Be that way!"
	default:
		return "Whatever."
	}
}
