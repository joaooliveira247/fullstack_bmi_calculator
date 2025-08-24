package utils

import "math"

type Measures struct {
	Weigth float64
	Heigth float64
}

func (measures *Measures) BMI() float64 {
	bmi := measures.Weigth / (measures.Heigth * measures.Heigth)
	return math.Round(bmi*100) / 100
}
