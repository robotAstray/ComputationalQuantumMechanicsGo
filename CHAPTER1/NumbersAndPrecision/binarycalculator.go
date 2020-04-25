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

	if !strings.Contains(baseXNum, ".") {
		reps := len(baseXNum) * ratio //upper bound on operation needed
		fmt.Printf("Upper bound on No. of operations required ~ %d*ln(10)/ln(2) = %d\n\n", len(baseXNum), reps)
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
		integerReps := len(integerNum) * ratio //upper bound on operation needed
		fmt.Printf("STEP 1: Converting integer part: %s\n\n", integerNum)
		fmt.Printf("Upper bound on No. of operations required ~ %d*ln(10)/ln(2) = %d\n\n", len(integerNum), integerReps)
		for i := 0; i < integerReps; i++ {
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
		fmt.Printf("Integer Bin: %v\n\n", intBin)

		//"Converting factional part...\n\n"
		fractionalNum := baseXNum[index+1:]
		inputF64, _ := strconv.ParseFloat(baseXNum, 64)
		integerNumF64, _ := strconv.ParseFloat(integerNum, 64)
		fractionalPart := inputF64 - integerNumF64
		fmt.Printf("STEP 2: Converting factional part: %v\n\n", fractionalPart)
		var fractSlice []string
		var m float64
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
		fmt.Printf("Fractional Bin: %s\n\n", fractionalBinStr)
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
	fmt.Printf("\nGoodbye!\n")
	fmt.Printf("Copyright(c) 2020 rndmemex@cantab.net\n")
}

/* Sample Output for input = 23

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

/* Sample Output for input = 23.625
-----------------------
Command Shell
-----------------------
$ enter a number (to exit press X): 23.625

Converting 23.625 base-10 to binary representation..

STEP 1: Converting integer part: 23

Upper bound on No. of operations required ~ 2*ln(10)/ln(2) = 6

Operation 1: 23/2=11 with remainder 1

Operation 2: 11/2=5 with remainder 1

Operation 3: 5/2=2 with remainder 1

Operation 4: 2/2=1 with remainder 0

Operation 5: 1/2=0 with remainder 1

Integer Bin: 10111

STEP 2: Converting factional part: 0.625

0.625 * 2 = 1.25

0.25 * 2 = 0.5

0.5 * 2 = 1

Fractional Bin: 0.101

The binary representation of 23.625 base-10 is 10111.101

$ enter a number (to exit press X): X

Goodbye!
Copyright(c) 2020 rndmemex@cantab.net


*/
