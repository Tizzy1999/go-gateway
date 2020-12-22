package dto

import (
	"github.com/gin-gonic/gin"
	"go_gateway_demo/public"
	"time"
)

type AdminInfoOutput struct {
	ID           int       `json:"id"`
	UserName     string    `json:"name"`
	LoginTime    time.Time `json:"login_time"`
	Avatar       string    `json:"avatar"`
	Introduction string    `json:"introduction"`
	Roles        []string  `json:"roles"`
}

type ChangePwdInput struct {
	Password string `json:"password" form:"password" comment:"密码" example:"123456" validate:"required"`
}

// 绑定结构体，校验标签的required
func (params *ChangePwdInput) BindingValidParams(c *gin.Context) error {
	return public.DefaultGetValidParams(c, params)
}
