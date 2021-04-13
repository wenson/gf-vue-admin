package api

import (
	"gf-vue-admin/app/api/system/internal"
	"gf-vue-admin/app/model/system/request"
	service "gf-vue-admin/app/service/system"
	"gf-vue-admin/library/response"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/util/gconv"
)

var Admin = new(admin)

type admin struct{}

// @Tags SystemAdmin
// @Summary 用户注册账号
// @Produce  application/json
// @Param data body request.Register true "用户名, 昵称, 密码, 角色ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"注册成功"}"
// @Router /user/register [post]
func (a *admin) Register(r *ghttp.Request) *response.Response {
	var info request.Register
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorAdminRegister}
	}
	if err := service.Admin.Register(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorAdminRegister}
	}
	return &response.Response{MessageCode: response.SuccessAdminRegister}
}

// @Tags SystemAdmin
// @Summary 用户修改密码
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body request.ChangePassword true "用户名, 原密码, 新密码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /user/changePassword [put]
func (a *admin) ChangePassword(r *ghttp.Request) *response.Response {
	var info request.ChangePassword
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorChangePassword}
	}
	info.Uuid = internal.Info.GetUserUuid(r)
	if err := service.Admin.ChangePassword(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorChangePassword}
	}
	return &response.Response{MessageCode: response.SuccessChangePassword}
}

// @Tags SystemAdmin
// @Summary 分页获取用户列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "页码, 每页大小"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /user/getUserList [post]
func (a *admin) GetList(r *ghttp.Request) *response.Response {
	var info request.PageInfo
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorGetList}
	}
	list, total, err := service.Admin.GetAdminList(&info)
	if err != nil {
		return &response.Response{MessageCode: response.ErrorGetList, Error: err}
	}
	return &response.Response{Data: response.PageResult{
		List:     list,
		Total:    total,
		Page:     info.Page,
		PageSize: info.PageSize,
	}, MessageCode: response.SuccessGetList}
}

// @Tags SystemAdmin
// @Summary 设置用户权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SetAuthority true "用户UUID, 角色ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /user/setUserAuthority [post]
func (a *admin) SetAuthority(r *ghttp.Request) *response.Response {
	var info request.SetAuthority
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorSetAuthority}
	}
	if err := service.Admin.SetUserAuthority(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorSetAuthority}
	}
	return &response.Response{MessageCode: response.SuccessSetAuthority}
}

// @Tags SystemAdmin
// @Summary 删除用户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "用户ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /user/deleteUser [delete]
func (a *admin) Delete(r *ghttp.Request) *response.Response {
	var info request.GetById
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorDeleted}
	}
	if err := service.Admin.Delete(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorDeleted}
	}
	return &response.Response{MessageCode: response.SuccessDeleted}
}

// @Tags SystemAdmin
// @Summary 设置用户权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.UpdateAdmin true "用户UUID, 角色ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /user/setUserAuthority [post]
func (a *admin) Update(r *ghttp.Request) *response.Response {
	var info request.UpdateAdmin
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorSetAdminInfo}
	}
	info.Uuid = internal.Info.GetUserUuid(r)
	if data, err := service.Admin.SetAdminInfo(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorSetAdminInfo}
	} else {
		return &response.Response{Data: g.Map{"userInfo": data}, MessageCode: response.SuccessSetAdminInfo}
	}
}


//获得当前用户
//

func (a *admin) CurrentUser(r *ghttp.Request) *response.Response {
	
	token, _ := Auth.ParseToken(r); 

	claims := gconv.Map(token.Claims)
	uuid := gconv.String(claims["admin_uuid"])
	
	data, _ := service.Admin.FindAdmininfo(&request.GetByUuid{Uuid: uuid});
	glog.Debug(data)
	return &response.Response{Data: data}	

	
}
