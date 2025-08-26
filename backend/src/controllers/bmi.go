package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joaooliveira247/fullstack_bmi_calculator/backend/src/schemas"
	"github.com/joaooliveira247/fullstack_bmi_calculator/backend/src/utils"
)

type BMIController struct{}

func NewBMIController() *BMIController {
	return &BMIController{}
}

func (ctrl *BMIController) GetBMI(ctx *gin.Context) {
	var bmi schemas.BMIIn

	if err := ctx.ShouldBindJSON(&bmi); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	measures := utils.Measures(bmi)

	ctx.JSON(
		http.StatusOK,
		&schemas.BMIOut{BMI: measures.BMI(), Message: measures.Category()},
	)
}
