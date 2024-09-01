package p_gpio_sysfs

import (
	"context"
	"ems-plan/c_base"
)

type SDriverGpioImpl struct {
	IGpioSysfsProtocol

	Ctx         context.Context
	description *c_base.SDescription
}

func (l *SDriverGpioImpl) Init(client c_base.IProtocol, deviceConfig *c_base.SDriverConfig) {
	l.IGpioSysfsProtocol = client.(IGpioSysfsProtocol)

	l.description = &c_base.SDescription{
		Brand:  "Basic",
		Model:  string(l.IGpioSysfsProtocol.GetGpioDeviceConfig().Direction),
		Remark: l.IGpioSysfsProtocol.GetDeviceConfig().Name,
	}
}

func (l *SDriverGpioImpl) GetFunctionList() []*c_base.SFunction {
	return []*c_base.SFunction{
		{FunctionName: "status", Unit: "bool", Remark: "开关量"},
	}
}

func (l *SDriverGpioImpl) GetDescription() *c_base.SDescription {
	return l.description
}

func (l *SDriverGpioImpl) GetDriverType() c_base.EDeviceType {
	return c_base.EGpio
}
