package plug_protocol_gpio_sysfs

import (
	"context"
	"ems-plan/c_base"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
	"plug_protocol_gpio_sysfs/internal"
	"plug_protocol_gpio_sysfs/p_gpio_sysfs"
	"time"
)

func NewGpioSysfsProvider(ctx context.Context, protocolConfig *c_base.SProtocolConfig, deviceConfig *p_gpio_sysfs.SGpioSysfsDeviceConfig) (p_gpio_sysfs.IGpioSysfsProtocol, error) {
	return internal.NewGpioSysfsProvider(ctx, protocolConfig, deviceConfig)
}

type sMain struct {
	g.Meta `name:"main"`
}
type sInput p_gpio_sysfs.SGpioSysfsDeviceConfig

type sOutput struct {
}

func (m *sMain) Start(ctx context.Context, config sInput) (*sOutput, error) {
	g.Log().Infof(ctx, "gpio path: %s", config.Path)

	provider, err := internal.NewGpioSysfsProvider(context.TODO(), &c_base.SProtocolConfig{
		Name:           "",
		Protocol:       "",
		Address:        "",
		Timeout:        0,
		LogLevel:       "INFO",
		Config:         nil,
		Enable:         true,
		DeviceChildren: nil,
	}, &p_gpio_sysfs.SGpioSysfsDeviceConfig{
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

	provider.Init(c_base.EGpio)

	provider.RegisterHighHandler(func(ctx context.Context) {
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
