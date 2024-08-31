package internal

import (
	"context"
	"ems-plan/c_base"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/os/gtimer"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"plug_protocol_gpio_sysfs/p_gpio_sysfs"
	"sync"
	"time"
)

type SGpioSysfsProvider struct {
	once       sync.Once          // 只执行一次Init方法
	deviceId   string             // 设备名称
	deviceType c_base.EDeviceType // 设备类型
	//cache                 *gcache.Cache      // 点位缓存
	log             *glog.Logger // 日志
	printCacheValue bool         // 打印缓存值

	*c_base.SAlarmHandler            // 告警
	status                bool       // 结果
	lastUpdateTime        *time.Time // 最后更新时间

	deviceConfig *p_gpio_sysfs.SGpioSysfsDeviceConfig

	handler func(ctx context.Context, status bool) // 处理函数
}

func NewGpioSysfsProvider(ctx context.Context, clientConfig *c_base.SProtocolConfig, deviceConfig *p_gpio_sysfs.SGpioSysfsDeviceConfig) (p_gpio_sysfs.IGpioSysfsProtocol, error) {
	provider := &SGpioSysfsProvider{
		once:            sync.Once{},
		deviceId:        deviceConfig.Id,
		printCacheValue: deviceConfig.PrintCacheValue,
		deviceConfig:    deviceConfig,
		SAlarmHandler: &c_base.SAlarmHandler{
			Ctx: ctx,
		},
		log: g.Log(deviceConfig.Id),
	}

	if deviceConfig.LogLevel != "" {
		err := provider.log.SetLevelStr(deviceConfig.LogLevel)
		if err != nil {
			provider.log.Level(glog.LEVEL_INFO)
		}
	} else {
		// 设置默认的日志等级为上一级接口配置的日志等级
		err := provider.log.SetLevelStr(clientConfig.GetLogLevel())
		if err != nil {
			provider.log.Level(glog.LEVEL_INFO)
		}
	}

	return provider, nil
}

func (s *SGpioSysfsProvider) Init(deviceType c_base.EDeviceType) {
	if s.deviceConfig.Direction == p_gpio_sysfs.EGpioDirectionNone {
		panic(fmt.Errorf("direction方向未设置！"))
	}
	s.once.Do(func() {
		s.deviceType = deviceType
		// 先判断Path是否存在
		if !gfile.Exists(s.deviceConfig.Path) {
			// 判断一下export是否存在，如果存在，就export导出一下端口
			if gfile.Exists(s.deviceConfig.ExportPath) && s.deviceConfig.ExportPort != 0 {
				// 导出端口
				err := gfile.PutContents(s.deviceConfig.ExportPath, gconv.String(s.deviceConfig.ExportPort))
				if err != nil {
					panic(fmt.Errorf("导出端口失败！%v", err))
				}
			} else {
				panic(fmt.Errorf("path不存在！%s", s.deviceConfig.Path))
			}
		}
		// 设置方向
		err := gfile.PutContents(gfile.Join(s.deviceConfig.Path, GpioPathDirection), string(s.deviceConfig.Direction))
		if err != nil {
			panic(fmt.Errorf("设置方向失败！%v", err))
		}

		// 如果是In类型的，说明需要时刻监听
		if s.deviceConfig.Direction == p_gpio_sysfs.EGpioDirectionIn {
			gtimer.SetInterval(s.Ctx, 200*time.Millisecond, func(ctx context.Context) {
				// 读取值
				s.IsHigh()
			})
			if err != nil {
				panic(fmt.Errorf("监听失败！%v", err))
			}
		}
		s.log.Infof(s.Ctx, "GPIO %s init success,当前状态为: %v, 类型为: %s", s.deviceId, s.IsHigh(), s.deviceConfig.Direction)
	})
}

func (s *SGpioSysfsProvider) RegisterHandler(handler func(ctx context.Context, status bool)) {
	if s.deviceConfig.Direction != p_gpio_sysfs.EGpioDirectionIn {
		panic(fmt.Errorf("只有输入类型的GPIO才能注册Handler"))
	}
	s.handler = handler
}

func (s *SGpioSysfsProvider) GetId() string {
	return s.deviceId
}

func (s *SGpioSysfsProvider) GetType() c_base.EDeviceType {
	return s.deviceType
}

func (s *SGpioSysfsProvider) Close() error {
	return nil
}

func (s *SGpioSysfsProvider) IsActivate() bool {
	return true
}

func (s *SGpioSysfsProvider) PrintCacheValues() {

}

func (s *SGpioSysfsProvider) GetLastUpdateTime() *time.Time {
	return s.lastUpdateTime
}

func (s *SGpioSysfsProvider) IsHigh() bool {
	// 是否是高电平
	value := gfile.GetContents(gfile.Join(s.deviceConfig.Path, GpioPathValue))
	if gstr.Trim(value) == "1" {
		s.process(true)
		return true
	} else {
		s.process(false)
		return false
	}
}

func (s *SGpioSysfsProvider) process(status bool) {
	now := time.Now()
	s.lastUpdateTime = &now
	// 如果是反向的，就取反
	if s.deviceConfig.Reverse {
		status = !status
	}

	if s.status != status {
		s.status = status
		if s.handler != nil {
			s.handler(s.Ctx, status)
		}
	}
}

func (s *SGpioSysfsProvider) IsLow() bool {
	return !s.IsHigh()
}

func (s *SGpioSysfsProvider) SetHigh() error {
	// 设置为高电平
	return s.setValue("1")
}

func (s *SGpioSysfsProvider) SetLow() error {
	// 设置为低电平
	return s.setValue("0")
}

func (s *SGpioSysfsProvider) setValue(value string) error {
	// 设置为低电平
	if s.deviceConfig.Direction == p_gpio_sysfs.EGpioDirectionIn {
		s.log.Debugf(s.Ctx, "GPIO Direction %s is [in], can't output", s.deviceId)
		return nil
	}
	return gfile.PutContents(gfile.Join(s.deviceConfig.Path, GpioPathValue), value)
}
