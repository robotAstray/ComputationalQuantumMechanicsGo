package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	Quantum "github.com/robotAstray/ComputationalQuantumMechanicsGo/src/robotastray/QUANTUM"
	robotAstray "github.com/robotAstray/ComputationalQuantumMechanicsGo/src/robotastray/GREETINGS"
)

func main(){
	//declare flags
	realStr := flag.String("real", "", "real part")
	imagStr := flag.String("img", "", "imaginary part, enter just a number without i or j")
	flag.Parse()
	requiredFlags := []string{"real", "img"}
	seenFlag := make(map[string]bool)
    flag.Visit(func (f *flag.Flag){
		seenFlag[f.Name] = true
	})
	for _, reqFlag := range requiredFlags{
		if !seenFlag[reqFlag]{
			log.Printf("ERROR!!Missing mandatory flag: %s", reqFlag)
			robotAstray.Goodbye()
			os.Exit(1)
		}
	}
	if *realStr != "" && *imagStr != ""{
		//convert from string to float64
		real, err := strconv.ParseFloat(*realStr, 64)
		if err != nil{
			fmt.Println("This is not a number! Start Again")
			robotAstray.Goodbye()
			os.Exit(1)
		}
		imaginary, err := strconv.ParseFloat(*imagStr, 64)
		if err != nil{
			fmt.Println("This is not a number! Start Again")
			robotAstray.Goodbye()
			os.Exit(1)
		}
		psi := complex(real, imaginary)

		normalisedPsi := Quantum.NormaliseStateVector(psi)
		fmt.Printf("Normalised Wave Function: %#v\n", normalisedPsi)
	}
}

