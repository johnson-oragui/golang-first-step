package math_operations

import (
	"errors"
	"fmt"
)

/*
9. Simple Calculator with Switch Case
Problem: Write a program that takes two numbers and an
operator (like +, -, *, /) from the user. Use a switch
statement to perform the operation based on the operator
and print the result.

Example Output:
Enter first number: 8
Enter operator (+, -, *, /): *
Enter second number: 7
Result: 56
*/


func Calculator() (string, error) {
	var first, second, operation int
	var operator string

	fmt.Print("Enter first number: ")
	fmt.Scanf("%d", &first)

	fmt.Print("Enter operator (+, -, *, /): ")
	fmt.Scanf("%s", &operator)

	fmt.Print("Enter second number:")
	fmt.Scanf("%d", &second)


	switch operator {
	case "+": operation = first + second
	case "-": operation = first - second
	case "*": operation = first * second
	case "/":
		if second == 0 {
			return "", errors.New("Error: Division by zero is undefined")
		}
		operation = first / second
	default: return "", errors.New("Error: Invalid operator")
		
	}

	result := fmt.Sprintf("Result: %d", operation)

	return result, nil
}
