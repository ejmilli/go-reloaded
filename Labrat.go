package main

import (
	"strings"
)

func MyLab(sentence string) string {
	words := strings.Fields(sentence)
	if len(words) > 0 {
		words[0] = toCapitalize(words[0])
	}
	var num int
	for i := 0; i < len(words); i++ {
		switch words[i] {
		case "(hex)":
			if i > 0 {
				words[i-1] =  HexToDecimal(words[i-1])
				words[i] = ""
			}
		case "(bin)":
			if i > 0 {
				words[i-1] = BinToDecimal(words[i-1])
				words[i] = ""
			}
		case "(up)":
			if i > 0 {
				words[i-1] = toUpper(words[i-1])
				words[i] = ""
			}
		case "(low)":
			if i > 0 {
				words[i-1] = ToLower(words[i-1])
				words[i] = ""
			}
		case "(cap)":
			if i > 0 {
				words[i-1] = toCapitalize(words[i-1])
				words[i] = ""
			}
		case "(low,":
			if i < len(words)-1 {
				num = extractNum(words[i+1])
				NumModify(num, i, words, ToLower)
				words[i] = ""
				words[i+1] = ""
			}
		case "(up,":
			if i < len(words)-1 {
				num = extractNum(words[i+1])
				NumModify(num, i, words, toUpper)
				words[i] = ""
				words[i+1] = ""
			}
		case "(cap,":
			if i < len(words)-1 {
				num = extractNum(words[i+1])
				NumModify(num, i, words, toCapitalize)
				words[i] = ""
				words[i+1] = ""
			}
		}
	}
	edited := addAn(RemoveQuote(removePunctuation(removeTheSlice(words))))
	return strings.Join(edited, " ")
}
