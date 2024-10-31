package main

import (
	"os"
)

func main() {

	if len(os.Args) != 3 {
		return
	}

	fileB, err := os.ReadFile(os.Args[1])
	IsErr(err)

	ModifiedText := MyLab(string(fileB))

	os.WriteFile(os.Args[2], []byte(ModifiedText), 0644)

}

