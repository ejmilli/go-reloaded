package main

import (
	"fmt"
	"strings"
)

func QuoteHandling(input string) string {
	var result string
	inQuotes := false
	var quotedContent string

	for i := 0; i < len(input); i++ {
		char := input[i]

		// Check for the opening single quote
		if char == '\'' {
			if inQuotes {
				// We found a closing quote
				inQuotes = false
				// Add the cleaned quoted content back to the result
				result += "'" + strings.TrimSpace(quotedContent) + "'"
				// Reset the quoted content for future use
				quotedContent = ""
			} else {
				// We found an opening quote
				inQuotes = true
			}
			// Continue to the next character
			continue
		}


		if inQuotes {
			// If we're inside quotes, collect characters
	 
				quotedContent += string(char)
	
		} else {
			// If we're outside quotes, add the character to the result
			result += string(char)
		}
	}
	return strings.TrimSpace(result)
}


func main() {
	// Test cases for the QuoteHandling function
	testCases := []struct {
		input    string
		expected string
	}{
		{" ' Hello World ' ", "'Hello World'"},
		{" 'Hello World' ", "'Hello World'"},
		{"Hello ' World ' !", "Hello 'World'!"},
		{"This is a test ' string  '  with quotes", "This is a test 'string' with quotes"},
		{"A ' single ' test case.", "A 'single' test case."},
		{"This is not ' a ' test", "This is not 'a' test"},
		{" 'Leading and trailing spaces ' ", "'Leading and trailing spaces'"},
		{"'  Too many   spaces  ' ", "'Too many spaces'"},
		{"This is a sentence without quotes.", "This is a sentence without quotes."},
		{"   ' Just spaces before and after '   ", "'Just spaces before and after'"},
		{"I am exactly how they describe me: ' awesome '" , "I am exactly how they describe me: 'awesome'"},
		{"As Elton John said: ' I am the most well-known homosexual in the world '", "As Elton John said: 'I am the most well-known homosexual in the world'"},
	}

	// Iterate through each test case
	for _, tc := range testCases {
		result := QuoteHandling(tc.input)
		fmt.Printf("Input: %q\nExpected: %q\nOutput: %q\n\n", tc.input, tc.expected, result)
	}
}