package dao

import (
	"errors"
	"github.com/e421083458/golang_common/lib"
	"github.com/gin-gonic/gin"
	"go_gateway_demo/dto"
	"go_gateway_demo/public"
	"net/http/httptest"
	"strings"
	"sync"
)

type ServiceDetail struct {
	Info          *ServiceInfo   `json:"info" description:"基本信息"`
	HTTPRule      *HttpRule      `json:"http_rule" description:"http_rule"`
	TCPRule       *TcpRule       `json:"tcp_rule" description:"tcp_rule"`
	GRPCRule      *GrpcRule      `json:"grpc_rule" description:"grpc_rule"`
	LoadBalance   *LoadBalance   `json:"load_balance" description:"load_balance"`
	AccessControl *AccessControl `json:"access_control" description:"access_control"`
}

// 让ServiceManager以handler的形式暴露出去
var ServiceManagerHandler = NewServiceManager()

func init() {
	ServiceManagerHandler = NewServiceManager()
}

type ServiceManager struct {
	ServiceMap   map[string]*ServiceDetail //通过map方式获得服务详情
	ServiceSlice []*ServiceDetail          //通过遍历数组的方式获得服务详情
	Locker       sync.RWMutex
	init         sync.Once //初始化一次
	err          error
}

func NewServiceManager() *ServiceManager {
	return &ServiceManager{
		ServiceMap:   map[string]*ServiceDetail{},
		ServiceSlice: []*ServiceDetail{},
		Locker:       sync.RWMutex{},
		init:         sync.Once{},
	}
}

// http服务接入方式
func (s *ServiceManager) HTTPAccessMode(c *gin.Context) (*ServiceDetail, error) {
	/**
	匹配方式
	1）前缀匹配 /abc ==> serviceSlice.rule
	2) 域名匹配 www.test.com ==> serviceSlice.rule

	host c.Request.Host
	path c.Request.URL.Path
	*/
	host := c.Request.Host // www.test.com:8080
	host = host[0:strings.Index(host, ":")]
	//fmt.Println("host", host)
	path := c.Request.URL.Path
	for _, serviceItem := range s.ServiceSlice {
		if serviceItem.Info.LoadType != public.LoadTypeHTTP {
			continue
		}
		// 域名匹配
		if serviceItem.HTTPRule.RuleType == public.HTTPRuleTypeDomain {
			if serviceItem.HTTPRule.Rule == host {
				return serviceItem, nil
			}
		}
		// 前缀匹配
		if serviceItem.HTTPRule.RuleType == public.HTTPRuleTypePrefixURL {
			if strings.HasPrefix(path, serviceItem.HTTPRule.Rule) {
				return serviceItem, nil
			}
		}
	}
	return nil, errors.New("no matched service")
}

// 把所有服务载入内存
func (s *ServiceManager) LoadOnce() error {
	// 只会执行一次，有锁的机制
	s.init.Do(func() {
		// 从db中分页读取信息
		serviceInfo := &ServiceInfo{}
		c, _ := gin.CreateTestContext(httptest.NewRecorder())
		tx, err := lib.GetGormPool("default")
		if err != nil {
			s.err = err
			return
		}
		params := &dto.ServiceListInput{PageNo: 1, PageSize: 99999}
		list, _, err := serviceInfo.PageList(c, tx, params)
		if err != nil {
			s.err = err
			return
		}
		for _, listItem := range list {

			tmpItem := listItem
			serviceDetail, err := tmpItem.ServiceDetail(c, tx, &tmpItem)
			//serviceDetail, err := listItem.ServiceDetail(c, tx, &listItem)
			//fmt.Println("serviceDetail")
			//fmt.Println(public.Obj2Json(serviceDetail))
			if err != nil {
				s.err = err
				return
			}
			s.ServiceMap[listItem.ServiceName] = serviceDetail
			s.ServiceSlice = append(s.ServiceSlice, serviceDetail)
		}
	})
	return s.err
}
