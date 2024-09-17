package math_operations

import "fmt"

/*
6. Sum of All Even Numbers in a Range
Problem: Write a function sumEvenNumbers(start, end int) int that calculates the sum of
all even numbers in the range from start to end (inclusive). Print the result.

Example Output:

Sum of even numbers between 1 and 10 is: 30

*/

func SumEvenNumbers(start, end int) string {
	even := 0
	for i := start; i <= end; i++ {
		if i % 2 == 0 {
			even += i
		}
	}
	result := fmt.Sprintf("Sum of even numbers between %d and %d is: %d", start, end, even)
	return result
}
