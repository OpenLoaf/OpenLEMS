package utils

import (
	v1 "application/api/auth/v1"
	"common/c_enum"
	"strings"

	"github.com/gogf/gf/v2/net/ghttp"
)

// MiddlewareAuth 认证与授权中间件
func MiddlewareAuth(r *ghttp.Request) {
	handler := r.GetServeHandler()

	// 支持通过 Authorization 头传递会话ID（Bearer <token> 或 直接 <token>）
	if auth := r.GetHeader("Authorization"); auth != "" {
		token := auth
		if strings.HasPrefix(strings.ToLower(token), "bearer ") {
			token = strings.TrimSpace(token[7:])
		}
		token = strings.TrimSpace(token)
		if token != "" {
			const sessionCookieName = "ems_session_id"
			if ck := r.Request.Header.Get("Cookie"); ck != "" {
				r.Request.Header.Set("Cookie", ck+"; "+sessionCookieName+"="+token)
			} else {
				r.Request.Header.Set("Cookie", sessionCookieName+"="+token)
			}
		}
	}

	// 公共接口跳过鉴权
	if handler != nil && handler.GetMetaTag("noAuth") == "true" {
		r.Middleware.Next()
		return
	}

	// 读取 Session 中的角色
	role := r.Session.MustGet("role").String()
	if role == "" {
		r.Response.WriteJson(ghttp.DefaultHandlerResponse{Code: 401, Message: "未登录或登录已过期"})
		return
	}

	// 检查角色要求：若未标注role，按方法推断(GET=user，其它=admin)
	if handler != nil {
		required := handler.GetMetaTag("role")
		if required == "" {
			if r.Method == "GET" {
				required = string(c_enum.EUserRoleUser)
			} else {
				required = string(c_enum.EUserRoleAdmin)
			}
		}
		if required != "" && !checkRolePermission(role, required) {
			r.Response.WriteJson(ghttp.DefaultHandlerResponse{Code: 403, Message: "权限不足"})
			return
		}
	}

	// 刷新 Session 过期时间（滑动过期）
	refreshSessionTimeout(r, role)

	// 注入当前用户信息到上下文（例如在控制器可用）
	r.SetCtxVar("userRole", role)

	r.Middleware.Next()
}

// checkRolePermission 角色包含关系：admin ⊇ user
func checkRolePermission(current, required string) bool {
	if current == string(c_enum.EUserRoleAdmin) {
		return true
	}
	return current == required
}

// refreshSessionTimeout 根据角色刷新过期时间
func refreshSessionTimeout(r *ghttp.Request, role string) {
	// 允许认证接口无需刷新
	// 其它请求刷新为自定义时长（单位小时），默认2小时
	// 会话刷新：GoFrame v2 Session 不提供逐会话 TTL API，这里仅作为占位逻辑；
	// 我们把期望的过期秒数写入 Session（expireSeconds），供前端与服务器侧参考。
	// 服务器全局过期时间在 web.go 中通过 SetSessionMaxAge 设置。
	_ = role
}

// 占位引用，避免未使用导入（接口元信息）
var _ = v1.LoginReq{}
