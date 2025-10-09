package gpio_in_basic_v1

import (
	"common/c_base"
	"common/c_device"
	"common/c_enum"
	"common/c_log"
	"common/c_proto"
	"common/c_type"

	"github.com/shockerli/cvt"
)

type sBasicGpioIn struct {
	*c_device.SRealDeviceImpl[c_proto.IGpiodProtocol]

	GpioDeviceConfig *c_proto.SGpioDeviceConfig
}

func (s *sBasicGpioIn) SetHigh() error {
	//  调用协议层的SetHigh方法
	return s.ExecuteProtocolMethod(func(protocol c_proto.IGpiodProtocol) error {
		return protocol.SetHigh()
	})
}

func (s *sBasicGpioIn) SetLow() error {
	// 调用协议层的SetLow方法
	return s.ExecuteProtocolMethod(func(protocol c_proto.IGpiodProtocol) error {
		return protocol.SetLow()
	})
}

var gpioPoint = &c_base.SPoint{
	Key:     "status",
	Name:    "状态",
	Group:   c_base.GroupTotal,
	Precise: 0,
	Hidden:  true,
}

var _ c_type.IGpioOut = (*sBasicGpioIn)(nil)

func (s *sBasicGpioIn) Shutdown() {

}

func (s *sBasicGpioIn) Init() error {

	err := s.GetConfig().ScanParams(s.GpioDeviceConfig)
	if err != nil {
		c_log.BizErrorf(s.DeviceCtx, "Init Device GpioDeviceConfig Error: %s", err.Error())
		return err
	}

	if s.GpioDeviceConfig.Level != c_enum.EAlarmLevelNone {
		// 触发告警
		gpioPoint.Trigger = func(value interface{}) (trigger bool, level c_enum.EAlarmLevel, err error) {
			trigger, err = cvt.BoolE(value)
			if !s.GpioDeviceConfig.HighTrigger {
				trigger = !trigger
			}
			level = s.GpioDeviceConfig.Level
			return
		}
	}

	_ = s.ExecuteProtocolMethod(func(protocol c_proto.IGpiodProtocol) error {
		protocol.InitGpioPoint(gpioPoint)
		return nil
	})
	return nil
}

func (s *sBasicGpioIn) RegisterHandler(handler func(status bool, isChange bool)) {
	_ = s.ExecuteProtocolMethod(func(protocol c_proto.IGpiodProtocol) error {
		protocol.RegisterHandler(handler)
		return nil
	})
}

func (s *sBasicGpioIn) GetStatus() *bool {
	v, err := s.GetFromProtocolBool(func(protocol c_proto.IGpiodProtocol) (any, error) {
		return protocol.GetStatus(), nil
	})
	if err != nil {
		return nil
	}
	return v
}

// 实现新的IDevice接口方法
func (s *sBasicGpioIn) GetTelemetryPoints() []c_base.IPoint {
	// GPIO驱动没有遥测点位，返回空列表
	return []c_base.IPoint{}
}

func (s *sBasicGpioIn) GetProtocolPoints() []c_base.IPoint {
	// 返回GPIO协议点位
	return []c_base.IPoint{
		gpioPoint,
	}
}

func (s *sBasicGpioIn) GetConfigPoints() []*c_base.SConfigPoint {
	// 从配置结构体转换而来
	configPoints, err := c_base.BuildConfigPoints(s.GpioDeviceConfig)
	if err != nil {
		c_log.Errorf(s.DeviceCtx, "构建配置点位失败: %v", err)
		return []*c_base.SConfigPoint{}
	}
	return configPoints
}
