// Package leap implements functions for working with leap years.
package leap

// IsLeapYear determines if a given year is or is not a leap year.
func IsLeapYear(year int) bool {
	return divisibleBy(year, 4) &&
		(!divisibleBy(year, 100) || divisibleBy(year, 400))
}

func divisibleBy(year, divisor int) bool {
	return year%divisor == 0
}
