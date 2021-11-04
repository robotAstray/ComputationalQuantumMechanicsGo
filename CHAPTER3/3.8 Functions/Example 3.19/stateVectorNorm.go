//stateVectorNorm.go normalises a two-basis state vector Ψ =|0> + i|1> such that <Ψ|Ψ> = Sum|Ψi|^2 =1
//Usage:./stateVectorNorm
//Copyright (c) robotAstray

package main

import (
	"fmt"
	"log"
	"math"
)

func main() {
	// var Ψ = 0 + 1i

	var Ψ = complex(0, 1)
	log.Printf("initialising a two basis state vector Ψ: %v", Ψ)
	var Ψ_vector = []float64{real(Ψ), imag(Ψ)}
	N := Normalise(Ψ_vector)
	fmt.Printf("The normalise vector is %v/%v", Ψ, N)
}

//Calculate the vector norm of Ψ =|0> + i|1>
func Normalise(Ψ_vector []float64) (N float64) {
	var sum float64
	for _, base := range Ψ_vector {
		absSqr := math.Pow(2, math.Abs(base))
		sum += absSqr
	}
	N = math.Sqrt(sum)

	return N
}
