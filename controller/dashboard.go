package controller

import (
	"errors"
	"github.com/e421083458/golang_common/lib"
	"github.com/gin-gonic/gin"
	"go_gateway_demo/dao"
	"go_gateway_demo/dto"
	"go_gateway_demo/middleware"
	"go_gateway_demo/public"
	"time"
)

type DashboardController struct{}

func DashboardRegister(group *gin.RouterGroup) {
	service := &DashboardController{}
	group.GET("/panel_group_data", service.PanelGroupData)
	group.GET("/flow_stat", service.FlowStat)
	group.GET("/service_stat", service.ServiceStat)
}

// PanelGroupData godoc
// @Summary panel group data
// @Description panel group data
// @Tags Dashboard
// @ID /dashboard/panel_group_data
// @Accept  json
// @Produce  json
// @Success 200 {object} middleware.Response{data=dto.PanelGroupDataOutput} "success"
// @Router /dashboard/panel_group_data [get]
func (service *DashboardController) PanelGroupData(c *gin.Context) {
	tx, err := lib.GetGormPool("default")
	if err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}

	serviceInfo := &dao.ServiceInfo{}
	// serviceNum 为service总数
	_, serviceNum, err := serviceInfo.PageList(c, tx, &dto.ServiceListInput{PageSize: 1, PageNo: 1})
	if err != nil {
		middleware.ResponseError(c, 2002, err)
		return
	}
	app := &dao.App{}
	// appNum 为app总数
	_, appNum, err := app.APPList(c, tx, &dto.APPListInput{PageNo: 1, PageSize: 1})
	if err != nil {
		middleware.ResponseError(c, 2003, err)
		return
	}
	counter, err := public.FlowCounterHandler.GetCounter(public.FlowTotal)
	if err != nil {
		middleware.ResponseError(c, 2003, err)
		return
	}
	out := &dto.PanelGroupDataOutput{
		ServiceNum:      serviceNum,
		AppNum:          appNum,
		TodayRequestNum: counter.TotalCount,
		CurrentQPS:      counter.QPS,
	}
	middleware.ResponseSuccess(c, out)
}

// FlowStat godoc
// @Summary service flow data
// @Description service flow data
// @Tags Dashboard
// @ID /dashboard/flow_stat
// @Accept  json
// @Produce  json
// @Success 200 {object} middleware.Response{data=dto.ServiceStatOutput} "success"
// @Router /dashboard/flow_stat [get]
func (service *DashboardController) FlowStat(c *gin.Context) {
	counter, err := public.FlowCounterHandler.GetCounter(public.FlowTotal)
	if err != nil {
		middleware.ResponseError(c, 2001, err)
	}
	todayList := []int64{}
	currentTime := time.Now()
	for i := 0; i <= currentTime.Hour(); i++ {
		datetTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), i, 0, 0, 0, lib.TimeLocation)
		hourDate, _ := counter.GetHourData(datetTime)
		todayList = append(todayList, hourDate)
	}
	yesterdayList := []int64{}
	// 从现在开始倒退24个小时
	yesterTime := currentTime.Add(-1 * time.Duration(time.Hour*24))
	for i := 0; i <= 23; i++ {
		datetTime := time.Date(yesterTime.Year(), yesterTime.Month(), yesterTime.Day(), i, 0, 0, 0, lib.TimeLocation)
		hourData, _ := counter.GetHourData(datetTime)
		yesterdayList = append(yesterdayList, hourData)
	}
	middleware.ResponseSuccess(c, &dto.ServiceStatOutput{
		Today:     todayList,
		Yesterday: yesterdayList,
	})
}

// ServiceStat godoc
// @Summary service percentage
// @Description service percentage
// @Tags Dashboard
// @ID /dashboard/service_stat
// @Accept  json
// @Produce  json
// @Success 200 {object} middleware.Response{data=dto.DashServiceStatOutput} "success"
// @Router /dashboard/service_stat [get]
func (service *DashboardController) ServiceStat(c *gin.Context) {
	tx, err := lib.GetGormPool("default")
	if err != nil {
		middleware.ResponseError(c, 2001, err)
	}
	serviceInfo := &dao.ServiceInfo{}
	list, err := serviceInfo.GroupByLoadType(c, tx)
	if err != nil {
		middleware.ResponseError(c, 2002, err)
	}
	legend := []string{}
	for index, item := range list {
		name, ok := public.LoadTypeMap[item.LoadType]
		if !ok {
			middleware.ResponseError(c, 2003, errors.New("load_type not found"))
		}
		list[index].Name = name
		legend = append(legend, name)
	}
	out := &dto.DashServiceStatOutput{
		Legend: legend,
		Series: list,
	}
	middleware.ResponseSuccess(c, out)
}
