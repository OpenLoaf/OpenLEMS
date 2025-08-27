package pcs_star_charge_100E_v1

import (
	"common/c_base"
	"common/c_device"
	"common/c_proto"
	_ "embed"
)

//go:embed build.yaml
var buildYaml []byte

func NewPlugin(device c_base.IDevice) c_base.IDevice {
	return &sPcsStarCharge100E{
		SRealDeviceImpl: device.(*c_device.SRealDeviceImpl[c_proto.IModbusProtocol]),
	}
}

func GetDriverInfo() *c_base.SDriverInfo {
	return c_base.BuildDescriptionFromYaml(buildYaml)
}
