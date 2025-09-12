package gpio_basic_v1

import (
	"common/c_device"
)

type sBasicGpio struct {
	c_device.SRealGpio
}

func (s *sBasicGpio) Shutdown() {

}

func (s *sBasicGpio) Init() error {

	return nil
}
