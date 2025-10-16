package v1

import "github.com/gogf/gf/v2/frame/g"

type LogoutReq struct {
	g.Meta `path:"/auth/logout" method:"post" tags:"认证" summary:"用户退出" role:"user"`
}
type LogoutRes struct{}

type ChangePasswordReq struct {
	g.Meta      `path:"/auth/change-password" method:"post" tags:"认证" summary:"修改密码" role:"user"`
	Role        string `json:"role" v:"required|in:user,admin#角色不能为空|角色必须是user或admin"`
	OldPassword string `json:"oldPassword" v:"required#旧密码不能为空"`
	NewPassword string `json:"newPassword" v:"required#新密码不能为空"`
}
type ChangePasswordRes struct{}

type GetCurrentUserReq struct {
	g.Meta `path:"/auth/current" method:"get" tags:"认证" summary:"获取当前用户信息"`
}
type GetCurrentUserRes struct {
	Role string `json:"role"`
}
