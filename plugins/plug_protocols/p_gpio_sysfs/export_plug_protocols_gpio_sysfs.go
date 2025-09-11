package p_gpio_sysfs

import (
	"common/c_base"
	"common/c_proto"
	"context"
	"p_gpio_sysfs/internal"
	"p_gpio_sysfs/p_gpio_sysfs"
)

func NewGpioSysfsProvider(ctx context.Context, protocolConfig *c_base.SProtocolConfig, deviceConfig *c_base.SDeviceConfig) (c_proto.IGpioSysfsProtocol, error) {
	return internal.NewGpioSysfsProvider(ctx, protocolConfig, deviceConfig)
}

// ScanGpioSysfs 扫描指定路径（如 /sys/class/gpio/）下的 gpiochip* 与 gpio* 信息
func ScanGpioSysfs(ctx context.Context, root string) (*p_gpio_sysfs.SGpioScanResult, error) {
	return internal.ScanGpioSysfs(ctx, root)
}
