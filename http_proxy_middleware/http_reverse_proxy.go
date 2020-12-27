package http_proxy_middleware

import (
	"github.com/gin-gonic/gin"
)

// 基于请求信息 匹配接入方式
func HTTPReverseProxyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 创建reverseproxy，使用reverseproxy.ServerHTTP(c.Request, c.Response)
	}
}
