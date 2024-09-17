package math_operations

import "fmt"

/*
7. Fibonacci Sequence
Problem: Write a function fibonacci(n int) []int that returns the first
n numbers of the Fibonacci sequence. Print the sequence.

Example Output:
First 5 Fibonacci numbers: [0 1 1 2 3]

*/

func Fibonacci(n int) string {

	list := make([]int, 0, n)
	if n >= 1 {
		list = append(list, 0)
	}
	if n >= 2 {
		list = append(list, 1)
	}

	for i := 2; i <= n; i++ {
		next := list[i-1] + list[i-2]
		list = append(list, next)
	}
	result := fmt.Sprintf("First %d Fibonacci numbers: %v", n, list)

	return result
}
