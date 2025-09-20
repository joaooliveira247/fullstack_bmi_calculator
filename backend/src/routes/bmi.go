package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joaooliveira247/fullstack_bmi_calculator/backend/src/controllers"
)

func BMIRoutes(eng *gin.Engine) {
	controller := controllers.NewBMIController()

	bmiGroup := eng.Group("/bmi")

	bmiGroup.Use(cors.New(
		cors.Config{
			AllowOrigins:     []string{"http://localhost:3000"},
			AllowMethods:     []string{"POST"},
			AllowHeaders:     []string{"Origin", "Content-Type"},
			AllowCredentials: true,
		},
	))

	{
		bmiGroup.POST("/", controller.GetBMI)
	}
}
