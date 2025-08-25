package cmd

import (
	"application/internal/consts"
	"application/internal/controller/control"
	"application/internal/controller/device"
	"application/internal/controller/driver"
	"application/internal/controller/log"
	"application/internal/controller/network"
	"application/internal/controller/protocol"
	"application/internal/controller/system"
	"application/internal/controller/telemetry"

	"application/manifest"
	"common/c_base"
	"context"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/i18n/gi18n"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/net/goai"
	"github.com/gogf/gf/v2/util/guid"
)

func startWeb(ctx context.Context) *ghttp.Server {
	ctx = context.WithValue(ctx, c_base.ConstCtxKeyGroupName, "Web")
	// todo 添加关闭服务，允许使用接口关闭web

	g.Log().Infof(ctx, "准备启动web程序！")

	s := g.Server()
	// 中间件顺序：响应封装 -> 永不超时上下文 -> 请求ID -> 错误处理 -> 访问日志
	// 关闭框架内置访问/错误日志，避免与自定义日志重复
	s.SetAccessLogEnabled(false)
	s.SetErrorLogEnabled(false)
	s.Use(
		ghttp.MiddlewareHandlerResponse,
		ghttp.MiddlewareNeverDoneCtx,
		MiddlewareRequestID,
		MiddlewareErrorHandler,
		MiddlewareAccessLog,
	)

	s.Group("/api", func(group *ghttp.RouterGroup) {

		group.Middleware(
			ghttp.MiddlewareCORS,
			func(r *ghttp.Request) {
				httpLanguage := r.GetHeader("Accept-Language")
				if httpLanguage == "" {
					httpLanguage = "zh-CN"
				}
				r.SetCtx(gi18n.WithLanguage(r.Context(), httpLanguage))
				r.Middleware.Next()
			},
		)

		group.Bind(telemetry.NewV1())
		group.Bind(device.NewV1())
		group.Bind(driver.NewV1())
		group.Bind(network.NewV1())
		group.Bind(system.NewV1())
		group.Bind(protocol.NewV1())
		group.Bind(control.NewV1())
		group.Bind(log.NewV1())

	})

	// Custom enhance API document.
	enhanceOpenAPIDoc(s)

	// 静态站点：将 `application/manifest/web` 打包进可执行文件并作为根路径提供
	if webfs, err := manifest.WebFS(); err != nil {
		g.Log().Warningf(ctx, "Web 静态资源初始化失败: %v", err)
	} else {
		fileServer := http.FileServer(http.FS(webfs))

		// 静态资源路径 - 优先处理静态资源
		s.BindHandler("GET:/assets/*", func(r *ghttp.Request) {
			fileServer.ServeHTTP(r.Response.Writer, r.Request)
		})
		s.BindHandler("GET:/images/*", func(r *ghttp.Request) {
			fileServer.ServeHTTP(r.Response.Writer, r.Request)
		})
		s.BindHandler("GET:/favicon.ico", func(r *ghttp.Request) {
			fileServer.ServeHTTP(r.Response.Writer, r.Request)
		})

		// 根路径：直接输出 index.html
		s.BindHandler("GET:/", func(r *ghttp.Request) {
			f, err := webfs.Open("index.html")
			if err != nil {
				r.Response.WriteStatus(http.StatusNotFound)
				return
			}
			defer f.Close()
			// 明确 Content-Type
			r.Response.Header().Set("Content-Type", "text/html; charset=utf-8")
			if _, err := io.Copy(r.Response.Writer, f); err != nil {
				g.Log().Warningf(ctx, "写入 index.html 失败: %v", err)
			}
		})

		// 支持Vue history模式：所有其他GET请求都返回index.html
		s.BindHandler("GET:/*", func(r *ghttp.Request) {
			// 跳过API路由
			if strings.HasPrefix(r.URL.Path, "/api/") {
				r.Response.WriteStatus(http.StatusNotFound)
				return
			}
			// 跳过WebSocket路由
			if strings.HasPrefix(r.URL.Path, "/station") || strings.HasPrefix(r.URL.Path, "/telemetry") {
				r.Response.WriteStatus(http.StatusNotFound)
				return
			}

			f, err := webfs.Open("index.html")
			if err != nil {
				r.Response.WriteStatus(http.StatusNotFound)
				return
			}
			defer f.Close()
			// 明确 Content-Type
			r.Response.Header().Set("Content-Type", "text/html; charset=utf-8")
			if _, err := io.Copy(r.Response.Writer, f); err != nil {
				g.Log().Warningf(ctx, "写入 index.html 失败: %v", err)
			}
		})
	}
	// Just run the server.
	return s
}

func MiddlewareErrorHandler(r *ghttp.Request) {
	r.Middleware.Next()
	if err := r.GetError(); err != nil {
		// 更详细的错误日志：包含请求方法/路径、请求体、堆栈
		ctx := r.Context()
		logx := g.Log().Clone()
		logx.SetStack(false)
		if stack := gerror.Stack(err); stack != "" {
			logx.Errorf(ctx, "HTTP %s %s - Error: %v\nStack:\n%s", r.Method, r.URL.Path, err, stack)
		} else {
			logx.Errorf(ctx, "HTTP %s %s - Error: %v", r.Method, r.URL.Path, err)
		}
		if body := r.GetBodyString(); body != "" {
			// 仅在调试或出现错误时打印请求体（注意敏感字段）
			logx.Debugf(ctx, "Request Body: %s", body)
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
	if strings.HasPrefix(path, "/assets/") || strings.HasPrefix(path, "/images/") || path == "/favicon.ico" {
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
	logx := g.Log().Clone()
	logx.SetStack(false)
	if status >= 400 {
		logx.Warningf(ctx, msg, method, path, status, used.Milliseconds(), clientIP, requestID)
		return
	}
	logx.Infof(ctx, msg, method, path, status, used.Milliseconds(), clientIP, requestID)
}

func enhanceOpenAPIDoc(s *ghttp.Server) {
	openapi := s.GetOpenApi()
	openapi.Config.CommonResponse = ghttp.DefaultHandlerResponse{}
	openapi.Config.CommonResponseDataField = `Data`

	// API description.
	openapi.Info = goai.Info{
		Title:       consts.OpenAPITitle,
		Description: consts.OpenAPIDescription,
		Contact: &goai.Contact{
			Name: "Hex-Ems",
			URL:  "",
		},
	}
}
