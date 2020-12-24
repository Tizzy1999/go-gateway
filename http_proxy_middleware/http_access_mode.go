package http_proxy_middleware

import (
	"github.com/gin-gonic/gin"
)

// 基于请求信息 匹配接入方式
func HTTPAccessModeMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
