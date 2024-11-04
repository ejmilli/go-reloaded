package main

import "strconv"

func extractNum(s string) int {

	Num, err := strconv.Atoi(string(s[:len(s)-1]))
	IsErr(err)
	return Num

}

/////////////   example  //////////////


// supposedly i have "I am exciting (up 2)" it removes "2" and converts it into int 2 n then use it in labrat to actually convert to up, low, capitalize etc