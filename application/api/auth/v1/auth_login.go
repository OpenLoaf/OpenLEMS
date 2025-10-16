package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type LoginReq struct {
	g.Meta   `path:"/auth/login" method:"post" tags:"认证" summary:"用户登录" noAuth:"true"`
	Password string `json:"password" v:"required#密码不能为空" dc:"密码"`
	Role     string `json:"role" v:"required|in:user,admin#角色不能为空|角色必须是user或admin" dc:"用户角色：user-普通用户, admin-管理员"`
}

type LoginRes struct {
	Token    string `json:"token" dc:"Session Token"`
	Role     string `json:"role" dc:"用户角色"`
	ExpireAt int64  `json:"expireAt" dc:"过期时间戳（秒）"`
}
