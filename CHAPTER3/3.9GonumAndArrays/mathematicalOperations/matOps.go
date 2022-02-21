package main


import (
	"gonum.org/v1/gonum/mat"
	"fmt"
)

func main(){
    //cross product
	a := mat.NewDense(3,3, []float64{
		2,3,4,
		4,5,6,
		5,6,7,
	})
	b := mat.NewDense(3,3, []float64{
		2,3,4,
		4,5,6,
		5,6,7,
	})
	var c mat.Dense
	c.Add(a, b)
	fmt.Printf("Addition:\n %v\n", mat.Formatted(&c, mat.Prefix("   "), mat.Squeeze()))
	var d mat.Dense
	d.Sub(a, b)
	fmt.Printf("Subtraction:\n %v\n", mat.Formatted(&d, mat.Prefix("   "), mat.Squeeze()))
	var e mat.Dense
	e.Mul(a,b)
	fmt.Printf("Multiplication:\n %v\n", mat.Formatted(&e, mat.Prefix("   "), mat.Squeeze()))
	var f mat.Dense
	f.DivElem(a,b)
	fmt.Printf("Division:\n %v\n", mat.Formatted(&f, mat.Prefix("   "), mat.Squeeze()))
	var g mat.Dense
	g.Exp(a)
	fmt.Printf("Exponential:\n %v\n", mat.Formatted(&g, mat.Prefix("   "), mat.Squeeze()))
	var h mat.Dense
	h.Pow(a, 0)
	fmt.Printf("a power of 0:\n %v\n", mat.Formatted(&h, mat.Prefix("   "), mat.Squeeze()))

}