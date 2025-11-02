package strain

// Implement the "Keep" and "Discard" function in this file.

// You will need typed parameters (aka "Generics") to solve this exercise.
// They are not part of the Exercism syllabus yet but you can learn about
// them here: https://go.dev/tour/generics/1

// Keep returns a new collection containing those elements where the predicate
// is true.
func Keep[T any](in []T, predicate func(T) bool) []T {
	var res = []T{}

	for _, c := range in {
		if predicate(c) {
			res = append(res, c)
		}
	}

	return res
}

// Discard returns a new collection containing those elements where the
// predicate is false.
func Discard[T any](in []T, predicate func(T) bool) []T {
	var res = []T{}

	for _, c := range in {
		if !predicate(c) {
			res = append(res, c)
		}
	}

	return res
}
