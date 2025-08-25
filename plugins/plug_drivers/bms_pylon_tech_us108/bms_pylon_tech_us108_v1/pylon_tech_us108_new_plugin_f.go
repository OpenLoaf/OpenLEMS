package bms_pylon_tech_us108_v1

import (
	"common/c_base"
	"common/c_device"
	"common/c_proto"
	_ "embed"
)

//go:embed build.yaml
var buildYaml []byte

func NewPlugin(device c_base.IDevice) c_base.IDevice {
	return &sBmsPylonTechUs108{
		SRealDevice: device.(*c_device.SRealDevice[c_proto.IModbusProtocol]),
	}
}

func GetDriverInfo() *c_base.SDriverInfo {
	return c_base.BuildDescriptionFromYaml(buildYaml)
}
