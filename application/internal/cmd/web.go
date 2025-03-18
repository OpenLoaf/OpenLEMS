package cmd

import (
	"application/internal/consts"
	"application/internal/controller/device"
	"application/internal/controller/station_ess"
	"application/internal/controller/telemetry"
	"application/internal/ws"
	"common/c_base"
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/i18n/gi18n"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/net/goai"
)

func startWeb(ctx context.Context) *ghttp.Server {
	ctx = context.WithValue(ctx, c_base.ConstCtxKeyGroupName, "Web")
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
	})

	s.BindObject("/station", ws.NewStationWebsocket())
	s.BindObject("/telemetry", ws.NewTelemetryWebsocket())

	// Custom enhance API document.
	enhanceOpenAPIDoc(s)
	// Just run the server.
	return s
}

func MiddlewareErrorHandler(r *ghttp.Request) {
	r.Middleware.Next()
	if err := r.GetError(); err != nil {
		// 记录到自定义错误日志文件
		//g.Log("exception").Error(r.Context(), err)
		//返回固定的友好信息
		r.Response.ClearBuffer()
		if gerr, ok := err.(*gerror.Error); ok {
			r.Response.WriteJson(ghttp.DefaultHandlerResponse{
				Code: gerr.Code().Code(),
				//Message: g.I18n().T(r.Context(), gerr.Code().Message()),
				Message: err.Error(),
			})
			return
		} else {
			r.Response.WriteJson(ghttp.DefaultHandlerResponse{
				Code: 500,
				//Message: g.I18n().T(r.Context(), "serverError"),
				Message: err.Error(),
			})
		}

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
