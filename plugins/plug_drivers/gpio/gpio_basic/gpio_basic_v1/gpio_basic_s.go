package gpio_basic_v1

import (
	"common/c_base"
	"common/c_device"
	"common/c_enum"
	"common/c_log"
	"common/c_proto"
	"common/c_type"

	"github.com/shockerli/cvt"
)

type sBasicGpio struct {
	*c_device.SRealDeviceImpl[c_proto.IGpiodProtocol]

	GpioDeviceConfig *c_proto.SGpioDeviceConfig
}

var gpioPoint = &c_base.SPoint{
	Key:     "pin",
	Name:    "状态",
	Group:   c_base.GroupTotal,
	Precise: 0,
	Hidden:  true,
}

var _ c_type.IGpio = (*sBasicGpio)(nil)

func (s *sBasicGpio) Shutdown() {

}

func (s *sBasicGpio) Init() error {

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

func (s *sBasicGpio) RegisterHandler(handler func(status bool, isChange bool)) {
	_ = s.ExecuteProtocolMethod(func(protocol c_proto.IGpiodProtocol) error {
		protocol.RegisterHandler(handler)
		return nil
	})
}

func (s *sBasicGpio) GetStatus() *bool {
	v, err := s.GetFromProtocolBool(func(protocol c_proto.IGpiodProtocol) (any, error) {
		return protocol.GetStatus(), nil
	})
	if err != nil {
		return nil
	}
	return v
}

func (s *sBasicGpio) SetHigh() error {
	return s.ExecuteProtocolMethod(func(protocol c_proto.IGpiodProtocol) error {
		return protocol.SetHigh()
	})
}

func (s *sBasicGpio) SetLow() error {
	return s.ExecuteProtocolMethod(func(protocol c_proto.IGpiodProtocol) error {
		return protocol.SetLow()
	})
}
