//typeConversion.go - string to float64 to string
//Usage: go build typeConversion.go; ./typeConversion
//Copyright (c) 2020 rndMemex

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var sign string

func addXXIAndConvertBackToStr(num string) (float64, string) {
	floatNum, _ := strconv.ParseFloat(num, 64)
	fmt.Printf("converting the string into float64 and adding 21...\n\n")
	fmt.Printf("%#v + 21.0 = %#v\n\n", floatNum, floatNum+21)
	floatNum = floatNum + 21.0

	strNum := fmt.Sprintf("%f", floatNum)
	if sign == "negative" {
		strNum = "-" + strNum
	}
	fmt.Printf("converting float64: %#v to string: %#v\n\n", floatNum, strNum)
	return floatNum, strNum
}
func enterNumber(reader *bufio.Reader) string {
	var num string
	reader = bufio.NewReader(os.Stdin)
	for {
		fmt.Print("$ Enter a number (to exit press X): ")
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
			num = n
			if sign == "negative" {
				num = n[1:]
			}
		}
		return num
	}
}
func isFloat64(input string) bool {
	_, err := strconv.ParseFloat(input, 64)
	if strings.HasPrefix(input, "-") {
		sign = "negative"
	} else {
		sign = "positive"
	}
	return err == nil
}
func goodbye() {
	fmt.Printf("\nGoodbye!\n")
	fmt.Printf("Copyright(c) 2020 rndmemex@cantab.net\n")
}

func main() {
	var inputNumber *bufio.Reader
	numStr := enterNumber(inputNumber)
	floatConv, strConv := addXXIAndConvertBackToStr(numStr)
	fmt.Printf("Output: %#v, %v, %#v\n\n", numStr, floatConv, strConv)
}
