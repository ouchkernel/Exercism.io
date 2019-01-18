package leap

// IsLeapYear returns a bool depending on if year is a leap or not
func IsLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}
