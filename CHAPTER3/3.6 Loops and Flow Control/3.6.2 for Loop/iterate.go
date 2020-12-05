//iterate.go - iterate over runes
//Usage: go build iterate.go; ./iterate
//Copyright (c) 2020 rndMemex

package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	s := "Hello World! This is rndMemex."
	for i := 0; i < len(s); i++ {
		log.Printf("%#U, at position %d\n", s[i], i)

	}
	goodbye()
}

func goodbye() {
	fmt.Printf("\nGoodbye!\n")
	fmt.Printf("Copyright(c) 2020 rndmemex@cantab.net\n")
	os.Exit(0)
}
