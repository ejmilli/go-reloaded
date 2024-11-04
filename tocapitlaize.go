package main

import (
	"strings"
)

func toCapitalize(cap string) string {
var result string 
	for index, char := range cap {

		if index == 0 {
			result += strings.ToUpper(string(char))
		} else {
			result += strings.ToLower(string(char))
		}
		
	}
	return result
}
