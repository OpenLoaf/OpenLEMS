package ess_pylon_checkwatt_v1

import (
	"common/c_base"
	"common/c_device"
	_ "embed"
)

//go:embed build.yaml
var buildYaml []byte

func NewPlugin(device c_base.IDevice) c_base.IDevice {
	return &sEssPylonCheckwatt{
		SVirtualDeviceImpl: device.(*c_device.SVirtualDeviceImpl),
	}
}

func GetDriverInfo() *c_base.SDriverInfo {
	return c_base.BuildDescriptionFromYaml(buildYaml)
}
