package cmd

import (
	applog "application/internal/log"
	"application/manifest"
	"common"
	"common/c_base"
	"common/c_log"
	"context"
	"os"
	"runtime"
	"s_db"
	"s_driver"
	"s_storage"
	"time"
	"tsdb"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/i18n/gi18n"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gproc"
)

const (
	ArgDeviceConfigName   = "device-name"           // 驱动配置文件
	ArgDriverConfigName   = "driver-path"           // 驱动文件的存放路径
	ArgEnableWeb          = "web"                   // 是否启动web
	ArgLanguage           = "language"              // 全局语言设置
	ArgPebbleDbPath       = "runtime-path"          // pebble数据库路径
	ArgSqliteDbPath       = "db-path"               // sqlite数据库路径
	ArgActiveDeviceRootId = "active-device-root-id" // 强制激活根设备
	ArgProfile            = "profile"               // 配置profile: default/dev/prod等
)

var MainCtx context.Context

var (
	Main = gcmd.Command{
		Name:  "Start",
		Usage: "start",
		Arguments: []gcmd.Argument{
			{Name: ArgEnableWeb, Short: "w", Brief: "Default: false 启动web端 ", IsArg: false, Orphan: false},
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

			// 注入系统日志（GoFrame）
			c_log.SetSystemLogger(applog.NewGoFrameLoggerAdapter(g.Log()))
			// 注入业务日志（基于上下文路由到不同类型与ID文件）
			c_log.SetBusinessLogger(applog.NewBizRouterLoggerAdapter())
			// 启用异步日志输出提高性能
			g.Log().SetAsync(true)

			// 优先加载嵌入式配置，支持 --profile 或 APP_PROFILE 环境变量
			profile := parser.GetOpt(ArgProfile, os.Getenv("APP_PROFILE")).String()
			if profile == "" || profile == "default" {
				profile = "prod"
			}
			g.Log().Infof(ctx, "Active Profile: %s", profile)
			manifest.LoadEmbeddedConfig(profile)

			// 设置默认语言为中文(简体)
			gi18n.SetLanguage("zh-CN")

			// 初始化数据库
			s_db.Init()

			// 初始化存储（切换为 TSDB，无外部配置时采用默认路径与策略）
			storageInst := tsdb.NewStorageInstance(ctx, &c_base.SStorageConfig{Enable: true, Type: c_base.EStorageTypePebbledb, Url: "", Params: map[string]string{}})
			s_storage.NewSingleStorageManager(nil, storageInst)
			common.RegisterStorageInstance(storageInst)

			common.RegisterDeviceManager(s_driver.NewDriverManagerImpl(ctx))
			if err != nil {
				panic(err)
			}

			g.Log().Infof(ctx, "EMS Start！PID：%d", os.Getpid())

			// 启动设备
			go func() {
				common.GetDeviceManager().Start()
				g.Log().Infof(ctx, "DeviceManger State : %s", common.GetDeviceManager().Status())
			}()

			gproc.AddSigHandlerShutdown(func(sig os.Signal) {
				g.Log().Infof(ctx, "接收到关闭服务信号：%s", sig.String())
				cancelFunc()
				time.Sleep(1 * time.Second)
				g.Log().Infof(ctx, "程序退出！剩余Goroutine数量：%d", runtime.NumGoroutine())
			})

			if parser.GetOpt(ArgEnableWeb).Bool() {
				g.Log().Infof(ctx, "启动web服务！")
				var web *ghttp.Server
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
