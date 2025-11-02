package sorting

import (
	"fmt"
	"strconv"
)

// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
	return fmt.Sprintf("This is the number %.1f", f)
}

type NumberBox interface {
	Number() int
}

// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nBox NumberBox) string {
	return fmt.Sprintf(
		"This is a box containing the number %.1f",
		float64(nBox.Number()),
	)
}

type FancyNumber struct {
	n string
}

func (i FancyNumber) Value() string {
	return i.n
}

type FancyNumberBox interface {
	Value() string
}

// ExtractFancyNumber should return the integer value for a FancyNumber
// and 0 if any other FancyNumberBox is supplied.
func ExtractFancyNumber(fnBox FancyNumberBox) int {
	fn, ok := fnBox.(FancyNumber)
	if !ok {
		return 0
	}
	fancyNum, err := strconv.Atoi(fn.n)
	if err != nil {
		return 0
	}

	return fancyNum
}

// DescribeFancyNumberBox should return a string describing the FancyNumberBox.
func DescribeFancyNumberBox(fnBox FancyNumberBox) string {
	return fmt.Sprintf(
		"This is a fancy box containing the number %d.0",
		ExtractFancyNumber(fnBox),
	)
}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i any) string {
	switch v := i.(type) {
	case int:
		return DescribeNumber(float64(v))
	case float64:
		return DescribeNumber(v)
	case NumberBox:
		return DescribeNumberBox(v)
	case FancyNumberBox:
		return DescribeFancyNumberBox(v)

	default:
		return "Return to sender"
	}
}
