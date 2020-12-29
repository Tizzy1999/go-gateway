package grpc_proxy_middleware

import (
	"go_gateway_demo/dao"
	"go_gateway_demo/public"
	"google.golang.org/grpc"
	"log"
)

// 基于请求信息 匹配接入方式
func GrpcFlowCountMiddleware(serviceDetail *dao.ServiceDetail) func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		totalCounter, err := public.FlowCounterHandler.GetCounter(public.FlowTotal)
		if err != nil {
			return err
		}
		totalCounter.Increase()

		///dayCount, _ := totalCounter.GetDayData(time.Now())
		//fmt.Printf("totalCounter qps: %v, dayCount: %v \n", totalCounter.QPS, dayCount)
		serviceCounter, err := public.FlowCounterHandler.GetCounter(public.FlowServicePrefix + serviceDetail.Info.ServiceName)
		if err != nil {
			return err
		}
		serviceCounter.Increase()
		if err := handler(srv, ss); err != nil {
			log.Printf("GrpcFlowCountMiddleware failed with error %v\n", err)
			return err
		}
		//dayServiceCount, _ := serviceCounter.GetDayData(time.Now())
		//fmt.Printf("serviceCounter qps: %v, dayCount: %v \n", serviceCounter.QPS, dayServiceCount)
		return nil
	}
}
