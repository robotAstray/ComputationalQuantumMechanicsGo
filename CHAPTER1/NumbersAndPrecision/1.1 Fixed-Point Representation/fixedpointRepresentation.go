//fixedpointRep2baseX.go - base-10 to base-2 calculator
//Usage: go build fixedpointRep2baseX.go; ./fixedpointRep2baseX
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

func XVIfixedPointRep(bin float64) string {
	binStr := fmt.Sprintf("%f", bin)
	//	fmt.Printf("binStr %s\n", binStr)
	var XVIPointRepStr string
	N := 16 //Total Number of bits
	I := 4  //number of integer bits
	if !strings.Contains(binStr, ".") {
		if len(binStr) < I {
			zeros := I - len(binStr)
			fmt.Printf("I - len(%v) = %d", bin, zeros)
			for i := 0; i < zeros; i++ {
				binStr = "0" + binStr
				fmt.Printf("Adding zeros: %s", binStr)
			}
			fractionalZeros := N - I
			binStr = binStr + "."
			fmt.Printf("Adding a decimal point: %s", binStr)
			for k := 0; k < fractionalZeros; k++ {
				binStr = binStr + "0"
			}
			fmt.Printf("16-bit fixed Point Rep: %s", binStr)
			// XVIPointRep, _ = strconv.ParseFloat(binStr, 64)
			XVIPointRepStr = binStr
			return XVIPointRepStr

		} else if len(binStr) == I {
			fractionalZeros := N - I

			XVIPointRepStr = binStr + "."
			for i := 0; i < fractionalZeros; i++ {
				XVIPointRepStr = XVIPointRepStr + "0"
			}

			return XVIPointRepStr
		}
		return XVIPointRepStr
	} else {
		index := strings.Index(binStr, ".")
		fmt.Println(index)
		integerBits := binStr[:index]
		fmt.Printf("integerBits %s\n", integerBits)
		fractionalBits := binStr[index+1:]
		fmt.Printf("fractionalBits %s\n", fractionalBits)
		if len(integerBits) < I {
			zeros := I - len(integerBits)
			for i := 0; i < zeros; i++ {
				integerBits = "0" + integerBits
			}
			fractionalZeros := N - I - len(fractionalBits)
			for k := 0; k < fractionalZeros; k++ {
				fractionalBits = fractionalBits + "0"
			}

			XVIPointRepStr = integerBits + "." + fractionalBits

		} else if len(integerBits) == I {
			fractionalZeros := N - I - len(fractionalBits)
			for k := 0; k < fractionalZeros; k++ {
				fractionalBits = fractionalBits + "0"
			}
			XVIPointRepStr = integerBits + "." + fractionalBits

		} else {
			fmt.Printf("Truncation error or Round-off error\n")
			goodbye()
			os.Exit(0)
		}
		return XVIPointRepStr
	}
	return XVIPointRepStr
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

func main() {
	var inputNum *bufio.Reader
	fmt.Println("-----------------------")
	fmt.Println("Command Shell")
	fmt.Println("-----------------------")
	for {
		baseX := enterNumber(inputNum)
		bin := binaryCalculator(baseX)
		fmt.Printf("The binary representation of %s base-10 is %#v\n\n", baseX, bin)
		fmt.Printf("Converting to 16-bit fixed point representation...\n\n")
		XVIFixPointRep := XVIfixedPointRep(bin)
		fmt.Printf("The 16-bit fixed-point representation of %v is %#v\n\n", bin, XVIFixPointRep)
	}

}

func goodbye() {
	fmt.Printf("\nGoodbye!\n")
	fmt.Printf("Copyright(c) 2020 rndmemex@cantab.net\n")
}
