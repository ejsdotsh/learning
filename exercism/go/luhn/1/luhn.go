package luhn

import (
	"slices"
	"strconv"
	"strings"
)

func Valid(id string) bool {
	noSpaces := strings.ReplaceAll(id, " ", "")
	if len(noSpaces) <= 1 {
		return false
	}
	numList := strings.Split(noSpaces, "")
	slices.Reverse(numList)
	numSum := 0

	for idx, num := range numList {
		n, err := strconv.Atoi(num)
		if err != nil {
			return false
		}
		if idx%2 == 0 {
			numSum += n
		} else if 2*n > 9 {
			numSum += (2*n - 9)
		} else {
			numSum += 2 * n
		}
	}

	return numSum%10 == 0
}
