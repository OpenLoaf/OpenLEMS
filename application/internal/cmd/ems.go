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
	"p_policy_energy_storage"
	"p_policy_mircogrid"
	"runtime"
	"s_automation"
	"s_db"
	"s_db/s_db_basic"
	"s_driver"
	s_export_modbus "s_export_modbus"
	s_export_mqtt "s_mqtt"
	"s_policy"
	"s_price"
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
	if err := s_db.Init(); err != nil {
		c_log.Errorf(ctx, "数据库初始化失败: %+v", err)
		return err
	}

	// 初始化存储管理器
	storageManager := s_storage.NewStorageManagerWithConfig(ctx, &c_base.SStorageConfig{
		Enable: true,
		Type:   c_enum.EStorageTypeTsdb,
		Url:    "",
		Params: map[string]string{},
	})
	common.RegisterStorageInstance(storageManager.GetStorageInstance())

	// 注册设备管理器
	common.RegisterDeviceManager(s_driver.NewDriverManagerImpl(ctx))

	// 注册告警管理器（直接注入 s_db 的告警服务实现，满足 common 的告警接口）
	common.RegisterAlarmManager(s_db.GetAlarmService())

	// 初始化自动化服务
	s_automation.Init()

	// 注册自动化服务
	service.RegisterAutomation(logic.NewAutomation())

	// 创建并注册策略管理器
	policyManager := s_policy.NewPolicyManager(ctx)
	common.RegisterPolicyManager(policyManager)

	// 创建并注册电价管理器
	priceManager := s_price.NewPriceManager(ctx)
	common.RegisterPriceManager(priceManager)

	// 注册策略插件（开发环境直接注册，生产环境通过插件加载）
	err := policyManager.RegisterPolicy("policy_microgrid", p_policy_mircogrid.NewPolicyMircogrid())
	if err != nil {
		c_log.Errorf(ctx, "注册微电网策略失败: %+v", err)
	}
	err = policyManager.RegisterPolicy("policy_ess", p_policy_energy_storage.NewPolicyEnergyStorage())
	if err != nil {
		c_log.Errorf(ctx, "注册储能站策略失败: %+v", err)
	}

	// 初始化MQTT服务（在其他服务初始化之后）
	s_export_mqtt.Init()

	// 初始化Modbus服务
	s_export_modbus.Init()

	return nil
}

// StartServices 启动所有服务
func StartServices(ctx context.Context) {
	pid := os.Getpid()
	c_log.Infof(ctx, "EMS After！PID：%d", pid)

	// 启动存储服务（系统指标保存）
	go func() {
		if err := s_storage.StartStorageManager(ctx); err != nil {
			c_log.Errorf(ctx, "启动存储服务失败: %+v", err)
		} else {
			c_log.BizInfof(ctx, "存储服务启动成功！")
		}
	}()

	// 启动设备
	go func() {
		common.GetDeviceManager().Start()

		c_log.BizInfof(ctx, "EMS系统启动成功！")
		c_log.Infof(ctx, "DeviceManger State : %s", common.GetDeviceManager().Status())
	}()

	// 启动自动化管理器
	go func() {
		// 等待设备管理器启动完成
		time.Sleep(2 * time.Second)

		internalMillisecondsStr := s_db.GetSettingService().GetSettingValueBySystemSettingDefine(ctx,
			s_db_basic.SystemSettingAutomationInternalMilliseconds)

		var internalMilliseconds int64
		if internalMillisecondsStr != nil {
			internalMilliseconds = cvt.Int64(*internalMillisecondsStr)
		}
		if internalMilliseconds < 0 {
			internalMilliseconds = cvt.Int64(s_db_basic.SystemSettingAutomationInternalMilliseconds.DefaultValue)
		}

		// 启动自动化管理器，按配置间隔执行
		err := s_automation.StartAutomationManager(ctx, time.Duration(internalMilliseconds)*time.Millisecond)
		if err != nil {
			c_log.Errorf(ctx, "启动自动化管理器失败: %+v", err)
		} else {
			c_log.BizInfof(ctx, "自动化服务启动成功！")
		}
	}()

	// 启动MQTT服务
	go func() {
		// 等待设备管理器启动完成
		time.Sleep(3 * time.Second)

		err := s_export_mqtt.StartMqtt(ctx)
		if err != nil {
			c_log.Errorf(ctx, "启动MQTT服务失败: %+v", err)
		} else {
			c_log.BizInfof(ctx, "MQTT服务启动成功！")
		}
	}()

	// 启动Modbus服务
	go func() {
		// 等待设备管理器启动完成
		time.Sleep(4 * time.Second)

		err := s_export_modbus.StartModbus(ctx)
		if err != nil {
			c_log.Errorf(ctx, "启动Modbus服务失败: %+v", err)
		} else {
			c_log.BizInfof(ctx, "Modbus服务启动成功！")
		}
	}()

	// 启动策略管理器
	go func() {
		// 等待设备管理器启动完成
		time.Sleep(5 * time.Second)

		err := common.GetPolicyManager().Start(ctx)
		if err != nil {
			c_log.Errorf(ctx, "启动策略管理器失败: %+v", err)
		} else {
			activePolicyId := common.GetPolicyManager().GetActivePolicyId()
			if activePolicyId != "" {
				c_log.BizInfof(ctx, "策略管理器启动成功！当前激活策略: %s", activePolicyId)
			} else {
				c_log.BizInfof(ctx, "策略管理器启动成功！未配置激活策略")
			}
		}
	}()

	// 启动电价管理器
	go func() {
		// 等待设备管理器启动完成
		time.Sleep(6 * time.Second)

		err := common.GetPriceManager().Start(ctx)
		if err != nil {
			c_log.Errorf(ctx, "启动电价管理器失败: %+v", err)
		} else {
			c_log.BizInfof(ctx, "电价管理器启动成功！")
		}
	}()

}

// SetupShutdownHandler 设置关闭信号处理
func SetupShutdownHandler(ctx context.Context, cancelFunc context.CancelFunc) {
	gproc.AddSigHandlerShutdown(func(sig os.Signal) {
		c_log.Infof(ctx, "接收到关闭服务信号：%s", sig.String())

		// 停止自动化管理器
		err := s_automation.StopAutomationManager(ctx)
		if err != nil {
			c_log.Errorf(ctx, "停止自动化管理器失败: %+v", err)
		} else {
			c_log.Infof(ctx, "自动化管理器已停止")
			c_log.BizInfof(ctx, "自动化服务已停止")
		}

		// 停止MQTT服务
		err = s_export_mqtt.StopMqtt(ctx)
		if err != nil {
			c_log.Errorf(ctx, "停止MQTT服务失败: %+v", err)
		} else {
			c_log.Infof(ctx, "MQTT服务已停止")
			c_log.BizInfof(ctx, "MQTT服务已停止")
		}

		// 停止Modbus服务
		err = s_export_modbus.StopModbus(ctx)
		if err != nil {
			c_log.Errorf(ctx, "停止Modbus服务失败: %+v", err)
		} else {
			c_log.Infof(ctx, "Modbus服务已停止")
			c_log.BizInfof(ctx, "Modbus服务已停止")
		}

		// 停止策略管理器
		common.GetPolicyManager().Shutdown()
		c_log.Infof(ctx, "策略管理器已停止")
		c_log.BizInfof(ctx, "策略管理器已停止")

		// 停止电价管理器
		common.GetPriceManager().Shutdown()
		c_log.Infof(ctx, "电价管理器已停止")
		c_log.BizInfof(ctx, "电价管理器已停止")

		common.GetDeviceManager().Shutdown()

		// 清理PID文件
		pidFile := utils.GetPidFilePath(ctx)
		if err := utils.RemovePidFile(pidFile); err != nil {
			c_log.Warningf(ctx, "清理PID文件失败: %v", err)
		} else {
			c_log.Infof(ctx, "PID文件已清理: %s", pidFile)
		}

		cancelFunc()
		time.Sleep(1 * time.Second)
		c_log.Infof(ctx, "程序退出！剩余Goroutine数量：%d", runtime.NumGoroutine())
		c_log.BizWarningf(ctx, "EMS系统关闭！")
	})
}
