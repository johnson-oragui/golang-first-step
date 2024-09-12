package math_operations

/*
7. Fibonacci Sequence
Problem: Write a function fibonacci(n int) []int that returns the first
n numbers of the Fibonacci sequence. Print the sequence.

Example Output:
First 5 Fibonacci numbers: [0 1 1 2 3]

*/

func fibonacci(n int) []int {
	list := [...]int{1, 2, 3, 4, 5}
	return list[1:n]
}
