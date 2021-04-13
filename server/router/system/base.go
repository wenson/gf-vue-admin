package router

import (
	api "gf-vue-admin/app/api/system"
	"gf-vue-admin/interfaces"
	"gf-vue-admin/library/response"

	"github.com/gogf/gf/net/ghttp"
)

type base struct {
	router   *ghttp.RouterGroup
	response *response.Handler
}

func NewBaseGroup(router *ghttp.RouterGroup) interfaces.Router {
	return &base{router: router, response: &response.Handler{}}
}

func (b *base) Init() {
	group := b.router.Group("/api")
	{
		group.POST("/account/captcha", b.response.Handler()(api.Base.Captcha))
		group.POST("/account/login", api.Auth.LoginHandler) // 登录
		group.GET("/account/logout", api.Auth.LogoutHandler) // 登出
	}
}
