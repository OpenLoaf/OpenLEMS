package gpio_basic_v1

import (
	"common/c_base"
	"context"
	_ "embed"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"gpio_sysfs/p_gpio_sysfs"
)

type sDriverGpioImpl struct {
	p_gpio_sysfs.IGpioSysfsProtocol

	ctx context.Context
	*c_base.SDescription
}

func (l *sDriverGpioImpl) Init(protocol c_base.IProtocol, deviceConfig *c_base.SDriverConfig) {
	if protocol == nil {
		panic(gerror.Newf("GPIO设备需要配置加载对应的协议！请检查设备：[%s]%s 的protocol相关配置！", deviceConfig.Name, deviceConfig.Id))
	}
	l.IGpioSysfsProtocol = protocol.(p_gpio_sysfs.IGpioSysfsProtocol)
	g.Log().Infof(l.ctx, "初始化GPIO驱动[%s]成功！", l.GetDeviceConfig().Name)
}

func (l *sDriverGpioImpl) Destroy() {

}

func (l *sDriverGpioImpl) GetDriverType() c_base.EDeviceType {
	return c_base.EDeviceGpio
}
