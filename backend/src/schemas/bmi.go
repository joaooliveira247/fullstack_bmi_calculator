package schemas

type BMIIn struct {
	Weigth float64 `json:"weigth" binding:"required,gt=0"`
	Heigth float64 `json:"weight" binding:"required,gt=0"`
}

type BMIOut struct {
	BMI     float64 `json:"BMI"`
	Message string  `json:"message"`
}
