package internal

import (
	"context"
	"ems-plan/c_base"
	"fmt"
	"github.com/gogf/gf/v2/container/gvar"
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
	once sync.Once // 只执行一次Init方法
	//deviceId   string             // 设备名称
	deviceType c_base.EDeviceType // 设备类型
	//cache                 *gcache.Cache      // 点位缓存
	log             *glog.Logger // 日志
	printCacheValue bool         // 打印缓存值

	*c_base.SAlarmHandler            // 告警
	status                bool       // 结果
	lastUpdateTime        *time.Time // 最后更新时间

	deviceConfig  *p_gpio_sysfs.SGpioSysfsDeviceConfig
	protocolParam *p_gpio_sysfs.SGpioProtocolConfig

	handler func(ctx context.Context, status bool) // 处理函数

	meta  *c_base.Meta
	mutex sync.Mutex
}

func NewGpioSysfsProvider(ctx context.Context, protocolConfig *c_base.SProtocolConfig, deviceConfig *p_gpio_sysfs.SGpioSysfsDeviceConfig) (p_gpio_sysfs.IGpioSysfsProtocol, error) {
	provider := &SGpioSysfsProvider{
		once:            sync.Once{},
		printCacheValue: deviceConfig.PrintCacheValue,
		deviceConfig:    deviceConfig,
		protocolParam:   &p_gpio_sysfs.SGpioProtocolConfig{},
		SAlarmHandler: &c_base.SAlarmHandler{
			Ctx: ctx,
		},
		log: g.Log(deviceConfig.Id),

		meta: &c_base.Meta{
			Debug:      false,
			Name:       deviceConfig.Id,
			Cn:         deviceConfig.Name,
			Addr:       uint16(deviceConfig.ExportPort),
			BitLength:  1,
			Endianness: c_base.EBigEndian,
			ReadType:   c_base.RBit0,
			SystemType: c_base.SBool,
			Level:      deviceConfig.Level,
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

	err := gconv.Scan(protocolConfig.Config, provider.protocolParam)
	if err != nil {
		panic("modbus rtu配置文件解析失败")
	}

	return provider, nil
}

func (s *SGpioSysfsProvider) GetMetaValueList() []*c_base.MetaValueWrapper {
	return []*c_base.MetaValueWrapper{{
		DeviceId:   s.deviceConfig.GetId(),
		DeviceType: s.deviceType,
		Meta:       s.meta,
		Value:      gvar.New(s.status),
		HappenTime: s.lastUpdateTime,
	}}
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
				s.isHighForce()
			})
			if err != nil {
				panic(fmt.Errorf("监听失败！%v", err))
			}
		}
		s.log.Infof(s.Ctx, "GPIO %s 初始化完毕,当前状态为: %v, 类型为: %s", s.deviceConfig.GetName(), s.GetStatus(), s.deviceConfig.Direction)
	})
}

func (s *SGpioSysfsProvider) RegisterHandler(handler func(ctx context.Context, status bool)) {
	if s.deviceConfig.Direction != p_gpio_sysfs.EGpioDirectionIn {
		panic(fmt.Errorf("只有输入类型的GPIO才能注册Handler"))
	}
	s.handler = handler
}

func (s *SGpioSysfsProvider) GetId() string {
	return s.deviceConfig.Id
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

func (s *SGpioSysfsProvider) GetStatus() bool {
	// 如果缓存中的数据获取时间大于1秒或者，或者缓存中无值就force获取
	if s.lastUpdateTime == nil || time.Now().Sub(*s.lastUpdateTime) > time.Second {
		return s.isHighForce()
	}
	return s.status
}

func (s *SGpioSysfsProvider) isHighForce() bool {
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
	s.mutex.Lock()
	defer s.mutex.Unlock()

	now := time.Now()
	s.lastUpdateTime = &now
	// 如果是反向的，就取反
	if s.deviceConfig.Reverse {
		status = !status
	}

	if s.status != status {
		if s.GetId() == "button-scram" {
			g.Log().Debugf(s.Ctx, "GPIO %s value: %v, func:%v", s.deviceConfig.Id, status, s.handler)
		}
		s.status = status
		if s.handler != nil {
			s.handler(s.Ctx, status)
			g.Log().Debugf(s.Ctx, "22GPIO %s value: %v, func:%v", s.deviceConfig.Id, status, s.handler)
		}

		// 触发告警
		if s.deviceConfig.Level != c_base.ENone && s.deviceConfig.Direction == p_gpio_sysfs.EGpioDirectionIn {
			s.SAlarmHandler.ProcessAlarmDetail(&c_base.SAlarmDetail{
				DeviceId:   s.GetId(),
				DeviceType: s.deviceType,
				Level:      s.deviceConfig.Level,
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
	if s.deviceConfig.Direction == p_gpio_sysfs.EGpioDirectionIn {
		s.log.Debugf(s.Ctx, "GPIO Direction %s is [in], can't output", s.GetId())
		return nil
	}
	return gfile.PutContents(gfile.Join(s.deviceConfig.Path, GpioPathValue), value)
}
