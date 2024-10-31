package main

import "log"

func IsErr(err error) {
	if err != nil {
		log.Fatal("Error: ", err)
	}
}