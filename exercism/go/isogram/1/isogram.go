package isogram

import (
	"slices"
	"unicode"
)

func IsIsogram(word string) bool {
	var found []rune
	for _, l := range word {
		if l == ' ' || l == '-' {
			continue
		}
		if slices.Contains(found, unicode.ToLower(l)) {
			return false
		}
		found = append(found, unicode.ToLower(l))
	}
	return true
}
