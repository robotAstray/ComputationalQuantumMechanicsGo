//ifStatements.go.go Find the smallest of three real numbers a, b, c
//Usage: go build ifStatements.go; ./ifStatements
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

var version = "0.1"

const exit = "X"
const lengthLimit = 3

func main() {
	fmt.Println("-----------------------")
	fmt.Println("Command Shell")
	fmt.Println("-----------------------")
	//ask for inpuut number
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("$ Input three numbers separated by a comma (to exit press X):")
		n, err := reader.ReadString('\n')
		if err != nil {
			log.Print(err)
			continue
		}
		//if user provide number separeted by a comma
		n = strings.TrimSuffix(n, "\n")
		reComma := regexp.MustCompile(`,-?[0-9]+`)
		commaExists := reComma.MatchString(n)
		commaSpaceExists := strings.Contains(n, ", ")
		//correspond to carriage return
		carriageReturnExists := strings.HasSuffix(n, "\r")
		var numFloatArray []float64
		if commaExists {
			//this is necessary since a byte[0xd] may be added to the input.
			//New lines behave differently across platforms
			if carriageReturnExists {
				n = n[:len(n)-1]
			}

			numArray := strings.Split(n, ",")
			limitExceeded := limitReached(numArray, lengthLimit)
			if limitExceeded || len(numArray) == 1 {
				fmt.Printf("You have entered %v numbers, please start again and follow the instructions\n", len(numArray))
				goodbye()
				break
			}
			for i := 0; i < len(numArray); i++ {
				numStr := strings.TrimPrefix(strings.TrimSuffix(numArray[i], " "), " ")
				nF, err := strconv.ParseFloat(numStr, 64) //number float
				if err != nil {
					fmt.Print(err)
				}
				numFloatArray = append(numFloatArray, nF)
			}
			//find the smallest of the three numbers
			a := numFloatArray[0]
			b := numFloatArray[1]
			c := numFloatArray[2]
			var smallest float64
			if a < b && a < c {
				smallest = a
			} else if b < a && b < c {
				smallest = b
			} else {
				smallest = c
			}

			fmt.Printf("Smallest of three numbers: %v\n", smallest)
		} else if commaSpaceExists {
			fmt.Printf("carriageReturnExists:%v\n", carriageReturnExists)
			//this is necessary since a byte[0xd] may be added to the input.
			//New lines behave differently across platforms
			if carriageReturnExists {
				n = n[:len(n)-1]
			}
			numArray := strings.Split(n, ", ")
			limitExceeded := limitReached(numArray, lengthLimit)
			if limitExceeded || len(numArray) == 1 {
				fmt.Printf("You have entered %v numbers, please start again and follow the instructions\n", len(numArray))
				goodbye()
				break
			}
			for i := 0; i < len(numArray); i++ {
				numStr := strings.TrimPrefix(strings.TrimSuffix(numArray[i], " "), " ")
				nF, err := strconv.ParseFloat(numStr, 64) //number float
				if err != nil {
					fmt.Print(err)
				}
				numFloatArray = append(numFloatArray, nF)
			}
			//find the smallest of the three numbers
			a := numFloatArray[0]
			b := numFloatArray[1]
			c := numFloatArray[2]
			var smallest float64
			if a < b && a < c {
				smallest = a
			} else if b < a && b < c {
				smallest = b
			} else {
				smallest = c
			}
			// log.Printf("array input: %#v", numArray)
			// log.Printf("Length of input array: %d", len(numArray))
			// log.Printf("Length of last input: %#v", len(numArray[1]))
			// log.Printf("Array of float64:%v", numFloatArray)
			fmt.Printf("Smallest of three numbers: %v\n", smallest)
		} else {
			if carriageReturnExists {
				n = n[:len(n)-1]
				if n == exit {
					goodbye()
					os.Exit(0)
				}
			}

			log.Printf("ERROR! Please follow the instructions below.")
		}

	}
}

func goodbye() {
	fmt.Printf("\nGoodbye!\n")
	fmt.Printf("Copyright(c) 2020 rndmemex@cantab.net\n")
}

func limitReached(n []string, limit int) (limitReach bool) {
	limitReach = false
	if len(n) > limit {
		limitReach = true
	}
	return limitReach
}
