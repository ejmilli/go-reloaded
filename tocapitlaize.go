package main

import (
	"strings"
)

func toCapitalize(cap string) string {
	result := ""
	for index, char := range cap {

		if index == 0 {
			result += strings.ToUpper(string(char))
		} else {
			result += strings.ToLower(string(char))
		}
		
	}
	return result
}
