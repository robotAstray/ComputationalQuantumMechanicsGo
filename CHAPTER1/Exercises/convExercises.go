//convExercises.go - base-2 to base-10 calculator and viceverse from commandline
//USAGE: e.g. .\convExercises.exe -base10 23.625 -conv
//Copyright (c) 2020 rndMemex

package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var sign string

func binaryCalculator(baseXNum string) (bin float64) {
	fmt.Printf("\nConverting %s base-10 to binary representation..\n\n", baseXNum)
	ratio := int(math.Log10(10) / math.Log10(2))

	if !strings.Contains(baseXNum, ".") {
		reps := len(baseXNum) * ratio //upper bound on operations needed
		fmt.Printf("Upper bound on No. of operations required ~ %d*log(10)/log(2)\n\n", len(baseXNum))
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
		if sign == "negative" {
			bin = -bin
		}
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
		//the following should be improved
		for k := 0; k < len(fractionalNum); k++ {
			m = fractionalPart * 2
			fmt.Printf("%v * 2 = %v \n\n", fractionalPart, m)
			if strings.HasPrefix(fmt.Sprintf("%f", fractionalPart*2), "1") {
				fractSlice = append(fractSlice, "1")
				fractionalPart = m - 1
			} else if strings.HasPrefix(fmt.Sprintf("%f", fractionalPart*2), "0") {
				fractSlice = append(fractSlice, "0")
				fractionalPart = m
			}
		}
		fractionalBinStr := "0." + strings.Join(fractSlice, "")
		fmt.Printf("Fractional Bits: %s\n\n", fractionalBinStr)
		fractionalBin, _ := strconv.ParseFloat(fractionalBinStr, 64)
		bin = intBin + fractionalBin
		if sign == "negative" {
			bin = -bin
		}
	}

	return bin
}
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
func main() {
	binary := flag.String("bin", "", "binary number e.g. '1001'")
	//precision := flag.String("precision", "", "fixed point or floating point precision, e.g. 'floating', 'fixed'")
	baseX := flag.String("base10", "", "base-10 number e.g. '10.2")
	conv := flag.Bool("conv", false, "convert")
	mandatoryFlags := []string{"conv"}
	flag.Parse()
	exists := make(map[string]bool)
	flag.Visit(func(f *flag.Flag) { exists[f.Name] = true })
	for _, flag := range mandatoryFlags {
		if !exists[flag] && (!exists["baseX"] || !exists["bin"]) {
			fmt.Fprintf(os.Stderr, "missing required -%s argument/flag\n", flag)
			os.Exit(2)
		}
	}
	if *conv {
		if *binary != "" {
			if isBin(*binary) {
				baseX := baseIIConverter(*binary)
				fmt.Printf("The Binary number %s is %#v in base-10\n\n", *binary, baseX)
			} else {
				fmt.Printf("\n%s is not a binary number\n\n", *binary)
			}
		} else if *baseX != "" {
			if isFloat64(*baseX) {
				bin := binaryCalculator(*baseX)
				if sign == "negative" {
					fmt.Printf("The binary representation of -%s base-10 is -%#v\n\n", *baseX, bin)
				} else {
					fmt.Printf("The binary representation of %s base-10 is %#v\n\n", *baseX, bin)
				}
			} else {
				fmt.Printf("ERROR!! %s is not a number!\n\n", *baseX)
			}
		}
	}

}

func goodbye() {
	fmt.Printf("\nGoodbye!\n")
	fmt.Printf("Copyright(c) 2020 rndmemex@cantab.net\n")
}

/*isFloat64() checks if the string is a number*/
func isFloat64(input string) bool {
	_, err := strconv.ParseFloat(input, 64)
	if strings.HasPrefix(input, "-") {
		sign = "negative"
	} else {
		sign = "positive"
	}
	return err == nil
}

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

/*
--------------------------------------
SAMPLE OUTPUT
-------------------------------------
.\convExercises.exe -base10 23.625 -conv

Converting 23.625 base-10 to binary representation..

STEP 1: Converting integer part: 23

Upper bound on No. of operations required ~ 2*log(10)/log(2) = 6

Operation 1: 23/2=11 with remainder 1

Operation 2: 11/2=5 with remainder 1

Operation 3: 5/2=2 with remainder 1

Operation 4: 2/2=1 with remainder 0

Operation 5: 1/2=0 with remainder 1

Integer Bits: 10111

STEP 2: Converting fractional part: 0.625

0.625 * 2 = 1.25

0.25 * 2 = 0.5

0.5 * 2 = 1

Fractional Bits: 0.101

The binary representation of 23.625 base-10 is 10111.101
*/
