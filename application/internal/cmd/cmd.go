package cmd

import (
	"application/internal/collect"
	"application/internal/consts"
	"application/internal/controller/device"
	"application/internal/controller/telemetry"
	"application/internal/ws"
	"context"
	"ems-plan/c_base"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/i18n/gi18n"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/net/goai"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gproc"
	"os"
	"runtime"
	"time"
)

var (
	Main = gcmd.Command{
		Name:  "Ems",
		Usage: "Ems",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			// 先获取所有配置文件
			ctx, cancelFunc := context.WithCancel(context.Background())

			ctx = context.WithValue(ctx, "I18nName", "Main")

			pid := os.Getpid()
			g.Log().Infof(ctx, "程序启动！PID：%d", pid)

			linkConfigList, configPath, err := c_base.GetConfigList[c_base.SProtocolConfig](ctx, c_base.DevicesKey)
			if err != nil {
				g.Log().Error(ctx, err)
				cancelFunc()
				return err
			}

			g.Log().Infof(ctx, "查询到配置文件：%s", configPath)
			// 从配置文件中获取link的信息用来创建link
			err = collect.Create(ctx, linkConfigList)
			if err != nil {
				g.Log().Errorf(ctx, "启动采集程序失败！%+v", err)
				exit(ctx, cancelFunc)
				return
			}
			// 设置默认语言为中文(简体)
			gi18n.SetLanguage("zh-CN")
			s := g.Server()
			s.Use(ghttp.MiddlewareHandlerResponse, ghttp.MiddlewareNeverDoneCtx, MiddlewareErrorHandler)

			s.Group("/", func(group *ghttp.RouterGroup) {
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
				//group.Bind(station_ess.NewV1())
				group.Bind(device.NewV1())
			})

			s.BindObject("/telemetry", ws.NewTelemetryWebsocket())

			gproc.AddSigHandlerShutdown(func(sig os.Signal) {
				exit(ctx, cancelFunc)
				_ = s.Shutdown()
			})

			// Custom enhance API document.
			enhanceOpenAPIDoc(s)
			// Just run the server.
			s.Run()
			return nil
		},
	}
)

func MiddlewareErrorHandler(r *ghttp.Request) {
	r.Middleware.Next()
	if err := r.GetError(); err != nil {
		// 记录到自定义错误日志文件
		g.Log("exception").Error(r.Context(), err)
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

func exit(ctx context.Context, cancel context.CancelFunc) {
	cancel()
	time.Sleep(1 * time.Second)

	g.Log().Infof(ctx, "程序退出！剩余Goroutine数量：%d", runtime.NumGoroutine())
}
