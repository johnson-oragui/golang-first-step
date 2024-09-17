package math_operations

import "errors"

/*
Problem: Write a program that declares two integer variables, x and y, and performs the following operations:

Initialize x to 10 and y to 20.
Calculate the sum, difference, product, and quotient of x and y.
Print the results in a readable format.
Example Output:
Sum: 30
Difference: -10
Product: 200
Quotient: 0.5
*/
func Add(a, b int) int {

	return a + b
}

func Substract(a, b int) int {
	return a - b
}

func Multiply(a, b int) int {
	
	return a * b
}

func Quotient(a, b int) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	} else {
		return float64(a) / float64(b), nil
	}
}
