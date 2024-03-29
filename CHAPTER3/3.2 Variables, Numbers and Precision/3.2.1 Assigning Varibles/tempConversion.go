//tempConversion.go - Celsius to Fahrenheit calculator
//Usage: go build tempConversion.go; ./tempConversion
//Copyright (c) robotAstray

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"robotAstray"
	"strconv"
	"strings"
)

var sign string //global variable

func fahrenheit(temp string) float64 {
	//convert string to float64
	DegC, _ := strconv.ParseFloat(temp, 64)
	if sign == "negative" {
		DegC = -DegC
	}
	DegF := DegC*(9.0/5.0) + 32
	return DegF
}

func enterTemp(reader *bufio.Reader) string {
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

		//the following is necessary since a byte[0xd] may be added to the input.
		//New lines behave differently across platforms.
		if whitespaceExists {
			n = strings.TrimSpace(n)
		}
		if !isFloat64(n) {
			exit := regexp.MustCompile(`(?i)^[X]+$`)
			if exit.MatchString(n) {
				robotAstray.Goodbye()
				os.Exit(0)
			}
			fmt.Printf("ERROR!! %s is  not a number!\n\n", n)
			continue
		}
		temp = n
		if sign == "negative" {
			temp = n[1:]
			log.Printf("-%s is a negative number\n", temp)
		}
		return temp
	}
}

func isFloat64(input string) bool {
	_, err := strconv.ParseFloat(input, 64)
	sign = "positive"
	if strings.HasPrefix(input, "-") {
		sign = "negative"
	}
	// log.Printf("Sign: %s", sign)
	return err == nil
}

func main() {
	var inputTemp *bufio.Reader
	fmt.Println("-----------------------")
	fmt.Println("Command Shell")
	fmt.Println("-----------------------")
	for {
		tempCelsius := enterTemp(inputTemp)
		tempFahrenheit := fahrenheit(tempCelsius)
		if sign == "negative" {
			fmt.Printf("-%s Celsius is equal to %v Fahrenheit\n\n", tempCelsius, tempFahrenheit)
		} else {
			fmt.Printf("%s Celsius is equal to %v Fahrenheit\n\n", tempCelsius, tempFahrenheit)
		}
	}
}
