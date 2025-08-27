package main

import (
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/joaooliveira247/fullstack_bmi_calculator/backend/docs"
	"github.com/joaooliveira247/fullstack_bmi_calculator/backend/src/routes"
)

// @title BMI API
// @version 1.0
// @description BMI Calculator API.
// @host localhost:8000
// @BasePath /
func main() {
	api := gin.Default()
	routes.RegisterRoutes(api)

	if err := api.Run(":8000"); err != nil {
		log.Fatal(err)
	}
}
