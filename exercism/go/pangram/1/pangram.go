package pangram

import "strings"

func IsPangram(input string) bool {
	foundLetters := make(map[rune]bool)
	lowerInput := strings.ToLower(input)

	for _, c := range lowerInput {
		if c >= 'a' && c <= 'z' {
			foundLetters[c] = true
		}
	}

	return len(foundLetters) == 26
}
