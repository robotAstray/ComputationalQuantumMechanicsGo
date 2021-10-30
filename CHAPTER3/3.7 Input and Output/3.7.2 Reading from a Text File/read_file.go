//read_file.go reads from a text file
//Usage: go build read_file.go; ./read_file -path ./folder/file.txt
//Copyright (c) 2021 rndMemex 

package main


import (
	"fmt"
	"log"
	"os"
	"flag"
)

func main(){
	//declare flags
	pathToFile := flag.String("path", "", "path to file (e.g. /folder/file.txt")
	requiredFlags := []string{"path"}
	flag.Parse()
	seenFlag :=make(map[string]bool)
	flag.Visit(func (f *flag.Flag){
		seenFlag[f.Name] = true
	})
	for _, reqFlag := range requiredFlags{
		if !seenFlag[reqFlag]{
			log.Printf("ERROR!Missing mandatory flag:%s", reqFlag)
			os.Exit(1)
		}
	}
	if *pathToFile != ""{
		file , err := os.ReadFile(*pathToFile)
		if err != nil{
			log.Fatal(err)
		}
		fmt.Println(string(file))
		goodbye()
	}
}

func goodbye(){
	fmt.Printf("\nGoodbye\n")
	fmt.Printf("Copyright (c) 2021 RobotAstray\n")
}