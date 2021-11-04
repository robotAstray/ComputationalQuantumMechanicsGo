//calculateSum.go reads numbers from a file and calculates the sum
//Usage: go build calculateSum.go; ./calculateSum -path /folder/readme.txt
//Copyright (c) 2021 robotAstray

package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"robotAstray"
	"strconv"
)

func checkIsInputNumber(content string) (isNum bool) {
	//default is true
	isNum = true
	_, err := strconv.Atoi(content)
	if err != nil {
		isNum = false
		log.Printf("This entry is not a number\n")
	}
	return isNum
}

func main() {
	//declare flags
	pathToFile := flag.String("path", " ", "path to file (e.g. /folder/filename.txt)")
	requiredFlags := []string{"path"}
	flag.Parse()
	seenFlag := make(map[string]bool)
	flag.Visit(func(f *flag.Flag) {
		seenFlag[f.Name] = true
	})
	for _, reqFlag := range requiredFlags {
		if !seenFlag[reqFlag] {
			log.Printf("ERROR!Missing required flag: %s\n", reqFlag)
			os.Exit(1)
		}
	}
	if *pathToFile != " " {
		//open file
		file, err := os.Open(*pathToFile)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		var sumNum int
		for scanner.Scan() {
			content := scanner.Text()
			if checkIsInputNumber(content) {
				num, _ := strconv.Atoi(content)
				sumNum += num
			}
		}
		fmt.Printf("The Sum is %d\n", sumNum)
		robotAstray.Goodbye()
	}
}
