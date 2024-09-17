package math_operations

import "fmt"

/*
10. Determine Leap Year
Problem: Write a function isLeapYear(year int) bool that determines if a
given year is a leap year. A year is a leap year if it is divisible by 4,
but not divisible by 100 unless it is also divisible by 400.
Test the function with various years and print the results.

Example Output:
Year 2000 is a leap year: true
Year 1900 is a leap year: false
*/

func IsLeapYear(year int) string {
	var leap bool

	if year % 400 == 0 {
		leap = true
	} else if year % 100 == 0 {
		leap = false
	} else if year % 4 == 0 {
		leap = true
	} else {
		leap = false
	}
	result := fmt.Sprintf("Year %d is a leap year: %v", year, leap)
	return result
}
