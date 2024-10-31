package main

import (
	"strings"
)

func RemoveQuote(input string) string {
	var result string
	inQuotes := false
	var quotedContent string

	for i := 0; i < len(input); i++ {
		char := input[i]

		// It detects if the char is a single quote, after this it skips " inQuotes = false " part and sets " inQuotes = true"
		if char == '\'' {
			if inQuotes {
				// then if it  detects the closing quote, sets " inQuotes = false " 
				inQuotes = false
				// adds the single quotes manually to ensure that the quotes are added to the result with our input which is quoteContent
				result += "'" + strings.TrimSpace(quotedContent) + "'"
				// Reset the quoted content for future use
				quotedContent = ""
			} else {
				// We found an opening quote
				inQuotes = true
			}
			// after detecting the opening quotes, continues back to the loop 
			continue
		}


		if inQuotes {
			// If we're inside quotes, it clloects all the characters 
	 
				quotedContent += string(char)
	
		} else {
			// If we're outside quotes, add the character to the result
			result += string(char)
		}
	}
	return strings.TrimSpace(result)
}

