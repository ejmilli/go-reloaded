package main

import (
	"strings"
)


func addAn(words []string) []string {
	var modifiedWords []string

	for i := 0; i < len(words); i++ {
		word := words[i]

		// Check if the word is "a" or "A" and the next word starts with a vowel or 'h'
		if (word == "a" || word == "A") && i+1 < len(words) && len(words[i+1]) > 0 &&
			strings.Contains("aeiouhAEIOUH", string(words[i+1][0])) {
			
			// Add "an" or "An" based on the original case
			if word == "a" {
				modifiedWords = append(modifiedWords, "an")
			} else {
				modifiedWords = append(modifiedWords, "An")
			}
		} else {
			// Add the word unchanged if the condition does not match
			modifiedWords = append(modifiedWords, word)
		}
	}
	return modifiedWords
}




