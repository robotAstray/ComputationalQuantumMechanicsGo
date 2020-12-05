//iterate.go - iterate over runes
//Usage: go build iterate.go; ./iterate
//Copyright (c) 2020 rndMemex

package main

import (
	"log"
)

func main() {
	s := "Hello World! This is rndMemex."
	for i := 0; i < len(s); i++ {
		log.Printf("%#U, at position %d\n", s[i], i)

	}
}
