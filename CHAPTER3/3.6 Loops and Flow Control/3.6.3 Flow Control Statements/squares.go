//squares.go - print all squares numbers between 10 and 60
//Usage: go build squares.go; ./squares
//Copyright (c) 2020 rndMemex

package main

import "fmt"

func main() {
	i := 0
	for {

		i++
		sqr := i * i
		if sqr < 10 {
			continue
		} else if sqr > 60 {
			break
		}
		fmt.Println(sqr)
	}
}
