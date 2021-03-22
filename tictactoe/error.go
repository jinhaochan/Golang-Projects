package main

import (
	"fmt"
	"os"
)

func errorCheck(e error) {
	if e != nil {
		fmt.Println("Error: ", e)
		os.Exit(1)
	}
}
