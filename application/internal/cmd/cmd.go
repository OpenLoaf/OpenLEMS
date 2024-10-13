package cmd

import (
	"common"
	"common/c_base"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/i18n/gi18n"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gproc"
	"os"
	"runtime"
	"services"
	"time"
)

const (
	argDeviceConfigName = "device-name" // 驱动配置文件
	argDriverConfigName = "driver-path" // 驱动文件的存放路径
	argEnableWeb        = "web"         // 是否启动web
	argLanguage         = "language"    // 全局语言设置
	argTimeZone         = "time-zone"   // 全局时区设置
)

var (
	Main = gcmd.Command{
		Name:  "Start",
		Usage: "start",
		Arguments: []gcmd.Argument{
			{Name: argDeviceConfigName, Short: "d", Brief: "Default: device 设备配置文件 ", IsArg: false, Orphan: false},
			{Name: argDriverConfigName, Short: "dr", Brief: "Default: ./driver 驱动存放路径 ", IsArg: false, Orphan: false},
			{Name: argEnableWeb, Short: "w", Brief: "Default: false 启动web端 ", IsArg: false, Orphan: false},
			{Name: argLanguage, Short: "l", Brief: "Default: zh-CN 设置语言 ", IsArg: false, Orphan: false},
			{Name: argTimeZone, Short: "t", Brief: "Default: zh-CN 设置语言 ", IsArg: false, Orphan: false},
		},
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			pwd, err := os.Getwd()
			if err != nil {
				return err
			}

			// 设置默认语言为中文(简体)
			gi18n.SetLanguage("zh-CN")
			deviceConfigName := parser.GetOpt(argDeviceConfigName, "device").String()
			driverConfigName := parser.GetOpt(argDriverConfigName, gfile.Join(pwd, "driver")).String()
			common.SystemInitConfigInstance(deviceConfigName, driverConfigName)

			// 初始化context
			ctx, cancelFunc := context.WithCancel(context.Background())
			ctx = context.WithValue(ctx, c_base.ConstCtxKeyGroupName, "Main")

			g.Log().Infof(ctx, "程序启动！PID：%d", os.Getpid())
			g.Log().Infof(ctx, "加载驱动文件路径：%s", driverConfigName)

			// 启动设备
			deviceCmd := services.NewDeviceCmd(ctx)
			deviceCmd.Start()

			var web *ghttp.Server

			gproc.AddSigHandlerShutdown(func(sig os.Signal) {
				g.Log().Noticef(ctx, "接收到信号：%s", sig.String())
				if web != nil {
					_ = web.Shutdown()
				}
				if deviceCmd != nil {
					deviceCmd.Stop()
				}
				cancelFunc()
				time.Sleep(1 * time.Second)
				g.Log().Infof(ctx, "程序退出！剩余Goroutine数量：%d", runtime.NumGoroutine())
			})

			if parser.GetOpt(argEnableWeb).Bool() {
				g.Log().Infof(ctx, "启动web服务！")
				web = startWeb(ctx)
				web.Run()
			} else {
				g.Log().Infof(ctx, "未启动web服务！")
				gproc.Listen()
			}

			return nil
		},
	}
)
