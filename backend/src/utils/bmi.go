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

func (measures *Measures) Category() string {
	bmi := measures.BMI()

	switch {
	case bmi < 16:
		return "Severe thinness ⚠️ 🪦"
	case bmi >= 16 && bmi < 17:
		return "Modarate thinness ⚠️ 🏃‍♂️"
	case bmi >= 17 && bmi < 18.5:
		return "Mild thinness ⚠️ 🍵"
	case bmi >= 18.5 && bmi < 25:
		return "Normal ✅ 🟢"
	case bmi >= 25 && bmi < 30:
		return "Overweight ⚠️ 🍔"
	case bmi >= 30 && bmi < 35:
		return "Obese class I 🔴 ⚠️"
	case bmi >= 35 && bmi < 40:
		return "Obese class II 🔴🚨"
	case bmi >= 40:
		return "Obese class III 🛑 🏥"
	default:
		return ""
	}
}
