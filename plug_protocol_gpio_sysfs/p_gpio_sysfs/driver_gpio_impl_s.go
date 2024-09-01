package p_gpio_sysfs

import (
	"context"
	"ems-plan/c_base"
	"github.com/gogf/gf/v2/frame/g"
)

type SDriverGpioImpl struct {
	*SGpioSysfsDeviceConfig
	IGpioSysfsProtocol
	Ctx         context.Context
	description c_base.SDescription
}

func (l *SDriverGpioImpl) GetDescription() c_base.SDescription {
	return l.description
}

func (l *SDriverGpioImpl) Init(client c_base.IProtocol, cfg any) error {
	l.IGpioSysfsProtocol = client.(IGpioSysfsProtocol)
	l.SGpioSysfsDeviceConfig = c_base.ConvertConfig[*SGpioSysfsDeviceConfig](cfg)

	g.Log().Infof(l.Ctx, "配置信息:%+v", l.SGpioSysfsDeviceConfig)

	l.description = c_base.SDescription{
		Brand:  "Basic",
		Model:  string(l.SGpioSysfsDeviceConfig.Direction),
		Type:   c_base.EGpio,
		Remark: l.SGpioSysfsDeviceConfig.Name,
	}

	return nil
}

func (l *SDriverGpioImpl) GetType() c_base.EDeviceType {
	return c_base.EGpio
}

func (l *SDriverGpioImpl) GetFunctionList() []*c_base.SFunction {
	return []*c_base.SFunction{
		{FunctionName: "status", Unit: "bool", Remark: "开关量"},
	}
}
