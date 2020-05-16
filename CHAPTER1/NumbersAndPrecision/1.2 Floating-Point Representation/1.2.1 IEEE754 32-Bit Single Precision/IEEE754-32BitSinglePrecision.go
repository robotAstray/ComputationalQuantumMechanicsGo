//floatingpointRepresentation.go - base-10 to base-2 calculator to floating point representation
//Usage: go build floatingpointRepresentation.go; ./floatingpointRepresentation
//Copyright (c) 2020 rndMemex

package main

import (
	"bufio"
	"fmt"
)

func main() {
	var inputNum *bufio.Reader
	fmt.Println("-----------------------")
	fmt.Println("Command Shell")
	fmt.Println("-----------------------")
	for {
		baseX := enterNumber(inputNum)
		bin := binaryCalculator(baseX)
		fmt.Printf("The binary representation of %s base-10 is %#v\n\n", baseX, bin)
		fmt.Printf("Converting to floating point representation...\n\n")
		//always using 8 significant figures
		VIIIFloatingPointRep := floatingPointRep(bin)
		fmt.Printf("The floating-point representation of %v is %#v\n\n", bin, VIIIFloatingPointRep)
	}

}

func goodbye() {
	fmt.Printf("\nGoodbye!\n")
	fmt.Printf("Copyright(c) 2020 rndmemex@cantab.net\n")
}
