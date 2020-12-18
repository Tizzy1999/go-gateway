package controller

import (
	"github.com/gin-gonic/gin"
	"go_gateway_demo/dto"
	"go_gateway_demo/middleware"
)

type ServiceController struct {
}

func ServiceRegister(group *gin.RouterGroup) {
	service := &ServiceController{}
	// 请求前缀 + action方法
	group.GET("/service_list", service.ServiceList)
}

// Service godoc
// @Summary 服务列表
// @Description 服务列表
// @Tags 服务管理
// @ID /service/service_list
// @Accept  json
// @Produce  json
// @Param info query string false "关键词"
// @Param page_no query int true "当前页数"
// @Param page_size query int true "每页个数"
// @Success 200 {object} middleware.Response{data=dto.ServiceListOutput} "success"
// @Router /service/service_list [get]
func (service *ServiceController) ServiceList(c *gin.Context) {
	params := &dto.ServiceListInput{}
	if err := params.BindValidParam(c); err != nil {
		middleware.ResponseError(c, 2000, err)
	}
	// 1. 读取sessionKey对应json转换为结构体
	// 2. 取出数据然后封装输出结构体
	//ID          int64  `json:"id" form:"id"`                     //id
	//ServiceName string `json:"service_name" form:"service_name"` //服务名称
	//ServiceDesc string `json:"service_desc" form:"service_desc"` //服务描述
	//LoadType    int    `json:"load_type" form:"load_type"`       //类型
	//ServiceAddr string `json:"service_addr" form:"service_addr"` //服务地址
	//Qps         int64  `json:"qps" form:"qps"`                   //qps
	//Qpd         int64  `json:"qpd" form:"qpd"`                   //qpd
	//TotalNode   int    `json:"total_node" form:"total_node"`     //节点数
	out := &dto.ServiceListOutput{}
	middleware.ResponseSuccess(c, out)
}
