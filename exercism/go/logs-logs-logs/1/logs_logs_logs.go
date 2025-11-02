package logs

import (
	"unicode/utf8"
)

// Application identifies the application emitting the given log.
func Application(log string) string {
	for _, ch := range log {
		if ch == rune('❗') {
			return "recommendation"
		}
		if ch == rune('🔍') {
			return "search"
		}
		if ch == rune('☀') {
			return "weather"
		}
	}

	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	r := []rune{}
	for _, ch := range log {
		if ch == oldRune {
			r = append(r, newRune)
		} else {
			r = append(r, ch)
		}
	}
	return string(r)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	numberOfRunes := utf8.RuneCountInString(log)
	return numberOfRunes <= limit
}
