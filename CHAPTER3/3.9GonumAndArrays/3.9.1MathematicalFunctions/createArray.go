package main 


import(
	"fmt"
	"gonum.org/v1/gonum/mat"
)

func main(){
	//declare one dimensional array
	array := mat.NewDense(1,6, nil)
	fmt.Printf("array: %v", array)
}