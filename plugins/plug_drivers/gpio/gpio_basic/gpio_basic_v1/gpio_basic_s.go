package gpio_basic_v1

import (
	"common/c_base"
	"common/c_device"
	"common/c_proto"
)

type sBasicGpio struct {
	c_device.SRealGpio

	config *c_proto.SGpioDeviceConfig
}

var gpioPoint = &c_base.SPoint{
	Key:     "pin",
	Name:    "状态",
	Group:   c_base.GroupTotal,
	Precise: 0,
}

func (s *sBasicGpio) Shutdown() {

}

func (s *sBasicGpio) Init() error {

	err := s.GetConfig().ScanParams(s.config)
	if err != nil {
		return err
	}

	s.InitGpioPoint(gpioPoint)

	s.RegisterHandler(func(status bool) {

	})

	return nil
}
