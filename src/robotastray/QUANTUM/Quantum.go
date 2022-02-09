package Quantum 

import (
	"math"
	"fmt"
)

func NormaliseStateVector(psi complex128)(normalisedPsi complex128){
	re := real(psi) //real part 
	im := imag(psi) //imaginary part 
	var components = []float64{re, im}
	var sumAbsComponent float64
	for _, component := range components{
		sumAbsComponent += math.Pow(math.Abs(component),2)
	}
	var N float64
	N = math.Sqrt(sumAbsComponent)
	fmt.Printf("N: %v\n", N)
	re = real(psi)/N
	im = imag(psi)/N
	normalisedPsi = complex(re, im)
	return normalisedPsi
}