package auth

import (
	v1 "application/api/auth/v1"
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

// GetCurrentUser 获取当前登录用户信息
func (c *Controller) GetCurrentUser(ctx context.Context, req *v1.GetCurrentUserReq) (res *v1.GetCurrentUserRes, err error) {
	r := g.RequestFromCtx(ctx)
	if r == nil {
		return nil, gerror.NewCode(gcode.CodeInternalError, "请求上下文无效")
	}
	role := r.Session.MustGet("role").String()
	if role == "" {
		return &v1.GetCurrentUserRes{Role: "guest"}, nil
	}
	return &v1.GetCurrentUserRes{Role: role}, nil
}
