package http_proxy_middleware

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go_gateway_demo/dao"
	"go_gateway_demo/middleware"
	"go_gateway_demo/public"
)

// 基于请求信息 匹配接入方式
func HTTPJwtFlowCountMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 拿不到就说明没有开启Jwt验证
		//fmt.Printf("This is jwt flow count middleware")
		appInterface, ok := c.Get("app")
		if !ok {
			c.Next()
			return
		}
		appInfo := appInterface.(*dao.App)
		appCounter, err := public.FlowCounterHandler.GetCounter(public.FlowAppPrefix + appInfo.AppID)
		if err != nil {
			middleware.ResponseError(c, 2002, err)
			c.Abort()
			return
		}
		appCounter.Increase()
		if appInfo.Qpd > 0 && appCounter.TotalCount > appInfo.Qpd {
			middleware.ResponseError(c, 2003, errors.New(fmt.Sprintf("App qpd limit:%v current:%v", appInfo.Qpd, appCounter.TotalCount)))
			c.Abort()
			return
		}
		//fmt.Printf("appCounter qps: %v, dayCount: %v",appCounter.QPS, appCounter.TotalCount)
		c.Next()
	}
}
