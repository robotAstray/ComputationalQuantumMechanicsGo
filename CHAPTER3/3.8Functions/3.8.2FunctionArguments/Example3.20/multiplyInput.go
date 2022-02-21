package main

import (
	"flag"
)


func main(){
	numbers := flag.String("num", "", "numbers separated by a comma")
	requiredFlags := []string{"num"}
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
	if *numbers{
		if IsInputNumberArray(*numbers){
			result := multiplyNumbers(*numbers)
		}
	}

}
func IsInputNumberArray(*numbers) (isNum bool, arrayInt []int){
	isNum  = true
	var arrayInputStr []string{}
	arrayInputStr = strings.Split(str, ",")
	for _, element := range arrayInputStr{
		
	}

 
}

func multiplyNumbers(*args) (mult float64){

}