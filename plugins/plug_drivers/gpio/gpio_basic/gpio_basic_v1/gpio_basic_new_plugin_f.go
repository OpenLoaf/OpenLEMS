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
	d := &sBasicGpio{
		SRealDeviceImpl:  device.(*c_device.SRealDeviceImpl[c_proto.IGpiodProtocol]),
		GpioDeviceConfig: &c_proto.SGpioDeviceConfig{},
	}

	return d
}

func GetDriverInfo() *c_base.SDriverInfo {
	return c_base.BuildDescriptionFromYaml(buildYaml, &c_proto.SGpioDeviceConfig{})
}
