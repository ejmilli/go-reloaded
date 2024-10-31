package main

import (
	"math"
	"strconv"
)

func BinCharToDec(bin rune) int {

	if bin == '0' || bin == '1' {
		return int(bin - '0')

	}
	return -1
}

func BinToDecimal(bina string) string {
	result := 0
	length := len(bina) // Get the length of the binary string

	for i := length - 1; i >= 0; i-- {
			char := bina[i] // Get the current character
			value := BinCharToDec(rune(char)) // Convert to decimal value

			// Error handling for invalid characters
			if value == -1 {
					return "-1" // Return "-1" for invalid input; change this if needed
			}

			position := length - 1 - i // Calculate position
			
			// Use multiplication to calculate the contribution
			result += value * int(math.Pow(2, float64(position)))
	}

	return strconv.Itoa(result) // Convert result to string and return it
}

