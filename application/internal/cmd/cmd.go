package cmd

import (
	"common"
	"common/c_base"
	"context"
	"os"
	"pebbledb"
	"runtime"
	"services"
	database "sqlite"
	"time"

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
)

var DeviceStartCancel context.CancelFunc = nil // 设备启动取消函数，设备启动时候回堵塞进程

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
		},
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {

			// 设置默认语言为中文(简体)
			gi18n.SetLanguage("zh-CN")
			// deviceConfigName := parser.GetOpt(ArgDeviceConfigName, "device").String()
			// driverConfigName := parser.GetOpt(ArgDriverConfigName, gfile.Join(pwd, "drivers")).String()

			// 初始化数据库表
			database.Init()

			common.SystemInitConfigInstance(ctx)

			// 初始化存储
			common.RegisterStorageInstance(func(ctx context.Context) c_base.IStorage {
				return pebbledb.NewStorageInstance(ctx)
			})

			// 初始化context
			ctx, cancelFunc := context.WithCancel(context.Background())
			ctx = context.WithValue(ctx, c_base.ConstCtxKeyGroupName, "Main")

			g.Log().Infof(ctx, "Hex EMS程序启动！PID：%d", os.Getpid())
			// g.Log().Infof(ctx, "加载驱动文件路径：%s", driverConfigName)

			// 启动设备
			deviceCmd := services.NewDeviceCmd(ctx)

			deviceStartCtx, DeviceStartCancel := context.WithCancel(context.Background())
			go deviceStart(deviceStartCtx, deviceCmd, parser.GetOpt(ArgActiveDeviceRootId).String())

			var web *ghttp.Server

			gproc.AddSigHandlerShutdown(func(sig os.Signal) {
				g.Log().Infof(ctx, "接收到关闭服务信号：%s", sig.String())
				if web != nil {
					_ = web.Shutdown()
				}
				if DeviceStartCancel != nil {
					DeviceStartCancel()
				}
				if deviceCmd != nil {
					deviceCmd.Stop()
				}
				cancelFunc()
				// 关闭存储
				common.CloseStorage()

				time.Sleep(1 * time.Second)
				g.Log().Infof(ctx, "程序退出！剩余Goroutine数量：%d", runtime.NumGoroutine())
			})

			if parser.GetOpt(ArgEnableWeb).Bool() {
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

func deviceStart(deviceStartCtx context.Context, deviceCmd c_base.IService, activeDeviceRootId string) {
	select {
	case <-deviceStartCtx.Done():
		deviceCmd.Stop()
		return
	default:
		deviceCmd.Start(activeDeviceRootId)
	}
}
