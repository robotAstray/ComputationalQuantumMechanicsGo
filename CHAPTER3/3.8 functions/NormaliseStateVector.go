package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"math"
)


func main(){
	//define flags 
	realStr := flag.String("real", "", "real part")
	imagStr := flag.String("img", "", "imaginary part, enter just a number without i or j")
	flag.Parse()
	requiredFlags := []string{"real", "img"}
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
	if *realStr != "" && *imagStr != ""{
		//convert between string to float64
		real, err := strconv.ParseFloat(*realStr, 64)
		if err != nil{
			fmt.Println("This is not a number! Start Again")
			goodbye()
			os.Exit(1)
		}
		imaginary, err := strconv.ParseFloat(*imagStr, 64)
		if err != nil{
			fmt.Println("This is not a number! Start Again")
			goodbye()
			os.Exit(1)
		}
		psi := complex(real,imaginary)

		normalisedPsi := NormaliseStateVector(psi)
		fmt.Printf("Normilesed Wave Function: %#v\n", normalisedPsi)
	}
}

func goodbye(){
	fmt.Printf("\nGoodbye\n")
	fmt.Printf("Copyright (c) 2022 RobotAstray\n")
}


func NormaliseStateVector(psi complex128) (normalisedPsi complex128){
	re := real(psi) //real part
	im := imag(psi) //imaginary part 
	var waveComponents = []float64{re, im}
    var sumAbsComponent float64
	for _, component := range waveComponents{
		sumAbsComponent += math.Pow(math.Abs(component),2)
	}
	var N float64
 	N = math.Sqrt(sumAbsComponent)
	fmt.Println(N)
	re = real(psi)/N
	im = imag(psi)/N
	normalisedPsi = complex(re,im)

	return normalisedPsi
}