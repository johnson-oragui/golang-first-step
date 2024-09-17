package math_operations


func BitwiseMultiply(a, b int) int {
    result := 0
    for b != 0 {
        if b&1 != 0 {           // If the least significant bit is set
            result = BitwiseAdd(result, a)  // Add a to the result
        }
        a <<= 1                 // Left shift a
        b >>= 1                 // Right shift b
    }
    return result
}

func BitwiseSubtract(a, b int) int {
    return BitwiseAdd(a, BitwiseAdd(^b, 1))  // Add a and -b using two's complement
}

func BitwiseAdd(a, b int) int {
	for b != 0 {
		carry := a & b
		a = a ^ b

		b = carry << 1
	}
	return a
}
