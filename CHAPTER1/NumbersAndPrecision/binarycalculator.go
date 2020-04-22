//binarycalculator.go - base-10 to base-2 calculator
//Usage: go build binarycalculator.go; ./binarycalculator
//Copyright (c) 2020 rndMemex

package main

import (
	"bufio"
	"fmt"
)

func binaryCalculator(baseXNum string) (bin float64) {
	fmt.Printf("\nConverting %s to binary representation..\n\n", baseXNum)
	if !strings.Contains(bin, ".") {
	} else {

	}
	return bin
}

//enterNumber() reads from a standard input
func enterNumber(reader *bufio.Reader) string {
	var baseXNum string
	reader = bufio.NewReader(os.Stdin)
	for {
		fmt.Print("$ enter a number (to exit press X): ")
		n, err := reader.ReadString('\n')
		if err != nil {
			log.Println(err)
			continue
		}
		n = strings.TrimSuffix(n, "\n")
		if !isFloat(n) {
			exit := regexp.MustCompile(`(?i)^[X]+$`)
			if exit.MatchString(n) {
				goodbye()
				os.Exit(0)
			} else {
				fmt.Printf("ERROR!! %s is not a number!\n\n", n)
				continue
			}
		} else {
			baseXNum = n
		}
		return baseXNum
	}
	return baseXNum
}

/*iFloat64() checks if the string is a number*/
func isFloat64(input string) bool {
	_, err := strconv.ParseFloat(input, 64)
	return err == nil
}

func main() {
	var inputNum *bufio.Reader
	fmt.Println("-----------------------")
	fmt.Println("Command Shell")
	fmt.Println("-----------------------")
	for {
		baseX := enterNumber(inputNum)
		bin := binaryCalculator(baseX)
		fmt.Printf("The %s base-10 is %#v base-2 \n\n", baseX, bin)
	}
}

func goodbye() {
	fmt.Println("\nGoodbye!\n")
	fmt.Println("Copyright(c) 2020 rndmemex@cantab.net\n")
}
