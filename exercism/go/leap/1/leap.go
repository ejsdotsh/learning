package leap

// IsLeapYear returns `true` if the number is a leap year.
func IsLeapYear(year int) bool {
	if year%4 == 0 {
		if year%400 == 0 {
			return true
		}
		if year%100 != 0 {
			return true
		}
	}
	return false
}
