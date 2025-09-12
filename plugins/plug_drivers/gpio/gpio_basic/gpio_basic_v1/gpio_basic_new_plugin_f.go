package gpio_basic_v1

import (
	"common/c_base"
	"common/c_device"
	"common/c_proto"
	_ "embed"
)

//go:embed build.yaml
var buildYaml []byte

func NewPlugin(device c_base.IDevice) c_base.IDevice {
	return &sBasicGpio{
		SRealGpio: *device.(*c_device.SRealGpio),
	}
}

func GetDriverInfo() *c_base.SDriverInfo {
	return c_base.BuildDescriptionFromYaml(buildYaml, &c_proto.SGpioDeviceConfig{})
}
