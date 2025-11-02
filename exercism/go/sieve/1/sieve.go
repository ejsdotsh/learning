package sieve

func Sieve(limit int) []int {
	primeList := make([]bool, limit+1)
	for i := 2; i <= limit; i++ {
		primeList[i] = true
	}

	for p := 2; p*p <= limit; p++ {
		if primeList[p] == true {
			for i := p * p; i <= limit; i += p {
				primeList[i] = false
			}
		}
	}

	var primeNumbers []int
	for pn := 2; pn <= limit; pn++ {
		if primeList[pn] {
			primeNumbers = append(primeNumbers, pn)
		}
	}

	return primeNumbers
}
