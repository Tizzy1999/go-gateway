package controller

import (
	"github.com/gin-gonic/gin"
	"go_gateway_demo/dto"
	"go_gateway_demo/middleware"
)

type AdminLoginController struct {
}

func AdminLoginRegister(group *gin.RouterGroup) {
	adminLogin := &AdminLoginController{}
	// 请求前缀 + action方法
	group.POST("/login", adminLogin.AdminLogin)
}

// ListPage godoc
// @Summary 管理员登陆
// @Description 管理员登陆
// @Tags 管理员接口
// @ID /admin_login/login
// @Accept  json
// @Produce  json
// @Param polygon body dto.AdminLoginInput true "body"
// @Success 200 {object} middleware.Response{data=dto.AdminLoginOutput} "success"
// @Router /admin_login/login [post]
func (adminlogin *AdminLoginController) AdminLogin(c *gin.Context) {
	params := &dto.AdminLoginInput{}
	if err := params.BindingValidParams(c); err != nil {
		middleware.ResponseError(c, 1001, err)
		return
	}
	out := &dto.AdminLoginOutput{Token: params.UserName}
	middleware.ResponseSuccess(c, out)
}
