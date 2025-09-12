package gpio_basic_v1

import (
	"common/c_base"
	"common/c_device"
)

type sBasicGpio struct {
	c_device.SRealGpio
}

func (s *sBasicGpio) GetConfig() *SDeviceConfig {
	//TODO implement me
	panic("implement me")
}

func (s *sBasicGpio) GetConfig() *SDeviceConfig {
	//TODO implement me
	panic("implement me")
}

func (s *sBasicGpio) Init() error {

	return nil
}

func (s *sBasicGpio) Shutdown() {
	//TODO implement me
	panic("implement me")
}

func (s *sBasicGpio) GetPointValueList() []*c_base.SPointValue {
	return nil
}
