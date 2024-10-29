package main

import (
	"strings"
)



func addAn(words []string) []string {
	var modifiedWords []string

	for i := 0; i < len(words); i++ {
		word := words[i]

		// If word is "a" and the next word starts with a vowel or 'h', replace it with "an"
		if word == "a" && i+1 < len(words) && len(words[i+1]) > 0 &&
			strings.Contains("aeiouhAEIOUH", string(words[i+1][0])) {
			modifiedWords = append(modifiedWords, "an")
		} else {
			modifiedWords = append(modifiedWords, word)
		}
	}
	return modifiedWords
}




