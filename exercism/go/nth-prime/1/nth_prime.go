// Package prime returns the Nth prime number given an input, `N`.
package prime

import (
	"fmt"
	"math"
)

// Nth returns the nth prime number. An error must be returned if the nth prime
// number can't be calculated ('n' is equal or less than zero)
func Nth(n int) (int, error) {
	if n < 1 {
		return 0, fmt.Errorf("number must be greater than 1")
	}

	nth := 0
	num := 0
	for {
		num++
		// sieve/filter prime multipliers to slightly improve efficiency
		if (num > 2 && num%2 == 0) ||
			(num > 3 && num%3 == 0) ||
			(num > 5 && num%5 == 0) ||
			(num > 7 && num%7 == 0) ||
			(num > 11 && num%11 == 0) ||
			(num > 13 && num%13 == 0) {
			continue
		}
		if isPrime(num) {
			nth++
		}
		if nth == n {
			return num, nil
		}
	}
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	sqrt := int(math.Sqrt(float64(n)))
	for i := 2; i <= sqrt; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
