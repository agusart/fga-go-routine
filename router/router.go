package router

import (
	"github.com/gin-gonic/gin"
	"weather/handler"
)

func StartServer(filePath string) *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("template/**")
	router.GET("/", handler.Home(filePath))

	return router
}
