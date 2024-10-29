package main

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
				result += "'" + quotedContent + "'"
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
			if char != ' ' {
				quotedContent += string(char)
			}
		} else {
			// If we're outside quotes, add the character to the result
			result += string(char)
		}
	}

	// Return the final cleaned up string
	return result
}