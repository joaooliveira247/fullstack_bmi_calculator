package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(eng *gin.Engine) {
	BMIRoutes(eng)
}
