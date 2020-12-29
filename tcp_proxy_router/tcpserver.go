package tcp_proxy_router

import (
	"context"
	"fmt"
	"go_gateway_demo/dao"
	"go_gateway_demo/tcp_proxy_middleware"
	"go_gateway_demo/tcp_server"
	"log"
	"net"
)

// 全局切片，这样可以被开启和关闭函数调用
var tcpServerList = []*tcp_server.TcpServer{}

type tcpHandler struct {
}

// 用于测试tcp服务器启动的简易handler
func (t *tcpHandler) ServeTCP(ctx context.Context, src net.Conn) {
	src.Write([]byte("tcpHandler by Tizzy\n"))
}

func TcpServerRun() {
	serviceList := dao.ServiceManagerHandler.GetTcpServiceList()
	for _, serviceItem := range serviceList {
		tempItem := serviceItem
		// 利用协程并行启动多个程序v
		go func(serviceDetail *dao.ServiceDetail) {
			addr := fmt.Sprintf(":%d", serviceDetail.TCPRule.Port)
			// 暂时关闭反向代理， 使用简易handler
			//rb, err := dao.LoadBalancerHandler.GetLoadBalancer(serviceDetail)
			//if err != nil{
			//	log.Fatalf(" [INFO] GetTcpLoadBalancer %v err:%v\n", addr, err)
			//	return
			//}
			//构建路由及设置中间件
			router := tcp_proxy_middleware.NewTcpSliceRouter()
			router.Group("/").Use(
				tcp_proxy_middleware.TCPFlowCountMiddleware(),
				tcp_proxy_middleware.TCPFlowLimitMiddleware(),
				tcp_proxy_middleware.TCPWhiteListMiddleware(),
				tcp_proxy_middleware.TCPBlackListMiddleware(),
			)

			//todo 构建回调handler
			routerHandler := tcp_proxy_middleware.NewTcpSliceRouterHandler(
				func(c *tcp_proxy_middleware.TcpSliceRouterContext) tcp_server.TCPHandler {
					//return reverse_proxy.NewTcpLoadBalanceReverseProxy(c, rb)
					return &tcpHandler{}
				}, router)
			// 下游context
			baseCtx := context.WithValue(context.Background(), "service", serviceDetail)
			tcpServer := &tcp_server.TcpServer{
				Addr:    addr,
				Handler: routerHandler,
				BaseCtx: baseCtx,
			}
			tcpServerList = append(tcpServerList, tcpServer)
			log.Printf(" [INFO] tcp_proxy_run %v\n", addr)
			// 非关闭连接的错误才中止进程， fatal相当与panic
			if err := tcpServer.ListenAndServe(); err != nil && err != tcp_server.ErrServerClosed {
				//中止进程直到端口空出来
				log.Fatalf(" [INFO] tcp_proxy_run %v err:%v\n", addr, err)
			}
		}(tempItem)
	}
}

func TcpServerStop() {
	for _, tcpServer := range tcpServerList {
		tcpServer.Close()
		log.Printf(" [INFO] tcp_proxy_stop %v stopped\n", tcpServer.Addr)
	}
}
