package main

import (
	"flag"
	"fmt"
	"github.com/e421083458/golang_common/lib"
	"go_gateway_demo/dao"
	"go_gateway_demo/http_proxy_router"
	"go_gateway_demo/router"
	"os"
	"os/signal"
	"syscall"
)

/**
endpoint dashboard后台管理 server代理服务器
config ./conf/prod/ 对应配置文件夹
*/
var (
	endpoint = flag.String("endpoint", "", "input endpoint dashboard or server")
	config   = flag.String("config", "", "input config file like ./conf/dev/")
)

func main() {
	flag.Parse() //参数配置
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
		// 加载服务列表到内存中
		dao.ServiceManagerHandler.LoadOnce()
		// 利用协程，可同时启动多个服务器http/https
		go func() {
			http_proxy_router.HttpServerRun()
		}()
		go func() {
			http_proxy_router.HttpsServerRun()
		}()
		fmt.Println("start server")
		quit := make(chan os.Signal)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit
		fmt.Println("server stopping")
		http_proxy_router.HttpServerStop()
		http_proxy_router.HttpsServerStop()
	}

}
