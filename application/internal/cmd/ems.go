package cmd

import (
	applog "application/internal/log"
	"application/internal/logic"
	"application/internal/service"
	"application/internal/utils"
	_ "application/manifest"
	"common"
	"common/c_base"
	"common/c_enum"
	"common/c_log"
	"context"
	"os"
	"p_tsdb"
	"runtime"
	"s_automation"
	"s_db"
	"s_db/s_db_basic"
	"s_driver"
	"s_storage"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/i18n/gi18n"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gproc"
	"github.com/shockerli/cvt"
)

// InitSystem 初始化系统
func InitSystem(ctx context.Context, parser *gcmd.Parser) error {

	// 注入系统日志（GoFrame）- 在配置加载后创建，确保使用正确的日志级别
	c_log.SetSystemLogger(applog.NewSystemAdapter(g.Log()))
	// 注入业务日志（输出到数据库）
	c_log.SetBusinessLogger(applog.NewDatabaseAdapter())
	// 启用异步日志输出提高性能
	g.Log().SetAsync(true)

	// 设置默认语言为中文(简体)
	gi18n.SetLanguage("zh-CN")

	// 初始化数据库
	s_db.Init()

	// 初始化存储（切换为 TSDB，无外部配置时采用默认路径与策略）
	storageInst := p_tsdb.NewStorageInstance(ctx, &c_base.SStorageConfig{Enable: true, Type: c_enum.EStorageTypeTsdb, Url: "", Params: map[string]string{}})
	s_storage.NewSingleStorageManager(ctx, storageInst)
	common.RegisterStorageInstance(storageInst)

	// 注册设备管理器
	common.RegisterDeviceManager(s_driver.NewDriverManagerImpl(ctx))

	// 注册告警管理器（直接注入 s_db 的告警服务实现，满足 common 的告警接口）
	common.RegisterAlarmManager(s_db.GetAlarmService())

	// 初始化自动化服务
	s_automation.Init()

	// 注册自动化服务
	service.RegisterAutomation(logic.NewAutomation())

	return nil
}

// StartServices 启动所有服务
func StartServices(ctx context.Context) {
	pid := os.Getpid()
	g.Log().Infof(ctx, "EMS After！PID：%d", pid)

	// 启动设备
	go func() {
		common.GetDeviceManager().Start()

		c_log.BizInfof(ctx, "EMS系统启动成功！")
		g.Log().Infof(ctx, "DeviceManger State : %s", common.GetDeviceManager().Status())
	}()

	// 启动自动化管理器
	go func() {
		// 等待设备管理器启动完成
		time.Sleep(2 * time.Second)

		internalMillisecondsStr := s_db.GetSettingService().GetSettingValueBySystemSettingDefine(ctx,
			s_db_basic.SystemSettingAutomationInternalMilliseconds)

		internalMilliseconds := cvt.Int64(internalMillisecondsStr)
		if internalMilliseconds < 0 {
			internalMilliseconds = cvt.Int64(s_db_basic.SystemSettingAutomationInternalMilliseconds.DefaultValue)
		}

		// 启动自动化管理器，按配置间隔执行
		err := s_automation.StartAutomationManager(ctx, time.Duration(internalMilliseconds)*time.Millisecond)
		if err != nil {
			g.Log().Errorf(ctx, "启动自动化管理器失败: %+v", err)
		} else {
			c_log.BizInfof(ctx, "自动化服务启动成功！")
		}
	}()
}

// SetupShutdownHandler 设置关闭信号处理
func SetupShutdownHandler(ctx context.Context, cancelFunc context.CancelFunc) {
	gproc.AddSigHandlerShutdown(func(sig os.Signal) {
		g.Log().Infof(ctx, "接收到关闭服务信号：%s", sig.String())

		// 停止自动化管理器
		err := s_automation.StopAutomationManager(ctx)
		if err != nil {
			g.Log().Errorf(ctx, "停止自动化管理器失败: %+v", err)
		} else {
			g.Log().Infof(ctx, "自动化管理器已停止")
			c_log.BizInfof(ctx, "自动化服务已停止")
		}

		common.GetDeviceManager().Shutdown()

		// 清理PID文件
		pidFile := utils.GetPidFilePath(ctx)
		if err := utils.RemovePidFile(pidFile); err != nil {
			g.Log().Warningf(ctx, "清理PID文件失败: %v", err)
		} else {
			g.Log().Infof(ctx, "PID文件已清理: %s", pidFile)
		}

		cancelFunc()
		time.Sleep(1 * time.Second)
		g.Log().Infof(ctx, "程序退出！剩余Goroutine数量：%d", runtime.NumGoroutine())
		c_log.BizWarningf(ctx, "EMS系统关闭！")
	})
}
