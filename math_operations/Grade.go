package math_operations

import "fmt"

/*
5. Simple Grade Calculator
Problem: Write a function calculateGrade(score int) string that takes a numeric
score and returns a letter grade based on the following scale:

90 and above: A
80 to 89: B
70 to 79: C
60 to 69: D
Below 60: F
Print the grade for various scores.

Example Output:
Score: 85, Grade: B
Score: 72, Grade: C
Score: 45, Grade: F
*/

func getGrade(grade int) string {
	if grade < 60 {
		return "F"
	}else if grade >= 60 && grade <= 69 {
		return "D"
	} else if grade >= 70 && grade <= 79 {
		return "C"
	} else if grade >= 80 && grade <= 89 {
		return "B"
	} else {
		return "A"
	}
}

func CalculateGrade(score int) string {
	grade := getGrade(score)
	result := fmt.Sprintf("Score %d, Grade %s", score, grade)
	return result 
}
