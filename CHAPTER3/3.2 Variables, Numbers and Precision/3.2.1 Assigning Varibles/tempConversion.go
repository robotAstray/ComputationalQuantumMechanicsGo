//tempConversion.go - Celsius to Farnheit calculator
//Usage: go build tempConversion.go; ./tempconversion
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

func farnheitConversion(temp string) float64 {
	//convert string to float
	DegC, _ := strconv.ParseFloat(temp, 64)
	if sign == "negative" {
		DegC = -DegC
	}
	DegF := DegC*(9.0/5.0) + 32

	return DegF
}

func enterTemperature(reader *bufio.Reader) string {
	var temp string
	reader = bufio.NewReader(os.Stdin)
	for {
		fmt.Print("$ Enter temperature in Celsius (to exit press X): ")
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
			}
			// } else {
			fmt.Printf("ERROR!! %s is not a number!\n\n", n)
			continue
			//}
		} else {
			temp = n
			if sign == "negative" {
				temp = n[1:]

				log.Printf("Negative number is %s\n", temp)
			}
		}
		return temp
	}
}

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

func goodbye() {
	fmt.Printf("\nGoodbye!\n")
	fmt.Printf("Copyright(c) 2020 rndmemex@cantab.net\n")
}

func main() {
	var inputTemp *bufio.Reader
	fmt.Println("-----------------------")
	fmt.Println("Command Shell")
	fmt.Println("-----------------------")
	for {
		tempCelsius := enterTemperature(inputTemp)
		tempFarenheit := farnheitConversion(tempCelsius)
		if sign == "negative" {
			fmt.Printf("-%s Celsius is equals to %v Fahrenheit\n\n", tempCelsius, tempFarenheit)
		} else {
			fmt.Printf("%s Celsius is equals to %v Fahrenheit\n\n", tempCelsius, tempFarenheit)
		}
	}
}
