package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joaooliveira247/fullstack_bmi_calculator/backend/src/routes"
)

func main() {
	api := gin.Default()
	routes.RegisterRoutes(api)

	if err := api.Run(":8000"); err != nil {
		log.Fatal(err)
	}
}
