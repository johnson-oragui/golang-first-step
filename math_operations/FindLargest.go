package math_operations

/*
2. Find the Largest of Three Numbers
Problem: Write a function findLargest(a, b, c int) int that returns
the largest of three integers a, b, and c. Test the function with
different values.

Example Output:
The largest number is: 15

*/

func FindLargest(a, b, c int) int {
	largest := a

	if b > largest {
		largest = b
	}

	if c > largest {
		largest = c
	}

	return largest
}
