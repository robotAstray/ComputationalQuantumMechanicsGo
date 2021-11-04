//writeToFile.go calculates the 5 times table and outputs the result into a file named '5times.txt'
//Usage: go build writeToFile.go; ./writeToFile
//Copyright (c) robotAstray

package main

import (
	"fmt"
	"log"
	"os"
	"robotAstray"
	"strconv"
)

//STEP 1: Create a file named '5times.txt' in the current directory
//STEP 2: Write the 5 times table to '5times.txt'

func main() {
	file, err := os.Create("5times.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	i := 1
	for i <= 10 {
		product := 5 * i
		num := strconv.Itoa(i)
		productStr := strconv.Itoa(product)

		file.WriteString("5 x " + num + " = " + productStr + "\n")
		i++
	}
	fmt.Print("The file '5times.txt' has been created, check your current directory! \n")
	robotAstray.Goodbye()
}
