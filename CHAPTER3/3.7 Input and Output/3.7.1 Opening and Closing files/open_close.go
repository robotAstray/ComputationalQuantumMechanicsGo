//open_close.go opens and closes a file
//Usage: go build open_close.go; ./open_close -path /folder/file.txt
//Copyright (c) 2021 rndMemex

package main

import (

	"os"
	"log"
	"fmt"
	"flag"
)


func main(){
//declare flag
	pathToFile := flag.String("path", " ", "path to file (e.g. /folder/file.txt)")
	requiredFlags := []string{"path"}
	flag.Parse()
	seenFlag := make(map[string]bool)
	flag.Visit(func(f *flag.Flag){
		seenFlag[f.Name] = true
	})

	for _, reqFlag := range requiredFlags{
		if !seenFlag[reqFlag]{
			log.Print("ERROR!! Missing required flag: %s", reqFlag)
			os.Exit(1)
		}
	}
	if *pathToFile != " "{
		//Open File from path provided
		file, err := os.OpenFile(*pathToFile, os.O_APPEND, 0666)
		if err != nil{
			log.Fatal(err)
		}
		log.Printf("%s was opened", *pathToFile)
		file.Close()
		log.Printf("%s was closed", *pathToFile)
		goodbye()
	}
}

func goodbye(){
	fmt.Printf("\nGoodbye!\n")
	fmt.Printf("Copyright (c) 2021 rndMemex\n\n")
}