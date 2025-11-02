package diffsquares

func SquareOfSum(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	sos := sum * sum
	return sos
}

func SumOfSquares(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i * i
	}
	sos := sum
	return sos
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
