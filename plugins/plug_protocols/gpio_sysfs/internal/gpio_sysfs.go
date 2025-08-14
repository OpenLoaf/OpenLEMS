package internal

import (
	"common/c_base"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/os/gproc"
	"github.com/gogf/gf/v2/os/gtimer"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"gpio_sysfs/p_gpio_sysfs"
	"sync"
	"time"
)

type SGpioSysfsProvider struct {
	*c_base.SAlarmHandler                                                       // 告警
	once                  sync.Once                                             // 只执行一次Init方法
	meta                  *c_base.Meta                                          // 元数据
	mutex                 sync.Mutex                                            // 互斥锁，修改数据防止并发
	log                   *glog.Logger                                          // 日志
	status                bool                                                  // 结果
	lastUpdateTime        *time.Time                                            // 最后更新时间
	deviceConfig          *c_base.SDriverConfig                                 // 设备基础配置
	deviceGpioConfig      *p_gpio_sysfs.SDeviceGpioConfig                       // 设备扩展配置
	protocolConfig        *c_base.SProtocolConfig                               // 协议配置
	protocolGpioConfig    *p_gpio_sysfs.SProtocolGpioConfig                     // 协议扩展配置
	gpioDirPath           string                                                // gpio路径
	gpioValuePath         string                                                // gpio值路径
	handler               func(ctx context.Context, status bool, isChange bool) // 处理函数
}

func NewGpioSysfsProvider(ctx context.Context, protocolConfig *c_base.SProtocolConfig, deviceConfig *c_base.SDriverConfig) (p_gpio_sysfs.IGpioSysfsProtocol, error) {

	if protocolConfig == nil {
		panic(gerror.Newf("GPIO设备：[%s]%s 的协议配置不能为空！", deviceConfig.Id, deviceConfig.Name))
	}
	if protocolConfig.Address == "" {
		panic(gerror.Newf("GPIO设备：[%s]%s 的协议配置地址不能为空！", deviceConfig.Id, deviceConfig.Name))
	}
	if deviceConfig == nil {
		panic(gerror.Newf("GPIO协议：%s 的设备配置不能为空！", protocolConfig.Id))
	}
	var deviceGpioConfig = &p_gpio_sysfs.SDeviceGpioConfig{}
	err := gconv.Scan(deviceConfig.Params, deviceGpioConfig)
	if err != nil {
		panic(gerror.Newf("设备[%s]的Param参数配置错误：%v 无法转换为SGpioSysfsDeviceConfig", deviceConfig.Id, err))
	}
	if deviceGpioConfig.Io == 0 {
		// io默认禁止为0
		panic(gerror.Newf("GPIO设备[%s]的IO不能为空或者0", deviceConfig.Id))
	}

	provider := &SGpioSysfsProvider{
		once:             sync.Once{},
		deviceConfig:     deviceConfig,
		deviceGpioConfig: deviceGpioConfig,
		protocolConfig:   protocolConfig,
		gpioDirPath:      gfile.Join(protocolConfig.Address, fmt.Sprintf("gpio%d", deviceGpioConfig.Io)),
		SAlarmHandler: &c_base.SAlarmHandler{
			Ctx: ctx,
		},
		log: g.Log(deviceConfig.Id),

		meta: &c_base.Meta{
			Debug:      false,
			Name:       deviceConfig.Id,
			Cn:         deviceConfig.Name,
			Addr:       uint16(deviceGpioConfig.Io),
			BitLength:  1,
			Endianness: c_base.EBigEndian,
			ReadType:   c_base.RBit0,
			SystemType: c_base.SBool,
			Level:      deviceGpioConfig.Level,
			Factor:     1,
			Offset:     0,
			Min:        0,
			Max:        1,
			Precise:    0,
			Unit:       "",
			Desc:       fmt.Sprintf("Gpio Id: %s %s", deviceConfig.Id, deviceConfig.Name),
			Trigger:    nil,
		},
	}
	provider.gpioValuePath = gfile.Join(provider.gpioDirPath, GpioPathValue)
	if deviceConfig.LogLevel != "" {
		err := provider.log.SetLevelStr(deviceConfig.LogLevel)
		if err != nil {
			provider.log.Level(glog.LEVEL_INFO)
		}
	} else {
		// 设置默认的日志等级为上一级接口配置的日志等级
		err := provider.log.SetLevelStr(protocolConfig.GetLogLevel())
		if err != nil {
			provider.log.Level(glog.LEVEL_INFO)
		}
	}

	provider.protocolGpioConfig = &p_gpio_sysfs.SProtocolGpioConfig{}
	err = gconv.Scan(protocolConfig.Params, provider.protocolGpioConfig)
	if err != nil {
		panic(gerror.Newf("协议[%s]的Param参数配置错误：%v 无法转换为SGpioProtocolConfig", protocolConfig.Id, err))
	}

	return provider, nil
}

func (s *SGpioSysfsProvider) GetDeviceConfig() *c_base.SDriverConfig {
	return s.deviceConfig
}

func (s *SGpioSysfsProvider) GetProtocolConfig() *c_base.SProtocolConfig {
	return s.protocolConfig
}

func (s *SGpioSysfsProvider) GetMetaValueList() []*c_base.MetaValueWrapper {
	return []*c_base.MetaValueWrapper{{
		DeviceId:   s.deviceConfig.Id,
		DeviceType: c_base.EDeviceGpio,
		Meta:       s.meta,
		Value:      gvar.New(s.status),
		HappenTime: s.lastUpdateTime,
	}}
}

func (s *SGpioSysfsProvider) Init() {
	if s.deviceGpioConfig.Direction == p_gpio_sysfs.EGpioDirectionNone {
		panic(gerror.Newf("direction方向未设置！"))
	}
	s.once.Do(func() {

		// 先判断Path是否存在
		if !gfile.Exists(s.gpioDirPath) {
			// 判断一下export是否存在，如果存在，就export导出一下端口
			exportFilePath := gfile.Join(s.protocolConfig.Address, "export")
			if gfile.Exists(exportFilePath) {
				// 执行导出方法
				gproc.MustShellExec(s.Ctx, fmt.Sprintf("echo %d > %s", s.deviceGpioConfig.Io, exportFilePath))
			} else {
				panic(gerror.Newf("导出的方法不存在！%s", exportFilePath))
			}
			// 再次判断是否成功
			if !gfile.Exists(s.gpioDirPath) {
				panic(gerror.Newf("export导出后%s还是不存在！", s.gpioDirPath))
			}
		}
		if !gfile.Exists(s.gpioValuePath) {
			panic(gerror.Newf("gpio值路径不存在！%s", s.gpioValuePath))
		}

		// 设置方向
		err := gfile.PutContents(gfile.Join(s.gpioDirPath, GpioPathDirection), string(s.deviceGpioConfig.Direction))
		if err != nil {
			panic(gerror.Newf("设置方向失败！%v", err))
		}

		// 如果是In类型的，说明需要时刻监听
		if s.deviceGpioConfig.Direction == p_gpio_sysfs.EGpioDirectionIn {
			gtimer.SetInterval(s.Ctx, 200*time.Millisecond, func(ctx context.Context) {
				// 读取值
				s.isHighForce()
			})
			if err != nil {
				panic(gerror.Newf("监听失败！%v", err))
			}
		} else {
			gtimer.SetInterval(s.Ctx, 3000*time.Millisecond, func(ctx context.Context) {
				// out类型的 3秒读取一次值
				s.isHighForce()
			})
		}
	})
}

func (s *SGpioSysfsProvider) RegisterHandler(handler func(ctx context.Context, status bool, isChange bool)) {
	if s.deviceGpioConfig.Direction != p_gpio_sysfs.EGpioDirectionIn {
		panic(gerror.Newf("只有输入类型的GPIO才能注册Handler"))
	}
	s.handler = handler
}

func (s *SGpioSysfsProvider) GetId() string {
	return s.deviceConfig.Id
}

func (s *SGpioSysfsProvider) Close() {

}

func (s *SGpioSysfsProvider) IsActivate() bool {
	return true
}

func (s *SGpioSysfsProvider) GetLastUpdateTime() *time.Time {
	return s.lastUpdateTime
}

func (s *SGpioSysfsProvider) GetStatus() bool {
	// 如果缓存中的数据获取时间大于1秒或者，或者缓存中无值就force获取
	if s.lastUpdateTime == nil || time.Now().Sub(*s.lastUpdateTime) > time.Second {
		return s.isHighForce()
	}
	return s.status
}

func (s *SGpioSysfsProvider) isHighForce() bool {
	value := gfile.GetContents(s.gpioValuePath)

	if gstr.Trim(value) == "1" {
		s.process(true)
		return true
	} else {
		s.process(false)
		return false
	}
}

func (s *SGpioSysfsProvider) process(status bool) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	now := time.Now()
	s.lastUpdateTime = &now
	// 如果是反向的，就取反
	if s.deviceGpioConfig.Reverse {
		status = !status
	}

	if s.handler != nil {
		s.handler(s.Ctx, status, s.status != status)
		//g.Log().Debugf(s.ctx, "GPIO %s value: %v, func:%v", s.deviceConfig.Id, status, s.handler)
	}

	if s.status != status {
		s.status = status

		// 触发告警
		if s.deviceGpioConfig.Level != c_base.ENone && s.deviceGpioConfig.Direction == p_gpio_sysfs.EGpioDirectionIn {
			s.SAlarmHandler.TriggerAlarm(&c_base.SAlarmDetail{
				DeviceId:   s.GetId(),
				DeviceType: c_base.EDeviceGpio,
				Level:      s.deviceGpioConfig.Level,
				Meta:       s.meta,
				HappenTime: now,
				IsTrigger:  status,
				Value:      status,
			})
		}

	}
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
	if s.deviceGpioConfig.Direction == p_gpio_sysfs.EGpioDirectionIn {
		s.log.Debugf(s.Ctx, "GPIO Direction %s is [in], can't output", s.GetId())
		return nil
	}
	return gfile.PutContents(s.gpioValuePath, value)
}

func (s *SGpioSysfsProvider) GetGpioDeviceConfig() *p_gpio_sysfs.SDeviceGpioConfig {
	return s.deviceGpioConfig
}
