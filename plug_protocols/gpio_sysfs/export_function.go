package gpio_sysfs

import (
	"common/c_base"
	"context"
	"gpio_sysfs/internal"
	"gpio_sysfs/p_gpio_sysfs"
)

func NewGpioSysfsProvider(ctx context.Context, protocolConfig *c_base.SProtocolConfig, deviceConfig *c_base.SDriverConfig) (p_gpio_sysfs.IGpioSysfsProtocol, error) {
	return internal.NewGpioSysfsProvider(ctx, protocolConfig, deviceConfig)
}

/*
type sMain struct {
	g.Meta `name:"main"`
}
type sInput p_gpio_sysfs.SDeviceGpioConfig

type sOutput struct {
}

func (m *sMain) Start(ctx context.Context, config sInput) (*sOutput, error) {
	g.Log().Infof(ctx, "gpio path: %s", config.Path)

	provider, err := internal.NewGpioSysfsProvider(context.TODO(), &c_base.SProtocolConfig{
		Id:           "",
		Protocol:       "",
		Address:        "",
		Timeout:        0,
		LogLevel:       "INFO",
		Params:         nil,
		Enable:         true,
		DeviceChildren: nil,
	}, &p_gpio_sysfs.SDeviceGpioConfig{
		SDriverConfig: c_base.SDriverConfig{
			Id: "GpioTest",
		},
		Direction:  config.Direction,
		Path:       config.Path,
		ExportPath: config.ExportPath,
		ExportPort: config.ExportPort,
	})
	if err != nil {
		panic(err)
	}

	provider.Init(c_base.EDeviceGpio)

	provider.RegisterHandler(func(ctx context.Context, status bool) {
		g.Log().Infof(ctx, "high")
	})

	for {
		time.Sleep(time.Second)
		_ = provider.SetHigh()
		time.Sleep(time.Second)
		_ = provider.SetLow()
	}

}

func main() {

	cmd, err := gcmd.NewFromObject(&sMain{})
	if err != nil {
		panic(err)
	}

	cmd.Run(context.TODO())

}
*/
