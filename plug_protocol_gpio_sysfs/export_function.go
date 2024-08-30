package main

import (
	"context"
	"ems-plan/c_base"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
	"plug_protocol_gpio_sysfs/internal"
	"plug_protocol_gpio_sysfs/p_gpio_sysfs"
	"time"
)

func main() {

	Main := gcmd.Command{
		Name: "gpio",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			g.Log().Infof(ctx, "gpio path: %s", parser.GetOpt("path").String())

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
				Direction: p_gpio_sysfs.EGpioDirection(parser.GetOpt("direction").String()),
				//Path:          "/sys/class/gpio/gpio291",
				Path:       parser.GetOpt("path").String(),
				ExportPath: "",
				ExportPort: 0,
			}, nil)
			if err != nil {
				panic(err)
			}

			provider.Init(c_base.EGpio)

			for {
				time.Sleep(time.Second)
				_ = provider.SetHigh()
				time.Sleep(time.Second)
				_ = provider.SetLow()
			}

			return err
		},
	}

	Main.Run(context.TODO())

}
