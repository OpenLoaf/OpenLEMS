//go:generate ../build.sh
package main

import (
	"common/c_base"
	"context"
	_ "embed"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gtimer"
	"pcs_lnxall/pcs_lnxall_v1"
	"services"
	"time"
)

// 通过构建脚本自动注入
var (
	buildTime  string
	commitHash string
)

// NewPlugin 必须的方法，不能取消。修改实现只需修改此方法
func NewPlugin(ctx context.Context) c_base.IDevice {
	plugin := pcs_lnxall_v1.NewPlugin(ctx)

	plugin.GetDriverDescription().BuildTime = buildTime
	plugin.GetDriverDescription().CommitHash = commitHash
	return plugin
}

func main() {
	command := c_base.PluginDriverCommand(func() c_base.IDevice {
		return NewPlugin(context.Background())
	})

	// 此处可添加自定义命令
	_ = command.AddCommand(&gcmd.Command{
		Name:  "test",
		Usage: "test",
		Brief: "测试启动",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			cmd := services.NewDeviceCmd(ctx)

			device := cmd.InitDriver(make(map[string]any), &c_base.SDeviceConfig{
				Id:            "pcs",
				Name:          "协能PCS",
				ProtocolId:    "test_modbus",
				StorageEnable: false,
				Driver:        "pcs_lnxall_v1",
				Enabled:       true,
				LogLevel:      "",
				Params: map[string]string{
					"unitId": "2",
				},
				DeviceChildren: nil,
			}, []*c_base.SProtocolConfig{{
				Id:       "test_modbus",
				Type:     "modbus_tcp",
				Address:  "127.0.0.1:2509",
				Timeout:  30,
				LogLevel: "DEBUG",
				Params:   nil,
			}})

			gtimer.SetInterval(ctx, 1*time.Second, func(ctx context.Context) {
				g.Log().Noticef(ctx, "设备[%s]数据:\n%v", device.GetDeviceConfig().Name, device.GetAllTelemetry(device))
			})

			<-ctx.Done()
			return err
		},
	})
	command.Run(context.Background())
}
