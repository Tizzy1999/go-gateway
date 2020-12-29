package main

import (
	"flag"
	"fmt"
	"github.com/e421083458/golang_common/lib"
	"go_gateway_demo/dao"
	"go_gateway_demo/grpc_proxy_router"
	"go_gateway_demo/http_proxy_router"
	"go_gateway_demo/router"
	"go_gateway_demo/tcp_proxy_router"
	"os"
	"os/signal"
	"syscall"
)

/**
endpoint dashboard 后台管理 server 代理服务器
config ./conf/prod/ 对应配置文件夹
*/
var (
	endpoint = flag.String("endpoint", "", "input endpoint dashboard or server")
	config   = flag.String("config", "", "input config file like ./conf/dev/")
)

func main() {
	flag.Parse()
	if *endpoint == "" {
		flag.Usage()
		os.Exit(1)
	}
	if *config == "" {
		flag.Usage()
		os.Exit(1)
	}
	if *endpoint == "dashboard" {
		// 如果 configPath 为空，则从命令行中 config = ./conf/prod/ 中读取
		lib.InitModule(*config, []string{"base", "mysql", "redis"})
		defer lib.Destroy()
		router.HttpServerRun()

		quit := make(chan os.Signal)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit

		router.HttpServerStop()
	} else {
		lib.InitModule(*config, []string{"base", "mysql", "redis"})
		defer lib.Destroy()
		dao.ServiceManagerHandler.LoadOnce()
		dao.AppManagerHandler.LoadOnce()
		// 利用协程，启动多个服务
		go func() {
			http_proxy_router.HttpServerRun()
		}()
		go func() {
			http_proxy_router.HttpsServerRun()
		}()
		go func() {
			tcp_proxy_router.TcpServerRun()
		}()
		go func() {
			grpc_proxy_router.GrpcServerRun()
		}()
		fmt.Println("start server")
		quit := make(chan os.Signal)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit
		fmt.Println("server stopping")
		grpc_proxy_router.GrpcServerStop()
		tcp_proxy_router.TcpServerStop()
		http_proxy_router.HttpServerStop()
		http_proxy_router.HttpsServerStop()

	}

}
