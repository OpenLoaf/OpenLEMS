package auth

import (
	v1 "application/api/auth/v1"
	"common/c_log"
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

// Logout 退出登录
func (c *Controller) Logout(ctx context.Context, req *v1.LogoutReq) (res *v1.LogoutRes, err error) {
	r := g.RequestFromCtx(ctx)
	if r == nil {
		return nil, gerror.NewCode(gcode.CodeInternalError, "请求上下文无效")
	}
	role := r.Session.MustGet("role").String()
	_ = r.Session.RemoveAll()
	c_log.BizInfof(ctx, "用户退出 - 角色:%s", role)
	return &v1.LogoutRes{}, nil
}
