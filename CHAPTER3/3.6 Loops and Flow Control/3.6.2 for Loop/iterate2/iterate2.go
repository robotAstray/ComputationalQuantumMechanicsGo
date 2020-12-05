//iterate2.go - iterate over runes
//Usage: go build iterate2.go; ./iterate2
//Copyright (c) 2020 rndMemex

package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	s := "Hello World! This is rndMemex."
	for i, c := range s {
		log.Printf("%#U, at position %d\n", c, i)

	}
	goodbye()
}

func goodbye() {
	fmt.Printf("\nGoodbye!\n")
	fmt.Printf("Copyright(c) 2020 rndmemex@cantab.net\n")
	os.Exit(0)
}
