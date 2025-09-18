package utils

import "math"

type Measures struct {
	Weight float64
	Height float64
}

type BMICategory struct {
	Status  string
	Message string
}

func (measures *Measures) BMI() float64 {
	bmi := measures.Weight / (measures.Height * measures.Height)
	return math.Round(bmi*100) / 100
}

func (measures *Measures) Category() BMICategory {
	bmi := measures.BMI()

	switch {
	case bmi < 16:
		return BMICategory{
			Status:  "Severe thinness",
			Message: "Very low BMI ⚠️🪦 — severe health risk, seek urgent medical help.",
		}
	case bmi >= 16 && bmi < 17:
		return BMICategory{
			Status:  "Moderate thinness",
			Message: "Low BMI ⚠️🏃‍♂️ — moderate malnutrition risk, pay attention to nutrition.",
		}
	case bmi >= 17 && bmi < 18.5:
		return BMICategory{
			Status:  "Mild thinness",
			Message: "Slightly below normal ⚠️🍵 — improve diet to avoid deficiencies.",
		}
	case bmi >= 18.5 && bmi < 25:
		return BMICategory{
			Status:  "Normal",
			Message: "Healthy BMI ✅🟢 — keep up the good habits.",
		}
	case bmi >= 25 && bmi < 30:
		return BMICategory{
			Status:  "Overweight",
			Message: "Above the ideal range ⚠️🍔 — higher risk of health issues if maintained.",
		}
	case bmi >= 30 && bmi < 35:
		return BMICategory{
			Status:  "Obese class I",
			Message: "Obesity class I 🔴⚠️ — health risks present, consider professional support.",
		}
	case bmi >= 35 && bmi < 40:
		return BMICategory{
			Status:  "Obese class II",
			Message: "Obesity class II 🔴🚨 — high risk of complications, seek medical guidance.",
		}
	case bmi >= 40:
		return BMICategory{
			Status:  "Obese class III",
			Message: "Obesity class III 🛑🏥 — very high risk, medical treatment is essential.",
		}
	default:
		return BMICategory{
			Status:  "Unknown",
			Message: "Invalid BMI value ❓",
		}
	}
}
