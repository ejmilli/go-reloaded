package main

import (
	"fmt"
	"math"
)

func BinCharToDec(bin rune) int {

	if bin == '0' || bin == '1' {
		return int(bin - '0')

	}
	return -1
}

func BinToDecimal(bina string) int {
	result := 0
	length := len(bina) // Get the length of the hex string

	for i := length - 1; i >= 0; i-- {
			char := bina[i] // Get the current character
			value := BinCharToDec(rune(char)) // Convert to decimal value

			// Error handling for invalid characters
			if value == -1 {
					return -1 // Return -1 for invalid input; change this if needed
			}

			position := length - 1 - i // Calculate position
			
			// Use multiplication to calculate the contribution
			result += value * int(math.Pow(2, float64(position)))
	}

	return result // Return the final decimal value
}


func main() {
	// Test cases
	tests := []string{
		"101",   // Expected output: 5
		"111",   // Expected output: 7
		"1001",  // Expected output: 9
		"0010",  // Expected output: 2
		"00001", // Expected output: 1
		"00000", // Expected output: 0
		"0",     // Expected output: 0
		"1",     // Expected output: 1
		"102",   // Expected output: Invalid input, returns -1
	}

	// Running the test cases
	for _, test := range tests {
		fmt.Printf("Binary: %s -> Decimal: %d\n", test, BinToDecimal(test))
	}
}