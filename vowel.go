package main

import "strings"

func addAn(vowel []string) []string {

	for i := 0; i < len(vowel); i++ {
		if len(vowel[i]) > 0 && strings.Contains("aeiouh", string(vowel[i][0])) {
       vowel[i] = "an " + vowel[i]
		}

	}
   return removeTheSlice(vowel) 
}