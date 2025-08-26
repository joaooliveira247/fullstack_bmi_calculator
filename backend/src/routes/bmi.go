package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/joaooliveira247/fullstack_bmi_calculator/backend/src/controllers"
)

func BMIRoutes(eng *gin.Engine) {
	controller := controllers.NewBMIController()

	bmiGroup := eng.Group("/bmi")
	{
		bmiGroup.POST("/", controller.GetBMI)
	}
}
