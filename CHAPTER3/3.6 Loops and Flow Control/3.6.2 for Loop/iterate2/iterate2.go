//iterate2.go - iterate over runes
//Usage: go build iterate2.go; ./iterate2
//Copyright (c) 2020 rndMemex

package main

import (
	"log"
)

func main() {
	s := "Hello World! This is rndMemex."
	for i, c := range s {
		log.Printf("%#U, at position %d\n", c, i)

	}
}
