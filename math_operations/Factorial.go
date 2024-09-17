package math_operations

/*
3. Factorial Calculation
Problem: Write a function factorial(n int) int that calculates the
factorial of a non-negative integer n. Use a loop to perform the
calculation. Print the result.

Example Output:
Factorial of 5 is: 120

*/

func Factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	var factorial int = n
	for i := 1; i <n; i++ {
		factorial *= i
	}
	return factorial
}
