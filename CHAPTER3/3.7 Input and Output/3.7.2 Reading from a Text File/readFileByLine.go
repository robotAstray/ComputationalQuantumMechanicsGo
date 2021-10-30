// opens a text file and reads its content line by line 
//Usage: go build readFileByLine.g; ./readsFileByLine -path ./folder/readme.txt
//Copyright (c) 2021 rndMemex


package main

import (

	"os"
	"log"
	"fmt"
	"bufio"
	"flag"
)


func main(){

	//declare flags

	pathToFile := flag.String("path", " ", "path to file (e.g. /folder/filename.txt)")
	requiredFlags := []string{"path"}
	flag.Parse()
	seenFlag:= make(map[string]bool)
	flag.Visit(func (f *flag.Flag){
		seenFlag[f.Name] = true
	})
	for _, reqFlag := range requiredFlags{
		if !seenFlag[reqFlag]{
			log.Printf("ERROR! Missing required flag: %s\n", reqFlag)
			os.Exit(1)

		}
	}
	if *pathToFile != ""{
		//open file 
		file, err := os.Open(*pathToFile)
		if err != nil{
			log.Fatal(err)
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		i := 0
		for scanner.Scan(){
			//print line
			i++
			fmt.Printf("line #%d:%s\n", i, scanner.Text())
		}
		goodbye()

	}


}


func goodbye(){
	fmt.Printf("\nGoodbye\n")
	fmt.Printf("Copyright (c) 2021 RobotAstray\n")
}