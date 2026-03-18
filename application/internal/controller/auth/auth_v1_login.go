package auth

import (
	v1 "application/api/auth/v1"
	"common/c_enum"
	"common/c_log"
	"common/c_util"
	"context"
	"s_db"
	"s_db/s_db_basic"
	"time"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

// Login 用户登录
func (c *Controller) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	role := req.Role
	var settingDef *s_db_basic.SSystemSettingDefine
	switch role {
	case string(c_enum.EUserRoleAdmin):
		settingDef = s_db_basic.SystemSettingAdminPassword
	case string(c_enum.EUserRoleUser):
		settingDef = s_db_basic.SystemSettingUserPassword
	default:
		return nil, gerror.NewCode(gcode.CodeInvalidParameter, "角色不支持")
	}

	// 获取密码并比对，兼容旧版明文存储。
	pwdPtr := s_db.GetSettingService().GetSettingValueBySystemSettingDefine(ctx, settingDef)
	storedPassword := ""
	if pwdPtr != nil {
		storedPassword = *pwdPtr
	}
	ok, err := c_util.VerifyPassword(storedPassword, req.Password)
	if err != nil {
		c_log.Errorf(ctx, "用户登录密码校验失败 - 角色:%s, 错误:%+v", role, err)
		return nil, gerror.NewCode(gcode.CodeInternalError, "密码校验失败")
	}
	if !ok {
		c_log.BizWarningf(ctx, "用户登录失败 - 角色:%s", role)
		return nil, gerror.NewCode(gcode.CodeNotAuthorized, "用户名或密码错误")
	}
	if storedPassword != "" && !c_util.IsPasswordHash(storedPassword) {
		tryUpgradePasswordHash(ctx, role, settingDef.Id, req.Password)
	}

	// 读取会话过期时间（小时）
	timeoutHours := 2
	if role == string(c_enum.EUserRoleAdmin) {
		if v := s_db.GetSettingService().GetSettingValueBySystemSettingDefine(ctx, s_db_basic.SystemSettingSessionAdminTimeout); v != nil {
			if n := g.Cfg().MustGet(ctx, "int:"+*v).Int(); n > 0 {
				timeoutHours = n
			}
		}
	} else {
		if v := s_db.GetSettingService().GetSettingValueBySystemSettingDefine(ctx, s_db_basic.SystemSettingSessionUserTimeout); v != nil {
			if n := g.Cfg().MustGet(ctx, "int:"+*v).Int(); n > 0 {
				timeoutHours = n
			}
		}
	}

	// 写入 Session
	r := g.RequestFromCtx(ctx)
	if r == nil {
		return nil, gerror.NewCode(gcode.CodeInternalError, "请求上下文无效")
	}
	r.Session.Set("role", role)
	expireSeconds := timeoutHours * 3600
	r.Session.Set("expireSeconds", expireSeconds)

	// 业务日志
	c_log.BizInfof(ctx, "用户登录成功 - 角色:%s", role)

	// 返回 Token 信息（SessionID）
	token, _ := r.Session.Id()
	expireAt := time.Now().Add(time.Duration(expireSeconds) * time.Second).Unix()

	return &v1.LoginRes{Token: token, Role: role, ExpireAt: expireAt}, nil
}

func tryUpgradePasswordHash(ctx context.Context, role string, settingID string, plaintext string) {
	hashedPassword, err := c_util.HashPassword(plaintext)
	if err != nil {
		c_log.Warningf(ctx, "登录后升级密码哈希失败 - 角色:%s, 错误:%+v", role, err)
		return
	}
	if err := s_db.GetSettingService().SetSettingValueById(ctx, settingID, hashedPassword); err != nil {
		c_log.Warningf(ctx, "登录后保存密码哈希失败 - 角色:%s, 错误:%+v", role, err)
		return
	}
	c_log.Infof(ctx, "已完成密码哈希升级 - 角色:%s", role)
}
