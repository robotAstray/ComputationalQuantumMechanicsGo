//floatingpointRepresentation.go - base-10 to base-2 calculator to floating point representation
//Usage: go build floatingpointRepresentation.go; ./floatingpointRepresentation
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

func floatingPointRep(bin float64) (floatingPointRepStr string) {
	binStr := fmt.Sprintf("%f", bin)
	index := strings.Index(binStr, ".")
	intBits := binStr[:index]
	fractionalBits := binStr[index+1:]
	significantFigures := 8
	if len(intBits) == 1 {
		floatingPointRepStr = intBits + "." + fractionalBits
		zeros := significantFigures - len(fractionalBits) - 1
		for k := 0; k < zeros; k++ {
			floatingPointRepStr = floatingPointRepStr + "0"
		}
		return floatingPointRepStr
	} else {
		power := len(intBits[1:])
		floatingPointRepStr = intBits[:1] + "." + intBits[1:] + fractionalBits + "x 2^(" + strconv.Itoa(power) + ")"
		return floatingPointRepStr
	}
}

func binaryCalculator(baseXNum string) (bin float64) {
	fmt.Printf("\nConverting %s base-10 to binary representation..\n\n", baseXNum)
	ratio := int(math.Log10(10) / math.Log10(2))

	if !strings.Contains(baseXNum, ".") {
		reps := len(baseXNum) * ratio //upper bound on operations needed
		fmt.Printf("Upper bound on No. of operations required ~ %d*log(10)/log(2) = %d\n\n", len(baseXNum), reps)
		var s []string
		n, _ := strconv.Atoi(baseXNum)
		for i := 0; i < reps+1; i++ {
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
		var reverseStr []string
		//reverse order in s slice
		for j := range s {
			integerBitStr := s[len(s)-1-j]
			reverseStr = append(reverseStr, integerBitStr)
		}
		var binStr string
		binStr = strings.Join(reverseStr, "")
		bin, _ := strconv.ParseFloat(binStr, 64)
		return bin
	} else {
		index := strings.Index(baseXNum, ".")
		integerNum := baseXNum[:index]
		var s []string
		n, _ := strconv.Atoi(integerNum)
		integerReps := len(integerNum) * ratio //upper bound on operations needed
		fmt.Printf("STEP 1: Converting integer part: %s\n\n", integerNum)
		fmt.Printf("Upper bound on No. of operations required ~ %d*log(10)/log(2) = %d\n\n", len(integerNum), integerReps)
		for i := 0; i < integerReps+1; i++ {
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
		var reverseIntStr []string
		//reverse order in s slice
		for j := range s {
			integerBitStr := s[len(s)-1-j]
			reverseIntStr = append(reverseIntStr, integerBitStr)
		}
		var binIntStr string
		binIntStr = strings.Join(reverseIntStr, "")
		intBin, _ := strconv.ParseFloat(binIntStr, 64)
		fmt.Printf("Integer Bits: %v\n\n", intBin)

		//"Converting fractional part...\n\n"
		fractionalNum := baseXNum[index+1:]
		inputF64, _ := strconv.ParseFloat(baseXNum, 64)
		integerNumF64, _ := strconv.ParseFloat(integerNum, 64)
		fractionalPart := inputF64 - integerNumF64
		fmt.Printf("STEP 2: Converting fractional part: %v\n\n", fractionalPart)
		var fractSlice []string
		var m float64

		//multiply the fractional part by 2 each time, storing each integer until we have a fractional part of 0
		//the following should be improved
		for k := 0; k < len(fractionalNum); k++ {
			m = fractionalPart * 2
			fmt.Printf("%v * 2 = %v \n\n", fractionalPart, m)
			if strings.HasPrefix(fmt.Sprintf("%f", fractionalPart*2), "1") {
				fractSlice = append(fractSlice, "1")
				fractionalPart = m - 1
				//fmt.Printf("m = %v, so next fractional part = %v\n\n", m, fractionalPart)
			} else if strings.HasPrefix(fmt.Sprintf("%f", fractionalPart*2), "0") {
				fractSlice = append(fractSlice, "0")
				fractionalPart = m
			}
		}
		fractionalBinStr := "0." + strings.Join(fractSlice, "")
		fmt.Printf("Fractional Bits: %s\n\n", fractionalBinStr)
		fractionalBin, _ := strconv.ParseFloat(fractionalBinStr, 64)
		bin = intBin + fractionalBin
	}
	return bin
}

//enterNumber() reads from a standard input
func enterNumber(reader *bufio.Reader) string {
	var baseXNum string
	reader = bufio.NewReader(os.Stdin)
	for {
		fmt.Print("$ Enter a number (to exit press X): ")
		n, err := reader.ReadString('\n')
		if err != nil {
			log.Println(err)
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
			baseXNum = n
		}
		return baseXNum
	}
	return baseXNum
}

/*isFloat64() checks if the string is a number*/
func isFloat64(input string) bool {
	_, err := strconv.ParseFloat(input, 64)
	return err == nil
}
func isFloat64Negative(input string) bool {
	if strings.HasPrefix(input, "-") {
		return true
	}
	return false
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

/*
Sample output:

-----------------------
Command Shell
-----------------------
$ Enter a number (to exit press X): 9.2

Converting 9.2 base-10 to binary representation..

STEP 1: Converting integer part: 9

Upper bound on No. of operations required ~ 1*log(10)/log(2) = 3

Operation 1: 9/2=4 with remainder 1

Operation 2: 4/2=2 with remainder 0

Operation 3: 2/2=1 with remainder 0

Operation 4: 1/2=0 with remainder 1

Integer Bits: 1001

STEP 2: Converting fractional part: 0.1999999999999993

0.1999999999999993 * 2 = 0.3999999999999986

Fractional Bits: 0.0

The binary representation of 9.2 base-10 is 1001

Converting to floating point representation...

The floating-point representation of 1001 is "1.001000000x 2^(3)"*/
