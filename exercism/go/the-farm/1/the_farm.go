package thefarm

import (
	"errors"
	"fmt"
)

// DivideFood accepts a `FodderCalculator` and a number of cows (int), and
// returns a float64 and an error.
func DivideFood(fc FodderCalculator, numCows int) (float64, error) {
	totalFodder, err := fc.FodderAmount(numCows)
	if err != nil {
		return 0, err
	}

	fatteningFactor, err := fc.FatteningFactor()
	if err != nil {
		return 0, err
	}

	foodPerCow := (totalFodder * fatteningFactor) / float64(numCows)

	return foodPerCow, nil
}

// ValidateInputAndDivideFood is a very simple value-checking wrapper around
// DivideFood.
func ValidateInputAndDivideFood(fc FodderCalculator, numCows int) (float64, error) {
	if numCows <= 0 {
		return 0, errors.New("invalid number of cows")
	}

	return DivideFood(fc, numCows)
}

// InvalidCowsError is a custom error type to improve error handling by adding
// a custom message.
type InvalidCowsError struct {
	numCows   int
	customMsg string
}

func (invCowsErr *InvalidCowsError) Error() string {
	return fmt.Sprintf("%v cows are invalid: %s", invCowsErr.numCows, invCowsErr.customMsg)
}

// ValidateNumberOfCows returns custom messages is the number of cows is invalid.
func ValidateNumberOfCows(numCows int) error {
	if numCows < 0 {
		return &InvalidCowsError{
			numCows:   numCows,
			customMsg: "there are no negative cows",
		}
	}
	if numCows == 0 {
		return &InvalidCowsError{
			numCows:   numCows,
			customMsg: "no cows don't need food",
		}
	}
	return nil
}
