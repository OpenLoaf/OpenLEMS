package internal

import (
	"context"
	"ems-plan/c_base"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"plug_protocol_gpio_sysfs/p_gpio_sysfs"
)

type SDriverGpioImpl struct {
	p_gpio_sysfs.IGpioSysfsProtocol

	ctx         context.Context
	description *c_base.SDescription
}

func NewPlugin(ctx context.Context) c_base.IDriver {
	return &SDriverGpioImpl{
		ctx: ctx,
	}
}

func (l *SDriverGpioImpl) Init(protocol c_base.IProtocol, deviceConfig *c_base.SDriverConfig) {
	if protocol == nil {
		panic(gerror.Newf("GPIO设备需要配置加载对应的协议！请检查设备：[%s]%s 的protocol相关配置！", deviceConfig.Name, deviceConfig.Id))
	}
	l.IGpioSysfsProtocol = protocol.(p_gpio_sysfs.IGpioSysfsProtocol)

	l.description = &c_base.SDescription{
		Brand:  "Basic",
		Model:  string(l.IGpioSysfsProtocol.GetGpioDeviceConfig().Direction),
		Remark: l.IGpioSysfsProtocol.GetDeviceConfig().Name,
	}
	g.Log().Infof(l.ctx, "初始化GPIO驱动[%s]成功！", l.GetDeviceConfig().Name)
}

func (l *SDriverGpioImpl) GetFunctionList() []*c_base.STelemetry {
	return []*c_base.STelemetry{
		{Name: "status", Unit: "bool", Remark: "开关量"},
	}
}

func (l *SDriverGpioImpl) GetDescription() *c_base.SDescription {
	return l.description
}

func (l *SDriverGpioImpl) GetDriverType() c_base.EDeviceType {
	return c_base.EDeviceGpio
}
