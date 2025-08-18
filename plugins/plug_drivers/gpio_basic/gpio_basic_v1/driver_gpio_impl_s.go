package gpio_basic_v1

import (
	"common/c_base"
	"common/c_device"
	"common/c_gpio"
	"common/c_log"
	"context"
	_ "embed"
	"fmt"
)

type sDriverGpioImpl struct {
	c_gpio.IGpioSysfsProtocol

	ctx context.Context
	*c_base.SDriverDescription
}

var _ c_device.IGpio = (*sDriverGpioImpl)(nil)

func (l *sDriverGpioImpl) InitDevice(deviceConfig *c_base.SDeviceConfig, protocol c_base.IProtocol, childDevice []c_base.IDevice) {
	if protocol == nil {
		panic(fmt.Errorf("GPIO设备需要配置加载对应的协议！请检查设备：[%s]%s 的protocol相关配置！", deviceConfig.Name, deviceConfig.Id))
	}
	l.IGpioSysfsProtocol = protocol.(c_gpio.IGpioSysfsProtocol)
	c_log.Infof(l.ctx, "初始化GPIO驱动[%s]成功！", l.GetDeviceConfig().Name)
}

func (l *sDriverGpioImpl) Shutdown() {

}

func (l *sDriverGpioImpl) GetDriverType() c_base.EDeviceType {
	return c_base.EDeviceGpio
}
