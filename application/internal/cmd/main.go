package cmd

import (
	"application/internal/utils"
	"common/c_enum"
	"common/c_log"
	"context"
	"os"
	"runtime"
	"time"

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
	ArgTest               = "test"                  // 测试模式：启动3秒后自动关闭
	ArgForce              = "force"                 // 强制启动：忽略PID检查
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
			{Name: ArgSqliteDbPath, Short: "cp", Brief: "Default: ./out/data 设置配置数据库路径 ", IsArg: false, Orphan: false},
			{Name: ArgLanguage, Short: "l", Brief: "Default: zh-CN 设置语言 ", IsArg: false, Orphan: false},
			{Name: ArgActiveDeviceRootId, Brief: "强制激活根设备ID ", IsArg: false, Orphan: false},
			{Name: ArgProfile, Brief: "Default: prod 选择配置profile (dev/prod等)", IsArg: false, Orphan: false},
			{Name: ArgTest, Short: "t", Brief: "Default: false 测试模式：启动3秒后自动关闭", IsArg: false, Orphan: false},
			{Name: ArgForce, Brief: "Default: false 强制启动：忽略PID检查，允许重复启动", IsArg: false, Orphan: false},
		},
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			// 初始化context
			ctx = context.WithValue(ctx, c_enum.ELogTypeEms, "Main")
			MainCtx = ctx
			ctx, cancelFunc := context.WithCancel(context.Background())
			defer cancelFunc() // 确保在所有退出路径中调用 cancelFunc

			// 检查是否强制启动
			forceStart := parser.GetOpt(ArgForce).Bool()

			// 写入PID文件（带检查）
			// 注意：必须在 InitSystem 之前执行，避免 TSDB 数据库锁冲突
			pid := os.Getpid()
			pidFile := utils.GetPidFilePath(ctx)
			if err := utils.WritePidFileWithCheck(pidFile, pid, forceStart); err != nil {
				if forceStart {
					g.Log().Warningf(ctx, "强制启动：%v", err)
				} else {
					g.Log().Errorf(ctx, "启动失败：%v", err)
					return err
				}
			} else {
				g.Log().Infof(ctx, "PID已保存到文件: %s", pidFile)
			}

			// 初始化系统（在 PID 检查之后执行，确保旧进程已被终止）
			if err := InitSystem(ctx, parser); err != nil {
				panic(err)
			}

			// 启动服务
			StartServices(ctx)

			// 设置关闭信号处理
			SetupShutdownHandler(ctx, cancelFunc)

			// 检查是否为测试模式
			enableTest := parser.GetOpt(ArgTest).Bool()
			if enableTest {
				c_log.Infof(ctx, "===> 测试模式已启用，程序将在5秒后自动关闭")
				go func() {
					time.Sleep(3 * time.Second)
					// 倒计时5秒，每秒显示剩余时间
					for i := 5; i > 0; i-- {
						time.Sleep(1 * time.Second)
						if i > 1 {
							c_log.Infof(ctx, "===> 测试模式倒计时：%d秒后自动关闭", i-1)
						} else {
							c_log.Infof(ctx, "===> 测试模式：即将结束进程")
						}
					}
					c_log.Infof(ctx, "===> 测试模式：发送shutdown信号")
					// 使用跨平台的进程终止函数
					if err := utils.KillProcess(); err != nil {
						c_log.Errorf(ctx, "发送终止信号失败: %v", err)
					}
				}()
			}

			// 启动Web服务或等待信号
			enableWeb := parser.GetOpt(ArgEnableWeb).Bool()
			enableGui := parser.GetOpt(ArgEnableGui).Bool()
			enableGuiFull := parser.GetOpt(ArgEnableGuiFull).Bool()

			// Windows 环境下默认启动 GUI
			if runtime.GOOS == "windows" && !enableWeb && !enableGui && !enableGuiFull {
				c_log.Infof(ctx, "Windows 环境：默认启动 GUI 界面")
				enableGui = true
			}

			if enableWeb || enableGui || enableGuiFull {
				// 如果启用了Web或GUI，都需要启动Web服务
				if (enableGui || enableGuiFull) && !enableWeb {
					// 如果只启用GUI，启动本地绑定的Web服务
					c_log.Infof(ctx, "启动web服务（GUI模式，仅本地访问）！")
					web := startWebWithBinding(ctx, true)
					go web.Run()

					// 启动GUI界面
					if enableGuiFull {
						c_log.Infof(ctx, "启动GUI界面（全屏模式）！")
						startGuiFullscreen(ctx)
					} else {
						c_log.Infof(ctx, "启动GUI界面！")
						startGui(ctx)
					}
				} else {
					// 如果启用了Web（无论是否启用GUI），启动正常的Web服务
					c_log.Infof(ctx, "启动web服务！")
					web := startWeb(ctx)

					// 如果同时启用了GUI，则启动GUI界面
					if enableGui {
						c_log.Infof(ctx, "启动GUI界面！")
						go startGui(ctx)
					} else if enableGuiFull {
						c_log.Infof(ctx, "启动GUI界面（全屏模式）！")
						go startGuiFullscreen(ctx)
					}

					web.Run()
				}
			} else {
				c_log.Infof(ctx, "未启动web服务和GUI界面！")
				gproc.Listen()
			}

			return nil
		},
	}
)
