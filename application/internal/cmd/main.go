package cmd

import (
	"application/internal/utils"
	"common/c_base"
	"context"
	"os"

	"github.com/gogf/gf/v2/frame/g"
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
	DefaultPidFile        = "out/ems.pid"           // 默认PID文件路径
)

var MainCtx context.Context

var (
	Main = gcmd.Command{
		Name:  "After",
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
			if parser.GetOpt(ArgEnableWeb).Bool() {
				g.Log().Infof(ctx, "启动web服务！")
				web := startWeb(ctx)
				web.Run()
			} else {
				g.Log().Infof(ctx, "未启动web服务！")
				gproc.Listen()
			}

			return nil
		},
	}
)
