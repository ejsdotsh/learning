// Package anagram takes a target word and one or more candidate words, and
// returns, in original order, any candidate words that are anagrams of the
// target word.
package anagram

import (
	"slices"
	"strings"
	"unicode/utf8"
)

// Detect takes a target word and a list of candidate words, and returns, in
// order, any candidate words that are anagrams of the target words.
func Detect(subject string, candidates []string) []string {
	anagrams := []string{}

	for _, word := range candidates {
		if strings.EqualFold(subject, word) {
			continue
		}
		if utf8.RuneCountInString(word) != utf8.RuneCountInString(subject) {
			continue
		}
		if isAnagram(subject, word) {
			anagrams = append(anagrams, word)
		}
	}

	return anagrams
}

// isAnagram returns a bool if two words are anagrams.
func isAnagram(w1, w2 string) bool {
	anagram := false
	r1, r2 := []rune(strings.ToLower(w1)), []rune(strings.ToLower(w2))
	slices.Sort(r1)
	slices.Sort(r2)
	if slices.Compare(r1, r2) == 0 {
		anagram = true
	}

	return anagram
}
