//squares.go
//Printing out all squares numbers less than 60

package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	for i := 0; i < 60; i++ {
		square := i * i
		if square > 60 {
			goodbye()
		}
		log.Println(square)
	}
}

func goodbye() {
	fmt.Printf("\nGoodbye!\n")
	fmt.Printf("Copyright(c) 2020 rndmemex@cantab.net\n")
	os.Exit(0)
}
