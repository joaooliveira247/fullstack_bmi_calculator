package schemas

type BMIIn struct {
	Weight float64 `json:"weight" binding:"required,gt=0"`
	Height float64 `json:"height" binding:"required,gt=0"`
}

type BMIOut struct {
	BMI     float64 `json:"BMI"`
	Status  string  `json:"status"`
	Message string  `json:"message"`
}
