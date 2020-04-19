//binary2base10.go
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

//baseXConverter() convert binary to base-10
func baseXConverter(bin string) (baseX float64) {
	fmt.Printf("\nConverting %s to base-10..\n\n", bin)
	if !strings.Contains(bin, ".") {
		var s float64
		for k := len(bin) - 1; k >= 0; k-- {
			i := len(bin) - 1
			n, _ := strconv.ParseFloat(string(bin[i-k]), 64)
			e := float64(k)
			s = s + n*math.Pow(2, e)
		}
		baseX = s
	} else {
		index := strings.Index(bin, ".")
		var s float64    //sum
		var dSum float64 //decimal points values sum
		binBeforeDecimalPoint := bin[:index]
		binAfterDecimalPoint := bin[index+1:]
		for k := len(bin[:index]) - 1; k >= 0; k-- {
			i := len(bin[:index]) - 1
			nDigit, _ := strconv.ParseFloat(string(binBeforeDecimalPoint[i-k]), 64)
			e := float64(k)
			s = s + nDigit*math.Pow(2, e)
		}
		for i := 0; i < len(bin[index+1:]); i++ {
			d, _ := strconv.ParseFloat(string(binAfterDecimalPoint[i]), 64)
			e := -float64((i + 1))
			dSum = dSum + d*math.Pow(2, e)
		}
		baseX = s + dSum
	}
	return baseX
}

/*isBinary() checks if input is a binary*/
func isBin(n string) bool {
	reDot := regexp.MustCompile(`^[01]+\.?[01]+$`)
	if reDot.MatchString(n) {
		return true
	} else {
		return false
	}
}

/*isInt() checks if a string is float64*/
func isFloat64(input string) bool {
	_, err := strconv.ParseFloat(input, 64)
	return err == nil
}

/*enterNumber() reads from standard input*/
func enterNumber(reader *bufio.Reader) string {
	var bin string
	reader = bufio.NewReader(os.Stdin)
	fmt.Println("-----------------------")
	fmt.Println("Command Shell")
	fmt.Println("-----------------------")
	for {
		fmt.Print("$ Enter a binary # (to exit press X): ")
		n, err := reader.ReadString('\n')
		if err != nil {
			log.Print(err)
			continue
		}
		n = strings.TrimSuffix(n, "\n")
		if !isFloat64(n) {
			exit := regexp.MustCompile(`(?i)^[X]+$`)
			if exit.MatchString(n) {
				goodbye()
				os.Exit(0)
			} else {
				fmt.Printf("ERROR!! %s is not a number!\n", n)
				continue
			}
		} else {
			if isBin((n)) {
				bin = n
			} else {
				fmt.Printf("%s is not a binary number\n", n)
				continue
			}
			return bin
		}
	}
	return bin
}

func main() {
	var inputNum *bufio.Reader
	bin := enterNumber(inputNum)
	baseX := baseXConverter(bin)
	fmt.Printf("The Binary number %s is %#v in base-10\n", bin, baseX)
	goodbye()

}

func goodbye() {
	fmt.Println("\nGoodbye!\n")
	fmt.Println("Copyright(c) 2020 rndmemex@cantab.net\n")
}
