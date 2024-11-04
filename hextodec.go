package main

import (
	"math"
	"strconv"
)

func HexCharToDecimal(char rune) int {

	var valueStr int

	// here it converts 0 -9 by saying that if my char is 0, i will do int(char-'0') which means int(48-48) = 0
	if char >= '0' && char <= '9' {
		valueStr = int(char - '0')
		// here it converts A- F by saying if that if my char is A, i will do int(char-'A') + 10 which means int(65-65) + 10 = 10
	} else if char >= 'A' && char <= 'F' {
		valueStr = int(char-'A') + 10
	}  else if char >= 'a' && char <= 'f' {
		valueStr = int(char-'a') + 10
	} else {
		return -1
	}
	return valueStr
}


func HexToDecimal(hexa string) string {
	result := 0
	length := len(hexa) // Get the length of the hex string

	for i := length - 1; i >= 0; i-- {
			char := hexa[i] // Get the current character
			value := HexCharToDecimal(rune(char)) // Convert to decimal value

			// Error handling for invalid characters
			if value == -1 {
					return "-1" // Return -1 for invalid input; change this if needed
			}

			position := length - 1 - i // Calculate position
			
			// Use multiplication to calculate the contribution
			result += value * int(math.Pow(16, float64(position)))
	}

	return strconv.Itoa(result) // Return the final decimal value
}