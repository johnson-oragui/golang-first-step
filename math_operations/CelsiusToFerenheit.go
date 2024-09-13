package math_operations

import "fmt"

/*
8. Temperature Converter
Problem: Write a function celsiusToFahrenheit(celsius float64) float64 that converts a
temperature from Celsius to Fahrenheit. Test the function with different temperatures
and print the results.

Example Output:
25°C is 77°F
0°C is 32°F
*/

func CelsiusToFahrenheit(celsius float64) string {
	fahrenheit := ((celsius * 9) / 5 ) + 32

	return fmt.Sprintf("%2.f°C is %.2f°F", celsius, fahrenheit)
}
