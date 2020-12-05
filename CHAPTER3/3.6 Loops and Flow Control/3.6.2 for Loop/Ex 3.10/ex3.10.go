//ex3.10.go - ex3.10 over runes
//Usage: go build ex3.10.go; ./ex3.10
//Copyright (c) 2020 rndMemex

package main

import (
	"fmt"
	"os"
)

func main() {
	var interval = []int{5, 11}
	for i := interval[0]; i <= interval[1]; i++ {
		fmt.Println(i * i)
	}
	goodbye()
}

func goodbye() {
	fmt.Printf("\nGoodbye!\n")
	fmt.Printf("Copyright(c) 2020 rndmemex@cantab.net\n")
	os.Exit(0)
}
