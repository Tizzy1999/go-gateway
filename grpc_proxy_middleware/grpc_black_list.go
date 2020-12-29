package grpc_proxy_middleware

import (
	"fmt"
	"github.com/pkg/errors"
	"go_gateway_demo/dao"
	"go_gateway_demo/public"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
	"log"
	"strings"
)

//匹配接入方式 基于请求信息
func GrpcBlackListMiddleware(serviceDetail *dao.ServiceDetail) func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		whiteIplist := []string{}
		if serviceDetail.AccessControl.WhiteList != "" {
			whiteIplist = strings.Split(serviceDetail.AccessControl.WhiteList, ",")
		}
		blackIplist := []string{}
		if serviceDetail.AccessControl.BlackList != "" {
			blackIplist = strings.Split(serviceDetail.AccessControl.BlackList, ",")
		}

		peerCtx, ok := peer.FromContext(ss.Context())
		if !ok {
			return errors.New("peer not found with context")
		}
		peerAddr := peerCtx.Addr.String()
		addrPos := strings.LastIndex(peerAddr, ":")
		clientIP := peerAddr[0:addrPos]
		if serviceDetail.AccessControl.OpenAuth == 1 && len(whiteIplist) == 0 && len(blackIplist) > 0 {
			if public.InStringSlice(blackIplist, clientIP) {
				return errors.New(fmt.Sprintf("%s in black ip list", clientIP))
			}
		}
		if err := handler(srv, ss); err != nil {
			log.Printf("RPC failed with error %v\n", err)
			return err
		}
		return nil
	}
}
