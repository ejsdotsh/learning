package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
	res := make(map[string]int, 0)

	for pointVal, letters := range in {
		for _, letter := range letters {
			res[strings.ToLower(letter)] = pointVal
		}
	}

	return res
}
