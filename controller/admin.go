package controller

import (
	"encoding/json"
	"fmt"
	"github.com/e421083458/golang_common/lib"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"go_gateway_demo/dao"
	"go_gateway_demo/dto"
	"go_gateway_demo/middleware"
	"go_gateway_demo/public"
)

type AdminController struct {
}

func AdminRegister(group *gin.RouterGroup) {
	adminLogin := &AdminController{}
	// 请求前缀 + action方法
	group.GET("/admin_info", adminLogin.AdminInfo)
	group.POST("/change_pwd", adminLogin.ChangePwd)
}

// Admin godoc
// @Summary 管理员信息
// @Description 管理员信息
// @Tags 管理员接口
// @ID /admin/admin_info
// @Accept  json
// @Produce  json
// @Success 200 {object} middleware.Response{data=dto.AdminLoginOutput} "success"
// @Router /admin/admin_info [get]
func (admininfo *AdminController) AdminInfo(c *gin.Context) {
	sess := sessions.Default(c)
	sessInfo := sess.Get(public.AdminSessionInfoKey)
	adminSessionInfo := &dto.AdminSessionInfo{}
	if err := json.Unmarshal([]byte(fmt.Sprint(sessInfo)), adminSessionInfo); err != nil {
		middleware.ResponseError(c, 2000, err)
		return
	}
	// 1. 读取sessionKey对应json转换为结构体
	// 2. 取出数据然后封装输出结构体
	out := &dto.AdminInfoOutput{
		ID:           adminSessionInfo.ID,
		UserName:     adminSessionInfo.UserName,
		LoginTime:    adminSessionInfo.LoginTime,
		Avatar:       "https://www.flaticon.com/svg/static/icons/svg/194/194938.svg",
		Introduction: "This is administrator Tizzy",
		Roles:        []string{"admin"},
	}
	middleware.ResponseSuccess(c, out)
}

// ChangePwd godoc
// @Summary 修改密码
// @Description 修改密码
// @Tags 管理员接口
// @ID /admin/change_pwd
// @Accept  json
// @Produce  json
// @Param polygon body dto.ChangePwdInput true "body"
// @Success 200 {object} middleware.Response{data=string} "success"
// @Router /admin/change_pwd [post]
func (admininfo *AdminController) ChangePwd(c *gin.Context) {
	// 定义接收参数用到的结构体
	params := &dto.ChangePwdInput{}
	// 进行数据绑定，将接收的参数绑定到param中
	if err := params.BindingValidParams(c); err != nil {
		middleware.ResponseError(c, 2000, err)
		return
	}

	// 1. 通过 session 读取信息到结构体 sessInfo
	// 2. sessInfo.ID 读取数据库信息 adminInfo
	// 3. params.password+adminInfo.salt sha256 saltPassword
	// 4. saltPassword ==> adminInfo.password 执行数据保存

	// session读取用户信息到结构体
	sess := sessions.Default(c)
	sessInfo := sess.Get(public.AdminSessionInfoKey)
	adminSessionInfo := &dto.AdminSessionInfo{}
	if err := json.Unmarshal([]byte(fmt.Sprint(sessInfo)), adminSessionInfo); err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}
	adminInfo := &dao.Admin{}
	// 从数据库中读取 adminInfo
	tx, err := lib.GetGormPool("default")
	if err != nil {
		middleware.ResponseError(c, 2002, err)
		return
	}
	adminInfo, err = adminInfo.Find(c, tx, &dao.Admin{UserName: adminSessionInfo.UserName})
	if err != nil {
		middleware.ResponseError(c, 2003, err)
		return
	}
	// 生成新密码 saltPassword
	saltPassword := public.GenSaltPassword(adminInfo.Salt, params.Password)
	adminInfo.Password = saltPassword
	// 执行数据保存
	if err = adminInfo.Save(c, tx); err != nil {
		middleware.ResponseError(c, 2004, err)
		return
	}
	middleware.ResponseSuccess(c, "")
}
