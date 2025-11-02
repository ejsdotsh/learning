package isogram

import (
	"slices"
	"strings"
)

func IsIsogram(word string) bool {
	var found []rune
	lw := strings.ToLower(word)
	for _, l := range lw {
		if l == ' ' || l == '-' {
			continue
		}
		if slices.Contains(found, l) {
			return false
		}
		found = append(found, l)
	}
	return true
}
