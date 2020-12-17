package dto

import (
	"github.com/gin-gonic/gin"
	"go_gateway_demo/public"
)

// 管理员登陆结构体
type AdminLoginInput struct {
	// gin的tag，json为json标签，form为表单标签
	UserName string `json:"username" form:"username" comment:"姓名" example:"admin" validate:"is_valid_username"`
	Password string `json:"password" form:"password" comment:"密码" example:"123456" validate:"required"`
}

// 绑定结构体，校验标签
func (params *AdminLoginInput) BindingValidParams(c *gin.Context) error {
	return public.DefaultGetValidParams(c, params)
}

type AdminLoginOutput struct {
	// 和前端整合的时候需要用token
	Token string `json:"token" form:"token" comment:"token" example:"token" validate:""`
}
