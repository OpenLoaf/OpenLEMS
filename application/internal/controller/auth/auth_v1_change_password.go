package auth

import (
	v1 "application/api/auth/v1"
	"common/c_enum"
	"common/c_log"
	"common/c_util"
	"context"
	"s_db"
	"s_db/s_db_basic"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

// ChangePassword 修改密码
func (c *Controller) ChangePassword(ctx context.Context, req *v1.ChangePasswordReq) (res *v1.ChangePasswordRes, err error) {
	// 从 Session 读取当前角色，若传入角色与当前不一致，按当前角色处理
	r := g.RequestFromCtx(ctx)
	if r == nil {
		return nil, gerror.NewCode(gcode.CodeInternalError, "请求上下文无效")
	}
	currentRole := r.Session.MustGet("role").String()
	role := req.Role
	if role == "" {
		role = currentRole
	}

	var def *s_db_basic.SSystemSettingDefine
	switch role {
	case string(c_enum.EUserRoleAdmin):
		def = s_db_basic.SystemSettingAdminPassword
	case string(c_enum.EUserRoleUser):
		def = s_db_basic.SystemSettingUserPassword
	default:
		return nil, gerror.NewCode(gcode.CodeInvalidParameter, "角色不支持")
	}

	hashedPassword, err := c_util.HashPassword(req.NewPassword)
	if err != nil {
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "生成密码哈希失败")
	}

	// 更新新密码
	if err := s_db.GetSettingService().SetSettingValueById(ctx, def.Id, hashedPassword); err != nil {
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "更新密码失败")
	}

	c_log.BizInfof(ctx, "用户修改密码 - 角色:%s", role)
	return &v1.ChangePasswordRes{}, nil
}
