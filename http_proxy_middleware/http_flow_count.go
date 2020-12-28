package http_proxy_middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go_gateway_demo/dao"
	"go_gateway_demo/middleware"
	"go_gateway_demo/public"
)

// 基于请求信息 匹配接入方式
func HTTPFlowCountMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		serviceInterface, ok := c.Get("service")
		if !ok {
			middleware.ResponseError(c, 2001, errors.New("service not found"))
			c.Abort()
			return
		}
		// 统计项 1 全站 2 服务 3 租户
		serviceDetail := serviceInterface.(*dao.ServiceDetail)
		totalCounter, err := public.FlowCounterHandler.GetCounter(public.FlowTotal)
		if err != nil {
			middleware.ResponseError(c, 4001, err)
			c.Abort()
			return
		}
		totalCounter.Increase()

		///dayCount, _ := totalCounter.GetDayData(time.Now())
		//fmt.Printf("totalCounter qps: %v, dayCount: %v \n", totalCounter.QPS, dayCount)
		serviceCounter, err := public.FlowCounterHandler.GetCounter(public.FlowServicePrefix + serviceDetail.Info.ServiceName)
		if err != nil {
			middleware.ResponseError(c, 4002, err)
			c.Abort()
			return
		}
		serviceCounter.Increase()
		//dayServiceCount, _ := serviceCounter.GetDayData(time.Now())
		//fmt.Printf("serviceCounter qps: %v, dayCount: %v \n", serviceCounter.QPS, dayServiceCount)
		c.Next()
	}
}
