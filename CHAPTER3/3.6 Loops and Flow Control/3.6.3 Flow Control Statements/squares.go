//squares.go - squares over runes
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
