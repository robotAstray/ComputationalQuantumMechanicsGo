package Quantum 

import ()

func NormaliseStateVector(psi complex64) (normalisedPsi complex64){
	re : = real(psi) //real part
	im := imag(psi) //imaginary part 
	var waveComponents []float64{re, im}
	var N float64
    var sumAbsComponent float64
	for i, component := range waveComponents{
		sumAbsComponent += math.Pow(math.Abs(component),2)
	}
	var N float64
 	N = math.Sqrt(absComponent)
	normalisedPsi = psi/ N

	return normalisedPsi
}