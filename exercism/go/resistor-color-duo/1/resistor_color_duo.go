// Package resistorcolorduo takes a list of colors as input, and returns the
// the two digit resistance value.
package resistorcolorduo

import "strconv"

// Value should return the resistance value of a resistor with a given colors.
func Value(colors []string) int {

	res := ""
	for i := range len(colors) {
		switch colors[i] {
		case "black":
			res += "0"
		case "brown":
			res += "1"
		case "red":
			res += "2"
		case "orange":
			res += "3"
		case "yellow":
			res += "4"
		case "green":
			res += "5"
		case "blue":
			res += "6"
		case "violet":
			res += "7"
		case "grey":
			res += "8"
		case "white":
			res += "9"

		}
	}

	// we only want the first two numbers
	value, err := strconv.Atoi(res[:2])
	if err != nil {
		return -1
	}

	return value
}
