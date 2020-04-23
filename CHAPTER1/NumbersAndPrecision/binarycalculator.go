//binarycalculator.go - base-10 to base-2 calculator
//Usage: go build binarycalculator.go; ./binarycalculator
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

func binaryCalculator(baseXNum string) (bin float64) {
	fmt.Printf("\nConverting %s base-10 to binary representation..\n\n", baseXNum)
	ratio := int(math.Log10(10) / math.Log10(2))
	reps := len(baseXNum) * ratio
	fmt.Printf("Upper bound on No. of operations required ~ %d*ln(10)/ln(2) = %d\n\n", len(baseXNum), reps)
	if !strings.Contains(baseXNum, ".") {
		var s []string
		n, _ := strconv.Atoi(baseXNum)
		for i := 0; i < reps; i++ {
			numerator := n
			if n == 0 {
				break
			}
			remainder := n % 2
			integerBit := strconv.Itoa(remainder)
			n = n / 2
			fmt.Printf("Operation %d: %d/2=%d with remainder %d\n\n", i+1, numerator, n, remainder)
			s = append(s, integerBit)

		}
		//fmt.Printf("remainder slice %v\n", s)
		var reverseStr []string
		//reverse order in s slice
		for j := range s {
			integerBitStr := s[len(s)-1-j]
			reverseStr = append(reverseStr, integerBitStr)
		}
		//fmt.Printf("Integers Bit %v", reverseStr)
		var binStr string
		binStr = strings.Join(reverseStr, "")
		bin, _ := strconv.ParseFloat(binStr, 64)
		return bin
	} else {
		log.Println("floating number")
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
		fmt.Printf("The binary representation of %s base-10 is %#v\n\n", baseX, bin)
	}

}

func goodbye() {
	fmt.Println("\nGoodbye!\n")
	fmt.Println("Copyright(c) 2020 rndmemex@cantab.net\n")
}

/* Sample Output:

-----------------------
Command Shell
-----------------------
$ enter a number (to exit press X): 23

Converting 23 base-10 to binary representation..

Upper bound on No. of operations required ~ 2*ln(10)/ln(2) = 6

Operation 1: 23/2=11 with remainder 1

Operation 2: 11/2=5 with remainder 1

Operation 3: 5/2=2 with remainder 1

Operation 4: 2/2=1 with remainder 0

Operation 5: 1/2=0 with remainder 1

The binary representation of 23 base-10 is 10111
-----------------------
Command Shell
-----------------------
$ enter a number (to exit press X): x

Goodbye!

Copyright(c) 2020 rndmemex@cantab.net
*/
