package gpio_out_basic_v1

import (
	"common/c_base"
	"common/c_device"
	"common/c_log"
	"common/c_proto"
	"common/c_type"
)

type sBasicGpioOut struct {
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

var _ c_type.IGpioOut = (*sBasicGpioOut)(nil)

func (s *sBasicGpioOut) Shutdown() {

}

func (s *sBasicGpioOut) Init() error {

	err := s.GetConfig().ScanParams(s.GpioDeviceConfig)
	if err != nil {
		c_log.BizErrorf(s.DeviceCtx, "Init Device GpioDeviceConfig Error: %s", err.Error())
		return err
	}

	_ = s.ExecuteProtocolMethod(func(protocol c_proto.IGpiodProtocol) error {
		protocol.InitGpioPoint(gpioPoint)
		return nil
	})
	return nil
}

func (s *sBasicGpioOut) SetHigh() error {
	return s.ExecuteProtocolMethod(func(protocol c_proto.IGpiodProtocol) error {
		return protocol.SetHigh()
	})
}

func (s *sBasicGpioOut) SetLow() error {
	return s.ExecuteProtocolMethod(func(protocol c_proto.IGpiodProtocol) error {
		return protocol.SetLow()
	})
}

func (s *sBasicGpioOut) RegisterHandler(handler func(status bool, isChange bool)) {
	_ = s.ExecuteProtocolMethod(func(protocol c_proto.IGpiodProtocol) error {
		protocol.RegisterHandler(handler)
		return nil
	})
}

func (s *sBasicGpioOut) GetStatus() *bool {
	v, err := s.GetFromProtocolBool(func(protocol c_proto.IGpiodProtocol) (any, error) {
		return protocol.GetStatus(), nil
	})
	if err != nil {
		return nil
	}
	return v
}
