package utils

import (
	"common/c_log"
	"net/http"
	"strings"
	"time"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/guid"
)

// MiddlewareRequestID 为每个请求生成并注入一个请求ID，同时透传到响应头中
func MiddlewareRequestID(r *ghttp.Request) {
	requestID := r.GetHeader("X-Request-Id")
	if requestID == "" {
		requestID = guid.S()
	}
	r.Response.Header().Set("X-Request-Id", requestID)
	r.SetCtxVar("requestId", requestID)
	r.Middleware.Next()
}

// MiddlewareErrorHandler 错误处理中间件
func MiddlewareErrorHandler(r *ghttp.Request) {
	r.Middleware.Next()
	if err := r.GetError(); err != nil {
		// 更详细的错误日志：包含请求方法/路径、请求体、堆栈
		ctx := r.Context()
		if stack := gerror.Stack(err); stack != "" {
			c_log.Errorf(ctx, "HTTP %s %s - Error: %v\nStack:\n%s", r.Method, r.URL.Path, err, stack)
		} else {
			c_log.Errorf(ctx, "HTTP %s %s - Error: %v", r.Method, r.URL.Path, err)
		}
		if body := r.GetBodyString(); body != "" {
			// 仅在调试或出现错误时打印请求体（注意敏感字段）
			c_log.Debugf(ctx, "Request Body: %s", body)
		}

		// 返回统一JSON，尽量带上Code与详细Message
		r.Response.ClearBuffer()
		if gerr, ok := err.(*gerror.Error); ok {
			r.Response.WriteJson(ghttp.DefaultHandlerResponse{
				Code:    gerr.Code().Code(),
				Message: err.Error(),
			})
			return
		}
		r.Response.WriteJson(ghttp.DefaultHandlerResponse{
			Code:    500,
			Message: err.Error(),
		})
	}
}

// MiddlewareAccessLog 记录访问摘要日志，避免与错误日志重复
func MiddlewareAccessLog(r *ghttp.Request) {
	startAt := time.Now()
	r.Middleware.Next()
	used := time.Since(startAt)

	status := r.Response.Status
	method := r.Method
	path := r.URL.Path
	clientIP := r.GetClientIp()
	requestID := r.Response.Header().Get("X-Request-Id")

	// 跳过预检、静态资源与WebSocket升级
	if method == http.MethodOptions {
		return
	}
	if strings.HasPrefix(path, "/assets/") || strings.HasPrefix(path, "/images/") || strings.HasPrefix(path, "/static/") || path == "/favicon.ico" {
		return
	}
	if status == http.StatusSwitchingProtocols {
		return
	}
	// 跳过API路由和WebSocket路由的访问日志
	if strings.HasPrefix(path, "/api/") || strings.HasPrefix(path, "/station") || strings.HasPrefix(path, "/telemetry") {
		return
	}

	// 业务错误(5xx)由错误处理中间件打印详细日志；或请求上有错误时，这里跳过，避免重复
	if status >= 500 || r.GetError() != nil {
		return
	}

	ctx := r.Context()
	msg := "%s %s -> %d | %dms | ip=%s | rid=%s"
	if status >= 400 {
		c_log.Warningf(ctx, msg, method, path, status, used.Milliseconds(), clientIP, requestID)
		return
	}
	c_log.Infof(ctx, msg, method, path, status, used.Milliseconds(), clientIP, requestID)
}
