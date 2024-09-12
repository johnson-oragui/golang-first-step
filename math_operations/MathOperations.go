package math_operations

import "fmt"

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
func Add(a int, b int) int {

	return a + b
}

func Substract(a int, b int) int {
	return a - b
}

func Multiply(a int, b int) int {
	
	return a * b
}

func Quotient(a int, b int) int {
	if b == 0 {
		fmt.Println("Cannot divide by Zero")
		return 0
	}
	return a / b
}
