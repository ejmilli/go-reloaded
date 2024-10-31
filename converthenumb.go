package main

import "strconv"

func extractNum(s string) int {

	Num, err := strconv.Atoi(string(s[:len(s)-1]))
	IsErr(err)
	return Num

}

