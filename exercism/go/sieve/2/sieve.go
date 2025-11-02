package sieve

func Sieve(limit int) []int {
	numList := make([]bool, limit+1)
	primeList := make([]int, limit/2)
	primeIndex := 0

	for i := 2; i <= limit; i++ {
		if !numList[i] {
			primeList[primeIndex] = i
			primeIndex++
			for j := i + i; j <= limit; j += i {
				numList[j] = true
			}
		}
	}

	return primeList[:primeIndex]
}
