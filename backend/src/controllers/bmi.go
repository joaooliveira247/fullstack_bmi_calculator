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

// GetBMI Returns BMI send in JSON
// @Summary Returns BMI
// @Description Returns BMI send in JSON
// @Tags BMI
// @Accept json
// @Produce json
// @Param data body schemas.BMIIn true "Widht and Height to calcule BMI"
// @Success 200 {object} schemas.BMIOut
// @Failure 400 {object} map[string]string
// @Router /bmi/ [post]
func (ctrl *BMIController) GetBMI(ctx *gin.Context) {
	var bmi schemas.BMIIn

	if err := ctx.ShouldBindJSON(&bmi); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	measures := utils.Measures(bmi)

	category := measures.Category()

	ctx.JSON(
		http.StatusOK,
		&schemas.BMIOut{
			BMI:     measures.BMI(),
			Status:  category.Status,
			Message: category.Message,
		},
	)
}
