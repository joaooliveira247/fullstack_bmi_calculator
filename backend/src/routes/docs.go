package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func DocsRoutes(eng *gin.Engine) {
	docsGroup := eng.Group("/docs")
	{
		docsGroup.GET("", func(c *gin.Context) {
			c.Redirect(http.StatusMovedPermanently, "/docs/index.html")
		})
		docsGroup.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
}
