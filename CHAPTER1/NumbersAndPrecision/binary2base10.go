//binary2base10.go - base-2 to base-10 calculator
//Usage: go build binary2base10.go; ./binary2base10
//Copyright (c) 2020 rndMemex

package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var version = "0.0"

const exit = "X"

var sign string

//baseXConverter() converts base-2 to base-10
func baseIIConverter(bin string) (baseX float64) {
	fmt.Printf("\nConverting %s to base-10..\n\n", bin)
	if !strings.Contains(bin, ".") {
		var s float64
		for k := len(bin) - 1; k >= 0; k-- {
			i := len(bin) - 1
			n, _ := strconv.ParseFloat(string(bin[i-k]), 64)
			e := float64(k)
			s += n * math.Pow(2, e)
		}
		baseX = s
		if sign == "negative" {
			baseX = -baseX
		}
	} else {
		index := strings.Index(bin, ".")
		var s float64    //sum
		var dSum float64 //decimal point values sum
		integerBits := bin[:index]
		fractionalBits := bin[index+1:]
		for k := len(bin[:index]) - 1; k >= 0; k-- {
			i := len(bin[:index]) - 1
			nDigit, _ := strconv.ParseFloat(string(integerBits[i-k]), 64)
			e := float64(k)
			s += nDigit * math.Pow(2, e)
		}
		for i := 0; i < len(bin[index+1:]); i++ {
			d, _ := strconv.ParseFloat(string(fractionalBits[i]), 64)
			e := -float64((i + 1))
			dSum += d * math.Pow(2, e)
		}
		baseX = s + dSum
		if sign == "negative" {
			baseX = -baseX
		}
	}
	return baseX
}

/*isBinary() checks if input is a binary*/
func isBin(n string) bool {
	reDot := regexp.MustCompile(`^[01]+\.?[01]*$`)
	if sign == "negative" {
		n = n[1:]
	}
	if reDot.MatchString(n) {
		return true
	}
	return false
}

/*iFloat64() checks if the string is a number*/
func isFloat64(input string) bool {
	_, err := strconv.ParseFloat(input, 64)
	if strings.HasPrefix(input, "-") {
		sign = "negative"
	} else {
		sign = "positive"
	}
	log.Printf("Sign: %s", sign)
	return err == nil
}

/*enterNumber() reads from standard input and */
func enterNumber(reader *bufio.Reader) string {
	var bin string
	reader = bufio.NewReader(os.Stdin)
	for {
		fmt.Print("$ Enter a binary # (to exit press X): ")
		n, err := reader.ReadString('\n')
		if err != nil {
			log.Print(err)
			continue
		}
		n = strings.TrimSuffix(n, "\n")
		whitespaceExists := strings.HasSuffix(n, "")

		//this is necessary since a byte[0xd] may be added to the input.
		//New lines behave differently across platforms
		if whitespaceExists {
			n = strings.TrimSpace(n)
		}
		if !isFloat64(n) {
			exit := regexp.MustCompile(`(?i)^[X]+$`)
			if exit.MatchString(n) {
				goodbye()
				os.Exit(0)
			} else {
				fmt.Printf("ERROR!! %s is not a number!\n\n", n)
				continue
			}
		} else {
			if isBin((n)) {

				bin = n
			} else {
				fmt.Printf("\n%s is not a binary number\n\n", n)
				continue
			}
			return bin
		}
	}
	return bin
}

func main() {
	var inputNum *bufio.Reader
	fmt.Println("-----------------------")
	fmt.Println("Command Shell")
	fmt.Println("-----------------------")
	for {
		bin := enterNumber(inputNum)
		baseX := baseIIConverter(bin)
		fmt.Printf("The Binary number %s is %#v in base-10\n\n", bin, baseX)
	}
}
func goodbye() {
	fmt.Println("\nGoodbye!\n")
	fmt.Println("Copyright(c) 2020 rndmemex@cantab.net\n")
}

/*Sample Output of base-2 to base-10 calculator

-----------------------
Command Shell
-----------------------
$ Enter a binary # (to exit press X): 10110.1

Converting 10110.1 to base-10..

The Binary number 10110.1 is 22.5 in base-10

Goodbye!

Copyright(c) 2020 rndmemex@cantab.net

*/
