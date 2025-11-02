package raindrops

import (
	"strconv"
	"strings"
)

func Convert(number int) string {
	res := []string{}

	if number%3 == 0 {
		res = append(res, "Pling")
	}
	if number%5 == 0 {
		res = append(res, "Plang")
	}
	if number%7 == 0 {
		res = append(res, "Plong")
	}

	if len(res) > 0 {
		return strings.Join(res, "")
	}
	return strconv.Itoa(number)
}
