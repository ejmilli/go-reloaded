package main

import "strings"

func QuoteHandle(input []string) []string {
	var resultSlice []string
	inQuotes := false
	var quotedContent []string

	for _, word := range input {
		// Check for single quote at the beginning or end of the word
		if word == "'" {
			if inQuotes {
				// Closing quote found
				inQuotes = false
				// Join quoted content into a single string and add it to the result slice
				resultSlice = append(resultSlice, "'"+strings.Join(quotedContent, " ")+"'")
				quotedContent = []string{} // Reset quoted content
			} else {
				// Opening quote found
				inQuotes = true
			}
		}
		if inQuotes {
			// Collect words inside quotes
			quotedContent = append(quotedContent, word)
		} else {
			// Add words outside quotes directly to result slice
			resultSlice = append(resultSlice, word)
		}
	}

	return resultSlice
}