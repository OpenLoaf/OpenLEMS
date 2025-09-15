package cmd

import (
	"application/internal/utils"
	"common/c_base"
	"context"
	"os"
	"runtime"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gproc"
)

const (
	ArgDeviceConfigName   = "device-name"           // 驱动配置文件
	ArgDriverConfigName   = "driver-path"           // 驱动文件的存放路径
	ArgEnableWeb          = "web"                   // 是否启动web
	ArgEnableGui          = "gui"                   // 是否启动GUI
	ArgEnableGuiFull      = "gui-full"              // 是否启动GUI并全屏
	ArgLanguage           = "language"              // 全局语言设置
	ArgPebbleDbPath       = "runtime-path"          // pebble数据库路径
	ArgSqliteDbPath       = "db-path"               // sqlite数据库路径
	ArgActiveDeviceRootId = "active-device-root-id" // 强制激活根设备
	ArgProfile            = "profile"               // 配置profile: default/dev/prod等
	DefaultPidFile        = "out/ems.pid"           // 默认PID文件路径
)

var MainCtx context.Context

var (
	Main = gcmd.Command{
		Name:  "After",
		Usage: "start",
		Arguments: []gcmd.Argument{
			{Name: ArgEnableWeb, Short: "w", Brief: "Default: false 启动web端 ", IsArg: false, Orphan: false},
			{Name: ArgEnableGui, Short: "g", Brief: "Default: false 启动GUI界面 ", IsArg: false, Orphan: false},
			{Name: ArgEnableGuiFull, Short: "f", Brief: "Default: false 启动GUI界面并全屏 ", IsArg: false, Orphan: false},
			{Name: ArgDeviceConfigName, Short: "d", Brief: "Default: device 设备配置文件 ", IsArg: false, Orphan: false},
			{Name: ArgDriverConfigName, Short: "dp", Brief: "Default: ./driver 驱动存放路径 ", IsArg: false, Orphan: false},
			{Name: ArgPebbleDbPath, Short: "rp", Brief: "Default: ./out/runtime 设置实时数据库路径 ", IsArg: false, Orphan: false},
			{Name: ArgSqliteDbPath, Short: "cp", Brief: "Default: ./out/db.sqlite3 设置配置数据库路径 ", IsArg: false, Orphan: false},
			{Name: ArgLanguage, Short: "l", Brief: "Default: zh-CN 设置语言 ", IsArg: false, Orphan: false},
			{Name: ArgActiveDeviceRootId, Brief: "强制激活根设备ID ", IsArg: false, Orphan: false},
			{Name: ArgProfile, Short: "p", Brief: "Default: default 选择配置profile (default/dev/prod等)", IsArg: false, Orphan: false},
		},
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			// 初始化context
			ctx = context.WithValue(ctx, c_base.ConstCtxKeyGroupName, "Main")
			MainCtx = ctx
			ctx, cancelFunc := context.WithCancel(context.Background())

			// 初始化系统
			if err := InitSystem(ctx, parser); err != nil {
				panic(err)
			}

			// 写入PID文件
			pid := os.Getpid()
			pidFile := DefaultPidFile
			if err := utils.WritePidFile(pidFile, pid); err != nil {
				g.Log().Warningf(ctx, "写入PID文件失败: %v", err)
			} else {
				g.Log().Infof(ctx, "PID已保存到文件: %s", pidFile)
			}

			// 启动服务
			StartServices(ctx)

			// 设置关闭信号处理
			SetupShutdownHandler(ctx, cancelFunc)

			// 启动Web服务或等待信号
			enableWeb := parser.GetOpt(ArgEnableWeb).Bool()
			enableGui := parser.GetOpt(ArgEnableGui).Bool()
			enableGuiFull := parser.GetOpt(ArgEnableGuiFull).Bool()

			// Windows 环境下默认启动 GUI
			if runtime.GOOS == "windows" && !enableWeb && !enableGui && !enableGuiFull {
				g.Log().Infof(ctx, "Windows 环境：默认启动 GUI 界面")
				enableGui = true
			}

			if enableWeb || enableGui || enableGuiFull {
				// 如果启用了Web或GUI，都需要启动Web服务
				if (enableGui || enableGuiFull) && !enableWeb {
					// 如果只启用GUI，启动本地绑定的Web服务
					g.Log().Infof(ctx, "启动web服务（GUI模式，仅本地访问）！")
					web := startWebWithBinding(ctx, true)
					go web.Run()

					// 启动GUI界面
					if enableGuiFull {
						g.Log().Infof(ctx, "启动GUI界面（全屏模式）！")
						startGuiFullscreen(ctx)
					} else {
						g.Log().Infof(ctx, "启动GUI界面！")
						startGui(ctx)
					}
				} else {
					// 如果启用了Web（无论是否启用GUI），启动正常的Web服务
					g.Log().Infof(ctx, "启动web服务！")
					web := startWeb(ctx)

					// 如果同时启用了GUI，则启动GUI界面
					if enableGui {
						g.Log().Infof(ctx, "启动GUI界面！")
						go startGui(ctx)
					} else if enableGuiFull {
						g.Log().Infof(ctx, "启动GUI界面（全屏模式）！")
						go startGuiFullscreen(ctx)
					}

					web.Run()
				}
			} else {
				g.Log().Infof(ctx, "未启动web服务和GUI界面！")
				gproc.Listen()
			}

			return nil
		},
	}
)
