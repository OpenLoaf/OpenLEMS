package cmd

import (
	"application/internal/consts"
	"application/internal/controller/device"
	"application/internal/controller/driver"
	"application/internal/controller/network"
	"application/internal/controller/protocol"
	"application/internal/controller/station_ess"
	"application/internal/controller/system"
	"application/internal/controller/telemetry"
	"application/internal/ws"
	"application/manifest"
	"common/c_base"
	"context"
	"io"
	"net/http"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/i18n/gi18n"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/net/goai"
)

func startWeb(ctx context.Context) *ghttp.Server {
	ctx = context.WithValue(ctx, c_base.ConstCtxKeyGroupName, "Web")
	// todo 添加关闭服务，允许使用接口关闭web

	g.Log().Infof(ctx, "准备启动web程序！")

	s := g.Server()
	s.Use(ghttp.MiddlewareHandlerResponse, ghttp.MiddlewareNeverDoneCtx, MiddlewareErrorHandler)

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
		group.Bind(station_ess.NewV1())
		group.Bind(device.NewV1())
		group.Bind(device.NewV2())
		group.Bind(driver.NewV1())
		group.Bind(network.NewV1())
		group.Bind(system.NewV1())
		group.Bind(protocol.NewV1())
	})

	s.BindObject("/station", ws.NewStationWebsocket())
	s.BindObject("/telemetry", ws.NewTelemetryWebsocket())

	// Custom enhance API document.
	enhanceOpenAPIDoc(s)

	// 静态站点：将 `application/manifest/web` 打包进可执行文件并作为根路径提供
	if webfs, err := manifest.WebFS(); err != nil {
		g.Log().Warningf(ctx, "Web 静态资源初始化失败: %v", err)
	} else {
		fileServer := http.FileServer(http.FS(webfs))
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
		// 静态资源路径
		s.BindHandler("GET:/assets/*", func(r *ghttp.Request) {
			fileServer.ServeHTTP(r.Response.Writer, r.Request)
		})
		s.BindHandler("GET:/images/*", func(r *ghttp.Request) {
			fileServer.ServeHTTP(r.Response.Writer, r.Request)
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
		g.Log().Errorf(ctx, "HTTP %s %s - Error: %v", r.Method, r.URL.Path, err)
		if stack := gerror.Stack(err); stack != "" {
			g.Log().Errorf(ctx, "Stack:\n%s", stack)
		}
		if body := r.GetBodyString(); body != "" {
			// 仅在调试或出现错误时打印请求体（注意敏感字段）
			g.Log().Debugf(ctx, "Request Body: %s", body)
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
