package main

import "strings"

func QuoteHandle(words []string) []string {

	var QuoteCount int

	for i := 0; i < len(words); i++ {
		if QuoteCount%2 == 0 && i < len(words)-1 && strings.Contains("'", words[i]) {
			words[i+1] = "'" + words[i+1]
			words[i] = ""
			QuoteCount++
		} else if QuoteCount%2 == 1 && i > 0 && strings.Contains("'", words[i]) {
			words[i-1] = words[i-1] + "'"
			words[i] = ""
			QuoteCount++
		}
	}

	return removeTheSlice(words)
}