package http_proxy_middleware

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go_gateway_demo/dao"
	"go_gateway_demo/middleware"
	"go_gateway_demo/public"
	"strings"
)

// 基于请求信息 匹配接入方式
func HTTPBlackListMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		serviceInterface, ok := c.Get("service")
		if !ok {
			middleware.ResponseError(c, 2001, errors.New("service not found"))
			c.Abort()
			return
		}
		serviceDetail := serviceInterface.(*dao.ServiceDetail)
		//因为白名单优先于黑名单，所以只有白名单为空的情况才考虑黑名单
		//whiteIpList := []string{}
		//if serviceDetail.AccessControl.WhiteList != "" {
		//	whiteIpList = strings.Split(serviceDetail.AccessControl.WhiteList, ",")
		//}
		blackIpList := []string{}
		if serviceDetail.AccessControl.BlackList != "" {
			blackIpList = strings.Split(serviceDetail.AccessControl.BlackList, ",")
		}
		if serviceDetail.AccessControl.OpenAuth == 1 && len(blackIpList) > 0 {
			if public.InStringSlice(blackIpList, c.ClientIP()) {
				middleware.ResponseError(c, 3001, errors.New(fmt.Sprintf("%s in black ip List", c.ClientIP())))
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
