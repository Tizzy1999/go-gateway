package http_proxy_router

import (
	"github.com/gin-gonic/gin"
	"go_gateway_demo/http_proxy_middleware"
)

func InitRouter(middlewares ...gin.HandlerFunc) *gin.Engine {
	router := gin.New()
	//router := gin.Default()
	router.Use(middlewares...)
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Use(http_proxy_middleware.HTTPAccessModeMiddleware())
	return router
}
